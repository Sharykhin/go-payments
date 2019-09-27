package request

import (
	types "github.com/Sharykhin/go-payments/core/type"
	identityEntity "github.com/Sharykhin/go-payments/domain/identity/entity"
)

type (
	// UserCreateRequest represents all the necessary data that should be used for creating a new user.
	// It also accepts a role
	UserCreateRequest struct {
		FirstName string
		LastName  types.NullString
		Password  string
		Email     string
		Role      identityEntity.Role
	}
)
