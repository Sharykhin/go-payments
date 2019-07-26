package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type (
	Payment struct {
		ID            int64 `gorm:"primary_key"`
		TransactionID string
		User          User `gorm:"foreignkey:User,association_foreignkey:ID"`
		//UserID        int64
		Amount      decimal.Decimal
		Description string
		Status      string
		ChargeDate  time.Time `gorm:"column:created_at"`
	}
)

func (p Payment) TableName() string {
	return "transactions"
}
