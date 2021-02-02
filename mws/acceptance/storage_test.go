package acceptance

import (
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
)

func TestMwsAccStorageConfigurations(t *testing.T) {
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
			resource "databricks_mws_storage_configurations" "this" {
				account_id                 = "{env.DATABRICKS_ACCOUNT_ID}"
				storage_configuration_name = "terraform-{var.RANDOM}"
				bucket_name                = "terraform-{var.RANDOM}"
			}`,
		},
	})
}
