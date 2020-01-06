package locator

import (
	"github.com/Sharykhin/go-payments/core/database/gorm"
	"github.com/Sharykhin/go-payments/core/file"
	"github.com/Sharykhin/go-payments/core/file/local"
	"github.com/Sharykhin/go-payments/core/logger"
	"github.com/Sharykhin/go-payments/core/queue"
	"github.com/Sharykhin/go-payments/core/queue/rabbitmq"
	identityRepository "github.com/Sharykhin/go-payments/domain/identity/repository"
	identityService "github.com/Sharykhin/go-payments/domain/identity/service/identity"
	"github.com/Sharykhin/go-payments/domain/identity/service/token"
	paymentFactory "github.com/Sharykhin/go-payments/domain/payment/factory"
	"github.com/Sharykhin/go-payments/domain/payment/repository"
	paymentService "github.com/Sharykhin/go-payments/domain/payment/service"
	"github.com/Sharykhin/go-payments/domain/user/auth"
	userService "github.com/Sharykhin/go-payments/domain/user/service"
	gorm2 "github.com/jinzhu/gorm"
	"sync"
)

type (
	ServiceLocator struct {
		instances   map[string]interface{}
		initialized bool
		queue       queue.QueueManager
		log         logger.Logger
		db          *gorm2.DB
	}
)

func NewServiceLocator() *ServiceLocator {
	s := ServiceLocator{
		initialized: true,
		instances:   make(map[string]interface{}),
		queue:       rabbitmq.NewQueue(),
		log:         logger.NewLogger(),
		db:          gorm.NewGORMConnection(),
	}

	return &s
}

func (s *ServiceLocator) GetPublisherService() queue.Publisher {
	var once sync.Once

	once.Do(func() {
		s.instances["PublisherService"] = s.queue
	})

	return s.queue
}

func (s *ServiceLocator) GetLoggerService() logger.Logger {
	var once sync.Once

	once.Do(func() {
		s.instances["LoggerService"] = s.log
	})

	return s.log
}

func (s *ServiceLocator) GetSubscriberService() queue.Subscriber {
	var once sync.Once

	once.Do(func() {
		s.instances["SubscriberService"] = s.queue
	})

	return s.queue
}

func (s *ServiceLocator) GetIdentityService() identityService.UserIdentity {
	var once sync.Once
	once.Do(func() {
		inst := identityService.NewIdentityService(
			identityRepository.NewGORMRepository(
				s.db,
			),
			s.GetLoggerService(),
			s.GetPublisherService(),
		)
		s.instances["IdentityService"] = inst
	})

	return s.instances["IdentityService"].(identityService.UserIdentity)
}

func (s *ServiceLocator) GetUserService() userService.UserService {
	var once sync.Once
	once.Do(func() {
		inst := userService.NewUserService()
		s.instances["UserService"] = inst
	})

	return s.instances["UserService"].(userService.UserService)
}

func (s *ServiceLocator) GetTokenService() token.Tokener {
	var once sync.Once
	once.Do(func() {
		s.instances["TokenService"] = token.NewTokenService(token.TypeJWF)
	})

	return s.instances["TokenService"].(token.Tokener)
}

func (s *ServiceLocator) GetUploaderService() file.Uploader {
	var once sync.Once
	once.Do(func() {
		s.instances["UploaderService"] = local.NewUploader()
	})

	return s.instances["UploaderService"].(file.Uploader)
}

func (s *ServiceLocator) GetPaymentService() paymentService.PaymentService {
	var once sync.Once

	once.Do(func() {
		inst := struct {
			paymentService.PaymentCommander
			paymentService.PaymentRetriever
		}{
			paymentService.NewAppPaymentCommander(
				repository.NewGORMRepository(
					s.db,
				),
				s.queue,
				paymentFactory.NewPaymentFactory(
					s.GetUploaderService(),
					s.GetPublisherService(),
				),
			),
			paymentService.NewAppPaymentRetriever(
				repository.NewGORMRepository(
					s.db,
				),
				paymentFactory.NewPaymentFactory(
					s.GetUploaderService(),
					s.GetPublisherService()),
			),
		}

		s.instances["PaymentService"] = inst
	})

	return s.instances["PaymentService"].(paymentService.PaymentService)
}

func (s *ServiceLocator) GetAuthService() auth.UserAuth {
	var once sync.Once
	once.Do(func() {
		s.instances["AuthService"] = auth.NewUserAuth(
			s.GetUserService(),
			s.GetIdentityService(),
			s.GetTokenService(),
			s.GetPublisherService(),
		)
	})

	return s.instances["AuthService"].(auth.UserAuth)
}
