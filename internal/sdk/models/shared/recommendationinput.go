// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type RecommendationInput struct {
	Criteria string `json:"criteria"`
}

func (o *RecommendationInput) GetCriteria() string {
	if o == nil {
		return ""
	}
	return o.Criteria
}
