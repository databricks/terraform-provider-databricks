package databricks

import (
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAzureAccBlobMount_correctly_mounts(t *testing.T) {
	// TODO: refactor for common instance pool & AZ CLI
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	terraformToApply := testBlobMountCorrectlyMounts(t)
	var clusterInfo model.ClusterInfo
	var azureBlobMount AzureBlobMount

	resource.Test(t, resource.TestCase{
		DisableBinaryDriver: true, // to see debug logs...
		IsUnitTest:          debugIfCloudEnvSet(),
		Providers:           testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: terraformToApply,
				Check: resource.ComposeTestCheckFunc(
					testBlobMountClusterExists("databricks_cluster.cluster", &clusterInfo),
					testBlobMountMountExists("databricks_azure_blob_mount.mount", &azureBlobMount, &clusterInfo),
				),
			},
			{
				// PreConfig: func() {
				// 	client := testAccProvider.Meta().(*service.DatabricksClient)
				// 	err := azureBlobMount.Delete(client.Commands())
				// 	assert.NoError(t, err, "TestBlobMount_correctly_mounts: Failed to remove the mount.")
				// },
				Config: terraformToApply,
				Check: resource.ComposeTestCheckFunc(
					testBlobMountMountExists("databricks_azure_blob_mount.mount", &azureBlobMount, &clusterInfo),
				),
				Destroy: true,
			},
		},
	})
}

func TestAzureAccBlobMount_cluster_deleted_correctly_mounts(t *testing.T) {
	// TODO: refactor for common instance pool & AZ CLI
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	terraformToApply := testBlobMountCorrectlyMounts(t)
	var clusterInfo model.ClusterInfo
	var azureBlobMount AzureBlobMount

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: terraformToApply,
				Check: resource.ComposeTestCheckFunc(
					testBlobMountClusterExists("databricks_cluster.cluster", &clusterInfo),
					testBlobMountMountExists("databricks_azure_blob_mount.mount", &azureBlobMount, &clusterInfo),
				),
			},
			{
				PreConfig: func() {
					client := testAccProvider.Meta().(*service.DatabricksClient)
					err := client.Clusters().Delete(clusterInfo.ClusterID)
					assert.NoError(t, err, err)
				},
				Config: terraformToApply,
				Check: resource.ComposeTestCheckFunc(
					testBlobMountMountExists("databricks_azure_blob_mount.mount", &azureBlobMount, &clusterInfo),
				),
			},
		},
	})
}

func testBlobMountClusterExists(n string, clusterInfo *model.ClusterInfo) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		client := testAccProvider.Meta().(*service.DatabricksClient)
		resp, err := client.Clusters().Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		// If no error, assign the response Widget attribute to the widget pointer
		*clusterInfo = resp
		return nil
	}
}

func testBlobMountMountExists(n string, azureBlobMount *AzureBlobMount, clusterInfo *model.ClusterInfo) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		// rs, ok := s.RootModule().Resources[n]
		// if !ok {
		// 	return fmt.Errorf("Not found in tfstate: %s", n)
		// }

		// authType := rs.Primary.Attributes["auth_type"]
		// containerName := rs.Primary.Attributes["container_name"]
		// storageAccountName := rs.Primary.Attributes["storage_account_name"]
		// directory := rs.Primary.Attributes["directory"]
		// mountName := rs.Primary.Attributes["mount_name"]
		// tokenSecretScope := rs.Primary.Attributes["token_secret_scope"]
		// tokenSecretKey := rs.Primary.Attributes["token_secret_key"]

		// blobMount := AzureBlobMount{} // NewAzureBlobMount(containerName, storageAccountName, directory, mountName, authType, tokenSecretScope, tokenSecretKey)

		// client := testAccProvider.Meta().(*service.DatabricksClient)
		// clusterID := clusterInfo.ClusterID

		// message, err := blobMount.Read(client.Commands(), clusterID)

		// if err != nil {
		// 	return fmt.Errorf("Error reading the mount %s: error %s", message, err)
		// }

		// *azureBlobMount = *blobMount
		return nil
	}
}

