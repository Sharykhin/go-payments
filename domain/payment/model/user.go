package model

import (
	"context"
	userApplicationEntity "github.com/Sharykhin/go-payments/domain/user/application/entity"
	userService "github.com/Sharykhin/go-payments/domain/user/service"
)

type (
	User struct {
		id    int64
		email string
	}

	UserProxy struct {
		id          int64
		user        *userApplicationEntity.User
		userService userService.UserRetriever
	}

	UserInterface interface {
		GetID() int64
		GetEmail() string
	}
)

func (u User) GetID() int64 {
	return u.id
}

func (u User) GetEmail() string {
	return u.email
}

func (up *UserProxy) GetID() int64 {
	return up.id
}

func (up *UserProxy) GetEmail() string {
	if up.user == nil {
		up.user, _ = up.userService.FindByID(context.Background(), up.id)
	}

	return up.user.Email
}

func NewUserProxy(ID int64) *UserProxy {
	return &UserProxy{
		id:          ID,
		user:        nil,
		userService: userService.NewAppUserRetriever(),
	}
}
