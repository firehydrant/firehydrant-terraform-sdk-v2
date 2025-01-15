// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type PostV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsRequest struct {
	TicketingProjectID string `pathParam:"style=simple,explode=false,name=ticketing_project_id"`
}

func (o *PostV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsRequest) GetTicketingProjectID() string {
	if o == nil {
		return ""
	}
	return o.TicketingProjectID
}

type PostV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Creates configuration for a ticketing project
	TicketingProjectConfigEntity *shared.TicketingProjectConfigEntity
}

func (o *PostV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsResponse) GetTicketingProjectConfigEntity() *shared.TicketingProjectConfigEntity {
	if o == nil {
		return nil
	}
	return o.TicketingProjectConfigEntity
}
