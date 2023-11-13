package acceptance

import (
	"testing"
)

func TestAccDataSourceInstanceProfiles(t *testing.T) {
	GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	workspaceLevel(t, step{
		Template: `
		resource "databricks_instance_profile" "this" {
			instance_profile_arn = "{env.TEST_EC2_INSTANCE_PROFILE}"
		}

  		# Delay the data call, otherwise it will return null
  		resource "time_sleep" "wait" {
  			depends_on = [databricks_instance_profile.this]

  			create_duration = "10s"
		}

		data "databricks_instance_profiles" "this" {
			depends_on = [
				databricks_instance_profile.this,
    				time_sleep.wait,
			]
		}
		
		output "instance_profiles" {
			value = data.databricks_instance_profiles.this.instance_profiles
		}

		output "instance_profile" {
			value = data.databricks_instance_profiles.this.instance_profiles[0]
		}`,
	})
}
