package auth

import (
	"context"
	"fmt"

	"github.com/Sharykhin/go-payments/core/queue"

	"github.com/Sharykhin/go-payments/domain/identity/service/identity"
	"github.com/Sharykhin/go-payments/domain/user/service"

	identityApplicationEntity "github.com/Sharykhin/go-payments/domain/identity/application/entity"
	userApplicationEntity "github.com/Sharykhin/go-payments/domain/user/application/entity"
	"github.com/Sharykhin/go-payments/domain/user/application/request"
)

type (
	UserAuth interface {
		SingIn(ctx context.Context, req request.UserSignInRequest) (*userApplicationEntity.User, identityApplicationEntity.Token, error)
	}

	AppUserAuth struct {
		userService         service.UserService
		userIdentityService identity.UserIdentity
		dispatcher          queue.Publisher
	}
)

func (s AppUserAuth) SingIn(ctx context.Context, req request.UserSignInRequest) (*userApplicationEntity.User, identityApplicationEntity.Token, error) {
	au, err := s.userService.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, "", fmt.Errorf("failed to find user by its email: %v", err)
	}

	if isValid, err := s.userIdentityService.ValidatePassword(ctx, au.Identity.Password, req.Password); !isValid || err != nil {
		return nil, "", fmt.Errorf("failed to find user by its email: %v", err)
	}

	return nil, "", nil
}
