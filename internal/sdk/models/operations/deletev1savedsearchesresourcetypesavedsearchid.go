// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResourceType string

const (
	ResourceTypeChangeEvents          ResourceType = "change_events"
	ResourceTypeIncidents             ResourceType = "incidents"
	ResourceTypeServices              ResourceType = "services"
	ResourceTypeScheduledMaintenances ResourceType = "scheduled_maintenances"
	ResourceTypeTicketTasks           ResourceType = "ticket_tasks"
	ResourceTypeTicketFollowUps       ResourceType = "ticket_follow_ups"
	ResourceTypeAnalytics             ResourceType = "analytics"
	ResourceTypeImpactAnalytics       ResourceType = "impact_analytics"
	ResourceTypeAlerts                ResourceType = "alerts"
	ResourceTypeIncidentEvents        ResourceType = "incident_events"
)

func (e ResourceType) ToPointer() *ResourceType {
	return &e
}
func (e *ResourceType) UnmarshalJSON(data []byte) error {
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
		*e = ResourceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ResourceType: %v", v)
	}
}

type DeleteV1SavedSearchesResourceTypeSavedSearchIDRequest struct {
	ResourceType  ResourceType `pathParam:"style=simple,explode=false,name=resource_type"`
	SavedSearchID string       `pathParam:"style=simple,explode=false,name=saved_search_id"`
}

func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDRequest) GetResourceType() ResourceType {
	if o == nil {
		return ResourceType("")
	}
	return o.ResourceType
}

func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDRequest) GetSavedSearchID() string {
	if o == nil {
		return ""
	}
	return o.SavedSearchID
}

type DeleteV1SavedSearchesResourceTypeSavedSearchIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}