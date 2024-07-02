// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ErrorResponse - Error information if api invocation resulted in an error.
type ErrorResponse struct {
	// Details specific to the error occurrence.
	Details map[string]any `json:"details,omitempty"`
	// Message summarizing error information and possible remediation if applicable.
	Message *string `json:"message,omitempty"`
	// HTTP status code providing indication of the type of error encountered.
	Status *int64 `json:"status,omitempty"`
}

func (o *ErrorResponse) GetDetails() map[string]any {
	if o == nil {
		return nil
	}
	return o.Details
}

func (o *ErrorResponse) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ErrorResponse) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}
