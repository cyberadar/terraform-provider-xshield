// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// UnmanagedDevicesSearchResults definition Returned by SearchUnmanagedDevices api.
type UnmanagedDevicesSearchResults struct {
	Items    []UnmanagedDeviceSummary `json:"items,omitempty"`
	Metadata *PaginationSummary       `json:"metadata,omitempty"`
}

func (o *UnmanagedDevicesSearchResults) GetItems() []UnmanagedDeviceSummary {
	if o == nil {
		return nil
	}
	return o.Items
}

func (o *UnmanagedDevicesSearchResults) GetMetadata() *PaginationSummary {
	if o == nil {
		return nil
	}
	return o.Metadata
}
