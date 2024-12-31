// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type ListAwsConnectionsRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// AWS account ID containing the role to be assumed
	AwsAccountID *string `queryParam:"style=form,explode=true,name=aws_account_id"`
	// ARN of the role to be assumed
	TargetArn *string `queryParam:"style=form,explode=true,name=target_arn"`
	// The external ID supplied when assuming the role
	ExternalID *string `queryParam:"style=form,explode=true,name=external_id"`
}

func (o *ListAwsConnectionsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListAwsConnectionsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *ListAwsConnectionsRequest) GetAwsAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AwsAccountID
}

func (o *ListAwsConnectionsRequest) GetTargetArn() *string {
	if o == nil {
		return nil
	}
	return o.TargetArn
}

func (o *ListAwsConnectionsRequest) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

type ListAwsConnectionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Lists the available and configured AWS integration connections for the authenticated organization.
	IntegrationsAwsConnectionEntityPaginated *shared.IntegrationsAwsConnectionEntityPaginated
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

func (o *ListAwsConnectionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAwsConnectionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAwsConnectionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListAwsConnectionsResponse) GetIntegrationsAwsConnectionEntityPaginated() *shared.IntegrationsAwsConnectionEntityPaginated {
	if o == nil {
		return nil
	}
	return o.IntegrationsAwsConnectionEntityPaginated
}

func (o *ListAwsConnectionsResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *ListAwsConnectionsResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *ListAwsConnectionsResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *ListAwsConnectionsResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *ListAwsConnectionsResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *ListAwsConnectionsResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}
