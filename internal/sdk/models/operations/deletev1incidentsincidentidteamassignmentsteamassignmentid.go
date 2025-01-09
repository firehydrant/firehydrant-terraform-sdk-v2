// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteV1IncidentsIncidentIDTeamAssignmentsTeamAssignmentIDRequestBody struct {
	// Team role assignments to unassign from the incident
	RoleAssignmentIds []string `form:"name=role_assignment_ids"`
}

func (o *DeleteV1IncidentsIncidentIDTeamAssignmentsTeamAssignmentIDRequestBody) GetRoleAssignmentIds() []string {
	if o == nil {
		return nil
	}
	return o.RoleAssignmentIds
}

type DeleteV1IncidentsIncidentIDTeamAssignmentsTeamAssignmentIDRequest struct {
	IncidentID       string                                                                 `pathParam:"style=simple,explode=false,name=incident_id"`
	TeamAssignmentID string                                                                 `pathParam:"style=simple,explode=false,name=team_assignment_id"`
	RequestBody      *DeleteV1IncidentsIncidentIDTeamAssignmentsTeamAssignmentIDRequestBody `request:"mediaType=application/x-www-form-urlencoded"`
}

func (o *DeleteV1IncidentsIncidentIDTeamAssignmentsTeamAssignmentIDRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *DeleteV1IncidentsIncidentIDTeamAssignmentsTeamAssignmentIDRequest) GetTeamAssignmentID() string {
	if o == nil {
		return ""
	}
	return o.TeamAssignmentID
}

func (o *DeleteV1IncidentsIncidentIDTeamAssignmentsTeamAssignmentIDRequest) GetRequestBody() *DeleteV1IncidentsIncidentIDTeamAssignmentsTeamAssignmentIDRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type DeleteV1IncidentsIncidentIDTeamAssignmentsTeamAssignmentIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteV1IncidentsIncidentIDTeamAssignmentsTeamAssignmentIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteV1IncidentsIncidentIDTeamAssignmentsTeamAssignmentIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteV1IncidentsIncidentIDTeamAssignmentsTeamAssignmentIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}