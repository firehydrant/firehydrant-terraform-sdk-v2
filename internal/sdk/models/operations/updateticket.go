// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type UpdateTicketRequest struct {
	TicketID                        string                                 `pathParam:"style=simple,explode=false,name=ticket_id"`
	PatchV1TicketingTicketsTicketID shared.PatchV1TicketingTicketsTicketID `request:"mediaType=application/json"`
}

func (o *UpdateTicketRequest) GetTicketID() string {
	if o == nil {
		return ""
	}
	return o.TicketID
}

func (o *UpdateTicketRequest) GetPatchV1TicketingTicketsTicketID() shared.PatchV1TicketingTicketsTicketID {
	if o == nil {
		return shared.PatchV1TicketingTicketsTicketID{}
	}
	return o.PatchV1TicketingTicketsTicketID
}

type UpdateTicketResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Update a ticket's attributes
	TicketingTicketEntity *shared.TicketingTicketEntity
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

func (o *UpdateTicketResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateTicketResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateTicketResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateTicketResponse) GetTicketingTicketEntity() *shared.TicketingTicketEntity {
	if o == nil {
		return nil
	}
	return o.TicketingTicketEntity
}

func (o *UpdateTicketResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *UpdateTicketResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *UpdateTicketResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *UpdateTicketResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *UpdateTicketResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *UpdateTicketResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}
