package acceptance

import (
	"testing"
)

// "databricks_instance_profile" is a singleton. To avoid multiple tests using this resource
// from interfering with each other, we run them in sequence as steps of a single test.
func TestAccInstanceProfileIntegrationSuite(t *testing.T) {
	WorkspaceLevel(t,
		// Assign instance profile to group
		Step{
			Template: `
		resource "databricks_instance_profile" "this" {
			instance_profile_arn = "{env.DUMMY_EC2_INSTANCE_PROFILE}"
		}
		resource "databricks_group" "this" {
			display_name = "tf-{var.RANDOM}"
		}
		resource "databricks_group_instance_profile" "this" {
			group_id = databricks_group.this.id
			instance_profile_id = databricks_instance_profile.this.id
		}`},
		// Assign instance profile to mount
		Step{
			Template: `
			resource "databricks_instance_profile" "this" {
				instance_profile_arn = "{env.DUMMY_EC2_INSTANCE_PROFILE}"
			}
			resource "databricks_mount" "this" {
				name = "experiments"
				cluster_id = "{env.TEST_DEFAULT_CLUSTER_ID}"
				s3 {
					instance_profile = databricks_instance_profile.this.id
					bucket_name      = "{env.TEST_S3_BUCKET}"
				}
			}`,
		},
		// ServicePrincipal resource on Aws with role
		Step{
			Template: `
			resource "databricks_service_principal" "this" {
				display_name = "SPN {var.RANDOM}"
			}
			resource "databricks_instance_profile" "this" {
				instance_profile_arn = "{env.DUMMY_EC2_INSTANCE_PROFILE}"
			}
			resource "databricks_service_principal_role" "this" {
				service_principal_id = databricks_service_principal.this.id
				role                 = databricks_instance_profile.this.id
			  }
			`,
		})
}
