package entity

import "time"

type (
	User struct {
		ID        int64      `gorm:"primary_key;AUTO_INCREMENT"`
		FirstName string     `gorm:"type:varchar(80);not null"`
		LastName  NullString `gorm:"type:varchar(80)"`
		Email     string     `gorm:"type:varchar(80);not null;unique_index"`
		CreatedAt time.Time
	}
)
