package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

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

type EventCallback func(e event.Event)

type (
	Queue struct {
		ch     *amqp.Channel
		events map[string][]EventCallback
	}
)

func init() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("failed to connect to rabbitmq: %v", err)
	}
	log.Println("Connected to RabbitMQ")

	//defer conn.Close()

	ch, err = conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	//defer ch.Close()

	log.Println("Created a channel")

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

	if _, ok := q.events[eventName]; !ok {
		q.events[eventName] = []EventCallback{}
	}
	q.events[eventName] = append(q.events[eventName], fn)

	// TODO: need to think how not to declare queue each time Subscribe is called
	//qd, err := q.ch.QueueInspect(tag)
	//log.Println("Error:", err)
	//if err == nil {
	//	return nil
	//}

	qd, err := q.ch.QueueDeclare(
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
			log.Printf("Got message from rabbitmq: %s, %v", ev.Name, ev.Data)
			if _, ok := q.events[ev.Name]; ok {
				log.Printf("Raise all callbacks for event name %s: %v", ev.Name, q.events[ev.Name])
				for _, fn := range q.events[ev.Name] {
					go fn(ev)
				}
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
		events: map[string][]EventCallback{},
	}
}
