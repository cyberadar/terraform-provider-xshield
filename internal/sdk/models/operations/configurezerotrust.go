// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/colortokens/terraform-provider-xshield/internal/sdk/models/shared"
	"net/http"
)

type ConfigureZeroTrustRequest struct {
	// Asset ID
	AssetID string `pathParam:"style=simple,explode=false,name=assetId"`
	// To State
	AssetStateTransitionInput shared.AssetStateTransitionInput `request:"mediaType=application/json"`
}

func (o *ConfigureZeroTrustRequest) GetAssetID() string {
	if o == nil {
		return ""
	}
	return o.AssetID
}

func (o *ConfigureZeroTrustRequest) GetAssetStateTransitionInput() shared.AssetStateTransitionInput {
	if o == nil {
		return shared.AssetStateTransitionInput{}
	}
	return o.AssetStateTransitionInput
}

type ConfigureZeroTrustResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Accepted. Async processing in-progress and can be tracked using value in header 'x-ct-workrequest-id'
	TwoHundredAndTwoApplicationJSONString *string
	// Not Modified
	ThreeHundredAndFourApplicationJSONString *string
	// Bad Request
	ErrorResponse *shared.ErrorResponse
	Headers       map[string][]string
}

func (o *ConfigureZeroTrustResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ConfigureZeroTrustResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ConfigureZeroTrustResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ConfigureZeroTrustResponse) GetTwoHundredAndTwoApplicationJSONString() *string {
	if o == nil {
		return nil
	}
	return o.TwoHundredAndTwoApplicationJSONString
}

func (o *ConfigureZeroTrustResponse) GetThreeHundredAndFourApplicationJSONString() *string {
	if o == nil {
		return nil
	}
	return o.ThreeHundredAndFourApplicationJSONString
}

func (o *ConfigureZeroTrustResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *ConfigureZeroTrustResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
