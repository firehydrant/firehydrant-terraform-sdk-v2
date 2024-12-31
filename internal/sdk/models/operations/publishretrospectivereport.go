// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type PublishRetrospectiveReportRequest struct {
	ReportID                                string                                         `pathParam:"style=simple,explode=false,name=report_id"`
	PostV1PostMortemsReportsReportIDPublish shared.PostV1PostMortemsReportsReportIDPublish `request:"mediaType=application/json"`
}

func (o *PublishRetrospectiveReportRequest) GetReportID() string {
	if o == nil {
		return ""
	}
	return o.ReportID
}

func (o *PublishRetrospectiveReportRequest) GetPostV1PostMortemsReportsReportIDPublish() shared.PostV1PostMortemsReportsReportIDPublish {
	if o == nil {
		return shared.PostV1PostMortemsReportsReportIDPublish{}
	}
	return o.PostV1PostMortemsReportsReportIDPublish
}

type PublishRetrospectiveReportResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Marks an incident retrospective as published and emails all of the participants in the report the summary
	PostMortemsPostMortemReportEntity *shared.PostMortemsPostMortemReportEntity
	// Bad Request
	ErrorEntity *shared.ErrorEntity
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

func (o *PublishRetrospectiveReportResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PublishRetrospectiveReportResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PublishRetrospectiveReportResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PublishRetrospectiveReportResponse) GetPostMortemsPostMortemReportEntity() *shared.PostMortemsPostMortemReportEntity {
	if o == nil {
		return nil
	}
	return o.PostMortemsPostMortemReportEntity
}

func (o *PublishRetrospectiveReportResponse) GetErrorEntity() *shared.ErrorEntity {
	if o == nil {
		return nil
	}
	return o.ErrorEntity
}

func (o *PublishRetrospectiveReportResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *PublishRetrospectiveReportResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *PublishRetrospectiveReportResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *PublishRetrospectiveReportResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *PublishRetrospectiveReportResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *PublishRetrospectiveReportResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}
