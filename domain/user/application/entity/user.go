package entity

import (
	"time"

	types "github.com/Sharykhin/go-payments/core/type"
)

type (
	User struct {
		ID        int64            `json:"id"`
		FirstName string           `json:"first_name"`
		LastName  types.NullString `json:"last_name"`
		Email     string           `json:"email"`
		CreatedAt time.Time        `json:"created_at"`
		Identity  Identity         `json:"-"`
		Payments  []Payment        `json:"payments,omitempty"`
	}

	Identity struct {
		Password string `json:"-"`
	}
)

func (u *User) GetPayments() []Payment {
	return nil
}
