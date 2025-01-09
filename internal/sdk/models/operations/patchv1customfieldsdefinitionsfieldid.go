// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type PatchV1CustomFieldsDefinitionsFieldIDRequest struct {
	FieldID                               string                                       `pathParam:"style=simple,explode=false,name=field_id"`
	PatchV1CustomFieldsDefinitionsFieldID shared.PatchV1CustomFieldsDefinitionsFieldID `request:"mediaType=application/json"`
}

func (o *PatchV1CustomFieldsDefinitionsFieldIDRequest) GetFieldID() string {
	if o == nil {
		return ""
	}
	return o.FieldID
}

func (o *PatchV1CustomFieldsDefinitionsFieldIDRequest) GetPatchV1CustomFieldsDefinitionsFieldID() shared.PatchV1CustomFieldsDefinitionsFieldID {
	if o == nil {
		return shared.PatchV1CustomFieldsDefinitionsFieldID{}
	}
	return o.PatchV1CustomFieldsDefinitionsFieldID
}

type PatchV1CustomFieldsDefinitionsFieldIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Update a single custom field definition
	OrganizationsCustomFieldDefinitionEntity *shared.OrganizationsCustomFieldDefinitionEntity
}

func (o *PatchV1CustomFieldsDefinitionsFieldIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchV1CustomFieldsDefinitionsFieldIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchV1CustomFieldsDefinitionsFieldIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchV1CustomFieldsDefinitionsFieldIDResponse) GetOrganizationsCustomFieldDefinitionEntity() *shared.OrganizationsCustomFieldDefinitionEntity {
	if o == nil {
		return nil
	}
	return o.OrganizationsCustomFieldDefinitionEntity
}
