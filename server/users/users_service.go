package users

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	pb "sgsg/proto"
	"sgsg/utils"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var githubOAuthConfig = oauth2.Config{
	ClientID:     utils.GITHUB_CLIENT_ID,
	ClientSecret: utils.GITHUB_CLIENT_SECRET,
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://github.com/login/oauth/authorize",
		TokenURL: "https://github.com/login/oauth/access_token",
	},
	RedirectURL: utils.SERVER_HTTP + "/oauth-callback/github",
	Scopes:      []string{"user:email"},
}

var googleOAuthConfig = oauth2.Config{
	ClientID:     utils.GOOGLE_CLIENT_ID,
	ClientSecret: utils.GOOGLE_CLIENT_SECRET,
	Endpoint:     google.Endpoint,
	RedirectURL:  utils.SERVER_HTTP + "/oauth-callback/google",
	Scopes:       []string{"profile", "email", "openid"},
}

type OAuthConfigProvider interface {
	getOAuthConfig(provider string) (*oauth2.Config, error)
	getUserInfo(provider string, accessToken string) (*UserInfo, error)
}

type UserInfo struct {
	email  string
	sub    string
	avatar string
}

type OAuthConfigImpl struct{}

func (o *OAuthConfigImpl) getOAuthConfig(provider string) (*oauth2.Config, error) {
	if provider == "github" {
		return &githubOAuthConfig, nil
	} else if provider == "google" {
		return &googleOAuthConfig, nil
	}
	return nil, fmt.Errorf("Invalid provider")
}

func (o *OAuthConfigImpl) getUserInfo(provider string, accessToken string) (*UserInfo, error) {
	var url string
	if provider == "github" {
		url = "https://api.github.com/user"
	} else if provider == "google" {
		url = "https://www.googleapis.com/oauth2/v2/userinfo"
	} else {
		return nil, fmt.Errorf("Invalid provider")
	}

	// Create a GET request to fetch user information from GitHub
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}

	// Set the "Authorization" header with the access token
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	// Send the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client.Do: %w", err)
	}
	defer resp.Body.Close()

	// Parse the JSON response
	var userInfo map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return nil, fmt.Errorf("json.NewDecoder: %w", err)
	}

	var sub string
	var email string
	var avatar string
	if provider == "github" {
		sub = fmt.Sprintf("%.0f", userInfo["id"].(float64))
		email = userInfo["email"].(string)
		avatar = userInfo["avatar_url"].(string)
	} else if provider == "google" {
		sub = userInfo["id"].(string)
		email = userInfo["email"].(string)
		avatar = userInfo["picture"].(string)
	} else {
		return nil, fmt.Errorf("Invalid provider")
	}

	return &UserInfo{
		email:  email,
		sub:    sub,
		avatar: avatar,
	}, nil
}

var storage = utils.NewStorage()

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
		return nil, "", err
	}
	// get oauth token from database
	token, err := getToken(claims.TokenId)
	if err != nil {
		return nil, "", fmt.Errorf("getToken: %w", err)
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
	token, err = createToken(*token)
	if err != nil {
		return nil, "", fmt.Errorf("createToken: %w", err)
	}
	// get user from database
	user, err := getUserById(token.UserId)
	if err != nil {
		return nil, "", fmt.Errorf("AuthUser: %w", err)
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
	token, err := getToken(claims.TokenId)
	if err != nil {
		return "", fmt.Errorf("getToken: %w", err)
	}
	user, err := getUserById(token.UserId)
	if err != nil {
		return "", fmt.Errorf("getUserById: %w", err)
	}
	return user.Id, nil
}

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

	config, err := configProvider.getOAuthConfig(provider)
	if err != nil {
		slog.Error("Error getting provider", "configProvider.getOAuthConfig", err)
		return c.Redirect(http.StatusTemporaryRedirect, utils.CLIENT_URL+"/auth?error=unauthorized")
	}

	// make request to github oauth api
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

	// create user
	user, err := createUser(userInfo.email, userInfo.sub, userInfo.avatar)
	if err != nil {
		slog.Error("Error creating user", "createUser", err)
		return c.Redirect(http.StatusTemporaryRedirect, utils.CLIENT_URL+"/auth?error=invalid_user")
	}

	// create token
	newToken := Token{
		UserId:       user.Id,
		Provider:     provider,
		AccessToken:  oauthToken.AccessToken,
		RefreshToken: oauthToken.RefreshToken,
		TokenType:    oauthToken.TokenType,
		Expires:      oauthToken.Expiry,
	}
	token, err := createToken(newToken)
	if err != nil {
		slog.Error("Error creating token", "createToken", err)
		return c.Redirect(http.StatusTemporaryRedirect, utils.CLIENT_URL+"/auth?error=unauthorized")
	}

	// set header with set-cookie
	// c.Response().Header().Set("Set-Cookie", fmt.Sprintf("token=%s; Path=/; Secure; SameSite=Lax; HttpOnly; Max-Age=%d", token.Id, 7*24*60*60))

	// set cookie
	cookie := &http.Cookie{}
	cookie.Domain = utils.COOKIE_DOMAIN
	cookie.Name = "token"
	cookie.Value = token.Id
	cookie.Path = "/"
	cookie.Secure = true
	cookie.SameSite = http.SameSiteLaxMode
	cookie.HttpOnly = true
	// // 7 days
	cookie.MaxAge = 7 * 24 * 60 * 60
	c.SetCookie(cookie)

	// redirect to home page
	return c.Redirect(http.StatusTemporaryRedirect, utils.CLIENT_URL)
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
		return nil, err
	}

	newOauthToken, err := config.TokenSource(context.Background(), &oauthToken).Token()
	if err != nil {
		return nil, err
	}
	return newOauthToken, nil
}
