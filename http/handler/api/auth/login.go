package auth

import (
	"fmt"
	"net/http"

	"github.com/Sharykhin/go-payments/database"
	"github.com/Sharykhin/go-payments/entity"
	"github.com/Sharykhin/go-payments/request"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
		if err == bcrypt.ErrMismatchedHashAndPassword {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Credentials are invalid"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	err := session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wring"})
		return
	}
	c.SetCookie("pizda", "da", 1200, "/", "", true, true)

	fmt.Println("session before", session.Get("user_id"))
	c.JSON(http.StatusOK, gin.H{"user": &user})
}
