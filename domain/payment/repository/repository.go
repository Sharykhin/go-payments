package repository

import (
	"context"

	"github.com/Sharykhin/go-payments/domain/user/repository/entity"
)

type (
	// Repository provides all the methods to works with a storage layer
	PaymentRepository interface {
		Create(cxt context.Context, payment entity.Payment) (*entity.Payment, error)
	}
)
