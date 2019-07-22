package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.JSON(http.StatusCreated, `{"user":"1"}`)
}
