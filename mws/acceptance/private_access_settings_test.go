package acceptance

import (
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
)

func TestMwsAccPrivateAccessSettings(t *testing.T) {
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
			resource "databricks_mws_private_access_settings" "my_pas" {
				account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
				private_access_settings_name = "my_pas"
				region = "{env.TEST_REGION}"
			}`,
		},
	})
}
