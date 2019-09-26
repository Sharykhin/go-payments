package web

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	c.HTML(http.StatusOK, "sign_in.tmpl", gin.H{
		"API_ADDRESS": os.Getenv("API_ADDR"),
	})
}
