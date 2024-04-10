package auth

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"service-auth/system"
	"time"

	"github.com/stripe/stripe-go/v76"
	portal_session "github.com/stripe/stripe-go/v76/billingportal/session"
	checkout_session "github.com/stripe/stripe-go/v76/checkout/session"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "service-auth/proto"
)

type AuthService interface {
	// Helper
	GetUser(ctx context.Context) (*pb.User, error)
	// gRPC
	Auth(ctx context.Context) (*pb.AuthResponse, error)
	CreateStripeCheckout(ctx context.Context) (*pb.StripeUrlResponse, error)
	CreateStripePortal(ctx context.Context) (*pb.StripeUrlResponse, error)
	// HTTP
	OauthLogin(http.ResponseWriter, *http.Request)
	OauthCallback(http.ResponseWriter, *http.Request)
	// Task
	CleanTokens(ctx context.Context) error
}

type AuthServiceImpl struct {
	AuthDB
}

func NewAuthService(authDB AuthDB) AuthService {
	return &AuthServiceImpl{authDB}
}

func (a *AuthServiceImpl) GetUser(ctx context.Context) (*pb.User, error) {
	claims, err := system.ExtractToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("extractToken: %w", err)
	}
	user, err := a.selectUserByTokenId(claims.Id)
	if err != nil {
		return nil, fmt.Errorf("selectUserByTokenId: %w", err)
	}
	subscribed := checkIfSubscribed(user, a.AuthDB)
	user.SubscriptionActive = subscribed
	return user, nil
}

/**
 * 1. Extract phantom token from context
 * 2. Using it's id, get oauth token from database
 * 3. Check if oauth token is valid
 * 4. Refresh oauth token if it's expired
 * 5. Create new phantom token
 * 6. Get user from database
 * 7. Return user and new phantom token
 */
