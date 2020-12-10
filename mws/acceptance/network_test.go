package acceptance

import (
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
)

func TestMwsAccNetworks(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv != "MWS" {
		t.Skip("Cannot run test on non-MWS environment")
	}
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `provider "databricks" {
				host     = "{env.DATABRICKS_HOST}"
				username = "{env.DATABRICKS_USERNAME}"
				password = "{env.DATABRICKS_PASSWORD}"
			}
			resource "databricks_mws_networks" "my_network" {
				account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
				network_name = "network-test-{var.RANDOM}"
				vpc_id       = "vpc-11111111"
				subnet_ids   = [
					"subnet-11111111",
					"subnet-99999999"
				]
				security_group_ids = [
					"sg-99999999"
				]
			}`,
		},
	})
}
