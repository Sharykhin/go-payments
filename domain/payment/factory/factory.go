package factory

import (
	"github.com/Sharykhin/go-payments/core/file/local"
	"github.com/Sharykhin/go-payments/core/queue/rabbitmq"
	"github.com/Sharykhin/go-payments/core/type"
	"github.com/Sharykhin/go-payments/domain/payment/model"
	"github.com/Sharykhin/go-payments/domain/payment/value"
)

type (
	PaymentFactory interface {
		NewPayment(
			id int64,
			amount value.Amount,
			description string,
			createdAt types.Time,
			user model.UserInterface,
		) *model.Payment
	}

	paymentFactory struct {
	}
)

func NewPaymentFactory() PaymentFactory {
	return &paymentFactory{}
}

func (f paymentFactory) NewPayment(
	id int64,
	amount value.Amount,
	description string,
	createdAt types.Time,
	user model.UserInterface,
) *model.Payment {
	return model.NewPayment(
		id,
		amount,
		description,
		createdAt,
		user,
		local.NewUploader(),
		rabbitmq.NewQueue(),
	)
}
