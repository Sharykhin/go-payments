package queue

import (
	"github.com/Sharykhin/go-payments/core/event"
	"github.com/Sharykhin/go-payments/core/queue/local"
	"github.com/Sharykhin/go-payments/core/queue/rabbitmq"
)

const (
	TypeLocal = iota
	RabbitMQ
)

type (
	Publisher interface {
		RaiseEvent(e event.Event) error
	}

	Subscriber interface {
		Subscribe(name string, fn func(e event.Event)) error
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
	case RabbitMQ:
		return rabbitmq.NewQueue()
	default:
		panic("invalid queue type")
	}
}
