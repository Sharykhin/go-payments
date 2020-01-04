package model

import "github.com/shopspring/decimal"

type (
	Payment struct {
		Amount      decimal.Decimal
		Description string
	}
)
