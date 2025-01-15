// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type GetV1ChangeTypesRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *GetV1ChangeTypesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1ChangeTypesRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type GetV1ChangeTypesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Lists all change types
	ChangeTypeEntityPaginated *shared.ChangeTypeEntityPaginated
}

func (o *GetV1ChangeTypesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1ChangeTypesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1ChangeTypesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1ChangeTypesResponse) GetChangeTypeEntityPaginated() *shared.ChangeTypeEntityPaginated {
	if o == nil {
		return nil
	}
	return o.ChangeTypeEntityPaginated
}
