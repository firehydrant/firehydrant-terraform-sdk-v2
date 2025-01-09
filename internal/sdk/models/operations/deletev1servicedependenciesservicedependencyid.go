// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type DeleteV1ServiceDependenciesServiceDependencyIDRequest struct {
	ServiceDependencyID string `pathParam:"style=simple,explode=false,name=service_dependency_id"`
}

func (o *DeleteV1ServiceDependenciesServiceDependencyIDRequest) GetServiceDependencyID() string {
	if o == nil {
		return ""
	}
	return o.ServiceDependencyID
}

type DeleteV1ServiceDependenciesServiceDependencyIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Deletes a single service dependency
	ServiceDependencyEntity *shared.ServiceDependencyEntity
}

func (o *DeleteV1ServiceDependenciesServiceDependencyIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteV1ServiceDependenciesServiceDependencyIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteV1ServiceDependenciesServiceDependencyIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteV1ServiceDependenciesServiceDependencyIDResponse) GetServiceDependencyEntity() *shared.ServiceDependencyEntity {
	if o == nil {
		return nil
	}
	return o.ServiceDependencyEntity
}
