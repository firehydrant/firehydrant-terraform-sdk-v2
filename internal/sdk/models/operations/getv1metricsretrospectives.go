// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/types"
	"net/http"
)

type GetV1MetricsRetrospectivesRequest struct {
	// The start date to return metrics from
	StartDate *types.Date `queryParam:"style=form,explode=true,name=start_date"`
	// The end date to return metrics from
	EndDate *types.Date `queryParam:"style=form,explode=true,name=end_date"`
}

func (g GetV1MetricsRetrospectivesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetV1MetricsRetrospectivesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetV1MetricsRetrospectivesRequest) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *GetV1MetricsRetrospectivesRequest) GetEndDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.EndDate
}

type GetV1MetricsRetrospectivesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns a report with retrospective analytics data
	MetricsRetrospectiveEntity *shared.MetricsRetrospectiveEntity
}

func (o *GetV1MetricsRetrospectivesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1MetricsRetrospectivesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1MetricsRetrospectivesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1MetricsRetrospectivesResponse) GetMetricsRetrospectiveEntity() *shared.MetricsRetrospectiveEntity {
	if o == nil {
		return nil
	}
	return o.MetricsRetrospectiveEntity
}
