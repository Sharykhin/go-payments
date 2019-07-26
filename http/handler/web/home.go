package web

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":       "Posts",
		"API_ADDRESS": os.Getenv("API_ADDR"),
	})
}
