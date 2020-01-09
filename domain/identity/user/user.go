package user

import (
	"context"
	"time"

	"github.com/Sharykhin/go-payments/core/logger"
	"github.com/Sharykhin/go-payments/core/queue"
	"github.com/Sharykhin/go-payments/domain/identity/repository"
)

// domain/identity/user.Identifier

type (
	Identifier interface {
		CreatePassword(ctx context.Context, userID int64, password string) (string, error)
		FindUserPassword(ctx context.Context, userID int64) (string, error)
		ValidatePassword(ctx context.Context, password string, compare string) (bool, error)
		UpdateLastLogin(ctx context.Context, userID int64, lastLogin time.Time) error
	}

	// identity is a default struct that implements UserIdentity interface
	identity struct {
		repository repository.IdentityRepository
		logger     logger.Logger
		dispatcher queue.Publisher
	}
)

func NewIdentifier(
	repository repository.IdentityRepository,
	logger logger.Logger,
	dispatcher queue.Publisher,
) Identifier {
	i := identity{
		repository: repository,
		logger:     logger,
		dispatcher: dispatcher,
	}
	return &i
}

func (i identity) CreatePassword(ctx context.Context, userID int64, pass string) (string, error) {
	return "", nil
}

func (i identity) FindUserPassword(ctx context.Context, userID int64) (string, error) {
	return "", nil
}

func (i identity) ValidatePassword(ctx context.Context, pass string, compare string) (bool, error) {
	return true, nil
}

func (i identity) UpdateLastLogin(ctx context.Context, userID int64, lastLogin time.Time) error {
	return nil
}
