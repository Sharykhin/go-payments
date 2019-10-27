package payment

import (
	"github.com/Sharykhin/go-payments/domain/payment/model"
	"github.com/gin-gonic/gin"

	"github.com/Sharykhin/go-payments/core/locator"
	"github.com/Sharykhin/go-payments/http"
)

func GetUserPayments(c *gin.Context) {
	service := locator.GetPaymentService()
	payments, err := service.All(c.Request.Context())
	if err != nil {
		http.ServerError(c, http.Errors{err.Error()})
		return
	}

	var vm []model.PaymentView
	for _, p := range payments {
		vm = append(vm, model.NewPaymentViewModel(p, "list"))
	}

	http.OK(c, map[string]interface{}{
		"Payments": vm,
	}, nil)
}
