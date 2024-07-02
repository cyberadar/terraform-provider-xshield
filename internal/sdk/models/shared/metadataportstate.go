// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type MetadataPortState string

const (
	MetadataPortStatePortUnreviewed              MetadataPortState = "PortUnreviewed"
	MetadataPortStatePortDenied                  MetadataPortState = "PortDenied"
	MetadataPortStatePortAllowIntranet           MetadataPortState = "PortAllowIntranet"
	MetadataPortStatePortAllowAny                MetadataPortState = "PortAllowAny"
	MetadataPortStatePortPathRestricted          MetadataPortState = "PortPathRestricted"
	MetadataPortStatePortDeniedByTemplate        MetadataPortState = "PortDeniedByTemplate"
	MetadataPortStatePortAllowIntranetByTemplate MetadataPortState = "PortAllowIntranetByTemplate"
	MetadataPortStatePortAllowAnyByTemplate      MetadataPortState = "PortAllowAnyByTemplate"
	MetadataPortStatePortAllowAnyByProgressive   MetadataPortState = "PortAllowAnyByProgressive"
)

func (e MetadataPortState) ToPointer() *MetadataPortState {
	return &e
}
func (e *MetadataPortState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PortUnreviewed":
		fallthrough
	case "PortDenied":
		fallthrough
	case "PortAllowIntranet":
		fallthrough
	case "PortAllowAny":
		fallthrough
	case "PortPathRestricted":
		fallthrough
	case "PortDeniedByTemplate":
		fallthrough
	case "PortAllowIntranetByTemplate":
		fallthrough
	case "PortAllowAnyByTemplate":
		fallthrough
	case "PortAllowAnyByProgressive":
		*e = MetadataPortState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MetadataPortState: %v", v)
	}
}
