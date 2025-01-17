// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type GetV1PostMortemsReportsReportIDReasonsRequest struct {
	ReportID string `pathParam:"style=simple,explode=false,name=report_id"`
	Page     *int   `queryParam:"style=form,explode=true,name=page"`
	PerPage  *int   `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *GetV1PostMortemsReportsReportIDReasonsRequest) GetReportID() string {
	if o == nil {
		return ""
	}
	return o.ReportID
}

func (o *GetV1PostMortemsReportsReportIDReasonsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1PostMortemsReportsReportIDReasonsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type GetV1PostMortemsReportsReportIDReasonsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List all contributing factors to an incident
	PostMortemsReasonEntityPaginated *shared.PostMortemsReasonEntityPaginated
}

func (o *GetV1PostMortemsReportsReportIDReasonsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1PostMortemsReportsReportIDReasonsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1PostMortemsReportsReportIDReasonsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1PostMortemsReportsReportIDReasonsResponse) GetPostMortemsReasonEntityPaginated() *shared.PostMortemsReasonEntityPaginated {
	if o == nil {
		return nil
	}
	return o.PostMortemsReasonEntityPaginated
}
