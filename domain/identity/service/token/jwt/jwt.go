package jwt

import (
	"fmt"
	"time"

	tokenErrors "github.com/Sharykhin/go-payments/domain/identity/service/token/error"

	jwtGo "github.com/dgrijalva/jwt-go"
)

var (
	// TODO: take from ENV variable
	secret = []byte("secret")
)

type (
	// JWT is a struct that can generate and validate JWT tokens
	JWT struct {
	}
)

// Generate generates a new jwt token with a provided data and expiration
func (JWT) Generate(claims map[string]interface{}, expiration time.Duration) (string, error) {
	jwtClaims := jwtGo.MapClaims(claims)
	jwtClaims["exp"] = time.Now().UTC().Add(expiration).Unix()

	token := jwtGo.NewWithClaims(jwtGo.SigningMethodHS256, jwtClaims)
	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("failded to generate a token: %v", err)
	}

	return tokenStr, nil
}

// Validate validates token and if it is valid returns its claims
func (JWT) Validate(tokenString string) (map[string]interface{}, error) {
	token, err := jwtGo.Parse(tokenString, func(token *jwtGo.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwtGo.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		jwtErr := err.(*jwtGo.ValidationError)
		if jwtErr.Errors == jwtGo.ValidationErrorExpired {
			return nil, tokenErrors.TokenIsExpired
		}

		return nil, fmt.Errorf("could not parse token string: %v", err)
	}

	if claims, ok := token.Claims.(jwtGo.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("failed to parse claims: %v", err)
}

// NewJWT is a function constructor
// that returns a new instance of JWT struct
func NewJWT() *JWT {
	return &JWT{}
}
