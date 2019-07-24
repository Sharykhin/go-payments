package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"

	"golang.org/x/crypto/bcrypt"

	"github.com/Sharykhin/go-payments/database"
	"github.com/Sharykhin/go-payments/entity"

	"github.com/Sharykhin/go-payments/request"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var lr request.LoginRequest
	if err := c.BindJSON(&lr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user entity.User
	database.G.Where("email = ?", lr.Email).First(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(lr.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wring"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": &user})
}
