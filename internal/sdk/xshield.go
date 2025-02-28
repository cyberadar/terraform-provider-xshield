// Code originally generated by Speakeasy (https://www.speakeasyapi.dev).

package sdk

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/colortokens/terraform-provider-xshield/internal/sdk/internal/hooks"
	"github.com/colortokens/terraform-provider-xshield/internal/sdk/internal/utils"
	"github.com/colortokens/terraform-provider-xshield/internal/sdk/models/shared"
	"github.com/colortokens/terraform-provider-xshield/internal/sdk/retry"
)

const (
	// US Region (IAD)
	ServerIad string = "iad"
	// India Region (BOM)
	ServerBom string = "bom"
	// European Region (FRA)
	ServerFra string = "fra"
	// Australia Region (SYD)
	ServerSyd string = "syd"
)

// ServerList contains the list of servers available to the SDK
var ServerList = map[string]string{
	ServerIad: "https://ng.colortokens.com",
	ServerBom: "https://bom.colortokens.com",
	ServerFra: "https://fra.colortokens.com",
	ServerSyd: "https://syd.colortokens.com",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// Pointer provides a helper function to return a pointer to a type
func Pointer[T any](v T) *T { return &v }

type sdkConfiguration struct {
	Client            HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	Server            string
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *retry.Config
	Hooks             *hooks.Hooks
	Timeout           *time.Duration
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	if c.Server == "" {
		c.Server = "iad"
	}

	return ServerList[c.Server], nil
}

// Xshield - ColorTokens Core API: API for managing lifecycle of core micro-segmentation resources (tags, assets & groups)
type Xshield struct {
	Assets           *Assets
	Openports        *Openports
	Paths            *Paths
	Tags             *Tags
	Namednetworks    *Namednetworks
	Templates        *Templates
	Events           *Events
	Metadata         *Metadata
	Recommendations  *Recommendations
	Tagbasedpolicies *Tagbasedpolicies
	Tagrules         *Tagrules
	Unmanageddevice  *Unmanageddevice
	Usergroups       *Usergroups
	Workrequests     *Workrequests

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Xshield)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Xshield) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Xshield) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServer allows the overriding of the default server by name
func WithServer(server string) SDKOption {
	return func(sdk *Xshield) {
		_, ok := ServerList[server]
		if !ok {
			panic(fmt.Errorf("server %s not found", server))
		}

		sdk.sdkConfiguration.Server = server
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Xshield) {
		sdk.sdkConfiguration.Client = client
	}
}

// // WithSecurity configures the SDK to use the provided security details
// func WithSecurity(security shared.Security) SDKOption {
// 	return func(sdk *Xshield) {
// 		sdk.sdkConfiguration.Security = utils.AsSecuritySource(security)
// 	}
// }

// // WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
// func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
// 	return func(sdk *Xshield) {
// 		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
// 			return security(ctx)
// 		}
// 	}
// }

// WithSecurity configures the SDK to use the provided configuration details
func WithConfigProvider(config shared.ConfigurationProvider) SDKOption {
	return func(sdk *Xshield) {
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(config)
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *Xshield) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *Xshield) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Xshield {
	sdk := &Xshield{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "202202",
			SDKVersion:        "0.3.2",
			GenVersion:        "2.529.2",
			UserAgent:         "speakeasy-sdk/terraform 0.3.2 2.529.2 202202 github.com/colortokens/terraform-provider-xshield/internal/sdk",
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.Assets = newAssets(sdk.sdkConfiguration)

	sdk.Openports = newOpenports(sdk.sdkConfiguration)

	sdk.Paths = newPaths(sdk.sdkConfiguration)

	sdk.Tags = newTags(sdk.sdkConfiguration)

	sdk.Namednetworks = newNamednetworks(sdk.sdkConfiguration)

	sdk.Templates = newTemplates(sdk.sdkConfiguration)

	sdk.Events = newEvents(sdk.sdkConfiguration)

	sdk.Metadata = newMetadata(sdk.sdkConfiguration)

	sdk.Recommendations = newRecommendations(sdk.sdkConfiguration)

	sdk.Tagbasedpolicies = newTagbasedpolicies(sdk.sdkConfiguration)

	sdk.Tagrules = newTagrules(sdk.sdkConfiguration)

	sdk.Unmanageddevice = newUnmanageddevice(sdk.sdkConfiguration)

	sdk.Usergroups = newUsergroups(sdk.sdkConfiguration)

	sdk.Workrequests = newWorkrequests(sdk.sdkConfiguration)

	return sdk
}
