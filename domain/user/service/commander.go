package service

import (
	"context"
	"fmt"

	"github.com/Sharykhin/go-payments/core/event"
	"github.com/Sharykhin/go-payments/core/logger"
	"github.com/Sharykhin/go-payments/core/queue"
	"github.com/Sharykhin/go-payments/domain/identity/service/identity"
	userApplicationEntity "github.com/Sharykhin/go-payments/domain/user/application/entity"
	"github.com/Sharykhin/go-payments/domain/user/application/request"
	"github.com/Sharykhin/go-payments/domain/user/repository"
)

type (
	// AppUserCommander
	AppUserCommander struct {
		userRepository      repository.UserRepository
		userIdentityService identity.UserIdentity
		dispatcher          queue.Publisher
	}
)

// NewAppUserCommander returns a new instance of AppUserCommander
func NewAppUserCommander() *AppUserCommander {
	return &AppUserCommander{
		userRepository:      repository.NewGORMRepository(),
		userIdentityService: identity.NewUserIdentityService(),
		dispatcher:          queue.Default(),
	}
}

func (s AppUserCommander) Create(ctx context.Context, req request.UserCreateRequest) (*userApplicationEntity.User, error) {

	userEntity, err := s.userRepository.Create(ctx, repository.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	})

	if err != nil {
		return nil, fmt.Errorf("could not create a new user: %v", err)
	}

	pass, err := s.userIdentityService.CreatePassword(ctx, userEntity.ID, req.Password)
	if err != nil {
		s.raiseFailedPasswordCreation(userEntity.ID)
		return nil, fmt.Errorf("could not create user's password: %v", err)
	}

	appUser := userApplicationEntity.NewUserFromRepository(userEntity, pass)

	//userModel := model.User{}

	s.raiseUserSuccessCreation(userEntity.ID)

	return appUser, err
}

func (s AppUserCommander) raiseFailedPasswordCreation(userId int64) {
	err := s.dispatcher.RaiseEvent(event.NewEvent(event.UserPasswordCreationFailedEvent, map[string]interface{}{
		"ID": userId,
	}))

	if err != nil {
		logger.Log.Error("failed to dispatch %s event: %v", event.UserPasswordCreationFailedEvent, err)
	}
}

func (s AppUserCommander) raiseUserSuccessCreation(userId int64) {
	err := s.dispatcher.RaiseEvent(event.NewEvent(event.UserCreatedEvent, map[string]interface{}{
		"ID": userId,
	}))

	if err != nil {
		logger.Log.Error("failed to dispatch %s event: %v", event.UserCreatedEvent, err)
	}
}
