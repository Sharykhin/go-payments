package repository

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/Sharykhin/go-payments/domain/user/repository/entity"
)

type (
	Payment struct {
		ID            int64 `gorm:"primary_key"`
		TransactionID string
		User          entity.User `gorm:"association_autoupdate:false" json:"-" `
		UserID        int64
		Amount        decimal.Decimal
		Description   string
		Status        string
		ChargeDate    time.Time `gorm:"column:created_at"`
	}
)

func (p Payment) TableName() string {
	return "transactions"
}
