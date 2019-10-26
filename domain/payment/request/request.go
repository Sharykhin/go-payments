package request

import "github.com/Sharykhin/go-payments/domain/payment/value"

type (
	NewPayment struct {
		Amount      value.Amount
		UserID      int64
		Description string
	}
)
