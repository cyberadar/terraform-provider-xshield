// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetFirewallRulesRequest struct {
	// assetId to list log files for
	AssetID string `pathParam:"style=simple,explode=false,name=assetId"`
	// name of the log file to download
	Objectname string `pathParam:"style=simple,explode=false,name=objectname"`
}

func (o *GetFirewallRulesRequest) GetAssetID() string {
	if o == nil {
		return ""
	}
	return o.AssetID
}

func (o *GetFirewallRulesRequest) GetObjectname() string {
	if o == nil {
		return ""
	}
	return o.Objectname
}

type GetFirewallRulesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Bytes []byte
	Body  []byte
}

func (o *GetFirewallRulesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetFirewallRulesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetFirewallRulesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetFirewallRulesResponse) GetBytes() []byte {
	if o == nil {
		return nil
	}
	return o.Bytes
}

func (o *GetFirewallRulesResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}
