package acceptance

import (
	"os"

	"github.com/databrickslabs/databricks-terraform/internal/acceptance"

	"testing"
)

func TestMwsAccCredentials(t *testing.T) {
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
			resource "databricks_mws_credentials" "this" {
				account_id       = "{env.DATABRICKS_ACCOUNT_ID}"
				credentials_name = "creds-test-{var.RANDOM}"
				role_arn         = "arn:aws:iam::999999999999:role/tf-test-{var.RANDOM}"
			}`,
		},
	})
}
