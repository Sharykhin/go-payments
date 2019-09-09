package api

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	handler "github.com/Sharykhin/go-payments/http/api/handler"
	handlerAuth "github.com/Sharykhin/go-payments/http/api/handler/auth"
	handlerPayment "github.com/Sharykhin/go-payments/http/api/handler/payment"
	handlerUser "github.com/Sharykhin/go-payments/http/api/handler/user"
	"github.com/Sharykhin/go-payments/http/middleware"
)

// ListenAndServe starts serving http requests
func ListenAndServe() error {
	r := gin.New()

	r.Use(cors.Default())
	r.Use(middleware.JsonContentType())
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", handler.Ping)
		v1.POST("/register", handlerUser.Register)
		v1.POST("/login", handlerAuth.Login)
	}

	auth := v1.Group("/")
	{
		auth.Use(middleware.AuthByToken())
		auth.GET("/users/:id/payments", handlerUser.GetUserPayments)
		auth.POST("/payments", handlerPayment.CreateTransaction)
	}

	return r.Run(os.Getenv("API_ADDR"))
}
