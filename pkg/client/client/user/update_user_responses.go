// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github/user-manager/pkg/client/models"
)

// UpdateUserReader is a Reader for the UpdateUser structure.
type UpdateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewUpdateUserUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateUserOK creates a UpdateUserOK with default headers values
func NewUpdateUserOK() *UpdateUserOK {
	return &UpdateUserOK{}
}

/*
UpdateUserOK describes a response with status code 200, with default header values.

User was created successfully
*/
type UpdateUserOK struct {
	Payload *models.User
}

// IsSuccess returns true when this update user o k response has a 2xx status code
func (o *UpdateUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update user o k response has a 3xx status code
func (o *UpdateUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user o k response has a 4xx status code
func (o *UpdateUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update user o k response has a 5xx status code
func (o *UpdateUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update user o k response a status code equal to that given
func (o *UpdateUserOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateUserOK) Error() string {
	return fmt.Sprintf("[PATCH /users][%d] updateUserOK  %+v", 200, o.Payload)
}

func (o *UpdateUserOK) String() string {
	return fmt.Sprintf("[PATCH /users][%d] updateUserOK  %+v", 200, o.Payload)
}

func (o *UpdateUserOK) GetPayload() *models.User {
	return o.Payload
}

func (o *UpdateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserUnprocessableEntity creates a UpdateUserUnprocessableEntity with default headers values
func NewUpdateUserUnprocessableEntity() *UpdateUserUnprocessableEntity {
	return &UpdateUserUnprocessableEntity{}
}

/*
UpdateUserUnprocessableEntity describes a response with status code 422, with default header values.

Server couldn't handle request
*/
type UpdateUserUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this update user unprocessable entity response has a 2xx status code
func (o *UpdateUserUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user unprocessable entity response has a 3xx status code
func (o *UpdateUserUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user unprocessable entity response has a 4xx status code
func (o *UpdateUserUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user unprocessable entity response has a 5xx status code
func (o *UpdateUserUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update user unprocessable entity response a status code equal to that given
func (o *UpdateUserUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *UpdateUserUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /users][%d] updateUserUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateUserUnprocessableEntity) String() string {
	return fmt.Sprintf("[PATCH /users][%d] updateUserUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateUserUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateUserUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserInternalServerError creates a UpdateUserInternalServerError with default headers values
func NewUpdateUserInternalServerError() *UpdateUserInternalServerError {
	return &UpdateUserInternalServerError{}
}

/*
UpdateUserInternalServerError describes a response with status code 500, with default header values.

Internal server error. Something went wrong
*/
type UpdateUserInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this update user internal server error response has a 2xx status code
func (o *UpdateUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user internal server error response has a 3xx status code
func (o *UpdateUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user internal server error response has a 4xx status code
func (o *UpdateUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update user internal server error response has a 5xx status code
func (o *UpdateUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update user internal server error response a status code equal to that given
func (o *UpdateUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateUserInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /users][%d] updateUserInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateUserInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /users][%d] updateUserInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateUserInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
