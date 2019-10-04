package middleware

import (
	"strings"

	"github.com/Sharykhin/go-payments/http"

	"github.com/gin-gonic/gin"
)

func JsonContentType() gin.HandlerFunc {
	return func(c *gin.Context) {
		cc := c.GetHeader("Content-Type")
		if (c.Request.Method == "POST" || c.Request.Method == "PUT") &&
			!strings.Contains(cc, "application/json") {
			http.BadRequest(c, http.Errors{"Content-Type must be application/json"})
			c.Abort()
		}

		c.Next()
	}
}
