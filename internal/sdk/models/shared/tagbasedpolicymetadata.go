// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type TagBasedPolicyMetadata struct {
	Description        *string `json:"description,omitempty"`
	TagBasedPolicyName *string `json:"tagBasedPolicyName,omitempty"`
}

func (o *TagBasedPolicyMetadata) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *TagBasedPolicyMetadata) GetTagBasedPolicyName() *string {
	if o == nil {
		return nil
	}
	return o.TagBasedPolicyName
}
