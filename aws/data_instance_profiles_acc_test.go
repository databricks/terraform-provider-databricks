package aws_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccDataSourceInstanceProfiles(t *testing.T) {
	acceptance.GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	acceptance.WorkspaceLevel(t, acceptance.Step{
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
