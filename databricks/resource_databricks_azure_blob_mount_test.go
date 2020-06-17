package databricks

import (
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccAzureBlobMount_correctly_mounts(t *testing.T) {
	terraformToApply := testAccAzureBlobMountCorrectlyMounts()
	var clusterInfo model.ClusterInfo
	var azureBlobMount AzureBlobMount

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: terraformToApply,
				Check: resource.ComposeTestCheckFunc(
					testAccAzureBlobMountClusterExists("databricks_cluster.cluster", &clusterInfo),
					testAccAzureBlobMountMountExists("databricks_azure_blob_mount.mount", &azureBlobMount, &clusterInfo),
				),
			},
			{
				PreConfig: func() {
					client := testAccProvider.Meta().(*service.DBApiClient)
					err := azureBlobMount.Delete(client.Commands(), clusterInfo.ClusterID)
					assert.NoError(t, err, "TestAccAzureBlobMount_correctly_mounts: Failed to remove the mount.")
				},
				Config: terraformToApply,
				Check: resource.ComposeTestCheckFunc(
					testAccAzureBlobMountMountExists("databricks_azure_blob_mount.mount", &azureBlobMount, &clusterInfo),
				),
			},
		},
	})
}

func TestAccAzureBlobMount_cluster_deleted_correctly_mounts(t *testing.T) {
	terraformToApply := testAccAzureBlobMountCorrectlyMounts()
	var clusterInfo model.ClusterInfo
	var azureBlobMount AzureBlobMount

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: terraformToApply,
				Check: resource.ComposeTestCheckFunc(
					testAccAzureBlobMountClusterExists("databricks_cluster.cluster", &clusterInfo),
					testAccAzureBlobMountMountExists("databricks_azure_blob_mount.mount", &azureBlobMount, &clusterInfo),
				),
			},
			{
				PreConfig: func() {
					client := testAccProvider.Meta().(*service.DBApiClient)
					err := client.Clusters().Delete(clusterInfo.ClusterID)
					assert.NoError(t, err, err)
				},
				Config: terraformToApply,
				Check: resource.ComposeTestCheckFunc(
					testAccAzureBlobMountMountExists("databricks_azure_blob_mount.mount", &azureBlobMount, &clusterInfo),
				),
			},
		},
	})
}

func testAccAzureBlobMountClusterExists(n string, clusterInfo *model.ClusterInfo) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		client := testAccProvider.Meta().(*service.DBApiClient)
		resp, err := client.Clusters().Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		// If no error, assign the response Widget attribute to the widget pointer
		*clusterInfo = resp
		return nil
	}
}

func testAccAzureBlobMountMountExists(n string, azureBlobMount *AzureBlobMount, clusterInfo *model.ClusterInfo) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found in tfstate: %s", n)
		}

		authType := rs.Primary.Attributes["auth_type"]
		containerName := rs.Primary.Attributes["container_name"]
		storageAccountName := rs.Primary.Attributes["storage_account_name"]
		directory := rs.Primary.Attributes["directory"]
		mountName := rs.Primary.Attributes["mount_name"]
		tokenSecretScope := rs.Primary.Attributes["token_secret_scope"]
		tokenSecretKey := rs.Primary.Attributes["token_secret_key"]

		blobMount := NewAzureBlobMount(containerName, storageAccountName, directory, mountName, authType,
			tokenSecretScope, tokenSecretKey)

		client := testAccProvider.Meta().(*service.DBApiClient)
		clusterID := clusterInfo.ClusterID

		message, err := blobMount.Read(client.Commands(), clusterID)

		if err != nil {
			return fmt.Errorf("Error reading the mount %s: error %s", message, err)
		}

		*azureBlobMount = *blobMount
		return nil
	}
}

func testAccAzureBlobMountCorrectlyMounts() string {
	blobAccountKey := os.Getenv("TEST_STORAGE_ACCOUNT_KEY")
	blobAccountName := os.Getenv("TEST_STORAGE_ACCOUNT_NAME")

	definition := fmt.Sprintf(`
	resource "databricks_cluster" "cluster" {
		num_workers = 1
		spark_version = "6.4.x-scala2.11"
		node_type_id = "Standard_D3_v2"
		# Don't spend too much, turn off cluster after 15mins
		autotermination_minutes = 15
		spark_conf = {
			"spark.databricks.delta.preview.enabled": "false"
		}
	} 

	resource "databricks_secret_scope" "terraform" {
	  # Add the cluster ID into the secret scope to ensure 
	  # it doesn't clash with one used by another test
	  name                     = "terraform${databricks_cluster.cluster.cluster_id}"
	  initial_manage_principal = "users"
	}
	
	resource "databricks_secret" "storage_key" {
	  key          = "blob_storage_key"
	  string_value = "%[1]s"
	  scope        = databricks_secret_scope.terraform.name
	}
	
	resource "databricks_azure_blob_mount" "mount" {
		cluster_id           = databricks_cluster.cluster.id
		container_name       = "dev" # Created by prereqs.tf
		storage_account_name = "%[2]s"
		mount_name           = "dev"
		auth_type            = "ACCESS_KEY"
		token_secret_scope   = databricks_secret_scope.terraform.name
		token_secret_key     = databricks_secret.storage_key.key
	}

`, blobAccountKey, blobAccountName)
	return definition
}
