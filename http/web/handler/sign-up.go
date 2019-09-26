package web

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "sign_up.tmpl", gin.H{
		"API_ADDRESS": os.Getenv("API_ADDR"),
	})
}
