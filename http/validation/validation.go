package validation

import (
	"fmt"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

// ValidateRequest validates any income request according to the rules
// specified in a request struct.
func ValidateRequest(c *gin.Context, request interface{}) (bool, []string) {
	err := c.ShouldBindJSON(request)

	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errors []string
			for _, v := range validationErrors {
				errors = append(
					errors,
					strings.TrimSpace(fmt.Sprintf("%s %s %s", v.Field, v.ActualTag, v.Param)),
				)
			}
			return false, errors
		}
		if err == io.EOF {
			return false, []string{"Request Body is Empty"}
		}

		return false, []string{err.Error()}
	}

	return true, nil
}
