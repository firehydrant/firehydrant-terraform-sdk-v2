// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type ImportsImportErrorEntity struct {
	CreatedAt types.String                            `tfsdk:"created_at"`
	Data      *TeamEntity1                            `tfsdk:"data"`
	ID        types.String                            `tfsdk:"id"`
	Message   types.String                            `tfsdk:"message"`
	Resource  *ImportsImportErrorEntityResourceEntity `tfsdk:"resource"`
}
