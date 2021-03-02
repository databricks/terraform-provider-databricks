package storage

import (
	"strings"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/compute"
	"github.com/databrickslabs/terraform-provider-databricks/internal"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAzureAccADLSv1Mount(t *testing.T) {
	client, mp := mountPointThroughReusedCluster(t)
	if !client.AzureAuth.IsClientSecretSet() {
		t.Skip("Test is meant only for client-secret conf Azure")
	}
	storageResource := qa.GetEnvOrSkipTest(t, "TEST_DATA_LAKE_STORE_NAME")
	testWithNewSecretScope(t, func(scope, key string) {
		testMounting(t, mp, AzureADLSGen1Mount{
			ClientID:        client.AzureAuth.ClientID,
			TenantID:        client.AzureAuth.TenantID,
			PrefixType:      "dfs.adls",
			StorageResource: storageResource,
			Directory:       "/",
			SecretScope:     scope,
			SecretKey:       key,
		})
	}, client, mp.name, client.AzureAuth.ClientSecret)
}

func TestResourceAdlsGen1Mount_Create(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceAzureAdlsGen1Mount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, "adl://test-adls.azuredatalakestore.net")
				assert.Contains(t, trunc, `"fs.adl.oauth2.credential":dbutils.secrets.get("c", "d")`)
			}
			assert.Contains(t, trunc, "/mnt/this_mount")
			return common.CommandResults{
				ResultType: "text",
				Data:       testS3BucketPath,
			}
		},
		State: map[string]interface{}{
			"cluster_id":            "this_cluster",
			"mount_name":            "this_mount",
			"storage_resource_name": "test-adls",
			"tenant_id":             "a",
			"client_id":             "b",
			"client_secret_scope":   "c",
			"client_secret_key":     "d",
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, testS3BucketPath, d.Get("source"))
}
