package service

import (
	"context"

	identityApplicationEntity "github.com/Sharykhin/go-payments/domain/identity/application/entity"
	userApplicationEntity "github.com/Sharykhin/go-payments/domain/user/application/entity"
	"github.com/Sharykhin/go-payments/domain/user/application/request"
)

type (
	AuthService interface {
		SingIn(ctx context.Context, req request.UserSignInRequest) (*userApplicationEntity.User, identityApplicationEntity.Token, error)
	}
)
