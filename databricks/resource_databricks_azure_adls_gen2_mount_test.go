package databricks

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccAzureAdlsGen2Mount_correctly_mounts(t *testing.T) {
	terraformToApply := testAccAzureAdlsGen2Mount_correctly_mounts()

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: terraformToApply,
			},
		},
	})
}

func TestAccAzureAdlsGen2Mount_capture_error(t *testing.T) {
	terraformToApply := testAccAzureAdlsGen2Mount_capture_error()

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:             terraformToApply,
				ExpectNonEmptyPlan: true,
				ExpectError:        regexp.MustCompile("java.lang.IllegalArgumentException: Secret does not exist with scope"),
				Destroy:            false,
			},
		},
	})
}

func testAccAzureAdlsGen2Mount_correctly_mounts() string {
	clientID := os.Getenv("ARM_CLIENT_ID")
	clientSecret := os.Getenv("ARM_CLIENT_SECRET")
	tenantID := os.Getenv("ARM_TENANT_ID")
	gen2AdalName := os.Getenv("TEST_GEN2_ADAL_NAME")

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
	
	resource "databricks_secret" "client_secret" {
	  key          = "datalake_sp_secret"
	  string_value = "%[2]s"
	  scope        = databricks_secret_scope.terraform.name
	}
	
	resource "databricks_azure_adls_gen2_mount" "mount" {
	  cluster_id             = databricks_cluster.cluster.id
	  container_name         = "dev" # Created by prereqs.tf
	  storage_account_name   = "%[4]s"
	  directory              = ""
	  mount_name             = "localdir${databricks_cluster.cluster.cluster_id}"
	  tenant_id              = "%[3]s"
	  client_id              = "%[1]s"
	  client_secret_scope    = databricks_secret_scope.terraform.name
	  client_secret_key      = databricks_secret.client_secret.key
	  initialize_file_system = true
	}

`, clientID, clientSecret, tenantID, gen2AdalName)
	return definition
}

func testAccAzureAdlsGen2Mount_capture_error() string {
	clientID := os.Getenv("ARM_CLIENT_ID")
	clientSecret := os.Getenv("ARM_CLIENT_SECRET")
	tenantID := os.Getenv("ARM_TENANT_ID")
	subscriptionID := os.Getenv("ARM_SUBSCRIPTION_ID")
	workspaceName := os.Getenv("TEST_WORKSPACE_NAME")
	resourceGroupName := os.Getenv("TEST_RESOURCE_GROUP")
	managedResourceGroupName := os.Getenv("TEST_MANAGED_RESOURCE_GROUP")
	location := os.Getenv("TEST_LOCATION")
	gen2AdalName := os.Getenv("TEST_GEN2_ADAL_NAME")

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
	
	resource "databricks_secret" "client_secret" {
	  key          = "datalake_sp_secret"
	  string_value = "%[2]s"
	  scope        = databricks_secret_scope.terraform.name
	}
	
	resource "databricks_azure_adls_gen2_mount" "mount" {
	  cluster_id             = databricks_cluster.cluster.id
	  container_name         = "dev" # Created by prereqs.tf
	  storage_account_name   = "%[9]s"
	  directory              = ""
	  mount_name             = "localdir${databricks_cluster.cluster.cluster_id}"
	  tenant_id              = "%[3]s"
	  client_id              = "%[1]s"
	  client_secret_scope    = databricks_secret_scope.terraform.name
	  client_secret_key      = "SECRET_KEY_WRONG_ON_PURPOSE"
	  initialize_file_system = true
	}

`, clientID, clientSecret, tenantID, subscriptionID, workspaceName, resourceGroupName, managedResourceGroupName, location, gen2AdalName)
	return definition
}
