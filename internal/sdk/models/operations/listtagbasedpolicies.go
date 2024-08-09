// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/colortokens/terraform-provider-xshield/internal/sdk/models/shared"
	"net/http"
)

type ListTagBasedPoliciesRequest struct {
	// file download
	Download *string `queryParam:"style=form,explode=true,name=download"`
	// search details
	SearchInput shared.SearchInput `request:"mediaType=application/json"`
}

func (o *ListTagBasedPoliciesRequest) GetDownload() *string {
	if o == nil {
		return nil
	}
	return o.Download
}

func (o *ListTagBasedPoliciesRequest) GetSearchInput() shared.SearchInput {
	if o == nil {
		return shared.SearchInput{}
	}
	return o.SearchInput
}

type ListTagBasedPoliciesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TagBasedPolicies *shared.TagBasedPolicies
	Body             []byte
	// Bad Request
	ErrorResponse *shared.ErrorResponse
}

func (o *ListTagBasedPoliciesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListTagBasedPoliciesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListTagBasedPoliciesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListTagBasedPoliciesResponse) GetTagBasedPolicies() *shared.TagBasedPolicies {
	if o == nil {
		return nil
	}
	return o.TagBasedPolicies
}

func (o *ListTagBasedPoliciesResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *ListTagBasedPoliciesResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}
