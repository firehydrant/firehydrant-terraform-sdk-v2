// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetV1SignalsOnCallRequest struct {
	// An optional comma separated list of team IDs to filter currently on-call users by
	TeamID *string `queryParam:"style=form,explode=true,name=team_id"`
}

func (o *GetV1SignalsOnCallRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

type GetV1SignalsOnCallResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetV1SignalsOnCallResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1SignalsOnCallResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1SignalsOnCallResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
