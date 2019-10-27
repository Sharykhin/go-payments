package service

import (
	"context"
	"github.com/Sharykhin/go-payments/domain/payment/model"
	"github.com/Sharykhin/go-payments/domain/payment/request"
)

type (
	PaymentService interface {
		PaymentCommander
		PaymentRetriever
	}

	PaymentRetriever interface {
		All(cxt context.Context, criteria ...SearchCriteria) ([]model.Payment, error)
	}

	PaymentCommander interface {
		Create(ctx context.Context, r request.NewPayment) (*model.Payment, error)
	}

	SearchCriteria interface {
		ApplyCriteria()
	}
)
