package main

import (
	"log"

	http "github.com/Sharykhin/go-payments/http/api"
)

func main() {
	log.Fatal(http.ListenAndServe())
}
