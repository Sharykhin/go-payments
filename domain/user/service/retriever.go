package service

import (
	"context"
	"fmt"

	"github.com/Sharykhin/go-payments/domain/identity/service/identity"

	"github.com/Sharykhin/go-payments/domain/user/model"
	"github.com/Sharykhin/go-payments/domain/user/repository"
)

type (
	// AppUserRetriever
	AppUserRetriever struct {
		userRepository      repository.UserRepository
		userIdentityService identity.UserIdentity
	}
)

// NewAppUserRetriever creates a new instance of AppUserRetriever struct
func NewAppUserRetriever() *AppUserRetriever {
	return &AppUserRetriever{
		userRepository:      repository.NewGORMRepository(),
		userIdentityService: identity.NewUserIdentityService(),
	}
}

func (s AppUserRetriever) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	ua, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from a repository: %v", err)
	}

	password, err := s.userIdentityService.FindUserPassword(ctx, ua.ID)

	if err != nil {
		return nil, fmt.Errorf("failed to find user password: %v", err)
	}

	user := model.NewUser(ua.ID, ua.FirstName, ua.Email, ua.LastName, model.Identity{
		Password: password,
	})

	return user, nil
}

func (s AppUserRetriever) FindByID(ctx context.Context, ID int64) (*model.User, error) {
	ua, err := s.userRepository.FindByID(ctx, ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from a repository: %v", err)
	}

	password, err := s.userIdentityService.FindUserPassword(ctx, ua.ID)

	if err != nil {
		return nil, fmt.Errorf("failed to find user password: %v", err)
	}

	user := model.NewUser(ua.ID, ua.FirstName, ua.Email, ua.LastName, model.Identity{
		Password: password,
	})

	return user, nil
}
