// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/colortokens/terraform-provider-xshield/internal/sdk/models/shared"
	"net/http"
)

type AddToNamedNetworkRequest struct {
	// the Named Network ID
	NamedNetworkID string `pathParam:"style=simple,explode=false,name=namedNetworkId"`
	// named network ranges
	RequestBody []shared.NamednetworkRange `request:"mediaType=application/json"`
}

func (o *AddToNamedNetworkRequest) GetNamedNetworkID() string {
	if o == nil {
		return ""
	}
	return o.NamedNetworkID
}

func (o *AddToNamedNetworkRequest) GetRequestBody() []shared.NamednetworkRange {
	if o == nil {
		return []shared.NamednetworkRange{}
	}
	return o.RequestBody
}

type AddToNamedNetworkResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Accepted
	NamednetworkRanges []shared.NamednetworkRange
	// Bad Request
	ErrorResponse *shared.ErrorResponse
	Headers       map[string][]string
}

func (o *AddToNamedNetworkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddToNamedNetworkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddToNamedNetworkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AddToNamedNetworkResponse) GetNamednetworkRanges() []shared.NamednetworkRange {
	if o == nil {
		return nil
	}
	return o.NamednetworkRanges
}

func (o *AddToNamedNetworkResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *AddToNamedNetworkResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
