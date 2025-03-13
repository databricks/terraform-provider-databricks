package mws_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestMwsAccStorageConfigurations(t *testing.T) {
	acceptance.GetEnvOrSkipTest(t, "TEST_ROOT_BUCKET") // marker for AWS test env
	acceptance.AccountLevel(t, acceptance.Step{
		Template: `
		resource "databricks_mws_storage_configurations" "this" {
			account_id                 = "{env.DATABRICKS_ACCOUNT_ID}"
			storage_configuration_name = "terraform-{var.RANDOM}"
			bucket_name                = "terraform-{var.RANDOM}"
		}`,
	})
}
