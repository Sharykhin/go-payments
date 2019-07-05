package payment

import (
	"fmt"
	"log"

	"github.com/Sharykhin/go-payments/request/payment"
	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var r payment.CreateTransactionRequest
	err := c.ShouldBind(&r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
	c.JSON(200, nil)
}
