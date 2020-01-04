package proxy

import (
	"context"

	"github.com/Sharykhin/go-payments/domain/payment/adapter"
	"github.com/Sharykhin/go-payments/domain/payment/model"
)

type (
	UserProxy struct {
		id          int64
		user        model.UserInterface
		userAdapter adapter.UserAdapter
	}
)

// GetID returns current user ID
func (up *UserProxy) GetID() int64 {
	return up.id
}

// GetEmail return user's email
func (up *UserProxy) GetEmail() string {
	if up.user == nil {
		up.user, _ = up.userAdapter.GetUser(context.Background(), up.id)
	}

	return up.user.GetEmail()
}

func NewUserProxy(ID int64) *UserProxy {
	return &UserProxy{
		id:          ID,
		user:        nil,
		userAdapter: adapter.NewUserAdapter(),
	}
}
