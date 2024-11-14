package acceptance

import (
	"testing"
)

func TestMwsAccCredentials(t *testing.T) {
	AccountLevel(t, Step{
		Template: `resource "databricks_mws_credentials" "this" {
			account_id       = "{env.DATABRICKS_ACCOUNT_ID}"
			credentials_name = "creds-test-{var.RANDOM}"
			role_arn         = "{env.TEST_CROSSACCOUNT_ARN}"
		}`,
	})
}
