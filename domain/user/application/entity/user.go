package entity

import (
	"github.com/Sharykhin/go-payments/domain/user/repository/entity"

	types "github.com/Sharykhin/go-payments/core/type"
)

type (
	// TODO: do I really need to store Identity in the scope of User struct, looks like it is redundant
	User struct {
		ID        int64            `json:"ID"`
		FirstName string           `json:"FirstName"`
		LastName  types.NullString `json:"LastName"`
		Email     string           `json:"Email"`
		CreatedAt types.Time       `json:"Created"`
		Identity  Identity         `json:"-"`
		Payments  []Payment        `json:"Payments,omitempty"`
	}

	Identity struct {
		Password string `json:"-"`
	}
)

func (u *User) GetPayments() []Payment {
	return nil
}

// NewUserFromRepository creates a new application user model
// based on one that repository returned. It also requires a password
func NewUserFromRepository(user *entity.User, pass string) *User {
	return &User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: types.Time(user.CreatedAt),
		Identity: Identity{
			Password: pass,
		},
	}
}
