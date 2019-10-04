package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/Sharykhin/go-payments/core/event"
	"github.com/Sharykhin/go-payments/core/queue"
)

func main() {
	q := queue.New(queue.RabbitMQ)

	err := q.Subscribe("notification", event.UserRegisteredEvent, func(e event.Event) {
		fmt.Println("User Registered", e)
	})

	if err != nil {
		log.Fatalf("faield to subscribe on event: %v", err)
	}

	err = q.Subscribe("notification", event.UserSignIn, func(e event.Event) {
		log.Println("HA HA HA I KNOW I COULD", e.Data, e.Time)
	})
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	<-quit
	fmt.Printf("Server should be gracefully shutdown!!")
}
