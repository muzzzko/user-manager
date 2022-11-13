// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github/user-manager/internal/generated/server/models"
)

// GetUsersByFiltersHandlerFunc turns a function with the right signature into a get users by filters handler
type GetUsersByFiltersHandlerFunc func(GetUsersByFiltersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUsersByFiltersHandlerFunc) Handle(params GetUsersByFiltersParams) middleware.Responder {
	return fn(params)
}

// GetUsersByFiltersHandler interface for that can handle valid get users by filters params
type GetUsersByFiltersHandler interface {
	Handle(GetUsersByFiltersParams) middleware.Responder
}

// NewGetUsersByFilters creates a new http.Handler for the get users by filters operation
func NewGetUsersByFilters(ctx *middleware.Context, handler GetUsersByFiltersHandler) *GetUsersByFilters {
	return &GetUsersByFilters{Context: ctx, Handler: handler}
}

/*
	GetUsersByFilters swagger:route GET /users user getUsersByFilters

GetUsersByFilters get users by filters API
*/
type GetUsersByFilters struct {
	Context *middleware.Context
	Handler GetUsersByFiltersHandler
}

func (o *GetUsersByFilters) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetUsersByFiltersParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetUsersByFiltersBody get users by filters body
//
// swagger:model GetUsersByFiltersBody
type GetUsersByFiltersBody struct {

	// filters
	// Required: true
	Filters models.Filters `json:"filters"`

	// User count in response
	// Maximum: 20
	// Minimum: 1
	Limit int64 `json:"limit,omitempty"`

	// Value is used for getting next user list. It is returned in response
	Next *string `json:"next,omitempty"`
}

// Validate validates this get users by filters body
func (o *GetUsersByFiltersBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUsersByFiltersBody) validateFilters(formats strfmt.Registry) error {

	if err := o.Filters.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "filters")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "filters")
		}
		return err
	}

	return nil
}

func (o *GetUsersByFiltersBody) validateLimit(formats strfmt.Registry) error {
	if swag.IsZero(o.Limit) { // not required
		return nil
	}

	if err := validate.MinimumInt("body"+"."+"limit", "body", o.Limit, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("body"+"."+"limit", "body", o.Limit, 20, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get users by filters body based on the context it is used
func (o *GetUsersByFiltersBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUsersByFiltersBody) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Filters.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "filters")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "filters")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUsersByFiltersBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUsersByFiltersBody) UnmarshalBinary(b []byte) error {
	var res GetUsersByFiltersBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}