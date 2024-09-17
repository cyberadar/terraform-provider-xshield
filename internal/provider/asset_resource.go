// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	tfTypes "github.com/colortokens/terraform-provider-xshield/internal/provider/types"
	"github.com/colortokens/terraform-provider-xshield/internal/sdk"
	"github.com/colortokens/terraform-provider-xshield/internal/sdk/models/operations"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AssetResource{}
var _ resource.ResourceWithImportState = &AssetResource{}

func NewAssetResource() resource.Resource {
	return &AssetResource{}
}

// AssetResource defines the resource implementation.
type AssetResource struct {
	client *sdk.Xshield
}

// AssetResourceModel describes the resource data model.
type AssetResourceModel struct {
	AgentID                             types.String                            `tfsdk:"agent_id"`
	AgentName                           types.String                            `tfsdk:"agent_name"`
	AgentStatus                         types.String                            `tfsdk:"agent_status"`
	AgentVersion                        types.String                            `tfsdk:"agent_version"`
	AssetAvailability                   types.String                            `tfsdk:"asset_availability"`
	AssetInternetFacing                 types.Bool                              `tfsdk:"asset_internet_facing"`
	AssetName                           types.String                            `tfsdk:"asset_name"`
	AssetRisk                           types.String                            `tfsdk:"asset_risk"`
	AttackSurface                       types.String                            `tfsdk:"attack_surface"`
	AttackSurfacePendingChanges         *tfTypes.PendingChanges                 `tfsdk:"attack_surface_pending_changes"`
	AutoSynchronizeEnabled              types.Bool                              `tfsdk:"auto_synchronize_enabled"`
	BlastRadius                         types.String                            `tfsdk:"blast_radius"`
	BlastRadiusPendingChanges           *tfTypes.PendingChanges                 `tfsdk:"blast_radius_pending_changes"`
	BusinessValue                       types.String                            `tfsdk:"business_value"`
	CloudTags                           []tfTypes.Tag                           `tfsdk:"cloud_tags"`
	ClusterIdentifier                   types.String                            `tfsdk:"cluster_identifier"`
	ContainerNamespace                  types.String                            `tfsdk:"container_namespace"`
	CoreTags                            map[string]types.String                 `tfsdk:"core_tags"`
	CPUCoreCount                        types.Int64                             `tfsdk:"cpu_core_count"`
	CPUName                             types.String                            `tfsdk:"cpu_name"`
	CurrentTrafficConfiguration         types.String                            `tfsdk:"current_traffic_configuration"`
	DeterministicID                     types.String                            `tfsdk:"deterministic_id"`
	DiskCapacityInGB                    types.Int64                             `tfsdk:"disk_capacity_in_gb"`
	HostName                            types.String                            `tfsdk:"host_name"`
	ID                                  types.String                            `tfsdk:"id"`
	InboundAssetStatus                  types.String                            `tfsdk:"inbound_asset_status"`
	InboundInternetPaths                *tfTypes.ReviewCoverage                 `tfsdk:"inbound_internet_paths"`
	InboundInternetPorts                *tfTypes.ReviewCoverage                 `tfsdk:"inbound_internet_ports"`
	InboundIntranetPaths                *tfTypes.ReviewCoverage                 `tfsdk:"inbound_intranet_paths"`
	InboundIntranetPorts                *tfTypes.ReviewCoverage                 `tfsdk:"inbound_intranet_ports"`
	Interfaces                          []tfTypes.NetworkInterface              `tfsdk:"interfaces"`
	KernelArchitecture                  types.String                            `tfsdk:"kernel_architecture"`
	KernelVersion                       types.String                            `tfsdk:"kernel_version"`
	LowestInboundAssetStatus            types.String                            `tfsdk:"lowest_inbound_asset_status"`
	LowestOutboundAssetStatus           types.String                            `tfsdk:"lowest_outbound_asset_status"`
	LowestProgressiveInboundAssetStatus types.String                            `tfsdk:"lowest_progressive_inbound_asset_status"`
	NamedNetworkChanges                 []tfTypes.MetadataNamedNetworkReference `tfsdk:"named_network_changes"`
	NamednetworksAssigned               types.Int64                             `tfsdk:"namednetworks_assigned"`
	OsName                              types.String                            `tfsdk:"os_name"`
	OutboundAssetStatus                 types.String                            `tfsdk:"outbound_asset_status"`
	OutboundInternetPaths               *tfTypes.ReviewCoverage                 `tfsdk:"outbound_internet_paths"`
	OutboundIntranetPaths               *tfTypes.ReviewCoverage                 `tfsdk:"outbound_intranet_paths"`
	PendingAttackSurfaceChanges         types.Bool                              `tfsdk:"pending_attack_surface_changes"`
	PendingBlastRadiusChanges           types.Bool                              `tfsdk:"pending_blast_radius_changes"`
	Platform                            types.String                            `tfsdk:"platform"`
	PoliciesAssigned                    types.Int64                             `tfsdk:"policies_assigned"`
	PolicyStatus                        types.String                            `tfsdk:"policy_status"`
	RAMCapacityInMB                     types.Int64                             `tfsdk:"ram_capacity_in_mb"`
	SecurityPatches                     types.Int64                             `tfsdk:"security_patches"`
	Tags                                []tfTypes.Tag                           `tfsdk:"tags"`
	TemplateChanges                     []tfTypes.TemplateReference             `tfsdk:"template_changes"`
	TemplatesAssigned                   types.Int64                             `tfsdk:"templates_assigned"`
	TotalComments                       types.Int64                             `tfsdk:"total_comments"`
	TotalInboundComments                types.Int64                             `tfsdk:"total_inbound_comments"`
	TotalOutboundComments               types.Int64                             `tfsdk:"total_outbound_comments"`
	TotalPaths                          types.Int64                             `tfsdk:"total_paths"`
	TotalPorts                          types.Int64                             `tfsdk:"total_ports"`
	TotalPortsPathRestricted            types.Int64                             `tfsdk:"total_ports_path_restricted"`
	Type                                types.String                            `tfsdk:"type"`
	UnreviewedPaths                     types.Int64                             `tfsdk:"unreviewed_paths"`
	UnreviewedPorts                     types.Int64                             `tfsdk:"unreviewed_ports"`
	UsergroupOutboundInternetPaths      *tfTypes.ReviewCoverage                 `tfsdk:"usergroup_outbound_internet_paths"`
	UsergroupOutboundIntranetPaths      *tfTypes.ReviewCoverage                 `tfsdk:"usergroup_outbound_intranet_paths"`
	UsergroupTotalPaths                 types.Int64                             `tfsdk:"usergroup_total_paths"`
	UsergroupUnreviewedPaths            types.Int64                             `tfsdk:"usergroup_unreviewed_paths"`
	Usergroups                          []tfTypes.AssetGroup                    `tfsdk:"usergroups"`
	Users                               []tfTypes.AssetUser                     `tfsdk:"users"`
	VendorInfo                          types.String                            `tfsdk:"vendor_info"`
	VirtualizationSystem                types.String                            `tfsdk:"virtualization_system"`
	Vulnerabilities                     types.Int64                             `tfsdk:"vulnerabilities"`
}

