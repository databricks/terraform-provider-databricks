package mws_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestMwsAccAwsCustomerManagedKeys(t *testing.T) {
	acceptance.AccountLevel(t, acceptance.Step{
		Template: `resource "databricks_mws_customer_managed_keys" "this" {
			account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
			aws_key_info {
				key_arn   = "{env.TEST_MANAGED_KMS_KEY_ARN}"
				key_alias = "{env.TEST_MANAGED_KMS_KEY_ALIAS}"
			}
			use_cases = ["MANAGED_SERVICES"]
		}`,
	})
}

func TestMwsAccGcpCustomerManagedKeysForStorage(t *testing.T) {
	acceptance.AccountLevel(t, acceptance.Step{
		Template: `resource "databricks_mws_customer_managed_keys" "this" {
				account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
				gcp_key_info {
					kms_key_id   = "{env.TEST_GCP_KMS_KEY_ID}"
				}
				use_cases = ["STORAGE"]
			}`,
	})
}
