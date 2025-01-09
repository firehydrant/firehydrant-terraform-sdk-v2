// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type DeleteV1SeverityMatrixConditionsConditionIDRequest struct {
	ConditionID string `pathParam:"style=simple,explode=false,name=condition_id"`
}

func (o *DeleteV1SeverityMatrixConditionsConditionIDRequest) GetConditionID() string {
	if o == nil {
		return ""
	}
	return o.ConditionID
}

type DeleteV1SeverityMatrixConditionsConditionIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Delete a specific condition
	SeverityMatrixConditionEntity *shared.SeverityMatrixConditionEntity
}

func (o *DeleteV1SeverityMatrixConditionsConditionIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteV1SeverityMatrixConditionsConditionIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteV1SeverityMatrixConditionsConditionIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteV1SeverityMatrixConditionsConditionIDResponse) GetSeverityMatrixConditionEntity() *shared.SeverityMatrixConditionEntity {
	if o == nil {
		return nil
	}
	return o.SeverityMatrixConditionEntity
}
