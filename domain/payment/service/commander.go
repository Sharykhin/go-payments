package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Sharykhin/go-payments/domain/payment/proxy"

	"github.com/Sharykhin/go-payments/core/queue"
	types "github.com/Sharykhin/go-payments/core/type"
	"github.com/Sharykhin/go-payments/domain/payment/model"
	"github.com/Sharykhin/go-payments/domain/payment/repository"
	"github.com/Sharykhin/go-payments/domain/payment/value"
)

type (
	// AppPaymentCommander is a concrete struct that implements PaymentCommander interface
	AppPaymentCommander struct {
		repository repository.PaymentRepository
		dispatcher queue.Publisher
	}

	NewPaymentRequest struct {
		Amount      value.Amount
		UserID      int64
		Description string
	}
)

func NewAppPaymentCommander(
	repo repository.PaymentRepository,
	dispatcher queue.Publisher,
) *AppPaymentCommander {
	return &AppPaymentCommander{
		repository: repo,
		dispatcher: dispatcher,
	}
}

// Create creates a new payment model
func (a AppPaymentCommander) Create(ctx context.Context, req NewPaymentRequest) (*model.Payment, error) {

	p, err := a.repository.Create(ctx, repository.PaymentAggregate{
		UserID:        req.UserID,
		TransactionID: value.NewTransactionID(),
		Amount:        req.Amount.Value,
		Description:   req.Description,
		ChargeDate:    time.Now().UTC(),
		Status:        value.StatusOpen.String(),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create a new payment")
	}

	payment := model.NewPayment(
		p.ID,
		req.Amount,
		req.Description,
		types.Time(time.Now().UTC()),
		proxy.NewUserProxy(req.UserID),
	)

	return payment, nil
}
