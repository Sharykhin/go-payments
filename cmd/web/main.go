package main

import (
	"log"

	http "github.com/Sharykhin/go-payments/http/web"
)

func main() {
	log.Fatal(http.ListenAndServe())
}
