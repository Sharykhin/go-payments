package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		fmt.Println("session", session.Get("user_id"))
		if _, valid := session.Get("user_id").(int64); !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not Authenticated"})
			c.Abort()
		}
		c.Next()
	}
}
