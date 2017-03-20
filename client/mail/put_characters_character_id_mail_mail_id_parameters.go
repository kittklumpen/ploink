package mail

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

// NewPutCharactersCharacterIDMailMailIDParams creates a new PutCharactersCharacterIDMailMailIDParams object
// with the default values initialized.
func NewPutCharactersCharacterIDMailMailIDParams() *PutCharactersCharacterIDMailMailIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PutCharactersCharacterIDMailMailIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCharactersCharacterIDMailMailIDParamsWithTimeout creates a new PutCharactersCharacterIDMailMailIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCharactersCharacterIDMailMailIDParamsWithTimeout(timeout time.Duration) *PutCharactersCharacterIDMailMailIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PutCharactersCharacterIDMailMailIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewPutCharactersCharacterIDMailMailIDParamsWithContext creates a new PutCharactersCharacterIDMailMailIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCharactersCharacterIDMailMailIDParamsWithContext(ctx context.Context) *PutCharactersCharacterIDMailMailIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PutCharactersCharacterIDMailMailIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewPutCharactersCharacterIDMailMailIDParamsWithHTTPClient creates a new PutCharactersCharacterIDMailMailIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCharactersCharacterIDMailMailIDParamsWithHTTPClient(client *http.Client) *PutCharactersCharacterIDMailMailIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PutCharactersCharacterIDMailMailIDParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*PutCharactersCharacterIDMailMailIDParams contains all the parameters to send to the API endpoint
for the put characters character id mail mail id operation typically these are written to a http.Request
*/
type PutCharactersCharacterIDMailMailIDParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
	/*CharacterID
	  An EVE character ID

	*/
	CharacterID int32
	/*Contents
	  Data used to update the mail

	*/
	Contents PutCharactersCharacterIDMailMailIDBody
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*MailID
	  An EVE mail ID

	*/
	MailID int32
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

// WithTimeout adds the timeout to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) WithTimeout(timeout time.Duration) *PutCharactersCharacterIDMailMailIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) WithContext(ctx context.Context) *PutCharactersCharacterIDMailMailIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) WithHTTPClient(client *http.Client) *PutCharactersCharacterIDMailMailIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) WithXUserAgent(xUserAgent *string) *PutCharactersCharacterIDMailMailIDParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithCharacterID adds the characterID to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) WithCharacterID(characterID int32) *PutCharactersCharacterIDMailMailIDParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithContents adds the contents to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) WithContents(contents PutCharactersCharacterIDMailMailIDBody) *PutCharactersCharacterIDMailMailIDParams {
	o.SetContents(contents)
	return o
}

// SetContents adds the contents to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) SetContents(contents PutCharactersCharacterIDMailMailIDBody) {
	o.Contents = contents
}

// WithDatasource adds the datasource to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) WithDatasource(datasource *string) *PutCharactersCharacterIDMailMailIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithMailID adds the mailID to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) WithMailID(mailID int32) *PutCharactersCharacterIDMailMailIDParams {
	o.SetMailID(mailID)
	return o
}

// SetMailID adds the mailId to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) SetMailID(mailID int32) {
	o.MailID = mailID
}

// WithToken adds the token to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) WithToken(token *string) *PutCharactersCharacterIDMailMailIDParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) SetToken(token *string) {
	o.Token = token
}

// WithUserAgent adds the userAgent to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) WithUserAgent(userAgent *string) *PutCharactersCharacterIDMailMailIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *PutCharactersCharacterIDMailMailIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.XUserAgent != nil {

		// header param X-User-Agent
		if err := r.SetHeaderParam("X-User-Agent", *o.XUserAgent); err != nil {
			return err
		}

	}

	// path param character_id
	if err := r.SetPathParam("character_id", swag.FormatInt32(o.CharacterID)); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Contents); err != nil {
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

	// path param mail_id
	if err := r.SetPathParam("mail_id", swag.FormatInt32(o.MailID)); err != nil {
		return err
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
