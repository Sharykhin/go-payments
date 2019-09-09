package event

import "time"

const (
	UserCreatedEvent = "UserCreated"
)

type (
	Event struct {
		Time time.Time
		Name string
		Data map[string]interface{}
	}
)

func NewEvent(name string, data map[string]interface{}) Event {
	return Event{
		Time: time.Now().UTC(),
		Name: name,
		Data: data,
	}
}
