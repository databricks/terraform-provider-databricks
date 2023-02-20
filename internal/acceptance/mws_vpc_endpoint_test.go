package acceptance

import (
	"testing"
)

func TestMwsAccVpcEndpoint(t *testing.T) {
	t.SkipNow()
	accountLevel(t, step{
		Template: `
		resource "databricks_mws_vpc_endpoint" "this" {
			account_id = "{env.DATABRICKS_ACCOUNT_ID}"
			vpc_endpoint_name = "tf-{var.RANDOM}"
			region = "{env.AWS_REGION}"
			aws_vpc_endpoint_id = "{env.TEST_RELAY_VPC_ENDPOINT}"
			aws_account_id = "{env.AWS_ACCOUNT_ID}"
		}`,
	})
}
