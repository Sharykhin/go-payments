package service

import (
	"context"

	userApplicationEntity "github.com/Sharykhin/go-payments/domain/user/application/entity"
	"github.com/Sharykhin/go-payments/domain/user/application/request"
)

type (
	// UserService provides create user method interface that is responsible for
	// fully creation user flow
	UserService interface {
		UserRetriever
		UserCommander
	}

	UserRetriever interface {
		FindUserByEmail(ctx context.Context, email string) (*userApplicationEntity.User, error)
		FindByID(ctx context.Context, ID int64) (*userApplicationEntity.User, error)
	}

	UserCommander interface {
		Create(ctx context.Context, req request.UserCreateRequest) (*userApplicationEntity.User, error)
	}

	// AppUserService is a main instance that would satisfy UserService interface
	AppUserService struct {
		AppUserRetriever
		AppUserCommander
	}
)

// NewUserService returns a new instance of AppUserService
// that actually implements UserService interface
func NewUserService() *AppUserService {
	return &AppUserService{
		AppUserRetriever: *NewAppUserRetriever(),
		AppUserCommander: *NewAppUserCommander(),
	}
}
