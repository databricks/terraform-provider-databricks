package acceptance

import (
	"testing"
)

func TestUcAccMetastoreAssignment(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `resource "databricks_metastore_assignment" "this" {
			metastore_id = "{env.TEST_METASTORE_ID}"
			workspace_id = {env.DUMMY_WORKSPACE_ID}
		}`,
	})
}

func TestUcAccAccountMetastoreAssignment(t *testing.T) {
	unityAccountLevel(t, step{
		Template: `resource "databricks_metastore_assignment" "this" {
			metastore_id = "{env.TEST_METASTORE_ID}"
			workspace_id = {env.DUMMY_WORKSPACE_ID}
		}`,
	})
}
