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
		StatusCode int  `json:"StatusCode"`
		Data       Data `json:"Data"`
		Meta       Meta `json:"Meta"`
	}
	Data map[string]interface{}
	Meta map[string]interface{}
)

func OK(c *gin.Context, data Data, meta Meta) {
	response := newResponse(httpCodes.StatusOK, data, meta)
	c.JSON(httpCodes.StatusOK, response)
}

func Created(c *gin.Context, data Data, meta Meta) {
	response := newResponse(httpCodes.StatusCreated, data, meta)
	c.JSON(httpCodes.StatusCreated, response)
}

func BadRequest(c *gin.Context, data Data) {
	response := newResponse(httpCodes.StatusBadRequest, data, nil)
	c.JSON(httpCodes.StatusBadRequest, response)
}

func newResponse(code int, data Data, meta Meta) Response {
	if meta == nil {
		meta = Meta{}
	}
	meta["ServerTimestamp"] = time.Now().Format(core.ISO8601)
	return Response{
		StatusCode: code,
		Data:       data,
		Meta:       meta,
	}
}
