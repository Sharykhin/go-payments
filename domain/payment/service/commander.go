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
	"github.com/Sharykhin/go-payments/domain/payment/request"
	"github.com/Sharykhin/go-payments/domain/payment/value"
)

type (
	// AppPaymentCommander is a concrete struct that implements PaymentCommander interface
	AppPaymentCommander struct {
		repository repository.PaymentRepository
		dispatcher queue.Publisher
	}
)

func NewAppPaymentCommander() *AppPaymentCommander {
	return &AppPaymentCommander{
		repository: repository.NewGORMRepository(),
		dispatcher: queue.Default(),
	}
}

// Create creates a new payment model
func (a AppPaymentCommander) Create(ctx context.Context, r request.NewPayment) (*model.Payment, error) {

	p, err := a.repository.Create(ctx, repository.Payment{
		UserID:        r.UserID,
		TransactionID: value.NewTransactionID(),
		Amount:        r.Amount.Value,
		Description:   r.Description,
		ChargeDate:    time.Now().UTC(),
		Status:        value.StatusOpen.String(),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create a new payment")
	}

	payment := model.NewPayment(
		p.ID,
		r.Amount,
		r.Description,
		types.Time(time.Now().UTC()),
		proxy.NewUserProxy(r.UserID),
	)

	return payment, nil
}
