// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type GetV1RunbooksExecutionsExecutionIDVotesStatusRequest struct {
	ExecutionID string `pathParam:"style=simple,explode=false,name=execution_id"`
}

func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusRequest) GetExecutionID() string {
	if o == nil {
		return ""
	}
	return o.ExecutionID
}

type GetV1RunbooksExecutionsExecutionIDVotesStatusResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns the current vote counts for an object
	VotesEntity *shared.VotesEntity
}

func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusResponse) GetVotesEntity() *shared.VotesEntity {
	if o == nil {
		return nil
	}
	return o.VotesEntity
}