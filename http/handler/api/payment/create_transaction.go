package payment

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Sharykhin/go-payments/database"
	"github.com/Sharykhin/go-payments/entity"
	"github.com/Sharykhin/go-payments/request/payment"
)

func CreateTransaction(c *gin.Context) {
	var r payment.CreateTransactionRequest
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := entity.User{ID: r.UserID}
	database.G.Find(&u)
	fmt.Println("User", u)
	fmt.Println()
	p := entity.Payment{
		TransactionID: "123451223",
		User:          u,
		Status:        "Accepted",
		Description:   r.Description,
		Amount:        r.Amount,
		ChargeDate:    time.Now().UTC(),
	}

	fmt.Println("Payment Before", p)

	database.G.Save(&p)
	fmt.Println("Payment After", p)
	c.JSON(http.StatusCreated, gin.H{"payment": p})

}
