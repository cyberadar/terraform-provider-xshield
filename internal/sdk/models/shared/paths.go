// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Paths struct {
	Items    []Path             `json:"items,omitempty"`
	Metadata *PaginationSummary `json:"metadata,omitempty"`
}

func (o *Paths) GetItems() []Path {
	if o == nil {
		return nil
	}
	return o.Items
}

func (o *Paths) GetMetadata() *PaginationSummary {
	if o == nil {
		return nil
	}
	return o.Metadata
}
