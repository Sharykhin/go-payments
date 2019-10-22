package entity

import (
	types "github.com/Sharykhin/go-payments/core/type"
)

type (
	// User is a user representation in the context of identity domain
	User struct {
		ID        int64 `gorm:"primary_key"`
		LastLogin types.NullTime
	}
)
