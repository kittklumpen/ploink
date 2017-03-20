package universe

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

// NewGetUniverseGraphicsGraphicIDParams creates a new GetUniverseGraphicsGraphicIDParams object
// with the default values initialized.
func NewGetUniverseGraphicsGraphicIDParams() *GetUniverseGraphicsGraphicIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseGraphicsGraphicIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUniverseGraphicsGraphicIDParamsWithTimeout creates a new GetUniverseGraphicsGraphicIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUniverseGraphicsGraphicIDParamsWithTimeout(timeout time.Duration) *GetUniverseGraphicsGraphicIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseGraphicsGraphicIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetUniverseGraphicsGraphicIDParamsWithContext creates a new GetUniverseGraphicsGraphicIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUniverseGraphicsGraphicIDParamsWithContext(ctx context.Context) *GetUniverseGraphicsGraphicIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseGraphicsGraphicIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetUniverseGraphicsGraphicIDParamsWithHTTPClient creates a new GetUniverseGraphicsGraphicIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUniverseGraphicsGraphicIDParamsWithHTTPClient(client *http.Client) *GetUniverseGraphicsGraphicIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseGraphicsGraphicIDParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetUniverseGraphicsGraphicIDParams contains all the parameters to send to the API endpoint
for the get universe graphics graphic id operation typically these are written to a http.Request
*/
type GetUniverseGraphicsGraphicIDParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*GraphicID
	  graphic_id integer

	*/
	GraphicID int32
	/*UserAgent
	  Client identifier, takes precedence over headers

	*/
	UserAgent *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get universe graphics graphic id params
func (o *GetUniverseGraphicsGraphicIDParams) WithTimeout(timeout time.Duration) *GetUniverseGraphicsGraphicIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get universe graphics graphic id params
func (o *GetUniverseGraphicsGraphicIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get universe graphics graphic id params
func (o *GetUniverseGraphicsGraphicIDParams) WithContext(ctx context.Context) *GetUniverseGraphicsGraphicIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get universe graphics graphic id params
func (o *GetUniverseGraphicsGraphicIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get universe graphics graphic id params
func (o *GetUniverseGraphicsGraphicIDParams) WithHTTPClient(client *http.Client) *GetUniverseGraphicsGraphicIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get universe graphics graphic id params
func (o *GetUniverseGraphicsGraphicIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get universe graphics graphic id params
func (o *GetUniverseGraphicsGraphicIDParams) WithXUserAgent(xUserAgent *string) *GetUniverseGraphicsGraphicIDParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get universe graphics graphic id params
func (o *GetUniverseGraphicsGraphicIDParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithDatasource adds the datasource to the get universe graphics graphic id params
func (o *GetUniverseGraphicsGraphicIDParams) WithDatasource(datasource *string) *GetUniverseGraphicsGraphicIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get universe graphics graphic id params
func (o *GetUniverseGraphicsGraphicIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithGraphicID adds the graphicID to the get universe graphics graphic id params
func (o *GetUniverseGraphicsGraphicIDParams) WithGraphicID(graphicID int32) *GetUniverseGraphicsGraphicIDParams {
	o.SetGraphicID(graphicID)
	return o
}

// SetGraphicID adds the graphicId to the get universe graphics graphic id params
func (o *GetUniverseGraphicsGraphicIDParams) SetGraphicID(graphicID int32) {
	o.GraphicID = graphicID
}

// WithUserAgent adds the userAgent to the get universe graphics graphic id params
func (o *GetUniverseGraphicsGraphicIDParams) WithUserAgent(userAgent *string) *GetUniverseGraphicsGraphicIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get universe graphics graphic id params
func (o *GetUniverseGraphicsGraphicIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetUniverseGraphicsGraphicIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.XUserAgent != nil {

		// header param X-User-Agent
		if err := r.SetHeaderParam("X-User-Agent", *o.XUserAgent); err != nil {
			return err
		}

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

	// path param graphic_id
	if err := r.SetPathParam("graphic_id", swag.FormatInt32(o.GraphicID)); err != nil {
		return err
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
