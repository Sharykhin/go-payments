package locator

import (
	"github.com/Sharykhin/go-payments/core/queue"
	"github.com/Sharykhin/go-payments/domain/identity/service/identity"
	"github.com/Sharykhin/go-payments/domain/user/service"
)

var (
	instances map[string]interface{}
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
func GetUserService() service.UserService {
	if _, ok := instances["UserService"]; ok {
		return instances["UserService"].(service.UserService)
	}
	inst := service.NewUserService()
	instances["UserService"] = inst
	return inst
}

func NeUserAuthenticationService() identity.UserAuthentication {
	if _, ok := instances["UserAuthentication"]; ok {
		return instances["UserAuthentication"].(identity.UserAuthentication)
	}
	inst := identity.NewUserAuthenticationService()
	instances["UserAuthentication"] = inst
	return inst
}
