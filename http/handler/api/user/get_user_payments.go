package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/Sharykhin/go-payments/database"
	"github.com/Sharykhin/go-payments/entity"
	"github.com/gin-gonic/gin"
)

func GetUserPayments(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("ID", id)
	user := entity.User{
		ID: int64(id),
	}

	database.G.Preload("Payments", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at DESC")
	}).Find(&user)
	fmt.Println("User", user)
	c.JSON(http.StatusOK, gin.H{"payments": user.Payments})

}
