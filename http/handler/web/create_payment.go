package web

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CreatePayment(c *gin.Context) {
	c.HTML(http.StatusOK, "create_payment.tmpl", gin.H{
		"API_ADDRESS": os.Getenv("API_ADDR"),
	})
}
