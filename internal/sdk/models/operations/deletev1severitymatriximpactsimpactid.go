// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type DeleteV1SeverityMatrixImpactsImpactIDRequest struct {
	ImpactID string `pathParam:"style=simple,explode=false,name=impact_id"`
}

func (o *DeleteV1SeverityMatrixImpactsImpactIDRequest) GetImpactID() string {
	if o == nil {
		return ""
	}
	return o.ImpactID
}

type DeleteV1SeverityMatrixImpactsImpactIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Delete a specific impact
	SeverityMatrixImpactEntity *shared.SeverityMatrixImpactEntity
}

func (o *DeleteV1SeverityMatrixImpactsImpactIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteV1SeverityMatrixImpactsImpactIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteV1SeverityMatrixImpactsImpactIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteV1SeverityMatrixImpactsImpactIDResponse) GetSeverityMatrixImpactEntity() *shared.SeverityMatrixImpactEntity {
	if o == nil {
		return nil
	}
	return o.SeverityMatrixImpactEntity
}
