package payment

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Sharykhin/go-payments/core/logger"
	"github.com/Sharykhin/go-payments/database"
	paymentEntity "github.com/Sharykhin/go-payments/domain/payment/repository/entity"
	userEntity "github.com/Sharykhin/go-payments/domain/user/repository/entity"
	"github.com/Sharykhin/go-payments/http/request/payment"
)

func CreatePayment(c *gin.Context) {
	var r payment.CreateTransactionRequest
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := userEntity.User{ID: r.UserID}
	database.G.Find(&u)

	p := paymentEntity.Payment{
		TransactionID: "123451223",
		User:          u,
		Status:        "Accepted",
		Description:   r.Description,
		Amount:        r.Amount,
		ChargeDate:    time.Now().UTC(),
	}

	logger.Info("Payment Before: %v", p)
	database.G.Save(&p)
	logger.Info("Payment After: %v", p)
	c.JSON(http.StatusCreated, gin.H{"payment": p})

}
