package acceptance

import (
	"testing"
)

func TestAccServicePrincipalResourceOnAwsWithRole(t *testing.T) {
	GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	awsSpn := `resource "databricks_service_principal" "this" {
		display_name = "SPN {var.RANDOM}"
	}
	resource "databricks_instance_profile" "this" {
		instance_profile_arn = "{env.TEST_EC2_INSTANCE_PROFILE}"
	}
	resource "databricks_service_principal_role" "this" {
		service_principal_id = databricks_service_principal.this.id
		role                 = databricks_instance_profile.this.id
	  }
	`
	workspaceLevel(t, step{
		Template: awsSpn,
	})
}
