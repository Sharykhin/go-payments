package password

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// GeneratePassword generate hashed password based on plain text
// Currently it uses Bcrypt algorithm with a default cost of 10
func GeneratePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("could not generate bcrypt password: %v", err)
	}

	return string(hash), nil
}

func ComparePasswords(password, hash string) error {
	return nil
}
