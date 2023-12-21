package acceptance

import (
	"sync"
	"testing"
)

var dummyWsMetastoreMutex sync.Mutex

func TestUcAccMetastoreAssignment(t *testing.T) {
	dummyWsMetastoreMutex.Lock()
	defer dummyWsMetastoreMutex.Unlock()
	unityWorkspaceLevel(t, step{
		Template: `resource "databricks_metastore_assignment" "this" {
			metastore_id = "{env.TEST_METASTORE_ID}"
			workspace_id = {env.DUMMY_WORKSPACE_ID}
		}`,
	})
}

func TestUcAccAccountMetastoreAssignment(t *testing.T) {
	dummyWsMetastoreMutex.Lock()
	defer dummyWsMetastoreMutex.Unlock()
	unityAccountLevel(t, step{
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
