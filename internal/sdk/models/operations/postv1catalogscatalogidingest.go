// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type PostV1CatalogsCatalogIDIngestRequest struct {
	CatalogID                     string                               `pathParam:"style=simple,explode=false,name=catalog_id"`
	PostV1CatalogsCatalogIDIngest shared.PostV1CatalogsCatalogIDIngest `request:"mediaType=application/json"`
}

func (o *PostV1CatalogsCatalogIDIngestRequest) GetCatalogID() string {
	if o == nil {
		return ""
	}
	return o.CatalogID
}

func (o *PostV1CatalogsCatalogIDIngestRequest) GetPostV1CatalogsCatalogIDIngest() shared.PostV1CatalogsCatalogIDIngest {
	if o == nil {
		return shared.PostV1CatalogsCatalogIDIngest{}
	}
	return o.PostV1CatalogsCatalogIDIngest
}

type PostV1CatalogsCatalogIDIngestResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Accepts catalog data in the configured format and asyncronously processes the data to incorporate changes into service catalog.
	ImportsImportEntity *shared.ImportsImportEntity
}

func (o *PostV1CatalogsCatalogIDIngestResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostV1CatalogsCatalogIDIngestResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostV1CatalogsCatalogIDIngestResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostV1CatalogsCatalogIDIngestResponse) GetImportsImportEntity() *shared.ImportsImportEntity {
	if o == nil {
		return nil
	}
	return o.ImportsImportEntity
}
