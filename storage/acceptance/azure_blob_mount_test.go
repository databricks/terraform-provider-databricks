package acceptance

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/compute"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	. "github.com/databrickslabs/databricks-terraform/storage"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

func mountResourceCheck(name string,
	cb func(*common.DatabricksClient, MountPoint) error) resource.TestCheckFunc {
	return acceptance.ResourceCheck(name,
		func(client *common.DatabricksClient, id string) error {
			client.WithCommandExecutor(func(ctx context.Context, client *common.DatabricksClient) common.CommandExecutor {
				return compute.NewCommandsAPI(ctx, client)
			})
			clusterInfo := compute.NewTinyClusterInCommonPoolPossiblyReused()
			mp := NewMountPoint(client.CommandExecutor(context.Background()), id, clusterInfo.ClusterID)
			return cb(client, mp)
		})
}

func TestAzureAccBlobMount_correctly_mounts(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	config := qa.EnvironmentTemplate(t, `
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
					func(client *common.DatabricksClient, mp MountPoint) error {
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
					mp := NewMountPoint(client.CommandExecutor(context.Background()),
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
