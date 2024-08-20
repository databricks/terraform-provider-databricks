package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/exporter"
	"github.com/databricks/terraform-provider-databricks/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
)

const startMessageFormat = `Databricks Terraform Provider
	
Version %s
	
https://registry.terraform.io/providers/databricks/databricks/latest/docs
	
`

func main() {
	log.SetFlags(0)
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println(common.Version())
		return
	}
	if len(os.Args) > 1 && os.Args[1] == "exporter" {
		if err := exporter.Run(os.Args...); err != nil {
			log.Printf("[ERROR] %s", err.Error())
			os.Exit(1)
		}
		return
	}

	log.Printf(startMessageFormat, common.Version())

	sdkPluginProvider := provider.DatabricksProvider()

	upgradedSdkPluginProvider, err := tf5to6server.UpgradeServer(
		context.Background(),
		sdkPluginProvider.GRPCProvider,
	)
	if err != nil {
		log.Fatal(err)
	}
	pluginFrameworkProvider := provider.GetDatabricksProviderPluginFramework()

	providers := []func() tfprotov6.ProviderServer{
		func() tfprotov6.ProviderServer {
			return upgradedSdkPluginProvider
		},
		providerserver.NewProtocol6(pluginFrameworkProvider),
	}

	ctx := context.Background()
	muxServer, err := tf6muxserver.NewMuxServer(ctx, providers...)

	if err != nil {
		log.Fatal(err)
	}

	var serveOpts []tf6server.ServeOpt
	if len(os.Args) > 1 && os.Args[1] == "debug" { // debug mode
		serveOpts = append(serveOpts, tf6server.WithManagedDebug())
	}

	err = tf6server.Serve(
		"registry.terraform.io/databricks/databricks",
		muxServer.ProviderServer,
		serveOpts...,
	)

	if err != nil {
		log.Fatal(err)
	}
}
