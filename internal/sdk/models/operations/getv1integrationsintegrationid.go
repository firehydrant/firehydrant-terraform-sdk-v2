// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type GetV1IntegrationsIntegrationIDRequest struct {
	// Integration UUID
	IntegrationID string `pathParam:"style=simple,explode=false,name=integration_id"`
}

func (o *GetV1IntegrationsIntegrationIDRequest) GetIntegrationID() string {
	if o == nil {
		return ""
	}
	return o.IntegrationID
}

type GetV1IntegrationsIntegrationIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Retrieve a single integration
	IntegrationsIntegrationEntity *shared.IntegrationsIntegrationEntity
}

func (o *GetV1IntegrationsIntegrationIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1IntegrationsIntegrationIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1IntegrationsIntegrationIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1IntegrationsIntegrationIDResponse) GetIntegrationsIntegrationEntity() *shared.IntegrationsIntegrationEntity {
	if o == nil {
		return nil
	}
	return o.IntegrationsIntegrationEntity
}
