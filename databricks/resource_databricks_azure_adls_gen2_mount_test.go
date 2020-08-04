package databricks

import (
	"os"
	"regexp"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/stretchr/testify/assert"
)

func TestAzureAccAdlsGen2Mount_correctly_mounts(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	resource.Test(t, resource.TestCase{
		Providers:  testAccProviders,
		IsUnitTest: debugIfCloudEnvSet(),
		Steps: []resource.TestStep{
			{
				Config: EnvironmentTemplate(t, `
				resource "databricks_secret_scope" "terraform" {
					name                     = "terraform-{var.RANDOM}"
					initial_manage_principal = "users"
				}
				resource "databricks_secret" "client_secret" {
					key          = "datalake_sp_secret"
					string_value = "{env.ARM_CLIENT_SECRET}"
					scope        = databricks_secret_scope.terraform.name
				}
				resource "databricks_azure_adls_gen2_mount" "mount" {
					container_name         = "dev"
					storage_account_name   = "{env.TEST_GEN2_ADAL_NAME}"
					mount_name             = "localdir{var.RANDOM}"
					tenant_id              = "{env.ARM_TENANT_ID}"
					client_id              = "{env.ARM_CLIENT_ID}"
					client_secret_scope    = databricks_secret_scope.terraform.name
					client_secret_key      = databricks_secret.client_secret.key
					initialize_file_system = true
				}`),
			},
		},
	})
}

func TestAzureAccAdlsGen2Mount_capture_error(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := service.CommonEnvironmentClient()
	if !client.IsAzure() {
		t.Skip("Test is meant only for Azure")
	}
	resource.Test(t, resource.TestCase{
		IsUnitTest: debugIfCloudEnvSet(),
		Providers:  testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: EnvironmentTemplate(t, `
				resource "databricks_secret_scope" "terraform" {
					name                     = "terraform-{var.RANDOM}"
					initial_manage_principal = "users"
				}
				resource "databricks_secret" "client_secret" {
					key          = "datalake_sp_secret"
					string_value = "{env.ARM_CLIENT_SECRET}"
					scope        = databricks_secret_scope.terraform.name
				}
				resource "databricks_azure_adls_gen2_mount" "mount" {
					container_name         = "dev"
					storage_account_name   = "{env.TEST_GEN2_ADAL_NAME}"
					mount_name             = "localdir{var.RANDOM}"
					tenant_id              = "{env.ARM_TENANT_ID}"
					client_id              = "{env.ARM_CLIENT_ID}"
					client_secret_scope    = databricks_secret_scope.terraform.name
					client_secret_key      = "SECRET_KEY_WRONG_ON_PURPOSE"
					initialize_file_system = true
				}`),
				ExpectNonEmptyPlan: true,
				ExpectError:        regexp.MustCompile("java.lang.IllegalArgumentException: Secret does not exist with scope"),
				Destroy:            false,
			},
		},
	})
}

func TestAzureAccADLSv2Mount(t *testing.T) {
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
	gen2AdalName := os.Getenv("TEST_GEN2_ADAL_NAME")
	if gen2AdalName == "" {
		t.Skip("No ADLS account given")
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

	source, err = mp.Mount(AzureADLSGen2Mount{
		ClientID:             client.AzureAuth.ClientID,
		TenantID:             client.AzureAuth.TenantID,
		ContainerName:        "dev",
		Directory:            "/",
		InitializeFileSystem: true,
		SecretKey:            "e",
		SecretScope:          "f",
		StorageAccountName:   gen2AdalName,
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

	err = client.Secrets().Create(client.AzureAuth.ClientSecret, randomScope, "key")
	assert.NoError(t, err)

	m := AzureADLSGen2Mount{
		ClientID:             client.AzureAuth.ClientID,
		TenantID:             client.AzureAuth.TenantID,
		ContainerName:        "dev",
		Directory:            "/",
		InitializeFileSystem: true,
		SecretKey:            "key",
		SecretScope:          randomScope,
		StorageAccountName:   gen2AdalName,
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
