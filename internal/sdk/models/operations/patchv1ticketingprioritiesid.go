// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type PatchV1TicketingPrioritiesIDRequest struct {
	ID                           string                              `pathParam:"style=simple,explode=false,name=id"`
	PatchV1TicketingPrioritiesID shared.PatchV1TicketingPrioritiesID `request:"mediaType=application/json"`
}

func (o *PatchV1TicketingPrioritiesIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PatchV1TicketingPrioritiesIDRequest) GetPatchV1TicketingPrioritiesID() shared.PatchV1TicketingPrioritiesID {
	if o == nil {
		return shared.PatchV1TicketingPrioritiesID{}
	}
	return o.PatchV1TicketingPrioritiesID
}

type PatchV1TicketingPrioritiesIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Update a single ticketing priority's attributes
	TicketingPriorityEntity *shared.TicketingPriorityEntity
}

func (o *PatchV1TicketingPrioritiesIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchV1TicketingPrioritiesIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchV1TicketingPrioritiesIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchV1TicketingPrioritiesIDResponse) GetTicketingPriorityEntity() *shared.TicketingPriorityEntity {
	if o == nil {
		return nil
	}
	return o.TicketingPriorityEntity
}
