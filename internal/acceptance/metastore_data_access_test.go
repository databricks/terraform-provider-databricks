package acceptance

import (
	"testing"
)

func TestUcAccMetastoreDataAccessOnAws(t *testing.T) {
	unityAccountLevel(t, step{
		Template: `
		resource "databricks_metastore" "this" {
			name          = "primary-{var.RANDOM}"
			storage_root  = "s3://{env.TEST_BUCKET}/test{var.RANDOM}"
			force_destroy = true
		}
		
		resource "databricks_metastore_data_access" "this" {
			metastore_id = databricks_metastore.this.id
			name         = "{var.RANDOM}"
			aws_iam_role {
			role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
		}`,
	})
}
