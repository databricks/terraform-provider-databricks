package provider

import (
	"context"
	"log"

	pluginframeworkprovider "github.com/databricks/terraform-provider-databricks/internal/pluginframework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
)

// GetProviderServer initializes and returns a Terraform Protocol v6 ProviderServer.
// The function begins by initializing the Databricks provider using the SDK plugin
// and then upgrades this provider to be compatible with Terraform's Protocol v6 using
// the v5-to-v6 upgrade mechanism. Additionally, it retrieves the Databricks provider
// that is implemented using the Terraform Plugin Framework. These different provider
// implementations are then combined using a multiplexing server, which allows multiple
// Protocol v6 providers to be served together. The function returns the multiplexed
// ProviderServer, or an error if any part of the process fails.
func GetProviderServer(ctx context.Context) (tfprotov6.ProviderServer, error) {
	sdkPluginProvider := DatabricksProvider()

	upgradedSdkPluginProvider, err := tf5to6server.UpgradeServer(
		context.Background(),
		sdkPluginProvider.GRPCProvider,
	)
	if err != nil {
		log.Fatal(err)
	}

	pluginFrameworkProvider := pluginframeworkprovider.GetDatabricksProviderPluginFramework()

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
