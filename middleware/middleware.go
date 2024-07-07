package middleware

import (
	"bioskuy/auth"
	"bioskuy/exception"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService auth.Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			c.Error(exception.ForbiddenError{Message: "Unauthorized"}).SetType(gin.ErrorTypePublic)
			c.AbortWithStatus(403)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		_, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.Error(exception.ForbiddenError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
			c.AbortWithStatus(403)
			return
		}

		c.Next()
	}
}