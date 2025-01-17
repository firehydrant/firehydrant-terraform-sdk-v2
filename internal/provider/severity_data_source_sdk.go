// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SeverityDataSourceModel) RefreshFromSharedSeverityEntity(resp *shared.SeverityEntity) {
	if resp != nil {
		r.Color = types.StringPointerValue(resp.Color)
		if resp.CreatedAt != nil {
			r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		} else {
			r.CreatedAt = types.StringNull()
		}
		r.Description = types.StringPointerValue(resp.Description)
		if resp.Position != nil {
			r.Position = types.Int64Value(int64(*resp.Position))
		} else {
			r.Position = types.Int64Null()
		}
		r.Slug = types.StringPointerValue(resp.Slug)
		r.SystemRecord = types.BoolPointerValue(resp.SystemRecord)
		r.Type = types.StringPointerValue(resp.Type)
		if resp.UpdatedAt != nil {
			r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
		} else {
			r.UpdatedAt = types.StringNull()
		}
	}
}
