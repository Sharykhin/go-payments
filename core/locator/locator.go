package locator

import (
	"github.com/Sharykhin/go-payments/core/queue"
	"github.com/Sharykhin/go-payments/domain/user/auth"
	"github.com/Sharykhin/go-payments/domain/user/service"
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
func GetUserService() service.UserService {
	if _, ok := instances["UserService"]; ok {
		return instances["UserService"].(service.UserService)
	}
	inst := service.NewUserService()
	instances["UserService"] = inst
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

func GetUserCommanderService() service.UserCommander {
	if _, ok := instances["UserCommander"]; ok {
		return instances["UserCommander"].(service.UserCommander)
	}
	inst := service.NewAppUserCommander()
	instances["UserCommander"] = inst

	return inst
}

func GetUserRetrieverService() service.UserRetriever {
	if _, ok := instances["UserRetriever"]; ok {
		return instances["UserRetriever"].(service.UserRetriever)
	}
	inst := service.NewAppUserRetriever()
	instances["UserRetriever"] = inst

	return inst
}
