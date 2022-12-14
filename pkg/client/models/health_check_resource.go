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

// HealthCheckResource health check resource
//
// swagger:model HealthCheckResource
type HealthCheckResource struct {

	// Show resource availability
	// Required: true
	IsAvailable *bool `json:"is_available"`

	// Checked resource
	// Required: true
	Resource *string `json:"resource"`
}

// Validate validates this health check resource
func (m *HealthCheckResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsAvailable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HealthCheckResource) validateIsAvailable(formats strfmt.Registry) error {

	if err := validate.Required("is_available", "body", m.IsAvailable); err != nil {
		return err
	}

	return nil
}

func (m *HealthCheckResource) validateResource(formats strfmt.Registry) error {

	if err := validate.Required("resource", "body", m.Resource); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this health check resource based on context it is used
func (m *HealthCheckResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HealthCheckResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HealthCheckResource) UnmarshalBinary(b []byte) error {
	var res HealthCheckResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
