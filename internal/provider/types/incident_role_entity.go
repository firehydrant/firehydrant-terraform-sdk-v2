// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type IncidentRoleEntity struct {
	CreatedAt   types.String `tfsdk:"created_at"`
	Description types.String `tfsdk:"description"`
	DiscardedAt types.String `tfsdk:"discarded_at"`
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Summary     types.String `tfsdk:"summary"`
	UpdatedAt   types.String `tfsdk:"updated_at"`
}