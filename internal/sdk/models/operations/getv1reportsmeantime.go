// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/types"
	"net/http"
)

type GetV1ReportsMeanTimeRequest struct {
	// A comma separated list of environment IDs
	Environments *string `queryParam:"style=form,explode=true,name=environments"`
	// A comma separated list of team IDs
	Teams *string `queryParam:"style=form,explode=true,name=teams"`
	// A comma separated list of service IDs
	Services *string `queryParam:"style=form,explode=true,name=services"`
	// Incident status
	Status *string `queryParam:"style=form,explode=true,name=status"`
	// The start date to return incidents from
	StartDate *types.Date `queryParam:"style=form,explode=true,name=start_date"`
	// The end date to return incidents from
	EndDate *types.Date `queryParam:"style=form,explode=true,name=end_date"`
	// A text query for an incident that searches on name, summary, and desciption
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// The id of a previously saved search.
	SavedSearchID *string `queryParam:"style=form,explode=true,name=saved_search_id"`
	// A comma separated list of priorities
	Priorities *string `queryParam:"style=form,explode=true,name=priorities"`
	// Flag for including incidents where priority has not been set
	PriorityNotSet *bool `queryParam:"style=form,explode=true,name=priority_not_set"`
	// A comma separated list of severities
	Severities *string `queryParam:"style=form,explode=true,name=severities"`
	// Flag for including incidents where severity has not been set
	SeverityNotSet *bool `queryParam:"style=form,explode=true,name=severity_not_set"`
	// A comma separated list of current milestones
	CurrentMilestones *string `queryParam:"style=form,explode=true,name=current_milestones"`
}

func (g GetV1ReportsMeanTimeRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetV1ReportsMeanTimeRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetV1ReportsMeanTimeRequest) GetEnvironments() *string {
	if o == nil {
		return nil
	}
	return o.Environments
}

func (o *GetV1ReportsMeanTimeRequest) GetTeams() *string {
	if o == nil {
		return nil
	}
	return o.Teams
}

func (o *GetV1ReportsMeanTimeRequest) GetServices() *string {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *GetV1ReportsMeanTimeRequest) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *GetV1ReportsMeanTimeRequest) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *GetV1ReportsMeanTimeRequest) GetEndDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *GetV1ReportsMeanTimeRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *GetV1ReportsMeanTimeRequest) GetSavedSearchID() *string {
	if o == nil {
		return nil
	}
	return o.SavedSearchID
}

func (o *GetV1ReportsMeanTimeRequest) GetPriorities() *string {
	if o == nil {
		return nil
	}
	return o.Priorities
}

func (o *GetV1ReportsMeanTimeRequest) GetPriorityNotSet() *bool {
	if o == nil {
		return nil
	}
	return o.PriorityNotSet
}

func (o *GetV1ReportsMeanTimeRequest) GetSeverities() *string {
	if o == nil {
		return nil
	}
	return o.Severities
}

func (o *GetV1ReportsMeanTimeRequest) GetSeverityNotSet() *bool {
	if o == nil {
		return nil
	}
	return o.SeverityNotSet
}

func (o *GetV1ReportsMeanTimeRequest) GetCurrentMilestones() *string {
	if o == nil {
		return nil
	}
	return o.CurrentMilestones
}

type GetV1ReportsMeanTimeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns a report with time bucketed analytics data
	ReportEntity *shared.ReportEntity
}

func (o *GetV1ReportsMeanTimeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1ReportsMeanTimeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1ReportsMeanTimeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1ReportsMeanTimeResponse) GetReportEntity() *shared.ReportEntity {
	if o == nil {
		return nil
	}
	return o.ReportEntity
}
