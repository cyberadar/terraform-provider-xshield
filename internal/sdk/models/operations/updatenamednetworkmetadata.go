// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-xshield-sdk/internal/sdk/models/shared"
	"net/http"
)

type UpdateNamedNetworkMetadataRequest struct {
	// Named Network ID
	NamedNetworkID string `pathParam:"style=simple,explode=false,name=namedNetworkId"`
	// named network definition
	NamednetworkNamedNetwork shared.NamednetworkNamedNetwork `request:"mediaType=application/json"`
}

func (o *UpdateNamedNetworkMetadataRequest) GetNamedNetworkID() string {
	if o == nil {
		return ""
	}
	return o.NamedNetworkID
}

func (o *UpdateNamedNetworkMetadataRequest) GetNamednetworkNamedNetwork() shared.NamednetworkNamedNetwork {
	if o == nil {
		return shared.NamednetworkNamedNetwork{}
	}
	return o.NamednetworkNamedNetwork
}

type UpdateNamedNetworkMetadataResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Bad Request
	ErrorResponse *shared.ErrorResponse
}

func (o *UpdateNamedNetworkMetadataResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateNamedNetworkMetadataResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateNamedNetworkMetadataResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateNamedNetworkMetadataResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}
