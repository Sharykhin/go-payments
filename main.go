package main

import (
	"github.com/Sharykhin/go-payments/http"
	"log"
)

func main() {
	log.Fatal(http.ListenAndServe())
}
