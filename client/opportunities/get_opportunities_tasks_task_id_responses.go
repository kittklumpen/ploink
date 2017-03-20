package opportunities

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

// GetOpportunitiesTasksTaskIDReader is a Reader for the GetOpportunitiesTasksTaskID structure.
type GetOpportunitiesTasksTaskIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOpportunitiesTasksTaskIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOpportunitiesTasksTaskIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetOpportunitiesTasksTaskIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOpportunitiesTasksTaskIDOK creates a GetOpportunitiesTasksTaskIDOK with default headers values
func NewGetOpportunitiesTasksTaskIDOK() *GetOpportunitiesTasksTaskIDOK {
	return &GetOpportunitiesTasksTaskIDOK{}
}

/*GetOpportunitiesTasksTaskIDOK handles this case with default header values.

Details of an opportunities task
*/
type GetOpportunitiesTasksTaskIDOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload GetOpportunitiesTasksTaskIDOKBody
}

func (o *GetOpportunitiesTasksTaskIDOK) Error() string {
	return fmt.Sprintf("[GET /opportunities/tasks/{task_id}/][%d] getOpportunitiesTasksTaskIdOK  %+v", 200, o.Payload)
}

func (o *GetOpportunitiesTasksTaskIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetOpportunitiesTasksTaskIDInternalServerError creates a GetOpportunitiesTasksTaskIDInternalServerError with default headers values
func NewGetOpportunitiesTasksTaskIDInternalServerError() *GetOpportunitiesTasksTaskIDInternalServerError {
	return &GetOpportunitiesTasksTaskIDInternalServerError{}
}

/*GetOpportunitiesTasksTaskIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetOpportunitiesTasksTaskIDInternalServerError struct {
	Payload GetOpportunitiesTasksTaskIDInternalServerErrorBody
}

func (o *GetOpportunitiesTasksTaskIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /opportunities/tasks/{task_id}/][%d] getOpportunitiesTasksTaskIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOpportunitiesTasksTaskIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetOpportunitiesTasksTaskIDInternalServerErrorBody get_opportunities_tasks_task_id_internal_server_error
//
// Internal server error
swagger:model GetOpportunitiesTasksTaskIDInternalServerErrorBody
*/
type GetOpportunitiesTasksTaskIDInternalServerErrorBody struct {

	// get_opportunities_tasks_task_id_500_internal_server_error
	//
	// Internal server error message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get opportunities tasks task ID internal server error body
func (o *GetOpportunitiesTasksTaskIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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

func (o *GetOpportunitiesTasksTaskIDInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getOpportunitiesTasksTaskIdInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*GetOpportunitiesTasksTaskIDOKBody get_opportunities_tasks_task_id_ok
//
// 200 ok object
swagger:model GetOpportunitiesTasksTaskIDOKBody
*/
type GetOpportunitiesTasksTaskIDOKBody struct {

	// get_opportunities_tasks_task_id_description
	//
	// description string
	// Required: true
	Description *string `json:"description"`

	// get_opportunities_tasks_task_id_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_opportunities_tasks_task_id_notification
	//
	// notification string
	// Required: true
	Notification *string `json:"notification"`

	// get_opportunities_tasks_task_id_task_id
	//
	// task_id integer
	// Required: true
	TaskID *int32 `json:"task_id"`
}

// Validate validates this get opportunities tasks task ID o k body
func (o *GetOpportunitiesTasksTaskIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateNotification(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateTaskID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOpportunitiesTasksTaskIDOKBody) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("getOpportunitiesTasksTaskIdOK"+"."+"description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *GetOpportunitiesTasksTaskIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("getOpportunitiesTasksTaskIdOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetOpportunitiesTasksTaskIDOKBody) validateNotification(formats strfmt.Registry) error {

	if err := validate.Required("getOpportunitiesTasksTaskIdOK"+"."+"notification", "body", o.Notification); err != nil {
		return err
	}

	return nil
}

func (o *GetOpportunitiesTasksTaskIDOKBody) validateTaskID(formats strfmt.Registry) error {

	if err := validate.Required("getOpportunitiesTasksTaskIdOK"+"."+"task_id", "body", o.TaskID); err != nil {
		return err
	}

	return nil
}
