// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type IncidentsContextObjectEntity struct {
	ContextDescription types.String `tfsdk:"context_description"`
	ContextTag         types.String `tfsdk:"context_tag"`
	ObjectID           types.String `tfsdk:"object_id"`
	ObjectType         types.String `tfsdk:"object_type"`
}
