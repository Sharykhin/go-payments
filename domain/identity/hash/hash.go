package hash

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type (
	Hash   string
	Hasher interface {
		GenerateHash(password string) (Hash, error)
		ValidateHash(plaint string, hash Hash) (bool, error)
	}

	BcryptHasher struct {
		cost int
	}
)

func (h BcryptHasher) GenerateHash(password string) (Hash, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)

	if err != nil {
		return "", fmt.Errorf("could not generate bcrypt password: %v", err)
	}

	return Hash(string(hash)), nil
}

func (h BcryptHasher) ValidateHash(plaint string, hash Hash) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plaint))
	if err != nil {
		return false, err
	}

	return true, nil
}

func NewHasher() Hasher {
	h := BcryptHasher{
		cost: bcrypt.DefaultCost,
	}

	return &h
}

func NewBcryptHasher(cost int) Hasher {
	h := BcryptHasher{
		cost: cost,
	}

	return &h
}
