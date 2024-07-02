// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type UserGroupHits struct {
	Items    []UserGroupHit     `json:"items,omitempty"`
	Metadata *PaginationSummary `json:"metadata,omitempty"`
}

func (o *UserGroupHits) GetItems() []UserGroupHit {
	if o == nil {
		return nil
	}
	return o.Items
}

func (o *UserGroupHits) GetMetadata() *PaginationSummary {
	if o == nil {
		return nil
	}
	return o.Metadata
}
