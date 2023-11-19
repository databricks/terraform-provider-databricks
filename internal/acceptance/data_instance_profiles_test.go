package acceptance

import (
	"testing"
)

func TestAccDataSourceInstanceProfiles(t *testing.T) {
	GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	workspaceLevel(t, step{
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

func TestAccDataSourceInstanceProfilesFilter(t *testing.T) {
	GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	workspaceLevel(t, step{
		Template: `
		data "databricks_instance_profiles" "this" {
			filter {
				name = "arn"
				pattern = "^{TEST_EC2_INSTANCE_PROFILE}$"
			}
		}

		output "instance_profile" {
			value = data.databricks_instance_profiles.this.instance_profiles[0]
		}`,
	})
}
