// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDRequest struct {
	ID         string `pathParam:"style=simple,explode=false,name=id"`
	TeamID     string `pathParam:"style=simple,explode=false,name=team_id"`
	ScheduleID string `pathParam:"style=simple,explode=false,name=schedule_id"`
}

func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDRequest) GetScheduleID() string {
	if o == nil {
		return ""
	}
	return o.ScheduleID
}

type GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
