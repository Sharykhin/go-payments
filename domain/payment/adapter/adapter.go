package adapter

import (
	"context"

	"github.com/Sharykhin/go-payments/domain/payment/model"
	"github.com/Sharykhin/go-payments/domain/user/service"
)

type (
	// UserAdapter transform user domain model into payment context representation
	UserAdapter interface {
		GetUser(ctx context.Context, userID int64) (model.UserInterface, error)
	}
	// TODO: semantically it's not correct it should be concrete UserAdapter implementation
	DefaultUserAdapter struct {
		userService service.UserRetriever
	}
)

// NewUserAdapter returns a concrete implementation of UserAdapter
func NewUserAdapter() DefaultUserAdapter {
	return DefaultUserAdapter{
		userService: service.NewUserService(),
	}
}

func (a DefaultUserAdapter) GetUser(ctx context.Context, userID int64) (model.UserInterface, error) {
	user, err := a.userService.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return model.NewUser(user.GetID(), user.GetEmail()), nil
}
