// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Recommendation struct {
	Attributes map[string]any          `json:"attributes,omitempty"`
	Category   *RecommendationCategory `json:"category,omitempty"`
	Name       *string                 `json:"name,omitempty"`
	Priority   *string                 `json:"priority,omitempty"`
}

func (o *Recommendation) GetAttributes() map[string]any {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *Recommendation) GetCategory() *RecommendationCategory {
	if o == nil {
		return nil
	}
	return o.Category
}

func (o *Recommendation) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Recommendation) GetPriority() *string {
	if o == nil {
		return nil
	}
	return o.Priority
}
