// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/colortokens/terraform-provider-xshield/internal/sdk/models/shared"
	"net/http"
)

type ListPathsWithFirewallLogsReportedRequest struct {
	// file download
	Download *string `queryParam:"style=form,explode=true,name=download"`
	// search criteria for filtering paths
	PathSearchInput shared.PathSearchInput `request:"mediaType=application/json"`
}

func (o *ListPathsWithFirewallLogsReportedRequest) GetDownload() *string {
	if o == nil {
		return nil
	}
	return o.Download
}

func (o *ListPathsWithFirewallLogsReportedRequest) GetPathSearchInput() shared.PathSearchInput {
	if o == nil {
		return shared.PathSearchInput{}
	}
	return o.PathSearchInput
}

type ListPathsWithFirewallLogsReportedResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Paths *shared.Paths
	// Bad Request
	ErrorResponse *shared.ErrorResponse
}

func (o *ListPathsWithFirewallLogsReportedResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListPathsWithFirewallLogsReportedResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListPathsWithFirewallLogsReportedResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListPathsWithFirewallLogsReportedResponse) GetPaths() *shared.Paths {
	if o == nil {
		return nil
	}
	return o.Paths
}

func (o *ListPathsWithFirewallLogsReportedResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}
