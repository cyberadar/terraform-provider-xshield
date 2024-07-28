// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/speakeasy/terraform-provider-xshield-sdk/internal/provider/types"
	"github.com/speakeasy/terraform-provider-xshield-sdk/internal/sdk/models/shared"
)

func (r *NamedNetworkDataSourceModel) RefreshFromSharedNamednetworkNamedNetwork(resp *shared.NamednetworkNamedNetwork) {
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
