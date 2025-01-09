// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type DeleteV1NuncConnectionsNuncConnectionIDImagesTypeRequest struct {
	NuncConnectionID string `pathParam:"style=simple,explode=false,name=nunc_connection_id"`
	Type             string `pathParam:"style=simple,explode=false,name=type"`
}

func (o *DeleteV1NuncConnectionsNuncConnectionIDImagesTypeRequest) GetNuncConnectionID() string {
	if o == nil {
		return ""
	}
	return o.NuncConnectionID
}

func (o *DeleteV1NuncConnectionsNuncConnectionIDImagesTypeRequest) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

type DeleteV1NuncConnectionsNuncConnectionIDImagesTypeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Delete an image attached to a FireHydrant status page
	NuncConnectionEntity *shared.NuncConnectionEntity
}

func (o *DeleteV1NuncConnectionsNuncConnectionIDImagesTypeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteV1NuncConnectionsNuncConnectionIDImagesTypeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteV1NuncConnectionsNuncConnectionIDImagesTypeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteV1NuncConnectionsNuncConnectionIDImagesTypeResponse) GetNuncConnectionEntity() *shared.NuncConnectionEntity {
	if o == nil {
		return nil
	}
	return o.NuncConnectionEntity
}
