package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"

	"github.com/Sharykhin/go-payments/core/event"
)

const (
	exchangeName = "payment-events"
)

var (
	ch *amqp.Channel
	Q  *Queue
)

type (
	Queue struct {
		ch     *amqp.Channel
		events map[string]struct{}
	}
)

func init() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failled to connect to rabbitmq: %v", err)
	}

	//defer conn.Close()

	ch, err = conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	//defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName, // name
		"fanout",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}

	Q = &Queue{
		ch: ch,
	}

}

func (q *Queue) Subscribe(tag, eventName string, fn func(e event.Event)) error {
	q.events[eventName] = struct{}{}

	qd, err := q.ch.QueueInspect(tag)
	if err == nil {
		return nil
	}

	qd, err = q.ch.QueueDeclare(
		tag,   // name
		false, // durable
		false, // delete when usused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		return fmt.Errorf("failed to declare a queue: %v", err)
	}

	err = q.ch.QueueBind(
		qd.Name,      // queue name
		"",           // routing key
		exchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to bind a queue: %v", err)
	}

	msgs, err := q.ch.Consume(
		qd.Name, // queue
		"",      // consumer
		true,    // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)

	if err != nil {
		return fmt.Errorf("failed to consume messages: %v", err)
	}

	go func() {
		for d := range msgs {
			var ev event.Event
			err := json.Unmarshal(d.Body, &ev)
			if err != nil {
				log.Printf("failed to parse income message: %v", err)
			}
			if _, ok := q.events[ev.Name]; ok {
				log.Printf("Consumer [%s] Received a message: %s", ev)
				fn(ev)
			}
		}
	}()

	return nil
}

func (q *Queue) RaiseEvent(e event.Event) error {
	b, err := json.Marshal(&e)
	if err != nil {
		return fmt.Errorf("failed to marshal event into json: %v", err)
	}

	err = q.ch.Publish(
		exchangeName, // exchange
		"",           // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		})

	if err != nil {
		return fmt.Errorf("failed to publish an event: %v", err)
	}

	return nil
}

func NewQueue() *Queue {
	return &Queue{
		ch:     ch,
		events: map[string]struct{}{},
	}
}
