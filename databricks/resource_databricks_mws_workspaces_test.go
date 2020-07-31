package databricks

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccMWSWorkspaces(t *testing.T) {
	// TODO: refactor for common instance pool & AZ CLI
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

func TestResourceMWSWorkspacesCreate(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/accounts/abc/workspaces",
			ExpectedRequest: model.MWSWorkspace{
				IsNoPublicIpEnabled:    true,
				WorkspaceName:          "labdata",
				DeploymentName:         "900150983cd24fb0",
				AwsRegion:              "us-east-1",
				CredentialsID:          "bcd",
				StorageConfigurationID: "ghi",
				NetworkID:              "fgh",
				CustomerManagedKeyID:   "def",
			},
			Response: model.MWSWorkspace{
				WorkspaceID:    1234,
				DeploymentName: "900150983cd24fb0",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/workspaces/1234?",
			Response: model.MWSWorkspace{
				WorkspaceStatus: model.WorkspaceStatusRunning,
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/workspaces/1234?",
			Response: model.MWSWorkspace{
				WorkspaceStatus: model.WorkspaceStatusRunning,
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/workspaces/1234?",
			Response: model.MWSWorkspace{
				WorkspaceStatus: model.WorkspaceStatusRunning,
			},
		},
	}, resourceMWSWorkspaces, map[string]interface{}{
		"account_id":                "abc",
		"aws_region":                "us-east-1",
		"credentials_id":            "bcd",
		"customer_managed_key_id":   "def",
		"deployment_name":           "900150983cd24fb0",
		"workspace_name":            "labdata",
		"is_no_public_ip_enabled":   true,
		"network_id":                "fgh",
		"storage_configuration_id":  "ghi",
		"verify_workspace_runnning": false,
	}, resourceMWSWorkspacesCreate)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/1234", d.Id())
}

func TestResourceMWSWorkspacesCreate_Error(t *testing.T) {
	t.Skipf("Making this test skip until we can configure sleep timings for test purposes")
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/accounts/abc/workspaces",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/accounts/abc/workspaces",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceMWSWorkspaces, map[string]interface{}{
		"account_id":                "abc",
		"aws_region":                "us-east-1",
		"credentials_id":            "bcd",
		"customer_managed_key_id":   "def",
		"deployment_name":           "900150983cd24fb0",
		"workspace_name":            "labdata",
		"is_no_public_ip_enabled":   true,
		"network_id":                "fgh",
		"storage_configuration_id":  "ghi",
		"verify_workspace_runnning": false,
	}, resourceMWSWorkspacesCreate)
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceMWSWorkspacesRead(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/workspaces/1234?",
			Response: model.MWSWorkspace{
				WorkspaceStatus:        model.WorkspaceStatusRunning,
				IsNoPublicIpEnabled:    true,
				WorkspaceName:          "labdata",
				DeploymentName:         "900150983cd24fb0",
				AwsRegion:              "us-east-1",
				CredentialsID:          "bcd",
				StorageConfigurationID: "ghi",
				NetworkID:              "fgh",
				CustomerManagedKeyID:   "def",
				WorkspaceID:            1234,
			},
		},
	}, resourceMWSWorkspaces, nil, actionWithID("abc/1234", resourceMWSWorkspacesRead))
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/1234", d.Id(), "Id should not be empty")
	assert.Equal(t, "us-east-1", d.Get("aws_region"))
	assert.Equal(t, "bcd", d.Get("credentials_id"))
	assert.Equal(t, "def", d.Get("customer_managed_key_id"))
	assert.Equal(t, "900150983cd24fb0", d.Get("deployment_name"))
	assert.Equal(t, false, d.Get("is_no_public_ip_enabled"))
	assert.Equal(t, "fgh", d.Get("network_id"))
	assert.Equal(t, "ghi", d.Get("storage_configuration_id"))
	assert.Equal(t, false, d.Get("verify_workspace_runnning"))
	assert.Equal(t, 1234, d.Get("workspace_id"))
	assert.Equal(t, "labdata", d.Get("workspace_name"))
	assert.Equal(t, "RUNNING", d.Get("workspace_status"))
	assert.Equal(t, "https://900150983cd24fb0.cloud.databricks.com", d.Get("workspace_url"))
}

