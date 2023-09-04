package acceptance

import "testing"

func TestAccCreateDatabricksMount(t *testing.T) {
	workspaceLevel(t,
		step{
			Template: `
				data "databricks_spark_version" "latest" {}
			  
				# Test cluster to create the mount using.
				resource "databricks_cluster" "this" {
					cluster_name = "acc-test-mounts-{var.RANDOM}"
					spark_version = data.databricks_spark_version.latest.id
					instance_pool_id = "{env.TEST_INSTANCE_POOL_ID}"
					num_workers = 1

					aws_attributes {
					instance_profile_arn = "{env.TEST_INSTANCE_PROFILE_ARN}"
					}
				}

				resource "databricks_mount" "my_mount" {
					name = "test-mount-{var.RANDOM}"
					cluster_id = databricks_cluster.this.id
					
					s3 {
						bucket_name      = "{env.TEST_S3_BUCKET_NAME}"
					}
				}`,
		})
}
