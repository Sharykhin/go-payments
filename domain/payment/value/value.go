package value

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math/rand"
)

const (
	USD = Currency(iota)
	EUR
	RYB

	StatusOpen = Status(iota)
	StatusCompleted
)

type (
	Currency int
	Status   int

	// Amount is a value object that describes amount of money
	Amount struct {
		Currency Currency
		Value    decimal.Decimal
	}

	TransactionID string
)

// NewAmount creates a new amount value object and it hasn't to be negative
func NewAmount(currency Currency, value decimal.Decimal) Amount {
	if value.IsNegative() {
		panic("amount can not be negative")
	}

	return Amount{
		Currency: currency,
		Value:    value,
	}
}

func NewTransactionID() string {
	// TODO: move it out
	var letter = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 20)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)
}

func (s Status) String() string {
	switch s {
	case StatusOpen:
		return "open"
	case StatusCompleted:
		return "completed"
	default:
		panic(fmt.Sprintf("status %d has invalid value", s))
	}
}
