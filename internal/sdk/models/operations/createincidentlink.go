// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type CreateIncidentLinkRequest struct {
	IncidentID                     string                                `pathParam:"style=simple,explode=false,name=incident_id"`
	PostV1IncidentsIncidentIDLinks shared.PostV1IncidentsIncidentIDLinks `request:"mediaType=application/json"`
}

func (o *CreateIncidentLinkRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *CreateIncidentLinkRequest) GetPostV1IncidentsIncidentIDLinks() shared.PostV1IncidentsIncidentIDLinks {
	if o == nil {
		return shared.PostV1IncidentsIncidentIDLinks{}
	}
	return o.PostV1IncidentsIncidentIDLinks
}

type CreateIncidentLinkResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Allows adding adhoc links to an incident as an attachment
	AttachmentsLinkEntity *shared.AttachmentsLinkEntity
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

func (o *CreateIncidentLinkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateIncidentLinkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateIncidentLinkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateIncidentLinkResponse) GetAttachmentsLinkEntity() *shared.AttachmentsLinkEntity {
	if o == nil {
		return nil
	}
	return o.AttachmentsLinkEntity
}

func (o *CreateIncidentLinkResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *CreateIncidentLinkResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *CreateIncidentLinkResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *CreateIncidentLinkResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *CreateIncidentLinkResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *CreateIncidentLinkResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}
