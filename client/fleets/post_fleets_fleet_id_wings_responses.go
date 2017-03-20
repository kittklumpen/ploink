package fleets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// PostFleetsFleetIDWingsReader is a Reader for the PostFleetsFleetIDWings structure.
type PostFleetsFleetIDWingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFleetsFleetIDWingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostFleetsFleetIDWingsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewPostFleetsFleetIDWingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostFleetsFleetIDWingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostFleetsFleetIDWingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostFleetsFleetIDWingsCreated creates a PostFleetsFleetIDWingsCreated with default headers values
func NewPostFleetsFleetIDWingsCreated() *PostFleetsFleetIDWingsCreated {
	return &PostFleetsFleetIDWingsCreated{}
}

/*PostFleetsFleetIDWingsCreated handles this case with default header values.

Wing created
*/
type PostFleetsFleetIDWingsCreated struct {
	Payload PostFleetsFleetIDWingsCreatedBody
}

func (o *PostFleetsFleetIDWingsCreated) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsCreated  %+v", 201, o.Payload)
}

func (o *PostFleetsFleetIDWingsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDWingsForbidden creates a PostFleetsFleetIDWingsForbidden with default headers values
func NewPostFleetsFleetIDWingsForbidden() *PostFleetsFleetIDWingsForbidden {
	return &PostFleetsFleetIDWingsForbidden{}
}

/*PostFleetsFleetIDWingsForbidden handles this case with default header values.

Forbidden
*/
type PostFleetsFleetIDWingsForbidden struct {
	Payload PostFleetsFleetIDWingsForbiddenBody
}

func (o *PostFleetsFleetIDWingsForbidden) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsForbidden  %+v", 403, o.Payload)
}

func (o *PostFleetsFleetIDWingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDWingsNotFound creates a PostFleetsFleetIDWingsNotFound with default headers values
func NewPostFleetsFleetIDWingsNotFound() *PostFleetsFleetIDWingsNotFound {
	return &PostFleetsFleetIDWingsNotFound{}
}

/*PostFleetsFleetIDWingsNotFound handles this case with default header values.

The fleet does not exist or you don't have access to it
*/
type PostFleetsFleetIDWingsNotFound struct {
	Payload PostFleetsFleetIDWingsNotFoundBody
}

func (o *PostFleetsFleetIDWingsNotFound) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsNotFound  %+v", 404, o.Payload)
}

func (o *PostFleetsFleetIDWingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDWingsInternalServerError creates a PostFleetsFleetIDWingsInternalServerError with default headers values
func NewPostFleetsFleetIDWingsInternalServerError() *PostFleetsFleetIDWingsInternalServerError {
	return &PostFleetsFleetIDWingsInternalServerError{}
}

/*PostFleetsFleetIDWingsInternalServerError handles this case with default header values.

Internal server error
*/
type PostFleetsFleetIDWingsInternalServerError struct {
	Payload PostFleetsFleetIDWingsInternalServerErrorBody
}

func (o *PostFleetsFleetIDWingsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostFleetsFleetIDWingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostFleetsFleetIDWingsCreatedBody post_fleets_fleet_id_wings_created
//
// 201 created object
swagger:model PostFleetsFleetIDWingsCreatedBody
*/
type PostFleetsFleetIDWingsCreatedBody struct {

	// post_fleets_fleet_id_wings_wing_id
	//
	// The wing_id of the newly created wing
	// Required: true
	WingID *int64 `json:"wing_id"`
}

// Validate validates this post fleets fleet ID wings created body
func (o *PostFleetsFleetIDWingsCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateWingID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostFleetsFleetIDWingsCreatedBody) validateWingID(formats strfmt.Registry) error {

	if err := validate.Required("postFleetsFleetIdWingsCreated"+"."+"wing_id", "body", o.WingID); err != nil {
		return err
	}

	return nil
}

/*PostFleetsFleetIDWingsForbiddenBody post_fleets_fleet_id_wings_forbidden
//
// Forbidden
swagger:model PostFleetsFleetIDWingsForbiddenBody
*/
type PostFleetsFleetIDWingsForbiddenBody struct {

	// post_fleets_fleet_id_wings_403_forbidden
	//
	// Forbidden message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this post fleets fleet ID wings forbidden body
func (o *PostFleetsFleetIDWingsForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostFleetsFleetIDWingsForbiddenBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("postFleetsFleetIdWingsForbidden"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*PostFleetsFleetIDWingsInternalServerErrorBody post_fleets_fleet_id_wings_internal_server_error
//
// Internal server error
swagger:model PostFleetsFleetIDWingsInternalServerErrorBody
*/
type PostFleetsFleetIDWingsInternalServerErrorBody struct {

	// post_fleets_fleet_id_wings_500_internal_server_error
	//
	// Internal server error message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this post fleets fleet ID wings internal server error body
func (o *PostFleetsFleetIDWingsInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostFleetsFleetIDWingsInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("postFleetsFleetIdWingsInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*PostFleetsFleetIDWingsNotFoundBody post_fleets_fleet_id_wings_not_found
//
// Not found
swagger:model PostFleetsFleetIDWingsNotFoundBody
*/
type PostFleetsFleetIDWingsNotFoundBody struct {

	// post_fleets_fleet_id_wings_404_not_found
	//
	// Not found message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this post fleets fleet ID wings not found body
func (o *PostFleetsFleetIDWingsNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostFleetsFleetIDWingsNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("postFleetsFleetIdWingsNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}
