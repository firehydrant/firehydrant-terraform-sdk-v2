// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type UpdateTeamScheduleShiftRequest struct {
	ID                                                  string                                                     `pathParam:"style=simple,explode=false,name=id"`
	TeamID                                              string                                                     `pathParam:"style=simple,explode=false,name=team_id"`
	ScheduleID                                          string                                                     `pathParam:"style=simple,explode=false,name=schedule_id"`
	PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID shared.PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID `request:"mediaType=application/json"`
}

func (o *UpdateTeamScheduleShiftRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateTeamScheduleShiftRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *UpdateTeamScheduleShiftRequest) GetScheduleID() string {
	if o == nil {
		return ""
	}
	return o.ScheduleID
}

func (o *UpdateTeamScheduleShiftRequest) GetPatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID() shared.PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID {
	if o == nil {
		return shared.PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID{}
	}
	return o.PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID
}

type UpdateTeamScheduleShiftResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
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

func (o *UpdateTeamScheduleShiftResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateTeamScheduleShiftResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateTeamScheduleShiftResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateTeamScheduleShiftResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *UpdateTeamScheduleShiftResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *UpdateTeamScheduleShiftResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *UpdateTeamScheduleShiftResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *UpdateTeamScheduleShiftResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *UpdateTeamScheduleShiftResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}
