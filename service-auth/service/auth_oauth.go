package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"service-auth/system"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var githubOAuthConfig = oauth2.Config{
	ClientID:     system.GITHUB_CLIENT_ID,
	ClientSecret: system.GITHUB_CLIENT_SECRET,
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://github.com/login/oauth/authorize",
		TokenURL: "https://github.com/login/oauth/access_token",
	},
	RedirectURL: system.SERVER_HTTP + "/oauth-callback/github",
	Scopes:      []string{"user:email"},
}

var googleOAuthConfig = oauth2.Config{
	ClientID:     system.GOOGLE_CLIENT_ID,
	ClientSecret: system.GOOGLE_CLIENT_SECRET,
	Endpoint:     google.Endpoint,
	RedirectURL:  system.SERVER_HTTP + "/oauth-callback/google",
	Scopes:       []string{"profile", "email", "openid"},
}

type UserInfo struct {
	email  string
	sub    string
	avatar string
}

type OAuth interface {
    getOAuthConfig() *oauth2.Config
	getUserInfo(accessToken string) (*UserInfo, error)
}

type OAuthGoogle struct {
	googleOAuthConfig oauth2.Config
}

func newOAuthGoogle() OAuth {
	return &OAuthGoogle{
		googleOAuthConfig: googleOAuthConfig,
	}
}

type OAuthGithub struct {
	githubOAuthConfig oauth2.Config
}

func newOAuthGithub() OAuth {
	return &OAuthGithub{
		githubOAuthConfig: githubOAuthConfig,
	}
}

func (o *OAuthGoogle) getOAuthConfig() *oauth2.Config {
    return &o.googleOAuthConfig
}

func (o *OAuthGoogle) getUserInfo(accessToken string) (*UserInfo, error) {
	url := "https://www.googleapis.com/oauth2/v2/userinfo"
	userInfo, err := httpCall("GET", url, accessToken)
	if err != nil {
		return nil, fmt.Errorf("httpCall: %w", err)
	}
	sub, ok := userInfo["id"].(string)
	if !ok {
		return nil, fmt.Errorf("Invalid user id")
	}
	email, ok := userInfo["email"].(string)
	if !ok {
		email = ""
	}
	avatar, ok := userInfo["picture"].(string)
	if !ok {
		avatar = ""
	}
	return &UserInfo{
		email:  email,
		sub:    sub,
		avatar: avatar,
	}, nil
}

func (o *OAuthGithub) getOAuthConfig() *oauth2.Config {
    return &o.githubOAuthConfig
}

func (o *OAuthGithub) getUserInfo(accessToken string) (*UserInfo, error) {
	url := "https://api.github.com/user"
	userInfo, err := httpCall("GET", url, accessToken)
	if err != nil {
		return nil, fmt.Errorf("httpCall: %w", err)
	}
	userId, ok := userInfo["id"].(float64)
	if !ok {
		return nil, fmt.Errorf("Invalid user id")
	}
	sub := fmt.Sprintf("%.0f", userId)
	email, ok := userInfo["email"].(string)
	if !ok {
		email = ""
	}
	avatar, ok := userInfo["avatar_url"].(string)
	if !ok {
		avatar = ""
	}
	return &UserInfo{
		email:  email,
		sub:    sub,
		avatar: avatar,
	}, nil
}

func httpCall(method, url, accessToken string) (map[string]interface{}, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client.Do: %w", err)
	}
	defer resp.Body.Close()
	var userInfo map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	return userInfo, nil
}
