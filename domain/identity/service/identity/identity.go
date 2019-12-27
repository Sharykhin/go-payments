package identity

import (
	"context"
	"time"

	"github.com/Sharykhin/go-payments/core/database/gorm"
	"github.com/Sharykhin/go-payments/core/logger"
	"github.com/Sharykhin/go-payments/core/queue"
	"github.com/Sharykhin/go-payments/domain/identity/repository"
)

type (
	// UserIdentity is a general interface around user identity domain
	UserIdentity interface {
		CreatePassword(ctx context.Context, userID int64, password string) (string, error)
		FindUserPassword(ctx context.Context, userID int64) (string, error)
		ValidatePassword(ctx context.Context, password string, compare string) (bool, error)
		UpdateLastLogin(ctx context.Context, userID int64, lastLogin time.Time) error
	}
)

// TODO: it'e better to have function constructor along with main struct like it is done for repositories
// NewUserIdentityService is a function constructor that returns
// a new instance of AppUserIdentity struct
func NewUserIdentityService() *AppUserIdentity {
	return &AppUserIdentity{
		repository: repository.NewGORMRepository(gorm.NewGORMConnection()),
		logger:     logger.Log,
		dispatcher: queue.Default(),
	}
}
