package main

import (
	"github.com/Sharykhin/go-payments/database"
	"github.com/Sharykhin/go-payments/entity"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		var user entity.User
		database.Conn.First(&user, 1)
		c.JSON(200,user)
	})

	log.Fatal(r.Run(os.Getenv("SERVER_ADDR")))
}
