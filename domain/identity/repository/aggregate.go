package repository

import (
	"time"

	types "github.com/Sharykhin/go-payments/core/type"
)

type (
	UserPassword struct {
		ID        uint64 `gorm:"primary_key"`
		UserID    int64
		Password  string
		CreatedAt time.Time
	}

	// User is a user representation in the context of identity domain
	User struct {
		ID        int64 `gorm:"primary_key"`
		LastLogin types.NullTime
	}
)
