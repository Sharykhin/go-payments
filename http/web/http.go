package web

import (
	"html/template"
	"os"

	handler "github.com/Sharykhin/go-payments/http/handler/web"
	"github.com/gin-gonic/gin"
)

// ListenAndServe starts serving http requests
func ListenAndServe() error {
	r := gin.New()

	html := template.Must(template.ParseFiles("templates/web/index.tmpl", "templates/web/sign_up.tmpl"))
	r.SetHTMLTemplate(html)
	r.GET("/", handler.HomePage)
	r.GET("/sign-up", handler.SignUp)

	return r.Run(os.Getenv("WEB_ADDR"))
}
