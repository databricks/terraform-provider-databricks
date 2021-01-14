package acceptance

import (
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
)

func TestMwsAccVpcEndpoint(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv != "MWS" {
		t.Skip("Cannot run test on non-MWS environment")
	}
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `resource "databricks_mws_vpc_endpoint" "my_vpc_endpoint_relay" {
				account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
				vpc_endpoint_name = "my_t2_vpce_relay_name"
				region = "{env.TEST_REGION}"
				aws_vpc_endpoint_id       = "{env.AWS_VPC_RELAY_ENDPOINT_ID}"
				aws_account_id = "997819012307"
			}`,
		},
	})
}
