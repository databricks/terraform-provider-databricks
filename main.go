package main

import (
	"fmt"
	"log"
	"os"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	log.SetFlags(0)
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println(common.Version())
		return
	}
	log.Printf(`Databricks Terraform Provider (experimental)

Version %s

https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs

`, common.Version())
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: provider.DatabricksProvider})
}
