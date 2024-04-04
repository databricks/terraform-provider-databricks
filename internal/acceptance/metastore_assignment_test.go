package acceptance

import (
	"context"
	"os"
	"testing"

	"github.com/databricks/databricks-sdk-go/qa/lock"
)

func lockForTest(t *testing.T) func() {
	return func() {
		lock.Acquire(context.Background(), lock.Workspace{WorkspaceId: os.Getenv("DUMMY_WORKSPACE_ID")}, lock.InTest(t))
	}
}

func TestUcAccMetastoreAssignment(t *testing.T) {
	unityWorkspaceLevel(t, step{
		PreConfig: lockForTest(t),
		Template: `resource "databricks_metastore_assignment" "this" {
			metastore_id = "{env.TEST_METASTORE_ID}"
			workspace_id = {env.DUMMY_WORKSPACE_ID}
		}`,
	})
}

func TestUcAccAccountMetastoreAssignment(t *testing.T) {
	unityAccountLevel(t, step{
		PreConfig: lockForTest(t),
		Template: `resource "databricks_metastore_assignment" "this" {
			metastore_id = "{env.TEST_METASTORE_ID}"
			workspace_id = {env.DUMMY_WORKSPACE_ID}
		}`,
	}, step{
		Template: `resource "databricks_metastore_assignment" "this" {
			metastore_id = "{env.TEST_METASTORE_ID}"
			workspace_id = {env.DUMMY2_WORKSPACE_ID}
		}`,
	})
}
