package acceptance

import (
	"testing"
)

func TestUcAccMetastoreAssignment(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `resource "databricks_metastore_assignment" "this" {
			metastore_id = "{env.TEST_METASTORE_ID}"
			workspace_id = {env.TEST_WORKSPACE_ID}
		}`,
	})
}

func TestUcAccAccMetastore(t *testing.T) {
	unityAccountLevel(t, step{
		Template: `resource "databricks_metastore_assignment" "this" {
			metastore_id = "{env.TEST_METASTORE_ID}"
			workspace_id = {env.TEST_WORKSPACE_ID}
		}`,
	})
}
