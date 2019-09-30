package identity

import (
	"context"
	"fmt"
	"time"

	"github.com/Sharykhin/go-payments/domain/identity/entity"
	"github.com/dgrijalva/jwt-go"

	"github.com/Sharykhin/go-payments/domain/identity/service/password"

	"github.com/Sharykhin/go-payments/core/logger"
	"github.com/Sharykhin/go-payments/domain/identity/repository"
)

type (
	UserIdentity interface {
		CreatePassword(ctx context.Context, userID int64, password string) (string, error)
	}

	UserAuthentication interface {
		SingIn(ctx context.Context, email, password string) (entity.Token, error)
	}

	AppUserIdentity struct {
		repository repository.IdentityRepository
		logger     logger.Logger
	}

	AppUserAuthentication struct {
		identityRepository repository.IdentityRepository
		logger             logger.Logger
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

func (a AppUserAuthentication) SingIn(ctx context.Context, email, password string) (entity.Token, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  23,
		"exp": time.Now().UTC().Add(1 * time.Second).Unix(),
	})

	tokenStr, err := token.SignedString([]byte("secret"))
	if err != nil {
		return entity.Token(""), fmt.Errorf("faield to parse token")
	}

	return entity.Token(tokenStr), nil
}

func NewUserIdentityService() *AppUserIdentity {
	return &AppUserIdentity{
		repository: repository.NewGORMRepository(),
		logger:     logger.Log,
	}
}

func NewUserAuthenticationService() *AppUserAuthentication {
	return &AppUserAuthentication{
		identityRepository: repository.NewGORMRepository(),
		logger:             logger.Log,
	}
}
