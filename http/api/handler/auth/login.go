package auth

import (
	"github.com/Sharykhin/go-payments/domain/user/auth"
	"github.com/gin-gonic/gin"

	"github.com/Sharykhin/go-payments/core/errors"
	"github.com/Sharykhin/go-payments/core/logger"
	"github.com/Sharykhin/go-payments/http"
	ar "github.com/Sharykhin/go-payments/http/request/auth"
)

func Login(c *gin.Context, authService auth.UserAuth, log logger.Logger) {
	var req ar.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		http.BadRequest(c, http.Errors{err.Error()})
		return
	}

	user, token, err := authService.SingIn(c.Request.Context(), auth.UserSignInRequest{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		log.Error("user with request data %v could not to sign in: %v", req, err)
		http.BadRequest(c, http.Errors{errors.CredentialsDoNotMatch.Error()})
		return
	}

	log.Info("user ID %d successfully signed in: %v", user.GetID())

	http.OK(c, http.Data{
		"User":  user,
		"Token": token,
	}, nil)
}
