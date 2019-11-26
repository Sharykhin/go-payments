package service

import (
	"context"

	"github.com/Sharykhin/go-payments/domain/payment/model"
	"github.com/Sharykhin/go-payments/domain/payment/request"
)

type (
	// PaymentService is a general interface that payment domain provides for
	// outer context to use
	PaymentService interface {
		PaymentCommander
		PaymentRetriever
	}

	PaymentRetriever interface {
		LimitedList(cxt context.Context, offset, limit int64) ([]model.Payment, error)
	}

	PaymentCommander interface {
		Create(ctx context.Context, r request.NewPayment) (*model.Payment, error)
	}

	SearchCriteria interface {
		ApplyCriteria()
	}
)
