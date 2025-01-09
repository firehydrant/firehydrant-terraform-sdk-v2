// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type DeleteV1IncidentsIncidentIDEventsEventIDRequest struct {
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
	EventID    string `pathParam:"style=simple,explode=false,name=event_id"`
}

func (o *DeleteV1IncidentsIncidentIDEventsEventIDRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *DeleteV1IncidentsIncidentIDEventsEventIDRequest) GetEventID() string {
	if o == nil {
		return ""
	}
	return o.EventID
}

type DeleteV1IncidentsIncidentIDEventsEventIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Delete a single event for an incident
	IncidentEventEntity *shared.IncidentEventEntity
}

func (o *DeleteV1IncidentsIncidentIDEventsEventIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteV1IncidentsIncidentIDEventsEventIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteV1IncidentsIncidentIDEventsEventIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteV1IncidentsIncidentIDEventsEventIDResponse) GetIncidentEventEntity() *shared.IncidentEventEntity {
	if o == nil {
		return nil
	}
	return o.IncidentEventEntity
}
