package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Sharykhin/go-payments/core"
	"github.com/Sharykhin/go-payments/core/event"
	"github.com/Sharykhin/go-payments/core/locator"
	"github.com/Sharykhin/go-payments/core/logger"
)

func main() {
	sl := locator.NewServiceLocator()
	subscriber := sl.GetSubscriberService()

	err := subscriber.Subscribe("notification", event.UserRegisteredEvent, func(e event.Event) {
		fmt.Println("User Registered", e)
	})

	if err != nil {
		log.Fatalf("faield to subscribe on event: %v", err)
	}

	err = subscriber.Subscribe("notification", event.UserSignIn, func(e event.Event) {
		log.Println("Goi event UserSingIn", e.Data, e.Time)

		identityService := sl.GetIdentityService()
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1*time.Minute))
		defer cancel()
		t, _ := time.Parse(core.ISO8601, e.Data["LoginAt"].(string))

		log.Printf("HA: %T %v %T %v", e.Data["UserID"], e.Data["UserID"], e.Data["LoginAt"], e.Data["LoginAt"])
		err := identityService.UpdateLastLogin(ctx, int64(e.Data["UserID"].(float64)), t)
		if err != nil {
			logger.Error("could not update users's last login: %v", err)
		}
	})

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	<-quit
	fmt.Printf("Server should be gracefully shutdown!!")
}
