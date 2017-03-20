package corporation

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

// GetCorporationsCorporationIDRolesReader is a Reader for the GetCorporationsCorporationIDRoles structure.
type GetCorporationsCorporationIDRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCorporationsCorporationIDRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCorporationsCorporationIDRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCorporationsCorporationIDRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDRolesOK creates a GetCorporationsCorporationIDRolesOK with default headers values
func NewGetCorporationsCorporationIDRolesOK() *GetCorporationsCorporationIDRolesOK {
	return &GetCorporationsCorporationIDRolesOK{}
}

/*GetCorporationsCorporationIDRolesOK handles this case with default header values.

List of member character ID's and roles
*/
type GetCorporationsCorporationIDRolesOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*GetCorporationsCorporationIDRolesOKBodyItems0
}

func (o *GetCorporationsCorporationIDRolesOK) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/roles/][%d] getCorporationsCorporationIdRolesOK  %+v", 200, o.Payload)
}

func (o *GetCorporationsCorporationIDRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDRolesForbidden creates a GetCorporationsCorporationIDRolesForbidden with default headers values
func NewGetCorporationsCorporationIDRolesForbidden() *GetCorporationsCorporationIDRolesForbidden {
	return &GetCorporationsCorporationIDRolesForbidden{}
}

/*GetCorporationsCorporationIDRolesForbidden handles this case with default header values.

Forbidden
*/
type GetCorporationsCorporationIDRolesForbidden struct {
	Payload GetCorporationsCorporationIDRolesForbiddenBody
}

func (o *GetCorporationsCorporationIDRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/roles/][%d] getCorporationsCorporationIdRolesForbidden  %+v", 403, o.Payload)
}

func (o *GetCorporationsCorporationIDRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDRolesInternalServerError creates a GetCorporationsCorporationIDRolesInternalServerError with default headers values
func NewGetCorporationsCorporationIDRolesInternalServerError() *GetCorporationsCorporationIDRolesInternalServerError {
	return &GetCorporationsCorporationIDRolesInternalServerError{}
}

/*GetCorporationsCorporationIDRolesInternalServerError handles this case with default header values.

Internal server error
*/
type GetCorporationsCorporationIDRolesInternalServerError struct {
	Payload GetCorporationsCorporationIDRolesInternalServerErrorBody
}

func (o *GetCorporationsCorporationIDRolesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/roles/][%d] getCorporationsCorporationIdRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCorporationsCorporationIDRolesForbiddenBody get_corporations_corporation_id_roles_forbidden
//
// Forbidden
swagger:model GetCorporationsCorporationIDRolesForbiddenBody
*/
type GetCorporationsCorporationIDRolesForbiddenBody struct {

	// get_corporations_corporation_id_roles_403_forbidden
	//
	// Forbidden message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get corporations corporation ID roles forbidden body
func (o *GetCorporationsCorporationIDRolesForbiddenBody) Validate(formats strfmt.Registry) error {
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

func (o *GetCorporationsCorporationIDRolesForbiddenBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getCorporationsCorporationIdRolesForbidden"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*GetCorporationsCorporationIDRolesInternalServerErrorBody get_corporations_corporation_id_roles_internal_server_error
//
// Internal server error
swagger:model GetCorporationsCorporationIDRolesInternalServerErrorBody
*/
type GetCorporationsCorporationIDRolesInternalServerErrorBody struct {

	// get_corporations_corporation_id_roles_500_internal_server_error
	//
	// Internal server error message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get corporations corporation ID roles internal server error body
func (o *GetCorporationsCorporationIDRolesInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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

func (o *GetCorporationsCorporationIDRolesInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getCorporationsCorporationIdRolesInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*GetCorporationsCorporationIDRolesOKBodyItems0 get_corporations_corporation_id_roles_200_ok
//
// 200 ok object
swagger:model GetCorporationsCorporationIDRolesOKBodyItems0
*/
type GetCorporationsCorporationIDRolesOKBodyItems0 struct {

	// get_corporations_corporation_id_roles_character_id
	//
	// character_id integer
	// Required: true
	CharacterID *int32 `json:"character_id"`

	// get_corporations_corporation_id_roles_grantable_roles
	//
	// grantable_roles array
	GrantableRoles []string `json:"grantable_roles"`

	// get_corporations_corporation_id_roles_grantable_roles_at_base
	//
	// grantable_roles_at_base array
	GrantableRolesAtBase []string `json:"grantable_roles_at_base"`

	// get_corporations_corporation_id_roles_grantable_roles_at_hq
	//
	// grantable_roles_at_hq array
	GrantableRolesAtHq []string `json:"grantable_roles_at_hq"`

	// get_corporations_corporation_id_roles_grantable_roles_at_other
	//
	// grantable_roles_at_other array
	GrantableRolesAtOther []string `json:"grantable_roles_at_other"`

	// get_corporations_corporation_id_roles_roles
	//
	// roles array
	Roles []string `json:"roles"`

	// get_corporations_corporation_id_roles_roles_at_base
	//
	// roles_at_base array
	RolesAtBase []string `json:"roles_at_base"`

	// get_corporations_corporation_id_roles_roles_at_hq
	//
	// roles_at_hq array
	RolesAtHq []string `json:"roles_at_hq"`

	// get_corporations_corporation_id_roles_roles_at_other
	//
	// roles_at_other array
	RolesAtOther []string `json:"roles_at_other"`
}

// Validate validates this get corporations corporation ID roles o k body items0
func (o *GetCorporationsCorporationIDRolesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCharacterID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateGrantableRoles(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateGrantableRolesAtBase(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateGrantableRolesAtHq(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateGrantableRolesAtOther(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateRoles(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateRolesAtBase(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateRolesAtHq(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateRolesAtOther(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateCharacterID(formats strfmt.Registry) error {

	if err := validate.Required("character_id", "body", o.CharacterID); err != nil {
		return err
	}

	return nil
}

var getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Director","Personnel_Manager","Accountant","Security_Officer","Factory_Manager","Station_Manager","Auditor","Hangar_Take_1","Hangar_Take_2","Hangar_Take_3","Hangar_Take_4","Hangar_Take_5","Hangar_Take_6","Hangar_Take_7","Hangar_Query_1","Hangar_Query_2","Hangar_Query_3","Hangar_Query_4","Hangar_Query_5","Hangar_Query_6","Hangar_Query_7","Account_Take_1","Account_Take_2","Account_Take_3","Account_Take_4","Account_Take_5","Account_Take_6","Account_Take_7","Diplomat","Config_Equipment","Container_Take_1","Container_Take_2","Container_Take_3","Container_Take_4","Container_Take_5","Container_Take_6","Container_Take_7","Rent_Office","Rent_Factory_Facility","Rent_Research_Facility","Junior_Accountant","Config_Starbase_Equipment","Trader","Communications_Officer","Contract_Manager","Starbase_Defense_Operator","Starbase_Fuel_Technician","Fitting_Manager","Terrestrial_Combat_Officer","Terrestrial_Logistics_Officer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesItemsEnum = append(getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesItemsEnum, v)
	}
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateGrantableRolesItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesItemsEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateGrantableRoles(formats strfmt.Registry) error {

	if swag.IsZero(o.GrantableRoles) { // not required
		return nil
	}

	for i := 0; i < len(o.GrantableRoles); i++ {

		// value enum
		if err := o.validateGrantableRolesItemsEnum("grantable_roles"+"."+strconv.Itoa(i), "body", o.GrantableRoles[i]); err != nil {
			return err
		}

	}

	return nil
}

var getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesAtBaseItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Director","Personnel_Manager","Accountant","Security_Officer","Factory_Manager","Station_Manager","Auditor","Hangar_Take_1","Hangar_Take_2","Hangar_Take_3","Hangar_Take_4","Hangar_Take_5","Hangar_Take_6","Hangar_Take_7","Hangar_Query_1","Hangar_Query_2","Hangar_Query_3","Hangar_Query_4","Hangar_Query_5","Hangar_Query_6","Hangar_Query_7","Account_Take_1","Account_Take_2","Account_Take_3","Account_Take_4","Account_Take_5","Account_Take_6","Account_Take_7","Diplomat","Config_Equipment","Container_Take_1","Container_Take_2","Container_Take_3","Container_Take_4","Container_Take_5","Container_Take_6","Container_Take_7","Rent_Office","Rent_Factory_Facility","Rent_Research_Facility","Junior_Accountant","Config_Starbase_Equipment","Trader","Communications_Officer","Contract_Manager","Starbase_Defense_Operator","Starbase_Fuel_Technician","Fitting_Manager","Terrestrial_Combat_Officer","Terrestrial_Logistics_Officer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesAtBaseItemsEnum = append(getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesAtBaseItemsEnum, v)
	}
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateGrantableRolesAtBaseItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesAtBaseItemsEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateGrantableRolesAtBase(formats strfmt.Registry) error {

	if swag.IsZero(o.GrantableRolesAtBase) { // not required
		return nil
	}

	for i := 0; i < len(o.GrantableRolesAtBase); i++ {

		// value enum
		if err := o.validateGrantableRolesAtBaseItemsEnum("grantable_roles_at_base"+"."+strconv.Itoa(i), "body", o.GrantableRolesAtBase[i]); err != nil {
			return err
		}

	}

	return nil
}

var getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesAtHqItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Director","Personnel_Manager","Accountant","Security_Officer","Factory_Manager","Station_Manager","Auditor","Hangar_Take_1","Hangar_Take_2","Hangar_Take_3","Hangar_Take_4","Hangar_Take_5","Hangar_Take_6","Hangar_Take_7","Hangar_Query_1","Hangar_Query_2","Hangar_Query_3","Hangar_Query_4","Hangar_Query_5","Hangar_Query_6","Hangar_Query_7","Account_Take_1","Account_Take_2","Account_Take_3","Account_Take_4","Account_Take_5","Account_Take_6","Account_Take_7","Diplomat","Config_Equipment","Container_Take_1","Container_Take_2","Container_Take_3","Container_Take_4","Container_Take_5","Container_Take_6","Container_Take_7","Rent_Office","Rent_Factory_Facility","Rent_Research_Facility","Junior_Accountant","Config_Starbase_Equipment","Trader","Communications_Officer","Contract_Manager","Starbase_Defense_Operator","Starbase_Fuel_Technician","Fitting_Manager","Terrestrial_Combat_Officer","Terrestrial_Logistics_Officer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesAtHqItemsEnum = append(getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesAtHqItemsEnum, v)
	}
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateGrantableRolesAtHqItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesAtHqItemsEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateGrantableRolesAtHq(formats strfmt.Registry) error {

	if swag.IsZero(o.GrantableRolesAtHq) { // not required
		return nil
	}

	for i := 0; i < len(o.GrantableRolesAtHq); i++ {

		// value enum
		if err := o.validateGrantableRolesAtHqItemsEnum("grantable_roles_at_hq"+"."+strconv.Itoa(i), "body", o.GrantableRolesAtHq[i]); err != nil {
			return err
		}

	}

	return nil
}

var getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesAtOtherItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Director","Personnel_Manager","Accountant","Security_Officer","Factory_Manager","Station_Manager","Auditor","Hangar_Take_1","Hangar_Take_2","Hangar_Take_3","Hangar_Take_4","Hangar_Take_5","Hangar_Take_6","Hangar_Take_7","Hangar_Query_1","Hangar_Query_2","Hangar_Query_3","Hangar_Query_4","Hangar_Query_5","Hangar_Query_6","Hangar_Query_7","Account_Take_1","Account_Take_2","Account_Take_3","Account_Take_4","Account_Take_5","Account_Take_6","Account_Take_7","Diplomat","Config_Equipment","Container_Take_1","Container_Take_2","Container_Take_3","Container_Take_4","Container_Take_5","Container_Take_6","Container_Take_7","Rent_Office","Rent_Factory_Facility","Rent_Research_Facility","Junior_Accountant","Config_Starbase_Equipment","Trader","Communications_Officer","Contract_Manager","Starbase_Defense_Operator","Starbase_Fuel_Technician","Fitting_Manager","Terrestrial_Combat_Officer","Terrestrial_Logistics_Officer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesAtOtherItemsEnum = append(getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesAtOtherItemsEnum, v)
	}
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateGrantableRolesAtOtherItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesAtOtherItemsEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateGrantableRolesAtOther(formats strfmt.Registry) error {

	if swag.IsZero(o.GrantableRolesAtOther) { // not required
		return nil
	}

	for i := 0; i < len(o.GrantableRolesAtOther); i++ {

		// value enum
		if err := o.validateGrantableRolesAtOtherItemsEnum("grantable_roles_at_other"+"."+strconv.Itoa(i), "body", o.GrantableRolesAtOther[i]); err != nil {
			return err
		}

	}

	return nil
}

var getCorporationsCorporationIdRolesOKBodyItems0RolesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Director","Personnel_Manager","Accountant","Security_Officer","Factory_Manager","Station_Manager","Auditor","Hangar_Take_1","Hangar_Take_2","Hangar_Take_3","Hangar_Take_4","Hangar_Take_5","Hangar_Take_6","Hangar_Take_7","Hangar_Query_1","Hangar_Query_2","Hangar_Query_3","Hangar_Query_4","Hangar_Query_5","Hangar_Query_6","Hangar_Query_7","Account_Take_1","Account_Take_2","Account_Take_3","Account_Take_4","Account_Take_5","Account_Take_6","Account_Take_7","Diplomat","Config_Equipment","Container_Take_1","Container_Take_2","Container_Take_3","Container_Take_4","Container_Take_5","Container_Take_6","Container_Take_7","Rent_Office","Rent_Factory_Facility","Rent_Research_Facility","Junior_Accountant","Config_Starbase_Equipment","Trader","Communications_Officer","Contract_Manager","Starbase_Defense_Operator","Starbase_Fuel_Technician","Fitting_Manager","Terrestrial_Combat_Officer","Terrestrial_Logistics_Officer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdRolesOKBodyItems0RolesItemsEnum = append(getCorporationsCorporationIdRolesOKBodyItems0RolesItemsEnum, v)
	}
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateRolesItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCorporationsCorporationIdRolesOKBodyItems0RolesItemsEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateRoles(formats strfmt.Registry) error {

	if swag.IsZero(o.Roles) { // not required
		return nil
	}

	for i := 0; i < len(o.Roles); i++ {

		// value enum
		if err := o.validateRolesItemsEnum("roles"+"."+strconv.Itoa(i), "body", o.Roles[i]); err != nil {
			return err
		}

	}

	return nil
}

var getCorporationsCorporationIdRolesOKBodyItems0RolesAtBaseItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Director","Personnel_Manager","Accountant","Security_Officer","Factory_Manager","Station_Manager","Auditor","Hangar_Take_1","Hangar_Take_2","Hangar_Take_3","Hangar_Take_4","Hangar_Take_5","Hangar_Take_6","Hangar_Take_7","Hangar_Query_1","Hangar_Query_2","Hangar_Query_3","Hangar_Query_4","Hangar_Query_5","Hangar_Query_6","Hangar_Query_7","Account_Take_1","Account_Take_2","Account_Take_3","Account_Take_4","Account_Take_5","Account_Take_6","Account_Take_7","Diplomat","Config_Equipment","Container_Take_1","Container_Take_2","Container_Take_3","Container_Take_4","Container_Take_5","Container_Take_6","Container_Take_7","Rent_Office","Rent_Factory_Facility","Rent_Research_Facility","Junior_Accountant","Config_Starbase_Equipment","Trader","Communications_Officer","Contract_Manager","Starbase_Defense_Operator","Starbase_Fuel_Technician","Fitting_Manager","Terrestrial_Combat_Officer","Terrestrial_Logistics_Officer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdRolesOKBodyItems0RolesAtBaseItemsEnum = append(getCorporationsCorporationIdRolesOKBodyItems0RolesAtBaseItemsEnum, v)
	}
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateRolesAtBaseItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCorporationsCorporationIdRolesOKBodyItems0RolesAtBaseItemsEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateRolesAtBase(formats strfmt.Registry) error {

	if swag.IsZero(o.RolesAtBase) { // not required
		return nil
	}

	for i := 0; i < len(o.RolesAtBase); i++ {

		// value enum
		if err := o.validateRolesAtBaseItemsEnum("roles_at_base"+"."+strconv.Itoa(i), "body", o.RolesAtBase[i]); err != nil {
			return err
		}

	}

	return nil
}

var getCorporationsCorporationIdRolesOKBodyItems0RolesAtHqItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Director","Personnel_Manager","Accountant","Security_Officer","Factory_Manager","Station_Manager","Auditor","Hangar_Take_1","Hangar_Take_2","Hangar_Take_3","Hangar_Take_4","Hangar_Take_5","Hangar_Take_6","Hangar_Take_7","Hangar_Query_1","Hangar_Query_2","Hangar_Query_3","Hangar_Query_4","Hangar_Query_5","Hangar_Query_6","Hangar_Query_7","Account_Take_1","Account_Take_2","Account_Take_3","Account_Take_4","Account_Take_5","Account_Take_6","Account_Take_7","Diplomat","Config_Equipment","Container_Take_1","Container_Take_2","Container_Take_3","Container_Take_4","Container_Take_5","Container_Take_6","Container_Take_7","Rent_Office","Rent_Factory_Facility","Rent_Research_Facility","Junior_Accountant","Config_Starbase_Equipment","Trader","Communications_Officer","Contract_Manager","Starbase_Defense_Operator","Starbase_Fuel_Technician","Fitting_Manager","Terrestrial_Combat_Officer","Terrestrial_Logistics_Officer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdRolesOKBodyItems0RolesAtHqItemsEnum = append(getCorporationsCorporationIdRolesOKBodyItems0RolesAtHqItemsEnum, v)
	}
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateRolesAtHqItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCorporationsCorporationIdRolesOKBodyItems0RolesAtHqItemsEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateRolesAtHq(formats strfmt.Registry) error {

	if swag.IsZero(o.RolesAtHq) { // not required
		return nil
	}

	for i := 0; i < len(o.RolesAtHq); i++ {

		// value enum
		if err := o.validateRolesAtHqItemsEnum("roles_at_hq"+"."+strconv.Itoa(i), "body", o.RolesAtHq[i]); err != nil {
			return err
		}

	}

	return nil
}

var getCorporationsCorporationIdRolesOKBodyItems0RolesAtOtherItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Director","Personnel_Manager","Accountant","Security_Officer","Factory_Manager","Station_Manager","Auditor","Hangar_Take_1","Hangar_Take_2","Hangar_Take_3","Hangar_Take_4","Hangar_Take_5","Hangar_Take_6","Hangar_Take_7","Hangar_Query_1","Hangar_Query_2","Hangar_Query_3","Hangar_Query_4","Hangar_Query_5","Hangar_Query_6","Hangar_Query_7","Account_Take_1","Account_Take_2","Account_Take_3","Account_Take_4","Account_Take_5","Account_Take_6","Account_Take_7","Diplomat","Config_Equipment","Container_Take_1","Container_Take_2","Container_Take_3","Container_Take_4","Container_Take_5","Container_Take_6","Container_Take_7","Rent_Office","Rent_Factory_Facility","Rent_Research_Facility","Junior_Accountant","Config_Starbase_Equipment","Trader","Communications_Officer","Contract_Manager","Starbase_Defense_Operator","Starbase_Fuel_Technician","Fitting_Manager","Terrestrial_Combat_Officer","Terrestrial_Logistics_Officer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdRolesOKBodyItems0RolesAtOtherItemsEnum = append(getCorporationsCorporationIdRolesOKBodyItems0RolesAtOtherItemsEnum, v)
	}
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateRolesAtOtherItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCorporationsCorporationIdRolesOKBodyItems0RolesAtOtherItemsEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateRolesAtOther(formats strfmt.Registry) error {

	if swag.IsZero(o.RolesAtOther) { // not required
		return nil
	}

	for i := 0; i < len(o.RolesAtOther); i++ {

		// value enum
		if err := o.validateRolesAtOtherItemsEnum("roles_at_other"+"."+strconv.Itoa(i), "body", o.RolesAtOther[i]); err != nil {
			return err
		}

	}

	return nil
}
