// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// AssetSummary definition Summary of host or application running on a host that can be observed to apply segmentation policies against
type AssetSummary struct {
	AssetID                             *string           `json:"assetId,omitempty"`
	AssetName                           string            `json:"assetName"`
	AutoSynchronizeEnabled              *bool             `json:"autoSynchronizeEnabled,omitempty"`
	ClusterIdentifier                   *string           `json:"clusterIdentifier,omitempty"`
	ContainerNamespace                  *string           `json:"containerNamespace,omitempty"`
	CoreTags                            map[string]string `json:"coreTags,omitempty"`
	InboundAssetStatus                  *string           `json:"inboundAssetStatus,omitempty"`
	LowestInboundAssetStatus            *string           `json:"lowestInboundAssetStatus,omitempty"`
	LowestOutboundAssetStatus           *string           `json:"lowestOutboundAssetStatus,omitempty"`
	LowestProgressiveInboundAssetStatus *string           `json:"lowestProgressiveInboundAssetStatus,omitempty"`
	OutboundAssetStatus                 *string           `json:"outboundAssetStatus,omitempty"`
	Type                                string            `json:"type"`
	VendorInfo                          *string           `json:"vendorInfo,omitempty"`
}

func (o *AssetSummary) GetAssetID() *string {
	if o == nil {
		return nil
	}
	return o.AssetID
}

func (o *AssetSummary) GetAssetName() string {
	if o == nil {
		return ""
	}
	return o.AssetName
}

func (o *AssetSummary) GetAutoSynchronizeEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.AutoSynchronizeEnabled
}

func (o *AssetSummary) GetClusterIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.ClusterIdentifier
}

func (o *AssetSummary) GetContainerNamespace() *string {
	if o == nil {
		return nil
	}
	return o.ContainerNamespace
}

func (o *AssetSummary) GetCoreTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.CoreTags
}

func (o *AssetSummary) GetInboundAssetStatus() *string {
	if o == nil {
		return nil
	}
	return o.InboundAssetStatus
}

func (o *AssetSummary) GetLowestInboundAssetStatus() *string {
	if o == nil {
		return nil
	}
	return o.LowestInboundAssetStatus
}

func (o *AssetSummary) GetLowestOutboundAssetStatus() *string {
	if o == nil {
		return nil
	}
	return o.LowestOutboundAssetStatus
}

func (o *AssetSummary) GetLowestProgressiveInboundAssetStatus() *string {
	if o == nil {
		return nil
	}
	return o.LowestProgressiveInboundAssetStatus
}

func (o *AssetSummary) GetOutboundAssetStatus() *string {
	if o == nil {
		return nil
	}
	return o.OutboundAssetStatus
}

func (o *AssetSummary) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *AssetSummary) GetVendorInfo() *string {
	if o == nil {
		return nil
	}
	return o.VendorInfo
}
