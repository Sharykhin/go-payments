package auth

import (
	"fmt"

	"github.com/Sharykhin/go-payments/http"

	"github.com/Sharykhin/go-payments/core/event"
	"github.com/Sharykhin/go-payments/core/logger"

	"github.com/Sharykhin/go-payments/core/locator"

	identityEntity "github.com/Sharykhin/go-payments/domain/identity/entity"
	"github.com/Sharykhin/go-payments/domain/user/application/request"

	"github.com/gin-gonic/gin"
	validator "gopkg.in/go-playground/validator.v8"

	ur "github.com/Sharykhin/go-payments/http/request/user"
)

// Register method creates a new consumer in the system.
// Then it raises a separate event so we can send a welcome email if it is necessary
func Register(c *gin.Context) {
	var rr ur.RegisterRequest

	// TODO: heck
	if err := c.ShouldBindJSON(&rr); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errors []string
		for _, v := range validationErrors {
			errors = append(errors, fmt.Sprintf("%s %s", v.Field, v.ActualTag))
		}
		http.BadRequest(c, http.Errors(errors))
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
		http.BadRequest(c, http.Errors{err.Error()})
		return
	}

	raiseSuccessfulRegistration(user.ID)

	http.Created(c, http.Data{
		"user": user,
	}, nil)
}

func raiseSuccessfulRegistration(userId int64) {
	dispatcher := locator.GetDefaultQueue()

	err := dispatcher.RaiseEvent(event.NewEvent(event.UserRegisteredEvent, event.Payload{
		"ID": userId,
	}))
	if err != nil {
		logger.Log.Error("failed to dispatch event %s: %v", event.UserRegisteredEvent, err)
	}

	logger.Log.Info("Raised an event %s", event.UserRegisteredEvent)
}
