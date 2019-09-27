package entity

import "time"

type (
	UserPassword struct {
		ID        uint64 `gorm:"primary_key"`
		UserID    int64
		Password  string
		CreatedAt time.Time
	}
)
