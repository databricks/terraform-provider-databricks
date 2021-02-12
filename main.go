package main

import (
	"fmt"
	"log"
	"os"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/importer"
	"github.com/databrickslabs/terraform-provider-databricks/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	log.SetFlags(0)
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println(common.Version())
		return
	}
	if len(os.Args) > 1 && os.Args[1] == "importer" {
		if err := importer.Run(os.Args...); err != nil {
			log.Printf("[ERROR] %s", err.Error())
			os.Exit(1)
		}
		return
	}
	log.Printf(`Databricks Terraform Provider (experimental)

Version %s

https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs

`, common.Version())
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: provider.DatabricksProvider})
}
