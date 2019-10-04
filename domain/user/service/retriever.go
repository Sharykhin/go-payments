package service

import (
	"context"
	"fmt"

	"github.com/Sharykhin/go-payments/domain/identity/service/identity"

	userApplicationEntity "github.com/Sharykhin/go-payments/domain/user/application/entity"
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

func (s AppUserRetriever) FindUserByEmail(ctx context.Context, email string) (*userApplicationEntity.User, error) {
	u, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from a repository: %v", err)
	}

	password, err := s.userIdentityService.FindUserPassword(ctx, u.ID)

	if err != nil {
		return nil, fmt.Errorf("failed to find user password: %v", err)
	}

	ua := userApplicationEntity.NewUserFromRepository(u, password)

	return ua, nil
}