func (a *AuthServiceImpl) Auth(ctx context.Context) (*pb.AuthResponse, error) {
	defer system.Perf("auth", time.Now())
	claims, err := system.ExtractToken(ctx)
	if err != nil {
		slog.Error("Error extracting token", "system.ExtractToken", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	// get token
	token, err := a.selectTokenById(claims.Id)
	if err != nil {
		slog.Error("Error selecting token by id", "authDB.selectTokenById", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	expires, err := time.Parse(time.RFC3339, token.Expires)
	if err != nil || time.Now().After(expires) {
		slog.Error("Token expired", "token.Expires", token.Expires)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	// get user from database
	user, err := a.selectUserById(token.UserId)
	if err != nil {
		slog.Error("Error selecting user by id", "authDB.selectUserById", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	// Update token expiration and user
	go func() {
		err = a.updateToken(claims.Id, time.Now().Add(7*24*time.Hour).Format(time.RFC3339))
		if err != nil {
			slog.Error("Error updating token", "authDB.updateToken", err)
		}
		err = a.updateUser(user.Id)
		if err != nil {
			slog.Error("Error updating user", "authDB.updateUser", err)
		}
	}()

	subscribed := checkIfSubscribed(user, a.AuthDB)
	user.SubscriptionActive = subscribed
	return &pb.AuthResponse{
		Token: claims.Id,
		User:  user,
	}, nil
}

func (a *AuthServiceImpl) CreateStripeCheckout(ctx context.Context) (*pb.StripeUrlResponse, error) {
	defer system.Perf("create_stripe_checkout", time.Now())
	user, err := a.GetUser(ctx)
	if err != nil {
		slog.Error("Error authorizing user", "auth.GetUser", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	customerId := user.SubscriptionId
	if customerId == "" {
		var err error
		customerId, err = createStripeUser(user.Id, user.Email, a.AuthDB)
		if err != nil {
			slog.Error("Error creating stripe user", "createStripeUser", err)
			return nil, status.Error(codes.Internal, "Error creating stripe user")
		}
	}

	stripe.Key = system.STRIPE_API_KEY

	params := &stripe.CheckoutSessionParams{
		SuccessURL: stripe.String(system.CLIENT_URL + "/billing?success"),
		CancelURL:  stripe.String(system.CLIENT_URL + "/billing?cancel"),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		Mode: stripe.String("subscription"),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(system.STRIPE_PRICE_ID),
				Quantity: stripe.Int64(1),
			},
		},
		Customer: stripe.String(customerId),
	}

	session, err := checkout_session.New(params)
	if err != nil {
		slog.Error("Error creating stripe checkout", "checkout_session.New", err)
		return nil, status.Error(codes.Internal, "Error creating stripe checkout")
	}

	err = a.updateSubscriptionCheck(user.Id, "1970-01-01T00:00:00Z")
	if err != nil {
		slog.Error("Error updating subscription check date", "updateSubscriptionCheck", err)
        return nil, status.Error(codes.Internal, "Error updating subscription check date")
	}
	return &pb.StripeUrlResponse{
		Url: session.URL,
	}, nil
}

func (a *AuthServiceImpl) CreateStripePortal(ctx context.Context) (*pb.StripeUrlResponse, error) {
	defer system.Perf("create_stripe_portal", time.Now())
	user, err := a.GetUser(ctx)
	if err != nil {
		slog.Error("Error authorizing user", "auth.GetUser", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	stripe.Key = system.STRIPE_API_KEY

	params := &stripe.BillingPortalSessionParams{
		Customer:  stripe.String(user.SubscriptionId),
		ReturnURL: stripe.String(system.CLIENT_URL + "/billing"),
	}
	session, err := portal_session.New(params)
	if err != nil {
		slog.Error("Error creating stripe portal", "portal_session.New", err)
		return nil, status.Error(codes.Internal, "Error creating stripe portal")
	}

	return &pb.StripeUrlResponse{
		Url: session.URL,
	}, nil
}

func (a *AuthServiceImpl) OauthLogin(w http.ResponseWriter, r *http.Request) {
	defer system.Perf("oauth_login", time.Now())
	provider := r.PathValue("provider")
	var OAuth OAuth
	if provider == "google" {
		OAuth = newOAuthGoogle()
	} else if provider == "github" {
		OAuth = newOAuthGithub()
	} else {
		slog.Error("Invalid provider", "provider", provider)
		http.Redirect(w, r, system.CLIENT_URL+"/auth?error=unauthorized", http.StatusTemporaryRedirect)
        return
	}

	// generate random state and verifier
	state, err := system.GenerateRandomState(32)
    if err != nil {
        slog.Error("Error generating random state", "GenerateRandomState", err)
		http.Redirect(w, r, system.CLIENT_URL+"/auth?error=unauthorized", http.StatusTemporaryRedirect)
		return
    }
	verifier := oauth2.GenerateVerifier()
	// store state and verifier
	_, err = a.insertToken(time.Now().Add(10*time.Second).Format(time.RFC3339), "", state, verifier)
	if err != nil {
		slog.Error("Error inserting token", "insertToken", err)
		http.Redirect(w, r, system.CLIENT_URL+"/auth?error=unauthorized", http.StatusTemporaryRedirect)
		return
	}
	// Redirect user to consent page to ask for permission
	url := OAuth.getOAuthConfig().AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (a *AuthServiceImpl) OauthCallback(w http.ResponseWriter, r *http.Request) {
	defer system.Perf("oauth_callback", time.Now())

	provider := r.PathValue("provider")
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")
	var OAuth OAuth
	if provider == "google" {
		OAuth = newOAuthGoogle()
	} else if provider == "github" {
		OAuth = newOAuthGithub()
	} else {
		slog.Error("Invalid provider", "provider", provider)
		http.Redirect(w, r, system.CLIENT_URL+"/auth?error=unauthorized", http.StatusTemporaryRedirect)
		return
	}

	// get verifier from state
	token, err := a.seleteTokenByState(state)
	if err != nil {
		slog.Error("Error getting token by state", "seleteTokenByState", err)
		http.Redirect(w, r, system.CLIENT_URL+"/auth?error=unauthorized", http.StatusTemporaryRedirect)
		return
	}
	expires, err := time.Parse(time.RFC3339, token.Expires)
	if err != nil || time.Now().After(expires) {
		slog.Error("Token expired", "token.Expires", token.Expires)
		http.Redirect(w, r, system.CLIENT_URL+"/auth?error=unauthorized", http.StatusTemporaryRedirect)
		return
	}

	// get oauth config
	config := OAuth.getOAuthConfig()

	// get oauth token
	oauthToken, err := config.Exchange(context.Background(), code, oauth2.VerifierOption(token.Verifier))
	if err != nil {
		slog.Error("Error exchanging code for token", "config.Exchange", err)
        http.Redirect(w, r, system.CLIENT_URL+"/auth?error=unauthorized", http.StatusTemporaryRedirect)
        return
	}

	// fetch user info from github
	userInfo, err := OAuth.getUserInfo(oauthToken.AccessToken)
	if err != nil {
		slog.Error("Error fetching user info", "configProvider.getUserInfo", err)
        http.Redirect(w, r, system.CLIENT_URL+"/auth?error=unauthorized", http.StatusTemporaryRedirect)
        return
	}

	// get user, create if not exists
	user, err := a.selectUserByEmailAndSub(userInfo.email, userInfo.sub)
	if err != nil {
		user, err = a.insertUser(userInfo.email, userInfo.sub, userInfo.avatar)
		if err != nil {
			slog.Error("Error inserting user", "insertUser", err)
			http.Redirect(w, r, system.CLIENT_URL+"/auth?error=invalid_user", http.StatusTemporaryRedirect)
            return
		}
	}

	// create auth token with a 10 seconds expiration
	token, err = a.insertToken(time.Now().Add(10*time.Second).Format(time.RFC3339), user.Id, "", "")
	if err != nil {
		slog.Error("Error inserting token", "insertToken", err)
        http.Redirect(w, r, system.CLIENT_URL+"/auth?error=unauthorized", http.StatusTemporaryRedirect)
        return
	}

	// redirect to home page
	http.Redirect(w, r, system.CLIENT_URL+"/?token="+token.Id, http.StatusTemporaryRedirect)
}

func (a *AuthServiceImpl) CleanTokens(ctx context.Context) error {
	err := a.AuthDB.CleanTokens()
	if err != nil {
		return fmt.Errorf("cleanTokens: %w", err)
	}
	return nil
}
