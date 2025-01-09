// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

// GetV1AlertsQueryParamTagMatchStrategy - The strategy to match tags. `any` will return alerts that have at least one of the supplied tags, `match_all` will return only alerts that have all of the supplied tags, and `exclude` will only return alerts that have none of the supplied tags. This currently only works for Signals alerts.
type GetV1AlertsQueryParamTagMatchStrategy string

const (
	GetV1AlertsQueryParamTagMatchStrategyAny      GetV1AlertsQueryParamTagMatchStrategy = "any"
	GetV1AlertsQueryParamTagMatchStrategyMatchAll GetV1AlertsQueryParamTagMatchStrategy = "match_all"
	GetV1AlertsQueryParamTagMatchStrategyExclude  GetV1AlertsQueryParamTagMatchStrategy = "exclude"
)

func (e GetV1AlertsQueryParamTagMatchStrategy) ToPointer() *GetV1AlertsQueryParamTagMatchStrategy {
	return &e
}
func (e *GetV1AlertsQueryParamTagMatchStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "any":
		fallthrough
	case "match_all":
		fallthrough
	case "exclude":
		*e = GetV1AlertsQueryParamTagMatchStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetV1AlertsQueryParamTagMatchStrategy: %v", v)
	}
}

type GetV1AlertsRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// A text query for alerts
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// A comma separated list of user IDs. This currently only works for Signals alerts.
	Users *string `queryParam:"style=form,explode=true,name=users"`
	// A comma separated list of team IDs. This currently only works for Signals alerts.
	Teams *string `queryParam:"style=form,explode=true,name=teams"`
	// A comma separated list of signals rule IDs. This currently only works for Signals alerts.
	SignalRules *string `queryParam:"style=form,explode=true,name=signal_rules"`
	// A comma separated list of environment IDs. This currently only works for Signals alerts.
	Environments *string `queryParam:"style=form,explode=true,name=environments"`
	// A comma separated list of functionality IDs. This currently only works for Signals alerts.
	Functionalities *string `queryParam:"style=form,explode=true,name=functionalities"`
	// A comma separated list of service IDs. This currently only works for Signals alerts.
	Services *string `queryParam:"style=form,explode=true,name=services"`
	// A comma separated list of tags. This currently only works for Signals alerts.
	Tags *string `queryParam:"style=form,explode=true,name=tags"`
	// The strategy to match tags. `any` will return alerts that have at least one of the supplied tags, `match_all` will return only alerts that have all of the supplied tags, and `exclude` will only return alerts that have none of the supplied tags. This currently only works for Signals alerts.
	TagMatchStrategy *GetV1AlertsQueryParamTagMatchStrategy `queryParam:"style=form,explode=true,name=tag_match_strategy"`
	// A comma separated list of statuses to filter by. Valid statuses are: opened, acknowledged, resolved, ignored, or expired
	Statuses *string `queryParam:"style=form,explode=true,name=statuses"`
}

func (o *GetV1AlertsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1AlertsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *GetV1AlertsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *GetV1AlertsRequest) GetUsers() *string {
	if o == nil {
		return nil
	}
	return o.Users
}

func (o *GetV1AlertsRequest) GetTeams() *string {
	if o == nil {
		return nil
	}
	return o.Teams
}

func (o *GetV1AlertsRequest) GetSignalRules() *string {
	if o == nil {
		return nil
	}
	return o.SignalRules
}

func (o *GetV1AlertsRequest) GetEnvironments() *string {
	if o == nil {
		return nil
	}
	return o.Environments
}

func (o *GetV1AlertsRequest) GetFunctionalities() *string {
	if o == nil {
		return nil
	}
	return o.Functionalities
}

func (o *GetV1AlertsRequest) GetServices() *string {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *GetV1AlertsRequest) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *GetV1AlertsRequest) GetTagMatchStrategy() *GetV1AlertsQueryParamTagMatchStrategy {
	if o == nil {
		return nil
	}
	return o.TagMatchStrategy
}

func (o *GetV1AlertsRequest) GetStatuses() *string {
	if o == nil {
		return nil
	}
	return o.Statuses
}

type GetV1AlertsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Retrieve all alerts from third parties
	AlertsAlertEntityPaginated *shared.AlertsAlertEntityPaginated
}

func (o *GetV1AlertsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1AlertsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1AlertsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1AlertsResponse) GetAlertsAlertEntityPaginated() *shared.AlertsAlertEntityPaginated {
	if o == nil {
		return nil
	}
	return o.AlertsAlertEntityPaginated
}