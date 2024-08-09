// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/colortokens/terraform-provider-xshield/internal/sdk/models/shared"
	"net/http"
)

type CloneTemplateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created
	Template *shared.Template
	// Bad Request
	ErrorResponse *shared.ErrorResponse
}

func (o *CloneTemplateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CloneTemplateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CloneTemplateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CloneTemplateResponse) GetTemplate() *shared.Template {
	if o == nil {
		return nil
	}
	return o.Template
}

func (o *CloneTemplateResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}
