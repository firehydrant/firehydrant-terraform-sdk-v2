// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type GetV1SeveritiesSeveritySlugRequest struct {
	SeveritySlug string `pathParam:"style=simple,explode=false,name=severity_slug"`
}

func (o *GetV1SeveritiesSeveritySlugRequest) GetSeveritySlug() string {
	if o == nil {
		return ""
	}
	return o.SeveritySlug
}

type GetV1SeveritiesSeveritySlugResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Retrieve a specific severity
	SeverityEntity *shared.SeverityEntity
}

func (o *GetV1SeveritiesSeveritySlugResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1SeveritiesSeveritySlugResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1SeveritiesSeveritySlugResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1SeveritiesSeveritySlugResponse) GetSeverityEntity() *shared.SeverityEntity {
	if o == nil {
		return nil
	}
	return o.SeverityEntity
}
