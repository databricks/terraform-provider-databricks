package mws_private_access_settings_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestMwsAccDataSourcePrivateAccessSettingsById(t *testing.T) {
	acceptance.AccountLevel(t, acceptance.Step{
		Template: `
		resource "databricks_mws_private_access_settings" "this" {
			private_access_settings_name = "tf-pas-{var.STICKY_RANDOM}"
			region = "{env.AWS_REGION}"
			public_access_enabled = true
		}

		data "databricks_mws_private_access_settings" "by_id" {
			private_access_settings_id = databricks_mws_private_access_settings.this.private_access_settings_id
		}`,
	})
}

func TestMwsAccDataSourcePrivateAccessSettingsByName(t *testing.T) {
	acceptance.AccountLevel(t, acceptance.Step{
		Template: `
		resource "databricks_mws_private_access_settings" "this" {
			private_access_settings_name = "tf-pas-{var.STICKY_RANDOM}"
			region = "{env.AWS_REGION}"
			public_access_enabled = true
		}

		data "databricks_mws_private_access_settings" "by_name" {
			private_access_settings_name = databricks_mws_private_access_settings.this.private_access_settings_name
		}`,
	})
}
