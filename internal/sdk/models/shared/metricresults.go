// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type MetricResults struct {
	Items map[string]any `json:"items,omitempty"`
}

func (o *MetricResults) GetItems() map[string]any {
	if o == nil {
		return nil
	}
	return o.Items
}
