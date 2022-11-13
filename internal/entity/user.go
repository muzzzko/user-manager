package entity

import (
	"fmt"

	"github.com/go-openapi/strfmt"

	"github/user-manager/internal/constant"
)

type User struct {
	ID           strfmt.UUID
	FirstName    string
	LastName     string
	Nickname     string
	PasswordHash string
	Email        string
	Country      Country
}

func (u User) GetSearchFields() []string {
	return []string{
		fmt.Sprintf("%s:%s", constant.CountryCodeFilter, u.Country.Code),
	}
}
