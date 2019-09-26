package api

import (
	"fmt"
	"github.com/Sharykhin/go-payments/core/event"
	"github.com/Sharykhin/go-payments/core/queue"
	"github.com/gin-gonic/gin"

	"github.com/Sharykhin/go-payments/http"
)

func Ping(c *gin.Context) {

	q := queue.New(queue.RabbitMQ)

	err := q.RaiseEvent(event.NewEvent(event.UserCreatedEvent, event.Payload{
		"foo": "bar",
		"baz": 12,
	}))
	if err != nil {
		fmt.Printf("failed to raise event: %v", err)
	}
	http.OK(c, map[string]interface{}{
		"Message": "pong",
	}, nil)
}
