// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type DeleteV1ChangesChangeIDIdentitiesIdentityIDRequest struct {
	IdentityID string `pathParam:"style=simple,explode=false,name=identity_id"`
	ChangeID   string `pathParam:"style=simple,explode=false,name=change_id"`
}

func (o *DeleteV1ChangesChangeIDIdentitiesIdentityIDRequest) GetIdentityID() string {
	if o == nil {
		return ""
	}
	return o.IdentityID
}

func (o *DeleteV1ChangesChangeIDIdentitiesIdentityIDRequest) GetChangeID() string {
	if o == nil {
		return ""
	}
	return o.ChangeID
}

type DeleteV1ChangesChangeIDIdentitiesIdentityIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Bad Request
	ErrorEntity *shared.ErrorEntity
}

func (o *DeleteV1ChangesChangeIDIdentitiesIdentityIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteV1ChangesChangeIDIdentitiesIdentityIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteV1ChangesChangeIDIdentitiesIdentityIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteV1ChangesChangeIDIdentitiesIdentityIDResponse) GetErrorEntity() *shared.ErrorEntity {
	if o == nil {
		return nil
	}
	return o.ErrorEntity
}