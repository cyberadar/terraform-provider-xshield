// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-xshield-sdk/internal/sdk/models/shared"
	"net/http"
)

type ListRecommendationsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Recommendations *shared.Recommendations
	// Bad Request
	ErrorResponse *shared.ErrorResponse
}

func (o *ListRecommendationsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListRecommendationsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListRecommendationsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListRecommendationsResponse) GetRecommendations() *shared.Recommendations {
	if o == nil {
		return nil
	}
	return o.Recommendations
}

func (o *ListRecommendationsResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}
