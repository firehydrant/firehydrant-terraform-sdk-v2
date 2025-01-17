// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type PostV1TeamsTeamIDSignalRulesRequest struct {
	TeamID                       string                              `pathParam:"style=simple,explode=false,name=team_id"`
	PostV1TeamsTeamIDSignalRules shared.PostV1TeamsTeamIDSignalRules `request:"mediaType=application/json"`
}

func (o *PostV1TeamsTeamIDSignalRulesRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *PostV1TeamsTeamIDSignalRulesRequest) GetPostV1TeamsTeamIDSignalRules() shared.PostV1TeamsTeamIDSignalRules {
	if o == nil {
		return shared.PostV1TeamsTeamIDSignalRules{}
	}
	return o.PostV1TeamsTeamIDSignalRules
}

type PostV1TeamsTeamIDSignalRulesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostV1TeamsTeamIDSignalRulesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostV1TeamsTeamIDSignalRulesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostV1TeamsTeamIDSignalRulesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
