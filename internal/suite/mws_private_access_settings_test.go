package acceptance

import (
	"os"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestMwsAccPrivateAccessSettings(t *testing.T) {
	t.SkipNow()
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv != "MWS" {
		t.Skip("Cannot run test on non-MWS environment")
	}
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			resource "databricks_mws_private_access_settings" "this" {
				account_id = "{env.DATABRICKS_ACCOUNT_ID}"
				private_access_settings_name = "tf-{var.RANDOM}"
				region = "{env.AWS_REGION}"
				public_access_enabled = true
			}`,
		},
	})
}
