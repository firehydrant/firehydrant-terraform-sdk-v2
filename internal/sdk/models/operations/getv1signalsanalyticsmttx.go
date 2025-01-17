// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"net/http"
	"time"
)

// GetV1SignalsAnalyticsMttxQueryParamGroupBy - String that determines how records are grouped
type GetV1SignalsAnalyticsMttxQueryParamGroupBy string

const (
	GetV1SignalsAnalyticsMttxQueryParamGroupBySignalRules  GetV1SignalsAnalyticsMttxQueryParamGroupBy = "signal_rules"
	GetV1SignalsAnalyticsMttxQueryParamGroupByTeams        GetV1SignalsAnalyticsMttxQueryParamGroupBy = "teams"
	GetV1SignalsAnalyticsMttxQueryParamGroupByServices     GetV1SignalsAnalyticsMttxQueryParamGroupBy = "services"
	GetV1SignalsAnalyticsMttxQueryParamGroupByEnvironments GetV1SignalsAnalyticsMttxQueryParamGroupBy = "environments"
	GetV1SignalsAnalyticsMttxQueryParamGroupByTags         GetV1SignalsAnalyticsMttxQueryParamGroupBy = "tags"
)

func (e GetV1SignalsAnalyticsMttxQueryParamGroupBy) ToPointer() *GetV1SignalsAnalyticsMttxQueryParamGroupBy {
	return &e
}
func (e *GetV1SignalsAnalyticsMttxQueryParamGroupBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "signal_rules":
		fallthrough
	case "teams":
		fallthrough
	case "services":
		fallthrough
	case "environments":
		fallthrough
	case "tags":
		*e = GetV1SignalsAnalyticsMttxQueryParamGroupBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetV1SignalsAnalyticsMttxQueryParamGroupBy: %v", v)
	}
}

// GetV1SignalsAnalyticsMttxQueryParamSortBy - String that determines how records are sorted
type GetV1SignalsAnalyticsMttxQueryParamSortBy string

const (
	GetV1SignalsAnalyticsMttxQueryParamSortByTotalOpenedAlerts   GetV1SignalsAnalyticsMttxQueryParamSortBy = "total_opened_alerts"
	GetV1SignalsAnalyticsMttxQueryParamSortByTotalAckedAlerts    GetV1SignalsAnalyticsMttxQueryParamSortBy = "total_acked_alerts"
	GetV1SignalsAnalyticsMttxQueryParamSortByTotalIncidents      GetV1SignalsAnalyticsMttxQueryParamSortBy = "total_incidents"
	GetV1SignalsAnalyticsMttxQueryParamSortByAckedPercentage     GetV1SignalsAnalyticsMttxQueryParamSortBy = "acked_percentage"
	GetV1SignalsAnalyticsMttxQueryParamSortByIncidentsPercentage GetV1SignalsAnalyticsMttxQueryParamSortBy = "incidents_percentage"
)

func (e GetV1SignalsAnalyticsMttxQueryParamSortBy) ToPointer() *GetV1SignalsAnalyticsMttxQueryParamSortBy {
	return &e
}
func (e *GetV1SignalsAnalyticsMttxQueryParamSortBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "total_opened_alerts":
		fallthrough
	case "total_acked_alerts":
		fallthrough
	case "total_incidents":
		fallthrough
	case "acked_percentage":
		fallthrough
	case "incidents_percentage":
		*e = GetV1SignalsAnalyticsMttxQueryParamSortBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetV1SignalsAnalyticsMttxQueryParamSortBy: %v", v)
	}
}

// GetV1SignalsAnalyticsMttxQueryParamSortDirection - String that determines how records are sorted
type GetV1SignalsAnalyticsMttxQueryParamSortDirection string

const (
	GetV1SignalsAnalyticsMttxQueryParamSortDirectionAsc  GetV1SignalsAnalyticsMttxQueryParamSortDirection = "asc"
	GetV1SignalsAnalyticsMttxQueryParamSortDirectionDesc GetV1SignalsAnalyticsMttxQueryParamSortDirection = "desc"
)

