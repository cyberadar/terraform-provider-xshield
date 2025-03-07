// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ApplyNamedNetworkSearchInput struct {
	Comment        *string  `json:"comment,omitempty"`
	Criteria       string   `json:"criteria"`
	NamedNetworkID *string  `json:"namedNetworkId,omitempty"`
	Namednetworks  []string `json:"namednetworks,omitempty"`
}

func (o *ApplyNamedNetworkSearchInput) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *ApplyNamedNetworkSearchInput) GetCriteria() string {
	if o == nil {
		return ""
	}
	return o.Criteria
}

func (o *ApplyNamedNetworkSearchInput) GetNamedNetworkID() *string {
	if o == nil {
		return nil
	}
	return o.NamedNetworkID
}

func (o *ApplyNamedNetworkSearchInput) GetNamednetworks() []string {
	if o == nil {
		return nil
	}
	return o.Namednetworks
}
