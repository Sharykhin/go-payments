package locator

import (
	"github.com/Sharykhin/go-payments/core/database/gorm"
	"github.com/Sharykhin/go-payments/core/logger"
	"github.com/Sharykhin/go-payments/core/queue"
	"github.com/Sharykhin/go-payments/core/queue/rabbitmq"
	identityRepository "github.com/Sharykhin/go-payments/domain/identity/repository"
	identityService "github.com/Sharykhin/go-payments/domain/identity/service/identity"
	"github.com/Sharykhin/go-payments/domain/identity/service/token"
	paymentFactory "github.com/Sharykhin/go-payments/domain/payment/factory"
	"github.com/Sharykhin/go-payments/domain/payment/repository"
	paymentService "github.com/Sharykhin/go-payments/domain/payment/service"
	userService "github.com/Sharykhin/go-payments/domain/user/service"
	"sync"
)

var (
	instances map[string]interface{} = make(map[string]interface{})
	logImpl   logger.Logger
	queueImpl queue.QueueManager
)

func init() {
	logImpl = logger.NewLogger()
	queueImpl = rabbitmq.NewQueue()
}

func GetQueueService() queue.QueueManager {
	return queueImpl
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
	var once sync.Once
	once.Do(func() {
		inst := identityService.NewIdentityService(
			identityRepository.NewGORMRepository(
				gorm.NewGORMConnection(),
			),
			logImpl,
			queueImpl,
		)
		instances["UserIdentity"] = inst
	})

	return instances["UserIdentity"].(identityService.UserIdentity)
}

func GetTokenService() token.Tokener {
	var once sync.Once
	once.Do(func() {
		inst := token.NewTokenService(token.TypeJWF)
		instances["Tokener"] = inst
	})

	return instances["Tokener"].(token.Tokener)
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

// GetPaymentService returns an instance of payment service
// that includes either as mutator as retriever.
func GetPaymentService() paymentService.PaymentService {
	var once sync.Once

	once.Do(func() {
		inst := struct {
			paymentService.PaymentCommander
			paymentService.PaymentRetriever
		}{
			paymentService.NewAppPaymentCommander(
				repository.NewGORMRepository(
					gorm.NewGORMConnection(),
				),
				queueImpl,
				paymentFactory.NewPaymentFactory(),
			),
			paymentService.NewAppPaymentRetriever(
				repository.NewGORMRepository(
					gorm.NewGORMConnection(),
				),
				paymentFactory.NewPaymentFactory(),
			),
		}

		instances["PaymentService"] = inst
	})

	return instances["PaymentService"].(paymentService.PaymentService)
}
