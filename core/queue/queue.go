package queue

import (
	"github.com/Sharykhin/go-payments/core/event"
	"github.com/Sharykhin/go-payments/core/queue/local"
)

func init() {
	publisher = NewPublisher(TypeLocal)
}

const (
	TypeLocal = iota
)

var (
	publisher Publisher
)

type (
	Publisher interface {
		RaiseEvent(e event.Event)
	}

	Subscriber interface {
		Subscribe(name string, fn func(e event.Event))
	}

	QueueManager interface {
		Subscriber
		Publisher
	}
)

func New(t int) QueueManager {
	switch t {
	case TypeLocal:
		return local.NewQueue()
	default:
		panic("invalid queue type")
	}
}

func NewPublisher(t int) Publisher {
	switch t {
	case TypeLocal:
		return local.NewQueue()
	default:
		panic("invalid queue type")
	}
}

func NewSubscriber(t int) Subscriber {
	switch t {
	case TypeLocal:
		return local.NewQueue()
	default:
		panic("invalid queue type")
	}
}
