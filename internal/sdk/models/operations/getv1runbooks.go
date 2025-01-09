// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

// QueryParamSort - Sort runbooks by their updated date. Accepts 'asc', 'desc'
type QueryParamSort string

const (
	QueryParamSortAsc  QueryParamSort = "asc"
	QueryParamSortDesc QueryParamSort = "desc"
)

func (e QueryParamSort) ToPointer() *QueryParamSort {
	return &e
}
func (e *QueryParamSort) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = QueryParamSort(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamSort: %v", v)
	}
}

type GetV1RunbooksRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// A query to search runbooks by their name
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// A query to search runbooks by their owners
	Owners *string `queryParam:"style=form,explode=true,name=owners"`
	// Sort runbooks by their updated date. Accepts 'asc', 'desc'
	Sort *QueryParamSort `queryParam:"style=form,explode=true,name=sort"`
}

func (o *GetV1RunbooksRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1RunbooksRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *GetV1RunbooksRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetV1RunbooksRequest) GetOwners() *string {
	if o == nil {
		return nil
	}
	return o.Owners
}

func (o *GetV1RunbooksRequest) GetSort() *QueryParamSort {
	if o == nil {
		return nil
	}
	return o.Sort
}

type GetV1RunbooksResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Lists all available runbooks.
	RunbookEntity *shared.RunbookEntity
}

func (o *GetV1RunbooksResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1RunbooksResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1RunbooksResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1RunbooksResponse) GetRunbookEntity() *shared.RunbookEntity {
	if o == nil {
		return nil
	}
	return o.RunbookEntity
}
