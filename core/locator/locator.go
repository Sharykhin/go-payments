package locator

import (
	"github.com/Sharykhin/go-payments/core/queue"
	identityService "github.com/Sharykhin/go-payments/domain/identity/service/identity"
	"github.com/Sharykhin/go-payments/domain/user/auth"
	userService "github.com/Sharykhin/go-payments/domain/user/service"
)

var (
	instances map[string]interface{} = make(map[string]interface{})
)

// GetDefaultQueue returns a default queue manager
// if that manager has already been initialized, return existing instance
func GetDefaultQueue() queue.QueueManager {
	if _, ok := instances["QueueManager"]; ok {
		return instances["QueueManager"].(queue.QueueManager)
	}
	inst := queue.Default()
	instances["QueueManager"] = inst
	return inst
}

// GetUserService returns an implementation of UserService interface
// if it already exists return the same instance
func GetUserService() userService.UserService {
	if _, ok := instances["UserService"]; ok {
		return instances["UserService"].(userService.UserService)
	}
	inst := userService.NewUserService()
	instances["UserService"] = inst
	return inst
}

// GetIdentityService returns implementation of UserIdentity interface
func GetIdentityService() identityService.UserIdentity {
	if _, ok := instances["UserIdentity"]; ok {
		return instances["UserIdentity"].(identityService.UserIdentity)
	}
	inst := identityService.NewUserIdentityService()
	instances["UserIdentity"] = inst
	return inst
}

func NeUserAuthenticationService() auth.UserAuth {
	if _, ok := instances["UserAuthentication"]; ok {
		return instances["UserAuthentication"].(auth.UserAuth)
	}
	inst := auth.NewAppUserAuth()
	instances["UserAuthentication"] = inst

	return inst
}

func GetUserCommanderService() userService.UserCommander {
	if _, ok := instances["UserCommander"]; ok {
		return instances["UserCommander"].(userService.UserCommander)
	}
	inst := userService.NewAppUserCommander()
	instances["UserCommander"] = inst

	return inst
}

func GetUserRetrieverService() userService.UserRetriever {
	if _, ok := instances["UserRetriever"]; ok {
		return instances["UserRetriever"].(userService.UserRetriever)
	}
	inst := userService.NewAppUserRetriever()
	instances["UserRetriever"] = inst

	return inst
}