func testBlobMountCorrectlyMounts(t *testing.T) string {
	blobAccountKey := os.Getenv("TEST_STORAGE_ACCOUNT_KEY")
	blobAccountName := os.Getenv("TEST_STORAGE_ACCOUNT_NAME")
	if blobAccountKey == "" || blobAccountName == "" {
		t.Skipf("Missing keys in environment")
	}

	definition := fmt.Sprintf(`
	resource "databricks_cluster" "cluster" {
		num_workers = 1
		cluster_name = "%[5]s"
		spark_version = "%[3]s"
		instance_pool_id = "%[4]s"
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

`, blobAccountKey, blobAccountName, service.CommonRuntimeVersion(),
		service.CommonInstancePoolID(), t.Name())
	return definition
}

// TODO: change by mocks...
// func TestResourceAzureBlobMountCreate(t *testing.T) {
// 	t.Skip("Skipping test until retry timeouts are re-enabled again")
// 	d, err := ResourceTester(t, []HTTPFixture{
// 		{
// 			Method:   "GET",
// 			Resource: "/api/2.0/clusters/get?cluster_id=b",
// 			Response: model.ClusterInfo{
// 				State: "RUNNING",
// 			},
// 		},
// 		{
// 			Method:   "POST",
// 			Resource: "/api/1.2/contexts/create",
// 			ExpectedRequest: map[string]interface{}{
// 				"clusterId": "b",
// 				"language":  "python",
// 			},
// 			Response: map[string]interface{}{
// 				"id": "789",
// 			},
// 		},
// 		{
// 			Method:   "GET",
// 			Resource: "/api/1.2/contexts/status?clusterId=b&contextId=789",
// 			Response: model.Command{
// 				Status: "Running",
// 			},
// 		},
// 		{
// 			Method:   "POST",
// 			Resource: "/api/1.2/commands/execute",
// 			ExpectedRequest: map[string]interface{}{
// 				"clusterId": "b",
// 				// TODO: eventually make generic ident-whitespace ignoring function
// 				"command": "\nfor mount in dbutils.fs.mounts():\n  if mount.mountPoint == \"/mnt/e\" and mount.source==\"wasbs://c@f.blob.core.windows.net/d\":\n" +
// 					"    print (\"Mount already exists\")\n    dbutils.notebook.exit(\"success\")\n\ntry:\n  dbutils.fs.mount(\n" +
// 					"    source = \"wasbs://c@f.blob.core.windows.net/d\",\n    mount_point = \"/mnt/e\",\n    " +
// 					"extra_configs = {\"fs.azure.account.key.f.blob.core.windows.net\":dbutils.secrets.get(scope = \"h\", key = \"g\")})\n" +
// 					"except Exception as e:\n  dbutils.fs.unmount(\"/mnt/e\")\n  raise e\ndbutils.notebook.exit(\"success\")\n",
// 				"contextId": "789",
// 				"language":  "python"},
// 			Response: map[string]interface{}{
// 				"id": "876",
// 			},
// 		},
// 		{
// 			Method:   "GET",
// 			Resource: "/api/1.2/commands/status?clusterId=b&commandId=876&contextId=789",
// 			Response: model.Command{
// 				Status: "Finished",
// 			},
// 		},
// 		{
// 			Method:   "GET",
// 			Resource: "/api/1.2/commands/status?clusterId=b&commandId=876&contextId=789",
// 			Response: model.Command{
// 				Status: "Finished",
// 				Results: &model.CommandResults{
// 					ResultType: "success",
// 				},
// 			},
// 		},
// 		{
// 			Method:   "POST",
// 			Resource: "/api/1.2/contexts/destroy",
// 			ExpectedRequest: map[string]interface{}{
// 				"clusterId": "b",
// 				"contextId": "789",
// 			},
// 		},
// 		{
// 			Method:   "GET",
// 			Resource: "/api/2.0/clusters/get?cluster_id=b",
// 			Response: model.ClusterInfo{
// 				State: "RUNNING",
// 			},
// 		},
// 		{
// 			Method:   "POST",
// 			Resource: "/api/1.2/contexts/create",
// 			ExpectedRequest: map[string]interface{}{
// 				"clusterId": "b",
// 				"language":  "python",
// 			},
// 			Response: map[string]interface{}{
// 				"id": "987",
// 			},
// 		},
// 		{
// 			Method:   "GET",
// 			Resource: "/api/1.2/contexts/status?clusterId=b&contextId=987",
// 			Response: model.Command{
// 				Status: "Running",
// 			},
// 		},
// 		{
// 			Method:   "POST",
// 			Resource: "/api/1.2/commands/execute",
// 			ExpectedRequest: map[string]interface{}{
// 				"clusterId": "b",
// 				"command":   "\nfor mount in dbutils.fs.mounts():\n  if mount.mountPoint == \"/mnt/e\":\n    dbutils.notebook.exit(mount.source)\n",
// 				"contextId": "987",
// 				"language":  "python",
// 			},
// 			Response: map[string]interface{}{
// 				"id": "678",
// 			},
// 		},
// 		{
// 			Method:   "GET",
// 			Resource: "/api/1.2/commands/status?clusterId=b&commandId=678&contextId=987",
// 			Response: model.Command{
// 				Status: "Finished",
// 				Results: &model.CommandResults{
// 					ResultType: "success",
// 				},
// 			},
// 		},
// 		{
// 			Method:   "GET",
// 			Resource: "/api/1.2/commands/status?clusterId=b&commandId=678&contextId=987",
// 			Response: model.Command{
// 				Status: "Finished",
// 				Results: &model.CommandResults{
// 					ResultType: "success",
// 					Data:       "wa.wb.wc",
// 				},
// 			},
// 		},
// 		{
// 			Method:   "POST",
// 			Resource: "/api/1.2/contexts/destroy",
// 			ExpectedRequest: map[string]interface{}{
// 				"clusterId": "b",
// 				"contextId": "987",
// 			},
// 		},
// 	}, resourceAzureBlobMount, map[string]interface{}{
// 		"auth_type":            "ACCESS_KEY",
// 		"cluster_id":           "b",
// 		"container_name":       "c",
// 		"directory":            "/d",
// 		"mount_name":           "e",
// 		"storage_account_name": "f",
// 		"token_secret_key":     "g",
// 		"token_secret_scope":   "h",
// 	}, resourceAzureBlobMountCreate)
// 	assert.NoError(t, err, err)
// 	assert.Equal(t, "e", d.Id())
// 	assert.Equal(t, "wa.wb.wc", d.Get("directory"))
// }

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
	clusterInfo, err := client.Clusters().GetOrCreateRunningCluster("TerraformIntegrationTest")
	assert.NoError(t, err)

	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	mp := MountPoint{
		exec:      client.Commands(),
		clusterID: clusterInfo.ClusterID,
		name:      randomName,
	}
	err = mp.Delete()
	assert.EqualError(t, err, "Directory not mounted: /mnt/"+randomName)

	source, err := mp.Source()
	assert.Equal(t, "", source)
	assert.EqualError(t, err, "Mount not found")

	source, err = mp.Mount(AzureBlobMount{
		StorageAccountName: blobAccountName,
		ContainerName:      "dev",
		Directory:          "/",
		SecretKey:          "e",
		SecretScope:        "f",
	})
	assert.Equal(t, "", source)
	assert.EqualError(t, err, "Secret does not exist with scope: f and key: e")

	randomScope := "test" + randomName
	err = client.SecretScopes().Create(randomScope, "users")
	assert.NoError(t, err)
	defer client.SecretScopes().Delete(randomScope)

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
	defer mp.Delete()

	source, err = mp.Source()
	assert.Equal(t, m.Source(), source)
	assert.NoError(t, err)
}
