package mws_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestMwsAccNetworkConnectivityConfig(t *testing.T) {
	if acceptance.IsAzure(t) {
		acceptance.AccountLevel(t, acceptance.Step{
			Template: `
			resource "databricks_mws_network_connectivity_config" "this" {
				name = "tf-{var.RANDOM}"
				region = "eastus2"
			}

			resource "databricks_mws_ncc_private_endpoint_rule" "this" {
				network_connectivity_config_id = databricks_mws_network_connectivity_config.this.network_connectivity_config_id
				resource_id = "/subscriptions/2a5a4578-9ca9-47e2-ba46-f6ee6cc731f2/resourceGroups/deco-prod-azure-eastus2-rg/providers/Microsoft.Storage/storageAccounts/decotestprodunity"
				group_id = "blob"
			}
			`,
		}, acceptance.Step{
			Template: `
			resource "databricks_mws_network_connectivity_config" "this" {
				name = "tf-{var.RANDOM}"
				region = "eastus2"
			}

			resource "databricks_mws_ncc_private_endpoint_rule" "this" {
				network_connectivity_config_id = databricks_mws_network_connectivity_config.this.network_connectivity_config_id
				resource_id = "/subscriptions/2a5a4578-9ca9-47e2-ba46-f6ee6cc731f2/resourceGroups/deco-prod-azure-eastus2-rg/providers/Microsoft.Storage/storageAccounts/decotestprodunity"
				group_id = "blob"
			}
			`,
		})
	}
	if acceptance.IsAws(t) {
		acceptance.AccountLevel(t, acceptance.Step{
			Template: `
			resource "databricks_mws_network_connectivity_config" "this" {
				account_id = "{env.DATABRICKS_ACCOUNT_ID}"
				name = "tf-{var.RANDOM}"
				region = "{env.AWS_REGION}"
			}
			`,
		}, acceptance.Step{
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
