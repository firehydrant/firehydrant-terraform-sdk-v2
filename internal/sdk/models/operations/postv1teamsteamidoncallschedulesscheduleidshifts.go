// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type PostV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsRequest struct {
	TeamID                                           string                                                  `pathParam:"style=simple,explode=false,name=team_id"`
	ScheduleID                                       string                                                  `pathParam:"style=simple,explode=false,name=schedule_id"`
	PostV1TeamsTeamIDOnCallSchedulesScheduleIDShifts shared.PostV1TeamsTeamIDOnCallSchedulesScheduleIDShifts `request:"mediaType=application/json"`
}

func (o *PostV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *PostV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsRequest) GetScheduleID() string {
	if o == nil {
		return ""
	}
	return o.ScheduleID
}

func (o *PostV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsRequest) GetPostV1TeamsTeamIDOnCallSchedulesScheduleIDShifts() shared.PostV1TeamsTeamIDOnCallSchedulesScheduleIDShifts {
	if o == nil {
		return shared.PostV1TeamsTeamIDOnCallSchedulesScheduleIDShifts{}
	}
	return o.PostV1TeamsTeamIDOnCallSchedulesScheduleIDShifts
}

type PostV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
