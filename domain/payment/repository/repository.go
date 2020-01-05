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
		Create(cxt context.Context, payment PaymentAggregate) (*PaymentAggregate, error)
		List(ctx context.Context, criteria ...Criteria) ([]PaymentAggregate, error)
		FindByID(ctx context.Context, ID int64) (*PaymentAggregate, error)
	}

	// Criteria describes general conditional criteria
	// that can be applied to a repository
	Criteria interface {
		Name() string
	}

	// LimitCriteria apply condition to return a limited number of records
	LimitCriteria struct {
		Offset int64
		Limit  int64
	}
)

// Name returns criteria name
func (c LimitCriteria) Name() string {
	return LimitCriteriaName
}
