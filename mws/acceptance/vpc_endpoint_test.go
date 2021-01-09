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
			Template: `provider "databricks" {
				host     = "{env.DATABRICKS_HOST}"
				username = "{env.DATABRICKS_USERNAME}"
				password = "{env.DATABRICKS_PASSWORD}"
			}
			resource "databricks_mws_vpc_endpoint" "my_vpc_endpoint" {
				account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
				vpc_endpoint_name = "my_vpce_name"
				region = "{env.TEST_REGION}"
				aws_vpc_endpoint_id       = "{env.AWS_VPC_RELAY_ENDPOINT_ID}"
			}`,
		},
	})
}
