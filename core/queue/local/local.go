package local

import "github.com/Sharykhin/go-payments/core/event"

type (
	Queue struct {
		events      []event.Event
		subscribers map[string][]func(e event.Event)
	}
)

func (q *Queue) RaiseEvent(e event.Event) {
	q.events = append(q.events, e)
}

func (q *Queue) Subscribe(name string, fn func(e event.Event)) {
	q.subscribers[name] = append(q.subscribers[name], fn)
}

func (q *Queue) ReleaseEvents() {
	for _, e := range q.events {
		if subscribers, ok := q.subscribers[e.Name]; ok {
			for _, fn := range subscribers {
				go fn(e)
			}
		}

	}
}

func NewQueue() *Queue {
	return &Queue{
		events:      []event.Event{},
		subscribers: make(map[string][]func(e event.Event)),
	}
}
