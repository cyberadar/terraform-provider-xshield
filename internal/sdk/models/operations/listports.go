// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/colortokens/terraform-provider-xshield/internal/sdk/models/shared"
	"net/http"
)

type ListPortsRequest struct {
	// file download
	Download *string `queryParam:"style=form,explode=true,name=download"`
	// search criteria for filtering open ports
	PathSearchInput shared.PathSearchInput `request:"mediaType=application/json"`
}

func (o *ListPortsRequest) GetDownload() *string {
	if o == nil {
		return nil
	}
	return o.Download
}

func (o *ListPortsRequest) GetPathSearchInput() shared.PathSearchInput {
	if o == nil {
		return shared.PathSearchInput{}
	}
	return o.PathSearchInput
}

type ListPortsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	OpenPorts *shared.OpenPorts
	Body      []byte
	// Bad Request
	ErrorResponse *shared.ErrorResponse
}

func (o *ListPortsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListPortsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListPortsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListPortsResponse) GetOpenPorts() *shared.OpenPorts {
	if o == nil {
		return nil
	}
	return o.OpenPorts
}

func (o *ListPortsResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *ListPortsResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}
