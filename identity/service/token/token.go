package jwt

import "time"

type (
	// Tokener is a general interface that provides method for working
	// with identification through the token string. Providers can be different, like
	// JWT, OAuth etc.
	Tokener interface {
		Generate(claims map[string]interface{}, expiration time.Duration) (string error)
		Validate(token string) (map[string]interface{}, error)
	}
)
