// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetV1CatalogsCatalogIDRefreshRequest struct {
	CatalogID string `pathParam:"style=simple,explode=false,name=catalog_id"`
}

func (o *GetV1CatalogsCatalogIDRefreshRequest) GetCatalogID() string {
	if o == nil {
		return ""
	}
	return o.CatalogID
}

type GetV1CatalogsCatalogIDRefreshResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetV1CatalogsCatalogIDRefreshResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1CatalogsCatalogIDRefreshResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1CatalogsCatalogIDRefreshResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
