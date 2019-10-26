package model

import (
	"encoding/json"

	"github.com/Sharykhin/go-payments/core/type"
	"github.com/Sharykhin/go-payments/domain/payment/value"
)

type (
	// Payment describes domain model
	Payment struct {
		id          int64
		amount      value.Amount
		description string
		user        UserInterface
		createdAt   types.Time
	}
)

func (p *Payment) SetID(ID int64) *Payment {
	p.id = ID
	return p
}

func (p *Payment) SetAmount(amount value.Amount) *Payment {
	p.amount = amount
	return p
}

func (p *Payment) SetDescription(description string) *Payment {
	p.description = description

	return p
}

func (p *Payment) SetCreatedAt(date types.Time) *Payment {
	p.createdAt = date
	return p
}

func (p *Payment) SetUser(user UserInterface) *Payment {
	p.user = user
	return p
}

// MarshalJSON implements json.Marshaler interface
func (p *Payment) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID          int64       `json:"ID"`
		Amount      string      `json:"amount"`
		Description string      `json:"description"`
		CreatedAt   types.Time  `json:"CreatedAt"`
		User        interface{} `json:"User"`
	}{
		ID:          p.id,
		Amount:      p.amount.Value.String(),
		Description: p.description,
		CreatedAt:   p.createdAt,
		User: struct {
			ID    int64  `json:"ID"`
			Email string `json:"Email"`
		}{
			ID:    p.user.GetID(),
			Email: p.user.GetEmail(),
		},
	})
}
