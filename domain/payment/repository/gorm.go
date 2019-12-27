package repository

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
)

type (
	// GORMRepository struct responsible for working with
	// storage layer by using GORM ORM
	GORMRepository struct {
		conn *gorm.DB
	}
)

// NewGORMRepository is a constructor function that returns a new instance of GORMRepository
// and satisfies PaymentRepository interface
func NewGORMRepository(conn *gorm.DB) *GORMRepository {
	return &GORMRepository{
		conn: conn,
	}
}

// Create creates a new payment record in a database and returns just created record
func (r GORMRepository) Create(cxt context.Context, payment PaymentAggregate) (*PaymentAggregate, error) {
	err := r.conn.Create(&payment).Error
	if err != nil {
		return nil, fmt.Errorf("failed to insert a new payment row: %v", err)
	}

	return &payment, nil
}

// List returns list of payments rows
// Can accept various of criteria, such as limit criteria and etc.
func (r GORMRepository) List(ctx context.Context, criteria ...Criteria) ([]PaymentAggregate, error) {
	var p []PaymentAggregate
	builder := r.conn.Order("created_at desc")

	for _, criteriaItem := range criteria {
		if criteriaItem.Name() == LimitCriteriaName {
			limitCriteria := criteriaItem.(LimitCriteria)

			builder.Offset(limitCriteria.Offset)
			builder.Limit(limitCriteria.Limit)
		}
	}

	err := builder.Find(&p).Error
	if err != nil {
		return nil, fmt.Errorf("failed to execute select statement to return a list of payments: %v", err)
	}

	return p, nil
}
