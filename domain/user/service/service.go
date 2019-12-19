package service

import (
	"context"

	"github.com/Sharykhin/go-payments/domain/user/application/request"
	userrModel "github.com/Sharykhin/go-payments/domain/user/model"
)

type (
	// UserService provides create user method interface that is responsible for
	// fully creation user flow
	UserService interface {
		UserRetriever
		UserCommander
	}

	UserRetriever interface {
		FindUserByEmail(ctx context.Context, email string) (*userrModel.User, error)
		FindByID(ctx context.Context, ID int64) (*userrModel.User, error)
	}

	UserCommander interface {
		Create(ctx context.Context, req request.UserCreateRequest) (*userrModel.User, error)
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
