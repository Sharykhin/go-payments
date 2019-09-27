package auth

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/Sharykhin/go-payments/core/database"
	"github.com/Sharykhin/go-payments/domain/user/repository/entity"
	identityRequest "github.com/Sharykhin/go-payments/http/request/identity"
)

func Login(c *gin.Context) {
	var lr identityRequest.LoginRequest
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().UTC().Add(1 * time.Second).Unix(),
	})

	tokenStr, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Printf("failed to create JWT token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": &user, "token": tokenStr})
}
