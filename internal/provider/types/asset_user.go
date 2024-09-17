// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type AssetUser struct {
	Assetid    types.String `tfsdk:"assetid"`
	Domainname types.String `tfsdk:"domainname"`
	Email      types.String `tfsdk:"email"`
	Logincount types.Int64  `tfsdk:"logincount"`
	Name       types.String `tfsdk:"name"`
	Scimuserid types.String `tfsdk:"scimuserid"`
	Signedin   types.Bool   `tfsdk:"signedin"`
}
