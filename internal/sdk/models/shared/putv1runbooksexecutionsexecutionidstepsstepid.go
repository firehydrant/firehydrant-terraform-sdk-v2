// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

// PutV1RunbooksExecutionsExecutionIDStepsStepID - Updates a runbook step execution, especially for changing the state of a step execution.
type PutV1RunbooksExecutionsExecutionIDStepsStepID struct {
	State       string     `json:"state"`
	ScheduleFor *time.Time `json:"schedule_for,omitempty"`
	// Data for execution of this step
	Data      map[string]any `json:"data,omitempty"`
	RepeatsAt *time.Time     `json:"repeats_at,omitempty"`
}

func (p PutV1RunbooksExecutionsExecutionIDStepsStepID) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PutV1RunbooksExecutionsExecutionIDStepsStepID) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PutV1RunbooksExecutionsExecutionIDStepsStepID) GetState() string {
	if o == nil {
		return ""
	}
	return o.State
}

func (o *PutV1RunbooksExecutionsExecutionIDStepsStepID) GetScheduleFor() *time.Time {
	if o == nil {
		return nil
	}
	return o.ScheduleFor
}

func (o *PutV1RunbooksExecutionsExecutionIDStepsStepID) GetData() map[string]any {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *PutV1RunbooksExecutionsExecutionIDStepsStepID) GetRepeatsAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RepeatsAt
}
