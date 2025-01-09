// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type PatchV1IntegrationsAwsCloudtrailBatchesIDRequest struct {
	ID                                        string                                           `pathParam:"style=simple,explode=false,name=id"`
	PatchV1IntegrationsAwsCloudtrailBatchesID shared.PatchV1IntegrationsAwsCloudtrailBatchesID `request:"mediaType=application/json"`
}

func (o *PatchV1IntegrationsAwsCloudtrailBatchesIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PatchV1IntegrationsAwsCloudtrailBatchesIDRequest) GetPatchV1IntegrationsAwsCloudtrailBatchesID() shared.PatchV1IntegrationsAwsCloudtrailBatchesID {
	if o == nil {
		return shared.PatchV1IntegrationsAwsCloudtrailBatchesID{}
	}
	return o.PatchV1IntegrationsAwsCloudtrailBatchesID
}

type PatchV1IntegrationsAwsCloudtrailBatchesIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Update a CloudTrail batch with new information.
	IntegrationsAwsCloudtrailBatchEntity *shared.IntegrationsAwsCloudtrailBatchEntity
}

func (o *PatchV1IntegrationsAwsCloudtrailBatchesIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchV1IntegrationsAwsCloudtrailBatchesIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchV1IntegrationsAwsCloudtrailBatchesIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchV1IntegrationsAwsCloudtrailBatchesIDResponse) GetIntegrationsAwsCloudtrailBatchEntity() *shared.IntegrationsAwsCloudtrailBatchEntity {
	if o == nil {
		return nil
	}
	return o.IntegrationsAwsCloudtrailBatchEntity
}
