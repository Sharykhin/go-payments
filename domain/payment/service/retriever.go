package service

import (
	"context"
	"github.com/Sharykhin/go-payments/domain/payment/model"
	"github.com/Sharykhin/go-payments/domain/payment/repository"
)

type (
	AppPaymentRetriever struct {
		repository repository.PaymentRepository
	}
)

func (a AppPaymentRetriever) all(cxt context.Context, criteria ...SearchCriteria) ([]model.Payment, error) {
	return nil, nil
}
