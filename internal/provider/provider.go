// Code originally generated by Speakeasy (https://www.speakeasyapi.dev).

package provider

import (
	"context"
	"net/http"

	"github.com/colortokens/terraform-provider-xshield/internal/sdk"
	"github.com/colortokens/terraform-provider-xshield/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &XshieldProvider{}

type XshieldProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// XshieldProviderModel describes the provider data model.
type XshieldProviderModel struct {
	ServerURL          types.String `tfsdk:"server_url"`
	TenancyId          types.String `tfsdk:"tenancy_id"`
	PrincipalId        types.String `tfsdk:"user_id"`
	FingerPrint        types.String `tfsdk:"fingerprint"`
	PrivateKeyLocation types.String `tfsdk:"private_key_path"`
}

func (p *XshieldProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "xshield"
	resp.Version = p.version
}

func (p *XshieldProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: `ColorTokens Core API: API for managing lifecycle of core micro-segmentation resources (tags, assets & groups)`,
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				MarkdownDescription: "Server URL (defaults to https://ng.colortokens.com)",
				Optional:            true,
				Required:            false,
			},
			"tenancy_id": schema.StringAttribute{
				Required:  true,
				Sensitive: false,
			},
			"user_id": schema.StringAttribute{
				Required:  true,
				Sensitive: false,
			}, "fingerprint": schema.StringAttribute{
				Required:  true,
				Sensitive: false,
			}, "private_key_path": schema.StringAttribute{
				Required:  true,
				Sensitive: false,
			},
		},
	}
}

func (p *XshieldProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data XshieldProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ServerURL := data.ServerURL.ValueString()

	if ServerURL == "" {
		ServerURL = "https://ng.colortokens.com"
	}

	config := buildConfigProvider(data, resp)

	opts := []sdk.SDKOption{
		sdk.WithServerURL(ServerURL),
		sdk.WithConfigProvider(config),
		sdk.WithClient(http.DefaultClient),
	}
	client := sdk.New(opts...)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func buildConfigProvider(data XshieldProviderModel, resp *provider.ConfigureResponse) shared.ConfigurationProvider {
	if data.TenancyId.ValueString() == "" {
		resp.Diagnostics.AddAttributeError(path.Root("provider").AtName("tenancy_id"), "required attribute is missing or empty", "")
	}

	if data.PrincipalId.ValueString() == "" {
		resp.Diagnostics.AddAttributeError(path.Root("provider").AtName("user_id"), "required attribute is missing or empty", "")
	}

	if data.FingerPrint.ValueString() == "" {
		resp.Diagnostics.AddAttributeError(path.Root("provider").AtName("fingerprint"), "required attribute is missing or empty", "")
	}

	if data.PrivateKeyLocation.ValueString() == "" {
		resp.Diagnostics.AddAttributeError(path.Root("provider").AtName("private_key_location"), "required attribute is missing or empty", "")
	}

	if resp.Diagnostics.HasError() {
		return nil
	}

	return shared.NewRawConfigProvider(data.TenancyId.ValueString(), data.PrincipalId.ValueString(), data.FingerPrint.ValueString(), data.PrivateKeyLocation.ValueString())
}

func (p *XshieldProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewAssetResource,
		NewNamedNetworkResource,
		NewSegmentResource,
		NewTagRuleResource,
		NewTemplateResource,
	}
}

func (p *XshieldProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAssetDataSource,
		NewNamedNetworkDataSource,
		NewSegmentDataSource,
		NewTagRuleDataSource,
		NewTemplateDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &XshieldProvider{
			version: version,
		}
	}
}
