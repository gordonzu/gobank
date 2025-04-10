package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	ID			int			`json:"id"`
	FirstName	string		`json:"first_name"`
	LastName	string		`json:"last_name"`
	AcctNumber	int64		`json:"acct_number"`
	Balance		int64		`json:"balance"`
	CreatedAt	time.Time	`json:"createdAt"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName:  firstName,
		LastName:   lastName,
		AcctNumber: int64(rand.Intn(1000000)),
		CreatedAt:	time.Now().UTC(),
	}
}


