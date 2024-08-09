// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	speakeasy_boolplanmodifier "github.com/colortokens/terraform-provider-xshield/internal/planmodifiers/boolplanmodifier"
	speakeasy_listplanmodifier "github.com/colortokens/terraform-provider-xshield/internal/planmodifiers/listplanmodifier"
	speakeasy_objectplanmodifier "github.com/colortokens/terraform-provider-xshield/internal/planmodifiers/objectplanmodifier"
	speakeasy_stringplanmodifier "github.com/colortokens/terraform-provider-xshield/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/colortokens/terraform-provider-xshield/internal/provider/types"
	"github.com/colortokens/terraform-provider-xshield/internal/sdk"
	"github.com/colortokens/terraform-provider-xshield/internal/sdk/models/operations"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &TemplateResource{}
var _ resource.ResourceWithImportState = &TemplateResource{}

func NewTemplateResource() resource.Resource {
	return &TemplateResource{}
}

// TemplateResource defines the resource implementation.
type TemplateResource struct {
	client *sdk.Xshield
}

// TemplateResourceModel describes the resource data model.
type TemplateResourceModel struct {
	AccessPolicyTemplate types.Bool             `tfsdk:"access_policy_template"`
	ColortokensManaged   types.Bool             `tfsdk:"colortokens_managed"`
	ID                   types.String           `tfsdk:"id"`
	TemplateCategory     types.String           `tfsdk:"template_category"`
	TemplateDescription  types.String           `tfsdk:"template_description"`
	TemplateName         types.String           `tfsdk:"template_name"`
	TemplatePaths        []tfTypes.MetadataPath `tfsdk:"template_paths"`
	TemplatePorts        []tfTypes.MetadataPort `tfsdk:"template_ports"`
	TemplateType         types.String           `tfsdk:"template_type"`
}

func (r *TemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_template"
}

