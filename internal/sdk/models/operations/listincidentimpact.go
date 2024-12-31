// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type PathParamType string

const (
	PathParamTypeEnvironments    PathParamType = "environments"
	PathParamTypeFunctionalities PathParamType = "functionalities"
	PathParamTypeServices        PathParamType = "services"
	PathParamTypeCustomers       PathParamType = "customers"
)

func (e PathParamType) ToPointer() *PathParamType {
	return &e
}
func (e *PathParamType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "environments":
		fallthrough
	case "functionalities":
		fallthrough
	case "services":
		fallthrough
	case "customers":
		*e = PathParamType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PathParamType: %v", v)
	}
}

type ListIncidentImpactRequest struct {
	IncidentID string        `pathParam:"style=simple,explode=false,name=incident_id"`
	Type       PathParamType `pathParam:"style=simple,explode=false,name=type"`
}

func (o *ListIncidentImpactRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *ListIncidentImpactRequest) GetType() PathParamType {
	if o == nil {
		return PathParamType("")
	}
	return o.Type
}

type ListIncidentImpactResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List impacted infrastructure on an incident by specifying type
	IncidentImpactEntityPaginated *shared.IncidentImpactEntityPaginated
	// A collection of codes that generally means the end user got something wrong in making the request
	BadRequest *shared.BadRequest
	// A collection of codes that generally means the client was not authenticated correctly for the request they want to make
	Unauthorized *shared.Unauthorized
	// Status codes relating to the resource/entity they are requesting not being found or endpoints/routes not existing
	NotFound *shared.NotFound
	// Status codes relating to the client being rate limited by the server
	RateLimited *shared.RateLimited
	// A collection of status codes that generally mean the server failed in an unexpected way
	InternalServerError *shared.InternalServerError
	// Timeouts occurred with the request
	Timeout *shared.Timeout
}

func (o *ListIncidentImpactResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListIncidentImpactResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListIncidentImpactResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListIncidentImpactResponse) GetIncidentImpactEntityPaginated() *shared.IncidentImpactEntityPaginated {
	if o == nil {
		return nil
	}
	return o.IncidentImpactEntityPaginated
}

func (o *ListIncidentImpactResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *ListIncidentImpactResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *ListIncidentImpactResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *ListIncidentImpactResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *ListIncidentImpactResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *ListIncidentImpactResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}
