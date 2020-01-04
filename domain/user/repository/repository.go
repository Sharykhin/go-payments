package repository

import (
	"context"
)

type (
	// Repository describes methods to work with user on
	// a storage layer
	UserRepository interface {
		Create(cxt context.Context, user User) (*User, error)
		FindByEmail(ctx context.Context, email string) (*User, error)
		FindByID(ctx context.Context, ID int64) (*User, error)
	}
)
