package http

import (
	"os"

	"github.com/Sharykhin/go-payments/http/handler"
	handlerPayment "github.com/Sharykhin/go-payments/http/handler/payment"

	"github.com/gin-gonic/gin"
)

// ListenAndServe starts serving http requests
func ListenAndServe() error {
	r := gin.New()

	r.GET("/ping", handler.Ping)
	r.POST("/payments", handlerPayment.CreateTransaction)

	return r.Run(os.Getenv("SERVER_ADDR"))
}
