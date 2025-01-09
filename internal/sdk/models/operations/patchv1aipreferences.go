// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type PatchV1AiPreferencesRequestBody struct {
	// Whether to enable AI features
	Ai *bool `form:"name=ai"`
	// Whether to enable incident summaries
	Summaries *bool `form:"name=summaries"`
	// Whether to enable incident descriptions
	Description *bool `form:"name=description"`
	// Whether to enable incident impact
	Impact *bool `form:"name=impact"`
	// Whether to enable incident updates
	Updates *bool `form:"name=updates"`
	// Whether to enable incident retrospectives
	Retros *bool `form:"name=retros"`
	// Whether to enable incident followups
	Followups *bool `form:"name=followups"`
	// Whether to enable similar incidents
	SimilarIncidents *bool `form:"name=similar_incidents"`
}

func (o *PatchV1AiPreferencesRequestBody) GetAi() *bool {
	if o == nil {
		return nil
	}
	return o.Ai
}

func (o *PatchV1AiPreferencesRequestBody) GetSummaries() *bool {
	if o == nil {
		return nil
	}
	return o.Summaries
}

func (o *PatchV1AiPreferencesRequestBody) GetDescription() *bool {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PatchV1AiPreferencesRequestBody) GetImpact() *bool {
	if o == nil {
		return nil
	}
	return o.Impact
}

func (o *PatchV1AiPreferencesRequestBody) GetUpdates() *bool {
	if o == nil {
		return nil
	}
	return o.Updates
}

func (o *PatchV1AiPreferencesRequestBody) GetRetros() *bool {
	if o == nil {
		return nil
	}
	return o.Retros
}

func (o *PatchV1AiPreferencesRequestBody) GetFollowups() *bool {
	if o == nil {
		return nil
	}
	return o.Followups
}

func (o *PatchV1AiPreferencesRequestBody) GetSimilarIncidents() *bool {
	if o == nil {
		return nil
	}
	return o.SimilarIncidents
}

type PatchV1AiPreferencesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Updates the AI preferences
	AIEntitiesPreferencesEntity *shared.AIEntitiesPreferencesEntity
}

func (o *PatchV1AiPreferencesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchV1AiPreferencesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchV1AiPreferencesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchV1AiPreferencesResponse) GetAIEntitiesPreferencesEntity() *shared.AIEntitiesPreferencesEntity {
	if o == nil {
		return nil
	}
	return o.AIEntitiesPreferencesEntity
}
