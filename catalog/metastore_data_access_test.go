package catalog_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestUcAccAccountMetastoreDataAccessOnAws(t *testing.T) {
	acceptance.UnityAccountLevel(t, acceptance.Step{
		Template: `
		resource "databricks_metastore" "this" {
			name          = "primary-{var.RANDOM}"
			storage_root  = "s3://{env.TEST_BUCKET}/test{var.RANDOM}"
			region        = "us-east-1"
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

func TestUcAccMetastoreDataAccessOnAws(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: `
		resource "databricks_metastore_data_access" "this" {
			metastore_id = "{env.TEST_METASTORE_ID}"
			name         = "{var.RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
		}`,
	})
}
