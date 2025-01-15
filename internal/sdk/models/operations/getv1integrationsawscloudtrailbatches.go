// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type GetV1IntegrationsAwsCloudtrailBatchesRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// AWS connection ID
	ConnectionID *string `queryParam:"style=form,explode=true,name=connection_id"`
}

func (o *GetV1IntegrationsAwsCloudtrailBatchesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1IntegrationsAwsCloudtrailBatchesRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *GetV1IntegrationsAwsCloudtrailBatchesRequest) GetConnectionID() *string {
	if o == nil {
		return nil
	}
	return o.ConnectionID
}

type GetV1IntegrationsAwsCloudtrailBatchesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Lists CloudTrail batches for the authenticated organization.
	IntegrationsAwsCloudtrailBatchEntityPaginated *shared.IntegrationsAwsCloudtrailBatchEntityPaginated
}

func (o *GetV1IntegrationsAwsCloudtrailBatchesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1IntegrationsAwsCloudtrailBatchesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1IntegrationsAwsCloudtrailBatchesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1IntegrationsAwsCloudtrailBatchesResponse) GetIntegrationsAwsCloudtrailBatchEntityPaginated() *shared.IntegrationsAwsCloudtrailBatchEntityPaginated {
	if o == nil {
		return nil
	}
	return o.IntegrationsAwsCloudtrailBatchEntityPaginated
}
