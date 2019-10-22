package event

import "time"

const (
	// TODO: replace with iota
	UserCreatedEvent                = "UserCreated"
	PaymentCreatedEvent             = "PaymentCreated"
	UserRegisteredEvent             = "UserRegisteredEvent"
	UserPasswordCreationFailedEvent = "UserPasswordCreationFailedEvent"
	UserPasswordCreatedEvent        = "UserPasswordCreatedEvent"
	UserSignIn                      = "UserSignIn"
)

type (
	// Event represents general event in the application
	// that can be transferred across different services
	Event struct {
		Time time.Time
		Name string
		Data map[string]interface{}
	}
	//Payload is a body payload of a queue message
	Payload map[string]interface{}

	// TODO: experiment not of using Payload map but interface of concrete payload for each event
	UserLoginPayload struct {
		UserID  int64
		LoginAt time.Time
	}
)

// NewEvent is a function construction to return
// a new instance of Event
func NewEvent(name string, data Payload) Event {
	return Event{
		Time: time.Now().UTC(),
		Name: name,
		Data: data,
	}
}
