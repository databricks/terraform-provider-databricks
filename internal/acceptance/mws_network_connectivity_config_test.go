package acceptance

import (
	"os"
	"testing"
)

func TestMwsAccNetworkConnectivityConfig(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	switch cloudEnv {
	case "azure-ucacct":
		accountLevel(t, step{
			Template: `
			resource "databricks_mws_network_connectivity_config" "this" {
				account_id = "{env.DATABRICKS_ACCOUNT_ID}"
				name = "tf-{var.RANDOM}"
				region = "eastus2"
			}

			resource "databricks_mws_ncc_private_endpoint_rule" "this" {
				network_connectivity_config_id = databricks_mws_network_connectivity_config.this.id
				resource_id = "/subscriptions/2a5a4578-9ca9-47e2-ba46-f6ee6cc731f2/resourceGroups/deco-prod-azure-eastus2-rg/providers/Microsoft.Storage/storageAccounts/decotestprodunity"
				group_id = "blob"
			}
			`,
		}, step{
			Template: `
			resource "databricks_mws_network_connectivity_config" "this" {
				account_id = "{env.DATABRICKS_ACCOUNT_ID}"
				name = "tf-{var.RANDOM}"
				region = "eastus2"
			}

			resource "databricks_mws_ncc_private_endpoint_rule" "this" {
				network_connectivity_config_id = databricks_mws_network_connectivity_config.this.id
				resource_id = "/subscriptions/2a5a4578-9ca9-47e2-ba46-f6ee6cc731f2/resourceGroups/deco-prod-azure-eastus2-rg/providers/Microsoft.Storage/storageAccounts/decotestprodunity"
				group_id = "blob"
			}
			`,
		})
	case "mws":
		accountLevel(t, step{
			Template: `
			resource "databricks_mws_network_connectivity_config" "this" {
				account_id = "{env.DATABRICKS_ACCOUNT_ID}"
				name = "tf-{var.RANDOM}"
				region = "{env.AWS_REGION}"
			}
			`,
		}, step{
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
