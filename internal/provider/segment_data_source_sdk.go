// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/speakeasy/terraform-provider-xshield-sdk/internal/provider/types"
	"github.com/speakeasy/terraform-provider-xshield-sdk/internal/sdk/models/shared"
	"time"
)

func (r *SegmentDataSourceModel) RefreshFromSharedTagBasedPolicy(resp *shared.TagBasedPolicy) {
	if resp != nil {
		r.AutoSynchronizeEnabled = types.BoolPointerValue(resp.AutoSynchronizeEnabled)
		r.Criteria = types.StringPointerValue(resp.Criteria)
		r.Description = types.StringPointerValue(resp.Description)
		r.ID = types.StringPointerValue(resp.ID)
		r.LowestInboundPolicyStatus = types.StringPointerValue(resp.LowestInboundPolicyStatus)
		r.LowestOutboundPolicyStatus = types.StringPointerValue(resp.LowestOutboundPolicyStatus)
		r.LowestProgressiveInboundPolicyStatus = types.StringPointerValue(resp.LowestProgressiveInboundPolicyStatus)
		r.MatchingAssets = types.Int64PointerValue(resp.MatchingAssets)
		r.Namednetworks = []tfTypes.NamednetworkNamedNetwork{}
		if len(r.Namednetworks) > len(resp.Namednetworks) {
			r.Namednetworks = r.Namednetworks[:len(resp.Namednetworks)]
		}
		for namednetworksCount, namednetworksItem := range resp.Namednetworks {
			var namednetworks1 tfTypes.NamednetworkNamedNetwork
			namednetworks1.AssignedByTagBasedPolicy = types.BoolPointerValue(namednetworksItem.AssignedByTagBasedPolicy)
			namednetworks1.ColortokensManaged = types.BoolPointerValue(namednetworksItem.ColortokensManaged)
			namednetworks1.ID = types.StringPointerValue(namednetworksItem.ID)
			namednetworks1.IPRanges = []tfTypes.NamednetworkRange{}
			for ipRangesCount, ipRangesItem := range namednetworksItem.IPRanges {
				var ipRanges1 tfTypes.NamednetworkRange
				ipRanges1.ID = types.StringPointerValue(ipRangesItem.ID)
				ipRanges1.IPCount = types.Int64PointerValue(ipRangesItem.IPCount)
				ipRanges1.IPRange = types.StringPointerValue(ipRangesItem.IPRange)
				if ipRangesCount+1 > len(namednetworks1.IPRanges) {
					namednetworks1.IPRanges = append(namednetworks1.IPRanges, ipRanges1)
				} else {
					namednetworks1.IPRanges[ipRangesCount].ID = ipRanges1.ID
					namednetworks1.IPRanges[ipRangesCount].IPCount = ipRanges1.IPCount
					namednetworks1.IPRanges[ipRangesCount].IPRange = ipRanges1.IPRange
				}
			}
			namednetworks1.NamedNetworkAssignments = types.Int64PointerValue(namednetworksItem.NamedNetworkAssignments)
			namednetworks1.NamedNetworkDescription = types.StringPointerValue(namednetworksItem.NamedNetworkDescription)
			namednetworks1.NamedNetworkName = types.StringPointerValue(namednetworksItem.NamedNetworkName)
			namednetworks1.NamednetworkTagBasedPolicyAssignments = types.Int64PointerValue(namednetworksItem.NamednetworkTagBasedPolicyAssignments)
			namednetworks1.ProgramAsInternet = types.BoolPointerValue(namednetworksItem.ProgramAsInternet)
			namednetworks1.ProgramAsIntranet = types.BoolPointerValue(namednetworksItem.ProgramAsIntranet)
			namednetworks1.Region = types.StringPointerValue(namednetworksItem.Region)
			namednetworks1.Service = types.StringPointerValue(namednetworksItem.Service)
			namednetworks1.TotalComments = types.Int64PointerValue(namednetworksItem.TotalComments)
			namednetworks1.TotalCount = types.Int64PointerValue(namednetworksItem.TotalCount)
			namednetworks1.UsergroupNamedNetworkAssignments = types.Int64PointerValue(namednetworksItem.UsergroupNamedNetworkAssignments)
			if namednetworksCount+1 > len(r.Namednetworks) {
				r.Namednetworks = append(r.Namednetworks, namednetworks1)
			} else {
				r.Namednetworks[namednetworksCount].AssignedByTagBasedPolicy = namednetworks1.AssignedByTagBasedPolicy
				r.Namednetworks[namednetworksCount].ColortokensManaged = namednetworks1.ColortokensManaged
				r.Namednetworks[namednetworksCount].ID = namednetworks1.ID
				r.Namednetworks[namednetworksCount].IPRanges = namednetworks1.IPRanges
				r.Namednetworks[namednetworksCount].NamedNetworkAssignments = namednetworks1.NamedNetworkAssignments
				r.Namednetworks[namednetworksCount].NamedNetworkDescription = namednetworks1.NamedNetworkDescription
				r.Namednetworks[namednetworksCount].NamedNetworkName = namednetworks1.NamedNetworkName
				r.Namednetworks[namednetworksCount].NamednetworkTagBasedPolicyAssignments = namednetworks1.NamednetworkTagBasedPolicyAssignments
				r.Namednetworks[namednetworksCount].ProgramAsInternet = namednetworks1.ProgramAsInternet
				r.Namednetworks[namednetworksCount].ProgramAsIntranet = namednetworks1.ProgramAsIntranet
				r.Namednetworks[namednetworksCount].Region = namednetworks1.Region
				r.Namednetworks[namednetworksCount].Service = namednetworks1.Service
				r.Namednetworks[namednetworksCount].TotalComments = namednetworks1.TotalComments
				r.Namednetworks[namednetworksCount].TotalCount = namednetworks1.TotalCount
				r.Namednetworks[namednetworksCount].UsergroupNamedNetworkAssignments = namednetworks1.UsergroupNamedNetworkAssignments
			}
		}
		r.PolicyAutomationConfigurable = types.BoolPointerValue(resp.PolicyAutomationConfigurable)
		if resp.PolicyProgressiveLastRefreshed != nil {
			r.PolicyProgressiveLastRefreshed = types.StringValue(resp.PolicyProgressiveLastRefreshed.Format(time.RFC3339Nano))
		} else {
			r.PolicyProgressiveLastRefreshed = types.StringNull()
		}
		r.TagBasedPolicyName = types.StringPointerValue(resp.TagBasedPolicyName)
		r.Templates = []tfTypes.TemplateSummary{}
		if len(r.Templates) > len(resp.Templates) {
			r.Templates = r.Templates[:len(resp.Templates)]
		}
		for templatesCount, templatesItem := range resp.Templates {
			var templates1 tfTypes.TemplateSummary
			templates1.AccessPolicyTemplate = types.BoolPointerValue(templatesItem.AccessPolicyTemplate)
			templates1.AssignedByTagBasedPolicy = types.BoolPointerValue(templatesItem.AssignedByTagBasedPolicy)
			templates1.OobTemplate = types.BoolPointerValue(templatesItem.OobTemplate)
			templates1.TemplateAssignments = types.Int64PointerValue(templatesItem.TemplateAssignments)
			templates1.TemplateCategory = types.StringPointerValue(templatesItem.TemplateCategory)
			templates1.TemplateDescription = types.StringPointerValue(templatesItem.TemplateDescription)
			templates1.TemplateID = types.StringPointerValue(templatesItem.TemplateID)
			templates1.TemplateName = types.StringPointerValue(templatesItem.TemplateName)
			templates1.TemplatePaths = types.Int64PointerValue(templatesItem.TemplatePaths)
			templates1.TemplatePorts = types.Int64PointerValue(templatesItem.TemplatePorts)
			templates1.TemplateTagBasedPolicyAssignments = types.Int64PointerValue(templatesItem.TemplateTagBasedPolicyAssignments)
			templates1.TemplateType = types.StringValue(templatesItem.TemplateType)
			templates1.TemplateUnassignmentsPendingFirewallSynchronize = types.Int64PointerValue(templatesItem.TemplateUnassignmentsPendingFirewallSynchronize)
			templates1.TotalComments = types.Int64PointerValue(templatesItem.TotalComments)
			templates1.UsergroupTemplateAssignments = types.Int64PointerValue(templatesItem.UsergroupTemplateAssignments)
			if templatesCount+1 > len(r.Templates) {
				r.Templates = append(r.Templates, templates1)
			} else {
				r.Templates[templatesCount].AccessPolicyTemplate = templates1.AccessPolicyTemplate
				r.Templates[templatesCount].AssignedByTagBasedPolicy = templates1.AssignedByTagBasedPolicy
				r.Templates[templatesCount].OobTemplate = templates1.OobTemplate
				r.Templates[templatesCount].TemplateAssignments = templates1.TemplateAssignments
				r.Templates[templatesCount].TemplateCategory = templates1.TemplateCategory
				r.Templates[templatesCount].TemplateDescription = templates1.TemplateDescription
				r.Templates[templatesCount].TemplateID = templates1.TemplateID
				r.Templates[templatesCount].TemplateName = templates1.TemplateName
				r.Templates[templatesCount].TemplatePaths = templates1.TemplatePaths
				r.Templates[templatesCount].TemplatePorts = templates1.TemplatePorts
				r.Templates[templatesCount].TemplateTagBasedPolicyAssignments = templates1.TemplateTagBasedPolicyAssignments
				r.Templates[templatesCount].TemplateType = templates1.TemplateType
				r.Templates[templatesCount].TemplateUnassignmentsPendingFirewallSynchronize = templates1.TemplateUnassignmentsPendingFirewallSynchronize
				r.Templates[templatesCount].TotalComments = templates1.TotalComments
				r.Templates[templatesCount].UsergroupTemplateAssignments = templates1.UsergroupTemplateAssignments
			}
		}
	}
}
