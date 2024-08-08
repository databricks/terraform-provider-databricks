package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
)

func GetProviderServer() []func() tfprotov6.ProviderServer {
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

	return providers
}
