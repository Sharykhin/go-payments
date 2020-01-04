package service

import (
	"context"
	"fmt"
	"github.com/Sharykhin/go-payments/core/file/local"

	types "github.com/Sharykhin/go-payments/core/type"
	"github.com/Sharykhin/go-payments/domain/payment/model"
	"github.com/Sharykhin/go-payments/domain/payment/proxy"
	"github.com/Sharykhin/go-payments/domain/payment/repository"
	"github.com/Sharykhin/go-payments/domain/payment/value"
)

type (
	AppPaymentRetriever struct {
		repository repository.PaymentRepository
	}
)

// NewAppPaymentRetriever is a function constructor
// that returns a concrete implementation of PaymentRetriever interface
func NewAppPaymentRetriever(repo repository.PaymentRepository) *AppPaymentRetriever {
	return &AppPaymentRetriever{
		repository: repo,
	}
}

// LimitedList returns limited number of payments records
func (a AppPaymentRetriever) LimitedList(ctx context.Context, offset, limit int64) ([]model.Payment, error) {
	payments, err := a.repository.List(ctx, repository.LimitCriteria{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get a list of payments from a repository: %v", err)
	}

	var pp []model.Payment
	// TODO: UserProxy works as lazy loading but what about eager loader
	for _, payment := range payments {
		pp = append(
			pp,
			*model.NewPayment(
				payment.ID,
				value.NewAmount(value.USD, payment.Amount),
				payment.Description,
				types.Time(payment.ChargeDate),
				proxy.NewUserProxy(payment.UserID),
				local.NewUploader(),
			),
		)
	}

	return pp, nil
}
