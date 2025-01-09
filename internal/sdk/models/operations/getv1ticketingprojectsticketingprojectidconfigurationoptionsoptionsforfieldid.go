// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetV1TicketingProjectsTicketingProjectIDConfigurationOptionsOptionsForFieldIDRequest struct {
	FieldID            string `pathParam:"style=simple,explode=false,name=field_id"`
	TicketingProjectID string `pathParam:"style=simple,explode=false,name=ticketing_project_id"`
}

func (o *GetV1TicketingProjectsTicketingProjectIDConfigurationOptionsOptionsForFieldIDRequest) GetFieldID() string {
	if o == nil {
		return ""
	}
	return o.FieldID
}

func (o *GetV1TicketingProjectsTicketingProjectIDConfigurationOptionsOptionsForFieldIDRequest) GetTicketingProjectID() string {
	if o == nil {
		return ""
	}
	return o.TicketingProjectID
}

type GetV1TicketingProjectsTicketingProjectIDConfigurationOptionsOptionsForFieldIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetV1TicketingProjectsTicketingProjectIDConfigurationOptionsOptionsForFieldIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1TicketingProjectsTicketingProjectIDConfigurationOptionsOptionsForFieldIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1TicketingProjectsTicketingProjectIDConfigurationOptionsOptionsForFieldIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
