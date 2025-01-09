// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/provider/types"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &FunctionalityDataSource{}
var _ datasource.DataSourceWithConfigure = &FunctionalityDataSource{}

func NewFunctionalityDataSource() datasource.DataSource {
	return &FunctionalityDataSource{}
}

// FunctionalityDataSource is the data source implementation.
type FunctionalityDataSource struct {
	client *sdk.FirehydrantTerraformSDK
}

// FunctionalityDataSourceModel describes the data model.
type FunctionalityDataSourceModel struct {
	ActiveIncidents       []types.String                   `tfsdk:"active_incidents"`
	AlertOnAdd            types.Bool                       `tfsdk:"alert_on_add"`
	AutoAddRespondingTeam types.Bool                       `tfsdk:"auto_add_responding_team"`
	CreatedAt             types.String                     `tfsdk:"created_at"`
	Description           types.String                     `tfsdk:"description"`
	ExternalResources     []tfTypes.ExternalResourceEntity `tfsdk:"external_resources"`
	ID                    types.String                     `tfsdk:"id"`
	Labels                map[string]types.String          `tfsdk:"labels"`
	Links                 []tfTypes.LinksEntity            `tfsdk:"links"`
	Name                  types.String                     `tfsdk:"name"`
	Owner                 *tfTypes.TeamEntity1             `tfsdk:"owner"`
	Slug                  types.String                     `tfsdk:"slug"`
	UpdatedAt             types.String                     `tfsdk:"updated_at"`
	UpdatedBy             *tfTypes.AuthorEntity            `tfsdk:"updated_by"`
}

// Metadata returns the data source type name.
func (r *FunctionalityDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_functionality"
}

// Schema defines the schema for the data source.
func (r *FunctionalityDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Functionality DataSource",

		Attributes: map[string]schema.Attribute{
			"active_incidents": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `List of active incident guids`,
			},
			"alert_on_add": schema.BoolAttribute{
				Computed: true,
			},
			"auto_add_responding_team": schema.BoolAttribute{
				Computed: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"external_resources": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"connection_id": schema.StringAttribute{
							Computed: true,
						},
						"connection_name": schema.StringAttribute{
							Computed: true,
						},
						"connection_type": schema.StringAttribute{
							Computed: true,
						},
						"created_at": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"remote_id": schema.StringAttribute{
							Computed: true,
						},
						"remote_url": schema.StringAttribute{
							Computed: true,
						},
						"updated_at": schema.StringAttribute{
							Computed: true,
						},
					},
				},
				Description: `Information about known linkages to representations of services outside of FireHydrant.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"labels": schema.MapAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `An object of label key and values`,
			},
			"links": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"href_url": schema.StringAttribute{
							Computed: true,
						},
						"icon_url": schema.StringAttribute{
							Computed: true,
						},
						"id": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
					},
				},
				Description: `List of links attached to this functionality.`,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"owner": schema.SingleNestedAttribute{
				Computed:    true,
				Description: `TeamEntity model`,
			},
			"slug": schema.StringAttribute{
				Computed: true,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
			"updated_by": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"email": schema.StringAttribute{
						Computed: true,
					},
					"id": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"source": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (r *FunctionalityDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *FunctionalityDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *FunctionalityDataSourceModel
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

	var functionalityID string
	functionalityID = data.ID.ValueString()

	request := operations.GetV1FunctionalitiesFunctionalityIDRequest{
		FunctionalityID: functionalityID,
	}
	res, err := r.client.Functionalities.Get(ctx, request)
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
	if !(res.FunctionalityEntity != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedFunctionalityEntity(res.FunctionalityEntity)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}