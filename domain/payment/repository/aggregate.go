package repository

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/Sharykhin/go-payments/domain/user/repository"
)

type (
	//PaymentAggregate is an aggregate model on a repository level. Right now it has GORM implementation
	PaymentAggregate struct {
		ID            int64           `gorm:"primary_key"`
		TransactionID string          `gorm:"column:transaction_id"`
		User          repository.User `gorm:"association_autoupdate:false" json:"-" `
		UserID        int64           `gorm:"column:user_id"`
		Amount        decimal.Decimal `gorm:"column:amount"`
		Description   string          `gorm:"column:description"`
		Status        string          `gorm:"column:status"`
		ChargeDate    time.Time       `gorm:"column:created_at"`
	}
)

// TableName is a method that GORM uses to identify table name
func (p PaymentAggregate) TableName() string {
	return "transactions"
}
