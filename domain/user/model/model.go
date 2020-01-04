package model

import (
	"github.com/Sharykhin/go-payments/core/type"
	"time"
)

type (
	User struct {
		id        int64            `json:"ID"`
		firstName string           `json:"FirstName"`
		lastName  types.NullString `json:"LastName"`
		email     string           `json:"Email"`
		createdAt types.Time       `json:"Created"`
		identity  Identity         `json:"-"`
	}

	Identity struct {
		Password string `json:"-"`
	}
)

func NewUser(
	id int64,
	firstName, email string,
	lastName types.NullString,
	identity Identity,
) *User {
	return &User{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		createdAt: types.Time(time.Now()),
		identity:  identity,
	}
}
