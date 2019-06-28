package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	fmt.Println("SERVER_ADDR", os.Getenv("SERVER_ADDR"))
	log.Fatal(r.Run(os.Getenv("SERVER_ADDR")))
}
