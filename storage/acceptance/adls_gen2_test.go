package acceptance

import (
	"os"
	"regexp"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAzureAccAdlsGen2Mount_correctly_mounts(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	if !common.CommonEnvironmentClient().AzureAuth.IsClientSecretSet() {
		t.Skip("Test is meant only for client-secret conf Azure")
	}
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: qa.EnvironmentTemplate(t, `
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
					storage_account_name   = "{env.TEST_STORAGE_V2_ACCOUNT}"
					container_name         = "{env.TEST_STORAGE_V2_ABFSS}"
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
	client := common.CommonEnvironmentClient()
	if !client.AzureAuth.IsClientSecretSet() {
		t.Skip("Test is meant only for SP Azure")
	}
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: qa.EnvironmentTemplate(t, `
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
					storage_account_name   = "{env.TEST_STORAGE_V2_ACCOUNT}"
					container_name         = "{env.TEST_STORAGE_V2_ABFSS}"
					mount_name             = "localdir{var.RANDOM}"
					tenant_id              = "{env.ARM_TENANT_ID}"
					client_id              = "{env.ARM_CLIENT_ID}"
					client_secret_scope    = databricks_secret_scope.terraform.name
					client_secret_key      = "SECRET_KEY_WRONG_ON_PURPOSE"
					initialize_file_system = true
				}`),
				ExpectNonEmptyPlan: true,
				ExpectError:        regexp.MustCompile("Secret does not exist with scope"),
				Destroy:            false,
			},
		},
	})
}
