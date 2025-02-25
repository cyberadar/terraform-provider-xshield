// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/colortokens/terraform-provider-xshield/internal/sdk/models/shared"
	"net/http"
)

// Scope - scope
type Scope string

const (
	ScopeAsset Scope = "asset"
	ScopePath  Scope = "path"
	ScopePort  Scope = "port"
	ScopeAgent Scope = "agent"
)

func (e Scope) ToPointer() *Scope {
	return &e
}
func (e *Scope) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asset":
		fallthrough
	case "path":
		fallthrough
	case "port":
		fallthrough
	case "agent":
		*e = Scope(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Scope: %v", v)
	}
}

type ListFieldsRequest struct {
	// scope
	Scope *Scope `queryParam:"style=form,explode=true,name=scope"`
}

func (o *ListFieldsRequest) GetScope() *Scope {
	if o == nil {
		return nil
	}
	return o.Scope
}

type ListFieldsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TypeaheadSuggestions *shared.TypeaheadSuggestions
	// Bad Request
	ErrorResponse *shared.ErrorResponse
}

func (o *ListFieldsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListFieldsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListFieldsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListFieldsResponse) GetTypeaheadSuggestions() *shared.TypeaheadSuggestions {
	if o == nil {
		return nil
	}
	return o.TypeaheadSuggestions
}

func (o *ListFieldsResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}
