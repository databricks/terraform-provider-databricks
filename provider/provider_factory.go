package provider

import (
	"context"
	"log"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
)

func GetProviderServer(ctx context.Context) (tfprotov6.ProviderServer, error) {
	sdkPluginProvider := DatabricksProvider()

	upgradedSdkPluginProvider, err := tf5to6server.UpgradeServer(
		context.Background(),
		sdkPluginProvider.GRPCProvider,
	)
	if err != nil {
		log.Fatal(err)
	}

	pluginFrameworkProvider := GetDatabricksProviderPluginFramework()

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

// Remove redundancy once verified that this works
func GetProviderServerWithConfiguredMockClient(ctx context.Context, client *common.DatabricksClient) (tfprotov6.ProviderServer, error) {
	sdkPluginProvider := DatabricksProvider()

	upgradedSdkPluginProvider, err := tf5to6server.UpgradeServer(
		context.Background(),
		sdkPluginProvider.GRPCProvider,
	)
	if err != nil {
		log.Fatal(err)
	}

	pluginFrameworkProvider := GetDatabricksProviderPluginFrameworkWithConfiguredMockClient(client)

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
