package http

import (
	"github.com/Sharykhin/go-payments/http/handler"
	"github.com/gin-gonic/gin"
	"os"
)

// ListenAndServe starts serving http requests
func ListenAndServe() error {
	r := gin.Default()

	r.GET("/ping", handler.Ping)

	return r.Run(os.Getenv("SERVER_ADDR"))
}
