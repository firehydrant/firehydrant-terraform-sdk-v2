// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type GetV1TicketingProjectsTicketingProjectIDRequest struct {
	TicketingProjectID string `pathParam:"style=simple,explode=false,name=ticketing_project_id"`
}

func (o *GetV1TicketingProjectsTicketingProjectIDRequest) GetTicketingProjectID() string {
	if o == nil {
		return ""
	}
	return o.TicketingProjectID
}

type GetV1TicketingProjectsTicketingProjectIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Retrieve a single ticketing project by ID
	TicketingProjectsProjectListItemEntity *shared.TicketingProjectsProjectListItemEntity
}

func (o *GetV1TicketingProjectsTicketingProjectIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1TicketingProjectsTicketingProjectIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1TicketingProjectsTicketingProjectIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1TicketingProjectsTicketingProjectIDResponse) GetTicketingProjectsProjectListItemEntity() *shared.TicketingProjectsProjectListItemEntity {
	if o == nil {
		return nil
	}
	return o.TicketingProjectsProjectListItemEntity
}
