// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-xshield-sdk/internal/sdk/models/shared"
	"net/http"
)

type AnnotateAssetRequest struct {
	// Asset ID
	AssetID string `pathParam:"style=simple,explode=false,name=assetId"`
	// asset annotation details
	AssetAnnotationDetails shared.AssetAnnotationDetails `request:"mediaType=application/json"`
}

func (o *AnnotateAssetRequest) GetAssetID() string {
	if o == nil {
		return ""
	}
	return o.AssetID
}

func (o *AnnotateAssetRequest) GetAssetAnnotationDetails() shared.AssetAnnotationDetails {
	if o == nil {
		return shared.AssetAnnotationDetails{}
	}
	return o.AssetAnnotationDetails
}

type AnnotateAssetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	AssetAnnotationDetails *shared.AssetAnnotationDetails
	// Bad Request
	ErrorResponse *shared.ErrorResponse
}

func (o *AnnotateAssetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AnnotateAssetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AnnotateAssetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AnnotateAssetResponse) GetAssetAnnotationDetails() *shared.AssetAnnotationDetails {
	if o == nil {
		return nil
	}
	return o.AssetAnnotationDetails
}

func (o *AnnotateAssetResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}
