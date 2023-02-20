package acceptance

import (
	"testing"
)

func TestMwsAccStorageConfigurations(t *testing.T) {
	accountLevel(t, step{
		Template: `
		resource "databricks_mws_storage_configurations" "this" {
			account_id                 = "{env.DATABRICKS_ACCOUNT_ID}"
			storage_configuration_name = "terraform-{var.RANDOM}"
			bucket_name                = "terraform-{var.RANDOM}"
		}`,
	})
}