func (r *TemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Template Resource",
		Attributes: map[string]schema.Attribute{
			"access_policy_template": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					speakeasy_boolplanmodifier.SuppressDiff(speakeasy_boolplanmodifier.ExplicitSuppress),
				},
			},
			"colortokens_managed": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					speakeasy_boolplanmodifier.SuppressDiff(speakeasy_boolplanmodifier.ExplicitSuppress),
				},
			},
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
			},
			"template_category": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Optional:    true,
				Description: `Requires replacement if changed. `,
			},
			"template_description": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Optional:    true,
				Description: `Requires replacement if changed. `,
			},
			"template_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Optional:    true,
				Description: `Requires replacement if changed. `,
			},
			"template_paths": schema.ListNestedAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
				},
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"destination_asset_id": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Optional:    true,
							Description: `Requires replacement if changed. `,
						},
						"destination_named_network": schema.SingleNestedAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
							},
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"named_network_id": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplaceIfConfigured(),
										speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
									},
									Optional:    true,
									Description: `Requires replacement if changed. `,
								},
								"named_network_name": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplaceIfConfigured(),
										speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
									},
									Optional:    true,
									Description: `Requires replacement if changed. `,
								},
							},
							Description: `Requires replacement if changed. `,
						},
						"destination_tag_based_policy": schema.SingleNestedAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
							},
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"criteria": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplaceIfConfigured(),
										speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
									},
									Optional:    true,
									Description: `Requires replacement if changed. `,
								},
								"criteria_as_params": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplaceIfConfigured(),
										speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
									},
									Optional:    true,
									Description: `Requires replacement if changed. `,
								},
								"tag_based_policy_id": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplaceIfConfigured(),
										speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
									},
									Optional:    true,
									Description: `Requires replacement if changed. `,
								},
								"tag_based_policy_name": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplaceIfConfigured(),
										speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
									},
									Optional:    true,
									Description: `Requires replacement if changed. `,
								},
							},
							Description: `Requires replacement if changed. `,
						},
						"direction": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Optional:    true,
							Description: `Requires replacement if changed. `,
						},
						"domain": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Optional:    true,
							Description: `Requires replacement if changed. `,
						},
						"dst_ip": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Optional:    true,
							Description: `Requires replacement if changed. `,
						},
						"dst_process": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Optional:    true,
							Description: `Requires replacement if changed. `,
						},
						"id": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
						},
						"method": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Optional:    true,
							Description: `Requires replacement if changed. `,
						},
						"port": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Optional:    true,
							Description: `Requires replacement if changed. `,
						},
						"port_name": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Optional:    true,
							Description: `Requires replacement if changed. `,
						},
						"protocol": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Optional:    true,
							Description: `Requires replacement if changed. `,
						},
						"source_asset_id": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Optional:    true,
							Description: `Requires replacement if changed. `,
						},
						"source_named_network": schema.SingleNestedAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
							},
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"named_network_id": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplaceIfConfigured(),
										speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
									},
									Optional:    true,
									Description: `Requires replacement if changed. `,
								},
								"named_network_name": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplaceIfConfigured(),
										speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
									},
									Optional:    true,
									Description: `Requires replacement if changed. `,
								},
							},
							Description: `Requires replacement if changed. `,
						},
						"source_tag_based_policy": schema.SingleNestedAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
							},
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"criteria": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplaceIfConfigured(),
										speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
									},
									Optional:    true,
									Description: `Requires replacement if changed. `,
								},
								"criteria_as_params": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplaceIfConfigured(),
										speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
									},
									Optional:    true,
									Description: `Requires replacement if changed. `,
								},
								"tag_based_policy_id": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplaceIfConfigured(),
										speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
									},
									Optional:    true,
									Description: `Requires replacement if changed. `,
								},
								"tag_based_policy_name": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplaceIfConfigured(),
										speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
									},
									Optional:    true,
									Description: `Requires replacement if changed. `,
								},
							},
							Description: `Requires replacement if changed. `,
						},
						"src_ip": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Optional:    true,
							Description: `Requires replacement if changed. `,
						},
						"src_process": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Optional:    true,
							Description: `Requires replacement if changed. `,
						},
						"uri": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Optional:    true,
							Description: `Requires replacement if changed. `,
						},
					},
				},
				Description: `Requires replacement if changed. `,
			},
			"template_ports": schema.ListNestedAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
				},
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
						},
						"listen_port": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Optional:    true,
							Description: `Requires replacement if changed. `,
						},
						"listen_port_name": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Optional:    true,
							Description: `Requires replacement if changed. `,
						},
						"listen_port_protocol": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Optional:    true,
							Description: `Requires replacement if changed. `,
						},
						"listen_port_reviewed": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Optional:    true,
							Description: `Requires replacement if changed. ; must be one of ["denied", "allow-intranet", "allow-any", "path-restricted"]`,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"denied",
									"allow-intranet",
									"allow-any",
									"path-restricted",
								),
							},
						},
						"listen_process_names": schema.ListAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.List{
								listplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
							},
							Optional:    true,
							ElementType: types.StringType,
							Description: `Requires replacement if changed. `,
						},
					},
				},
				Description: `Requires replacement if changed. `,
			},
			"template_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Optional:    true,
				Description: `Requires replacement if changed. ; must be one of ["application-template", "block-template"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"application-template",
						"block-template",
					),
				},
			},
		},
	}
}

func (r *TemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *TemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *TemplateResourceModel
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

	request := *data.ToSharedTemplate()
	res, err := r.client.Templates.CreateTemplate(ctx, request)
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
	if !(res.Template != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedTemplate(res.Template)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *TemplateResourceModel
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

	var templateid string
	templateid = data.ID.ValueString()

	request := operations.GetTemplateRequest{
		Templateid: templateid,
	}
	res, err := r.client.Templates.GetTemplate(ctx, request)
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
	if !(res.Template != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedTemplate(res.Template)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *TemplateResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *TemplateResourceModel
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

	var templateid string
	templateid = data.ID.ValueString()

	request := operations.DeleteTemplateRequest{
		Templateid: templateid,
	}
	res, err := r.client.Templates.DeleteTemplate(ctx, request)
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

}

func (r *TemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
