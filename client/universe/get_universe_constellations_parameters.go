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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetUniverseConstellationsParams creates a new GetUniverseConstellationsParams object
// with the default values initialized.
func NewGetUniverseConstellationsParams() *GetUniverseConstellationsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseConstellationsParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUniverseConstellationsParamsWithTimeout creates a new GetUniverseConstellationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUniverseConstellationsParamsWithTimeout(timeout time.Duration) *GetUniverseConstellationsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseConstellationsParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetUniverseConstellationsParamsWithContext creates a new GetUniverseConstellationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUniverseConstellationsParamsWithContext(ctx context.Context) *GetUniverseConstellationsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseConstellationsParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetUniverseConstellationsParamsWithHTTPClient creates a new GetUniverseConstellationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUniverseConstellationsParamsWithHTTPClient(client *http.Client) *GetUniverseConstellationsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseConstellationsParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetUniverseConstellationsParams contains all the parameters to send to the API endpoint
for the get universe constellations operation typically these are written to a http.Request
*/
type GetUniverseConstellationsParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
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

// WithTimeout adds the timeout to the get universe constellations params
func (o *GetUniverseConstellationsParams) WithTimeout(timeout time.Duration) *GetUniverseConstellationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get universe constellations params
func (o *GetUniverseConstellationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get universe constellations params
func (o *GetUniverseConstellationsParams) WithContext(ctx context.Context) *GetUniverseConstellationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get universe constellations params
func (o *GetUniverseConstellationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get universe constellations params
func (o *GetUniverseConstellationsParams) WithHTTPClient(client *http.Client) *GetUniverseConstellationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get universe constellations params
func (o *GetUniverseConstellationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get universe constellations params
func (o *GetUniverseConstellationsParams) WithXUserAgent(xUserAgent *string) *GetUniverseConstellationsParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get universe constellations params
func (o *GetUniverseConstellationsParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithDatasource adds the datasource to the get universe constellations params
func (o *GetUniverseConstellationsParams) WithDatasource(datasource *string) *GetUniverseConstellationsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get universe constellations params
func (o *GetUniverseConstellationsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithUserAgent adds the userAgent to the get universe constellations params
func (o *GetUniverseConstellationsParams) WithUserAgent(userAgent *string) *GetUniverseConstellationsParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get universe constellations params
func (o *GetUniverseConstellationsParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetUniverseConstellationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
