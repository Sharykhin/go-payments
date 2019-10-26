package repository

import (
	"context"

	"github.com/Sharykhin/go-payments/domain/user/repository/entity"
)

type (
	// Repository describes methods to work with user on
	// a storage layer
	UserRepository interface {
		Create(cxt context.Context, user entity.User) (*entity.User, error)
		FindByEmail(ctx context.Context, email string) (*entity.User, error)
		FindByID(ctx context.Context, ID int64) (*entity.User, error)
	}
)
