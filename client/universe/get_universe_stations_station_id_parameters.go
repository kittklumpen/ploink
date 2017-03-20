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

// NewGetUniverseStationsStationIDParams creates a new GetUniverseStationsStationIDParams object
// with the default values initialized.
func NewGetUniverseStationsStationIDParams() *GetUniverseStationsStationIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseStationsStationIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUniverseStationsStationIDParamsWithTimeout creates a new GetUniverseStationsStationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUniverseStationsStationIDParamsWithTimeout(timeout time.Duration) *GetUniverseStationsStationIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseStationsStationIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetUniverseStationsStationIDParamsWithContext creates a new GetUniverseStationsStationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUniverseStationsStationIDParamsWithContext(ctx context.Context) *GetUniverseStationsStationIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseStationsStationIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetUniverseStationsStationIDParamsWithHTTPClient creates a new GetUniverseStationsStationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUniverseStationsStationIDParamsWithHTTPClient(client *http.Client) *GetUniverseStationsStationIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseStationsStationIDParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetUniverseStationsStationIDParams contains all the parameters to send to the API endpoint
for the get universe stations station id operation typically these are written to a http.Request
*/
type GetUniverseStationsStationIDParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*StationID
	  station_id integer

	*/
	StationID int32
	/*UserAgent
	  Client identifier, takes precedence over headers

	*/
	UserAgent *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get universe stations station id params
func (o *GetUniverseStationsStationIDParams) WithTimeout(timeout time.Duration) *GetUniverseStationsStationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get universe stations station id params
func (o *GetUniverseStationsStationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get universe stations station id params
func (o *GetUniverseStationsStationIDParams) WithContext(ctx context.Context) *GetUniverseStationsStationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get universe stations station id params
func (o *GetUniverseStationsStationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get universe stations station id params
func (o *GetUniverseStationsStationIDParams) WithHTTPClient(client *http.Client) *GetUniverseStationsStationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get universe stations station id params
func (o *GetUniverseStationsStationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get universe stations station id params
func (o *GetUniverseStationsStationIDParams) WithXUserAgent(xUserAgent *string) *GetUniverseStationsStationIDParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get universe stations station id params
func (o *GetUniverseStationsStationIDParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithDatasource adds the datasource to the get universe stations station id params
func (o *GetUniverseStationsStationIDParams) WithDatasource(datasource *string) *GetUniverseStationsStationIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get universe stations station id params
func (o *GetUniverseStationsStationIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithStationID adds the stationID to the get universe stations station id params
func (o *GetUniverseStationsStationIDParams) WithStationID(stationID int32) *GetUniverseStationsStationIDParams {
	o.SetStationID(stationID)
	return o
}

// SetStationID adds the stationId to the get universe stations station id params
func (o *GetUniverseStationsStationIDParams) SetStationID(stationID int32) {
	o.StationID = stationID
}

// WithUserAgent adds the userAgent to the get universe stations station id params
func (o *GetUniverseStationsStationIDParams) WithUserAgent(userAgent *string) *GetUniverseStationsStationIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get universe stations station id params
func (o *GetUniverseStationsStationIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetUniverseStationsStationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param station_id
	if err := r.SetPathParam("station_id", swag.FormatInt32(o.StationID)); err != nil {
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
