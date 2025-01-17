// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type PostV1IncidentsIncidentIDTasksRequest struct {
	IncidentID                     string                                `pathParam:"style=simple,explode=false,name=incident_id"`
	PostV1IncidentsIncidentIDTasks shared.PostV1IncidentsIncidentIDTasks `request:"mediaType=application/json"`
}

func (o *PostV1IncidentsIncidentIDTasksRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *PostV1IncidentsIncidentIDTasksRequest) GetPostV1IncidentsIncidentIDTasks() shared.PostV1IncidentsIncidentIDTasks {
	if o == nil {
		return shared.PostV1IncidentsIncidentIDTasks{}
	}
	return o.PostV1IncidentsIncidentIDTasks
}

type PostV1IncidentsIncidentIDTasksResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Create a task
	TaskEntity *shared.TaskEntity
}

func (o *PostV1IncidentsIncidentIDTasksResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostV1IncidentsIncidentIDTasksResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostV1IncidentsIncidentIDTasksResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostV1IncidentsIncidentIDTasksResponse) GetTaskEntity() *shared.TaskEntity {
	if o == nil {
		return nil
	}
	return o.TaskEntity
}
