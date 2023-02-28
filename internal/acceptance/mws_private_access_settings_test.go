package acceptance

import (
	"testing"
)

func TestMwsAccPrivateAccessSettings(t *testing.T) {
	t.SkipNow()
	accountLevel(t, step{
		Template: `
		resource "databricks_mws_private_access_settings" "this" {
			account_id = "{env.DATABRICKS_ACCOUNT_ID}"
			private_access_settings_name = "tf-{var.RANDOM}"
			region = "{env.AWS_REGION}"
			public_access_enabled = true
		}`,
	})
}
