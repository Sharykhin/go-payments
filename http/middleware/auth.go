package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Sharykhin/go-payments/entity"
	tokenPkg "github.com/Sharykhin/go-payments/identity/service/token"
	tokenErrors "github.com/Sharykhin/go-payments/identity/service/token/error"
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
		uc := entity.UserContext{
			ID:    int64(claims["id"].(float64)),
			Roles: []entity.Role{entity.Role(claims["roles"].(string))},
		}

		c.Set("UserContext", uc)
		c.Next()
	}
}
