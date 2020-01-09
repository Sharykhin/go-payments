package repository

import (
	"context"

	types "github.com/Sharykhin/go-payments/core/type"
)

type (
	// Repository describes methods to work with user on
	// a storage layer
	IdentityRepository interface {
		CreatePassword(cxt context.Context, userID int64, password string) (*UserPassword, error)
		FindPasswordByUserID(cxt context.Context, userID int64) ([]UserPassword, error)
		Update(ctx context.Context, userID int64, fields UpdateFields) error
	}

	// UpdateFields represents fields that can be updated in term of Identity Domain
	UpdateFields struct {
		LastLogin types.NullTime
	}
)
