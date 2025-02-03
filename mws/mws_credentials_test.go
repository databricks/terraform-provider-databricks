package mws_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestMwsAccCredentials(t *testing.T) {
	acceptance.AccountLevel(t, acceptance.Step{
		Template: `resource "databricks_mws_credentials" "this" {
			account_id       = "{env.DATABRICKS_ACCOUNT_ID}"
			credentials_name = "creds-test-{var.RANDOM}"
			role_arn         = "{env.TEST_CROSSACCOUNT_ARN}"
		}`,
	})
}
