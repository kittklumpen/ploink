package search

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

// GetSearchReader is a Reader for the GetSearch structure.
type GetSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSearchOK creates a GetSearchOK with default headers values
func NewGetSearchOK() *GetSearchOK {
	return &GetSearchOK{}
}

/*GetSearchOK handles this case with default header values.

A list of search results
*/
type GetSearchOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload GetSearchOKBody
}

func (o *GetSearchOK) Error() string {
	return fmt.Sprintf("[GET /search/][%d] getSearchOK  %+v", 200, o.Payload)
}

func (o *GetSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSearchInternalServerError creates a GetSearchInternalServerError with default headers values
func NewGetSearchInternalServerError() *GetSearchInternalServerError {
	return &GetSearchInternalServerError{}
}

/*GetSearchInternalServerError handles this case with default header values.

Internal server error
*/
type GetSearchInternalServerError struct {
	Payload GetSearchInternalServerErrorBody
}

func (o *GetSearchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /search/][%d] getSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetSearchInternalServerErrorBody get_search_internal_server_error
//
// Internal server error
swagger:model GetSearchInternalServerErrorBody
*/
type GetSearchInternalServerErrorBody struct {

	// get_search_500_internal_server_error
	//
	// Internal server error message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get search internal server error body
func (o *GetSearchInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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

func (o *GetSearchInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getSearchInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*GetSearchOKBody get_search_ok
//
// 200 ok object
swagger:model GetSearchOKBody
*/
type GetSearchOKBody struct {

	// get_search_agent
	//
	// agent array
	// Required: true
	Agent []int32 `json:"agent"`

	// get_search_alliance
	//
	// alliance array
	// Required: true
	Alliance []int32 `json:"alliance"`

	// get_search_character
	//
	// character array
	// Required: true
	Character []int32 `json:"character"`

	// get_search_constellation
	//
	// constellation array
	// Required: true
	Constellation []int32 `json:"constellation"`

	// get_search_corporation
	//
	// corporation array
	// Required: true
	Corporation []int32 `json:"corporation"`

	// get_search_faction
	//
	// faction array
	// Required: true
	Faction []int32 `json:"faction"`

	// get_search_inventorytype
	//
	// inventorytype array
	// Required: true
	Inventorytype []int32 `json:"inventorytype"`

	// get_search_region
	//
	// region array
	// Required: true
	Region []int32 `json:"region"`

	// get_search_solarsystem
	//
	// solarsystem array
	// Required: true
	Solarsystem []int32 `json:"solarsystem"`

	// get_search_station
	//
	// station array
	// Required: true
	Station []int32 `json:"station"`
}

// Validate validates this get search o k body
func (o *GetSearchOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAgent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateAlliance(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateCharacter(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateConstellation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateCorporation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateFaction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateInventorytype(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateRegion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateSolarsystem(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateStation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSearchOKBody) validateAgent(formats strfmt.Registry) error {

	if err := validate.Required("getSearchOK"+"."+"agent", "body", o.Agent); err != nil {
		return err
	}

	return nil
}

func (o *GetSearchOKBody) validateAlliance(formats strfmt.Registry) error {

	if err := validate.Required("getSearchOK"+"."+"alliance", "body", o.Alliance); err != nil {
		return err
	}

	return nil
}

func (o *GetSearchOKBody) validateCharacter(formats strfmt.Registry) error {

	if err := validate.Required("getSearchOK"+"."+"character", "body", o.Character); err != nil {
		return err
	}

	return nil
}

func (o *GetSearchOKBody) validateConstellation(formats strfmt.Registry) error {

	if err := validate.Required("getSearchOK"+"."+"constellation", "body", o.Constellation); err != nil {
		return err
	}

	return nil
}

func (o *GetSearchOKBody) validateCorporation(formats strfmt.Registry) error {

	if err := validate.Required("getSearchOK"+"."+"corporation", "body", o.Corporation); err != nil {
		return err
	}

	return nil
}

func (o *GetSearchOKBody) validateFaction(formats strfmt.Registry) error {

	if err := validate.Required("getSearchOK"+"."+"faction", "body", o.Faction); err != nil {
		return err
	}

	return nil
}

func (o *GetSearchOKBody) validateInventorytype(formats strfmt.Registry) error {

	if err := validate.Required("getSearchOK"+"."+"inventorytype", "body", o.Inventorytype); err != nil {
		return err
	}

	return nil
}

func (o *GetSearchOKBody) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("getSearchOK"+"."+"region", "body", o.Region); err != nil {
		return err
	}

	return nil
}

func (o *GetSearchOKBody) validateSolarsystem(formats strfmt.Registry) error {

	if err := validate.Required("getSearchOK"+"."+"solarsystem", "body", o.Solarsystem); err != nil {
		return err
	}

	return nil
}

func (o *GetSearchOKBody) validateStation(formats strfmt.Registry) error {

	if err := validate.Required("getSearchOK"+"."+"station", "body", o.Station); err != nil {
		return err
	}

	return nil
}
