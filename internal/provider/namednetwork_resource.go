// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/speakeasy/terraform-provider-xshield-sdk/internal/provider/types"
	"github.com/speakeasy/terraform-provider-xshield-sdk/internal/sdk"
	"github.com/speakeasy/terraform-provider-xshield-sdk/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NamedNetworkResource{}
var _ resource.ResourceWithImportState = &NamedNetworkResource{}

func NewNamedNetworkResource() resource.Resource {
	return &NamedNetworkResource{}
}

// NamedNetworkResource defines the resource implementation.
type NamedNetworkResource struct {
	client *sdk.XshieldSDK
}

// NamedNetworkResourceModel describes the resource data model.
type NamedNetworkResourceModel struct {
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
	Region                                types.String                `tfsdk:"region"`
	Service                               types.String                `tfsdk:"service"`
	TotalComments                         types.Int64                 `tfsdk:"total_comments"`
	TotalCount                            types.Int64                 `tfsdk:"total_count"`
	UsergroupNamedNetworkAssignments      types.Int64                 `tfsdk:"usergroup_named_network_assignments"`
}

func (r *NamedNetworkResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_named_network"
}

func (r *NamedNetworkResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "NamedNetwork Resource",
		Attributes: map[string]schema.Attribute{
			"assigned_by_tag_based_policy": schema.BoolAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"ip_ranges": schema.ListNestedAttribute{
				Computed: true,
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"ip_count": schema.Int64Attribute{
							Computed: true,
						},
						"ip_range": schema.StringAttribute{
							Computed: true,
							Optional: true,
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
				Optional: true,
			},
			"named_network_name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"namednetwork_tag_based_policy_assignments": schema.Int64Attribute{
				Computed: true,
			},
			"program_as_internet": schema.BoolAttribute{
				Computed: true,
			},
			"program_as_intranet": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"region": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"service": schema.StringAttribute{
				Computed: true,
				Optional: true,
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

func (r *NamedNetworkResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.XshieldSDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.XshieldSDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *NamedNetworkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *NamedNetworkResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToSharedNamednetworkNamedNetwork()
	res, err := r.client.Namednetworks.CreateNamedNetwork(ctx, request)
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
	if res.StatusCode != 201 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NamedNetworkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *NamedNetworkResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

func (r *NamedNetworkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *NamedNetworkResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	namedNetworkID := data.ID.ValueString()
	namednetworkNamedNetwork := *data.ToSharedNamednetworkNamedNetwork()
	request := operations.UpdateNamedNetworkMetadataRequest{
		NamedNetworkID:           namedNetworkID,
		NamednetworkNamedNetwork: namednetworkNamedNetwork,
	}
	res, err := r.client.Namednetworks.UpdateNamedNetworkMetadata(ctx, request)
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
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NamedNetworkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *NamedNetworkResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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
	request := operations.DeleteNamedNetworkRequest{
		NamedNetworkID: namedNetworkID,
	}
	res, err := r.client.Namednetworks.DeleteNamedNetwork(ctx, request)
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
	if res.StatusCode != 202 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *NamedNetworkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
