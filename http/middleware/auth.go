package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	identityEntity "github.com/Sharykhin/go-payments/domain/identity/entity"
	tokenPkg "github.com/Sharykhin/go-payments/domain/identity/service/token"
	tokenErrors "github.com/Sharykhin/go-payments/domain/identity/service/token/error"
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
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header with Bearer type is required",
			})
			return
		}

		tokenString := authHeader[len(headerType):]

		claims, err := tokenService.Validate(tokenString)
		if err != nil {
			if err == tokenErrors.TokenIsExpired {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "Token is expired",
				})
				return
			}

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
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
