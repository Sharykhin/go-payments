package jwt

import (
	"time"

	"github.com/Sharykhin/go-payments/identity/service/token/jwt"
)

const (
	TypeJWF = iota
)

type (
	Claims map[string]interface{}
	// Tokener is a general interface that provides method for working
	// with identification through the token string. Providers can be different, like
	// JWT, OAuth etc.
	Tokener interface {
		Generate(claims map[string]interface{}, expiration time.Duration) (string, error)
		Validate(token string) (map[string]interface{}, error)
	}
)

// NewTokenService returns a concrete service that implements
// Tokener interface
func NewTokenService(tokenType int) Tokener {
	switch tokenType {
	case TypeJWF:
		return jwt.NewJWT()
	default:
		panic("unsupported token type")
	}
}
