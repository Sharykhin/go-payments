package user

import (
	"github.com/Sharykhin/go-payments/core/type"
)

type (
	RegisterRequest struct {
		FirstName string           `json:"first_name" binding:"required,max=80"`
		LastName  types.NullString `json:"last_name" binding:"max=80"`
		Email     string           `json:"email" binding:"required,email,max=80"`
		Password  string           `json:"password" binding:"required,min=8"`
	}
)
