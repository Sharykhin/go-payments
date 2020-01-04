package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	identityEntity "github.com/Sharykhin/go-payments/domain/identity/entity"
	tokenPkg "github.com/Sharykhin/go-payments/domain/identity/service/token"
	tokenErrors "github.com/Sharykhin/go-payments/domain/identity/service/token/error"
	"github.com/Sharykhin/go-payments/http"
	httpApp "github.com/Sharykhin/go-payments/http"
)

const (
	headerType = "Bearer "
	authHeader = "Authorization"
)

var tokenService = tokenPkg.NewTokenService(tokenPkg.TypeJWF)

// AuthByToken checks the income requests based on Authorization headers
// and after that populates the request with a user context
func AuthByToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(authHeader)
		if !strings.Contains(authHeader, headerType) {
			http.Unauthorized(c, http.Errors{"Authorization header with Bearer type is required"})
			c.Abort()
			return
		}

		tokenString := authHeader[len(headerType):]

		claims, err := tokenService.Validate(tokenString)
		if err != nil {
			if err == tokenErrors.TokenIsExpired {
				http.Unauthorized(c, http.Errors{"Token is expired"})
				c.Abort()
				return
			}

			http.ServerError(c, http.Errors{err.Error()})
			c.Abort()
			return

		}
		uc := identityEntity.UserContext{
			ID: int64(claims["userID"].(float64)),
			//Roles: []identityEntity.Role{identityEntity.Role(claims["role"].(int64))},
		}

		c.Set(httpApp.UserContext, uc)
		c.Next()
	}
}
