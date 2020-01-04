package payment

import (
	"github.com/gin-gonic/gin"

	"github.com/Sharykhin/go-payments/core/locator"
	"github.com/Sharykhin/go-payments/domain/payment/model"
	"github.com/Sharykhin/go-payments/http"
)

// GetUserPayments returns list of user's payments
func GetUserPayments(c *gin.Context) {
	service := locator.GetPaymentService()
	payments, err := service.LimitedList(c.Request.Context(), 0, 10)
	if err != nil {
		http.ServerError(c, http.Errors{err.Error()})
		return
	}

	var vm []model.PaymentView
	for _, p := range payments {
		vm = append(vm, *p.ViewModel("list"))
	}

	http.OK(c, map[string]interface{}{
		"Payments": vm,
	}, nil)
}
