package repository

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"

	GORMDB "github.com/Sharykhin/go-payments/core/database/gorm"
	"github.com/Sharykhin/go-payments/domain/payment/repository/entity"
)

type (
	// GORMRepository struct responsible for working with
	// storage layer by using GORM ORM
	GORMRepository struct {
		conn *gorm.DB
	}
)

// Create creates a new payment in a database and returns just created record
func (r GORMRepository) Create(cxt context.Context, payment entity.Payment) (*entity.Payment, error) {
	err := r.conn.Create(&payment).Error
	if err != nil {
		return nil, fmt.Errorf("failed to insert a new user row: %v", err)
	}

	return &payment, nil
}

// NewGORMRepository is a constructor function that returns a new instance of GORMRepository
// and satisfies PaymentRepository interface
func NewGORMRepository() *GORMRepository {
	return &GORMRepository{
		conn: GORMDB.G,
	}
}
