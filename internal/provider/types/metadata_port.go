// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type MetadataPort struct {
	ID                 types.String   `tfsdk:"id"`
	ListenPort         types.String   `tfsdk:"listen_port"`
	ListenPortName     types.String   `tfsdk:"listen_port_name"`
	ListenPortProtocol types.String   `tfsdk:"listen_port_protocol"`
	ListenPortReviewed types.String   `tfsdk:"listen_port_reviewed"`
	ListenProcessNames []types.String `tfsdk:"listen_process_names"`
}
