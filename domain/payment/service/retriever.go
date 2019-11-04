package service

import (
	"context"
	"fmt"

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

func NewAppPaymentRetriever() *AppPaymentRetriever {
	return &AppPaymentRetriever{
		repository: repository.NewGORMRepository(),
	}
}

func (a AppPaymentRetriever) All(ctx context.Context, criteria ...SearchCriteria) ([]model.Payment, error) {
	ps, _, err := a.repository.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not find list of payments: %v", err)
	}

	var pp []model.Payment
	// TODO: UserProxy works as lazy loading but what about eager loader
	for _, payment := range ps {
		pp = append(
			pp,
			*new(model.Payment).
				SetID(payment.ID).
				SetDescription(payment.Description).
				SetCreatedAt(types.Time(payment.ChargeDate)).
				SetAmount(value.NewAmount(value.USD, payment.Amount)).
				SetUser(proxy.NewUserProxy(payment.UserID)),
		)
	}

	return pp, nil
}
