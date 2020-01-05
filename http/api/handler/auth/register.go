package auth

import (
	"github.com/Sharykhin/go-payments/core/queue"
	"github.com/Sharykhin/go-payments/domain/user/service"
	"github.com/gin-gonic/gin"

	"github.com/Sharykhin/go-payments/core/event"
	"github.com/Sharykhin/go-payments/core/logger"
	identityEntity "github.com/Sharykhin/go-payments/domain/identity/entity"
	"github.com/Sharykhin/go-payments/domain/user/application/request"

	"github.com/Sharykhin/go-payments/http"
	ar "github.com/Sharykhin/go-payments/http/request/auth"
	"github.com/Sharykhin/go-payments/http/validation"
)

// Register method creates a new consumer in the system.
// Then it raises a separate event so we can send a welcome email if it is necessary
func Register(
	c *gin.Context,
	userService service.UserCommander,
	dispatcher queue.Publisher,
) {
	var req ar.RegisterRequest
	if isValid, errors := validation.ValidateRequest(c, &req); !isValid {
		http.BadRequest(c, http.Errors(errors))
		return
	}

	user, err := userService.Create(c.Request.Context(), request.UserCreateRequest{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Role:      identityEntity.RoleConsumer,
	})

	if err != nil {
		http.BadRequest(c, http.Errors{err.Error()})
		return
	}

	raiseSuccessfulRegistration(user.GetID(), dispatcher)

	http.Created(c, http.Data{
		"User": user,
	}, nil)
}

func raiseSuccessfulRegistration(userId int64, dispatcher queue.Publisher) {
	err := dispatcher.RaiseEvent(event.NewEvent(event.UserRegisteredEvent, event.Payload{
		"ID": userId,
	}))
	if err != nil {
		logger.Log.Error("failed to dispatch event %s: %v", event.UserRegisteredEvent, err)
	}

	logger.Log.Info("Raised an event %s", event.UserRegisteredEvent)
}
