// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type NamednetworkNamedNetwork struct {
	ID                                    *string             `json:"id,omitempty"`
	AssignedByTagBasedPolicy              *bool               `json:"assignedByTagBasedPolicy,omitempty"`
	IPRanges                              []NamednetworkRange `json:"ipRanges,omitempty"`
	IsOOBNetwork                          *bool               `json:"isOOBNetwork,omitempty"`
	NamedNetworkAssignments               *int64              `json:"namedNetworkAssignments,omitempty"`
	NamedNetworkDescription               *string             `json:"namedNetworkDescription,omitempty"`
	NamedNetworkName                      *string             `json:"namedNetworkName,omitempty"`
	NamednetworkTagBasedPolicyAssignments *int64              `json:"namednetworkTagBasedPolicyAssignments,omitempty"`
	ProgramAsInternet                     *bool               `json:"programAsInternet,omitempty"`
	ProgramAsIntranet                     *bool               `json:"programAsIntranet,omitempty"`
	Vendor                                *string             `json:"vendor,omitempty"`
	Region                                *string             `json:"region,omitempty"`
	Service                               *string             `json:"service,omitempty"`
	TotalComments                         *int64              `json:"totalComments,omitempty"`
	TotalCount                            *int64              `json:"totalCount,omitempty"`
	UsergroupNamedNetworkAssignments      *int64              `json:"usergroupNamedNetworkAssignments,omitempty"`
}

func (o *NamednetworkNamedNetwork) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *NamednetworkNamedNetwork) GetAssignedByTagBasedPolicy() *bool {
	if o == nil {
		return nil
	}
	return o.AssignedByTagBasedPolicy
}

func (o *NamednetworkNamedNetwork) GetIPRanges() []NamednetworkRange {
	if o == nil {
		return nil
	}
	return o.IPRanges
}

func (o *NamednetworkNamedNetwork) GetIsOOBNetwork() *bool {
	if o == nil {
		return nil
	}
	return o.IsOOBNetwork
}

func (o *NamednetworkNamedNetwork) GetNamedNetworkAssignments() *int64 {
	if o == nil {
		return nil
	}
	return o.NamedNetworkAssignments
}

func (o *NamednetworkNamedNetwork) GetNamedNetworkDescription() *string {
	if o == nil {
		return nil
	}
	return o.NamedNetworkDescription
}

func (o *NamednetworkNamedNetwork) GetNamedNetworkName() *string {
	if o == nil {
		return nil
	}
	return o.NamedNetworkName
}

func (o *NamednetworkNamedNetwork) GetNamednetworkTagBasedPolicyAssignments() *int64 {
	if o == nil {
		return nil
	}
	return o.NamednetworkTagBasedPolicyAssignments
}

func (o *NamednetworkNamedNetwork) GetProgramAsInternet() *bool {
	if o == nil {
		return nil
	}
	return o.ProgramAsInternet
}

func (o *NamednetworkNamedNetwork) GetProgramAsIntranet() *bool {
	if o == nil {
		return nil
	}
	return o.ProgramAsIntranet
}

func (o *NamednetworkNamedNetwork) GetVendor() *string {
	if o == nil {
		return nil
	}
	return o.Vendor
}

func (o *NamednetworkNamedNetwork) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *NamednetworkNamedNetwork) GetService() *string {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *NamednetworkNamedNetwork) GetTotalComments() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalComments
}

func (o *NamednetworkNamedNetwork) GetTotalCount() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalCount
}

func (o *NamednetworkNamedNetwork) GetUsergroupNamedNetworkAssignments() *int64 {
	if o == nil {
		return nil
	}
	return o.UsergroupNamedNetworkAssignments
}
