package entity

import "time"

type (
	//User is a central entity, a person who makes all the payments
	User struct {
		ID        int64
		FirstName string
		LastName  NullString
		Email     string
		CreatedAt time.Time
	}
)
