// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type PutV1RunbooksExecutionsExecutionIDStepsStepIDRequest struct {
	ExecutionID                                   string                                               `pathParam:"style=simple,explode=false,name=execution_id"`
	StepID                                        string                                               `pathParam:"style=simple,explode=false,name=step_id"`
	PutV1RunbooksExecutionsExecutionIDStepsStepID shared.PutV1RunbooksExecutionsExecutionIDStepsStepID `request:"mediaType=application/json"`
}

func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDRequest) GetExecutionID() string {
	if o == nil {
		return ""
	}
	return o.ExecutionID
}

func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDRequest) GetStepID() string {
	if o == nil {
		return ""
	}
	return o.StepID
}

func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDRequest) GetPutV1RunbooksExecutionsExecutionIDStepsStepID() shared.PutV1RunbooksExecutionsExecutionIDStepsStepID {
	if o == nil {
		return shared.PutV1RunbooksExecutionsExecutionIDStepsStepID{}
	}
	return o.PutV1RunbooksExecutionsExecutionIDStepsStepID
}

type PutV1RunbooksExecutionsExecutionIDStepsStepIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Updates a runbook step execution, especially for changing the state of a step execution.
	RunbooksExecutionEntity *shared.RunbooksExecutionEntity
}

func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDResponse) GetRunbooksExecutionEntity() *shared.RunbooksExecutionEntity {
	if o == nil {
		return nil
	}
	return o.RunbooksExecutionEntity
}