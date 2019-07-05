package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type (
	//User is a central entity, a person who makes all the payments
	User struct {
		ID        int64
		FirstName string
		LastName  NullString
		Email     string
		CreatedAt time.Time
	}

	Payment struct {
		TransactionID string
		User          User
		Date          time.Time
		Amount        decimal.Decimal
	}
)
