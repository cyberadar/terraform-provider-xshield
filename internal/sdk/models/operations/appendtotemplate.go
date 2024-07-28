// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-xshield-sdk/internal/sdk/models/shared"
	"net/http"
)

type AppendToTemplateRequest struct {
	// the template ID
	Templateid string `pathParam:"style=simple,explode=false,name=templateid"`
	// template rules
	TemplateBodyWithComment shared.TemplateBodyWithComment `request:"mediaType=application/json"`
}

func (o *AppendToTemplateRequest) GetTemplateid() string {
	if o == nil {
		return ""
	}
	return o.Templateid
}

func (o *AppendToTemplateRequest) GetTemplateBodyWithComment() shared.TemplateBodyWithComment {
	if o == nil {
		return shared.TemplateBodyWithComment{}
	}
	return o.TemplateBodyWithComment
}

type AppendToTemplateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Bad Request
	ErrorResponse *shared.ErrorResponse
	Headers       map[string][]string
}

func (o *AppendToTemplateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AppendToTemplateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AppendToTemplateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AppendToTemplateResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *AppendToTemplateResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
