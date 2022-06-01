package storage

import (
	"strings"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/clusters"
	"github.com/databrickslabs/terraform-provider-databricks/commands"
	"github.com/databrickslabs/terraform-provider-databricks/common"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResourceAdlsGen2Mount_Create(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceAzureAdlsGen2Mount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, "abfss://e@test-adls-gen2.dfs.core.windows.net")
				assert.Contains(t, trunc, `"fs.azure.account.oauth2.client.secret":dbutils.secrets.get("c", "d")`)
			}
			assert.Contains(t, trunc, "/mnt/this_mount")
			return common.CommandResults{
				ResultType: "text",
				Data:       "abfss://e@test-adls-gen2.dfs.core.windows.net",
			}
		},
		State: map[string]interface{}{
			"cluster_id":             "this_cluster",
			"container_name":         "e",
			"mount_name":             "this_mount",
			"storage_account_name":   "test-adls-gen2",
			"tenant_id":              "a",
			"client_id":              "b",
			"client_secret_scope":    "c",
			"client_secret_key":      "d",
			"initialize_file_system": true,
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, "abfss://e@test-adls-gen2.dfs.core.windows.net", d.Get("source"))
}
