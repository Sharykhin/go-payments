package main

import (
	"github.com/Sharykhin/go-payments/core/locator"
	"log"

	http "github.com/Sharykhin/go-payments/http/api"
)

func main() {
	sl := locator.NewServiceLocator()
	log.Fatal(http.ListenAndServe(sl))
}
