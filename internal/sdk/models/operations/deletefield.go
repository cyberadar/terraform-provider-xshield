// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-xshield-sdk/internal/sdk/models/shared"
	"net/http"
)

type DeleteFieldResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	MetadataColumnDescriptor *shared.MetadataColumnDescriptor
	// Bad Request
	ErrorResponse *shared.ErrorResponse
}

func (o *DeleteFieldResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteFieldResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteFieldResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteFieldResponse) GetMetadataColumnDescriptor() *shared.MetadataColumnDescriptor {
	if o == nil {
		return nil
	}
	return o.MetadataColumnDescriptor
}

func (o *DeleteFieldResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}
