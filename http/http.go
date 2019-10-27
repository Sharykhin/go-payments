package http

import (
	httpCodes "net/http"
	"time"

	"github.com/Sharykhin/go-payments/core"

	"github.com/gin-gonic/gin"
)

const (
	UserContext = "UserContext"
)

type (
	Response struct {
		StatusCode int    `json:"StatusCode"`
		Errors     Errors `json:"Errors"`
		Data       Data   `json:"Data"`
		Meta       Meta   `json:"Meta"`
	}
	Errors []string
	Data   map[string]interface{}
	Meta   map[string]interface{}
)

func OK(c *gin.Context, data Data, meta Meta) {
	response := newResponse(httpCodes.StatusOK, data, meta, nil)
	c.JSON(httpCodes.StatusOK, response)
}

func Created(c *gin.Context, data Data, meta Meta) {
	response := newResponse(httpCodes.StatusCreated, data, meta, nil)
	c.JSON(httpCodes.StatusCreated, response)
}

func BadRequest(c *gin.Context, errors Errors) {
	response := newResponse(httpCodes.StatusBadRequest, nil, nil, errors)
	c.JSON(httpCodes.StatusBadRequest, response)
}

func ServerError(c *gin.Context, errors Errors) {
	response := newResponse(httpCodes.StatusInternalServerError, nil, nil, errors)
	c.JSON(httpCodes.StatusInternalServerError, response)
}

func newResponse(code int, data Data, meta Meta, errors Errors) Response {
	if meta == nil {
		meta = Meta{}
	}
	meta["ServerTimestamp"] = time.Now().Format(core.ISO8601)
	return Response{
		StatusCode: code,
		Data:       data,
		Meta:       meta,
		Errors:     errors,
	}
}
