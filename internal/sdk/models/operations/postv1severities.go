// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type PostV1SeveritiesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Create a new severity
	SeverityEntity *shared.SeverityEntity
}

func (o *PostV1SeveritiesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostV1SeveritiesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostV1SeveritiesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostV1SeveritiesResponse) GetSeverityEntity() *shared.SeverityEntity {
	if o == nil {
		return nil
	}
	return o.SeverityEntity
}
