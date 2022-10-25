package acceptance

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestUcAccMetastoreDataAccessOnAws(t *testing.T) {
	qa.RequireCloudEnv(t, "ucws")
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			resource "databricks_metastore" "this" {
			  name          = "primary"
			  storage_root  = "s3://{env.TEST_BUCKET}/test{var.RANDOM}"
			  force_destroy = true
			}
			
			resource "databricks_metastore_data_access" "this" {
			  metastore_id = databricks_metastore.this.id
			  name         = "{var.RANDOM}"
			  aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			  }
			  is_default = true
			}`,
		},
	})
}
