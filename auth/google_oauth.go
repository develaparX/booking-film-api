package auth

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/oauth2/v2"
)

var (
	googleOauthConfig *oauth2.Config
	oauthStateString  = "random"
)

func init() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/api/v1/user/google/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{oauth2.ScopeEmail, oauth2.ScopeProfile},
		Endpoint:     google.Endpoint,
	}
}

func GetGoogleLoginURL(state string) string {
	return googleOauthConfig.AuthCodeURL(state)
}

func GetGoogleUser(code string) (*oauth2.Token, *oauth2.Userinfo, error) {
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	oauth2Service, err := oauth2.New(googleOauthConfig.Client(context.Background(), token))
	if err != nil {
		return nil, nil, fmt.Errorf("oauth2 service creation failed: %s", err.Error())
	}
	userinfo, err := oauth2Service.Userinfo.Get().Do()
	if err != nil {
		return nil, nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	return token, userinfo, nil
}
