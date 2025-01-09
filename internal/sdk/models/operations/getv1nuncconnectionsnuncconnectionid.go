// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type GetV1NuncConnectionsNuncConnectionIDRequest struct {
	NuncConnectionID string `pathParam:"style=simple,explode=false,name=nunc_connection_id"`
}

func (o *GetV1NuncConnectionsNuncConnectionIDRequest) GetNuncConnectionID() string {
	if o == nil {
		return ""
	}
	return o.NuncConnectionID
}

type GetV1NuncConnectionsNuncConnectionIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Retrieve the information displayed as part of your FireHydrant hosted status page.
	NuncConnectionEntity *shared.NuncConnectionEntity
}

func (o *GetV1NuncConnectionsNuncConnectionIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1NuncConnectionsNuncConnectionIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1NuncConnectionsNuncConnectionIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1NuncConnectionsNuncConnectionIDResponse) GetNuncConnectionEntity() *shared.NuncConnectionEntity {
	if o == nil {
		return nil
	}
	return o.NuncConnectionEntity
}
