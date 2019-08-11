package service

import (
	"context"
	"fmt"
	"time"

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
		events     []Event
	}

	Event struct {
		Time time.Time
		Name string
	}
)

func (us *UserService) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	newUser, err := us.repository.Create(ctx, user)

	if err != nil {
		return nil, fmt.Errorf("failed to create user")
	}
	us.events = append(us.events, NewEvent(UserCreatedEvent))
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
		events:     []Event{},
	}
}
