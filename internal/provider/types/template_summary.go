// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type TemplateSummary struct {
	AccessPolicyTemplate                            types.Bool   `tfsdk:"access_policy_template"`
	AssignedByTagBasedPolicy                        types.Bool   `tfsdk:"assigned_by_tag_based_policy"`
	OobTemplate                                     types.Bool   `tfsdk:"oob_template"`
	TemplateAssignments                             types.Int64  `tfsdk:"template_assignments"`
	TemplateCategory                                types.String `tfsdk:"template_category"`
	TemplateDescription                             types.String `tfsdk:"template_description"`
	TemplateID                                      types.String `tfsdk:"template_id"`
	TemplateName                                    types.String `tfsdk:"template_name"`
	TemplatePaths                                   types.Int64  `tfsdk:"template_paths"`
	TemplatePorts                                   types.Int64  `tfsdk:"template_ports"`
	TemplateTagBasedPolicyAssignments               types.Int64  `tfsdk:"template_tag_based_policy_assignments"`
	TemplateType                                    types.String `tfsdk:"template_type"`
	TemplateUnassignmentsPendingFirewallSynchronize types.Int64  `tfsdk:"template_unassignments_pending_firewall_synchronize"`
	TotalComments                                   types.Int64  `tfsdk:"total_comments"`
	UsergroupTemplateAssignments                    types.Int64  `tfsdk:"usergroup_template_assignments"`
}
