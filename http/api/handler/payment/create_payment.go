package payment

import (
	"github.com/Sharykhin/go-payments/core/locator"
	"github.com/Sharykhin/go-payments/domain/payment/request"
	"github.com/Sharykhin/go-payments/domain/payment/value"
	"github.com/Sharykhin/go-payments/http"
	"github.com/Sharykhin/go-payments/http/validation"

	"github.com/Sharykhin/go-payments/http/request/payment"
	"github.com/gin-gonic/gin"
)

func CreatePayment(c *gin.Context) {
	var req payment.CreateTransactionRequest
	if isValid, err := validation.ValidateRequest(c, &req); !isValid {
		http.BadRequest(c, http.Errors(err))
		return
	}

	service := locator.GetPaymentService()

	p, err := service.Create(c.Request.Context(), request.NewPayment{
		Amount:      value.NewAmount(value.USD, req.Amount),
		Description: req.Description,
		UserID:      req.UserID,
	})

	if err != nil {
		http.BadRequest(c, http.Errors{err.Error()})
		return
	}

	http.Created(c, http.Data{
		"Payment": p,
	}, nil)

}
