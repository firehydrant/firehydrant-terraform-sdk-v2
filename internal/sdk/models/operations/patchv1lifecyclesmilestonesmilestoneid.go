// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type PatchV1LifecyclesMilestonesMilestoneIDRequestBody struct {
	// The name of the milestone
	Name *string `form:"name=name"`
	// A long-form description of the milestone's purpose
	Description *string `form:"name=description"`
	// A unique identifier for the milestone. If not provided, one will be generated from the name.
	Slug *string `form:"name=slug"`
	// The position of the milestone within the phase. If not provided, the milestone will be added as the last milestone in the phase.
	Position *int `form:"name=position"`
	// The ID of a later milestone that cannot be started until this milestone has a timestamp populated
	RequiredAtMilestoneID *string `form:"name=required_at_milestone_id"`
}

func (o *PatchV1LifecyclesMilestonesMilestoneIDRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PatchV1LifecyclesMilestonesMilestoneIDRequestBody) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PatchV1LifecyclesMilestonesMilestoneIDRequestBody) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *PatchV1LifecyclesMilestonesMilestoneIDRequestBody) GetPosition() *int {
	if o == nil {
		return nil
	}
	return o.Position
}

func (o *PatchV1LifecyclesMilestonesMilestoneIDRequestBody) GetRequiredAtMilestoneID() *string {
	if o == nil {
		return nil
	}
	return o.RequiredAtMilestoneID
}

type PatchV1LifecyclesMilestonesMilestoneIDRequest struct {
	MilestoneID string                                             `pathParam:"style=simple,explode=false,name=milestone_id"`
	RequestBody *PatchV1LifecyclesMilestonesMilestoneIDRequestBody `request:"mediaType=application/x-www-form-urlencoded"`
}

func (o *PatchV1LifecyclesMilestonesMilestoneIDRequest) GetMilestoneID() string {
	if o == nil {
		return ""
	}
	return o.MilestoneID
}

func (o *PatchV1LifecyclesMilestonesMilestoneIDRequest) GetRequestBody() *PatchV1LifecyclesMilestonesMilestoneIDRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type PatchV1LifecyclesMilestonesMilestoneIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Update a milestone
	LifecyclesPhaseEntity *shared.LifecyclesPhaseEntity
}

func (o *PatchV1LifecyclesMilestonesMilestoneIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchV1LifecyclesMilestonesMilestoneIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchV1LifecyclesMilestonesMilestoneIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchV1LifecyclesMilestonesMilestoneIDResponse) GetLifecyclesPhaseEntity() *shared.LifecyclesPhaseEntity {
	if o == nil {
		return nil
	}
	return o.LifecyclesPhaseEntity
}