package entity

import (
	"time"

	"github.com/Sharykhin/go-payments/core/type"
)

type (
	User struct {
		ID        int64 `gorm:"primary_key"`
		FirstName string
		LastName  types.NullString
		Password  string
		Email     string
		CreatedAt time.Time
		DeletedAt types.NullTime
	}
)
