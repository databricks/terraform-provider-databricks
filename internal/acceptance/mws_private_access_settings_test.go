package acceptance

import (
	"testing"
)

func TestMwsAccPrivateAccessSettings(t *testing.T) {
	t.SkipNow()
	accountLevel(t, LegacyStep{
		Template: `
		resource "databricks_mws_private_access_settings" "this" {
			account_id = "{env.DATABRICKS_ACCOUNT_ID}"
			private_access_settings_name = "tf-{var.RANDOM}"
			region = "{env.AWS_REGION}"
			public_access_enabled = true
		}`,
	})
}

func TestMwsGcpAccPrivateAccessSettings(t *testing.T) {
	t.Skipf("skipping until feature is disabled")
	accountLevel(t, LegacyStep{
		Template: `
		resource "databricks_mws_private_access_settings" "this" {
			account_id = "{env.DATABRICKS_ACCOUNT_ID}"
			private_access_settings_name = "tf-{var.RANDOM}"
			region = "{env.GOOGLE_REGION}"
			public_access_enabled = true
			private_access_level = "ACCOUNT"
		}`,
	})
}
