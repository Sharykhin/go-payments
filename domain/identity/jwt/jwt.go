package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	jwtGo "github.com/dgrijalva/jwt-go"
)

var (
	TokenIsExpired = errors.New("TokenIsExpired")
)

const (
	SH256 = Algorithm(iota)
	RS256
)

type (
	//TokenManager is an interface that provides methods to work with JWT token
	TokenManager interface {
		Generate(claims Claims, expiration time.Duration) (Token, error)
		Validate(tokenString Token) (Claims, error)
	}
	// Algorithm is an alias of int to make this value more semantic and provide some restrictions
	Algorithm int
	// Claims is a payload that will be passed into token
	Claims map[string]interface{}
	// Token is an alias of string to make the token itself more semantic
	Token string
	// jwt is a struct that can generate and validate JWT tokens
	jwt struct {
		algorithm Algorithm
		secret    []byte
	}
)

// Generate generates a new jwt token with a provided data and expiration
func (j jwt) Generate(claims Claims, expiration time.Duration) (Token, error) {
	jwtClaims := jwtGo.MapClaims(claims)
	jwtClaims["exp"] = time.Now().UTC().Add(expiration).Unix()

	token := jwtGo.NewWithClaims(jwtGo.SigningMethodHS256, jwtClaims)
	tokenStr, err := token.SignedString(j.secret)
	if err != nil {
		return "", fmt.Errorf("failded to generate a token: %v", err)
	}

	return Token(tokenStr), nil
}

// Validate validates token and if it is valid returns its claims
func (j jwt) Validate(tokenStr Token) (Claims, error) {
	token, err := jwtGo.Parse(string(tokenStr), func(token *jwtGo.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwtGo.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.secret, nil
	})

	if err != nil {
		jwtErr := err.(*jwtGo.ValidationError)
		if jwtErr.Errors == jwtGo.ValidationErrorExpired {
			return nil, TokenIsExpired
		}

		return nil, fmt.Errorf("could not parse token string: %v", err)
	}

	if claims, ok := token.Claims.(jwtGo.MapClaims); ok && token.Valid {
		return Claims(claims), nil
	}

	return nil, fmt.Errorf("failed to parse claims: %v", err)
}

// NewTokenManager returns a concrete jwt implementation based on secret
// that satisfies an interface
func NewTokenManager(algorithm Algorithm) TokenManager {
	tm := jwt{
		algorithm: algorithm,
	}

	switch algorithm {
	case SH256:
		tm.secret = []byte(os.Getenv("JWT_SECRET"))
	case RS256:
		// Implement rs256
	}

	return &tm
}
