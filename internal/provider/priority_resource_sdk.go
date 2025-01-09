// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
)

func (r *PriorityResourceModel) ToSharedPostV1Priorities() *shared.PostV1Priorities {
	var slug string
	slug = r.Slug.ValueString()

	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	defaultVar := new(bool)
	if !r.Default.IsUnknown() && !r.Default.IsNull() {
		*defaultVar = r.Default.ValueBool()
	} else {
		defaultVar = nil
	}
	out := shared.PostV1Priorities{
		Slug:        slug,
		Description: description,
		Default:     defaultVar,
	}
	return &out
}

func (r *PriorityResourceModel) ToSharedPatchV1PrioritiesPrioritySlug() *shared.PatchV1PrioritiesPrioritySlug {
	slug := new(string)
	if !r.Slug.IsUnknown() && !r.Slug.IsNull() {
		*slug = r.Slug.ValueString()
	} else {
		slug = nil
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	defaultVar := new(bool)
	if !r.Default.IsUnknown() && !r.Default.IsNull() {
		*defaultVar = r.Default.ValueBool()
	} else {
		defaultVar = nil
	}
	out := shared.PatchV1PrioritiesPrioritySlug{
		Slug:        slug,
		Description: description,
		Default:     defaultVar,
	}
	return &out
}
