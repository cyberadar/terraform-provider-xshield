// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type NamednetworkRange struct {
	ID      *string `json:"id,omitempty"`
	IPCount *int64  `json:"ipCount,omitempty"`
	IPRange *string `json:"ipRange,omitempty"`
}

func (o *NamednetworkRange) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *NamednetworkRange) GetIPCount() *int64 {
	if o == nil {
		return nil
	}
	return o.IPCount
}

func (o *NamednetworkRange) GetIPRange() *string {
	if o == nil {
		return nil
	}
	return o.IPRange
}
