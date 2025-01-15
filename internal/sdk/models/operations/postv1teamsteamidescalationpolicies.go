// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type PostV1TeamsTeamIDEscalationPoliciesRequest struct {
	TeamID                              string                                     `pathParam:"style=simple,explode=false,name=team_id"`
	PostV1TeamsTeamIDEscalationPolicies shared.PostV1TeamsTeamIDEscalationPolicies `request:"mediaType=application/json"`
}

func (o *PostV1TeamsTeamIDEscalationPoliciesRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *PostV1TeamsTeamIDEscalationPoliciesRequest) GetPostV1TeamsTeamIDEscalationPolicies() shared.PostV1TeamsTeamIDEscalationPolicies {
	if o == nil {
		return shared.PostV1TeamsTeamIDEscalationPolicies{}
	}
	return o.PostV1TeamsTeamIDEscalationPolicies
}

type PostV1TeamsTeamIDEscalationPoliciesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostV1TeamsTeamIDEscalationPoliciesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostV1TeamsTeamIDEscalationPoliciesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostV1TeamsTeamIDEscalationPoliciesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
