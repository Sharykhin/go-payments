package facade

import (
	"context"

	"github.com/Sharykhin/go-payments/domain/identity/service/identity"
	"github.com/Sharykhin/go-payments/domain/user/service"

	applicationEntity "github.com/Sharykhin/go-payments/domain/user/application/entity"
	"github.com/Sharykhin/go-payments/domain/user/application/request"
)

type (
	UserFacade interface {
		Create(ctx context.Context, req request.UserCreateRequest) (*applicationEntity.User, error)
	}

	AppUserFacade struct {
		userService     service.UserService
		identityService identity.UserIdentity
	}
)

func (a AppUserFacade) Create(ctx context.Context, req request.UserCreateRequest) (*applicationEntity.User, error) {
	u, _ := a.userService.Create(ctx, req)

	pass, _ := a.identityService.CreatePassword(ctx, u.ID, req.Password)

	return &applicationEntity.User{
		ID: u.ID,
		Identity: applicationEntity.Identity{
			Password: pass,
		},
	}, nil

}
