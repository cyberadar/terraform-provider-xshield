// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type MetadataNamedNetworkReference struct {
	NamedNetworkID   *string `json:"namedNetworkId,omitempty"`
	NamedNetworkName *string `json:"namedNetworkName,omitempty"`
}

func (o *MetadataNamedNetworkReference) GetNamedNetworkID() *string {
	if o == nil {
		return nil
	}
	return o.NamedNetworkID
}

func (o *MetadataNamedNetworkReference) GetNamedNetworkName() *string {
	if o == nil {
		return nil
	}
	return o.NamedNetworkName
}
