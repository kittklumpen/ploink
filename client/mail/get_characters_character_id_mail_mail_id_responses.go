package mail

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// GetCharactersCharacterIDMailMailIDReader is a Reader for the GetCharactersCharacterIDMailMailID structure.
type GetCharactersCharacterIDMailMailIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDMailMailIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDMailMailIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDMailMailIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetCharactersCharacterIDMailMailIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDMailMailIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDMailMailIDOK creates a GetCharactersCharacterIDMailMailIDOK with default headers values
func NewGetCharactersCharacterIDMailMailIDOK() *GetCharactersCharacterIDMailMailIDOK {
	return &GetCharactersCharacterIDMailMailIDOK{}
}

/*GetCharactersCharacterIDMailMailIDOK handles this case with default header values.

Contents of a mail
*/
type GetCharactersCharacterIDMailMailIDOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload GetCharactersCharacterIDMailMailIDOKBody
}

func (o *GetCharactersCharacterIDMailMailIDOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/{mail_id}/][%d] getCharactersCharacterIdMailMailIdOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDMailMailIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMailMailIDForbidden creates a GetCharactersCharacterIDMailMailIDForbidden with default headers values
func NewGetCharactersCharacterIDMailMailIDForbidden() *GetCharactersCharacterIDMailMailIDForbidden {
	return &GetCharactersCharacterIDMailMailIDForbidden{}
}

/*GetCharactersCharacterIDMailMailIDForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDMailMailIDForbidden struct {
	Payload GetCharactersCharacterIDMailMailIDForbiddenBody
}

func (o *GetCharactersCharacterIDMailMailIDForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/{mail_id}/][%d] getCharactersCharacterIdMailMailIdForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDMailMailIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMailMailIDNotFound creates a GetCharactersCharacterIDMailMailIDNotFound with default headers values
func NewGetCharactersCharacterIDMailMailIDNotFound() *GetCharactersCharacterIDMailMailIDNotFound {
	return &GetCharactersCharacterIDMailMailIDNotFound{}
}

/*GetCharactersCharacterIDMailMailIDNotFound handles this case with default header values.

Mail not found
*/
type GetCharactersCharacterIDMailMailIDNotFound struct {
	Payload GetCharactersCharacterIDMailMailIDNotFoundBody
}

func (o *GetCharactersCharacterIDMailMailIDNotFound) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/{mail_id}/][%d] getCharactersCharacterIdMailMailIdNotFound  %+v", 404, o.Payload)
}

func (o *GetCharactersCharacterIDMailMailIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMailMailIDInternalServerError creates a GetCharactersCharacterIDMailMailIDInternalServerError with default headers values
func NewGetCharactersCharacterIDMailMailIDInternalServerError() *GetCharactersCharacterIDMailMailIDInternalServerError {
	return &GetCharactersCharacterIDMailMailIDInternalServerError{}
}

/*GetCharactersCharacterIDMailMailIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDMailMailIDInternalServerError struct {
	Payload GetCharactersCharacterIDMailMailIDInternalServerErrorBody
}

func (o *GetCharactersCharacterIDMailMailIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/{mail_id}/][%d] getCharactersCharacterIdMailMailIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDMailMailIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDMailMailIDForbiddenBody get_characters_character_id_mail_mail_id_forbidden
//
// Forbidden
swagger:model GetCharactersCharacterIDMailMailIDForbiddenBody
*/
type GetCharactersCharacterIDMailMailIDForbiddenBody struct {

	// get_characters_character_id_mail_mail_id_403_forbidden
	//
	// Forbidden message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get characters character ID mail mail ID forbidden body
func (o *GetCharactersCharacterIDMailMailIDForbiddenBody) Validate(formats strfmt.Registry) error {
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

func (o *GetCharactersCharacterIDMailMailIDForbiddenBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdMailMailIdForbidden"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDMailMailIDInternalServerErrorBody get_characters_character_id_mail_mail_id_internal_server_error
//
// Internal server error
swagger:model GetCharactersCharacterIDMailMailIDInternalServerErrorBody
*/
type GetCharactersCharacterIDMailMailIDInternalServerErrorBody struct {

	// get_characters_character_id_mail_mail_id_500_internal_server_error
	//
	// Internal server error message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get characters character ID mail mail ID internal server error body
func (o *GetCharactersCharacterIDMailMailIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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

func (o *GetCharactersCharacterIDMailMailIDInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdMailMailIdInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDMailMailIDNotFoundBody get_characters_character_id_mail_mail_id_not_found
//
// Not found
swagger:model GetCharactersCharacterIDMailMailIDNotFoundBody
*/
type GetCharactersCharacterIDMailMailIDNotFoundBody struct {

	// get_characters_character_id_mail_mail_id_404_not_found
	//
	// Not found message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get characters character ID mail mail ID not found body
func (o *GetCharactersCharacterIDMailMailIDNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *GetCharactersCharacterIDMailMailIDNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdMailMailIdNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDMailMailIDOKBody get_characters_character_id_mail_mail_id_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDMailMailIDOKBody
*/
type GetCharactersCharacterIDMailMailIDOKBody struct {

	// get_characters_character_id_mail_mail_id_body
	//
	// Mail's body
	// Required: true
	Body *string `json:"body"`

	// get_characters_character_id_mail_mail_id_from
	//
	// From whom the mail was sent
	// Required: true
	From *int32 `json:"from"`

	// get_characters_character_id_mail_mail_id_labels
	//
	// Labels attached to the mail
	// Required: true
	Labels []*int64 `json:"labels"`

	// get_characters_character_id_mail_mail_id_read
	//
	// Whether the mail is flagged as read
	// Required: true
	Read *bool `json:"read"`

	// get_characters_character_id_mail_mail_id_recipients
	//
	// Recipients of the mail
	// Required: true
	// Max Items: 50
	// Min Items: 1
	// Unique: true
	Recipients []*RecipientsItems0 `json:"recipients"`

	// get_characters_character_id_mail_mail_id_subject
	//
	// Mail subject
	// Required: true
	Subject *string `json:"subject"`

	// get_characters_character_id_mail_mail_id_timestamp
	//
	// When the mail was sent
	// Required: true
	Timestamp *strfmt.DateTime `json:"timestamp"`
}

// Validate validates this get characters character ID mail mail ID o k body
func (o *GetCharactersCharacterIDMailMailIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBody(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateFrom(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateLabels(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateRead(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateRecipients(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateSubject(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateTimestamp(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDMailMailIDOKBody) validateBody(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdMailMailIdOK"+"."+"body", "body", o.Body); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDMailMailIDOKBody) validateFrom(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdMailMailIdOK"+"."+"from", "body", o.From); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDMailMailIDOKBody) validateLabels(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdMailMailIdOK"+"."+"labels", "body", o.Labels); err != nil {
		return err
	}

	for i := 0; i < len(o.Labels); i++ {

		if swag.IsZero(o.Labels[i]) { // not required
			continue
		}

		if err := validate.MinimumInt("getCharactersCharacterIdMailMailIdOK"+"."+"labels"+"."+strconv.Itoa(i), "body", int64(*o.Labels[i]), 0, false); err != nil {
			return err
		}

	}

	return nil
}

func (o *GetCharactersCharacterIDMailMailIDOKBody) validateRead(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdMailMailIdOK"+"."+"read", "body", o.Read); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDMailMailIDOKBody) validateRecipients(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdMailMailIdOK"+"."+"recipients", "body", o.Recipients); err != nil {
		return err
	}

	iRecipientsSize := int64(len(o.Recipients))

	if err := validate.MinItems("getCharactersCharacterIdMailMailIdOK"+"."+"recipients", "body", iRecipientsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("getCharactersCharacterIdMailMailIdOK"+"."+"recipients", "body", iRecipientsSize, 50); err != nil {
		return err
	}

	if err := validate.UniqueItems("getCharactersCharacterIdMailMailIdOK"+"."+"recipients", "body", o.Recipients); err != nil {
		return err
	}

	for i := 0; i < len(o.Recipients); i++ {

		if swag.IsZero(o.Recipients[i]) { // not required
			continue
		}

		if o.Recipients[i] != nil {

			if err := o.Recipients[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getCharactersCharacterIdMailMailIdOK" + "." + "recipients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetCharactersCharacterIDMailMailIDOKBody) validateSubject(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdMailMailIdOK"+"."+"subject", "body", o.Subject); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDMailMailIDOKBody) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdMailMailIdOK"+"."+"timestamp", "body", o.Timestamp); err != nil {
		return err
	}

	return nil
}

/*RecipientsItems0 get_characters_character_id_mail_mail_id_recipient
//
// recipient object
swagger:model RecipientsItems0
*/
type RecipientsItems0 struct {

	// get_characters_character_id_mail_mail_id_recipient_id
	//
	// recipient_id integer
	// Required: true
	RecipientID *int32 `json:"recipient_id"`

	// get_characters_character_id_mail_mail_id_recipient_type
	//
	// recipient_type string
	// Required: true
	RecipientType *string `json:"recipient_type"`
}

// Validate validates this recipients items0
func (o *RecipientsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRecipientID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateRecipientType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RecipientsItems0) validateRecipientID(formats strfmt.Registry) error {

	if err := validate.Required("recipient_id", "body", o.RecipientID); err != nil {
		return err
	}

	return nil
}

var recipientsItems0TypeRecipientTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alliance","character","corporation","mailing_list"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recipientsItems0TypeRecipientTypePropEnum = append(recipientsItems0TypeRecipientTypePropEnum, v)
	}
}

const (
	// RecipientsItems0RecipientTypeAlliance captures enum value "alliance"
	RecipientsItems0RecipientTypeAlliance string = "alliance"
	// RecipientsItems0RecipientTypeCharacter captures enum value "character"
	RecipientsItems0RecipientTypeCharacter string = "character"
	// RecipientsItems0RecipientTypeCorporation captures enum value "corporation"
	RecipientsItems0RecipientTypeCorporation string = "corporation"
	// RecipientsItems0RecipientTypeMailingList captures enum value "mailing_list"
	RecipientsItems0RecipientTypeMailingList string = "mailing_list"
)

// prop value enum
func (o *RecipientsItems0) validateRecipientTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, recipientsItems0TypeRecipientTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *RecipientsItems0) validateRecipientType(formats strfmt.Registry) error {

	if err := validate.Required("recipient_type", "body", o.RecipientType); err != nil {
		return err
	}

	// value enum
	if err := o.validateRecipientTypeEnum("recipient_type", "body", *o.RecipientType); err != nil {
		return err
	}

	return nil
}
