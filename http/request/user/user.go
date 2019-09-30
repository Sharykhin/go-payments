package user

import (
	types "github.com/Sharykhin/go-payments/core/type"
)

type (
	RegisterRequest struct {
		FirstName string           `json:"FirstName" binding:"required,max=80"`
		LastName  types.NullString `json:"LastName" binding:"max=80"`
		Email     string           `json:"Email" binding:"required,email,max=80"`
		Password  string           `json:"Password" binding:"required,min=8"`
	}
)
