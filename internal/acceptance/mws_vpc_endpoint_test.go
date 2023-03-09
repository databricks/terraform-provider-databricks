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

func TestMwsAccVpcEndpoint_GCP(t *testing.T) {
	accountLevel(t, step{
		Template: `
		resource "databricks_mws_vpc_endpoint" "this" {
			account_id = "{env.DATABRICKS_ACCOUNT_ID}"
			vpc_endpoint_name = "{env.TEST_PREFIX}-{var.RANDOM}"
            
            gcp_vpc_endpoint_info {
			  project_id = "{env.GOOGLE_PROJECT}"
			  psc_endpoint_name = "{var.RANDOM}"
			  endpoint_region = "env.GOOGLE_REGION"
        }

		}`,
	})
}
