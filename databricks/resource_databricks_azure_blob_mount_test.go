package databricks

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccAzureBlobMount_correctly_mounts(t *testing.T) {
	terraformToApply := testAccAzureBlobMount_correctly_mounts()

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: terraformToApply,
			},
		},
	})
}

func testAddAzureBlobMount_correctly_mounts_unmount() {

	iamMountCommand := `
dbutils.fs.unmount("/mnt/dev")
dbutils.notebook.exit("success")
`
	resp, err := client.Commands().Execute(clusterID, "python", iamMountCommand)
	if err != nil {
		return err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return errors.New(resp.Results.Summary)
	}
}

func testAccAzureBlobMount_correctly_mounts() string {
	clientID := os.Getenv("ARM_CLIENT_ID")
	clientSecret := os.Getenv("ARM_CLIENT_SECRET")
	tenantID := os.Getenv("ARM_TENANT_ID")
	subscriptionID := os.Getenv("ARM_SUBSCRIPTION_ID")
	workspaceName := os.Getenv("TEST_WORKSPACE_NAME")
	resourceGroupName := os.Getenv("TEST_RESOURCE_GROUP")
	managedResourceGroupName := os.Getenv("TEST_MANAGED_RESOURCE_GROUP")
	location := os.Getenv("TEST_LOCATION")
	blobAccountKey := os.Getenv("TEST_STORAGE_ACCOUNT_KEY")
	blobAccountName := os.Getenv("TEST_STORAGE_ACCOUNT_NAME")

	definition := fmt.Sprintf(`
	provider "databricks" {
	  azure_auth = {
		client_id              = "%[1]s"
		client_secret          = "%[2]s"
		tenant_id              = "%[3]s"
		subscription_id        = "%[4]s"

		workspace_name         = "%[5]s"
		resource_group         = "%[6]s"
		managed_resource_group = "%[7]s"
		azure_region           = "%[8]s"
	  }
	}

	resource "databricks_cluster" "cluster" {
		num_workers = 1
		spark_version = "6.4.x-scala2.11"
		node_type_id = "Standard_D3_v2"
		# Don't spend too much, turn off cluster after 15mins
		autotermination_minutes = 15
	} 

	resource "databricks_secret_scope" "terraform" {
	  # Add the cluster ID into the secret scope to ensure 
	  # it doesn't clash with one used by another test
	  name                     = "terraform${databricks_cluster.cluster.cluster_id}"
	  initial_manage_principal = "users"
	}
	
	resource "databricks_secret" "storage_key" {
	  key          = "blob_storage_key"
	  string_value = "%[10]s"
	  scope        = databricks_secret_scope.terraform.name
	}
	
	resource "databricks_azure_blob_mount" "mount" {
		cluster_id           = databricks_cluster.cluster.id
		container_name       = "dev" # Created by prereqs.tf
		storage_account_name = "%[9]s"
		mount_name           = "dev"
		auth_type            = "ACCESS_KEY"
		token_secret_scope   = databricks_secret_scope.terraform.name
		token_secret_key     = databricks_secret.storage_key.key
	}

`, clientID, clientSecret, tenantID, subscriptionID, workspaceName, resourceGroupName, managedResourceGroupName, location, blobAccountName, blobAccountKey)
	return definition
}
