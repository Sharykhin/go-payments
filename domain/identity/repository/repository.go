package repository

import (
	"context"

	"github.com/Sharykhin/go-payments/domain/identity/repository/entity"
)

type (
	// Repository describes methods to work with user on
	// a storage layer
	IdentityRepository interface {
		CreatePassword(cxt context.Context, userID int64, password string) (*entity.UserPassword, error)
	}
)
