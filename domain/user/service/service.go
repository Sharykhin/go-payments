package service

import (
	"context"
	"fmt"

	"github.com/Sharykhin/go-payments/core/event"
	"github.com/Sharykhin/go-payments/core/queue"
	"github.com/Sharykhin/go-payments/domain/identity/service/password"
	"github.com/Sharykhin/go-payments/domain/user/repository"
	"github.com/Sharykhin/go-payments/domain/user/repository/entity"
)

type (
	Service interface {
		Create(ctx context.Context, user entity.User) (*entity.User, error)
	}

	UserService struct {
		repository repository.UserRepository
		dispatcher queue.Publisher
	}
)

func (us *UserService) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	hash, err := password.GeneratePassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("could not create a new user: %v", err)
	}
	user.Password = hash
	newUser, err := us.repository.Create(ctx, user)

	if err != nil {
		return nil, fmt.Errorf("could not create a new user: %v", err)
	}

	us.dispatcher.RaiseEvent(event.NewEvent(event.UserCreatedEvent, map[string]interface{}{
		"ID": newUser.ID,
	}))

	return newUser, err
}

func NewUserService() *UserService {
	return &UserService{
		repository: repository.NewGORMRepository(),
		dispatcher: queue.New(queue.RabbitMQ),
	}
}
