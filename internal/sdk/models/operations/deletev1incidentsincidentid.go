// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type DeleteV1IncidentsIncidentIDRequest struct {
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (o *DeleteV1IncidentsIncidentIDRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

type DeleteV1IncidentsIncidentIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Archives an incident which will hide it from lists and metrics
	IncidentEntity *shared.IncidentEntity
}

func (o *DeleteV1IncidentsIncidentIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteV1IncidentsIncidentIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteV1IncidentsIncidentIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteV1IncidentsIncidentIDResponse) GetIncidentEntity() *shared.IncidentEntity {
	if o == nil {
		return nil
	}
	return o.IncidentEntity
}
