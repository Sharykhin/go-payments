package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/Sharykhin/go-payments/core"

	"github.com/Sharykhin/go-payments/core/event"
	"github.com/Sharykhin/go-payments/core/logger"
	"github.com/Sharykhin/go-payments/core/queue"
	identityApplicationEntity "github.com/Sharykhin/go-payments/domain/identity/application/entity"
	"github.com/Sharykhin/go-payments/domain/identity/service/identity"
	"github.com/Sharykhin/go-payments/domain/identity/service/token"
	userApplicationEntity "github.com/Sharykhin/go-payments/domain/user/application/entity"
	"github.com/Sharykhin/go-payments/domain/user/application/request"
	"github.com/Sharykhin/go-payments/domain/user/service"
)

type (
	// UserAuth provides API for authentication and authorization purposes
	UserAuth interface {
		SingIn(ctx context.Context, req request.UserSignInRequest) (*userApplicationEntity.User, identityApplicationEntity.Token, error)
	}

	// AppUserAuth is a concrete implementation of UserAuth interface
	AppUserAuth struct {
		userRetriever service.UserRetriever
		userIdentity  identity.UserIdentity
		token         token.Tokener
		dispatcher    queue.Publisher
	}
)

// NewAppUserAuth this is a function constructor
// that returns a new instance of AppUserAuth struct
func NewAppUserAuth() *AppUserAuth {
	return &AppUserAuth{
		userRetriever: service.NewAppUserRetriever(),
		userIdentity:  identity.NewUserIdentityService(),
		token:         token.NewTokenService(token.TypeJWF),
		dispatcher:    queue.Default(),
	}
}

// SingIn signs user in by using general credentials such as email and password
// It also generate JWT token.
//TODO: just return JWT token it would be more semantic and obvious
func (s AppUserAuth) SingIn(
	ctx context.Context,
	req request.UserSignInRequest,
) (
	*userApplicationEntity.User,
	identityApplicationEntity.Token,
	error,
) {
	au, err := s.userRetriever.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, "", fmt.Errorf("failed to find user by email: %v", err)
	}

	if isValid, err := s.userIdentity.ValidatePassword(ctx, au.Identity.Password, req.Password); !isValid {
		return nil, "", fmt.Errorf("failed to validate password: %v", err)
	}

	tokenStr, err := s.token.Generate(map[string]interface{}{
		"userID": au.ID,
	}, time.Duration(1*time.Hour))
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate token: %v", err)
	}

	s.raiseSuccessSignInEvent(au.ID)

	return au, identityApplicationEntity.Token(tokenStr), nil
}

func (s AppUserAuth) raiseSuccessSignInEvent(userID int64) {
	err := s.dispatcher.RaiseEvent(event.NewEvent(event.UserSignIn, event.Payload{
		"UserID":  userID,
		"LoginAt": time.Now().UTC().Format(core.ISO8601),
	}))
	if err != nil {
		logger.Log.Error("failed to raise event %s: %v", event.UserSignIn, err)
	}
}
