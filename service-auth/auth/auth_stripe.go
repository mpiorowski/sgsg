package auth

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/customer"
	"github.com/stripe/stripe-go/v76/subscription"

	pb "service-auth/proto"
	"service-auth/system"
)

func checkIfSubscribed(user *pb.User, authDB AuthDB) bool {
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
			err := authDB.updateSubscriptionEnd(user.Id, subEnd)
			if err != nil {
				slog.Error("Error updating subscription end date", "updateSubscriptionEnd", err)
				return false
			}
			return true
		}
	}
	err := authDB.updateSubscriptionCheck(user.Id, time.Now().Add(time.Hour).Format(time.RFC3339))
	if err != nil {
		slog.Error("Error updating subscription check date", "updateSubscriptionCheck", err)
	}
	return false
}

func createStripeUser(userId string, email string, authDB AuthDB) (string, error) {
	stripe.Key = system.STRIPE_API_KEY

	params := &stripe.CustomerParams{
		Email: stripe.String(email),
	}
	customer, err := customer.New(params)
	if err != nil {
		return "", fmt.Errorf("customer.New: %w", err)
	}
	err = authDB.updateSubscriptionId(userId, customer.ID)
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
