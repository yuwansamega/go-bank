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
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
	CreatedAt string `json:"createdAt`
}

func NewAccount(firstName, lastName string) *Account {

	loc, _ := time.LoadLocation("Asia/Jakarta") // Set the location to Jakarta timezone
	createdAt := time.Now().UTC().In(loc)       // Get the current time in Jakarta timezone

	// Format the time in Indonesian format
	formattedTime := createdAt.Format("02 January 2006 15:04:05")
	// Generate a random ID between 0 and 9999
	// id, _ := rand.Int(rand.Reader, big.NewInt(10000))

	// Generate a random number between 0 and 9999999
	number, _ := rand.Int(rand.Reader, big.NewInt(10000000))

	return &Account{
		// ID:        int(id.Int64()),
		FirstName: firstName,
		LastName:  lastName,
		Number:    number.Int64(),
		CreatedAt: formattedTime,
	}
}
