package storage

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/databrickslabs/databricks-terraform/access"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/compute"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
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
		CommandMock: func(commandStr string) (string, error) {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)

			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, "wasbs://c@f.blob.core.windows.net/d")
				assert.Contains(t, trunc, `"fs.azure.account.key.f.blob.core.windows.net":dbutils.secrets.get("h", "g")`)
			}
			assert.Contains(t, trunc, "/mnt/e")
			return "wasbs://c@f.blob.core.windows.net/d", nil
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
		CommandMock: func(commandStr string) (string, error) {
			return "", errors.New("Some error")
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
		CommandMock: func(commandStr string) (string, error) {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			assert.Contains(t, trunc, "dbutils.fs.mounts()")
			assert.Contains(t, trunc, `mount.mountPoint == "/mnt/e"`)
			return "wasbs://c@f.blob.core.windows.net/d", nil
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
		CommandMock: func(commandStr string) (string, error) {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			return "", errors.New("Mount not found")
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
	assert.Equal(t, "", d.Id())
	assert.Equal(t, "", d.Get("source"))
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
		CommandMock: func(commandStr string) (string, error) {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			return "", errors.New("Some error")
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
		CommandMock: func(commandStr string) (string, error) {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			assert.Contains(t, trunc, "dbutils.fs.unmount(mount_point)")
			return "", nil
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
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := common.CommonEnvironmentClient()
	if !client.IsAzure() {
		t.Skip("Test is meant only for Azure")
	}
	if !client.AzureAuth.IsClientSecretSet() {
		t.Skip("Test is meant only for client-secret conf Azure")
	}
	blobAccountName := os.Getenv("TEST_STORAGE_ACCOUNT_NAME")
	if blobAccountName == "" {
		t.Skip("No TEST_STORAGE_ACCOUNT_NAME given")
	}
	blobAccountKey := os.Getenv("TEST_STORAGE_ACCOUNT_KEY")
	if blobAccountKey == "" {
		t.Skip("No TEST_STORAGE_ACCOUNT_KEY given")
	}
	clusterInfo := compute.NewTinyClusterInCommonPoolPossiblyReused()

	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	mp := MountPoint{
		exec:      client.CommandExecutor(),
		clusterID: clusterInfo.ClusterID,
		name:      randomName,
	}
	err := mp.Delete()
	qa.AssertErrorStartsWith(t, err, "Directory not mounted: /mnt/"+randomName)

	source, err := mp.Source()
	assert.Equal(t, "", source)
	qa.AssertErrorStartsWith(t, err, "Mount not found")
	source, err = mp.Mount(AzureBlobMount{
		StorageAccountName: blobAccountName,
		ContainerName:      "dev",
		Directory:          "/",
		SecretKey:          "e",
		SecretScope:        "f",
	})
	assert.Equal(t, "", source)
	qa.AssertErrorStartsWith(t, err, "Secret does not exist with scope: f and key: e")

	randomScope := "test" + randomName
	err = access.NewSecretScopesAPI(client).Create(randomScope, "users")
	assert.NoError(t, err)
	defer func() {
		err = access.NewSecretScopesAPI(client).Delete(randomScope)
		assert.NoError(t, err)
	}()

	err = access.NewSecretsAPI(client).Create(blobAccountKey, randomScope, "key")
	assert.NoError(t, err)

	m := AzureBlobMount{
		StorageAccountName: blobAccountName,
		ContainerName:      "dev",
		Directory:          "/",
		SecretKey:          "key",
		SecretScope:        randomScope,
	}

	source, err = mp.Mount(m)
	assert.Equal(t, m.Source(), source)
	assert.NoError(t, err)
	defer func() {
		err = mp.Delete()
		assert.NoError(t, err)
	}()

	source, err = mp.Source()
	assert.Equal(t, m.Source(), source)
	assert.NoError(t, err)
}
