package databricks

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAzureAccAdlsGen2Mount_correctly_mounts(t *testing.T) {
	// TODO: refactor for common instance pool & AZ CLI
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	terraformToApply := testAdlsGen2MountCorrectlyMounts(t)

	resource.Test(t, resource.TestCase{
		Providers:  testAccProviders,
		IsUnitTest: debugIfCloudEnvSet(),
		Steps: []resource.TestStep{
			{
				Config: terraformToApply,
			},
		},
	})
}

func TestAzureAccAdlsGen2Mount_cluster_deleted_correctly_mounts(t *testing.T) {
	// TODO: refactor for common instance pool & AZ CLI
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	terraformToApply := testAdlsGen2MountCorrectlyMounts(t)
	var cluster model.ClusterInfo

	resource.Test(t, resource.TestCase{
		Providers:  testAccProviders,
		IsUnitTest: debugIfCloudEnvSet(),
		Steps: []resource.TestStep{
			{
				Config: terraformToApply,
				Check:  testClusterResourceExists("databricks_cluster.cluster", &cluster, t),
			},
			{
				PreConfig: func() {
					client := testAccProvider.Meta().(*service.DatabricksClient)
					err := client.Clusters().Delete(cluster.ClusterID)
					assert.NoError(t, err, err)
				},
				Config: terraformToApply,
			},
		},
	})
}

func TestAzureAccAdlsGen2Mount_capture_error(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	clientID := os.Getenv("ARM_CLIENT_ID")             // make dual-env-var
	clientSecret := os.Getenv("ARM_CLIENT_SECRET")     // make dual-env-var
	tenantID := os.Getenv("ARM_TENANT_ID")             // make dual-env-var
	subscriptionID := os.Getenv("ARM_SUBSCRIPTION_ID") // make dual-env-var
	workspaceName := os.Getenv("TEST_WORKSPACE_NAME")
	resourceGroupName := os.Getenv("TEST_RESOURCE_GROUP")
	managedResourceGroupName := os.Getenv("TEST_MANAGED_RESOURCE_GROUP")
	location := os.Getenv("TEST_LOCATION")
	gen2AdalName := os.Getenv("TEST_GEN2_ADAL_NAME")

	if clientID == "" || clientSecret == "" || tenantID == "" ||
		subscriptionID == "" || workspaceName == "" || resourceGroupName == "" ||
		managedResourceGroupName == "" || location == "" || gen2AdalName == "" {
		t.Skipf("Missing configuration options for ADLSv2 mounts")
	}

	resource.Test(t, resource.TestCase{
		IsUnitTest: debugIfCloudEnvSet(),
		Providers:  testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "databricks_cluster" "cluster" {
						cluster_name = "%[12]s"
						num_workers = 1
						spark_version = "%[10]s"
						instance_pool_id = "%[11]s"
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
				`, clientID, clientSecret, tenantID, subscriptionID, workspaceName,
					resourceGroupName, managedResourceGroupName, location, gen2AdalName,
					service.CommonRuntimeVersion(), service.CommonInstancePoolID(),
					t.Name()),
				ExpectNonEmptyPlan: true,
				ExpectError:        regexp.MustCompile("java.lang.IllegalArgumentException: Secret does not exist with scope"),
				Destroy:            false,
			},
		},
	})
}

func testAdlsGen2MountCorrectlyMounts(t *testing.T) string {
	clientID := os.Getenv("ARM_CLIENT_ID")
	clientSecret := os.Getenv("ARM_CLIENT_SECRET")
	tenantID := os.Getenv("ARM_TENANT_ID")
	gen2AdalName := os.Getenv("TEST_GEN2_ADAL_NAME")

	if clientID == "" || clientSecret == "" || tenantID == "" || gen2AdalName == "" {
		t.Skip("Missing configuration options for ADLSv2 mounts")
	}

	// TODO: pass in ADLS container name by ENV

	definition := fmt.Sprintf(`
	resource "databricks_cluster" "cluster" {
		num_workers = 1
		cluster_name = "%[7]s"
		spark_version = "%[5]s"
		instance_pool_id = "%[6]s"
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
`, clientID, clientSecret, tenantID, gen2AdalName, service.CommonRuntimeVersion(),
		service.CommonInstancePoolID(), t.Name())
	return definition
}

// testClusterResourceExists queries the API and retrieves the matching Cluster.
func testClusterResourceExists(n string, cluster *model.ClusterInfo, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		conn := testAccProvider.Meta().(*service.DatabricksClient)
		resp, err := conn.Clusters().Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		*cluster = resp
		return nil
	}
}
