package market

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

// GetMarketsGroupsMarketGroupIDReader is a Reader for the GetMarketsGroupsMarketGroupID structure.
type GetMarketsGroupsMarketGroupIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMarketsGroupsMarketGroupIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMarketsGroupsMarketGroupIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetMarketsGroupsMarketGroupIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetMarketsGroupsMarketGroupIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMarketsGroupsMarketGroupIDOK creates a GetMarketsGroupsMarketGroupIDOK with default headers values
func NewGetMarketsGroupsMarketGroupIDOK() *GetMarketsGroupsMarketGroupIDOK {
	return &GetMarketsGroupsMarketGroupIDOK{}
}

/*GetMarketsGroupsMarketGroupIDOK handles this case with default header values.

Information about an item group
*/
type GetMarketsGroupsMarketGroupIDOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*The language used in the response
	 */
	ContentLanguage string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload GetMarketsGroupsMarketGroupIDOKBody
}

func (o *GetMarketsGroupsMarketGroupIDOK) Error() string {
	return fmt.Sprintf("[GET /markets/groups/{market_group_id}/][%d] getMarketsGroupsMarketGroupIdOK  %+v", 200, o.Payload)
}

func (o *GetMarketsGroupsMarketGroupIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Content-Language
	o.ContentLanguage = response.GetHeader("Content-Language")

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

// NewGetMarketsGroupsMarketGroupIDNotFound creates a GetMarketsGroupsMarketGroupIDNotFound with default headers values
func NewGetMarketsGroupsMarketGroupIDNotFound() *GetMarketsGroupsMarketGroupIDNotFound {
	return &GetMarketsGroupsMarketGroupIDNotFound{}
}

/*GetMarketsGroupsMarketGroupIDNotFound handles this case with default header values.

Market group not found
*/
type GetMarketsGroupsMarketGroupIDNotFound struct {
	Payload GetMarketsGroupsMarketGroupIDNotFoundBody
}

func (o *GetMarketsGroupsMarketGroupIDNotFound) Error() string {
	return fmt.Sprintf("[GET /markets/groups/{market_group_id}/][%d] getMarketsGroupsMarketGroupIdNotFound  %+v", 404, o.Payload)
}

func (o *GetMarketsGroupsMarketGroupIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsGroupsMarketGroupIDInternalServerError creates a GetMarketsGroupsMarketGroupIDInternalServerError with default headers values
func NewGetMarketsGroupsMarketGroupIDInternalServerError() *GetMarketsGroupsMarketGroupIDInternalServerError {
	return &GetMarketsGroupsMarketGroupIDInternalServerError{}
}

/*GetMarketsGroupsMarketGroupIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetMarketsGroupsMarketGroupIDInternalServerError struct {
	Payload GetMarketsGroupsMarketGroupIDInternalServerErrorBody
}

func (o *GetMarketsGroupsMarketGroupIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /markets/groups/{market_group_id}/][%d] getMarketsGroupsMarketGroupIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetMarketsGroupsMarketGroupIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetMarketsGroupsMarketGroupIDInternalServerErrorBody get_markets_groups_market_group_id_internal_server_error
//
// Internal server error
swagger:model GetMarketsGroupsMarketGroupIDInternalServerErrorBody
*/
type GetMarketsGroupsMarketGroupIDInternalServerErrorBody struct {

	// get_markets_groups_market_group_id_500_internal_server_error
	//
	// Internal server error message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get markets groups market group ID internal server error body
func (o *GetMarketsGroupsMarketGroupIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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

func (o *GetMarketsGroupsMarketGroupIDInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getMarketsGroupsMarketGroupIdInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*GetMarketsGroupsMarketGroupIDNotFoundBody get_markets_groups_market_group_id_not_found
//
// Not found
swagger:model GetMarketsGroupsMarketGroupIDNotFoundBody
*/
type GetMarketsGroupsMarketGroupIDNotFoundBody struct {

	// get_markets_groups_market_group_id_404_not_found
	//
	// Not found message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get markets groups market group ID not found body
func (o *GetMarketsGroupsMarketGroupIDNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *GetMarketsGroupsMarketGroupIDNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getMarketsGroupsMarketGroupIdNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*GetMarketsGroupsMarketGroupIDOKBody get_markets_groups_market_group_id_ok
//
// 200 ok object
swagger:model GetMarketsGroupsMarketGroupIDOKBody
*/
type GetMarketsGroupsMarketGroupIDOKBody struct {

	// get_markets_groups_market_group_id_description
	//
	// description string
	// Required: true
	Description *string `json:"description"`

	// get_markets_groups_market_group_id_market_group_id
	//
	// market_group_id integer
	// Required: true
	MarketGroupID *int32 `json:"market_group_id"`

	// get_markets_groups_market_group_id_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_markets_groups_market_group_id_parent_group_id
	//
	// parent_group_id integer
	// Required: true
	ParentGroupID *int32 `json:"parent_group_id"`

	// get_markets_groups_market_group_id_types
	//
	// types array
	// Required: true
	Types []int32 `json:"types"`
}

// Validate validates this get markets groups market group ID o k body
func (o *GetMarketsGroupsMarketGroupIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateMarketGroupID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateParentGroupID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateTypes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetMarketsGroupsMarketGroupIDOKBody) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("getMarketsGroupsMarketGroupIdOK"+"."+"description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsGroupsMarketGroupIDOKBody) validateMarketGroupID(formats strfmt.Registry) error {

	if err := validate.Required("getMarketsGroupsMarketGroupIdOK"+"."+"market_group_id", "body", o.MarketGroupID); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsGroupsMarketGroupIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("getMarketsGroupsMarketGroupIdOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsGroupsMarketGroupIDOKBody) validateParentGroupID(formats strfmt.Registry) error {

	if err := validate.Required("getMarketsGroupsMarketGroupIdOK"+"."+"parent_group_id", "body", o.ParentGroupID); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsGroupsMarketGroupIDOKBody) validateTypes(formats strfmt.Registry) error {

	if err := validate.Required("getMarketsGroupsMarketGroupIdOK"+"."+"types", "body", o.Types); err != nil {
		return err
	}

	return nil
}
