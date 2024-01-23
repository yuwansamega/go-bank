package main

import (
	"crypto/rand"
	"math/big"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	Number int64  `json: "number"`
	Token  string `json: "token"`
}

type LoginRequest struct {
	Number   int64  `json:"number"`
	Password string `json:"password"`
}

type TransferRequest struct {
	ToAccount int `json:"toAccount"`
	Amount    int `json:"amount"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

func (a *Account) ValidatePassword(pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(pass)) == nil
}

type Account struct {
	ID                int    `json:"id"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	Number            int64  `json:"number"`
	EncryptedPassword string `json:"-"`
	Balance           int64  `json:"balance"`
	CreatedAt         string `json:"createdAt"`
}

func NewAccount(firstName, lastName, password string) (*Account, error) {

	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

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
		FirstName:         firstName,
		LastName:          lastName,
		EncryptedPassword: string(encpw),
		Number:            number.Int64(),
		CreatedAt:         formattedTime,
	}, nil
}
