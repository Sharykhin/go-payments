package request

import "github.com/Sharykhin/go-payments/domain/payment/value"

type (
	// TODO: I think this package can be renamed into dto and may include it into another package?
	NewPayment struct {
		Amount      value.Amount
		UserID      int64
		Description string
	}
)