func (e GetV1SignalsAnalyticsMttxQueryParamSortDirection) ToPointer() *GetV1SignalsAnalyticsMttxQueryParamSortDirection {
	return &e
}
func (e *GetV1SignalsAnalyticsMttxQueryParamSortDirection) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetV1SignalsAnalyticsMttxQueryParamSortDirection(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetV1SignalsAnalyticsMttxQueryParamSortDirection: %v", v)
	}
}

type GetV1SignalsAnalyticsMttxRequest struct {
	// A comma separated list of signal rule IDs
	SignalRules *string `queryParam:"style=form,explode=true,name=signal_rules"`
	// A comma separated list of team IDs
	Teams *string `queryParam:"style=form,explode=true,name=teams"`
	// A comma separated list of environment IDs
	Environments *string `queryParam:"style=form,explode=true,name=environments"`
	// A comma separated list of service IDs
	Services *string `queryParam:"style=form,explode=true,name=services"`
	// A comma separated list of tags
	Tags *string `queryParam:"style=form,explode=true,name=tags"`
	// A comma separated list of user IDs
	Users *string `queryParam:"style=form,explode=true,name=users"`
	// String that determines how records are grouped
	GroupBy *GetV1SignalsAnalyticsMttxQueryParamGroupBy `queryParam:"style=form,explode=true,name=group_by"`
	// String that determines how records are sorted
	SortBy *GetV1SignalsAnalyticsMttxQueryParamSortBy `queryParam:"style=form,explode=true,name=sort_by"`
	// String that determines how records are sorted
	SortDirection *GetV1SignalsAnalyticsMttxQueryParamSortDirection `queryParam:"style=form,explode=true,name=sort_direction"`
	// The start date to return metrics from
	StartDate *time.Time `queryParam:"style=form,explode=true,name=start_date"`
	// The end date to return metrics from
	EndDate *time.Time `queryParam:"style=form,explode=true,name=end_date"`
}

func (g GetV1SignalsAnalyticsMttxRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetV1SignalsAnalyticsMttxRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetV1SignalsAnalyticsMttxRequest) GetSignalRules() *string {
	if o == nil {
		return nil
	}
	return o.SignalRules
}

func (o *GetV1SignalsAnalyticsMttxRequest) GetTeams() *string {
	if o == nil {
		return nil
	}
	return o.Teams
}

func (o *GetV1SignalsAnalyticsMttxRequest) GetEnvironments() *string {
	if o == nil {
		return nil
	}
	return o.Environments
}

func (o *GetV1SignalsAnalyticsMttxRequest) GetServices() *string {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *GetV1SignalsAnalyticsMttxRequest) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *GetV1SignalsAnalyticsMttxRequest) GetUsers() *string {
	if o == nil {
		return nil
	}
	return o.Users
}

func (o *GetV1SignalsAnalyticsMttxRequest) GetGroupBy() *GetV1SignalsAnalyticsMttxQueryParamGroupBy {
	if o == nil {
		return nil
	}
	return o.GroupBy
}

func (o *GetV1SignalsAnalyticsMttxRequest) GetSortBy() *GetV1SignalsAnalyticsMttxQueryParamSortBy {
	if o == nil {
		return nil
	}
	return o.SortBy
}

func (o *GetV1SignalsAnalyticsMttxRequest) GetSortDirection() *GetV1SignalsAnalyticsMttxQueryParamSortDirection {
	if o == nil {
		return nil
	}
	return o.SortDirection
}

func (o *GetV1SignalsAnalyticsMttxRequest) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *GetV1SignalsAnalyticsMttxRequest) GetEndDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndDate
}

type GetV1SignalsAnalyticsMttxResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetV1SignalsAnalyticsMttxResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1SignalsAnalyticsMttxResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1SignalsAnalyticsMttxResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
