// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/speakeasy/terraform-provider-xshield-sdk/internal/provider/types"
	"github.com/speakeasy/terraform-provider-xshield-sdk/internal/sdk/models/shared"
)

func (r *NamedNetworkResourceModel) ToSharedNamednetworkNamedNetwork() *shared.NamednetworkNamedNetwork {
	id := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueString()
	} else {
		id = nil
	}
	assignedByTagBasedPolicy := new(bool)
	if !r.AssignedByTagBasedPolicy.IsUnknown() && !r.AssignedByTagBasedPolicy.IsNull() {
		*assignedByTagBasedPolicy = r.AssignedByTagBasedPolicy.ValueBool()
	} else {
		assignedByTagBasedPolicy = nil
	}
	var ipRanges []shared.NamednetworkRange = []shared.NamednetworkRange{}
	for _, ipRangesItem := range r.IPRanges {
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
	if !r.IsOOBNetwork.IsUnknown() && !r.IsOOBNetwork.IsNull() {
		*isOOBNetwork = r.IsOOBNetwork.ValueBool()
	} else {
		isOOBNetwork = nil
	}
	namedNetworkAssignments := new(int64)
	if !r.NamedNetworkAssignments.IsUnknown() && !r.NamedNetworkAssignments.IsNull() {
		*namedNetworkAssignments = r.NamedNetworkAssignments.ValueInt64()
	} else {
		namedNetworkAssignments = nil
	}
	namedNetworkDescription := new(string)
	if !r.NamedNetworkDescription.IsUnknown() && !r.NamedNetworkDescription.IsNull() {
		*namedNetworkDescription = r.NamedNetworkDescription.ValueString()
	} else {
		namedNetworkDescription = nil
	}
	namedNetworkName := new(string)
	if !r.NamedNetworkName.IsUnknown() && !r.NamedNetworkName.IsNull() {
		*namedNetworkName = r.NamedNetworkName.ValueString()
	} else {
		namedNetworkName = nil
	}
	namednetworkTagBasedPolicyAssignments := new(int64)
	if !r.NamednetworkTagBasedPolicyAssignments.IsUnknown() && !r.NamednetworkTagBasedPolicyAssignments.IsNull() {
		*namednetworkTagBasedPolicyAssignments = r.NamednetworkTagBasedPolicyAssignments.ValueInt64()
	} else {
		namednetworkTagBasedPolicyAssignments = nil
	}
	programAsInternet := new(bool)
	if !r.ProgramAsInternet.IsUnknown() && !r.ProgramAsInternet.IsNull() {
		*programAsInternet = r.ProgramAsInternet.ValueBool()
	} else {
		programAsInternet = nil
	}
	programAsIntranet := new(bool)
	if !r.ProgramAsIntranet.IsUnknown() && !r.ProgramAsIntranet.IsNull() {
		*programAsIntranet = r.ProgramAsIntranet.ValueBool()
	} else {
		programAsIntranet = nil
	}
	region := new(string)
	if !r.Region.IsUnknown() && !r.Region.IsNull() {
		*region = r.Region.ValueString()
	} else {
		region = nil
	}
	service := new(string)
	if !r.Service.IsUnknown() && !r.Service.IsNull() {
		*service = r.Service.ValueString()
	} else {
		service = nil
	}
	totalComments := new(int64)
	if !r.TotalComments.IsUnknown() && !r.TotalComments.IsNull() {
		*totalComments = r.TotalComments.ValueInt64()
	} else {
		totalComments = nil
	}
	totalCount := new(int64)
	if !r.TotalCount.IsUnknown() && !r.TotalCount.IsNull() {
		*totalCount = r.TotalCount.ValueInt64()
	} else {
		totalCount = nil
	}
	usergroupNamedNetworkAssignments := new(int64)
	if !r.UsergroupNamedNetworkAssignments.IsUnknown() && !r.UsergroupNamedNetworkAssignments.IsNull() {
		*usergroupNamedNetworkAssignments = r.UsergroupNamedNetworkAssignments.ValueInt64()
	} else {
		usergroupNamedNetworkAssignments = nil
	}
	out := shared.NamednetworkNamedNetwork{
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
	}
	return &out
}

func (r *NamedNetworkResourceModel) RefreshFromSharedNamednetworkNamedNetwork(resp *shared.NamednetworkNamedNetwork) {
	if resp != nil {
		r.AssignedByTagBasedPolicy = types.BoolPointerValue(resp.AssignedByTagBasedPolicy)
		r.ID = types.StringPointerValue(resp.ID)
		r.IPRanges = []tfTypes.NamednetworkRange{}
		if len(r.IPRanges) > len(resp.IPRanges) {
			r.IPRanges = r.IPRanges[:len(resp.IPRanges)]
		}
		for ipRangesCount, ipRangesItem := range resp.IPRanges {
			var ipRanges1 tfTypes.NamednetworkRange
			ipRanges1.ID = types.StringPointerValue(ipRangesItem.ID)
			ipRanges1.IPCount = types.Int64PointerValue(ipRangesItem.IPCount)
			ipRanges1.IPRange = types.StringPointerValue(ipRangesItem.IPRange)
			if ipRangesCount+1 > len(r.IPRanges) {
				r.IPRanges = append(r.IPRanges, ipRanges1)
			} else {
				r.IPRanges[ipRangesCount].ID = ipRanges1.ID
				r.IPRanges[ipRangesCount].IPCount = ipRanges1.IPCount
				r.IPRanges[ipRangesCount].IPRange = ipRanges1.IPRange
			}
		}
		r.IsOOBNetwork = types.BoolPointerValue(resp.IsOOBNetwork)
		r.NamedNetworkAssignments = types.Int64PointerValue(resp.NamedNetworkAssignments)
		r.NamedNetworkDescription = types.StringPointerValue(resp.NamedNetworkDescription)
		r.NamedNetworkName = types.StringPointerValue(resp.NamedNetworkName)
		r.NamednetworkTagBasedPolicyAssignments = types.Int64PointerValue(resp.NamednetworkTagBasedPolicyAssignments)
		r.ProgramAsInternet = types.BoolPointerValue(resp.ProgramAsInternet)
		r.ProgramAsIntranet = types.BoolPointerValue(resp.ProgramAsIntranet)
		r.Region = types.StringPointerValue(resp.Region)
		r.Service = types.StringPointerValue(resp.Service)
		r.TotalComments = types.Int64PointerValue(resp.TotalComments)
		r.TotalCount = types.Int64PointerValue(resp.TotalCount)
		r.UsergroupNamedNetworkAssignments = types.Int64PointerValue(resp.UsergroupNamedNetworkAssignments)
	}
}
