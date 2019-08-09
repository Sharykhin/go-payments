package middleware

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len("Bearer "):]
		fmt.Println("token", tokenString)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			exp := int64(claims["exp"].(float64))
			if time.Now().UTC().Unix()-exp > 0 {
				fmt.Println("Token expired")
			} else {
				fmt.Println("token is valid")
			}
		} else {
			fmt.Println(err)
		}

		c.Next()
	}
}
