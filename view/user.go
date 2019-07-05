package view

import "time"

type (
	// UserView represents user view model
	UserView struct {
		ID        int64     `json:"id"`
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name,omitempty"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
	}
)
