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
var _ datasource.DataSource = &TaskListsDataSource{}
var _ datasource.DataSourceWithConfigure = &TaskListsDataSource{}

func NewTaskListsDataSource() datasource.DataSource {
	return &TaskListsDataSource{}
}

// TaskListsDataSource is the data source implementation.
type TaskListsDataSource struct {
	client *sdk.FirehydrantTerraformSDK
}

// TaskListsDataSourceModel describes the data model.
type TaskListsDataSourceModel struct {
	CreatedAt   types.String `tfsdk:"created_at"`
	Description types.String `tfsdk:"description"`
	Email       types.String `tfsdk:"email"`
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Page        types.Int64  `tfsdk:"page"`
	PerPage     types.Int64  `tfsdk:"per_page"`
	Source      types.String `tfsdk:"source"`
	Summary     types.String `tfsdk:"summary"`
	UpdatedAt   types.String `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *TaskListsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_task_lists"
}

// Schema defines the schema for the data source.
func (r *TaskListsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "TaskLists DataSource",

		Attributes: map[string]schema.Attribute{
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
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
			"summary": schema.StringAttribute{
				Computed: true,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *TaskListsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *TaskListsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *TaskListsDataSourceModel
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
	request := operations.GetV1TaskListsRequest{
		Page:    page,
		PerPage: perPage,
	}
	res, err := r.client.TaskLists.List(ctx, request)
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
