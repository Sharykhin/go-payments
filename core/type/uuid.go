package types

import "github.com/google/uuid"

type (
	UUID string
)

// NewUUID generates a new universally unique identity
func NewUUID() string {
	return uuid.New().String()
}
