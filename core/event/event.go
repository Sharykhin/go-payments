package event

import "time"

const (
	UserCreatedEvent    = "UserCreated"
	PaymentCreatedEvent = "PaymentCreated"
)

type (
	Event struct {
		Time time.Time
		Name string
		Data map[string]interface{}
	}

	Payload map[string]interface{}
)

func NewEvent(name string, data Payload) Event {
	return Event{
		Time: time.Now().UTC(),
		Name: name,
		Data: data,
	}
}
