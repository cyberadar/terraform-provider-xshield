// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/colortokens/terraform-provider-xshield/internal/sdk/models/shared"
	"net/http"
)

type ApplyTemplateRequest struct {
	// the template ID
	Templateid string `pathParam:"style=simple,explode=false,name=templateid"`
	// the asset ID to apply the template to
	Assetid string `pathParam:"style=simple,explode=false,name=assetid"`
	// template input
	TemplateInput shared.TemplateInput `request:"mediaType=application/json"`
}

func (o *ApplyTemplateRequest) GetTemplateid() string {
	if o == nil {
		return ""
	}
	return o.Templateid
}

func (o *ApplyTemplateRequest) GetAssetid() string {
	if o == nil {
		return ""
	}
	return o.Assetid
}

func (o *ApplyTemplateRequest) GetTemplateInput() shared.TemplateInput {
	if o == nil {
		return shared.TemplateInput{}
	}
	return o.TemplateInput
}

type ApplyTemplateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Bad Request
	ErrorResponse *shared.ErrorResponse
}

func (o *ApplyTemplateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ApplyTemplateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ApplyTemplateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ApplyTemplateResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}
