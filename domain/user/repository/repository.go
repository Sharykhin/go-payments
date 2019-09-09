package repository

import (
	"context"

	"github.com/Sharykhin/go-payments/user/entity"
)

type (
	// Repository describes methods to work with user on
	// a storage layer
	Repository interface {
		Create(cxt context.Context, user entity.User) (*entity.User, error)
	}
)
