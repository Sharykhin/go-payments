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
	"github.com/Sharykhin/go-payments/domain/user/application/request"
	userModel "github.com/Sharykhin/go-payments/domain/user/model"
	"github.com/Sharykhin/go-payments/domain/user/service"
)

type (
	// UserAuth provides API for authentication and authorization purposes
	UserAuth interface {
		SingIn(ctx context.Context, req request.UserSignInRequest) (*userModel.User, identityApplicationEntity.Token, error)
	}

	// AppUserAuth is a concrete implementation of UserAuth interface
	AppUserAuth struct {
		userRetriever service.UserRetriever
		userIdentity  identity.UserIdentity
		token         token.Tokener
		dispatcher    queue.Publisher
	}

	UserSignInRequest struct {
		Email    string
		Password string
	}
)

func NewUserAuth(
	userRetriever service.UserRetriever,
	userIdentitier identity.UserIdentity,
	tokener token.Tokener,
	dispatcher queue.Publisher,
) UserAuth {
	return &AppUserAuth{
		userRetriever: userRetriever,
		userIdentity:  userIdentitier,
		token:         tokener,
		dispatcher:    dispatcher,
	}
}

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
	*userModel.User,
	identityApplicationEntity.Token,
	error,
) {
	user, err := s.userRetriever.FindUserByEmail(ctx, req.Email)
	if err != nil {
		// TODO: replace by NotFoundError
		return nil, "", fmt.Errorf("failed to find user by email: %v", err)
	}
	// TODO: replace by IncorrectPasssword or to use errors.Is add create new error method
	if isValid, err := s.userIdentity.ValidatePassword(ctx, user.GetIdentity().Password, req.Password); !isValid {
		return nil, "", fmt.Errorf("failed to validate password: %v", err)
	}

	tokenStr, err := s.token.Generate(map[string]interface{}{
		"userID": user.GetID(),
	}, time.Duration(1*time.Hour))
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate token: %v", err)
	}

	s.raiseSuccessSignInEvent(user.GetID())

	return user, identityApplicationEntity.Token(tokenStr), nil
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
