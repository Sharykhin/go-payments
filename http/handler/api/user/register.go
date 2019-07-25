package user

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"

	"gopkg.in/go-playground/validator.v8"

	"github.com/Sharykhin/go-payments/database"
	"github.com/Sharykhin/go-payments/entity"
	"golang.org/x/crypto/bcrypt"

	"github.com/Sharykhin/go-payments/request"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var rr request.RegisterRequest

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

	// TODO: heck
	hash, err := bcrypt.GenerateFromPassword([]byte(rr.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}

	//TODO: heck
	user := entity.User{
		FirstName: rr.FirstName,
		LastName:  rr.LastName,
		Email:     rr.Email,
		Password:  string(hash),
		DeletedAt: entity.NullTime{
			Valid: false,
		},
	}
	fmt.Println("user", user)
	database.G.Save(&user)
	session := sessions.Default(c)
	session.Set("user_id", user.ID)

	c.JSON(http.StatusCreated, gin.H{"user": user})
}
