package api

import (
	"os"

	handlerUser "github.com/Sharykhin/go-payments/http/handler/api/user"

	handler "github.com/Sharykhin/go-payments/http/handler/api"
	handlerPayment "github.com/Sharykhin/go-payments/http/handler/api/payment"
	"github.com/gin-gonic/gin"
)

// ListenAndServe starts serving http requests
func ListenAndServe() error {
	r := gin.New()
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", handler.Ping)
		v1.POST("/payments", handlerPayment.CreateTransaction)
		v1.POST("/register", handlerUser.Register)
	}

	return r.Run(os.Getenv("API_ADDR"))
}
