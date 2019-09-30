package queue

import (
	"github.com/Sharykhin/go-payments/core/event"
	"github.com/Sharykhin/go-payments/core/queue/rabbitmq"
)

const (
	RabbitMQ     = iota
	DefaultQueue = RabbitMQ

	TagNotification = iota
)

type (
	Publisher interface {
		RaiseEvent(e event.Event) error
	}

	Subscriber interface {
		Subscribe(tag, eventName string, fn func(e event.Event)) error
	}

	QueueManager interface {
		Subscriber
		Publisher
	}
)

func New(t int) QueueManager {
	switch t {
	case RabbitMQ:
		return rabbitmq.NewQueue()
	default:
		panic("invalid queue type")
	}
}

//Default returns a default queue manager
func Default() QueueManager {
	switch DefaultQueue {
	case RabbitMQ:
		return rabbitmq.NewQueue()
	default:
		panic("invalid queue type")
	}
}
