package payment

import "github.com/shopspring/decimal"

type (
	// CreateTransactionRequest represents http request body to create a new payment transaction
	CreateTransactionRequest struct {
		UserID      int64           `json:"UserID" binding:"required"`
		Amount      decimal.Decimal `json:"Amount" binding:"required"`
		Description string          `json:"Description" binding:"required"`
	}
)
