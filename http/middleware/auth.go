package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		if _, valid := session.Get("user_id").(int64); !valid {
			c.Abort()
		}
		c.Next()
	}
}
