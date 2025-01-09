// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type DeleteV1IncidentRolesIncidentRoleIDRequest struct {
	IncidentRoleID string `pathParam:"style=simple,explode=false,name=incident_role_id"`
}

func (o *DeleteV1IncidentRolesIncidentRoleIDRequest) GetIncidentRoleID() string {
	if o == nil {
		return ""
	}
	return o.IncidentRoleID
}

type DeleteV1IncidentRolesIncidentRoleIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Archives an incident role which will hide it from lists and metrics
	IncidentRoleEntity *shared.IncidentRoleEntity
}

func (o *DeleteV1IncidentRolesIncidentRoleIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteV1IncidentRolesIncidentRoleIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteV1IncidentRolesIncidentRoleIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteV1IncidentRolesIncidentRoleIDResponse) GetIncidentRoleEntity() *shared.IncidentRoleEntity {
	if o == nil {
		return nil
	}
	return o.IncidentRoleEntity
}