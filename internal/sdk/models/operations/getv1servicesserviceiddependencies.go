// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type GetV1ServicesServiceIDDependenciesRequest struct {
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// If true, returns all dependencies in one array. If false, splits dependencies into different arrays for child and parent dependencies
	Flatten *bool `queryParam:"style=form,explode=true,name=flatten"`
}

func (o *GetV1ServicesServiceIDDependenciesRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetV1ServicesServiceIDDependenciesRequest) GetFlatten() *bool {
	if o == nil {
		return nil
	}
	return o.Flatten
}

type GetV1ServicesServiceIDDependenciesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Retrieves a service's dependencies
	ServiceWithAllDependenciesEntity *shared.ServiceWithAllDependenciesEntity
}

func (o *GetV1ServicesServiceIDDependenciesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1ServicesServiceIDDependenciesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1ServicesServiceIDDependenciesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1ServicesServiceIDDependenciesResponse) GetServiceWithAllDependenciesEntity() *shared.ServiceWithAllDependenciesEntity {
	if o == nil {
		return nil
	}
	return o.ServiceWithAllDependenciesEntity
}