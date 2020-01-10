package bcrypt

import (
	"fmt"

	"github.com/Sharykhin/go-payments/domain/identity/hash"

	bcryptLibrary "golang.org/x/crypto/bcrypt"
)

type (
	Hasher struct {
		cost int
	}
)

func (h Hasher) GenerateHash(password string) (hash.Hash, error) {
	hashBytes, err := bcryptLibrary.GenerateFromPassword([]byte(password), h.cost)

	if err != nil {
		return "", fmt.Errorf("could not generate bcrypt password: %v", err)
	}

	return hash.Hash(hashBytes), nil
}

func (h Hasher) ValidateHash(plaint string, hh hash.Hash) (bool, error) {
	err := bcryptLibrary.CompareHashAndPassword([]byte(hh), []byte(plaint))
	if err != nil {
		return false, err
	}

	return true, nil
}

func NewHasher(cost int) *Hasher {
	h := Hasher{
		cost: cost,
	}

	return &h
}
