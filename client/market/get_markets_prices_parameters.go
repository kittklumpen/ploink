package market

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

// NewGetMarketsPricesParams creates a new GetMarketsPricesParams object
// with the default values initialized.
func NewGetMarketsPricesParams() *GetMarketsPricesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetMarketsPricesParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMarketsPricesParamsWithTimeout creates a new GetMarketsPricesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMarketsPricesParamsWithTimeout(timeout time.Duration) *GetMarketsPricesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetMarketsPricesParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetMarketsPricesParamsWithContext creates a new GetMarketsPricesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMarketsPricesParamsWithContext(ctx context.Context) *GetMarketsPricesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetMarketsPricesParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetMarketsPricesParamsWithHTTPClient creates a new GetMarketsPricesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMarketsPricesParamsWithHTTPClient(client *http.Client) *GetMarketsPricesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetMarketsPricesParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetMarketsPricesParams contains all the parameters to send to the API endpoint
for the get markets prices operation typically these are written to a http.Request
*/
type GetMarketsPricesParams struct {

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

// WithTimeout adds the timeout to the get markets prices params
func (o *GetMarketsPricesParams) WithTimeout(timeout time.Duration) *GetMarketsPricesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get markets prices params
func (o *GetMarketsPricesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get markets prices params
func (o *GetMarketsPricesParams) WithContext(ctx context.Context) *GetMarketsPricesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get markets prices params
func (o *GetMarketsPricesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get markets prices params
func (o *GetMarketsPricesParams) WithHTTPClient(client *http.Client) *GetMarketsPricesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get markets prices params
func (o *GetMarketsPricesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get markets prices params
func (o *GetMarketsPricesParams) WithXUserAgent(xUserAgent *string) *GetMarketsPricesParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get markets prices params
func (o *GetMarketsPricesParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithDatasource adds the datasource to the get markets prices params
func (o *GetMarketsPricesParams) WithDatasource(datasource *string) *GetMarketsPricesParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get markets prices params
func (o *GetMarketsPricesParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithUserAgent adds the userAgent to the get markets prices params
func (o *GetMarketsPricesParams) WithUserAgent(userAgent *string) *GetMarketsPricesParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get markets prices params
func (o *GetMarketsPricesParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetMarketsPricesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
