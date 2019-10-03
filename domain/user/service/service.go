package service

import (
	"context"
	"fmt"

	"github.com/Sharykhin/go-payments/core/event"
	"github.com/Sharykhin/go-payments/core/logger"
	"github.com/Sharykhin/go-payments/core/queue"
	identityApplicationEntity "github.com/Sharykhin/go-payments/domain/identity/application/entity"
	"github.com/Sharykhin/go-payments/domain/identity/service/identity"
	userApplicationEntity "github.com/Sharykhin/go-payments/domain/user/application/entity"
	"github.com/Sharykhin/go-payments/domain/user/application/request"
	"github.com/Sharykhin/go-payments/domain/user/repository"
	userRepositoryEntity "github.com/Sharykhin/go-payments/domain/user/repository/entity"
)

type (
	// UserService provides create user method interface that is responsible for
	// fully creation user flow
	UserService interface {
		Create(ctx context.Context, req request.UserCreateRequest) (*userApplicationEntity.User, error)
		SingIn(ctx context.Context, req request.UserSignInRequest) (*userApplicationEntity.User, identityApplicationEntity.Token, error)
	}

	// AppUserService is a main instance that would satisfy UserService interface
	AppUserService struct {
		userRepository repository.UserRepository
		userIdentity   identity.UserIdentity
		dispatcher     queue.Publisher
		logger         logger.Logger
	}
)

// NewUserService returns a new instance of AppUserService
// that actually implements UserService interface
func NewUserService() *AppUserService {
	return &AppUserService{
		userRepository: repository.NewGORMRepository(),
		userIdentity:   identity.NewUserIdentityService(),
		dispatcher:     queue.New(queue.RabbitMQ),
		logger:         logger.Log,
	}
}

func (us *AppUserService) SingIn(ctx context.Context, req request.UserSignInRequest) (*userApplicationEntity.User, identityApplicationEntity.Token, error) {
	return nil, "", nil
}

// Create creates a new user and returns application user model
func (us *AppUserService) Create(ctx context.Context, req request.UserCreateRequest) (*userApplicationEntity.User, error) {

	user := userRepositoryEntity.NewUser(req.FirstName, req.LastName.String, req.Email)

	newUser, err := us.userRepository.Create(ctx, user)

	if err != nil {
		return nil, fmt.Errorf("could not create a new user: %v", err)
	}

	pass, err := us.userIdentity.CreatePassword(ctx, newUser.ID, req.Password)
	if err != nil {
		us.raiseFailedPasswordCreation(newUser.ID)
		return nil, fmt.Errorf("could not create user's password: %v", err)
	}

	appUser := userApplicationEntity.NewUserFromRepository(newUser, pass)

	us.raiseUserSuccessCreation(newUser.ID)

	return appUser, err
}

func (us *AppUserService) raiseFailedPasswordCreation(userId int64) {
	err := us.dispatcher.RaiseEvent(event.NewEvent(event.UserPasswordCreationFailedEvent, map[string]interface{}{
		"ID": userId,
	}))

	if err != nil {
		us.logger.Error("failed to dispatch %s event: %v", event.UserPasswordCreationFailedEvent, err)
	}
}

func (us *AppUserService) raiseUserSuccessCreation(userId int64) {
	err := us.dispatcher.RaiseEvent(event.NewEvent(event.UserCreatedEvent, map[string]interface{}{
		"ID": userId,
	}))

	if err != nil {
		us.logger.Error("failed to dispatch %s event: %v", event.UserCreatedEvent, err)
	}
}
