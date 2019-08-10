package entity

import (
	"time"

	"github.com/Sharykhin/go-payments/entity"
)

type (
	User struct {
		ID        int64 `gorm:"primary_key"`
		FirstName string
		LastName  entity.NullString
		Email     string
		CreatedAt time.Time
		DeletedAt entity.NullTime
	}
)
