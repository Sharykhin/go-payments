package service

import (
	"context"
	"fmt"

	types "github.com/Sharykhin/go-payments/core/type"
	"github.com/Sharykhin/go-payments/domain/user/application/request"

	re "github.com/Sharykhin/go-payments/domain/user/repository/entity"

	"github.com/Sharykhin/go-payments/core/event"
	"github.com/Sharykhin/go-payments/core/logger"
	"github.com/Sharykhin/go-payments/core/queue"
	"github.com/Sharykhin/go-payments/domain/identity/service/password"
	ae "github.com/Sharykhin/go-payments/domain/user/application/entity"
	"github.com/Sharykhin/go-payments/domain/user/repository"
)

type (
	// UserService provides create user method interface that is responsible for
	// fully creation user flow
	UserService interface {
		Create(ctx context.Context, req request.UserCreateRequest) (*ae.User, error)
	}

	// AppUserService is a main instance that would satisfy UserService interface
	AppUserService struct {
		repository repository.UserRepository
		dispatcher queue.Publisher
		logger     logger.Logger
	}
)

// Create creates a new user and returns application user model
func (us *AppUserService) Create(ctx context.Context, req request.UserCreateRequest) (*ae.User, error) {
	hash, err := password.GeneratePassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("could not create a new user: %v", err)
	}

	//TODO: use factory
	ure := re.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hash,
		DeletedAt: types.NullTime{
			Valid: false,
		},
	}

	newUser, err := us.repository.Create(ctx, ure)

	if err != nil {
		return nil, fmt.Errorf("could not create a new user: %v", err)
	}

	//TODO: use factory
	appUser := ae.User{
		ID: newUser.ID,
		Identity: ae.Identity{
			Password: newUser.Password,
		},
	}

	err = us.dispatcher.RaiseEvent(event.NewEvent(event.UserCreatedEvent, map[string]interface{}{
		"ID": newUser.ID,
	}))

	if err != nil {
		us.logger.Error("failed to dispatch %s event: %v", event.UserCreatedEvent, err)
	}

	return &appUser, err
}

func NewUserService() *AppUserService {
	return &AppUserService{
		repository: repository.NewGORMRepository(),
		dispatcher: queue.New(queue.RabbitMQ),
		logger:     logger.Log,
	}
}
