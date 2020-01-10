package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"

	"github.com/Sharykhin/go-payments/core/deferrer"
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
	defer func() {
		if err := recover(); err != nil {
			deferrer.Flush()
			log.Fatalf("Terminating: %v", err)
		}
	}()

	conn, err := amqp.Dial(fmt.Sprintf(
		"amqp://%s:%s@%s:5672/",
		os.Getenv("RABBITMQ_USER"),
		os.Getenv("RABBITMQ_PASS"),
		os.Getenv("RABBITMQ_HOST"),
	))
	if err != nil {
		log.Panicf("failed to connect to RabbitMQ: %v", err)
	}
	log.Println("Connected to RabbitMQ")

	deferrer.QueueDeffer(func() {
		log.Println("Gracefully closing RabbitMQ connection")
		err := conn.Close()
		if err != nil {
			log.Printf("failed to close RabbitMQ connection: %v", err)
		}
	})

	ch, err = conn.Channel()
	if err != nil {
		log.Panicf("Failed to open a channel: %v", err)
	}

	deferrer.QueueDeffer(func() {
		log.Println("Gracefully closing RabbitMQ channel")
		err := ch.Close()
		if err != nil {
			log.Printf("failed to close rabbitmq channel: %v", err)
		}
	})

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
		log.Panicf("Failed to declare an exchange: %v", err)
	}

	Q = &Queue{
		ch: ch,
	}

}

// Subscribe subscribes on an specific event with providing tag
// that in scope of RabbitMQ will create a new queue.
// This is done to allow multiple instances of the same service not to
// get the same message.
// TODO: think about using Subscribe not as method that creates goroutine but a sync one and run it in goroutine, or at least we should use context or channel to close goroutine
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
