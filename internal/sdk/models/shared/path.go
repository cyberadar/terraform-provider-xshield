// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy/terraform-provider-xshield-sdk/internal/sdk/internal/utils"
	"time"
)

type Path struct {
	BandwidthInBytes *int64  `json:"bandwidthInBytes,omitempty"`
	ChannelHash      *string `json:"channelHash,omitempty"`
	CompassDirection *string `json:"compassDirection,omitempty"`
	ConnectionCount  *int64  `json:"connectionCount,omitempty"`
	// AssetSummary definition Summary of host or application running on a host that can be observed to apply segmentation policies against
	DestinationAsset         *AssetSummary                  `json:"destinationAsset,omitempty"`
	DestinationNamedNetwork  *MetadataNamedNetworkReference `json:"destinationNamedNetwork,omitempty"`
	DestinationProcess       *string                        `json:"destinationProcess,omitempty"`
	Direction                *string                        `json:"direction,omitempty"`
	Domain                   *string                        `json:"domain,omitempty"`
	DstIP                    []string                       `json:"dstIp,omitempty"`
	Enforced                 *string                        `json:"enforced,omitempty"`
	FirewallAllowCount       *int64                         `json:"firewallAllowCount,omitempty"`
	FirewallDenyCount        *int64                         `json:"firewallDenyCount,omitempty"`
	FirewallLastReported     *time.Time                     `json:"firewallLastReported,omitempty"`
	InternetBandwidthInBytes *int64                         `json:"internetBandwidthInBytes,omitempty"`
	InternetFacing           *bool                          `json:"internetFacing,omitempty"`
	IntranetBandwidthInBytes *int64                         `json:"intranetBandwidthInBytes,omitempty"`
	ListenPortEnforced       *string                        `json:"listenPortEnforced,omitempty"`
	ListenPortReviewed       *string                        `json:"listenPortReviewed,omitempty"`
	MatchedByTemplates       []TemplateReference            `json:"matchedByTemplates,omitempty"`
	Method                   *string                        `json:"method,omitempty"`
	PathLastObserved         *time.Time                     `json:"pathLastObserved,omitempty"`
	Port                     *string                        `json:"port,omitempty"`
	PortName                 *string                        `json:"portName,omitempty"`
	Protocol                 *string                        `json:"protocol,omitempty"`
	Reviewed                 *string                        `json:"reviewed,omitempty"`
	// AssetSummary definition Summary of host or application running on a host that can be observed to apply segmentation policies against
	SourceAsset        *AssetSummary                  `json:"sourceAsset,omitempty"`
	SourceNamedNetwork *MetadataNamedNetworkReference `json:"sourceNamedNetwork,omitempty"`
	SourceProcess      *string                        `json:"sourceProcess,omitempty"`
	SrcIP              *string                        `json:"srcIp,omitempty"`
	TotalComments      *int64                         `json:"totalComments,omitempty"`
	URI                *string                        `json:"uri,omitempty"`
}

func (p Path) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *Path) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Path) GetBandwidthInBytes() *int64 {
	if o == nil {
		return nil
	}
	return o.BandwidthInBytes
}

func (o *Path) GetChannelHash() *string {
	if o == nil {
		return nil
	}
	return o.ChannelHash
}

func (o *Path) GetCompassDirection() *string {
	if o == nil {
		return nil
	}
	return o.CompassDirection
}

func (o *Path) GetConnectionCount() *int64 {
	if o == nil {
		return nil
	}
	return o.ConnectionCount
}

func (o *Path) GetDestinationAsset() *AssetSummary {
	if o == nil {
		return nil
	}
	return o.DestinationAsset
}

func (o *Path) GetDestinationNamedNetwork() *MetadataNamedNetworkReference {
	if o == nil {
		return nil
	}
	return o.DestinationNamedNetwork
}

func (o *Path) GetDestinationProcess() *string {
	if o == nil {
		return nil
	}
	return o.DestinationProcess
}

func (o *Path) GetDirection() *string {
	if o == nil {
		return nil
	}
	return o.Direction
}

func (o *Path) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *Path) GetDstIP() []string {
	if o == nil {
		return nil
	}
	return o.DstIP
}

func (o *Path) GetEnforced() *string {
	if o == nil {
		return nil
	}
	return o.Enforced
}

func (o *Path) GetFirewallAllowCount() *int64 {
	if o == nil {
		return nil
	}
	return o.FirewallAllowCount
}

func (o *Path) GetFirewallDenyCount() *int64 {
	if o == nil {
		return nil
	}
	return o.FirewallDenyCount
}

func (o *Path) GetFirewallLastReported() *time.Time {
	if o == nil {
		return nil
	}
	return o.FirewallLastReported
}

func (o *Path) GetInternetBandwidthInBytes() *int64 {
	if o == nil {
		return nil
	}
	return o.InternetBandwidthInBytes
}

func (o *Path) GetInternetFacing() *bool {
	if o == nil {
		return nil
	}
	return o.InternetFacing
}

func (o *Path) GetIntranetBandwidthInBytes() *int64 {
	if o == nil {
		return nil
	}
	return o.IntranetBandwidthInBytes
}

func (o *Path) GetListenPortEnforced() *string {
	if o == nil {
		return nil
	}
	return o.ListenPortEnforced
}

func (o *Path) GetListenPortReviewed() *string {
	if o == nil {
		return nil
	}
	return o.ListenPortReviewed
}

func (o *Path) GetMatchedByTemplates() []TemplateReference {
	if o == nil {
		return nil
	}
	return o.MatchedByTemplates
}

func (o *Path) GetMethod() *string {
	if o == nil {
		return nil
	}
	return o.Method
}

func (o *Path) GetPathLastObserved() *time.Time {
	if o == nil {
		return nil
	}
	return o.PathLastObserved
}

func (o *Path) GetPort() *string {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *Path) GetPortName() *string {
	if o == nil {
		return nil
	}
	return o.PortName
}

func (o *Path) GetProtocol() *string {
	if o == nil {
		return nil
	}
	return o.Protocol
}

func (o *Path) GetReviewed() *string {
	if o == nil {
		return nil
	}
	return o.Reviewed
}

func (o *Path) GetSourceAsset() *AssetSummary {
	if o == nil {
		return nil
	}
	return o.SourceAsset
}

func (o *Path) GetSourceNamedNetwork() *MetadataNamedNetworkReference {
	if o == nil {
		return nil
	}
	return o.SourceNamedNetwork
}

func (o *Path) GetSourceProcess() *string {
	if o == nil {
		return nil
	}
	return o.SourceProcess
}

func (o *Path) GetSrcIP() *string {
	if o == nil {
		return nil
	}
	return o.SrcIP
}

func (o *Path) GetTotalComments() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalComments
}

func (o *Path) GetURI() *string {
	if o == nil {
		return nil
	}
	return o.URI
}
