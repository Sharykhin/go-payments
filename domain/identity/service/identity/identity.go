package identity

import (
	"context"
	"fmt"

	"github.com/Sharykhin/go-payments/domain/identity/service/password"

	"github.com/Sharykhin/go-payments/core/logger"
	"github.com/Sharykhin/go-payments/domain/identity/repository"
)

type (
	UserIdentity interface {
		CreatePassword(ctx context.Context, userID int64, password string) (string, error)
	}

	AppUserIdentity struct {
		repository repository.IdentityRepository
		logger     logger.Logger
	}
)

func (a AppUserIdentity) CreatePassword(ctx context.Context, userID int64, pass string) (string, error) {
	hash, err := password.GeneratePassword(pass)
	if err != nil {
		return "", fmt.Errorf("failed to generate a hash based on a user password: %v", err)
	}
	up, err := a.repository.CreatePassword(ctx, userID, hash)
	if err != nil {
		return "", fmt.Errorf("failed to create a new user password: %v", err)
	}

	return up.Password, nil
}

func NewUserIdentityService() *AppUserIdentity {
	return &AppUserIdentity{
		repository: repository.GORMRepository{},
		logger:     logger.Log,
	}
}
