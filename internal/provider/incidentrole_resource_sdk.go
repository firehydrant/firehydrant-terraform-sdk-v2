// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
)

func (r *IncidentRoleResourceModel) ToSharedPostV1IncidentRoles() *shared.PostV1IncidentRoles {
	var name string
	name = r.Name.ValueString()

	var summary string
	summary = r.Summary.ValueString()

	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	out := shared.PostV1IncidentRoles{
		Name:        name,
		Summary:     summary,
		Description: description,
	}
	return &out
}

func (r *IncidentRoleResourceModel) ToSharedPatchV1IncidentRolesIncidentRoleID() *shared.PatchV1IncidentRolesIncidentRoleID {
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	summary := new(string)
	if !r.Summary.IsUnknown() && !r.Summary.IsNull() {
		*summary = r.Summary.ValueString()
	} else {
		summary = nil
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	out := shared.PatchV1IncidentRolesIncidentRoleID{
		Name:        name,
		Summary:     summary,
		Description: description,
	}
	return &out
}
