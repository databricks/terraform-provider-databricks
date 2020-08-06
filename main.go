package main

import (
	"log"

	"github.com/databrickslabs/databricks-terraform/provider"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var (
	// todo: put version into separate package, so that it could be used for UserAgent
	version = "dev"
	commit  = ""
	date    = "unknown"
)

func main() {
	log.SetFlags(0)
	log.Printf("%v, commit %v, built at %v\n", version, commit, date)
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return provider.DatabricksProvider(version)
		},
	})
}
