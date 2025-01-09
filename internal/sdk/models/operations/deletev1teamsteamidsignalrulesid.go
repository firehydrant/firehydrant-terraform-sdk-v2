// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteV1TeamsTeamIDSignalRulesIDRequest struct {
	TeamID string `pathParam:"style=simple,explode=false,name=team_id"`
	ID     string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteV1TeamsTeamIDSignalRulesIDRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *DeleteV1TeamsTeamIDSignalRulesIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteV1TeamsTeamIDSignalRulesIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteV1TeamsTeamIDSignalRulesIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteV1TeamsTeamIDSignalRulesIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteV1TeamsTeamIDSignalRulesIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
