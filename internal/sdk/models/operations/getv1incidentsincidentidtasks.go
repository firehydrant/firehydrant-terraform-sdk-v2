// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type GetV1IncidentsIncidentIDTasksRequest struct {
	Page       *int   `queryParam:"style=form,explode=true,name=page"`
	PerPage    *int   `queryParam:"style=form,explode=true,name=per_page"`
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (o *GetV1IncidentsIncidentIDTasksRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1IncidentsIncidentIDTasksRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *GetV1IncidentsIncidentIDTasksRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

type GetV1IncidentsIncidentIDTasksResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Retrieve a list of all tasks for a specific incident
	TaskEntityPaginated *shared.TaskEntityPaginated
}

func (o *GetV1IncidentsIncidentIDTasksResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1IncidentsIncidentIDTasksResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1IncidentsIncidentIDTasksResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1IncidentsIncidentIDTasksResponse) GetTaskEntityPaginated() *shared.TaskEntityPaginated {
	if o == nil {
		return nil
	}
	return o.TaskEntityPaginated
}
