package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/exporter"
	"github.com/databricks/terraform-provider-databricks/provider"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server"
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

	ctx := context.Background()
	providerServer, err := provider.GetProviderServer(ctx)
	if err != nil {
		log.Fatal(err)
	}

	var serveOpts []tf6server.ServeOpt
	if len(os.Args) > 1 && os.Args[1] == "debug" { // debug mode
		serveOpts = append(serveOpts, tf6server.WithManagedDebug())
	}

	err = tf6server.Serve(
		"registry.terraform.io/databricks/databricks",
		func() tfprotov6.ProviderServer { return providerServer },
		serveOpts...,
	)

	if err != nil {
		log.Fatal(err)
	}
}