func TestResourceMWSWorkspacesRead_NotFound(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/workspaces/1234?",
			Response: service.APIErrorBody{
				ErrorCode: "NOT_FOUND",
				Message:   "Item not found",
			},
			Status: 404,
		},
	}, resourceMWSWorkspaces, nil, actionWithID("abc/1234", resourceMWSWorkspacesRead))
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}

func TestResourceMWSWorkspacesRead_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{ // read log output for correct url...
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/workspaces/1234?",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceMWSWorkspaces, nil, actionWithID("abc/1234", resourceMWSWorkspacesRead))
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/1234", d.Id(), "Id should not be empty for error reads")
}

func TestResourceMWSWorkspacesUpdate(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "PATCH",
			Resource: "/api/2.0/accounts/abc/workspaces/1234",
			ExpectedRequest: model.MWSWorkspace{
				StorageConfigurationID: "ghi",
				NetworkID:              "fgh",
				CustomerManagedKeyID:   "def",
				IsNoPublicIpEnabled:    true,
				AwsRegion:              "us-east-1",
				CredentialsID:          "bcd",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/workspaces/1234?",
			Response: model.MWSWorkspace{
				WorkspaceStatus:        model.WorkspaceStatusRunning,
				IsNoPublicIpEnabled:    true,
				WorkspaceName:          "labdata",
				DeploymentName:         "900150983cd24fb0",
				AwsRegion:              "us-east-1",
				CredentialsID:          "bcd",
				StorageConfigurationID: "ghi",
				NetworkID:              "fgh",
				CustomerManagedKeyID:   "def",
				WorkspaceID:            1234,
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/workspaces/1234?",
			Response: model.MWSWorkspace{
				WorkspaceStatus:        model.WorkspaceStatusRunning,
				IsNoPublicIpEnabled:    true,
				WorkspaceName:          "labdata",
				DeploymentName:         "900150983cd24fb0",
				AwsRegion:              "us-east-1",
				CredentialsID:          "bcd",
				StorageConfigurationID: "ghi",
				NetworkID:              "fgh",
				CustomerManagedKeyID:   "def",
				WorkspaceID:            1234,
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/workspaces/1234?",
			Response: model.MWSWorkspace{
				WorkspaceStatus:        model.WorkspaceStatusRunning,
				IsNoPublicIpEnabled:    true,
				WorkspaceName:          "labdata",
				DeploymentName:         "900150983cd24fb0",
				AwsRegion:              "us-east-1",
				CredentialsID:          "bcd",
				StorageConfigurationID: "ghi",
				NetworkID:              "fgh",
				CustomerManagedKeyID:   "def",
				WorkspaceID:            1234,
			},
		},
	}, resourceMWSWorkspaces, map[string]interface{}{
		"account_id":                "abc",
		"aws_region":                "us-east-1",
		"credentials_id":            "bcd",
		"customer_managed_key_id":   "def",
		"deployment_name":           "900150983cd24fb0",
		"workspace_name":            "labdata",
		"is_no_public_ip_enabled":   true,
		"network_id":                "fgh",
		"storage_configuration_id":  "ghi",
		"verify_workspace_runnning": false,
	}, actionWithID("abc/1234", resourceMWSWorkspacesUpdate))
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/1234", d.Id(), "Id should be the same as in reading")
}

func TestResourceMWSWorkspacesUpdate_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "PATCH",
			Resource: "/api/2.0/accounts/abc/workspaces/1234",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceMWSWorkspaces, map[string]interface{}{
		"account_id":                "abc",
		"aws_region":                "us-east-1",
		"credentials_id":            "bcd",
		"customer_managed_key_id":   "def",
		"deployment_name":           "900150983cd24fb0",
		"workspace_name":            "labdata",
		"is_no_public_ip_enabled":   true,
		"network_id":                "fgh",
		"storage_configuration_id":  "ghi",
		"verify_workspace_runnning": false,
	}, actionWithID("abc/1234", resourceMWSWorkspacesUpdate))
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/1234", d.Id())
}

func TestResourceMWSWorkspacesDelete(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "DELETE",
			Resource: "/api/2.0/accounts/abc/workspaces/1234",
		},
	}, resourceMWSWorkspaces, nil, actionWithID("abc/1234", resourceMWSWorkspacesDelete))
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/1234", d.Id())
}

func TestResourceMWSWorkspacesDelete_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "DELETE",
			Resource: "/api/2.0/accounts/abc/workspaces/1234",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceMWSWorkspaces, nil, actionWithID("abc/1234", resourceMWSWorkspacesDelete))
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/1234", d.Id())
}
