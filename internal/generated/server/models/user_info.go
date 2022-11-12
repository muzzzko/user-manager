// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserInfo user info
//
// swagger:model UserInfo
type UserInfo struct {

	// User country identifier
	// Example: 1
	// Required: true
	// Minimum: 1
	CountryID int64 `json:"country_id"`

	// User email
	// Example: johnsmith@gmail.com
	// Required: true
	// Max Length: 256
	// Pattern: (?:[a-z0-9!#$%&'*+\=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+\=?^_`{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])
	Email string `json:"email"`

	// User first name
	// Example: Egor
	// Required: true
	// Max Length: 256
	// Min Length: 1
	FirstName string `json:"first_name"`

	// User last name
	// Example: Shestakov
	// Required: true
	// Max Length: 256
	// Min Length: 1
	LastName string `json:"last_name"`

	// User nickname
	// Example: muzzzko
	// Required: true
	// Max Length: 32
	// Min Length: 2
	Nickname string `json:"nickname"`

	// User password. It must contains capital, small letters and digit
	// Example: gD1wScAs
	// Required: true
	// Max Length: 256
	// Min Length: 8
	Password string `json:"password"`
}

// Validate validates this user info
func (m *UserInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountryID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNickname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserInfo) validateCountryID(formats strfmt.Registry) error {

	if err := validate.Required("country_id", "body", int64(m.CountryID)); err != nil {
		return err
	}

	if err := validate.MinimumInt("country_id", "body", m.CountryID, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *UserInfo) validateEmail(formats strfmt.Registry) error {

	if err := validate.RequiredString("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.MaxLength("email", "body", m.Email, 256); err != nil {
		return err
	}

	if err := validate.Pattern("email", "body", m.Email, `(?:[a-z0-9!#$%&'*+\=?^_`+"`"+`{|}~-]+(?:\.[a-z0-9!#$%&'*+\=?^_`+"`"+`{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])`); err != nil {
		return err
	}

	return nil
}

func (m *UserInfo) validateFirstName(formats strfmt.Registry) error {

	if err := validate.RequiredString("first_name", "body", m.FirstName); err != nil {
		return err
	}

	if err := validate.MinLength("first_name", "body", m.FirstName, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("first_name", "body", m.FirstName, 256); err != nil {
		return err
	}

	return nil
}

func (m *UserInfo) validateLastName(formats strfmt.Registry) error {

	if err := validate.RequiredString("last_name", "body", m.LastName); err != nil {
		return err
	}

	if err := validate.MinLength("last_name", "body", m.LastName, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("last_name", "body", m.LastName, 256); err != nil {
		return err
	}

	return nil
}

func (m *UserInfo) validateNickname(formats strfmt.Registry) error {

	if err := validate.RequiredString("nickname", "body", m.Nickname); err != nil {
		return err
	}

	if err := validate.MinLength("nickname", "body", m.Nickname, 2); err != nil {
		return err
	}

	if err := validate.MaxLength("nickname", "body", m.Nickname, 32); err != nil {
		return err
	}

	return nil
}

func (m *UserInfo) validatePassword(formats strfmt.Registry) error {

	if err := validate.RequiredString("password", "body", m.Password); err != nil {
		return err
	}

	if err := validate.MinLength("password", "body", m.Password, 8); err != nil {
		return err
	}

	if err := validate.MaxLength("password", "body", m.Password, 256); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user info based on context it is used
func (m *UserInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserInfo) UnmarshalBinary(b []byte) error {
	var res UserInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
