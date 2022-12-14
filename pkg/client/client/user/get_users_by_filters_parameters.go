// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetUsersByFiltersParams creates a new GetUsersByFiltersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUsersByFiltersParams() *GetUsersByFiltersParams {
	return &GetUsersByFiltersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersByFiltersParamsWithTimeout creates a new GetUsersByFiltersParams object
// with the ability to set a timeout on a request.
func NewGetUsersByFiltersParamsWithTimeout(timeout time.Duration) *GetUsersByFiltersParams {
	return &GetUsersByFiltersParams{
		timeout: timeout,
	}
}

// NewGetUsersByFiltersParamsWithContext creates a new GetUsersByFiltersParams object
// with the ability to set a context for a request.
func NewGetUsersByFiltersParamsWithContext(ctx context.Context) *GetUsersByFiltersParams {
	return &GetUsersByFiltersParams{
		Context: ctx,
	}
}

// NewGetUsersByFiltersParamsWithHTTPClient creates a new GetUsersByFiltersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUsersByFiltersParamsWithHTTPClient(client *http.Client) *GetUsersByFiltersParams {
	return &GetUsersByFiltersParams{
		HTTPClient: client,
	}
}

/*
GetUsersByFiltersParams contains all the parameters to send to the API endpoint

	for the get users by filters operation.

	Typically these are written to a http.Request.
*/
type GetUsersByFiltersParams struct {

	/* Body.

	   Filters for user searching
	*/
	Body GetUsersByFiltersBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get users by filters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsersByFiltersParams) WithDefaults() *GetUsersByFiltersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get users by filters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsersByFiltersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get users by filters params
func (o *GetUsersByFiltersParams) WithTimeout(timeout time.Duration) *GetUsersByFiltersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users by filters params
func (o *GetUsersByFiltersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users by filters params
func (o *GetUsersByFiltersParams) WithContext(ctx context.Context) *GetUsersByFiltersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users by filters params
func (o *GetUsersByFiltersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users by filters params
func (o *GetUsersByFiltersParams) WithHTTPClient(client *http.Client) *GetUsersByFiltersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users by filters params
func (o *GetUsersByFiltersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get users by filters params
func (o *GetUsersByFiltersParams) WithBody(body GetUsersByFiltersBody) *GetUsersByFiltersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get users by filters params
func (o *GetUsersByFiltersParams) SetBody(body GetUsersByFiltersBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersByFiltersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
