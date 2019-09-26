package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validator "gopkg.in/go-playground/validator.v8"

	"github.com/Sharykhin/go-payments/core/logger"
	"github.com/Sharykhin/go-payments/core/type"
	userEntity "github.com/Sharykhin/go-payments/domain/user/repository/entity"
	"github.com/Sharykhin/go-payments/domain/user/service"
	ur "github.com/Sharykhin/go-payments/http/request/user"
)

// Register
func Register(c *gin.Context) {
	var rr ur.RegisterRequest

	// TODO: heck
	if err := c.ShouldBindJSON(&rr); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errors := make(map[string]string, len(validationErrors))
		for _, v := range validationErrors {
			errors[v.Name] = v.ActualTag
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errors})
		return
	}

	userService := service.NewUserService()

	user := userEntity.User{
		FirstName: rr.FirstName,
		LastName:  rr.LastName,
		Email:     rr.Email,
		DeletedAt: types.NullTime{
			Valid: false,
		},
	}

	newUser, err := userService.Create(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("user: %v", newUser)

	c.JSON(http.StatusCreated, gin.H{"user": user})
}
