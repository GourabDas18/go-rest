package service

import (
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	Value string
}

func (p Password) GenerateHash() (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(p.Value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (p Password) ValidateHash(givenPassword []byte) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(p.Value), givenPassword); err != nil {
		return true
	}
	return false
}
