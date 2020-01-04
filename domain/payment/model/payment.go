package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Sharykhin/go-payments/core/errors"
	"github.com/Sharykhin/go-payments/core/event"
	"github.com/Sharykhin/go-payments/core/file"
	"github.com/Sharykhin/go-payments/core/logger"
	"github.com/Sharykhin/go-payments/core/queue"
	"github.com/Sharykhin/go-payments/core/queue/rabbitmq"
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
		files       []string

		fileUploader file.Uploader
		dispatcher   queue.Publisher
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

	fileUploader file.Uploader,
	dispatcher queue.Publisher,
) *Payment {

	p := Payment{
		id:          ID,
		amount:      Amount,
		description: Description,
		createdAt:   CreatedAt,
		user:        User,

		fileUploader: fileUploader,
		dispatcher:   dispatcher,
	}

	p.dispatcher = rabbitmq.NewQueue()

	return &p
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

func (p *Payment) AttachFile(ctx context.Context, f *os.File) error {
	//validate income file
	info, err := f.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file statistic: %v", err)
	}

	kilobytes := info.Size() / 1024
	megabytes := float64(kilobytes / 1024)

	if megabytes > 5.00 {
		return errors.FileIsTooBig
	}

	contentType, err := file.GetFileContentType(f)
	if err != nil {
		return fmt.Errorf("failed to get file content type: %v", err)
	}

	if contentType != "application/pdf" {
		return fmt.Errorf("unsupported file content type: %s", contentType)
	}

	url, err := p.fileUploader.UploadFile(ctx, f)
	if err != nil {
		return fmt.Errorf("failed to upload file: %v", err)
	}

	p.files = append(p.files, url)

	err = p.dispatcher.RaiseEvent(
		event.NewEvent(
			event.FileAttached,
			event.Payload{
				"PaymentID": p.id,
			},
		),
	)

	if err != nil {
		logger.Log.Error("failed to dispatch an event: %v", err)
	}

	return nil
}
