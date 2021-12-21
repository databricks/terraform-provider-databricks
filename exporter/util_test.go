package exporter

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/clusters"
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
