package service

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/stripe/stripe-go/v76"
	portal_session "github.com/stripe/stripe-go/v76/billingportal/session"
	checkout_session "github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/customer"
	"github.com/stripe/stripe-go/v76/subscription"

	pb "service-auth/proto"
	"service-auth/store"
	"service-auth/system"
)

type StripeStorage interface {
	UpdateSubscriptionId(ctx context.Context, userId string, subscriptionId string) error
	UpdateSubscriptionCheck(ctx context.Context, userId string, subscriptionCheck string) error
	UpdateSubscriptionEnd(ctx context.Context, userId string, subscriptionEnd string) error
}

func getUser(ctx context.Context, storage system.Storage) (*pb.User, error) {
	claims, err := system.ExtractToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("extractToken: %w", err)
	}

	var store AuthStorage = store.NewAuthDB(&storage)
	user, err := store.SelectUserById(ctx, claims.Id)
	if err != nil {
		return nil, fmt.Errorf("selectUserById: %w", err)
	}
	subscribed := checkIfSubscribed(ctx, storage, user)
	user.SubscriptionActive = subscribed
	return user, nil
}

func CreateStripeCheckout(
	ctx context.Context,
	storage system.Storage,
) (*pb.StripeUrlResponse, error) {
	defer system.Perf("create_stripe_checkout", time.Now())
	user, err := getUser(ctx, storage)
	if err != nil {
		slog.Error("Error authorizing user", "auth.GetUser", err)
		return nil, err
	}

	customerId := user.SubscriptionId
	if customerId == "" {
		var err error
		customerId, err = createStripeUser(
			ctx,
			storage,
			user.Id,
			user.Email,
		)
		if err != nil {
			slog.Error("Error creating stripe user", "createStripeUser", err)
			return nil, err
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
		return nil, err
	}

	var store StripeStorage = store.NewAuthDB(&storage)

	err = store.UpdateSubscriptionCheck(ctx, user.Id, "1970-01-01T00:00:00Z")
	if err != nil {
		slog.Error("Error updating subscription check date", "updateSubscriptionCheck", err)
		return nil, err
	}
	return &pb.StripeUrlResponse{
		Url: session.URL,
	}, nil
}

func CreateStripePortal(
	ctx context.Context,
	storage system.Storage,
) (*pb.StripeUrlResponse, error) {
	defer system.Perf("create_stripe_portal", time.Now())
	user, err := getUser(ctx, storage)
	if err != nil {
		slog.Error("Error authorizing user", "auth.GetUser", err)
		return nil, err
	}
	stripe.Key = system.STRIPE_API_KEY

	params := &stripe.BillingPortalSessionParams{
		Customer:  stripe.String(user.SubscriptionId),
		ReturnURL: stripe.String(system.CLIENT_URL + "/billing"),
	}
	session, err := portal_session.New(params)
	if err != nil {
		slog.Error("Error creating stripe portal", "portal_session.New", err)
		return nil, err
	}

	return &pb.StripeUrlResponse{
		Url: session.URL,
	}, nil
}

func checkIfSubscribed(
	ctx context.Context,
	storage system.Storage,
	user *pb.User,
) bool {
	defer system.Perf("check_if_subscribed", time.Now())
	if user.SubscriptionId == "" {
		return false
	}
	subEnd, _ := time.Parse(time.RFC3339, user.SubscriptionEnd)
	if subEnd.After(time.Now()) {
		return true
	}
	subCheck, _ := time.Parse(time.RFC3339, user.SubscriptionCheck)
	if subCheck.After(time.Now()) {
		return false
	}

	var store StripeStorage = store.NewAuthDB(&storage)

	stripe.Key = system.STRIPE_API_KEY

	params := &stripe.SubscriptionListParams{
		Customer: stripe.String(user.SubscriptionId),
		Status:   stripe.String("active"),
	}

	i := subscription.List(params)
	for i.Next() {
		if i.Subscription().Status == "active" {
			// get the subscription end date
			subEnd := time.Unix(i.Subscription().CurrentPeriodEnd, 0).Format(time.RFC3339)
			// update the user's subscription end date
			err := store.UpdateSubscriptionEnd(ctx, user.Id, subEnd)
			if err != nil {
				slog.Error("Error updating subscription end date", "updateSubscriptionEnd", err)
				return false
			}
			return true
		}
	}
	err := store.UpdateSubscriptionCheck(ctx, user.Id, time.Now().Add(time.Hour).Format(time.RFC3339))
	if err != nil {
		slog.Error("Error updating subscription check date", "updateSubscriptionCheck", err)
	}
	return false
}

func createStripeUser(
	ctx context.Context,
	storage system.Storage,
	userId string,
	email string,
) (string, error) {
	var store StripeStorage = store.NewAuthDB(&storage)
	stripe.Key = system.STRIPE_API_KEY
	params := &stripe.CustomerParams{
		Email: stripe.String(email),
	}
	customer, err := customer.New(params)
	if err != nil {
		return "", fmt.Errorf("customer.New: %w", err)
	}
	err = store.UpdateSubscriptionId(ctx, userId, customer.ID)
	if err != nil {
		return "", fmt.Errorf("updateSubscriptionId: %w", err)
	}
	return customer.ID, nil
}

// TODO: not used
// func oauthRefresh(token Token, configProvider OAuthConfigProvider) (*oauth2.Token, error) {
// 	oauthToken := oauth2.Token{
// 		AccessToken:  token.AccessToken,
// 		RefreshToken: token.RefreshToken,
// 		TokenType:    token.TokenType,
// 		Expiry:       token.Expires,
// 	}
//
// 	config, err := configProvider.getOAuthConfig(token.Provider)
// 	if err != nil {
// 		return nil, fmt.Errorf("getOAuthConfig: %w", err)
// 	}
//
// 	newOauthToken, err := config.TokenSource(context.Background(), &oauthToken).Token()
// 	if err != nil {
// 		return nil, fmt.Errorf("config.TokenSource: %w", err)
// 	}
// 	return newOauthToken, nil
// }
