package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Sharykhin/go-payments/identity/service/password"

	"github.com/Sharykhin/go-payments/core/queue"

	"github.com/Sharykhin/go-payments/user/repository"

	"github.com/Sharykhin/go-payments/user/entity"
)

const (
	UserCreatedEvent = "UserCreated"
)

type (
	Service interface {
		Create(ctx context.Context, user entity.User) (*entity.User, error)
	}

	UserService struct {
		repository repository.Repository
		dispatcher queue.Publisher
	}

	Event struct {
		Time time.Time
		Name string
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

	us.dispatcher.RaiseEvent(NewEvent(UserCreatedEvent))

	return newUser, err
}

func NewEvent(name string) Event {
	return Event{
		Time: time.Now().UTC(),
		Name: name,
	}
}

func NewUserService() *UserService {
	return &UserService{
		repository: repository.NewGORMRepository(),
		dispatcher: queue.New(queue.TypeLocal),
	}
}
