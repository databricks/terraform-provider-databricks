package acceptance

import (
	"os"
	"testing"
)

func TestAzureAccNetworkConnectivityConfig(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	switch cloudEnv {
	case "azure":
		accountLevel(t, step{
			Template: `
			resource "databricks_mws_network_connectivity_config" "this" {
				account_id = "{env.DATABRICKS_ACCOUNT_ID}"
				name = "tf-{var.RANDOM}"
				region = "eastus2"
			}

			resource "databricks_mws_ncc_private_endpoint_rule" "this" {
				network_connectivity_config_id = databricks_mws_network_connectivity_config.this.id
				resource_id = "/subscriptions/653bb673-1234-abcd-a90b-d064d5d53ca4/resourcegroups/example-resource-group/providers/Microsoft.Storage/storageAccounts/examplesa"
				group_id = "blob"
			}
			`,
		})
	case "aws":
		accountLevel(t, step{
			Template: `
			resource "databricks_mws_network_connectivity_config" "this" {
				account_id = "{env.DATABRICKS_ACCOUNT_ID}"
				name = "tf-{var.RANDOM}"
				region = "{env.AWS_REGION}"
			}
			`,
		})
	}
}
