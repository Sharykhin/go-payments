package web

import (
	"os"

	"github.com/Sharykhin/go-payments/http/handler/web"
	"github.com/gin-gonic/gin"
)

// ListenAndServe starts serving http requests
func ListenAndServe() error {
	r := gin.New()
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/", web.HomePage)

	return r.Run(os.Getenv("WEB_ADDR"))
}
