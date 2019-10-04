package auth

import (
	"github.com/Sharykhin/go-payments/core/errors"
	"github.com/Sharykhin/go-payments/core/logger"
	"github.com/Sharykhin/go-payments/domain/user/application/request"

	"github.com/Sharykhin/go-payments/http"

	"github.com/Sharykhin/go-payments/http/validation"

	"github.com/gin-gonic/gin"

	"github.com/Sharykhin/go-payments/core/locator"
	ar "github.com/Sharykhin/go-payments/http/request/auth"
)

func Login(c *gin.Context) {
	var req ar.LoginRequest
	if isValid, errors := validation.ValidateRequest(c, &req); !isValid {
		http.BadRequest(c, http.Errors(errors))
		return
	}

	auth := locator.NeUserAuthenticationService()
	user, token, err := auth.SingIn(c.Request.Context(), request.UserSignInRequest{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		logger.Error("user with request data %v could not to sign in: %v", req, err)
		http.BadRequest(c, http.Errors{errors.CredentialsDoNotMatch.Error()})
		return
	}
	logger.Info("user ID %d successfully signed in: %v", user.ID)
	http.OK(c, http.Data{
		"User":  user,
		"Token": token,
	}, nil)
}
