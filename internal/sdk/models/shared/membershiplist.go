// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type MembershipList struct {
	Items []string `json:"items,omitempty"`
}

func (o *MembershipList) GetItems() []string {
	if o == nil {
		return nil
	}
	return o.Items
}
