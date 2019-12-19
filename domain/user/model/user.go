package model

import (
	"time"

	types "github.com/Sharykhin/go-payments/core/type"
)

type (
	User struct {
		id        int64
		firstName string
		lastName  types.NullString
		email     string
		createdAt types.Time
		identity  Identity
		payments  []Payment
	}

	Identity struct {
		Password string `json:"-"`
	}
)

func (u *User) GetPayments() []Payment {
	return nil
}

func (u User) GetID() int64 {
	return u.id
}

func (u User) GetEmail() string {
	return u.email
}

func (u User) GetIdentity() Identity {
	return u.identity
}

func NewUser(id int64, firstName, email string, lastName types.NullString, identity Identity) *User {
	return &User{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		createdAt: types.Time(time.Now().UTC()),
		identity:  identity,
	}
}
