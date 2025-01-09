// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type PostV1PostMortemsReportsReportIDReasonsRequest struct {
	ReportID                                string                                         `pathParam:"style=simple,explode=false,name=report_id"`
	PostV1PostMortemsReportsReportIDReasons shared.PostV1PostMortemsReportsReportIDReasons `request:"mediaType=application/json"`
}

func (o *PostV1PostMortemsReportsReportIDReasonsRequest) GetReportID() string {
	if o == nil {
		return ""
	}
	return o.ReportID
}

func (o *PostV1PostMortemsReportsReportIDReasonsRequest) GetPostV1PostMortemsReportsReportIDReasons() shared.PostV1PostMortemsReportsReportIDReasons {
	if o == nil {
		return shared.PostV1PostMortemsReportsReportIDReasons{}
	}
	return o.PostV1PostMortemsReportsReportIDReasons
}

type PostV1PostMortemsReportsReportIDReasonsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Add a new contributing factor to an incident
	PostMortemsReasonEntity *shared.PostMortemsReasonEntity
}

func (o *PostV1PostMortemsReportsReportIDReasonsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostV1PostMortemsReportsReportIDReasonsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostV1PostMortemsReportsReportIDReasonsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostV1PostMortemsReportsReportIDReasonsResponse) GetPostMortemsReasonEntity() *shared.PostMortemsReasonEntity {
	if o == nil {
		return nil
	}
	return o.PostMortemsReasonEntity
}
