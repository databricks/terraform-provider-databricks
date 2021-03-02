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

func TestResourceAzureBlobMountCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceAzureBlobMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)

			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, "wasbs://c@f.blob.core.windows.net/d")
				assert.Contains(t, trunc, `"fs.azure.account.key.f.blob.core.windows.net":dbutils.secrets.get("h", "g")`)
			}
			assert.Contains(t, trunc, "/mnt/e")
			return common.CommandResults{
				ResultType: "text",
				Data:       "wasbs://c@f.blob.core.windows.net/d",
			}
		},
		State: map[string]interface{}{
			"auth_type":            "ACCESS_KEY",
			"cluster_id":           "b",
			"container_name":       "c",
			"directory":            "/d",
			"mount_name":           "e",
			"storage_account_name": "f",
			"token_secret_key":     "g",
			"token_secret_scope":   "h",
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err, err) // TODO: global search-replace for NoError
	assert.Equal(t, "e", d.Id())
	assert.Equal(t, "wasbs://c@f.blob.core.windows.net/d", d.Get("source"))
}

func TestResourceAzureBlobMountCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceAzureBlobMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			return common.CommandResults{
				ResultType: "error",
				Summary:    "Some error",
			}
		},
		State: map[string]interface{}{
			"auth_type":            "ACCESS_KEY",
			"cluster_id":           "b",
			"container_name":       "c",
			"directory":            "/d",
			"mount_name":           "e",
			"storage_account_name": "f",
			"token_secret_key":     "g",
			"token_secret_scope":   "h",
		},
		Create: true,
	}.Apply(t)
	require.EqualError(t, err, "Some error")
	assert.Equal(t, "e", d.Id())
	assert.Equal(t, "", d.Get("source"))
}

func TestResourceAzureBlobMountRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceAzureBlobMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			assert.Contains(t, trunc, "dbutils.fs.mounts()")
			assert.Contains(t, trunc, `mount.mountPoint == "/mnt/e"`)
			return common.CommandResults{
				ResultType: "text",
				Data:       "wasbs://c@f.blob.core.windows.net/d",
			}
		},
		State: map[string]interface{}{
			"auth_type":            "ACCESS_KEY",
			"cluster_id":           "b",
			"container_name":       "c",
			"directory":            "/d",
			"mount_name":           "e",
			"storage_account_name": "f",
			"token_secret_key":     "g",
			"token_secret_scope":   "h",
		},
		ID:   "e",
		Read: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "e", d.Id())
	assert.Equal(t, "wasbs://c@f.blob.core.windows.net/d", d.Get("source"))
}

func TestResourceAzureBlobMountRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceAzureBlobMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			return common.CommandResults{
				ResultType: "error",
				Summary:    "Mount not found",
			}
		},
		State: map[string]interface{}{
			"auth_type":            "ACCESS_KEY",
			"cluster_id":           "b",
			"container_name":       "c",
			"directory":            "/d",
			"mount_name":           "e",
			"storage_account_name": "f",
			"token_secret_key":     "g",
			"token_secret_scope":   "h",
		},
		ID:      "e",
		Read:    true,
		Removed: true,
	}.ApplyNoError(t)
}

func TestResourceAzureBlobMountRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceAzureBlobMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			return common.CommandResults{
				ResultType: "error",
				Summary:    "Some error",
			}
		},
		State: map[string]interface{}{
			"auth_type":            "ACCESS_KEY",
			"cluster_id":           "b",
			"container_name":       "c",
			"directory":            "/d",
			"mount_name":           "e",
			"storage_account_name": "f",
			"token_secret_key":     "g",
			"token_secret_scope":   "h",
		},
		ID:   "e",
		Read: true,
	}.Apply(t)
	require.EqualError(t, err, "Some error")
	assert.Equal(t, "e", d.Id())
	assert.Equal(t, "", d.Get("source"))
}

func TestResourceAzureBlobMountDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceAzureBlobMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			assert.Contains(t, trunc, "dbutils.fs.unmount(mount_point)")
			return common.CommandResults{
				ResultType: "Text",
				Data:       "",
			}
		},
		State: map[string]interface{}{
			"auth_type":            "ACCESS_KEY",
			"cluster_id":           "b",
			"container_name":       "c",
			"directory":            "/d",
			"mount_name":           "e",
			"storage_account_name": "f",
			"token_secret_key":     "g",
			"token_secret_scope":   "h",
		},
		ID:     "e",
		Delete: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "e", d.Id())
	assert.Equal(t, "", d.Get("source"))
}

func TestAzureAccBlobMount(t *testing.T) {
	client, mp := mountPointThroughReusedCluster(t)
	storageAccountName := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_ACCOUNT")
	accountKey := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_KEY")
	container := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_WASBS")
	testWithNewSecretScope(t, func(scope, key string) {
		testMounting(t, mp, AzureBlobMount{
			StorageAccountName: storageAccountName,
			ContainerName:      container,
			SecretScope:        scope,
			SecretKey:          key,
			Directory:          "/",
		})
	}, client, mp.name, accountKey)
}
