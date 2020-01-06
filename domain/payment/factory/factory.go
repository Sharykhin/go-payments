package factory

import (
	"github.com/Sharykhin/go-payments/core/file"
	"github.com/Sharykhin/go-payments/core/queue"
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
		fileUploader file.Uploader
		dispatcher   queue.Publisher
	}
)

func NewPaymentFactory(fileUploader file.Uploader, dispatcher queue.Publisher) PaymentFactory {
	return &paymentFactory{
		fileUploader: fileUploader,
		dispatcher:   dispatcher,
	}
}

func (f paymentFactory) NewPayment(
	id int64,
	amount value.Amount,
	description string,
	createdAt types.Time,
	user model.UserInterface,
) *model.Payment {

	payment := model.NewPayment(
		id,
		amount,
		description,
		createdAt,
		user,
		f.fileUploader,
		f.dispatcher,
	)

	return payment
}
