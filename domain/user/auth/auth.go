package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Sharykhin/go-payments/core"
	"github.com/Sharykhin/go-payments/core/event"
	"github.com/Sharykhin/go-payments/core/logger"
	"github.com/Sharykhin/go-payments/core/queue"
	"github.com/Sharykhin/go-payments/domain/identity/jwt"
	"github.com/Sharykhin/go-payments/domain/identity/service/identity"
	userModel "github.com/Sharykhin/go-payments/domain/user/model"
	"github.com/Sharykhin/go-payments/domain/user/service"
)

type (
	// UserAuth provides API for authentication and authorization purposes
	UserAuth interface {
		SingIn(ctx context.Context, req UserSignInRequest) (*userModel.User, jwt.Token, error)
	}

	// AppUserAuth is a concrete implementation of UserAuth interface
	userAuth struct {
		userRetriever  service.UserRetriever
		userIdentifier identity.UserIdentity
		tokenManager   jwt.TokenManager
		dispatcher     queue.Publisher
	}

	UserSignInRequest struct {
		Email    string
		Password string
	}
)

func NewUserAuth(
	userRetriever service.UserRetriever,
	userIdentifier identity.UserIdentity,
	tokenManager jwt.TokenManager,
	dispatcher queue.Publisher,
) UserAuth {
	return &userAuth{
		userRetriever:  userRetriever,
		userIdentifier: userIdentifier,
		tokenManager:   tokenManager,
		dispatcher:     dispatcher,
	}
}

func (req UserSignInRequest) Validate() error {
	if req.Email == "" {
		return errors.New("email field is required")
	}

	if req.Password == "" {
		return errors.New("password field is required")
	}

	return nil
}

// SingIn signs user in by using general credentials such as email and password
// It also generate JWT token.
//TODO: just return JWT token it would be more semantic and obvious
func (s userAuth) SingIn(
	ctx context.Context,
	req UserSignInRequest,
) (
	*userModel.User,
	jwt.Token,
	error,
) {
	if err := req.Validate(); err != nil {
		return nil, "", err
	}

	user, err := s.userRetriever.FindUserByEmail(ctx, req.Email)
	if err != nil {
		// TODO: replace by NotFoundError
		return nil, "", fmt.Errorf("failed to find user by email: %v", err)
	}
	// TODO: replace by IncorrectPassword or to use errors.Is add create new error method
	if isValid, err := s.userIdentifier.ValidatePassword(ctx, user.GetIdentity().Password, req.Password); !isValid {
		return nil, "", fmt.Errorf("failed to validate password: %v", err)
	}

	tokenStr, err := s.tokenManager.Generate(jwt.Claims(map[string]interface{}{
		"userID": user.GetID(),
	}), time.Duration(1*time.Hour))
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate token: %v", err)
	}

	s.raiseSuccessSignInEvent(user.GetID())

	return user, jwt.Token(tokenStr), nil
}

func (s userAuth) raiseSuccessSignInEvent(userID int64) {
	err := s.dispatcher.RaiseEvent(event.NewEvent(event.UserSignIn, event.Payload{
		"UserID":  userID,
		"LoginAt": time.Now().UTC().Format(core.ISO8601),
	}))
	if err != nil {
		logger.Log.Error("failed to raise event %s: %v", event.UserSignIn, err)
	}
}
