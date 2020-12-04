package acceptance

import (
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestMwsAccWorkspaces(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv != "MWS" {
		t.Skip("Acceptance tests skipped unless CLOUD_ENV=MWS is set")
	}
	config := qa.EnvironmentTemplate(t, `
	provider "databricks" {
		host     = "{env.DATABRICKS_HOST}"
		username = "{env.DATABRICKS_USERNAME}"
		password = "{env.DATABRICKS_PASSWORD}"
	}
	resource "databricks_mws_credentials" "this" {
		account_id       = "{env.DATABRICKS_ACCOUNT_ID}"
		credentials_name = "credentials-ws-{var.RANDOM}"
		role_arn         = "{env.TEST_CROSSACCOUNT_ARN}"
	}
	resource "databricks_mws_customer_managed_keys" "this" {
		account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
		aws_key_info {
			key_arn   = "{env.TEST_KMS_KEY_ARN}"
			key_alias = "{env.TEST_KMS_KEY_ALIAS}"
		}
	}
	resource "databricks_mws_storage_configurations" "this" {
		account_id                 = "{env.DATABRICKS_ACCOUNT_ID}"
		storage_configuration_name = "storage-ws-{var.RANDOM}"
		bucket_name                = "{env.TEST_ROOT_BUCKET}"
	}
	resource "databricks_mws_networks" "this" {
		account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
		network_name = "network-ws-{var.RANDOM}"
		vpc_id       = "{env.TEST_VPC_ID}"
		subnet_ids   = [
			"{env.TEST_SUBNET_PUBLIC}",
			"{env.TEST_SUBNET_PRIVATE}",
		]
		security_group_ids = [
			"{env.TEST_SECURITY_GROUP}",
		]
	}
	resource "databricks_mws_workspaces" "this" {
		account_id      = "{env.DATABRICKS_ACCOUNT_ID}"
		workspace_name  = "terra-{var.RANDOM}"
		deployment_name = "terra-{var.RANDOM}"
		aws_region      = "{env.TEST_REGION}"

		credentials_id = databricks_mws_credentials.this.credentials_id
		storage_configuration_id = databricks_mws_storage_configurations.this.storage_configuration_id
		customer_managed_key_id = databricks_mws_customer_managed_keys.this.customer_managed_key_id
		network_id = databricks_mws_networks.this.network_id
		verify_workspace_runnning = true
	}`)
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("databricks_mws_workspaces.this", "workspace_status", "RUNNING"),
				),
				Destroy: false,
			},
		},
		CheckDestroy: func(s *terraform.State) error {
			// client := common.CommonEnvironmentClient()
			// for _, rs := range s.RootModule().Resources {
			// 	if rs.Type != "databricks_mws_workspaces" {
			// 		continue
			// 	}
			// 	packagedMWSIds, err := UnpackMWSAccountID(rs.Primary.ID)
			// 	if err != nil {
			// 		return err
			// 	}
			// 	idInt64, err := strconv.ParseInt(packagedMWSIds.ResourceID, 10, 64)
			// 	if err != nil {
			// 		return err
			// 	}
			// 	_, err = NewWorkspacesAPI(context.Background(), client).Read(packagedMWSIds.MwsAcctID, idInt64)
			// 	if err != nil {
			// 		return nil
			// 	}
			// 	return errors.New("resource is not cleaned up")
			// }
			return nil
		},
	})
}
