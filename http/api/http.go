package api

import (
	"os"

	handler "github.com/Sharykhin/go-payments/http/handler/api"
	handlerAuth "github.com/Sharykhin/go-payments/http/handler/api/auth"
	handlerPayment "github.com/Sharykhin/go-payments/http/handler/api/payment"
	handlerUser "github.com/Sharykhin/go-payments/http/handler/api/user"
	"github.com/Sharykhin/go-payments/http/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// ListenAndServe starts serving http requests
func ListenAndServe() error {
	r := gin.New()
	store := cookie.NewStore([]byte("secret"))

	r.Use(sessions.Sessions("mysession", store))
	r.Use(cors.Default())
	r.Use(middleware.JsonContentType())
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", handler.Ping)
		v1.POST("/register", handlerUser.Register)
		v1.POST("/login", handlerAuth.Login)
		v1.GET("/users/:id/payments", handlerUser.GetUserPayments)
		auth := v1.Group("/")
		//auth.Use(middleware.Auth())
		auth.POST("/payments", handlerPayment.CreateTransaction)
	}

	return r.Run(os.Getenv("API_ADDR"))
}
