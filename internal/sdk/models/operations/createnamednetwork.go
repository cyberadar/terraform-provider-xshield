// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/colortokens/terraform-provider-xshield/internal/sdk/models/shared"
	"net/http"
)

type CreateNamedNetworkResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created
	NamednetworkNamedNetwork *shared.NamednetworkNamedNetwork
	// Bad Request
	ErrorResponse *shared.ErrorResponse
}

func (o *CreateNamedNetworkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateNamedNetworkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateNamedNetworkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateNamedNetworkResponse) GetNamednetworkNamedNetwork() *shared.NamednetworkNamedNetwork {
	if o == nil {
		return nil
	}
	return o.NamednetworkNamedNetwork
}

func (o *CreateNamedNetworkResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}
