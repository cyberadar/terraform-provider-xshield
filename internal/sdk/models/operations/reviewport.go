// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-xshield-sdk/internal/sdk/models/shared"
	"net/http"
)

type ReviewPortRequest struct {
	// Port ID
	PortID string `pathParam:"style=simple,explode=false,name=portId"`
	// To State
	StateTransitionInput shared.StateTransitionInput `request:"mediaType=application/json"`
}

func (o *ReviewPortRequest) GetPortID() string {
	if o == nil {
		return ""
	}
	return o.PortID
}

func (o *ReviewPortRequest) GetStateTransitionInput() shared.StateTransitionInput {
	if o == nil {
		return shared.StateTransitionInput{}
	}
	return o.StateTransitionInput
}

type ReviewPortResponse struct {
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

func (o *ReviewPortResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ReviewPortResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ReviewPortResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ReviewPortResponse) GetString() *string {
	if o == nil {
		return nil
	}
	return o.String
}

func (o *ReviewPortResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}
