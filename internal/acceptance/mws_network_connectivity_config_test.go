package acceptance

import (
	"os"
	"testing"
)

// Acceptance test for Network Connectivity Config on Azure.
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
			}`,
		})
	default:
		t.Skipf("not available on %s", cloudEnv)
	}
}

// TODO: Add Acceptance test for AWS NCC.
