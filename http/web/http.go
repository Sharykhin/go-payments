package web

import (
	"html/template"
	"os"

	handler "github.com/Sharykhin/go-payments/http/web/handler"
	"github.com/gin-gonic/gin"
)

// ListenAndServe starts serving http requests
func ListenAndServe() error {
	r := gin.New()

	html := template.Must(template.ParseFiles(
		"templates/web/index.tmpl",
		"templates/web/sign_up.tmpl",
		"templates/web/sign_in.tmpl",
		"templates/web/create_payment.tmpl",
	))
	r.SetHTMLTemplate(html)
	r.GET("/", handler.HomePage)
	r.GET("/sign-up", handler.SignUp)
	r.GET("/sign-in", handler.SignIn)
	r.GET("/payments/create", handler.CreatePayment)

	return r.Run(os.Getenv("WEB_ADDR"))
}
