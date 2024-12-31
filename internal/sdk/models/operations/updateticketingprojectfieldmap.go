// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type UpdateTicketingProjectFieldMapRequest struct {
	MapID              string `pathParam:"style=simple,explode=false,name=map_id"`
	TicketingProjectID string `pathParam:"style=simple,explode=false,name=ticketing_project_id"`
}

func (o *UpdateTicketingProjectFieldMapRequest) GetMapID() string {
	if o == nil {
		return ""
	}
	return o.MapID
}

func (o *UpdateTicketingProjectFieldMapRequest) GetTicketingProjectID() string {
	if o == nil {
		return ""
	}
	return o.TicketingProjectID
}

type UpdateTicketingProjectFieldMapResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Update field map for a ticketing project
	TicketingProjectFieldMapEntity *shared.TicketingProjectFieldMapEntity
	// A collection of codes that generally means the end user got something wrong in making the request
	BadRequest *shared.BadRequest
	// A collection of codes that generally means the client was not authenticated correctly for the request they want to make
	Unauthorized *shared.Unauthorized
	// Status codes relating to the resource/entity they are requesting not being found or endpoints/routes not existing
	NotFound *shared.NotFound
	// Status codes relating to the client being rate limited by the server
	RateLimited *shared.RateLimited
	// A collection of status codes that generally mean the server failed in an unexpected way
	InternalServerError *shared.InternalServerError
	// Timeouts occurred with the request
	Timeout *shared.Timeout
}

func (o *UpdateTicketingProjectFieldMapResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateTicketingProjectFieldMapResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateTicketingProjectFieldMapResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateTicketingProjectFieldMapResponse) GetTicketingProjectFieldMapEntity() *shared.TicketingProjectFieldMapEntity {
	if o == nil {
		return nil
	}
	return o.TicketingProjectFieldMapEntity
}

func (o *UpdateTicketingProjectFieldMapResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *UpdateTicketingProjectFieldMapResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *UpdateTicketingProjectFieldMapResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *UpdateTicketingProjectFieldMapResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *UpdateTicketingProjectFieldMapResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *UpdateTicketingProjectFieldMapResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}
