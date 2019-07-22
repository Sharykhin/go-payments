package api

import (
	"github.com/Sharykhin/go-payments/database"
	"github.com/Sharykhin/go-payments/entity"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	var user entity.User
	database.G.First(&user, 1)
	c.JSON(200, user)
}
