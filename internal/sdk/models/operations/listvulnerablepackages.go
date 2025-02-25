// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/colortokens/terraform-provider-xshield/internal/sdk/models/shared"
	"net/http"
)

type ListVulnerablePackagesRequest struct {
	// Asset ID
	AssetID string `pathParam:"style=simple,explode=false,name=assetId"`
	// vulnerable packages request
	PaginationConfig shared.PaginationConfig `request:"mediaType=application/json"`
}

func (o *ListVulnerablePackagesRequest) GetAssetID() string {
	if o == nil {
		return ""
	}
	return o.AssetID
}

func (o *ListVulnerablePackagesRequest) GetPaginationConfig() shared.PaginationConfig {
	if o == nil {
		return shared.PaginationConfig{}
	}
	return o.PaginationConfig
}

type ListVulnerablePackagesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	VulnerablePackagesResults *shared.VulnerablePackagesResults
	// Bad Request
	ErrorResponse *shared.ErrorResponse
}

func (o *ListVulnerablePackagesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListVulnerablePackagesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListVulnerablePackagesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListVulnerablePackagesResponse) GetVulnerablePackagesResults() *shared.VulnerablePackagesResults {
	if o == nil {
		return nil
	}
	return o.VulnerablePackagesResults
}

func (o *ListVulnerablePackagesResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}
