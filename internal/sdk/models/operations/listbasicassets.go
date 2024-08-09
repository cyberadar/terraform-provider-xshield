// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/colortokens/terraform-provider-xshield/internal/sdk/models/shared"
	"net/http"
)

type ListBasicAssetsRequest struct {
	// file download
	Download *string `queryParam:"style=form,explode=true,name=download"`
	// search details
	SearchInput shared.SearchInput `request:"mediaType=application/json"`
}

func (o *ListBasicAssetsRequest) GetDownload() *string {
	if o == nil {
		return nil
	}
	return o.Download
}

func (o *ListBasicAssetsRequest) GetSearchInput() shared.SearchInput {
	if o == nil {
		return shared.SearchInput{}
	}
	return o.SearchInput
}

type ListBasicAssetsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	AssetSearchResults *shared.AssetSearchResults
	Body               []byte
	// Bad Request
	ErrorResponse *shared.ErrorResponse
}

func (o *ListBasicAssetsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListBasicAssetsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListBasicAssetsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListBasicAssetsResponse) GetAssetSearchResults() *shared.AssetSearchResults {
	if o == nil {
		return nil
	}
	return o.AssetSearchResults
}

func (o *ListBasicAssetsResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *ListBasicAssetsResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}
