package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type (
	//User is a central entity, a person who makes all the payments
	User struct {
		ID        int64
		FirstName string
		LastName  NullString
		Email     string
		CreatedAt time.Time
		DeletedAt time.Time
	}

	Payment struct {
		TransactionID string
		User          User
		Date          time.Time
		Amount        decimal.Decimal
		Description   NullString
	}
)
