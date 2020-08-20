package acceptance

import (
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/compute"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	. "github.com/databrickslabs/databricks-terraform/storage"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/stretchr/testify/assert"
)

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
		string_value = "{env.TEST_STORAGE_ACCOUNT_KEY}"
		scope        = databricks_secret_scope.terraform.name
	}
	resource "databricks_azure_blob_mount" "mount" {
		container_name       = "dev"
		storage_account_name = "{env.TEST_STORAGE_ACCOUNT_NAME}"
		mount_name           = "{var.RANDOM}"
		auth_type            = "ACCESS_KEY"
		token_secret_scope   = databricks_secret_scope.terraform.name
		token_secret_key     = databricks_secret.storage_key.key
	}`)
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: acceptance.ResourceCheck("databricks_azure_blob_mount.mount", func(client *common.DatabricksClient, id string) error {
					clusterInfo, err := compute.NewClustersAPI(client).GetOrCreateRunningCluster("TerraformIntegrationTest")
					assert.NoError(t, err)
					mp := NewMountPoint(client, id, clusterInfo.ClusterID)
					source, err := mp.Source()
					assert.NoError(t, err)
					assert.Equal(t, fmt.Sprintf("wasbs://%s@%s.blob.core.windows.net/", "dev",
						qa.FirstKeyValue(t, config, "storage_account_name")), source)
					return nil
				}),
			},
			{
				PreConfig: func() {
					client := common.CommonEnvironmentClient()
					clusterInfo := compute.NewTinyClusterInCommonPoolPossiblyReused()
					mp := NewMountPoint(client, randomName, clusterInfo.ClusterID)
					// TODO: check correctness of remounting with different thing...

					// remove mount out of tf resource and see what happens
					err := mp.Delete()
					assert.NoError(t, err)
				},
				Config: config,
				// Destroy: true,
			},
		},
	})
}
