package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "sign_up.tmpl", gin.H{})
}
