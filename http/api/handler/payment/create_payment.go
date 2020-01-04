package payment

import (
	"context"
	"strconv"
	"time"

	paymentService "github.com/Sharykhin/go-payments/domain/payment/service"

	"github.com/Sharykhin/go-payments/core/locator"
	"github.com/Sharykhin/go-payments/domain/payment/value"
	"github.com/Sharykhin/go-payments/http"
	"github.com/Sharykhin/go-payments/http/validation"
	"github.com/shopspring/decimal"

	identityEntity "github.com/Sharykhin/go-payments/domain/identity/entity"
	pr "github.com/Sharykhin/go-payments/http/request/payment"
	"github.com/gin-gonic/gin"
)

// CreatePayment is a handler that services creating a new payment transaction endpoint
func CreatePayment(c *gin.Context) {
	userID, err := getRouteParamAsInt64(c, "id")
	if err != nil {
		http.BadRequest(c, http.Errors{"route parameter :id is invalid"})
		return
	}

	authUser, ok := c.Value(http.UserContext).(identityEntity.UserContext)
	if !ok {
		panic("create payment endpoint must be under auth middleware to consume identityEntity.UserContext")
	}

	if authUser.ID != int64(userID) {
		http.Forbidden(c)
		return
	}

	var req pr.CreateTransactionRequest
	if isValid, err := validation.ValidateRequest(c, &req); !isValid {
		http.BadRequest(c, http.Errors(err))
		return
	}

	service := locator.GetPaymentService()
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(15*time.Second))
	defer cancel()

	p, err := service.Create(
		ctx,
		paymentService.NewPaymentRequest{
			Amount:      value.NewAmount(value.USD, decimal.NewFromFloat(req.Amount)),
			Description: req.Description,
			UserID:      int64(userID),
		},
	)

	if err != nil {
		http.BadRequest(c, http.Errors{err.Error()})
		return
	}

	http.Created(c, http.Data{
		"Payment": p,
	}, nil)

}

func getRouteParamAsInt64(c *gin.Context, param string) (int64, error) {
	val, err := strconv.Atoi(c.Param(param))
	if err != nil {
		return 0, err
	}

	return int64(val), nil
}
