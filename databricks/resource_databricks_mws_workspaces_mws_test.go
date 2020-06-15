package databricks

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccMWSWorkspaces(t *testing.T) {
	//var MWSWorkspaces model.MWSWorkspace
	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest
	//scope := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	mwsAcctID := os.Getenv("DATABRICKS_MWS_ACCT_ID")
	mwsHost := os.Getenv("DATABRICKS_MWS_HOST")
	credentialsName := "tf-workspace-test-creds"
	roleArn := os.Getenv("TEST_MWS_CROSS_ACCT_ROLE")
	storageConfigName := "tf-workspace-storage-config"
	bucketName := os.Getenv("TEST_MWS_ROOT_BUCKET")
	workspaceName := fmt.Sprintf("tf-test-workspace-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	deploymentName := fmt.Sprintf("%s-dep", workspaceName)
	awsRegion := os.Getenv("DATABRICKS_MWS_AWS_REGION")
	networkName := "tf-workspace-test-network"
	vpcID := os.Getenv("TEST_MWS_VPC_ID")
	subnet1 := os.Getenv("TEST_MWS_SUBNET_1")
	subnet2 := os.Getenv("TEST_MWS_SUBNET_2")
	sg := os.Getenv("TEST_MWS_SG")

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testMWSWorkspacesResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testMWSWorkspacesCreate(mwsAcctID, mwsHost, credentialsName, roleArn, storageConfigName, bucketName,
					workspaceName, deploymentName, awsRegion, networkName, vpcID, subnet1, subnet2, sg),
				//// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// verify that after creation the workspace_status is in a running state.
					resource.TestCheckResourceAttr("databricks_mws_workspaces.my_mws_workspace", "workspace_status", "RUNNING"),
				),
				Destroy: false,
			},
		},
	})
}

func testMWSWorkspacesResourceDestroy(s *terraform.State) error {
	client := getMWSClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_mws_workspaces" {
			continue
		}
		packagedMWSIds, err := unpackMWSAccountID(rs.Primary.ID)
		if err != nil {
			return err
		}
		idInt64, err := strconv.ParseInt(packagedMWSIds.ResourceID, 10, 64)
		if err != nil {
			return err
		}
		_, err = client.MWSWorkspaces().Read(packagedMWSIds.MwsAcctID, idInt64)
		if err != nil {
			return nil
		}
		return errors.New("resource Scim Group is not cleaned up")
	}
	return nil
}

func testMWSWorkspacesCreate(mwsAcctID, mwsHost, credentialsName, roleArn, storageConfigName, bucketName,
	workspaceName, deploymentName, awsRegion, networkName, vpcID, subnet1, subnet2, sg string) string {
	return fmt.Sprintf(`
								provider "databricks" {
								  host = "%[1]s"
								  basic_auth {}
								}
								resource "databricks_mws_credentials" "my_mws_credentials" {
								  account_id = "%[2]s"
								  credentials_name = "%[3]s"
								  role_arn         = "%[4]s"
								}
								
								resource "databricks_mws_storage_configurations" "my_mws_storage_configurations" {
								  account_id = "%[2]s"
								  storage_configuration_name = "%[5]s"
								  bucket_name         = "%[6]s"
								}
								resource "databricks_mws_networks" "my_network" {
								  account_id = "%[2]s"
								  network_name = "%[10]s"
								  vpc_id = "%[11]s"
								  subnet_ids = [
									"%[12]s",
									"%[13]s",
								  ]
								  security_group_ids = [
									"%[14]s",
								  ]
								}
								resource "databricks_mws_workspaces" "my_mws_workspace" {
								  account_id = "%[2]s"
								  workspace_name = "%[7]s"
								  deployment_name = "%[8]s"
								  aws_region = "%[9]s"
								  credentials_id = databricks_mws_credentials.my_mws_credentials.credentials_id
								  storage_configuration_id = databricks_mws_storage_configurations.my_mws_storage_configurations.storage_configuration_id
								  network_id = databricks_mws_networks.my_network.network_id
								  verify_workspace_runnning = true
								}
								`, mwsHost, mwsAcctID, credentialsName, roleArn, storageConfigName, bucketName,
		workspaceName, deploymentName, awsRegion, networkName, vpcID, subnet1, subnet2, sg)
}
