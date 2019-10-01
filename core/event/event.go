package event

import "time"

const (
	// TODO: replace with iota
	UserCreatedEvent                = "UserCreated"
	PaymentCreatedEvent             = "PaymentCreated"
	UserRegisteredEvent             = "UserRegisteredEvent"
	UserPasswordCreationFailedEvent = "UserPasswordCreationFailedEvent"
	UserPasswordCreatedEvent        = "UserPasswordCreatedEvent"
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
