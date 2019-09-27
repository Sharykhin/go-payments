package entity

import "github.com/shopspring/decimal"

type (
	Payment struct {
		Amount      decimal.Decimal `json:"amount"`
		Description string          `json:"description"`
	}
)
