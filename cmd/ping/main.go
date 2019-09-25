package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Sharykhin/go-payments/core/logger"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sig := <-sigs
		log.Printf("Ping server is going to be closed gracefully because of a signal: %v\n", sig)
		cancel()
		done <- true
	}()

	go func(ctx context.Context) {
		for {
			select {
			case <-time.After(60 * time.Second):
				logger.Log.Info("PING")
			case <-ctx.Done():
				logger.Log.Info("Ping controller is closed.")
				return
			}
		}
	}(ctx)

	<-done
}
