package acceptance

import (
	"os"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"
)

func TestMwsAccCustomerManagedKeys(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv != "MWS" {
		t.Skip("Acceptance tests skipped unless CLOUD_ENV=MWS is set")
	}
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `provider "databricks" {
				host     = "{env.DATABRICKS_HOST}"
				username = "{env.DATABRICKS_USERNAME}"
				password = "{env.DATABRICKS_PASSWORD}"
			}
			resource "databricks_mws_customer_managed_keys" "this" {
				account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
				aws_key_info {
					key_arn   = "{env.TEST_KMS_KEY_ARN}"
					key_alias = "{env.TEST_KMS_KEY_ALIAS}"
				}
			}`,
		},
	})
}
