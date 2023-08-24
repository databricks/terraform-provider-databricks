package acceptance

import (
	"testing"
)

func TestUcAccMetastoreDataAccessOnAws(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_metastore_data_access" "this" {
			metastore_id = {env.TEST_METASTORE_ID}
			name         = "{var.RANDOM}"
			aws_iam_role {
			role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			is_default = true
		}`,
	})
}
