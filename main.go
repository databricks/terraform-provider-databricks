package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/exporter"
	"github.com/databrickslabs/terraform-provider-databricks/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
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
	if len(os.Args) > 1 && os.Args[1] == "debug" {
		err := plugin.Debug(context.Background(), "registry.terraform.io/databrickslabs/databricks",
			&plugin.ServeOpts{ProviderFunc: provider.DatabricksProvider})
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}
	log.Printf(`Databricks Terraform Provider (experimental)

Version %s

https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs

`, common.Version())
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: provider.DatabricksProvider})
}
