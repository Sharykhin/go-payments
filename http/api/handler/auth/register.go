package auth

import (
	"net/http"

	"github.com/Sharykhin/go-payments/core/event"
	"github.com/Sharykhin/go-payments/core/logger"

	"github.com/Sharykhin/go-payments/core/locator"

	identityEntity "github.com/Sharykhin/go-payments/domain/identity/entity"
	"github.com/Sharykhin/go-payments/domain/user/application/request"

	"github.com/gin-gonic/gin"
	validator "gopkg.in/go-playground/validator.v8"

	ur "github.com/Sharykhin/go-payments/http/request/user"
)

// Register
func Register(c *gin.Context) {
	var rr ur.RegisterRequest

	// TODO: heck
	if err := c.ShouldBindJSON(&rr); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errors := make(map[string]string, len(validationErrors))
		for _, v := range validationErrors {
			errors[v.Name] = v.ActualTag
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errors})
		return
	}

	userService := locator.GetUserService()

	user, err := userService.Create(c.Request.Context(), request.UserCreateRequest{
		FirstName: rr.FirstName,
		LastName:  rr.LastName,
		Email:     rr.Email,
		Password:  rr.Password,
		Role:      identityEntity.RoleConsumer,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	dispatcher := locator.GetDefaultQueue()

	err = dispatcher.RaiseEvent(event.NewEvent(event.UserRegisteredEvent, event.Payload{
		"ID": user.ID,
	}))
	if err != nil {
		logger.Log.Error("failed to dispatch event %s: %v", event.UserRegisteredEvent, err)
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}
