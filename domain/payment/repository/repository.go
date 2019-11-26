package repository

import (
	"context"
)

const (
	LimitCriteriaName = "LimitCriteria"
)

type (
	// Repository provides all the methods to works with a storage layer
	PaymentRepository interface {
		Create(cxt context.Context, payment Payment) (*Payment, error)
		List(ctx context.Context, criteria ...Criteria) ([]Payment, error)
	}

	Criteria interface {
		Name() string
	}

	LimitCriteria struct {
		Offset int64
		Limit  int64
	}
)

func (c LimitCriteria) Name() string {
	return LimitCriteriaName
}
