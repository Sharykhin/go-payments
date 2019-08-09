package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

const (
	headerType = "Bearer "
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, headerType) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header with Bearer type is required",
			})
			return
		}

		tokenString := authHeader[len(headerType):]
		fmt.Println("token", tokenString)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})

		if err != nil {
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
