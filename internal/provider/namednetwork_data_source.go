// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/speakeasy/terraform-provider-xshield-sdk/internal/provider/types"
	"github.com/speakeasy/terraform-provider-xshield-sdk/internal/sdk"
	"github.com/speakeasy/terraform-provider-xshield-sdk/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &NamedNetworkDataSource{}
var _ datasource.DataSourceWithConfigure = &NamedNetworkDataSource{}

func NewNamedNetworkDataSource() datasource.DataSource {
	return &NamedNetworkDataSource{}
}

// NamedNetworkDataSource is the data source implementation.
type NamedNetworkDataSource struct {
	client *sdk.XshieldSDK
}

// NamedNetworkDataSourceModel describes the data model.
type NamedNetworkDataSourceModel struct {
	AssignedByTagBasedPolicy              types.Bool                  `tfsdk:"assigned_by_tag_based_policy"`
	ID                                    types.String                `tfsdk:"id"`
	IPRanges                              []tfTypes.NamednetworkRange `tfsdk:"ip_ranges"`
	IsOOBNetwork                          types.Bool                  `tfsdk:"is_oob_network"`
	NamedNetworkAssignments               types.Int64                 `tfsdk:"named_network_assignments"`
	NamedNetworkDescription               types.String                `tfsdk:"named_network_description"`
	NamedNetworkName                      types.String                `tfsdk:"named_network_name"`
	NamednetworkTagBasedPolicyAssignments types.Int64                 `tfsdk:"namednetwork_tag_based_policy_assignments"`
	ProgramAsInternet                     types.Bool                  `tfsdk:"program_as_internet"`
	ProgramAsIntranet                     types.Bool                  `tfsdk:"program_as_intranet"`
	Provider                              types.String                `tfsdk:"provider"`
	Region                                types.String                `tfsdk:"region"`
	Service                               types.String                `tfsdk:"service"`
	TotalComments                         types.Int64                 `tfsdk:"total_comments"`
	TotalCount                            types.Int64                 `tfsdk:"total_count"`
	UsergroupNamedNetworkAssignments      types.Int64                 `tfsdk:"usergroup_named_network_assignments"`
}

// Metadata returns the data source type name.
func (r *NamedNetworkDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_named_network"
}

// Schema defines the schema for the data source.
func (r *NamedNetworkDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "NamedNetwork DataSource",

		Attributes: map[string]schema.Attribute{
			"assigned_by_tag_based_policy": schema.BoolAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ip_ranges": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"ip_count": schema.Int64Attribute{
							Computed: true,
						},
						"ip_range": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"is_oob_network": schema.BoolAttribute{
				Computed: true,
			},
			"named_network_assignments": schema.Int64Attribute{
				Computed: true,
			},
			"named_network_description": schema.StringAttribute{
				Computed: true,
			},
			"named_network_name": schema.StringAttribute{
				Computed: true,
			},
			"namednetwork_tag_based_policy_assignments": schema.Int64Attribute{
				Computed: true,
			},
			"program_as_internet": schema.BoolAttribute{
				Computed: true,
			},
			"program_as_intranet": schema.BoolAttribute{
				Computed: true,
			},
			"provider": schema.StringAttribute{
				Computed: true,
			},
			"region": schema.StringAttribute{
				Computed: true,
			},
			"service": schema.StringAttribute{
				Computed: true,
			},
			"total_comments": schema.Int64Attribute{
				Computed: true,
			},
			"total_count": schema.Int64Attribute{
				Computed: true,
			},
			"usergroup_named_network_assignments": schema.Int64Attribute{
				Computed: true,
			},
		},
	}
}

func (r *NamedNetworkDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.XshieldSDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.XshieldSDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *NamedNetworkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *NamedNetworkDataSourceModel
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

	namedNetworkID := data.ID.ValueString()
	request := operations.GetNamedNetworkRequest{
		NamedNetworkID: namedNetworkID,
	}
	res, err := r.client.Namednetworks.GetNamedNetwork(ctx, request)
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
	if !(res.NamednetworkNamedNetwork != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedNamednetworkNamedNetwork(res.NamednetworkNamedNetwork)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
