package main

import (
	"log"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	log.SetFlags(0)
	log.Printf(`Databricks Terraform Provider (experimental)

Version %s

https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs

`, common.Version())
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: provider.DatabricksProvider})
}
