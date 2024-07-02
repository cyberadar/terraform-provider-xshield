// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AssetIDInfo - Asset ids
type AssetIDInfo struct {
	AgentID         *string `json:"agentId,omitempty"`
	AssetID         *string `json:"assetId,omitempty"`
	DeterministicID string  `json:"deterministicId"`
}

func (o *AssetIDInfo) GetAgentID() *string {
	if o == nil {
		return nil
	}
	return o.AgentID
}

func (o *AssetIDInfo) GetAssetID() *string {
	if o == nil {
		return nil
	}
	return o.AssetID
}

func (o *AssetIDInfo) GetDeterministicID() string {
	if o == nil {
		return ""
	}
	return o.DeterministicID
}
