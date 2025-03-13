package mws_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestMwsAccNetworks(t *testing.T) {
	acceptance.GetEnvOrSkipTest(t, "TEST_ROOT_BUCKET") // marker for AWS test env
	acceptance.AccountLevel(t, acceptance.Step{
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

func TestMwsAccGcpPscNetworks(t *testing.T) {
	acceptance.AccountLevel(t, acceptance.Step{
		Template: `
		resource "databricks_mws_networks" "my_network" {
			account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
			network_name = "network-test-{var.RANDOM}"
			gcp_network_info {
			  network_project_id = "{env.GOOGLE_PROJECT}"
			  vpc_id = "{env.VPC_NETWORK_ID}"
			  subnet_id = "{env.SUBNET_ID}"
			  subnet_region = "{env.GOOGLE_REGION}"
			  pod_ip_range_name = "{env.POD_IP_RANGE_NAME}"
			  service_ip_range_name = "{env.SVC_IP_RANGE_NAME}"
            }
            vpc_endpoints {
        	  rest_api = ["{env.REST_API_VPCE_ID}"]
        	  dataplane_relay = ["{env.RELAY_VPCE_ID}"]
            }
		}`,
	})
}
