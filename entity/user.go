package entity

import (
	"time"
)

type (
	//User is a central entity, a person who makes all the payments
	User struct {
		ID        int64 `gorm:"primary_key"`
		FirstName string
		LastName  NullString
		Email     string
		Password  string
		CreatedAt time.Time
		DeletedAt NullTime
	}
)
