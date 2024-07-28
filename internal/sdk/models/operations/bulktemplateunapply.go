// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-xshield-sdk/internal/sdk/models/shared"
	"net/http"
)

type BulkTemplateUnApplyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Accepted. Async processing in-progress and can be tracked using value in header 'x-ct-workrequest-id'
	TwoHundredAndTwoApplicationJSONString *string
	// Not Modified
	ThreeHundredAndFourApplicationJSONString *string
	// Bad Request
	ErrorResponse *shared.ErrorResponse
	Headers       map[string][]string
}

func (o *BulkTemplateUnApplyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *BulkTemplateUnApplyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *BulkTemplateUnApplyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *BulkTemplateUnApplyResponse) GetTwoHundredAndTwoApplicationJSONString() *string {
	if o == nil {
		return nil
	}
	return o.TwoHundredAndTwoApplicationJSONString
}

func (o *BulkTemplateUnApplyResponse) GetThreeHundredAndFourApplicationJSONString() *string {
	if o == nil {
		return nil
	}
	return o.ThreeHundredAndFourApplicationJSONString
}

func (o *BulkTemplateUnApplyResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *BulkTemplateUnApplyResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
