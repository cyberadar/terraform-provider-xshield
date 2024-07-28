// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// WorkRequests definition Returned by work request listing
type WorkRequests struct {
	// List of workRequests that matched any filters specified
	Items    []WorkrequestWorkRequest `json:"items,omitempty"`
	Metadata *PaginationSummary       `json:"metadata,omitempty"`
}

func (o *WorkRequests) GetItems() []WorkrequestWorkRequest {
	if o == nil {
		return nil
	}
	return o.Items
}

func (o *WorkRequests) GetMetadata() *PaginationSummary {
	if o == nil {
		return nil
	}
	return o.Metadata
}
