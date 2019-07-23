package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JsonContentType() gin.HandlerFunc {
	return func(c *gin.Context) {
		cc := c.GetHeader("Content-Type")
		if !strings.Contains(cc, "application/json") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Content-Type must be application/json"})
			c.Abort()
		}

		c.Next()
	}
}
