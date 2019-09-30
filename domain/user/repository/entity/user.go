package entity

import (
	"time"

	types "github.com/Sharykhin/go-payments/core/type"
	//"github.com/Sharykhin/go-payments/domain/payment/repository/entity"
)

type (
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

	Payment struct {
		TransactionID string
	}
)

func NewUser(firstName, lastName, email string) User {
	return User{
		FirstName: firstName,
		LastName: types.NullString{
			Valid:  lastName != "",
			String: lastName,
		},
		Email: email,
		DeletedAt: types.NullTime{
			Valid: false,
		},
	}
}
