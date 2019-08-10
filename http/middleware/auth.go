package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	tokenPkg "github.com/Sharykhin/go-payments/identity/service/token"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

const (
	headerType = "Bearer "
	authHeader = "Authorization"
)

var tokenService = tokenPkg.NewTokenService(tokenPkg.TypeJWF)

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
		if err != nil && err == TokenIs {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			exp := int64(claims["exp"].(float64))
			if time.Now().UTC().Unix()-exp > 0 {
				fmt.Println("Token expired")
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "Token has been expired",
				})
				return
			}
		}

		c.Next()
	}
}
