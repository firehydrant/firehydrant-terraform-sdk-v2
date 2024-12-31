// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/internal/utils"
	"time"
)

type MetricsTicketFunnelMetricsEntityDataBucketEntity struct {
	// The start datetime for the period
	TimeBucket   *time.Time                                                    `json:"time_bucket,omitempty"`
	FilterParams *MetricsTicketFunnelMetricsEntityDataBucketFilterParamsEntity `json:"filter_params,omitempty"`
	// The number of tasks created
	TasksCreated *int `json:"tasks_created,omitempty"`
	// The number of tasks completed
	TasksDone *int `json:"tasks_done,omitempty"`
	// The number of follow ups created
	FollowUpsCreated *int `json:"follow_ups_created,omitempty"`
	// The number of follow ups completed
	FollowUpsDone *int `json:"follow_ups_done,omitempty"`
}

func (m MetricsTicketFunnelMetricsEntityDataBucketEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MetricsTicketFunnelMetricsEntityDataBucketEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MetricsTicketFunnelMetricsEntityDataBucketEntity) GetTimeBucket() *time.Time {
	if o == nil {
		return nil
	}
	return o.TimeBucket
}

func (o *MetricsTicketFunnelMetricsEntityDataBucketEntity) GetFilterParams() *MetricsTicketFunnelMetricsEntityDataBucketFilterParamsEntity {
	if o == nil {
		return nil
	}
	return o.FilterParams
}

func (o *MetricsTicketFunnelMetricsEntityDataBucketEntity) GetTasksCreated() *int {
	if o == nil {
		return nil
	}
	return o.TasksCreated
}

func (o *MetricsTicketFunnelMetricsEntityDataBucketEntity) GetTasksDone() *int {
	if o == nil {
		return nil
	}
	return o.TasksDone
}

func (o *MetricsTicketFunnelMetricsEntityDataBucketEntity) GetFollowUpsCreated() *int {
	if o == nil {
		return nil
	}
	return o.FollowUpsCreated
}

func (o *MetricsTicketFunnelMetricsEntityDataBucketEntity) GetFollowUpsDone() *int {
	if o == nil {
		return nil
	}
	return o.FollowUpsDone
}
