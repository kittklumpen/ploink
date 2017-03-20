package corporation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// GetCorporationsCorporationIDAlliancehistoryReader is a Reader for the GetCorporationsCorporationIDAlliancehistory structure.
type GetCorporationsCorporationIDAlliancehistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDAlliancehistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCorporationsCorporationIDAlliancehistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetCorporationsCorporationIDAlliancehistoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDAlliancehistoryOK creates a GetCorporationsCorporationIDAlliancehistoryOK with default headers values
func NewGetCorporationsCorporationIDAlliancehistoryOK() *GetCorporationsCorporationIDAlliancehistoryOK {
	return &GetCorporationsCorporationIDAlliancehistoryOK{}
}

/*GetCorporationsCorporationIDAlliancehistoryOK handles this case with default header values.

Alliance history for the given corporation
*/
type GetCorporationsCorporationIDAlliancehistoryOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*GetCorporationsCorporationIDAlliancehistoryOKBodyItems0
}

func (o *GetCorporationsCorporationIDAlliancehistoryOK) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/alliancehistory/][%d] getCorporationsCorporationIdAlliancehistoryOK  %+v", 200, o.Payload)
}

func (o *GetCorporationsCorporationIDAlliancehistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDAlliancehistoryInternalServerError creates a GetCorporationsCorporationIDAlliancehistoryInternalServerError with default headers values
func NewGetCorporationsCorporationIDAlliancehistoryInternalServerError() *GetCorporationsCorporationIDAlliancehistoryInternalServerError {
	return &GetCorporationsCorporationIDAlliancehistoryInternalServerError{}
}

/*GetCorporationsCorporationIDAlliancehistoryInternalServerError handles this case with default header values.

Internal server error
*/
type GetCorporationsCorporationIDAlliancehistoryInternalServerError struct {
	Payload GetCorporationsCorporationIDAlliancehistoryInternalServerErrorBody
}

func (o *GetCorporationsCorporationIDAlliancehistoryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/alliancehistory/][%d] getCorporationsCorporationIdAlliancehistoryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDAlliancehistoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCorporationsCorporationIDAlliancehistoryInternalServerErrorBody get_corporations_corporation_id_alliancehistory_internal_server_error
//
// Internal server error
swagger:model GetCorporationsCorporationIDAlliancehistoryInternalServerErrorBody
*/
type GetCorporationsCorporationIDAlliancehistoryInternalServerErrorBody struct {

	// get_corporations_corporation_id_alliancehistory_500_internal_server_error
	//
	// Internal server error message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get corporations corporation ID alliancehistory internal server error body
func (o *GetCorporationsCorporationIDAlliancehistoryInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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

func (o *GetCorporationsCorporationIDAlliancehistoryInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getCorporationsCorporationIdAlliancehistoryInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*GetCorporationsCorporationIDAlliancehistoryOKBodyItems0 get_corporations_corporation_id_alliancehistory_200_ok
//
// 200 ok object
swagger:model GetCorporationsCorporationIDAlliancehistoryOKBodyItems0
*/
type GetCorporationsCorporationIDAlliancehistoryOKBodyItems0 struct {

	// alliance
	Alliance *GetCorporationsCorporationIDAlliancehistoryOKBodyItems0Alliance `json:"alliance,omitempty"`

	// get_corporations_corporation_id_alliancehistory_record_id
	//
	// An incrementing ID that can be used to canonically establish order of records in cases where dates may be ambiguous
	// Required: true
	RecordID *int32 `json:"record_id"`

	// get_corporations_corporation_id_alliancehistory_start_date
	//
	// start_date string
	// Required: true
	StartDate *strfmt.DateTime `json:"start_date"`
}

// Validate validates this get corporations corporation ID alliancehistory o k body items0
func (o *GetCorporationsCorporationIDAlliancehistoryOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAlliance(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateRecordID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateStartDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCorporationsCorporationIDAlliancehistoryOKBodyItems0) validateAlliance(formats strfmt.Registry) error {

	if swag.IsZero(o.Alliance) { // not required
		return nil
	}

	if o.Alliance != nil {

		if err := o.Alliance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alliance")
			}
			return err
		}
	}

	return nil
}

func (o *GetCorporationsCorporationIDAlliancehistoryOKBodyItems0) validateRecordID(formats strfmt.Registry) error {

	if err := validate.Required("record_id", "body", o.RecordID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDAlliancehistoryOKBodyItems0) validateStartDate(formats strfmt.Registry) error {

	if err := validate.Required("start_date", "body", o.StartDate); err != nil {
		return err
	}

	return nil
}

/*GetCorporationsCorporationIDAlliancehistoryOKBodyItems0Alliance get_corporations_corporation_id_alliancehistory_alliance
//
// alliance object
swagger:model GetCorporationsCorporationIDAlliancehistoryOKBodyItems0Alliance
*/
type GetCorporationsCorporationIDAlliancehistoryOKBodyItems0Alliance struct {

	// get_corporations_corporation_id_alliancehistory_alliance_id
	//
	// alliance_id integer
	// Required: true
	AllianceID *int32 `json:"alliance_id"`

	// get_corporations_corporation_id_alliancehistory_is_deleted
	//
	// True if the alliance has been deleted
	// Required: true
	IsDeleted *bool `json:"is_deleted"`
}

// Validate validates this get corporations corporation ID alliancehistory o k body items0 alliance
func (o *GetCorporationsCorporationIDAlliancehistoryOKBodyItems0Alliance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAllianceID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateIsDeleted(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCorporationsCorporationIDAlliancehistoryOKBodyItems0Alliance) validateAllianceID(formats strfmt.Registry) error {

	if err := validate.Required("alliance"+"."+"alliance_id", "body", o.AllianceID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDAlliancehistoryOKBodyItems0Alliance) validateIsDeleted(formats strfmt.Registry) error {

	if err := validate.Required("alliance"+"."+"is_deleted", "body", o.IsDeleted); err != nil {
		return err
	}

	return nil
}
