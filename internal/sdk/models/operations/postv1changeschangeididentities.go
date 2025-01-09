// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type PostV1ChangesChangeIDIdentitiesRequest struct {
	ChangeID                        string                                 `pathParam:"style=simple,explode=false,name=change_id"`
	PostV1ChangesChangeIDIdentities shared.PostV1ChangesChangeIDIdentities `request:"mediaType=application/json"`
}

func (o *PostV1ChangesChangeIDIdentitiesRequest) GetChangeID() string {
	if o == nil {
		return ""
	}
	return o.ChangeID
}

func (o *PostV1ChangesChangeIDIdentitiesRequest) GetPostV1ChangesChangeIDIdentities() shared.PostV1ChangesChangeIDIdentities {
	if o == nil {
		return shared.PostV1ChangesChangeIDIdentities{}
	}
	return o.PostV1ChangesChangeIDIdentities
}

type PostV1ChangesChangeIDIdentitiesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Create an identity for this change
	ChangeIdentityEntity *shared.ChangeIdentityEntity
	// Bad Request
	ErrorEntity *shared.ErrorEntity
}

func (o *PostV1ChangesChangeIDIdentitiesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostV1ChangesChangeIDIdentitiesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostV1ChangesChangeIDIdentitiesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostV1ChangesChangeIDIdentitiesResponse) GetChangeIdentityEntity() *shared.ChangeIdentityEntity {
	if o == nil {
		return nil
	}
	return o.ChangeIdentityEntity
}

func (o *PostV1ChangesChangeIDIdentitiesResponse) GetErrorEntity() *shared.ErrorEntity {
	if o == nil {
		return nil
	}
	return o.ErrorEntity
}