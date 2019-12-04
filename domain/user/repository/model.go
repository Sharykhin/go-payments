package repository

import (
	"time"

	types "github.com/Sharykhin/go-payments/core/type"
)

type (
	// User describe user model on repository level
	User struct {
		ID        int64 `gorm:"primary_key"`
		FirstName string
		LastName  types.NullString
		Email     string
		CreatedAt time.Time
		DeletedAt types.NullTime
		//Payments  []entity.Payment `gorm:"PRELOAD:true;foreignkey:UserID" json:"-"`
		Payments []Payment `gorm:"PRELOAD:true;foreignkey:UserID" json:"-"`
	}

	// Payment describes payment model on repository level in the context of user domain
	Payment struct {
		TransactionID string
	}
)
