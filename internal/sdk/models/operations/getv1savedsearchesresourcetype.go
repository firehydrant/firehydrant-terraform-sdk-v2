// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type GetV1SavedSearchesResourceTypePathParamResourceType string

const (
	GetV1SavedSearchesResourceTypePathParamResourceTypeChangeEvents          GetV1SavedSearchesResourceTypePathParamResourceType = "change_events"
	GetV1SavedSearchesResourceTypePathParamResourceTypeIncidents             GetV1SavedSearchesResourceTypePathParamResourceType = "incidents"
	GetV1SavedSearchesResourceTypePathParamResourceTypeServices              GetV1SavedSearchesResourceTypePathParamResourceType = "services"
	GetV1SavedSearchesResourceTypePathParamResourceTypeScheduledMaintenances GetV1SavedSearchesResourceTypePathParamResourceType = "scheduled_maintenances"
	GetV1SavedSearchesResourceTypePathParamResourceTypeTicketTasks           GetV1SavedSearchesResourceTypePathParamResourceType = "ticket_tasks"
	GetV1SavedSearchesResourceTypePathParamResourceTypeTicketFollowUps       GetV1SavedSearchesResourceTypePathParamResourceType = "ticket_follow_ups"
	GetV1SavedSearchesResourceTypePathParamResourceTypeAnalytics             GetV1SavedSearchesResourceTypePathParamResourceType = "analytics"
	GetV1SavedSearchesResourceTypePathParamResourceTypeImpactAnalytics       GetV1SavedSearchesResourceTypePathParamResourceType = "impact_analytics"
	GetV1SavedSearchesResourceTypePathParamResourceTypeAlerts                GetV1SavedSearchesResourceTypePathParamResourceType = "alerts"
	GetV1SavedSearchesResourceTypePathParamResourceTypeIncidentEvents        GetV1SavedSearchesResourceTypePathParamResourceType = "incident_events"
)

func (e GetV1SavedSearchesResourceTypePathParamResourceType) ToPointer() *GetV1SavedSearchesResourceTypePathParamResourceType {
	return &e
}
func (e *GetV1SavedSearchesResourceTypePathParamResourceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "change_events":
		fallthrough
	case "incidents":
		fallthrough
	case "services":
		fallthrough
	case "scheduled_maintenances":
		fallthrough
	case "ticket_tasks":
		fallthrough
	case "ticket_follow_ups":
		fallthrough
	case "analytics":
		fallthrough
	case "impact_analytics":
		fallthrough
	case "alerts":
		fallthrough
	case "incident_events":
		*e = GetV1SavedSearchesResourceTypePathParamResourceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetV1SavedSearchesResourceTypePathParamResourceType: %v", v)
	}
}

type GetV1SavedSearchesResourceTypeRequest struct {
	ResourceType GetV1SavedSearchesResourceTypePathParamResourceType `pathParam:"style=simple,explode=false,name=resource_type"`
	// The user ID used to filter saved searches.
	UserID *string `queryParam:"style=form,explode=true,name=user_id"`
	// Filter saved searches with a query on their name
	Query   *string `queryParam:"style=form,explode=true,name=query"`
	Page    *int    `queryParam:"style=form,explode=true,name=page"`
	PerPage *int    `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *GetV1SavedSearchesResourceTypeRequest) GetResourceType() GetV1SavedSearchesResourceTypePathParamResourceType {
	if o == nil {
		return GetV1SavedSearchesResourceTypePathParamResourceType("")
	}
	return o.ResourceType
}

func (o *GetV1SavedSearchesResourceTypeRequest) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *GetV1SavedSearchesResourceTypeRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *GetV1SavedSearchesResourceTypeRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1SavedSearchesResourceTypeRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type GetV1SavedSearchesResourceTypeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Lists save searches
	SavedSearchEntity *shared.SavedSearchEntity
}

func (o *GetV1SavedSearchesResourceTypeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1SavedSearchesResourceTypeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1SavedSearchesResourceTypeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1SavedSearchesResourceTypeResponse) GetSavedSearchEntity() *shared.SavedSearchEntity {
	if o == nil {
		return nil
	}
	return o.SavedSearchEntity
}
