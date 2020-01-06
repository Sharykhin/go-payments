package service

import (
	"context"
	"fmt"

	"github.com/Sharykhin/go-payments/core/event"
	"github.com/Sharykhin/go-payments/core/logger"
	"github.com/Sharykhin/go-payments/core/queue"
	"github.com/Sharykhin/go-payments/domain/identity/service/identity"
	"github.com/Sharykhin/go-payments/domain/user/application/request"
	useModel "github.com/Sharykhin/go-payments/domain/user/model"
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
		userIdentityService: identity.NewIdentityService(),
		dispatcher:          queue.Default(),
	}
}

func (s AppUserCommander) Create(ctx context.Context, req request.UserCreateRequest) (*useModel.User, error) {

	ua, err := s.userRepository.Create(ctx, repository.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	})

	if err != nil {
		return nil, fmt.Errorf("could not create a new user: %v", err)
	}

	password, err := s.userIdentityService.CreatePassword(ctx, ua.ID, req.Password)
	if err != nil {
		s.raiseFailedPasswordCreation(ua.ID)
		return nil, fmt.Errorf("could not create user's password: %v", err)
	}

	appUser := useModel.NewUser(ua.ID, ua.FirstName, ua.Email, ua.LastName, useModel.Identity{
		Password: password,
	})

	s.raiseUserSuccessCreation(ua.ID)

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
