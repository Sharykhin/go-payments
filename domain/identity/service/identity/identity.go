package identity

import (
	"context"
	"fmt"

	"github.com/Sharykhin/go-payments/core/event"
	"github.com/Sharykhin/go-payments/core/logger"
	"github.com/Sharykhin/go-payments/core/queue"
	"github.com/Sharykhin/go-payments/domain/identity/repository"
	"github.com/Sharykhin/go-payments/domain/identity/service/password"
)

type (
	UserIdentity interface {
		CreatePassword(ctx context.Context, userID int64, password string) (string, error)
		FindUserPassword(ctx context.Context, userID int64) (string, error)
		ValidatePassword(ctx context.Context, password string, compare string) (bool, error)
	}

	AppUserIdentity struct {
		repository repository.IdentityRepository
		logger     logger.Logger
		dispatcher queue.Publisher
	}
)

func NewUserIdentityService() *AppUserIdentity {
	return &AppUserIdentity{
		repository: repository.NewGORMRepository(),
		logger:     logger.Log,
		dispatcher: queue.Default(),
	}
}

func (a AppUserIdentity) CreatePassword(ctx context.Context, userID int64, pass string) (string, error) {
	hash, err := password.GeneratePassword(pass)
	if err != nil {
		return "", fmt.Errorf("failed to generate a hash based on a user password: %v", err)
	}
	up, err := a.repository.CreatePassword(ctx, userID, hash)
	if err != nil {
		return "", fmt.Errorf("failed to create a new user password: %v", err)
	}

	a.raiseSuccessfulPasswordCreation(up.ID, userID)

	return up.Password, nil
}

func (a AppUserIdentity) FindUserPassword(ctx context.Context, userID int64) (string, error) {
	up, err := a.repository.FindPasswordByUserID(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("could not find user password: %v", err)
	}

	if len(up) == 0 {
		return "", fmt.Errorf("you have no a valid password")
	}

	return up[0].Password, nil
}

func (a AppUserIdentity) ValidatePassword(ctx context.Context, pass string, compare string) (bool, error) {
	if err := password.ComparePasswords(compare, pass); err != nil {
		return false, err
	}

	return true, nil
}

func (a AppUserIdentity) raiseSuccessfulPasswordCreation(userPasswordID uint64, userID int64) {
	err := a.dispatcher.RaiseEvent(event.NewEvent(event.UserPasswordCreatedEvent, event.Payload{
		"userPasswordID": userPasswordID,
		"userID":         userID,
	}))

	if err != nil {
		a.logger.Error("failed to raise an event %s: %v", event.UserPasswordCreatedEvent, err)
	}
}
