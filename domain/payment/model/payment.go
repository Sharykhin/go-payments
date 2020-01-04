package model

import (
	"encoding/json"
	"github.com/Sharykhin/go-payments/core/file"
	"os"

	types "github.com/Sharykhin/go-payments/core/type"
	"github.com/Sharykhin/go-payments/domain/payment/value"
)

type (
	// Payment describes the main domain model
	Payment struct {
		id          int64
		amount      value.Amount
		description string
		user        UserInterface
		createdAt   types.Time
		files       []file.FileURL
	}

	// TODO: payment view doesn't look really handy, think about how it should be changed if that possible
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

// NewPayment returns a new instance of Payment model
func NewPayment(
	ID int64,
	Amount value.Amount,
	Description string,
	CreatedAt types.Time,
	User UserInterface,
) *Payment {

	return &Payment{
		id:          ID,
		amount:      Amount,
		description: Description,
		createdAt:   CreatedAt,
		user:        User,
	}
}

// MarshalJSON implements json.Marshaler interface
// TODO: since we are using view model concept this method can be removed?
// TODO: eventually I like an idea of marshaling struct
func (p *Payment) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID          int64       `json:"ID"`
		Amount      string      `json:"Amount"`
		Description string      `json:"Description"`
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

// ViewModel is a view representation of Payment model
func (p *Payment) ViewModel(view string) *PaymentView {
	vm := &PaymentView{
		ID:          p.id,
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

func (p *Payment) AttachFile(file *os.File) {

}
