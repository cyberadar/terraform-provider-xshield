// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CreateAssetDetails struct {
	AgentID         *string           `json:"agentId,omitempty"`
	ID              *string           `json:"assetId,omitempty"`
	AssetName       string            `json:"assetName"`
	Type            string            `json:"type"`
	CoreTags        map[string]string `json:"coreTags,omitempty"`
	DeterministicID *string           `json:"deterministicId,omitempty"`
	VendorInfo      *string           `json:"vendorInfo,omitempty"`
}

func (o *CreateAssetDetails) GetAgentID() *string {
	if o == nil {
		return nil
	}
	return o.AgentID
}

func (o *CreateAssetDetails) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateAssetDetails) GetAssetName() string {
	if o == nil {
		return ""
	}
	return o.AssetName
}

func (o *CreateAssetDetails) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *CreateAssetDetails) GetCoreTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.CoreTags
}

func (o *CreateAssetDetails) GetDeterministicID() *string {
	if o == nil {
		return nil
	}
	return o.DeterministicID
}

func (o *CreateAssetDetails) GetVendorInfo() *string {
	if o == nil {
		return nil
	}
	return o.VendorInfo
}
