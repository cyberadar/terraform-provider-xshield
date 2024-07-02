// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-xshield-sdk/internal/sdk/models/shared"
	"net/http"
)

type EditTemplateMetadataRequest struct {
	// the template ID
	Templateid string `pathParam:"style=simple,explode=false,name=templateid"`
	// template metadata
	TemplateSummary shared.TemplateSummary `request:"mediaType=application/json"`
}

func (o *EditTemplateMetadataRequest) GetTemplateid() string {
	if o == nil {
		return ""
	}
	return o.Templateid
}

func (o *EditTemplateMetadataRequest) GetTemplateSummary() shared.TemplateSummary {
	if o == nil {
		return shared.TemplateSummary{}
	}
	return o.TemplateSummary
}

type EditTemplateMetadataResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Bad Request
	ErrorResponse *shared.ErrorResponse
}

func (o *EditTemplateMetadataResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *EditTemplateMetadataResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *EditTemplateMetadataResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *EditTemplateMetadataResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}
