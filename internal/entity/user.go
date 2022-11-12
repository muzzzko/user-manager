package entity

import "github.com/go-openapi/strfmt"

type User struct {
	ID           strfmt.UUID
	FirstName    string
	LastName     string
	Nickname     string
	PasswordHash string
	Email        string
	Country      Country
}
