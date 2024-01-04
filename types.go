package main

import (
	"crypto/rand"
	"math/big"
	"time"
)

type CreateAccountRequest struct {
	FirstName string `json:"firstName`
	LastName  string `json:"lastName`
}

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt`
}

func NewAccount(firstName, lastName string) *Account {

	// Generate a random ID between 0 and 9999
	id, _ := rand.Int(rand.Reader, big.NewInt(10000))

	// Generate a random number between 0 and 9999999
	number, _ := rand.Int(rand.Reader, big.NewInt(10000000))

	return &Account{
		ID:        int(id.Int64()),
		FirstName: firstName,
		LastName:  lastName,
		Number:    number.Int64(),
		CreatedAt: time.Now().UTC(),
	}
}
