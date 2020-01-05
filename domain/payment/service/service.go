package service

import (
	"context"

	"github.com/Sharykhin/go-payments/domain/payment/model"
)

type (
	// PaymentService is a general interface that payment domain provides for
	// clients that will work with payment domain
	PaymentService interface {
		PaymentCommander
		PaymentRetriever
	}

	//TODO: deprecated
	PaymentAttachmentService interface {
		AttachFile(ctx context.Context, p *model.Payment) error
	}

	// PaymentRetriever is an interface that is responsible for retrieving payment.
	// And it works like some sort of a factory. Take into account it has not side effect
	PaymentRetriever interface {
		LimitedList(cxt context.Context, offset, limit int64) ([]model.Payment, error)
	}

	// PaymentCommander is responsible for side affects regarding payment domain.
	PaymentCommander interface {
		Create(ctx context.Context, req NewPaymentRequest) (*model.Payment, error)
	}

	SearchCriteria interface {
		ApplyCriteria()
	}
)
