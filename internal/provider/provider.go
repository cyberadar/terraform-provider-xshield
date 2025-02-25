// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/colortokens/terraform-provider-xshield/internal/sdk"
	"github.com/colortokens/terraform-provider-xshield/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"net/http"
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
	Jwt       types.String `tfsdk:"jwt"`
	ServerURL types.String `tfsdk:"server_url"`
}

func (p *XshieldProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "xshield"
	resp.Version = p.version
}

func (p *XshieldProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"jwt": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"server_url": schema.StringAttribute{
				Description: `Server URL (defaults to https://ng.colortokens.com)`,
				Optional:    true,
			},
		},
		MarkdownDescription: `ColorTokens Core API: API for managing lifecycle of core micro-segmentation resources (tags, assets & groups)`,
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

	jwt := new(string)
	if !data.Jwt.IsUnknown() && !data.Jwt.IsNull() {
		*jwt = data.Jwt.ValueString()
	} else {
		jwt = nil
	}
	security := shared.Security{
		Jwt: jwt,
	}

	providerHTTPTransportOpts := ProviderHTTPTransportOpts{
		SetHeaders: make(map[string]string),
		Transport:  http.DefaultTransport,
	}

	httpClient := http.DefaultClient
	httpClient.Transport = NewProviderHTTPTransport(providerHTTPTransportOpts)

	opts := []sdk.SDKOption{
		sdk.WithServerURL(ServerURL),
		sdk.WithSecurity(security),
		sdk.WithClient(httpClient),
	}
	client := sdk.New(opts...)

	resp.DataSourceData = client
	resp.ResourceData = client
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
