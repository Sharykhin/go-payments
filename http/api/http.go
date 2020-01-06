package api

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Sharykhin/go-payments/core/locator"
	handler "github.com/Sharykhin/go-payments/http/api/handler"
	handlerAuth "github.com/Sharykhin/go-payments/http/api/handler/auth"
	handlerPayment "github.com/Sharykhin/go-payments/http/api/handler/payment"
	"github.com/Sharykhin/go-payments/http/middleware"
)

// ListenAndServe starts serving http requests
func ListenAndServe(sl *locator.ServiceLocator) error {
	r := gin.New()

	r.Use(cors.Default())
	r.Use(middleware.JsonContentType())
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", handler.Ping)
		v1.POST("/register", func(c *gin.Context) {
			handlerAuth.Register(
				c,
				sl.GetUserService(),
				sl.GetPublisherService(),
			)
		})
		v1.POST("/login", func(c *gin.Context) {
			handlerAuth.Login(c, sl.GetAuthService(), sl.GetLoggerService())
		})
	}

	auth := v1.Group("/")
	{
		auth.Use(middleware.AuthByToken())
		auth.GET("/users/:id/payments", func(c *gin.Context) {
			handlerPayment.GetUserPayments(
				c,
				sl.GetPaymentService(),
			)
		})
		auth.POST("/users/:id/payments", func(c *gin.Context) {
			handlerPayment.CreatePayment(
				c,
				sl.GetPaymentService(),
			)
		})
	}

	return r.Run(os.Getenv("API_ADDR"))
}
