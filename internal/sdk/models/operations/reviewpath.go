// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-xshield-sdk/internal/sdk/models/shared"
	"net/http"
)

type ReviewPathRequest struct {
	// Path ID
	PathID string `pathParam:"style=simple,explode=false,name=pathId"`
	// To State
	StateTransitionInput shared.StateTransitionInput `request:"mediaType=application/json"`
}

func (o *ReviewPathRequest) GetPathID() string {
	if o == nil {
		return ""
	}
	return o.PathID
}

func (o *ReviewPathRequest) GetStateTransitionInput() shared.StateTransitionInput {
	if o == nil {
		return shared.StateTransitionInput{}
	}
	return o.StateTransitionInput
}

type ReviewPathResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Not Modified
	String *string
	// Bad Request
	ErrorResponse *shared.ErrorResponse
}

func (o *ReviewPathResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ReviewPathResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ReviewPathResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ReviewPathResponse) GetString() *string {
	if o == nil {
		return nil
	}
	return o.String
}

func (o *ReviewPathResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}
