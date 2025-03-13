package mws_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestMwsAccVpcEndpoint(t *testing.T) {
	t.SkipNow()
	acceptance.AccountLevel(t, acceptance.Step{
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
	acceptance.AccountLevel(t, acceptance.Step{
		Template: `
		resource "databricks_mws_vpc_endpoint" "this" {
			account_id = "{env.DATABRICKS_ACCOUNT_ID}"
			vpc_endpoint_name = "vpce-{var.RANDOM}"
            
            gcp_vpc_endpoint_info {
			  project_id = "{env.GOOGLE_PROJECT}"
			  psc_endpoint_name = "{env.TEST_RELAY_PSC_ENDPOINT}"
			  endpoint_region = "{env.GOOGLE_REGION}"
        }

		}`,
	})
}
