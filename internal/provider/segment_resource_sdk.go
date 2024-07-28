// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/speakeasy/terraform-provider-xshield-sdk/internal/provider/types"
	"github.com/speakeasy/terraform-provider-xshield-sdk/internal/sdk/models/shared"
	"time"
)

func (r *SegmentResourceModel) ToSharedTagBasedPolicy() *shared.TagBasedPolicy {
	autoSynchronizeEnabled := new(bool)
	if !r.AutoSynchronizeEnabled.IsUnknown() && !r.AutoSynchronizeEnabled.IsNull() {
		*autoSynchronizeEnabled = r.AutoSynchronizeEnabled.ValueBool()
	} else {
		autoSynchronizeEnabled = nil
	}
	criteria := new(string)
	if !r.Criteria.IsUnknown() && !r.Criteria.IsNull() {
		*criteria = r.Criteria.ValueString()
	} else {
		criteria = nil
	}
	criteriaAsParams := new(string)
	if !r.CriteriaAsParams.IsUnknown() && !r.CriteriaAsParams.IsNull() {
		*criteriaAsParams = r.CriteriaAsParams.ValueString()
	} else {
		criteriaAsParams = nil
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	lowestInboundPolicyStatus := new(string)
	if !r.LowestInboundPolicyStatus.IsUnknown() && !r.LowestInboundPolicyStatus.IsNull() {
		*lowestInboundPolicyStatus = r.LowestInboundPolicyStatus.ValueString()
	} else {
		lowestInboundPolicyStatus = nil
	}
	lowestOutboundPolicyStatus := new(string)
	if !r.LowestOutboundPolicyStatus.IsUnknown() && !r.LowestOutboundPolicyStatus.IsNull() {
		*lowestOutboundPolicyStatus = r.LowestOutboundPolicyStatus.ValueString()
	} else {
		lowestOutboundPolicyStatus = nil
	}
	lowestProgressiveInboundPolicyStatus := new(string)
	if !r.LowestProgressiveInboundPolicyStatus.IsUnknown() && !r.LowestProgressiveInboundPolicyStatus.IsNull() {
		*lowestProgressiveInboundPolicyStatus = r.LowestProgressiveInboundPolicyStatus.ValueString()
	} else {
		lowestProgressiveInboundPolicyStatus = nil
	}
	matchingAssets := new(int64)
	if !r.MatchingAssets.IsUnknown() && !r.MatchingAssets.IsNull() {
		*matchingAssets = r.MatchingAssets.ValueInt64()
	} else {
		matchingAssets = nil
	}
	var namednetworks []shared.NamednetworkNamedNetwork = []shared.NamednetworkNamedNetwork{}
	for _, namednetworksItem := range r.Namednetworks {
		id := new(string)
		if !namednetworksItem.ID.IsUnknown() && !namednetworksItem.ID.IsNull() {
			*id = namednetworksItem.ID.ValueString()
		} else {
			id = nil
		}
		assignedByTagBasedPolicy := new(bool)
		if !namednetworksItem.AssignedByTagBasedPolicy.IsUnknown() && !namednetworksItem.AssignedByTagBasedPolicy.IsNull() {
			*assignedByTagBasedPolicy = namednetworksItem.AssignedByTagBasedPolicy.ValueBool()
		} else {
			assignedByTagBasedPolicy = nil
		}
		var ipRanges []shared.NamednetworkRange = []shared.NamednetworkRange{}
		for _, ipRangesItem := range namednetworksItem.IPRanges {
			id1 := new(string)
			if !ipRangesItem.ID.IsUnknown() && !ipRangesItem.ID.IsNull() {
				*id1 = ipRangesItem.ID.ValueString()
			} else {
				id1 = nil
			}
			ipCount := new(int64)
			if !ipRangesItem.IPCount.IsUnknown() && !ipRangesItem.IPCount.IsNull() {
				*ipCount = ipRangesItem.IPCount.ValueInt64()
			} else {
				ipCount = nil
			}
			ipRange := new(string)
			if !ipRangesItem.IPRange.IsUnknown() && !ipRangesItem.IPRange.IsNull() {
				*ipRange = ipRangesItem.IPRange.ValueString()
			} else {
				ipRange = nil
			}
			ipRanges = append(ipRanges, shared.NamednetworkRange{
				ID:      id1,
				IPCount: ipCount,
				IPRange: ipRange,
			})
		}
		isOOBNetwork := new(bool)
		if !namednetworksItem.IsOOBNetwork.IsUnknown() && !namednetworksItem.IsOOBNetwork.IsNull() {
			*isOOBNetwork = namednetworksItem.IsOOBNetwork.ValueBool()
		} else {
			isOOBNetwork = nil
		}
		namedNetworkAssignments := new(int64)
		if !namednetworksItem.NamedNetworkAssignments.IsUnknown() && !namednetworksItem.NamedNetworkAssignments.IsNull() {
			*namedNetworkAssignments = namednetworksItem.NamedNetworkAssignments.ValueInt64()
		} else {
			namedNetworkAssignments = nil
		}
		namedNetworkDescription := new(string)
		if !namednetworksItem.NamedNetworkDescription.IsUnknown() && !namednetworksItem.NamedNetworkDescription.IsNull() {
			*namedNetworkDescription = namednetworksItem.NamedNetworkDescription.ValueString()
		} else {
			namedNetworkDescription = nil
		}
		namedNetworkName := new(string)
		if !namednetworksItem.NamedNetworkName.IsUnknown() && !namednetworksItem.NamedNetworkName.IsNull() {
			*namedNetworkName = namednetworksItem.NamedNetworkName.ValueString()
		} else {
			namedNetworkName = nil
		}
		namednetworkTagBasedPolicyAssignments := new(int64)
		if !namednetworksItem.NamednetworkTagBasedPolicyAssignments.IsUnknown() && !namednetworksItem.NamednetworkTagBasedPolicyAssignments.IsNull() {
			*namednetworkTagBasedPolicyAssignments = namednetworksItem.NamednetworkTagBasedPolicyAssignments.ValueInt64()
		} else {
			namednetworkTagBasedPolicyAssignments = nil
		}
		programAsInternet := new(bool)
		if !namednetworksItem.ProgramAsInternet.IsUnknown() && !namednetworksItem.ProgramAsInternet.IsNull() {
			*programAsInternet = namednetworksItem.ProgramAsInternet.ValueBool()
		} else {
			programAsInternet = nil
		}
		programAsIntranet := new(bool)
		if !namednetworksItem.ProgramAsIntranet.IsUnknown() && !namednetworksItem.ProgramAsIntranet.IsNull() {
			*programAsIntranet = namednetworksItem.ProgramAsIntranet.ValueBool()
		} else {
			programAsIntranet = nil
		}
		region := new(string)
		if !namednetworksItem.Region.IsUnknown() && !namednetworksItem.Region.IsNull() {
			*region = namednetworksItem.Region.ValueString()
		} else {
			region = nil
		}
		service := new(string)
		if !namednetworksItem.Service.IsUnknown() && !namednetworksItem.Service.IsNull() {
			*service = namednetworksItem.Service.ValueString()
		} else {
			service = nil
		}
		totalComments := new(int64)
		if !namednetworksItem.TotalComments.IsUnknown() && !namednetworksItem.TotalComments.IsNull() {
			*totalComments = namednetworksItem.TotalComments.ValueInt64()
		} else {
			totalComments = nil
		}
		totalCount := new(int64)
		if !namednetworksItem.TotalCount.IsUnknown() && !namednetworksItem.TotalCount.IsNull() {
			*totalCount = namednetworksItem.TotalCount.ValueInt64()
		} else {
			totalCount = nil
		}
		usergroupNamedNetworkAssignments := new(int64)
		if !namednetworksItem.UsergroupNamedNetworkAssignments.IsUnknown() && !namednetworksItem.UsergroupNamedNetworkAssignments.IsNull() {
			*usergroupNamedNetworkAssignments = namednetworksItem.UsergroupNamedNetworkAssignments.ValueInt64()
		} else {
			usergroupNamedNetworkAssignments = nil
		}
		namednetworks = append(namednetworks, shared.NamednetworkNamedNetwork{
			ID:                                    id,
			AssignedByTagBasedPolicy:              assignedByTagBasedPolicy,
			IPRanges:                              ipRanges,
			IsOOBNetwork:                          isOOBNetwork,
			NamedNetworkAssignments:               namedNetworkAssignments,
			NamedNetworkDescription:               namedNetworkDescription,
			NamedNetworkName:                      namedNetworkName,
			NamednetworkTagBasedPolicyAssignments: namednetworkTagBasedPolicyAssignments,
			ProgramAsInternet:                     programAsInternet,
			ProgramAsIntranet:                     programAsIntranet,
			Region:                                region,
			Service:                               service,
			TotalComments:                         totalComments,
			TotalCount:                            totalCount,
			UsergroupNamedNetworkAssignments:      usergroupNamedNetworkAssignments,
		})
	}
	policyAutomationConfigurable := new(bool)
	if !r.PolicyAutomationConfigurable.IsUnknown() && !r.PolicyAutomationConfigurable.IsNull() {
		*policyAutomationConfigurable = r.PolicyAutomationConfigurable.ValueBool()
	} else {
		policyAutomationConfigurable = nil
	}
	policyProgressiveLastRefreshed := new(time.Time)
	if !r.PolicyProgressiveLastRefreshed.IsUnknown() && !r.PolicyProgressiveLastRefreshed.IsNull() {
		*policyProgressiveLastRefreshed, _ = time.Parse(time.RFC3339Nano, r.PolicyProgressiveLastRefreshed.ValueString())
	} else {
		policyProgressiveLastRefreshed = nil
	}
	id2 := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id2 = r.ID.ValueString()
	} else {
		id2 = nil
	}
	tagBasedPolicyName := new(string)
	if !r.TagBasedPolicyName.IsUnknown() && !r.TagBasedPolicyName.IsNull() {
		*tagBasedPolicyName = r.TagBasedPolicyName.ValueString()
	} else {
		tagBasedPolicyName = nil
	}
	var templates []shared.TemplateSummary = []shared.TemplateSummary{}
	for _, templatesItem := range r.Templates {
		accessPolicyTemplate := new(bool)
		if !templatesItem.AccessPolicyTemplate.IsUnknown() && !templatesItem.AccessPolicyTemplate.IsNull() {
			*accessPolicyTemplate = templatesItem.AccessPolicyTemplate.ValueBool()
		} else {
			accessPolicyTemplate = nil
		}
		assignedByTagBasedPolicy1 := new(bool)
		if !templatesItem.AssignedByTagBasedPolicy.IsUnknown() && !templatesItem.AssignedByTagBasedPolicy.IsNull() {
			*assignedByTagBasedPolicy1 = templatesItem.AssignedByTagBasedPolicy.ValueBool()
		} else {
			assignedByTagBasedPolicy1 = nil
		}
		oobTemplate := new(bool)
		if !templatesItem.OobTemplate.IsUnknown() && !templatesItem.OobTemplate.IsNull() {
			*oobTemplate = templatesItem.OobTemplate.ValueBool()
		} else {
			oobTemplate = nil
		}
		templateAssignments := new(int64)
		if !templatesItem.TemplateAssignments.IsUnknown() && !templatesItem.TemplateAssignments.IsNull() {
			*templateAssignments = templatesItem.TemplateAssignments.ValueInt64()
		} else {
			templateAssignments = nil
		}
		templateCategory := new(string)
		if !templatesItem.TemplateCategory.IsUnknown() && !templatesItem.TemplateCategory.IsNull() {
			*templateCategory = templatesItem.TemplateCategory.ValueString()
		} else {
			templateCategory = nil
		}
		templateDescription := new(string)
		if !templatesItem.TemplateDescription.IsUnknown() && !templatesItem.TemplateDescription.IsNull() {
			*templateDescription = templatesItem.TemplateDescription.ValueString()
		} else {
			templateDescription = nil
		}
		templateID := new(string)
		if !templatesItem.TemplateID.IsUnknown() && !templatesItem.TemplateID.IsNull() {
			*templateID = templatesItem.TemplateID.ValueString()
		} else {
			templateID = nil
		}
		templateName := new(string)
		if !templatesItem.TemplateName.IsUnknown() && !templatesItem.TemplateName.IsNull() {
			*templateName = templatesItem.TemplateName.ValueString()
		} else {
			templateName = nil
		}
		templatePaths := new(int64)
		if !templatesItem.TemplatePaths.IsUnknown() && !templatesItem.TemplatePaths.IsNull() {
			*templatePaths = templatesItem.TemplatePaths.ValueInt64()
		} else {
			templatePaths = nil
		}
		templatePorts := new(int64)
		if !templatesItem.TemplatePorts.IsUnknown() && !templatesItem.TemplatePorts.IsNull() {
			*templatePorts = templatesItem.TemplatePorts.ValueInt64()
		} else {
			templatePorts = nil
		}
		templateTagBasedPolicyAssignments := new(int64)
		if !templatesItem.TemplateTagBasedPolicyAssignments.IsUnknown() && !templatesItem.TemplateTagBasedPolicyAssignments.IsNull() {
			*templateTagBasedPolicyAssignments = templatesItem.TemplateTagBasedPolicyAssignments.ValueInt64()
		} else {
			templateTagBasedPolicyAssignments = nil
		}
		var templateType string
		templateType = templatesItem.TemplateType.ValueString()

		templateUnassignmentsPendingFirewallSynchronize := new(int64)
		if !templatesItem.TemplateUnassignmentsPendingFirewallSynchronize.IsUnknown() && !templatesItem.TemplateUnassignmentsPendingFirewallSynchronize.IsNull() {
			*templateUnassignmentsPendingFirewallSynchronize = templatesItem.TemplateUnassignmentsPendingFirewallSynchronize.ValueInt64()
		} else {
			templateUnassignmentsPendingFirewallSynchronize = nil
		}
		totalComments1 := new(int64)
		if !templatesItem.TotalComments.IsUnknown() && !templatesItem.TotalComments.IsNull() {
			*totalComments1 = templatesItem.TotalComments.ValueInt64()
		} else {
			totalComments1 = nil
		}
		usergroupTemplateAssignments := new(int64)
		if !templatesItem.UsergroupTemplateAssignments.IsUnknown() && !templatesItem.UsergroupTemplateAssignments.IsNull() {
			*usergroupTemplateAssignments = templatesItem.UsergroupTemplateAssignments.ValueInt64()
		} else {
			usergroupTemplateAssignments = nil
		}
		templates = append(templates, shared.TemplateSummary{
			AccessPolicyTemplate:              accessPolicyTemplate,
			AssignedByTagBasedPolicy:          assignedByTagBasedPolicy1,
			OobTemplate:                       oobTemplate,
			TemplateAssignments:               templateAssignments,
			TemplateCategory:                  templateCategory,
			TemplateDescription:               templateDescription,
			TemplateID:                        templateID,
			TemplateName:                      templateName,
			TemplatePaths:                     templatePaths,
			TemplatePorts:                     templatePorts,
			TemplateTagBasedPolicyAssignments: templateTagBasedPolicyAssignments,
			TemplateType:                      templateType,
			TemplateUnassignmentsPendingFirewallSynchronize: templateUnassignmentsPendingFirewallSynchronize,
			TotalComments:                totalComments1,
			UsergroupTemplateAssignments: usergroupTemplateAssignments,
		})
	}
	out := shared.TagBasedPolicy{
		AutoSynchronizeEnabled:               autoSynchronizeEnabled,
		Criteria:                             criteria,
		CriteriaAsParams:                     criteriaAsParams,
		Description:                          description,
		LowestInboundPolicyStatus:            lowestInboundPolicyStatus,
		LowestOutboundPolicyStatus:           lowestOutboundPolicyStatus,
		LowestProgressiveInboundPolicyStatus: lowestProgressiveInboundPolicyStatus,
		MatchingAssets:                       matchingAssets,
		Namednetworks:                        namednetworks,
		PolicyAutomationConfigurable:         policyAutomationConfigurable,
		PolicyProgressiveLastRefreshed:       policyProgressiveLastRefreshed,
		ID:                                   id2,
		TagBasedPolicyName:                   tagBasedPolicyName,
		Templates:                            templates,
	}
	return &out
}

func (r *SegmentResourceModel) RefreshFromSharedTagBasedPolicy(resp *shared.TagBasedPolicy) {
	if resp != nil {
		r.AutoSynchronizeEnabled = types.BoolPointerValue(resp.AutoSynchronizeEnabled)
		r.Criteria = types.StringPointerValue(resp.Criteria)
		r.CriteriaAsParams = types.StringPointerValue(resp.CriteriaAsParams)
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
			namednetworks1.IsOOBNetwork = types.BoolPointerValue(namednetworksItem.IsOOBNetwork)
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
				r.Namednetworks[namednetworksCount].ID = namednetworks1.ID
				r.Namednetworks[namednetworksCount].IPRanges = namednetworks1.IPRanges
				r.Namednetworks[namednetworksCount].IsOOBNetwork = namednetworks1.IsOOBNetwork
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
