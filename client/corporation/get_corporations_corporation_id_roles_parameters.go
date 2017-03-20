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

// NewGetCorporationsCorporationIDRolesParams creates a new GetCorporationsCorporationIDRolesParams object
// with the default values initialized.
func NewGetCorporationsCorporationIDRolesParams() *GetCorporationsCorporationIDRolesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDRolesParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCorporationsCorporationIDRolesParamsWithTimeout creates a new GetCorporationsCorporationIDRolesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCorporationsCorporationIDRolesParamsWithTimeout(timeout time.Duration) *GetCorporationsCorporationIDRolesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDRolesParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCorporationsCorporationIDRolesParamsWithContext creates a new GetCorporationsCorporationIDRolesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCorporationsCorporationIDRolesParamsWithContext(ctx context.Context) *GetCorporationsCorporationIDRolesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDRolesParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetCorporationsCorporationIDRolesParamsWithHTTPClient creates a new GetCorporationsCorporationIDRolesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCorporationsCorporationIDRolesParamsWithHTTPClient(client *http.Client) *GetCorporationsCorporationIDRolesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDRolesParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetCorporationsCorporationIDRolesParams contains all the parameters to send to the API endpoint
for the get corporations corporation id roles operation typically these are written to a http.Request
*/
type GetCorporationsCorporationIDRolesParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
	/*CorporationID
	  A corporation ID

	*/
	CorporationID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*Token
	  Access token to use, if preferred over a header

	*/
	Token *string
	/*UserAgent
	  Client identifier, takes precedence over headers

	*/
	UserAgent *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) WithTimeout(timeout time.Duration) *GetCorporationsCorporationIDRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) WithContext(ctx context.Context) *GetCorporationsCorporationIDRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) WithHTTPClient(client *http.Client) *GetCorporationsCorporationIDRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) WithXUserAgent(xUserAgent *string) *GetCorporationsCorporationIDRolesParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithCorporationID adds the corporationID to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) WithCorporationID(corporationID int32) *GetCorporationsCorporationIDRolesParams {
	o.SetCorporationID(corporationID)
	return o
}

// SetCorporationID adds the corporationId to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) SetCorporationID(corporationID int32) {
	o.CorporationID = corporationID
}

// WithDatasource adds the datasource to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) WithDatasource(datasource *string) *GetCorporationsCorporationIDRolesParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithToken adds the token to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) WithToken(token *string) *GetCorporationsCorporationIDRolesParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) SetToken(token *string) {
	o.Token = token
}

// WithUserAgent adds the userAgent to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) WithUserAgent(userAgent *string) *GetCorporationsCorporationIDRolesParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetCorporationsCorporationIDRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Token != nil {

		// query param token
		var qrToken string
		if o.Token != nil {
			qrToken = *o.Token
		}
		qToken := qrToken
		if qToken != "" {
			if err := r.SetQueryParam("token", qToken); err != nil {
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
