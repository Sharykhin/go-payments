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
	PaymentAdapter struct {
		userService service.UserRetriever
	}
)

// NewPaymentAdapter returns a concrete implementation of UserAdapter
func NewPaymentAdapter() PaymentAdapter {
	return PaymentAdapter{
		userService: service.NewAppUserRetriever(),
	}
}

func (a PaymentAdapter) GetUser(ctx context.Context, userID int64) (model.UserInterface, error) {
	user, err := a.userService.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return model.NewUser(user.ID, user.Email), nil
}
