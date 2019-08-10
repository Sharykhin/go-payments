package repository

import "github.com/Sharykhin/go-payments/user/entity"

type (
	Repository interface {
		Create(user entity.User) (*entity.User, error)
	}
)
