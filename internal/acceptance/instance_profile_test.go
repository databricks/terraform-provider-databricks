package acceptance

import (
	"testing"
)

func TestAccInstanceProfileAssignedToGroupAndMounts(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		resource "databricks_instance_profile" "this" {
			instance_profile_arn = "{env.TEST_EC2_INSTANCE_PROFILE}"
		}
		resource "databricks_group" "this" {
			display_name = "tf-{var.RANDOM}"
		}
		resource "databricks_group_instance_profile" "this" {
			group_id = databricks_group.this.id
			instance_profile_id = databricks_instance_profile.this.id
		}
		resource "databricks_mount" "this" {
			name = "experiments"
			cluster_id = "{env.TEST_DEFAULT_CLUSTER_ID}"
			s3 {
				instance_profile = databricks_instance_profile.this.id
				bucket_name      = "{env.TEST_S3_BUCKET}"
			}
		}`,
	})
}
