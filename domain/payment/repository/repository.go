package repository

import (
	"context"
)

type (
	// Repository provides all the methods to works with a storage layer
	PaymentRepository interface {
		Create(cxt context.Context, payment Payment) (*Payment, error)
	}
)
