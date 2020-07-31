package databricks

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAzureAccBlobMount_correctly_mounts(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	blobAccountKey := os.Getenv("TEST_STORAGE_ACCOUNT_KEY")
	blobAccountName := os.Getenv("TEST_STORAGE_ACCOUNT_NAME")
	if blobAccountKey == "" || blobAccountName == "" {
		t.Skipf("Missing keys in environment")
	}
	// assert.NoError(t, os.Setenv("TF_LOG", "TRACE"))
	// assert.NoError(t, os.Setenv("TF_LOG_PATH", "/tmp/tf-integration.log"))

	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resource.Test(t, resource.TestCase{
		IsUnitTest: debugIfCloudEnvSet(),
		Providers:  testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "databricks_secret_scope" "terraform" {
						name                     = "terraform%[3]s"
						initial_manage_principal = "users"
					}
					resource "databricks_secret" "storage_key" {
						key          = "blob_storage_key"
						string_value = "%[1]s"
						scope        = databricks_secret_scope.terraform.name
					}
					resource "databricks_azure_blob_mount" "mount" {
						container_name       = "dev" # Created by prereqs.tf
						storage_account_name = "%[2]s"
						mount_name           = "%[3]s"
						auth_type            = "ACCESS_KEY"
						token_secret_scope   = databricks_secret_scope.terraform.name
						token_secret_key     = databricks_secret.storage_key.key
					}
				`, blobAccountKey, blobAccountName, randomName),
				Check: epoch.ResourceCheck("databricks_azure_blob_mount.mount", func(client *service.DatabricksClient, id string) error {
					clusterInfo, err := client.Clusters().GetOrCreateRunningCluster("TerraformIntegrationTest")
					assert.NoError(t, err)
					mp := NewMountPoint(client, id, clusterInfo.ClusterID)
					source, err := mp.Source()
					assert.NoError(t, err)
					assert.Equal(t, fmt.Sprintf("wasbs://%s@%s.blob.core.windows.net/", "dev", blobAccountName), source)
					return nil
				}),
			},
			{
				PreConfig: func() {
					client := testAccProvider.Meta().(*service.DatabricksClient)
					clusterInfo := service.NewTinyClusterInCommonPoolPossiblyReused()
					mp := NewMountPoint(client, randomName, clusterInfo.ClusterID)
					// TODO: check correctness of remounting with different thing...

					// remove mount out of tf resource and see what happens
					err := mp.Delete()
					assert.NoError(t, err)
				},
				Config: fmt.Sprintf(`
					resource "databricks_secret_scope" "terraform" {
						name                     = "terraform%[3]s"
						initial_manage_principal = "users"
					}
					resource "databricks_secret" "storage_key" {
						key          = "blob_storage_key"
						string_value = "%[1]s"
						scope        = databricks_secret_scope.terraform.name
					}
					resource "databricks_azure_blob_mount" "mount" {
						container_name       = "dev" # Created by prereqs.tf
						storage_account_name = "%[2]s"
						mount_name           = "%[3]s"
						auth_type            = "ACCESS_KEY"
						token_secret_scope   = databricks_secret_scope.terraform.name
						token_secret_key     = databricks_secret.storage_key.key
					}
				`, blobAccountKey, blobAccountName, randomName),
				// Destroy: true,
			},
		},
	})
}

func TestResourceAzureBlobMountCreate(t *testing.T) {
	d, err := ResourceFixture{
		Fixtures: []HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: model.ClusterInfo{
					State: model.ClusterStateRunning,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: model.ClusterInfo{
					State: model.ClusterStateRunning,
				},
			},
		},
		Resource: resourceAzureBlobMount(),
		CommandMock: func(commandStr string) (string, error) {
			trunc := service.TrimLeadingWhitespace(commandStr)
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
	d, err := ResourceFixture{
		Fixtures: []HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: model.ClusterInfo{
					State: model.ClusterStateRunning,
				},
			},
		},
		Resource: resourceAzureBlobMount(),
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
	d, err := ResourceFixture{
		Fixtures: []HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: model.ClusterInfo{
					State: model.ClusterStateRunning,
				},
			},
		},
		Resource: resourceAzureBlobMount(),
		CommandMock: func(commandStr string) (string, error) {
			trunc := service.TrimLeadingWhitespace(commandStr)
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
	d, err := ResourceFixture{
		Fixtures: []HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: model.ClusterInfo{
					State: model.ClusterStateRunning,
				},
			},
		},
		Resource: resourceAzureBlobMount(),
		CommandMock: func(commandStr string) (string, error) {
			trunc := service.TrimLeadingWhitespace(commandStr)
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
	d, err := ResourceFixture{
		Fixtures: []HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: model.ClusterInfo{
					State: model.ClusterStateRunning,
				},
			},
		},
		Resource: resourceAzureBlobMount(),
		CommandMock: func(commandStr string) (string, error) {
			trunc := service.TrimLeadingWhitespace(commandStr)
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
	d, err := ResourceFixture{
		Fixtures: []HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: model.ClusterInfo{
					State: model.ClusterStateRunning,
				},
			},
		},
		Resource: resourceAzureBlobMount(),
		CommandMock: func(commandStr string) (string, error) {
			trunc := service.TrimLeadingWhitespace(commandStr)
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
	client := service.CommonEnvironmentClient()
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
	clusterInfo := service.NewTinyClusterInCommonPoolPossiblyReused()

	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	mp := MountPoint{
		exec:      client.Commands(),
		clusterID: clusterInfo.ClusterID,
		name:      randomName,
	}
	err := mp.Delete()
	assertErrorStartsWith(t, err, "Directory not mounted: /mnt/"+randomName)

	source, err := mp.Source()
	assert.Equal(t, "", source)
	assertErrorStartsWith(t, err, "Mount not found")
	source, err = mp.Mount(AzureBlobMount{
		StorageAccountName: blobAccountName,
		ContainerName:      "dev",
		Directory:          "/",
		SecretKey:          "e",
		SecretScope:        "f",
	})
	assert.Equal(t, "", source)
	assertErrorStartsWith(t, err, "Secret does not exist with scope: f and key: e")

	randomScope := "test" + randomName
	err = client.SecretScopes().Create(randomScope, "users")
	assert.NoError(t, err)
	defer func() {
		err = client.SecretScopes().Delete(randomScope)
		assert.NoError(t, err)
	}()

	err = client.Secrets().Create(blobAccountKey, randomScope, "key")
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
