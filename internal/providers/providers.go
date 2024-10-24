// Package providers contains the changes for both SDKv2 and Plugin Framework which are defined in their respective sub packages.
//
// Note: Top level files under providers package only contains the changes that depends on both internal/providers/sdkv2 and internal/providers/pluginfw packages
package providers

import (
	"context"
	"log"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw"
	"github.com/databricks/terraform-provider-databricks/internal/providers/sdkv2"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type serverOptions struct {
	sdkV2Provider           *schema.Provider
	pluginFrameworkProvider provider.Provider
	sdkV2fallbacks          []pluginfw.SdkV2FallbackOption
}

// ServerOption is a common interface for overriding providers in GetProviderServer functino call.
type ServerOption interface {
	Apply(*serverOptions)
}

type sdkV2ProviderOption struct {
	sdkV2Provider *schema.Provider
}

func (o *sdkV2ProviderOption) Apply(options *serverOptions) {
	options.sdkV2Provider = o.sdkV2Provider
}

// WithSdkV2Provider allows overriding the SDKv2 provider used when creating a Terraform provider with muxing.
// This is typically used in acceptance test for a test step to have a custom provider override.
func WithSdkV2Provider(sdkV2Provider *schema.Provider) ServerOption {
	return &sdkV2ProviderOption{sdkV2Provider: sdkV2Provider}
}

type sdkV2FallbackOption struct {
	sdkV2fallbacks []pluginfw.SdkV2FallbackOption
}

func (o *sdkV2FallbackOption) Apply(options *serverOptions) {
	options.sdkV2fallbacks = o.sdkV2fallbacks
}

// WithSdkV2FallbackOptions allows overriding the SDKv2 fallback options used when creating a Terraform provider with muxing.
// This is typically used in acceptance test for testing the compatibility between sdkv2 and plugin framework.
func WithSdkV2FallbackOptions(options []pluginfw.SdkV2FallbackOption) ServerOption {
	return &sdkV2FallbackOption{sdkV2fallbacks: options}
}

// GetProviderServer initializes and returns a Terraform Protocol v6 ProviderServer.
// The function begins by initializing the Databricks provider using the SDK plugin
// and then upgrades this provider to be compatible with Terraform's Protocol v6 using
// the v5-to-v6 upgrade mechanism. Additionally, it retrieves the Databricks provider
// that is implemented using the Terraform Plugin Framework. These different provider
// implementations are then combined using a multiplexing server, which allows multiple
// Protocol v6 providers to be served together. The function returns the multiplexed
// ProviderServer, or an error if any part of the process fails.
//
// GetProviderServer constructs the Databricks Terraform provider server. By default, it combines the default
// SDKv2-based provider and the default plugin framework-based provider using muxing.
// The providers used by the muxed server can be overridden using ServerOptions.
func GetProviderServer(ctx context.Context, options ...ServerOption) (tfprotov6.ProviderServer, error) {
	serverOptions := serverOptions{}
	for _, o := range options {
		o.Apply(&serverOptions)
	}
	sdkPluginProvider := serverOptions.sdkV2Provider
	if sdkPluginProvider == nil {
		sdkPluginProvider = sdkv2.DatabricksProvider(serverOptions.sdkV2fallbacks...)
	}
	pluginFrameworkProvider := serverOptions.pluginFrameworkProvider
	if pluginFrameworkProvider == nil {
		pluginFrameworkProvider = pluginfw.GetDatabricksProviderPluginFramework(serverOptions.sdkV2fallbacks...)
	}

	upgradedSdkPluginProvider, err := tf5to6server.UpgradeServer(
		context.Background(),
		sdkPluginProvider.GRPCProvider,
	)
	if err != nil {
		log.Fatal(err)
	}
	providers := []func() tfprotov6.ProviderServer{
		func() tfprotov6.ProviderServer {
			return upgradedSdkPluginProvider
		},
		providerserver.NewProtocol6(pluginFrameworkProvider),
	}

	muxServer, err := tf6muxserver.NewMuxServer(ctx, providers...)

	if err != nil {
		return nil, err
	}

	return muxServer.ProviderServer(), nil
}
