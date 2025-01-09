// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type PatchV1TeamsTeamIDSignalRulesIDRequest struct {
	TeamID                          string                                 `pathParam:"style=simple,explode=false,name=team_id"`
	ID                              string                                 `pathParam:"style=simple,explode=false,name=id"`
	PatchV1TeamsTeamIDSignalRulesID shared.PatchV1TeamsTeamIDSignalRulesID `request:"mediaType=application/json"`
}

func (o *PatchV1TeamsTeamIDSignalRulesIDRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *PatchV1TeamsTeamIDSignalRulesIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PatchV1TeamsTeamIDSignalRulesIDRequest) GetPatchV1TeamsTeamIDSignalRulesID() shared.PatchV1TeamsTeamIDSignalRulesID {
	if o == nil {
		return shared.PatchV1TeamsTeamIDSignalRulesID{}
	}
	return o.PatchV1TeamsTeamIDSignalRulesID
}

type PatchV1TeamsTeamIDSignalRulesIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchV1TeamsTeamIDSignalRulesIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchV1TeamsTeamIDSignalRulesIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchV1TeamsTeamIDSignalRulesIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}