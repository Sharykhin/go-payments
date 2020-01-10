package identity

import (
	"context"
	"fmt"
	"time"

	"github.com/Sharykhin/go-payments/domain/identity/hash"

	"github.com/Sharykhin/go-payments/core/event"
	"github.com/Sharykhin/go-payments/core/logger"
	"github.com/Sharykhin/go-payments/core/queue"
	types "github.com/Sharykhin/go-payments/core/type"
	"github.com/Sharykhin/go-payments/domain/identity/repository"
)

type (
	// UserIdentity is a default struct that implements UserIdentity interface
	UserIdentity struct {
		repository repository.IdentityRepository
		logger     logger.Logger
		dispatcher queue.Publisher
		hasher     hash.Hasher
	}
)

// NewIdentityService returns a new instance of identity service
// This service is used as port to integrate with other contexts
// or any component can use that UserIdentity API to interact with identity context
func NewIdentityService(
	repository repository.IdentityRepository,
	logger logger.Logger,
	dispatcher queue.Publisher,
	hasher hash.Hasher,
) *UserIdentity {

	return &UserIdentity{
		repository: repository,
		logger:     logger,
		dispatcher: dispatcher,
		hasher:     hasher,
	}
}

// CreatePassword creates a new password for a given user ID. it applies a corresponding hash function
// and raises an event that a user password has been created.
func (a UserIdentity) CreatePassword(ctx context.Context, userID int64, pass string) (string, error) {
	h, err := a.hasher.GenerateHash(pass)
	if err != nil {
		return "", fmt.Errorf("failed to generate a hash based on a user password: %v", err)
	}
	up, err := a.repository.CreatePassword(ctx, userID, string(h))
	if err != nil {
		return "", fmt.Errorf("failed to create a new user password: %v", err)
	}

	a.raiseSuccessfulPasswordCreation(up.ID, userID)

	return up.Password, nil
}

// FindUserPassword finds the latest user's password
func (a UserIdentity) FindUserPassword(ctx context.Context, userID int64) (string, error) {
	up, err := a.repository.FindPasswordByUserID(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("could not find user password: %v", err)
	}

	if len(up) == 0 {
		return "", fmt.Errorf("you have no a valid password")
	}

	return up[0].Password, nil
}

// ValidatePassword just validates whether a plaint text password is equal to its hashed one
func (a UserIdentity) ValidatePassword(ctx context.Context, pass string, compare string) (bool, error) {
	valid, err := a.hasher.ValidateHash(compare, hash.Hash(pass))

	return valid, err
}

// UpdateLastLogin updates user's last login
func (a UserIdentity) UpdateLastLogin(ctx context.Context, userID int64, lastLogin time.Time) error {
	err := a.repository.Update(ctx, userID, repository.UpdateFields{
		LastLogin: types.NullTime{
			Valid: true,
			Time:  lastLogin,
		},
	})

	if err != nil {
		return fmt.Errorf("failed to update user's last login: %v", err)
	}

	return nil
}

func (a UserIdentity) raiseSuccessfulPasswordCreation(userPasswordID uint64, userID int64) {
	err := a.dispatcher.RaiseEvent(event.NewEvent(event.UserPasswordCreatedEvent, event.Payload{
		"userPasswordID": userPasswordID,
		"userID":         userID,
	}))

	if err != nil {
		a.logger.Error("failed to raise an event %s: %v", event.UserPasswordCreatedEvent, err)
	}
}
