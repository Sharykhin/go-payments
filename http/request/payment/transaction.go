package payment

type (
	// CreateTransactionRequest represents http request body to create a new payment transaction
	CreateTransactionRequest struct {
		UserID      int64   `json:"UserID" binding:"required"`
		Amount      float64 `json:"Amount" binding:"required"`
		Description string  `json:"Description" binding:"required"`
	}
)
