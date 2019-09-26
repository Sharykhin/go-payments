package payment

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/Sharykhin/go-payments/database"
	userEntity "github.com/Sharykhin/go-payments/domain/user/repository/entity"
)

func GetUserPayments(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := userEntity.User{
		ID: int64(id),
	}

	database.G.Preload("Payments", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at DESC")
	}).Find(&user)
	fmt.Println("User", user)
	c.JSON(http.StatusOK, gin.H{"payments": user.Payments})
}
