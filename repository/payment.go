package repository

import (
	"context"
	"fmt"

	"github.com/Sharykhin/go-payments/database"

	"github.com/Sharykhin/go-payments/entity"
	"github.com/jinzhu/gorm"
)

// Type is a custom semantical type of a concrete repository that factory would return
type Type int64

const (
	GORM = Type(iota)
)

type (
	// PaymentRepository describes general method with storage layer around payment domain
	PaymentRepository interface {
		Create(ctx context.Context, payment entity.Payment) (*entity.Payment, error)
		Find(ctx context.Context, ID int64) (*entity.Payment, error)
	}
	// GormPaymentRepository is a concrete implementation of PaymentRepository
	// that uses Gorm ORM for managing records
	GormPaymentRepository struct {
		conn *gorm.DB
	}
)

// Create create a new record of payment and return a reference to that record
// If an error is occurred return nul and error itself
func (r GormPaymentRepository) Create(ctx context.Context, payment entity.Payment) (*entity.Payment, error) {
	err := r.conn.Create(&payment).Error
	if err != nil {
		return nil, fmt.Errorf("failed to insert a new payment row: %v", err)
	}

	return &payment, nil
}

// Find finds a record by its id. If an error is occurred function checks whether
// a corresponding row was not found and if so return repository not found error
func (r GormPaymentRepository) Find(ctx context.Context, ID int64) (*entity.Payment, error) {
	p := &entity.Payment{
		ID: ID,
	}

	err := r.conn.Find(&p).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, NotFound
		}

		return nil, fmt.Errorf("could not execute select statement to find a payment: %v", err)
	}

	return p, nil
}

// PaymentRepositoryFactory creates an appropriate repository instance
// that satisfies an interface
func PaymentRepositoryFactory(t Type) PaymentRepository {
	switch t {
	case GORM:
		return GormPaymentRepository{
			conn: database.G,
		}
	default:
		panic(fmt.Errorf("invalid repository type"))
	}
}
