package corporation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCorporationsCorporationIDParams creates a new GetCorporationsCorporationIDParams object
// with the default values initialized.
func NewGetCorporationsCorporationIDParams() *GetCorporationsCorporationIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCorporationsCorporationIDParamsWithTimeout creates a new GetCorporationsCorporationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCorporationsCorporationIDParamsWithTimeout(timeout time.Duration) *GetCorporationsCorporationIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCorporationsCorporationIDParamsWithContext creates a new GetCorporationsCorporationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCorporationsCorporationIDParamsWithContext(ctx context.Context) *GetCorporationsCorporationIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetCorporationsCorporationIDParamsWithHTTPClient creates a new GetCorporationsCorporationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCorporationsCorporationIDParamsWithHTTPClient(client *http.Client) *GetCorporationsCorporationIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetCorporationsCorporationIDParams contains all the parameters to send to the API endpoint
for the get corporations corporation id operation typically these are written to a http.Request
*/
type GetCorporationsCorporationIDParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
	/*CorporationID
	  An Eve corporation ID

	*/
	CorporationID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*UserAgent
	  Client identifier, takes precedence over headers

	*/
	UserAgent *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get corporations corporation id params
func (o *GetCorporationsCorporationIDParams) WithTimeout(timeout time.Duration) *GetCorporationsCorporationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get corporations corporation id params
func (o *GetCorporationsCorporationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get corporations corporation id params
func (o *GetCorporationsCorporationIDParams) WithContext(ctx context.Context) *GetCorporationsCorporationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get corporations corporation id params
func (o *GetCorporationsCorporationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get corporations corporation id params
func (o *GetCorporationsCorporationIDParams) WithHTTPClient(client *http.Client) *GetCorporationsCorporationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get corporations corporation id params
func (o *GetCorporationsCorporationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get corporations corporation id params
func (o *GetCorporationsCorporationIDParams) WithXUserAgent(xUserAgent *string) *GetCorporationsCorporationIDParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get corporations corporation id params
func (o *GetCorporationsCorporationIDParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithCorporationID adds the corporationID to the get corporations corporation id params
func (o *GetCorporationsCorporationIDParams) WithCorporationID(corporationID int32) *GetCorporationsCorporationIDParams {
	o.SetCorporationID(corporationID)
	return o
}

// SetCorporationID adds the corporationId to the get corporations corporation id params
func (o *GetCorporationsCorporationIDParams) SetCorporationID(corporationID int32) {
	o.CorporationID = corporationID
}

// WithDatasource adds the datasource to the get corporations corporation id params
func (o *GetCorporationsCorporationIDParams) WithDatasource(datasource *string) *GetCorporationsCorporationIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get corporations corporation id params
func (o *GetCorporationsCorporationIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithUserAgent adds the userAgent to the get corporations corporation id params
func (o *GetCorporationsCorporationIDParams) WithUserAgent(userAgent *string) *GetCorporationsCorporationIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get corporations corporation id params
func (o *GetCorporationsCorporationIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetCorporationsCorporationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.XUserAgent != nil {

		// header param X-User-Agent
		if err := r.SetHeaderParam("X-User-Agent", *o.XUserAgent); err != nil {
			return err
		}

	}

	// path param corporation_id
	if err := r.SetPathParam("corporation_id", swag.FormatInt32(o.CorporationID)); err != nil {
		return err
	}

	if o.Datasource != nil {

		// query param datasource
		var qrDatasource string
		if o.Datasource != nil {
			qrDatasource = *o.Datasource
		}
		qDatasource := qrDatasource
		if qDatasource != "" {
			if err := r.SetQueryParam("datasource", qDatasource); err != nil {
				return err
			}
		}

	}

	if o.UserAgent != nil {

		// query param user_agent
		var qrUserAgent string
		if o.UserAgent != nil {
			qrUserAgent = *o.UserAgent
		}
		qUserAgent := qrUserAgent
		if qUserAgent != "" {
			if err := r.SetQueryParam("user_agent", qUserAgent); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
