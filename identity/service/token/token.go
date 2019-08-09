package jwt

import (
	"time"

	"github.com/Sharykhin/go-payments/identity/service/token/jwt"
)

const (
	TYPE_JWF = iota
)

type (
	// Tokener is a general interface that provides method for working
	// with identification through the token string. Providers can be different, like
	// JWT, OAuth etc.
	Tokener interface {
		Generate(claims map[string]interface{}, expiration time.Duration) (string error)
		Validate(token string) (map[string]interface{}, error)
	}
)

func NewToken(tokenType string) Tokener {
	switch tokenType {
	case TYPE_JWF:
		return jwt.NewJWT()
	}
}
