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

type BucketSize string

const (
	BucketSizeWeek    BucketSize = "week"
	BucketSizeDay     BucketSize = "day"
	BucketSizeMonth   BucketSize = "month"
	BucketSizeAllTime BucketSize = "all_time"
)

func (e BucketSize) ToPointer() *BucketSize {
	return &e
}
func (e *BucketSize) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "week":
		fallthrough
	case "day":
		fallthrough
	case "month":
		fallthrough
	case "all_time":
		*e = BucketSize(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BucketSize: %v", v)
	}
}

type By string

const (
	ByTotal           By = "total"
	BySeverity        By = "severity"
	ByPriority        By = "priority"
	ByFunctionality   By = "functionality"
	ByService         By = "service"
	ByEnvironment     By = "environment"
	ByUser            By = "user"
	ByUserInvolvement By = "user_involvement"
)

func (e By) ToPointer() *By {
	return &e
}
func (e *By) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "total":
		fallthrough
	case "severity":
		fallthrough
	case "priority":
		fallthrough
	case "functionality":
		fallthrough
	case "service":
		fallthrough
	case "environment":
		fallthrough
	case "user":
		fallthrough
	case "user_involvement":
		*e = By(v)
		return nil
	default:
		return fmt.Errorf("invalid value for By: %v", v)
	}
}

type SortField string

const (
	SortFieldMttd      SortField = "mttd"
	SortFieldMtta      SortField = "mtta"
	SortFieldMttm      SortField = "mttm"
	SortFieldMttr      SortField = "mttr"
	SortFieldCount     SortField = "count"
	SortFieldTotalTime SortField = "total_time"
)

func (e SortField) ToPointer() *SortField {
	return &e
}
func (e *SortField) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mttd":
		fallthrough
	case "mtta":
		fallthrough
	case "mttm":
		fallthrough
	case "mttr":
		fallthrough
	case "count":
		fallthrough
	case "total_time":
		*e = SortField(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SortField: %v", v)
	}
}

type SortDirection string

const (
	SortDirectionAsc  SortDirection = "asc"
	SortDirectionDesc SortDirection = "desc"
)

func (e SortDirection) ToPointer() *SortDirection {
	return &e
}
func (e *SortDirection) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = SortDirection(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SortDirection: %v", v)
	}
}

type ListIncidentMetricsRequest struct {
	// The start date to return metrics from
	StartDate *types.Date `queryParam:"style=form,explode=true,name=start_date"`
	// The end date to return metrics from
	EndDate       *types.Date    `queryParam:"style=form,explode=true,name=end_date"`
	BucketSize    *BucketSize    `queryParam:"style=form,explode=true,name=bucket_size"`
	By            *By            `queryParam:"style=form,explode=true,name=by"`
	SortField     *SortField     `queryParam:"style=form,explode=true,name=sort_field"`
	SortDirection *SortDirection `queryParam:"style=form,explode=true,name=sort_direction"`
	SortLimit     *int           `queryParam:"style=form,explode=true,name=sort_limit"`
	Conditions    *string        `queryParam:"style=form,explode=true,name=conditions"`
}

func (l ListIncidentMetricsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListIncidentMetricsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListIncidentMetricsRequest) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *ListIncidentMetricsRequest) GetEndDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *ListIncidentMetricsRequest) GetBucketSize() *BucketSize {
	if o == nil {
		return nil
	}
	return o.BucketSize
}

func (o *ListIncidentMetricsRequest) GetBy() *By {
	if o == nil {
		return nil
	}
	return o.By
}

func (o *ListIncidentMetricsRequest) GetSortField() *SortField {
	if o == nil {
		return nil
	}
	return o.SortField
}

func (o *ListIncidentMetricsRequest) GetSortDirection() *SortDirection {
	if o == nil {
		return nil
	}
	return o.SortDirection
}

func (o *ListIncidentMetricsRequest) GetSortLimit() *int {
	if o == nil {
		return nil
	}
	return o.SortLimit
}

func (o *ListIncidentMetricsRequest) GetConditions() *string {
	if o == nil {
		return nil
	}
	return o.Conditions
}

type ListIncidentMetricsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns a report with time bucketed analytics data
	MetricsMetricsEntity *shared.MetricsMetricsEntity
	// A collection of codes that generally means the end user got something wrong in making the request
	BadRequest *shared.BadRequest
	// A collection of codes that generally means the client was not authenticated correctly for the request they want to make
	Unauthorized *shared.Unauthorized
	// Status codes relating to the resource/entity they are requesting not being found or endpoints/routes not existing
	NotFound *shared.NotFound
	// Status codes relating to the client being rate limited by the server
	RateLimited *shared.RateLimited
	// A collection of status codes that generally mean the server failed in an unexpected way
	InternalServerError *shared.InternalServerError
	// Timeouts occurred with the request
	Timeout *shared.Timeout
}

func (o *ListIncidentMetricsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListIncidentMetricsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListIncidentMetricsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListIncidentMetricsResponse) GetMetricsMetricsEntity() *shared.MetricsMetricsEntity {
	if o == nil {
		return nil
	}
	return o.MetricsMetricsEntity
}

func (o *ListIncidentMetricsResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *ListIncidentMetricsResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *ListIncidentMetricsResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *ListIncidentMetricsResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *ListIncidentMetricsResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *ListIncidentMetricsResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}
