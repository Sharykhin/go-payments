package model

import (
	"encoding/json"

	types "github.com/Sharykhin/go-payments/core/type"
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
	// TODO: payment view doesn't look reallly handy, think about how it should be changed if that possible
	// PaymentView represents how payment transaction
	// should be serialized in json
	PaymentView struct {
		ID          int64      `json:"ID"`
		Amount      string     `json:"Amount"`
		Description string     `json:"Description"`
		User        *UserView  `json:"User,omitempty"`
		CreatedAt   types.Time `json:"CreatedAt"`
	}

	// UserView represents user in a payment context
	UserView struct {
		ID    int64  `json:"ID"`
		Email string `json:"Email"`
	}
)

func (p *Payment) SetID(ID int64) *Payment {
	p.id = ID
	return p
}

func (p *Payment) GetID() int64 {
	return p.id
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
		User: &struct {
			ID    int64  `json:"ID"`
			Email string `json:"Email"`
		}{
			ID:    p.user.GetID(),
			Email: p.user.GetEmail(),
		},
	})
}

func (p *Payment) ViewModel() ([]byte, error) {
	return json.Marshal(struct {
		ID          int64      `json:"ID"`
		Amount      string     `json:"amount"`
		Description string     `json:"description"`
		CreatedAt   types.Time `json:"CreatedAt"`
	}{
		ID:          p.id,
		Amount:      p.amount.Value.String(),
		Description: p.description,
		CreatedAt:   p.createdAt,
	})
}

func NewPaymentViewModel(p Payment, view string) PaymentView {
	vm := PaymentView{
		ID:          p.GetID(),
		Amount:      p.amount.Value.String(),
		CreatedAt:   p.createdAt,
		Description: p.description,
	}

	if view == "list" {
		return vm
	}

	vm.User = &UserView{
		ID:    p.user.GetID(),
		Email: p.user.GetEmail(),
	}
	return vm
}