func (r *AssetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_asset"
}

func (r *AssetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Asset Resource",
		Attributes: map[string]schema.Attribute{
			"agent_id": schema.StringAttribute{
				Computed: true,
			},
			"agent_name": schema.StringAttribute{
				Computed: true,
			},
			"agent_status": schema.StringAttribute{
				Computed: true,
			},
			"agent_version": schema.StringAttribute{
				Computed: true,
			},
			"asset_availability": schema.StringAttribute{
				Computed: true,
			},
			"asset_internet_facing": schema.BoolAttribute{
				Computed: true,
			},
			"asset_name": schema.StringAttribute{
				Required: true,
			},
			"asset_risk": schema.StringAttribute{
				Computed: true,
			},
			"attack_surface": schema.StringAttribute{
				Computed: true,
			},
			"attack_surface_pending_changes": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"allow_templates": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"block_templates": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"internet_paths": schema.Int64Attribute{
						Computed: true,
					},
					"internet_ports": schema.Int64Attribute{
						Computed: true,
					},
					"intranet_change": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"intranet_paths": schema.Int64Attribute{
						Computed: true,
					},
					"intranet_ports": schema.Int64Attribute{
						Computed: true,
					},
					"namednetwork_change": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"progressive_sync_pending": schema.BoolAttribute{
						Computed: true,
					},
					"unassigned_allow_templates": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"unassigned_block_templates": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
				},
			},
			"auto_synchronize_enabled": schema.BoolAttribute{
				Computed: true,
			},
			"blast_radius": schema.StringAttribute{
				Computed: true,
			},
			"blast_radius_pending_changes": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"allow_templates": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"block_templates": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"internet_paths": schema.Int64Attribute{
						Computed: true,
					},
					"internet_ports": schema.Int64Attribute{
						Computed: true,
					},
					"intranet_change": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"intranet_paths": schema.Int64Attribute{
						Computed: true,
					},
					"intranet_ports": schema.Int64Attribute{
						Computed: true,
					},
					"namednetwork_change": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"progressive_sync_pending": schema.BoolAttribute{
						Computed: true,
					},
					"unassigned_allow_templates": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"unassigned_block_templates": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
				},
			},
			"business_value": schema.StringAttribute{
				Computed: true,
			},
			"cloud_tags": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Description: `Unique identifier for this tag instance. Should be provided to any other endpoints accepting tagId to reference this particular tag definition.`,
						},
						"is_cloud_tag": schema.BoolAttribute{
							Computed:    true,
							Description: `Set to true if tag is mirror of cloud provider resource tag (AWS / AZURE / GCP / OCI)`,
						},
						"key": schema.StringAttribute{
							Computed:    true,
							Description: `Tag Name, human readable name for this tag e.g Environment`,
						},
						"value": schema.StringAttribute{
							Computed:    true,
							Description: `Tag Value, human readable value for this tag e.g Development`,
						},
					},
				},
			},
			"cluster_identifier": schema.StringAttribute{
				Computed: true,
			},
			"container_namespace": schema.StringAttribute{
				Computed: true,
			},
			"core_tags": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"cpu_core_count": schema.Int64Attribute{
				Computed: true,
			},
			"cpu_name": schema.StringAttribute{
				Computed: true,
			},
			"current_traffic_configuration": schema.StringAttribute{
				Computed:    true,
				Description: `must be one of ["disabled", "enable-all", "enable-inbound-only", "enable-outbound-only"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"disabled",
						"enable-all",
						"enable-inbound-only",
						"enable-outbound-only",
					),
				},
			},
			"deterministic_id": schema.StringAttribute{
				Computed: true,
			},
			"disk_capacity_in_gb": schema.Int64Attribute{
				Computed: true,
			},
			"host_name": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"inbound_asset_status": schema.StringAttribute{
				Computed: true,
			},
			"inbound_internet_paths": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"allowed": schema.Int64Attribute{
						Computed: true,
					},
					"allowed_ports": schema.Int64Attribute{
						Computed: true,
					},
					"reviewed": schema.Int64Attribute{
						Computed: true,
					},
					"total": schema.Int64Attribute{
						Computed: true,
					},
					"unreviewed": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
			"inbound_internet_ports": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"allowed": schema.Int64Attribute{
						Computed: true,
					},
					"allowed_ports": schema.Int64Attribute{
						Computed: true,
					},
					"reviewed": schema.Int64Attribute{
						Computed: true,
					},
					"total": schema.Int64Attribute{
						Computed: true,
					},
					"unreviewed": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
			"inbound_intranet_paths": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"allowed": schema.Int64Attribute{
						Computed: true,
					},
					"allowed_ports": schema.Int64Attribute{
						Computed: true,
					},
					"reviewed": schema.Int64Attribute{
						Computed: true,
					},
					"total": schema.Int64Attribute{
						Computed: true,
					},
					"unreviewed": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
			"inbound_intranet_ports": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"allowed": schema.Int64Attribute{
						Computed: true,
					},
					"allowed_ports": schema.Int64Attribute{
						Computed: true,
					},
					"reviewed": schema.Int64Attribute{
						Computed: true,
					},
					"total": schema.Int64Attribute{
						Computed: true,
					},
					"unreviewed": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
			"interfaces": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"flags": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
						"ipaddresses": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
						"macaddress": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"kernel_architecture": schema.StringAttribute{
				Computed: true,
			},
			"kernel_version": schema.StringAttribute{
				Computed: true,
			},
			"lowest_inbound_asset_status": schema.StringAttribute{
				Computed: true,
			},
			"lowest_outbound_asset_status": schema.StringAttribute{
				Computed: true,
			},
			"lowest_progressive_inbound_asset_status": schema.StringAttribute{
				Computed: true,
			},
			"named_network_changes": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"named_network_id": schema.StringAttribute{
							Computed: true,
						},
						"named_network_name": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"namednetworks_assigned": schema.Int64Attribute{
				Computed: true,
			},
			"os_name": schema.StringAttribute{
				Computed: true,
			},
			"outbound_asset_status": schema.StringAttribute{
				Computed: true,
			},
			"outbound_internet_paths": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"allowed": schema.Int64Attribute{
						Computed: true,
					},
					"allowed_ports": schema.Int64Attribute{
						Computed: true,
					},
					"reviewed": schema.Int64Attribute{
						Computed: true,
					},
					"total": schema.Int64Attribute{
						Computed: true,
					},
					"unreviewed": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
			"outbound_intranet_paths": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"allowed": schema.Int64Attribute{
						Computed: true,
					},
					"allowed_ports": schema.Int64Attribute{
						Computed: true,
					},
					"reviewed": schema.Int64Attribute{
						Computed: true,
					},
					"total": schema.Int64Attribute{
						Computed: true,
					},
					"unreviewed": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
			"pending_attack_surface_changes": schema.BoolAttribute{
				Computed: true,
			},
			"pending_blast_radius_changes": schema.BoolAttribute{
				Computed: true,
			},
			"platform": schema.StringAttribute{
				Computed: true,
			},
			"policies_assigned": schema.Int64Attribute{
				Computed: true,
			},
			"policy_status": schema.StringAttribute{
				Computed: true,
			},
			"ram_capacity_in_mb": schema.Int64Attribute{
				Computed: true,
			},
			"security_patches": schema.Int64Attribute{
				Computed: true,
			},
			"tags": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Description: `Unique identifier for this tag instance. Should be provided to any other endpoints accepting tagId to reference this particular tag definition.`,
						},
						"is_cloud_tag": schema.BoolAttribute{
							Computed:    true,
							Description: `Set to true if tag is mirror of cloud provider resource tag (AWS / AZURE / GCP / OCI)`,
						},
						"key": schema.StringAttribute{
							Computed:    true,
							Description: `Tag Name, human readable name for this tag e.g Environment`,
						},
						"value": schema.StringAttribute{
							Computed:    true,
							Description: `Tag Value, human readable value for this tag e.g Development`,
						},
					},
				},
			},
			"template_changes": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"template_id": schema.StringAttribute{
							Computed: true,
						},
						"template_name": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"templates_assigned": schema.Int64Attribute{
				Computed: true,
			},
			"total_comments": schema.Int64Attribute{
				Computed: true,
			},
			"total_inbound_comments": schema.Int64Attribute{
				Computed: true,
			},
			"total_outbound_comments": schema.Int64Attribute{
				Computed: true,
			},
			"total_paths": schema.Int64Attribute{
				Computed: true,
			},
			"total_ports": schema.Int64Attribute{
				Computed: true,
			},
			"total_ports_path_restricted": schema.Int64Attribute{
				Computed: true,
			},
			"type": schema.StringAttribute{
				Required: true,
			},
			"unreviewed_paths": schema.Int64Attribute{
				Computed: true,
			},
			"unreviewed_ports": schema.Int64Attribute{
				Computed: true,
			},
			"usergroup_outbound_internet_paths": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"allowed": schema.Int64Attribute{
						Computed: true,
					},
					"allowed_ports": schema.Int64Attribute{
						Computed: true,
					},
					"reviewed": schema.Int64Attribute{
						Computed: true,
					},
					"total": schema.Int64Attribute{
						Computed: true,
					},
					"unreviewed": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
			"usergroup_outbound_intranet_paths": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"allowed": schema.Int64Attribute{
						Computed: true,
					},
					"allowed_ports": schema.Int64Attribute{
						Computed: true,
					},
					"reviewed": schema.Int64Attribute{
						Computed: true,
					},
					"total": schema.Int64Attribute{
						Computed: true,
					},
					"unreviewed": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
			"usergroup_total_paths": schema.Int64Attribute{
				Computed: true,
			},
			"usergroup_unreviewed_paths": schema.Int64Attribute{
				Computed: true,
			},
			"usergroups": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"groupid": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"users": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"assetid": schema.StringAttribute{
							Computed: true,
						},
						"domainname": schema.StringAttribute{
							Computed:    true,
							Description: `Domain Name the user belongs to`,
						},
						"email": schema.StringAttribute{
							Computed: true,
						},
						"logincount": schema.Int64Attribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: `Unique identifier for this user instance. Should be provided to any other endpoints accepting user to reference this particular user details.`,
						},
						"scimuserid": schema.StringAttribute{
							Computed: true,
						},
						"signedin": schema.BoolAttribute{
							Computed: true,
						},
					},
				},
			},
			"vendor_info": schema.StringAttribute{
				Computed: true,
			},
			"virtualization_system": schema.StringAttribute{
				Computed: true,
			},
			"vulnerabilities": schema.Int64Attribute{
				Computed: true,
			},
		},
	}
}

func (r *AssetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Xshield)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.Xshield, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *AssetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *AssetResourceModel
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

	var assetID string
	assetID = data.ID.ValueString()

	request := operations.GetAssetRequest{
		AssetID: assetID,
	}
	res, err := r.client.Assets.GetAsset(ctx, request)
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
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.AssetDetails != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAssetDetails(res.AssetDetails)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *AssetResourceModel
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

	var assetID string
	assetID = data.ID.ValueString()

	request := operations.GetAssetRequest{
		AssetID: assetID,
	}
	res, err := r.client.Assets.GetAsset(ctx, request)
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
	if !(res.AssetDetails != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAssetDetails(res.AssetDetails)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *AssetResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var assetID string
	assetID = data.ID.ValueString()

	createAssetDetails := *data.ToSharedCreateAssetDetails()
	request := operations.UpdateAssetRequest{
		AssetID:            assetID,
		CreateAssetDetails: createAssetDetails,
	}
	res, err := r.client.Assets.UpdateAsset(ctx, request)
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
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.AssetDetails != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAssetDetails(res.AssetDetails)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *AssetResourceModel
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

	// Not Implemented; entity does not have a configured DELETE operation
}

func (r *AssetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
