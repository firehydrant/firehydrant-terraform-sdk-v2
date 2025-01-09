// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/internal/utils"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/types"
	"net/http"
)

type InfraType string

const (
	InfraTypeEnvironments    InfraType = "environments"
	InfraTypeFunctionalities InfraType = "functionalities"
	InfraTypeServices        InfraType = "services"
	InfraTypeCustomers       InfraType = "customers"
)

func (e InfraType) ToPointer() *InfraType {
	return &e
}
func (e *InfraType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "environments":
		fallthrough
	case "functionalities":
		fallthrough
	case "services":
		fallthrough
	case "customers":
		*e = InfraType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InfraType: %v", v)
	}
}

type GetV1MetricsInfraTypeRequest struct {
	InfraType InfraType `pathParam:"style=simple,explode=false,name=infra_type"`
	// The start date to return metrics from; defaults to 30 days ago
	StartDate *types.Date `queryParam:"style=form,explode=true,name=start_date"`
	// The end date to return metrics from, defaults to today
	EndDate *types.Date `queryParam:"style=form,explode=true,name=end_date"`
}

func (g GetV1MetricsInfraTypeRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetV1MetricsInfraTypeRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetV1MetricsInfraTypeRequest) GetInfraType() InfraType {
	if o == nil {
		return InfraType("")
	}
	return o.InfraType
}

func (o *GetV1MetricsInfraTypeRequest) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *GetV1MetricsInfraTypeRequest) GetEndDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.EndDate
}

type GetV1MetricsInfraTypeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns metrics for all components of a given type
	MetricsInfrastructureListEntity *shared.MetricsInfrastructureListEntity
}

func (o *GetV1MetricsInfraTypeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1MetricsInfraTypeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1MetricsInfraTypeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1MetricsInfraTypeResponse) GetMetricsInfrastructureListEntity() *shared.MetricsInfrastructureListEntity {
	if o == nil {
		return nil
	}
	return o.MetricsInfrastructureListEntity
}
