package acceptance

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"
	"github.com/databrickslabs/terraform-provider-databricks/internal/compute"
	"github.com/databrickslabs/terraform-provider-databricks/storage"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

func TestAzureAccBlobMount(t *testing.T) {
	client, mp := mountPointThroughReusedCluster(t)
	storageAccountName := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_ACCOUNT")
	accountKey := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_KEY")
	container := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_WASBS")
	testWithNewSecretScope(t, func(scope, key string) {
		testMounting(t, mp, storage.AzureBlobMount{
			StorageAccountName: storageAccountName,
			ContainerName:      container,
			SecretScope:        scope,
			SecretKey:          key,
			Directory:          "/",
		})
	}, client, mp.Name, accountKey)
}

func TestAzureAccBlobMount_correctly_mounts(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	config := acceptance.EnvironmentTemplate(t, `
	resource "databricks_secret_scope" "terraform" {
		name                     = "terraform-{var.RANDOM}"
		initial_manage_principal = "users"
	}
	resource "databricks_secret" "storage_key" {
		key          = "blob_storage_key"
		string_value = "{env.TEST_STORAGE_V2_KEY}"
		scope        = databricks_secret_scope.terraform.name
	}
	resource "databricks_azure_blob_mount" "mount" {
		storage_account_name = "{env.TEST_STORAGE_V2_ACCOUNT}"
		container_name       = "{env.TEST_STORAGE_V2_WASBS}"
		mount_name           = "{var.RANDOM}"
		auth_type            = "ACCESS_KEY"
		token_secret_scope   = databricks_secret_scope.terraform.name
		token_secret_key     = databricks_secret.storage_key.key
	}`)
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: mountResourceCheck("databricks_azure_blob_mount.mount",
					func(client *common.DatabricksClient, mp storage.MountPoint) error {
						source, err := mp.Source()
						assert.NoError(t, err)
						assert.Equal(t, fmt.Sprintf(
							"wasbs://%s@%s.blob.core.windows.net/",
							qa.FirstKeyValue(t, config, "container_name"),
							qa.FirstKeyValue(t, config, "storage_account_name")), source)
						return nil
					}),
			},
			{
				PreConfig: func() {
					client := compute.CommonEnvironmentClientWithRealCommandExecutor()
					clusterInfo := compute.NewTinyClusterInCommonPoolPossiblyReused()
					mp := storage.NewMountPoint(client.CommandExecutor(context.Background()),
						qa.FirstKeyValue(t, config, "mount_name"),
						clusterInfo.ClusterID)
					err := mp.Delete()
					assert.NoError(t, err)
				},
				Config: config,
				// Destroy: true,
			},
		},
	})
}
