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
	var debug bool
	if len(os.Args) > 1 && os.Args[1] == "debug" {
		debug = true
	}
	log.Printf(`Databricks Terraform Provider
	
	Version %s
	
	https://registry.terraform.io/providers/databricks/databricks/latest/docs
	
	`, common.Version())

	sdkPluginProvider := provider.DatabricksProvider()

	// Translate terraform sdk plugin to protocol 6
	upgradedSdkPluginProvider, err := tf5to6server.UpgradeServer(
		context.Background(),
		sdkPluginProvider.GRPCProvider,
	)

	pluginFrameworkProvider := provider.GetDatabricksProviderPluginFramework()

	providers := []func() tfprotov6.ProviderServer{
		func() tfprotov6.ProviderServer {
			return upgradedSdkPluginProvider
		},
		providerserver.NewProtocol6(pluginFrameworkProvider),
	}

	// Translate plugin framework to protocol 5,
	// we would use tf5muxserver.NewMuxServer(ctx, providers...) and tf5server.Serve below
	// providers := []func() tfprotov5.ProviderServer{
	// 	sdkPluginProvider.GRPCProvider,
	// 	providerserver.NewProtocol5(
	// 		pluginFrameworkProvider,
	// 	),
	// }

	ctx := context.Background()
	muxServer, err := tf6muxserver.NewMuxServer(ctx, providers...)
	if err != nil {
		log.Fatal(err)
	}

	var serveOpts []tf6server.ServeOpt
	if debug {
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
