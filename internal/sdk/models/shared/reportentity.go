// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/types"
)

// ReportEntity model
type ReportEntity struct {
	Data         []ReportsBucketEntity `json:"data,omitempty"`
	StartDate    *types.Date           `json:"start_date,omitempty"`
	EndDate      *types.Date           `json:"end_date,omitempty"`
	BucketPeriod *string               `json:"bucket_period,omitempty"`
}

func (r ReportEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *ReportEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ReportEntity) GetData() []ReportsBucketEntity {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *ReportEntity) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *ReportEntity) GetEndDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *ReportEntity) GetBucketPeriod() *string {
	if o == nil {
		return nil
	}
	return o.BucketPeriod
}
