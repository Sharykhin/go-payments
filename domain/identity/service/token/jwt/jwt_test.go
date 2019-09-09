package jwt

import (
	"testing"
	"time"
)

func TestJWT_Generate(t *testing.T) {
	claims := make(map[string]interface{})
	exp := 1 * time.Second
	token := NewJWT()

	tokenString, err := token.Generate(claims, exp)
	if err != nil {
		t.Errorf("expected err nill but got: %v", err)
	}

	t.Logf("test passed and generated token: %v", tokenString)
}

func TestJWT_Validate(t *testing.T) {
	claims := map[string]interface{}{
		"id":   10,
		"name": "name",
	}
	exp := 1 * time.Minute
	token := NewJWT()
	tokenString, err := token.Generate(claims, exp)
	if err != nil {
		t.Errorf("unexpected error on token generate: %v", err)
	}

	parsedClaims, err := token.Validate(tokenString)
	if err != nil {
		t.Errorf("expected err to be nil on Validate method but got: %v", err)
	}

	v, ok := parsedClaims["id"]
	if !ok {
		t.Errorf("expected id key to be in claims but it is missing")
	}

	if _, ok := v.(float64); !ok {
		t.Errorf("expected id type float64 but got %T", v)
	}

	if parsedClaims["id"] != float64(claims["id"].(int)) {
		t.Errorf("expected to get id equals %d but got %T %v", claims["id"], parsedClaims["id"], parsedClaims["id"])
	}

	if parsedClaims["name"] != claims["name"] {
		t.Errorf("expected to get name equals %s but got %v", claims["id"], parsedClaims["id"])
	}

	exp = 1 * time.Second
	tokenString, err = token.Generate(claims, exp)
	if err != nil {
		t.Errorf("unexpected error on token generate: %v", err)
	}

	time.Sleep(2 * time.Second)
	parsedClaims, err = token.Validate(tokenString)
	if err == nil {
		t.Error("expected err not to be nil but it turned to be nil")
	}

	if err != TokenIsExpired {
		t.Errorf("expected get TokenIsExpired error but got: %v", err)
	}
}
