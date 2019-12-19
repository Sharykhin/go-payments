package auth

import (
	"github.com/Sharykhin/go-payments/domain/user/auth"
	"github.com/gin-gonic/gin"

	"github.com/Sharykhin/go-payments/core/errors"
	//"github.com/Sharykhin/go-payments/core/locator"
	"github.com/Sharykhin/go-payments/core/logger"
	"github.com/Sharykhin/go-payments/domain/user/application/request"
	"github.com/Sharykhin/go-payments/http"
	ar "github.com/Sharykhin/go-payments/http/request/auth"
	"github.com/Sharykhin/go-payments/http/validation"
)

func Login(c *gin.Context, auth auth.UserAuth) {
	// TODO: Do we really need to validate request maybe it's better to use more ValueObject and business rules?
	var req ar.LoginRequest
	if isValid, err := validation.ValidateRequest(c, &req); !isValid {
		http.BadRequest(c, http.Errors(err))
		return
	}

	//auth := locator.NeUserAuthenticationService()
	user, token, err := auth.SingIn(c.Request.Context(), request.UserSignInRequest{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		logger.Error("user with request data %v could not to sign in: %v", req, err)
		http.BadRequest(c, http.Errors{errors.CredentialsDoNotMatch.Error()})
		return
	}
	logger.Info("user ID %d successfully signed in: %v", user.GetID())
	http.OK(c, http.Data{
		"User":  user,
		"Token": token,
	}, nil)
}
