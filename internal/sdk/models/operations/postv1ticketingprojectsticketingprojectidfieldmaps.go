// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type PostV1TicketingProjectsTicketingProjectIDFieldMapsRequest struct {
	TicketingProjectID string `pathParam:"style=simple,explode=false,name=ticketing_project_id"`
}

func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsRequest) GetTicketingProjectID() string {
	if o == nil {
		return ""
	}
	return o.TicketingProjectID
}

type PostV1TicketingProjectsTicketingProjectIDFieldMapsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Creates field map for a ticketing project
	TicketingProjectFieldMapEntity *shared.TicketingProjectFieldMapEntity
}

func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsResponse) GetTicketingProjectFieldMapEntity() *shared.TicketingProjectFieldMapEntity {
	if o == nil {
		return nil
	}
	return o.TicketingProjectFieldMapEntity
}
