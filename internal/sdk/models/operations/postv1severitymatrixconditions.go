// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type PostV1SeverityMatrixConditionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Create a new condition
	SeverityMatrixConditionEntity *shared.SeverityMatrixConditionEntity
}

func (o *PostV1SeverityMatrixConditionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostV1SeverityMatrixConditionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostV1SeverityMatrixConditionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostV1SeverityMatrixConditionsResponse) GetSeverityMatrixConditionEntity() *shared.SeverityMatrixConditionEntity {
	if o == nil {
		return nil
	}
	return o.SeverityMatrixConditionEntity
}
