package authmock

import (
	"bioskuy/api/v1/user/entity"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/mock"
	"golang.org/x/oauth2"
	oauth2V2 "google.golang.org/api/oauth2/v2"
)

type AuthMock struct {
	mock.Mock
}

func (a *AuthMock) GenerateToken(user entity.User, c *gin.Context) (string, error) {
	args := a.Called(user, c)
	return args.String(0), args.Error(1)
}

func (a *AuthMock) ValidateToken(encodedToken string) (jwt.MapClaims, error) {
	args := a.Called(encodedToken)
	return args.Get(0).(jwt.MapClaims), args.Error(1)
}

type GoogleAuthMock struct {
	mock.Mock
}

func (g *GoogleAuthMock) GetGoogleLoginURL(state string) string {
	args := g.Called(state)
	return args.String(0)
}

func (g *GoogleAuthMock) GetGoogleUser(code string, c *gin.Context) (*oauth2.Token, *oauth2V2.Userinfo, error) {
	args := g.Called(code, c)
	return args.Get(0).(*oauth2.Token), args.Get(1).(*oauth2V2.Userinfo), args.Error(2)
}
