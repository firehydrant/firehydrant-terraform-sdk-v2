// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &WebhooksDataSource{}
var _ datasource.DataSourceWithConfigure = &WebhooksDataSource{}

func NewWebhooksDataSource() datasource.DataSource {
	return &WebhooksDataSource{}
}

// WebhooksDataSource is the data source implementation.
type WebhooksDataSource struct {
	client *sdk.FirehydrantTerraformSDK
}

// WebhooksDataSourceModel describes the data model.
type WebhooksDataSourceModel struct {
	CreatedAt     types.String `tfsdk:"created_at"`
	Email         types.String `tfsdk:"email"`
	ID            types.String `tfsdk:"id"`
	Name          types.String `tfsdk:"name"`
	Page          types.Int64  `tfsdk:"page"`
	PerPage       types.Int64  `tfsdk:"per_page"`
	Source        types.String `tfsdk:"source"`
	State         types.String `tfsdk:"state"`
	Subscriptions types.String `tfsdk:"subscriptions"`
	UpdatedAt     types.String `tfsdk:"updated_at"`
	URL           types.String `tfsdk:"url"`
}

// Metadata returns the data source type name.
func (r *WebhooksDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webhooks"
}

// Schema defines the schema for the data source.
func (r *WebhooksDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Webhooks DataSource",

		Attributes: map[string]schema.Attribute{
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"email": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"page": schema.Int64Attribute{
				Optional: true,
			},
			"per_page": schema.Int64Attribute{
				Optional: true,
			},
			"source": schema.StringAttribute{
				Computed: true,
			},
			"state": schema.StringAttribute{
				Computed: true,
			},
			"subscriptions": schema.StringAttribute{
				Computed: true,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
			"url": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *WebhooksDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.FirehydrantTerraformSDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.FirehydrantTerraformSDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *WebhooksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *WebhooksDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	page := new(int)
	if !data.Page.IsUnknown() && !data.Page.IsNull() {
		*page = int(data.Page.ValueInt64())
	} else {
		page = nil
	}
	perPage := new(int)
	if !data.PerPage.IsUnknown() && !data.PerPage.IsNull() {
		*perPage = int(data.PerPage.ValueInt64())
	} else {
		perPage = nil
	}
	request := operations.GetV1WebhooksRequest{
		Page:    page,
		PerPage: perPage,
	}
	res, err := r.client.Webhooks.List(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
