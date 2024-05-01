package service

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"service-auth/store"
	"service-auth/system"
	"strings"
	"time"

	"golang.org/x/oauth2"

	pb "service-auth/proto"
)

type AuthStorage interface {
	// Token
	SelectTokenById(ctx context.Context, id string) (*store.Token, error)
	SeleteTokenByState(ctx context.Context, state string) (*store.Token, error)
	InsertToken(ctx context.Context, expires string, userId string, state string, verifier string) (*store.Token, error)
	UpdateToken(ctx context.Context, id string, expires string) error
	// User
	SelectUserById(ctx context.Context, id string) (*pb.User, error)
	SelectUserByEmail(ctx context.Context, email string) (*pb.User, error)
	InsertUser(ctx context.Context, email string, sub string, avatar string) (*pb.User, error)
	UpdateUserSub(ctx context.Context, id string, sub string) error
	UpdateUserActivity(ctx context.Context, id string) error
}

type OAuth interface {
	getOAuthConfig() *oauth2.Config
	getUserInfo(accessToken string) (*UserInfo, error)
}


// 1. Extract phantom token from context
// 2. Using it's id, get oauth token from database
// 3. Check if oauth token is valid
// 4. Refresh oauth token if it's expired
// 5. Create new phantom token
// 6. Get user from database
// 7. Return user and new phantom token
func Auth(
	ctx context.Context,
	storage system.Storage,
) (*pb.AuthResponse, error) {
	defer system.Perf("auth", time.Now())
	claims, err := system.ExtractToken(ctx)
	if err != nil {
		slog.Error("Error extracting token", "system.ExtractToken", err)
		return nil, err
	}

	var store AuthStorage = store.NewAuthDB(&storage)
	// get token
	token, err := store.SelectTokenById(ctx, claims.Id)
	if err != nil {
		slog.Error("Error selecting token by id", "authDB.selectTokenById", err)
		return nil, err
	}
	expires, err := time.Parse(time.RFC3339, token.Expires)
	if err != nil || time.Now().After(expires) {
		slog.Error("Token expired", "token.Expires", token.Expires)
		return nil, err
	}

	// get user from database
	user, err := store.SelectUserById(ctx, token.UserId)
	if err != nil {
		slog.Error("Error selecting user by id", "authDB.selectUserById", err)
		return nil, err
	}

	// Update token expiration and user
	go func() {
		var ctx context.Context = context.Background()
		err = store.UpdateToken(ctx, claims.Id, time.Now().Add(7*24*time.Hour).Format(time.RFC3339))
		if err != nil {
			slog.Error("Error updating token", "authDB.updateToken", err)
		}
		err = store.UpdateUserActivity(ctx, user.Id)
		if err != nil {
			slog.Error("Error updating user", "authDB.updateUser", err)
		}
	}()

	subscribed := checkIfSubscribed(ctx, storage, user)
	user.SubscriptionActive = subscribed
	return &pb.AuthResponse{
		Token: claims.Id,
		User:  user,
	}, nil
}

func OauthLogin(
	storage system.Storage,
	w http.ResponseWriter,
	r *http.Request,
) {
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
	var store AuthStorage = store.NewAuthDB(&storage)
	_, err = store.InsertToken(r.Context(), time.Now().Add(10*time.Second).Format(time.RFC3339), "", state, verifier)
	if err != nil {
		slog.Error("Error inserting token", "insertToken", err)
		http.Redirect(w, r, system.CLIENT_URL+"/auth?error=unauthorized", http.StatusTemporaryRedirect)
		return
	}
	// Redirect user to consent page to ask for permission
	url := OAuth.getOAuthConfig().AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func OauthCallback(
	storage system.Storage,
	w http.ResponseWriter,
	r *http.Request,
) {
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
	var store AuthStorage = store.NewAuthDB(&storage)
	token, err := store.SeleteTokenByState(r.Context(), state)
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
	user, err := store.SelectUserByEmail(r.Context(), userInfo.email)
	if err != nil {
		// create new user if not found
		sub := provider + ":" + userInfo.sub
		user, err = store.InsertUser(r.Context(), userInfo.email, sub, userInfo.avatar)
		if err != nil {
			slog.Error("Error inserting user", "insertUser", err)
			http.Redirect(w, r, system.CLIENT_URL+"/auth?error=unauthorized", http.StatusTemporaryRedirect)
			return
		}
	} else {
		// validate provider
		err = validateProvider(r.Context(), user.Id, user.Sub, provider, userInfo.sub, store)
		if err != nil {
			slog.Error("Error validating provider", "validateProvider", err)
			http.Redirect(w, r, system.CLIENT_URL+"/auth?error=unauthorized", http.StatusTemporaryRedirect)
			return
		}
	}

	// create auth token with a 10 seconds expiration
	token, err = store.InsertToken(r.Context(), time.Now().Add(10*time.Second).Format(time.RFC3339), user.Id, "", "")
	if err != nil {
		slog.Error("Error inserting token", "insertToken", err)
		http.Redirect(w, r, system.CLIENT_URL+"/auth?error=unauthorized", http.StatusTemporaryRedirect)
		return
	}

	// redirect to home page
	http.Redirect(w, r, system.CLIENT_URL+"/?token="+token.Id, http.StatusTemporaryRedirect)
}

// User can have multiple providers, in form of "provider:sub"
// User cannot have multiple providers of the same type
// If user don't have provider, add it
func validateProvider(ctx context.Context, userId string, userSubs string, provider string, sub string, store AuthStorage) error {
	var err error
	// if doesn't have provider, add it
	if !strings.Contains(userSubs, provider) {
		err = store.UpdateUserSub(ctx, userId, userSubs+","+provider+":"+sub)
		if err != nil {
			return fmt.Errorf("updateUserSub: %w", err)
		}
		return nil
	}

	// if have provider, check if it's the same
	subs := strings.Split(userSubs, ",")
	for _, s := range subs {
		if strings.Contains(s, provider) {
			if s != provider+":"+sub {
				return fmt.Errorf("provider already exists")
			}
		}
	}
	return nil
}
