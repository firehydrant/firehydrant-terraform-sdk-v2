// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type EnvironmentEntryEntity struct {
	ActiveIncidents   types.String             `tfsdk:"active_incidents"`
	CreatedAt         types.String             `tfsdk:"created_at"`
	Description       types.String             `tfsdk:"description"`
	ExternalResources []ExternalResourceEntity `tfsdk:"external_resources"`
	ID                types.String             `tfsdk:"id"`
	Name              types.String             `tfsdk:"name"`
	Slug              types.String             `tfsdk:"slug"`
	UpdatedAt         types.String             `tfsdk:"updated_at"`
}
