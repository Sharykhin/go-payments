package validation

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

// ValidateRequest validates any income request according to the rules
// specified in a request struct.
func ValidateRequest(c *gin.Context, request interface{}) (bool, []string) {
	if err := c.ShouldBindJSON(&request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errors []string
		for _, v := range validationErrors {
			errors = append(errors, fmt.Sprintf("%s %s", v.Field, v.ActualTag))
		}
		return false, errors
	}

	return true, nil
}
