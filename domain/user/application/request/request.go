package request

import (
	types "github.com/Sharykhin/go-payments/core/type"
	identityEntity "github.com/Sharykhin/go-payments/domain/identity/entity"
)

type (
	// TODO: move rquestss to apporiate packages causse they must be aware of incomint request and no depend on a separate package you bastard!
	// UserCreateRequest represents all the necessary data that should be used for creating a new user.
	// It also accepts a role
	UserCreateRequest struct {
		FirstName string
		LastName  types.NullString
		Password  string
		Email     string
		Role      identityEntity.Role
	}

	UserSignInRequest struct {
		Email    string
		Password string
	}
)
