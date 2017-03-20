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

// PostFleetsFleetIDMembersReader is a Reader for the PostFleetsFleetIDMembers structure.
type PostFleetsFleetIDMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFleetsFleetIDMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPostFleetsFleetIDMembersNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewPostFleetsFleetIDMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostFleetsFleetIDMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPostFleetsFleetIDMembersUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostFleetsFleetIDMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostFleetsFleetIDMembersNoContent creates a PostFleetsFleetIDMembersNoContent with default headers values
func NewPostFleetsFleetIDMembersNoContent() *PostFleetsFleetIDMembersNoContent {
	return &PostFleetsFleetIDMembersNoContent{}
}

/*PostFleetsFleetIDMembersNoContent handles this case with default header values.

Fleet invitation sent
*/
type PostFleetsFleetIDMembersNoContent struct {
}

func (o *PostFleetsFleetIDMembersNoContent) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/members/][%d] postFleetsFleetIdMembersNoContent ", 204)
}

func (o *PostFleetsFleetIDMembersNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostFleetsFleetIDMembersForbidden creates a PostFleetsFleetIDMembersForbidden with default headers values
func NewPostFleetsFleetIDMembersForbidden() *PostFleetsFleetIDMembersForbidden {
	return &PostFleetsFleetIDMembersForbidden{}
}

/*PostFleetsFleetIDMembersForbidden handles this case with default header values.

Forbidden
*/
type PostFleetsFleetIDMembersForbidden struct {
	Payload PostFleetsFleetIDMembersForbiddenBody
}

func (o *PostFleetsFleetIDMembersForbidden) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/members/][%d] postFleetsFleetIdMembersForbidden  %+v", 403, o.Payload)
}

func (o *PostFleetsFleetIDMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDMembersNotFound creates a PostFleetsFleetIDMembersNotFound with default headers values
func NewPostFleetsFleetIDMembersNotFound() *PostFleetsFleetIDMembersNotFound {
	return &PostFleetsFleetIDMembersNotFound{}
}

/*PostFleetsFleetIDMembersNotFound handles this case with default header values.

The fleet does not exist or you don't have access to it
*/
type PostFleetsFleetIDMembersNotFound struct {
	Payload PostFleetsFleetIDMembersNotFoundBody
}

func (o *PostFleetsFleetIDMembersNotFound) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/members/][%d] postFleetsFleetIdMembersNotFound  %+v", 404, o.Payload)
}

func (o *PostFleetsFleetIDMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDMembersUnprocessableEntity creates a PostFleetsFleetIDMembersUnprocessableEntity with default headers values
func NewPostFleetsFleetIDMembersUnprocessableEntity() *PostFleetsFleetIDMembersUnprocessableEntity {
	return &PostFleetsFleetIDMembersUnprocessableEntity{}
}

/*PostFleetsFleetIDMembersUnprocessableEntity handles this case with default header values.

Errors in invitation
*/
type PostFleetsFleetIDMembersUnprocessableEntity struct {
	Payload PostFleetsFleetIDMembersUnprocessableEntityBody
}

func (o *PostFleetsFleetIDMembersUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/members/][%d] postFleetsFleetIdMembersUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostFleetsFleetIDMembersUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDMembersInternalServerError creates a PostFleetsFleetIDMembersInternalServerError with default headers values
func NewPostFleetsFleetIDMembersInternalServerError() *PostFleetsFleetIDMembersInternalServerError {
	return &PostFleetsFleetIDMembersInternalServerError{}
}

/*PostFleetsFleetIDMembersInternalServerError handles this case with default header values.

Internal server error
*/
type PostFleetsFleetIDMembersInternalServerError struct {
	Payload PostFleetsFleetIDMembersInternalServerErrorBody
}

func (o *PostFleetsFleetIDMembersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/members/][%d] postFleetsFleetIdMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *PostFleetsFleetIDMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostFleetsFleetIDMembersBody post_fleets_fleet_id_members_invitation
//
// invitation object
swagger:model PostFleetsFleetIDMembersBody
*/
type PostFleetsFleetIDMembersBody struct {

	// post_fleets_fleet_id_members_character_id
	//
	// The character you want to invite
	// Required: true
	CharacterID *int32 `json:"character_id"`

	// post_fleets_fleet_id_members_role
	//
	// - If a character is invited with the `fleet_commander` role, neither `wing_id` or `squad_id` should be specified - If a character is invited with the `wing_commander` role, only `wing_id` should be specified - If a character is invited with the `squad_commander` role, both `wing_id` and `squad_id` should be specified - If a character is invited with the `squad_member` role, `wing_id` and `squad_id` should either both be specified or not specified at all. If they aren’t specified, the invited character will join any squad with available positions
	//
	// Required: true
	Role *string `json:"role"`

	// post_fleets_fleet_id_members_squad_id
	//
	// squad_id integer
	// Required: true
	// Minimum: 0
	SquadID *int64 `json:"squad_id"`

	// post_fleets_fleet_id_members_wing_id
	//
	// wing_id integer
	// Required: true
	// Minimum: 0
	WingID *int64 `json:"wing_id"`
}

/*PostFleetsFleetIDMembersForbiddenBody post_fleets_fleet_id_members_forbidden
//
// Forbidden
swagger:model PostFleetsFleetIDMembersForbiddenBody
*/
type PostFleetsFleetIDMembersForbiddenBody struct {

	// post_fleets_fleet_id_members_403_forbidden
	//
	// Forbidden message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this post fleets fleet ID members forbidden body
func (o *PostFleetsFleetIDMembersForbiddenBody) Validate(formats strfmt.Registry) error {
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

func (o *PostFleetsFleetIDMembersForbiddenBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("postFleetsFleetIdMembersForbidden"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*PostFleetsFleetIDMembersInternalServerErrorBody post_fleets_fleet_id_members_internal_server_error
//
// Internal server error
swagger:model PostFleetsFleetIDMembersInternalServerErrorBody
*/
type PostFleetsFleetIDMembersInternalServerErrorBody struct {

	// post_fleets_fleet_id_members_500_internal_server_error
	//
	// Internal server error message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this post fleets fleet ID members internal server error body
func (o *PostFleetsFleetIDMembersInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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

func (o *PostFleetsFleetIDMembersInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("postFleetsFleetIdMembersInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*PostFleetsFleetIDMembersNotFoundBody post_fleets_fleet_id_members_not_found
//
// Not found
swagger:model PostFleetsFleetIDMembersNotFoundBody
*/
type PostFleetsFleetIDMembersNotFoundBody struct {

	// post_fleets_fleet_id_members_404_not_found
	//
	// Not found message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this post fleets fleet ID members not found body
func (o *PostFleetsFleetIDMembersNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *PostFleetsFleetIDMembersNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("postFleetsFleetIdMembersNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*PostFleetsFleetIDMembersUnprocessableEntityBody post_fleets_fleet_id_members_unprocessable_entity
//
// 422 unprocessable entity object
swagger:model PostFleetsFleetIDMembersUnprocessableEntityBody
*/
type PostFleetsFleetIDMembersUnprocessableEntityBody struct {

	// post_fleets_fleet_id_members_error
	//
	// error message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this post fleets fleet ID members unprocessable entity body
func (o *PostFleetsFleetIDMembersUnprocessableEntityBody) Validate(formats strfmt.Registry) error {
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

func (o *PostFleetsFleetIDMembersUnprocessableEntityBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("postFleetsFleetIdMembersUnprocessableEntity"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}
