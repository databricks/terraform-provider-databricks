package acceptance

import (
	"os"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/storage"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAzureAccAdlsGen1Mount_correctly_mounts(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	if !common.CommonEnvironmentClient().IsAzureClientSecretSet() {
		t.Skip("Test is meant only for client-secret conf Azure")
	}
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: acceptance.EnvironmentTemplate(t, `
				resource "databricks_secret_scope" "terraform" {
					name                     = "terraform-{var.RANDOM}"
					initial_manage_principal = "users"
				}
				resource "databricks_secret" "client_secret" {
					key          = "datalake_sp_secret"
					string_value = "{env.ARM_CLIENT_SECRET}"
					scope        = databricks_secret_scope.terraform.name
				}
				resource "databricks_azure_adls_gen1_mount" "mount" {
					storage_resource_name   = "{env.TEST_DATA_LAKE_STORE_NAME}"
					mount_name             = "localdir{var.RANDOM}"
					tenant_id              = "{env.ARM_TENANT_ID}"
					client_id              = "{env.ARM_CLIENT_ID}"
					client_secret_scope    = databricks_secret_scope.terraform.name
					client_secret_key      = databricks_secret.client_secret.key
				}`),
			},
		},
	})
}

func TestAzureAccADLSv1Mount(t *testing.T) {
	client, mp := mountPointThroughReusedCluster(t)
	if !client.IsAzureClientSecretSet() {
		t.Skip("Test is meant only for client-secret conf Azure")
	}
	storageResource := qa.GetEnvOrSkipTest(t, "TEST_DATA_LAKE_STORE_NAME")
	testWithNewSecretScope(t, func(scope, key string) {
		testMounting(t, mp, storage.AzureADLSGen1Mount{
			ClientID:        client.AzureClientID,
			TenantID:        client.AzureTenantID,
			PrefixType:      "dfs.adls",
			StorageResource: storageResource,
			Directory:       "/",
			SecretScope:     scope,
			SecretKey:       key,
		})
	}, client, mp.Name, client.AzureClientSecret)
}
