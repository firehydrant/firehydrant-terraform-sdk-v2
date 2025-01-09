// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type PostV1ServicesServiceLinksResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Creates a service with the appropriate integration for each external service ID passed
	ServiceLinkEntities []shared.ServiceLinkEntity
}

func (o *PostV1ServicesServiceLinksResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostV1ServicesServiceLinksResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostV1ServicesServiceLinksResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostV1ServicesServiceLinksResponse) GetServiceLinkEntities() []shared.ServiceLinkEntity {
	if o == nil {
		return nil
	}
	return o.ServiceLinkEntities
}