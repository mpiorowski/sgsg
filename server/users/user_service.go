package users

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	pb "sgsg/proto"
	"sgsg/utils"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

// TODO - use redis instead of in-memory storage
var storage = utils.NewStorage()

/**
 *  Oauth login
 *  @api {get} /oauth-login/:provider Oauth login
 */
func OauthLogin(c echo.Context, configProvider OAuthConfigProvider) error {
	provider := c.Param("provider")
	// generate random state
	state, err := utils.GenerateRandomState(32)
	if err != nil {
		slog.Error("Error generating random state", "utils.GenerateRandomState", err)
		return c.Redirect(http.StatusTemporaryRedirect, utils.CLIENT_URL+"/auth?error=unauthorized")
	}
	storage.Add(state)

	config, err := configProvider.getOAuthConfig(provider)
	if err != nil {
		slog.Error("Error getting provider", "configProvider.getOAuthConfig", err)
		return c.Redirect(http.StatusTemporaryRedirect, utils.CLIENT_URL+"/auth?error=unauthorized")
	}

	// Redirect user to GitHub login page
	url := config.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

/**
 *  Oauth callback
 *  @api {get} /oauth-callback/:provider Oauth callback
 */
func OauthCallback(c echo.Context, configProvider OAuthConfigProvider) error {
	provider := c.Param("provider")
	code := c.QueryParam("code")
	state := c.QueryParam("state")

	// check state
	if !storage.Check(state) {
		slog.Error("Invalid state", "storage.Check", state)
		return c.Redirect(http.StatusTemporaryRedirect, utils.CLIENT_URL+"/auth?error=unauthorized")
	}
	storage.Remove(state)

	// get oauth config
	config, err := configProvider.getOAuthConfig(provider)
	if err != nil {
		slog.Error("Error getting provider", "configProvider.getOAuthConfig", err)
		return c.Redirect(http.StatusTemporaryRedirect, utils.CLIENT_URL+"/auth?error=unauthorized")
	}

	// get oauth token
	oauthToken, err := config.Exchange(context.Background(), code)
	if err != nil {
		slog.Error("Error exchanging code for token", "config.Exchange", err)
		return c.Redirect(http.StatusTemporaryRedirect, utils.CLIENT_URL+"/auth?error=unauthorized")
	}

	// fetch user info from github
	userInfo, err := configProvider.getUserInfo(provider, oauthToken.AccessToken)
	if err != nil {
		slog.Error("Error fetching user info", "configProvider.getUserInfo", err)
		return c.Redirect(http.StatusTemporaryRedirect, utils.CLIENT_URL+"/auth?error=unauthorized")
	}

	// get user, create if not exists
	user, err := selectUserByEmailAndSub(userInfo.email, userInfo.sub)
	if err != nil {
		user, err = insertUser(userInfo.email, userInfo.sub, userInfo.avatar)
		if err != nil {
			slog.Error("Error inserting user", "insertUser", err)
			return c.Redirect(http.StatusTemporaryRedirect, utils.CLIENT_URL+"/auth?error=invalid_user")
		}
	}

	// delete old tokens, create new token
	err = deleteTokensByUserId(user.Id)
	if err != nil {
		slog.Error("Error deleting tokens", "deleteTokensByUserId", err)
		return c.Redirect(http.StatusTemporaryRedirect, utils.CLIENT_URL+"/auth?error=unauthorized")
	}
	newToken := Token{
		UserId:       user.Id,
		Provider:     provider,
		AccessToken:  oauthToken.AccessToken,
		RefreshToken: oauthToken.RefreshToken,
		TokenType:    oauthToken.TokenType,
		Expires:      oauthToken.Expiry,
	}
	token, err := insertToken(newToken)
	if err != nil {
		slog.Error("Error creating token", "createToken", err)
		return c.Redirect(http.StatusTemporaryRedirect, utils.CLIENT_URL+"/auth?error=unauthorized")
	}

	// set cookie
	cookie := &http.Cookie{}
	cookie.Domain = utils.COOKIE_DOMAIN
	cookie.Name = "token"
	cookie.Value = token.Id
	cookie.Path = "/"
	cookie.Secure = true
	cookie.SameSite = http.SameSiteLaxMode
	cookie.HttpOnly = true
	// 7 days
	cookie.MaxAge = 7 * 24 * 60 * 60
	c.SetCookie(cookie)

	// redirect to home page
	return c.Redirect(http.StatusTemporaryRedirect, utils.CLIENT_URL)
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
func UserAuth(ctx context.Context) (*pb.User, string, error) {
	claims, err := extractToken(ctx)
	if err != nil {
		return nil, "", fmt.Errorf("extractToken: %w", err)
	}

	// get oauth token from database
	token, err := selectTokenById(claims.TokenId)
	if err != nil {
		return nil, "", fmt.Errorf("selectTokenById: %w", err)
	}

	// check if oauth token is expired
	if token.Expires.Before(time.Now()) {
		// refresh oauth token
		configProvider := OAuthConfigImpl{}
		oauthToken, err := oauthRefresh(*token, &configProvider)
		if err != nil {
			return nil, "", fmt.Errorf("oauthRefresh: %w", err)
		}
		token.AccessToken = oauthToken.AccessToken
		token.RefreshToken = oauthToken.RefreshToken
		token.TokenType = oauthToken.TokenType
		token.Expires = oauthToken.Expiry
	}

	// create new phantom token
	err = deleteTokensByUserId(token.UserId)
	if err != nil {
		return nil, "", fmt.Errorf("deleteTokensByUserId: %w", err)
	}
	token, err = insertToken(*token)
	if err != nil {
		return nil, "", fmt.Errorf("insertToken: %w", err)
	}
	// get user from database
	user, err := selectUserById(token.UserId)
	if err != nil {
		return nil, "", fmt.Errorf("selectUserById: %w", err)
	}
	return user, token.Id, nil
}

/**
 * 1. Extract token from context
 * 2. Get oauth token from database
 * 3. Get user from database
 * 4. Return userId
 */
func UserCheck(ctx context.Context) (string, error) {
	claims, err := extractToken(ctx)
	if err != nil {
		return "", err
	}
	token, err := selectTokenById(claims.TokenId)
	if err != nil {
		return "", fmt.Errorf("selectTokenById: %w", err)
	}
	user, err := selectUserById(token.UserId)
	if err != nil {
		return "", fmt.Errorf("selectUserById: %w", err)
	}
	return user.Id, nil
}

func oauthRefresh(token Token, configProvider OAuthConfigProvider) (*oauth2.Token, error) {
	oauthToken := oauth2.Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		TokenType:    token.TokenType,
		Expiry:       token.Expires,
	}

	config, err := configProvider.getOAuthConfig(token.Provider)
	if err != nil {
		return nil, fmt.Errorf("getOAuthConfig: %w", err)
	}

	newOauthToken, err := config.TokenSource(context.Background(), &oauthToken).Token()
	if err != nil {
		return nil, fmt.Errorf("config.TokenSource: %w", err)
	}
	return newOauthToken, nil
}
