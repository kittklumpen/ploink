package contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new contacts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for contacts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteCharactersCharacterIDContacts deletes contacts

Bulk delete contacts

---

Alternate route: `/v1/characters/{character_id}/contacts/`

Alternate route: `/legacy/characters/{character_id}/contacts/`

Alternate route: `/latest/characters/{character_id}/contacts/`

*/
func (a *Client) DeleteCharactersCharacterIDContacts(params *DeleteCharactersCharacterIDContactsParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCharactersCharacterIDContactsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCharactersCharacterIDContactsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete_characters_character_id_contacts",
		Method:             "DELETE",
		PathPattern:        "/characters/{character_id}/contacts/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCharactersCharacterIDContactsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteCharactersCharacterIDContactsNoContent), nil

}

/*
GetCharactersCharacterIDContacts gets contacts

Return contacts of a character

---

Alternate route: `/v1/characters/{character_id}/contacts/`

Alternate route: `/legacy/characters/{character_id}/contacts/`

Alternate route: `/latest/characters/{character_id}/contacts/`


---

This route is cached for up to 300 seconds
*/
func (a *Client) GetCharactersCharacterIDContacts(params *GetCharactersCharacterIDContactsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCharactersCharacterIDContactsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDContactsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_characters_character_id_contacts",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/contacts/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDContactsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCharactersCharacterIDContactsOK), nil

}

/*
GetCharactersCharacterIDContactsLabels gets contact labels

Return custom labels for contacts the character defined

---

Alternate route: `/v1/characters/{character_id}/contacts/labels/`

Alternate route: `/legacy/characters/{character_id}/contacts/labels/`

Alternate route: `/latest/characters/{character_id}/contacts/labels/`


---

This route is cached for up to 300 seconds
*/
func (a *Client) GetCharactersCharacterIDContactsLabels(params *GetCharactersCharacterIDContactsLabelsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCharactersCharacterIDContactsLabelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDContactsLabelsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_characters_character_id_contacts_labels",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/contacts/labels/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDContactsLabelsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCharactersCharacterIDContactsLabelsOK), nil

}

/*
PostCharactersCharacterIDContacts adds contacts

Bulk add contacts with same settings

---

Alternate route: `/v1/characters/{character_id}/contacts/`

Alternate route: `/legacy/characters/{character_id}/contacts/`

Alternate route: `/latest/characters/{character_id}/contacts/`

*/
func (a *Client) PostCharactersCharacterIDContacts(params *PostCharactersCharacterIDContactsParams, authInfo runtime.ClientAuthInfoWriter) (*PostCharactersCharacterIDContactsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCharactersCharacterIDContactsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "post_characters_character_id_contacts",
		Method:             "POST",
		PathPattern:        "/characters/{character_id}/contacts/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCharactersCharacterIDContactsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostCharactersCharacterIDContactsCreated), nil

}

/*
PutCharactersCharacterIDContacts edits contacts

Bulk edit contacts with same settings

---

Alternate route: `/v1/characters/{character_id}/contacts/`

Alternate route: `/legacy/characters/{character_id}/contacts/`

Alternate route: `/latest/characters/{character_id}/contacts/`

*/
func (a *Client) PutCharactersCharacterIDContacts(params *PutCharactersCharacterIDContactsParams, authInfo runtime.ClientAuthInfoWriter) (*PutCharactersCharacterIDContactsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCharactersCharacterIDContactsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "put_characters_character_id_contacts",
		Method:             "PUT",
		PathPattern:        "/characters/{character_id}/contacts/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutCharactersCharacterIDContactsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutCharactersCharacterIDContactsNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
