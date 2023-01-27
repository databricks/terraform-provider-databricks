package exporter

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/stretchr/testify/assert"
)

func TestImportClusterEmitsInitScripts(t *testing.T) {
	ic := importContextForTest()
	ic.importCluster(&clusters.Cluster{
		InitScripts: []clusters.InitScriptStorageInfo{
			{
				Dbfs: &clusters.DbfsStorageInfo{
					Destination: "/mnt/abc/test.sh",
				},
			},
		},
	})
	assert.Equal(t, 1, len(ic.testEmits))
	assert.True(t, ic.testEmits["databricks_dbfs_file[<unknown>] (id: /mnt/abc/test.sh)"])
}

func TestAddAwsMounts(t *testing.T) {
	ic := importContextForTest()
	ic.mountMap = map[string]mount{}
	ic.addAwsMounts("abc", map[string]string{
		"foo": "bar",
		"baz": "🙄",
	})
	assert.Equal(t, 2, len(ic.mountMap))
}

func TestEmitUserOrServicePrincipal(t *testing.T) {
	ic := importContextForTest()
	assert.True(t, len(ic.testEmits) == 0)

	ic.emitUserOrServicePrincipal("user@domain.com")
	assert.True(t, len(ic.testEmits) == 1)
	assert.True(t, ic.testEmits["databricks_user[<unknown>] (user_name: user@domain.com)"])

	//
	ic = importContextForTest()
	ic.emitUserOrServicePrincipal("21aab5a7-ee70-4385-34d4-a77278be5cb6")
	assert.True(t, len(ic.testEmits) == 1)
	assert.True(t, ic.testEmits["databricks_service_principal[<unknown>] (application_id: 21aab5a7-ee70-4385-34d4-a77278be5cb6)"])

	// unsuccessfull test
	ic = importContextForTest()
	ic.emitUserOrServicePrincipal("abc")
	assert.True(t, len(ic.testEmits) == 0)
}

func TestEmitUserOrServicePrincipalForPath(t *testing.T) {
	ic := importContextForTest()

	ic.emitUserOrServicePrincipalForPath("/Users/user@domain.com/abc", "/Users")
	assert.True(t, len(ic.testEmits) == 1)
	assert.True(t, ic.testEmits["databricks_user[<unknown>] (user_name: user@domain.com)"])

	// Negative cases
	ic = importContextForTest()
	ic.emitUserOrServicePrincipalForPath("/Shared/abc", "/Users")
	assert.True(t, len(ic.testEmits) == 0)

	ic = importContextForTest()
	ic.emitUserOrServicePrincipalForPath("/Users/abc", "/Users")
	assert.True(t, len(ic.testEmits) == 0)
	ic = importContextForTest()
	ic.emitUserOrServicePrincipalForPath("/Users/", "/Users")
	assert.True(t, len(ic.testEmits) == 0)
}

func TestEmitNotebookOrRepo(t *testing.T) {
	ic := importContextForTest()
	ic.emitNotebookOrRepo("/Users/user@domain.com/abc")
	assert.True(t, len(ic.testEmits) == 1)
	assert.True(t, ic.testEmits["databricks_notebook[<unknown>] (id: /Users/user@domain.com/abc)"])

	// test for repository
	ic = importContextForTest()
	ic.emitNotebookOrRepo("/Repos/user@domain.com/repo/abc")
	assert.True(t, len(ic.testEmits) == 1)
	assert.True(t, ic.testEmits["databricks_repo[<unknown>] (path: /Repos/user@domain.com/repo)"])
}
