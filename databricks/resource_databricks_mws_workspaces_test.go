package databricks

import (
	"errors"
	"os"
	"strconv"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

func TestMwsAccWorkspaces(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv != "MWS" {
		t.Skip("Acceptance tests skipped unless CLOUD_ENV=MWS is set")
	}
	config := EnvironmentTemplate(t, `
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
		network_id = databricks_mws_networks.this.network_id
		verify_workspace_runnning = true
	}`)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
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
			client := service.CommonEnvironmentClient()
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
				return errors.New("resource is not cleaned up")
			}
			return nil
		},
	})
}

func TestResourceMWSWorkspacesCreate(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/accounts/abc/workspaces",
			ExpectedRequest: model.MWSWorkspace{
				IsNoPublicIPEnabled:    true,
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
			Resource: "/api/2.0/accounts/abc/workspaces/1234",
			Response: model.MWSWorkspace{
				WorkspaceStatus: model.WorkspaceStatusRunning,
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/workspaces/1234",
			Response: model.MWSWorkspace{
				WorkspaceStatus: model.WorkspaceStatusRunning,
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/workspaces/1234",
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
			Resource: "/api/2.0/accounts/abc/workspaces/1234",
			Response: model.MWSWorkspace{
				WorkspaceStatus:        model.WorkspaceStatusRunning,
				IsNoPublicIPEnabled:    true,
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
			Resource: "/api/2.0/accounts/abc/workspaces/1234",
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
			Resource: "/api/2.0/accounts/abc/workspaces/1234",
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
				IsNoPublicIPEnabled:    true,
				AwsRegion:              "us-east-1",
				CredentialsID:          "bcd",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/workspaces/1234",
			Response: model.MWSWorkspace{
				WorkspaceStatus:        model.WorkspaceStatusRunning,
				IsNoPublicIPEnabled:    true,
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
			Resource: "/api/2.0/accounts/abc/workspaces/1234",
			Response: model.MWSWorkspace{
				WorkspaceStatus:        model.WorkspaceStatusRunning,
				IsNoPublicIPEnabled:    true,
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
			Resource: "/api/2.0/accounts/abc/workspaces/1234",
			Response: model.MWSWorkspace{
				WorkspaceStatus:        model.WorkspaceStatusRunning,
				IsNoPublicIPEnabled:    true,
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
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/workspaces/1234",
			Response: model.MWSWorkspace{
				WorkspaceName:          "labdata",
				WorkspaceStatus:        model.WorkspaceStatusCanceled,
				WorkspaceStatusMessage: "Things are being removed",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/workspaces/1234",
			Response: service.APIErrorBody{
				ErrorCode: "NOT_FOUND",
				Message:   "Cannot find anything",
			},
			Status: 404,
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
