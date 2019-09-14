package payment

import "github.com/shopspring/decimal"

type (
	CreateTransactionRequest struct {
		UserID      int64           `form:"user_id"`
		Amount      decimal.Decimal `form:"amount"`
		Description string          `form:"description"`
	}
)
