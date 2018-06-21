package main

import (
	"golang.org/x/crypto/bcrypt"
)

type Token string

func (token Token) Hash() (Token, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(token),
		bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return Token(hash), nil
}

func (token Token) IsValid(newToken string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(token), []byte(newToken))
	if err != nil {
		return false
	}

	return true
}
