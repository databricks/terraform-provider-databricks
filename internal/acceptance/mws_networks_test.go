package acceptance

import (
	"testing"
)

func TestMwsAccNetworks(t *testing.T) {
	GetEnvOrSkipTest(t, "TEST_ROOT_BUCKET") // marker for AWS test env
	accountLevel(t, step{
		Template: `
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
	})
}
