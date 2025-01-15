// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SeverityResourceModel) ToSharedPostV1Severities() *shared.PostV1Severities {
	var slug string
	slug = r.Slug.ValueString()

	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	position := new(int)
	if !r.Position.IsUnknown() && !r.Position.IsNull() {
		*position = int(r.Position.ValueInt64())
	} else {
		position = nil
	}
	color := new(shared.Color)
	if !r.Color.IsUnknown() && !r.Color.IsNull() {
		*color = shared.Color(r.Color.ValueString())
	} else {
		color = nil
	}
	out := shared.PostV1Severities{
		Slug:        slug,
		Description: description,
		Position:    position,
		Color:       color,
	}
	return &out
}

func (r *SeverityResourceModel) RefreshFromSharedSeverityEntity(resp *shared.SeverityEntity) {
	if resp != nil {
		if resp.Color != nil {
			r.Color = types.StringValue(string(*resp.Color))
		} else {
			r.Color = types.StringNull()
		}
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

func (r *SeverityResourceModel) ToSharedPatchV1SeveritiesSeveritySlug() *shared.PatchV1SeveritiesSeveritySlug {
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
	position := new(int)
	if !r.Position.IsUnknown() && !r.Position.IsNull() {
		*position = int(r.Position.ValueInt64())
	} else {
		position = nil
	}
	color := new(shared.PatchV1SeveritiesSeveritySlugColor)
	if !r.Color.IsUnknown() && !r.Color.IsNull() {
		*color = shared.PatchV1SeveritiesSeveritySlugColor(r.Color.ValueString())
	} else {
		color = nil
	}
	out := shared.PatchV1SeveritiesSeveritySlug{
		Slug:        slug,
		Description: description,
		Position:    position,
		Color:       color,
	}
	return &out
}
