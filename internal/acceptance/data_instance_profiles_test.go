package acceptance

import (
	"testing"
)

func TestAccDataSourceInstanceProfiles(t *testing.T) {
	GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	workspaceLevel(t, LegacyStep{
		Template: `
		data "databricks_instance_profiles" "this" {
		}
		
		output "instance_profiles" {
			value = data.databricks_instance_profiles.this.instance_profiles
		}

		output "instance_profile" {
			value = data.databricks_instance_profiles.this.instance_profiles[0]
		}`,
	})
}
