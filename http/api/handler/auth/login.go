package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Sharykhin/go-payments/core/locator"
	identityRequest "github.com/Sharykhin/go-payments/http/request/identity"
)

func Login(c *gin.Context) {
	var req identityRequest.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auth := locator.NeUserAuthenticationService()
	token, err := auth.SingIn(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
