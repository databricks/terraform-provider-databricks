package main

import (
	"log"

	"github.com/databrickslabs/databricks-terraform/databricks"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var (
	version = "dev"
	commit  = ""
	date    = "unknown"
)

func main() {
	log.Printf("%v, commit %v, built at %v\n", version, commit, date)
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return databricks.Provider(version)
		},
	})
}
