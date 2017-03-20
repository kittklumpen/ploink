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

// PutFleetsFleetIDSquadsSquadIDReader is a Reader for the PutFleetsFleetIDSquadsSquadID structure.
type PutFleetsFleetIDSquadsSquadIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutFleetsFleetIDSquadsSquadIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPutFleetsFleetIDSquadsSquadIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewPutFleetsFleetIDSquadsSquadIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutFleetsFleetIDSquadsSquadIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPutFleetsFleetIDSquadsSquadIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutFleetsFleetIDSquadsSquadIDNoContent creates a PutFleetsFleetIDSquadsSquadIDNoContent with default headers values
func NewPutFleetsFleetIDSquadsSquadIDNoContent() *PutFleetsFleetIDSquadsSquadIDNoContent {
	return &PutFleetsFleetIDSquadsSquadIDNoContent{}
}

/*PutFleetsFleetIDSquadsSquadIDNoContent handles this case with default header values.

Squad renamed
*/
type PutFleetsFleetIDSquadsSquadIDNoContent struct {
}

func (o *PutFleetsFleetIDSquadsSquadIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/squads/{squad_id}/][%d] putFleetsFleetIdSquadsSquadIdNoContent ", 204)
}

func (o *PutFleetsFleetIDSquadsSquadIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutFleetsFleetIDSquadsSquadIDForbidden creates a PutFleetsFleetIDSquadsSquadIDForbidden with default headers values
func NewPutFleetsFleetIDSquadsSquadIDForbidden() *PutFleetsFleetIDSquadsSquadIDForbidden {
	return &PutFleetsFleetIDSquadsSquadIDForbidden{}
}

/*PutFleetsFleetIDSquadsSquadIDForbidden handles this case with default header values.

Forbidden
*/
type PutFleetsFleetIDSquadsSquadIDForbidden struct {
	Payload PutFleetsFleetIDSquadsSquadIDForbiddenBody
}

func (o *PutFleetsFleetIDSquadsSquadIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/squads/{squad_id}/][%d] putFleetsFleetIdSquadsSquadIdForbidden  %+v", 403, o.Payload)
}

func (o *PutFleetsFleetIDSquadsSquadIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDSquadsSquadIDNotFound creates a PutFleetsFleetIDSquadsSquadIDNotFound with default headers values
func NewPutFleetsFleetIDSquadsSquadIDNotFound() *PutFleetsFleetIDSquadsSquadIDNotFound {
	return &PutFleetsFleetIDSquadsSquadIDNotFound{}
}

/*PutFleetsFleetIDSquadsSquadIDNotFound handles this case with default header values.

The fleet does not exist or you don't have access to it
*/
type PutFleetsFleetIDSquadsSquadIDNotFound struct {
	Payload PutFleetsFleetIDSquadsSquadIDNotFoundBody
}

func (o *PutFleetsFleetIDSquadsSquadIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/squads/{squad_id}/][%d] putFleetsFleetIdSquadsSquadIdNotFound  %+v", 404, o.Payload)
}

func (o *PutFleetsFleetIDSquadsSquadIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDSquadsSquadIDInternalServerError creates a PutFleetsFleetIDSquadsSquadIDInternalServerError with default headers values
func NewPutFleetsFleetIDSquadsSquadIDInternalServerError() *PutFleetsFleetIDSquadsSquadIDInternalServerError {
	return &PutFleetsFleetIDSquadsSquadIDInternalServerError{}
}

/*PutFleetsFleetIDSquadsSquadIDInternalServerError handles this case with default header values.

Internal server error
*/
type PutFleetsFleetIDSquadsSquadIDInternalServerError struct {
	Payload PutFleetsFleetIDSquadsSquadIDInternalServerErrorBody
}

func (o *PutFleetsFleetIDSquadsSquadIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/squads/{squad_id}/][%d] putFleetsFleetIdSquadsSquadIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PutFleetsFleetIDSquadsSquadIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PutFleetsFleetIDSquadsSquadIDBody put_fleets_fleet_id_squads_squad_id_naming
//
// naming object
swagger:model PutFleetsFleetIDSquadsSquadIDBody
*/
type PutFleetsFleetIDSquadsSquadIDBody struct {

	// put_fleets_fleet_id_squads_squad_id_name
	//
	// name string
	// Required: true
	// Max Length: 10
	Name *string `json:"name"`
}

/*PutFleetsFleetIDSquadsSquadIDForbiddenBody put_fleets_fleet_id_squads_squad_id_forbidden
//
// Forbidden
swagger:model PutFleetsFleetIDSquadsSquadIDForbiddenBody
*/
type PutFleetsFleetIDSquadsSquadIDForbiddenBody struct {

	// put_fleets_fleet_id_squads_squad_id_403_forbidden
	//
	// Forbidden message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this put fleets fleet ID squads squad ID forbidden body
func (o *PutFleetsFleetIDSquadsSquadIDForbiddenBody) Validate(formats strfmt.Registry) error {
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

func (o *PutFleetsFleetIDSquadsSquadIDForbiddenBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("putFleetsFleetIdSquadsSquadIdForbidden"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*PutFleetsFleetIDSquadsSquadIDInternalServerErrorBody put_fleets_fleet_id_squads_squad_id_internal_server_error
//
// Internal server error
swagger:model PutFleetsFleetIDSquadsSquadIDInternalServerErrorBody
*/
type PutFleetsFleetIDSquadsSquadIDInternalServerErrorBody struct {

	// put_fleets_fleet_id_squads_squad_id_500_internal_server_error
	//
	// Internal server error message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this put fleets fleet ID squads squad ID internal server error body
func (o *PutFleetsFleetIDSquadsSquadIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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

func (o *PutFleetsFleetIDSquadsSquadIDInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("putFleetsFleetIdSquadsSquadIdInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*PutFleetsFleetIDSquadsSquadIDNotFoundBody put_fleets_fleet_id_squads_squad_id_not_found
//
// Not found
swagger:model PutFleetsFleetIDSquadsSquadIDNotFoundBody
*/
type PutFleetsFleetIDSquadsSquadIDNotFoundBody struct {

	// put_fleets_fleet_id_squads_squad_id_404_not_found
	//
	// Not found message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this put fleets fleet ID squads squad ID not found body
func (o *PutFleetsFleetIDSquadsSquadIDNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *PutFleetsFleetIDSquadsSquadIDNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("putFleetsFleetIdSquadsSquadIdNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}
