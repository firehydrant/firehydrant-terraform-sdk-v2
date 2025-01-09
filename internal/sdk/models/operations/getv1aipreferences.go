// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type GetV1AiPreferencesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Retrieves the current AI preferences
	AIEntitiesPreferencesEntity *shared.AIEntitiesPreferencesEntity
}

func (o *GetV1AiPreferencesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1AiPreferencesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1AiPreferencesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1AiPreferencesResponse) GetAIEntitiesPreferencesEntity() *shared.AIEntitiesPreferencesEntity {
	if o == nil {
		return nil
	}
	return o.AIEntitiesPreferencesEntity
}