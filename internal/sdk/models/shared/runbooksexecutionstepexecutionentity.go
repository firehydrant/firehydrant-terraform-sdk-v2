// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/internal/utils"
	"time"
)

type RunbooksExecutionStepExecutionEntityState string

const (
	RunbooksExecutionStepExecutionEntityStateInitial   RunbooksExecutionStepExecutionEntityState = "initial"
	RunbooksExecutionStepExecutionEntityStatePending   RunbooksExecutionStepExecutionEntityState = "pending"
	RunbooksExecutionStepExecutionEntityStateScheduled RunbooksExecutionStepExecutionEntityState = "scheduled"
	RunbooksExecutionStepExecutionEntityStateStarted   RunbooksExecutionStepExecutionEntityState = "started"
	RunbooksExecutionStepExecutionEntityStateDismissed RunbooksExecutionStepExecutionEntityState = "dismissed"
	RunbooksExecutionStepExecutionEntityStateCompleted RunbooksExecutionStepExecutionEntityState = "completed"
	RunbooksExecutionStepExecutionEntityStateErrored   RunbooksExecutionStepExecutionEntityState = "errored"
)

func (e RunbooksExecutionStepExecutionEntityState) ToPointer() *RunbooksExecutionStepExecutionEntityState {
	return &e
}
func (e *RunbooksExecutionStepExecutionEntityState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "initial":
		fallthrough
	case "pending":
		fallthrough
	case "scheduled":
		fallthrough
	case "started":
		fallthrough
	case "dismissed":
		fallthrough
	case "completed":
		fallthrough
	case "errored":
		*e = RunbooksExecutionStepExecutionEntityState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RunbooksExecutionStepExecutionEntityState: %v", v)
	}
}

type RunbooksExecutionStepExecutionEntityData struct {
}

type RunbooksExecutionStepExecutionEntity struct {
	State           *RunbooksExecutionStepExecutionEntityState `json:"state,omitempty"`
	Data            *RunbooksExecutionStepExecutionEntityData  `json:"data,omitempty"`
	PerformedBy     *ActorEntity                               `json:"performed_by,omitempty"`
	PerformedAt     *time.Time                                 `json:"performed_at,omitempty"`
	ScheduledFor    *time.Time                                 `json:"scheduled_for,omitempty"`
	Error           *string                                    `json:"error,omitempty"`
	WebhookDelivery *RunbooksWebhookDeliveryEntity             `json:"webhook_delivery,omitempty"`
}

func (r RunbooksExecutionStepExecutionEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RunbooksExecutionStepExecutionEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RunbooksExecutionStepExecutionEntity) GetState() *RunbooksExecutionStepExecutionEntityState {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *RunbooksExecutionStepExecutionEntity) GetData() *RunbooksExecutionStepExecutionEntityData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *RunbooksExecutionStepExecutionEntity) GetPerformedBy() *ActorEntity {
	if o == nil {
		return nil
	}
	return o.PerformedBy
}

func (o *RunbooksExecutionStepExecutionEntity) GetPerformedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.PerformedAt
}

func (o *RunbooksExecutionStepExecutionEntity) GetScheduledFor() *time.Time {
	if o == nil {
		return nil
	}
	return o.ScheduledFor
}

func (o *RunbooksExecutionStepExecutionEntity) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *RunbooksExecutionStepExecutionEntity) GetWebhookDelivery() *RunbooksWebhookDeliveryEntity {
	if o == nil {
		return nil
	}
	return o.WebhookDelivery
}
