package entity

import (
	"time"

	"github.com/Sharykhin/go-payments/domain/user/repository/entity"

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

// NewUserFromRepository creates a new application user model
// based on one that repository returned. It also requires a password
func NewUserFromRepository(user *entity.User, pass string) *User {
	return &User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		Identity: Identity{
			Password: pass,
		},
	}
}
