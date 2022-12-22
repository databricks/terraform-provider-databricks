package mws

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/tokens"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMwsAccWorkspace(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	acctID := qa.GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	client := common.CommonEnvironmentClient()
	workspaceList, err := NewWorkspacesAPI(context.Background(), client).List(acctID)
	assert.NoError(t, err, err)
	t.Log(workspaceList)
}

func TestGcpaAccWorkspace(t *testing.T) {
	acctID := qa.GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	client := common.CommonEnvironmentClient()
	workspacesAPI := NewWorkspacesAPI(context.Background(), client)

	workspaceList, err := workspacesAPI.List(acctID)
	require.NoError(t, err, err)
	t.Log(workspaceList)

	ws := Workspace{
		AccountID:     acctID,
		WorkspaceName: qa.RandomName(qa.GetEnvOrSkipTest(t, "TEST_PREFIX") + "-"),
		Location:      qa.GetEnvOrSkipTest(t, "GOOGLE_REGION"),
		CloudResourceBucket: &CloudResourceContainer{
			GCP: &GCP{
				ProjectID: qa.GetEnvOrSkipTest(t, "GOOGLE_PROJECT"),
			},
		},
	}
	err = workspacesAPI.Create(&ws, 5*time.Minute)
	require.NoError(t, err)

	err = workspacesAPI.Delete(acctID, fmt.Sprintf("%d", ws.WorkspaceID))
	require.NoError(t, err)
}

func TestResourceWorkspaceCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/workspaces",
				ExpectedRequest: Workspace{
					AccountID:                           "abc",
					IsNoPublicIPEnabled:                 true,
					WorkspaceName:                       "labdata",
					DeploymentName:                      "900150983cd24fb0",
					AwsRegion:                           "us-east-1",
					CredentialsID:                       "bcd",
					StorageConfigurationID:              "ghi",
					NetworkID:                           "fgh",
					ManagedServicesCustomerManagedKeyID: "def",
					StorageCustomerManagedKeyID:         "def",
				},
				Response: Workspace{
					WorkspaceID:    1234,
					AccountID:      "abc",
					DeploymentName: "900150983cd24fb0",
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/accounts/abc/workspaces/1234",
				Response: Workspace{
					WorkspaceID:                         1234,
					WorkspaceStatus:                     WorkspaceStatusRunning,
					WorkspaceName:                       "labdata",
					DeploymentName:                      "900150983cd24fb0",
					AwsRegion:                           "us-east-1",
					CredentialsID:                       "bcd",
					StorageConfigurationID:              "ghi",
					NetworkID:                           "fgh",
					ManagedServicesCustomerManagedKeyID: "def",
					StorageCustomerManagedKeyID:         "def",
					AccountID:                           "abc",
				},
			},
		},
		Resource: ResourceMwsWorkspaces(),
		State: map[string]any{
			"account_id":     "abc",
			"aws_region":     "us-east-1",
			"credentials_id": "bcd",
			"managed_services_customer_managed_key_id": "def",
			"storage_customer_managed_key_id":          "def",
			"deployment_name":                          "900150983cd24fb0",
			"workspace_name":                           "labdata",
			"network_id":                               "fgh",
			"storage_configuration_id":                 "ghi",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/1234", d.Id())
}

func TestResourceWorkspaceCreateGcp(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/workspaces",
				// retreating to raw JSON, as certain fields don't work well together
				ExpectedRequest: map[string]any{
					"account_id": "abc",
					"cloud":      "gcp",
					"cloud_resource_container": map[string]any{
						"gcp": map[string]any{
							"project_id": "def",
						},
					},
					"location":   "bcd",
					"network_id": "net_id_a",
					"gke_config": map[string]any{
						"master_ip_range":   "e",
						"connectivity_type": "d",
					},
					"gcp_managed_network_config": map[string]any{
						"gke_cluster_pod_ip_range":     "b",
						"gke_cluster_service_ip_range": "c",
						"subnet_cidr":                  "a",
					},
					"workspace_name": "labdata",
				},
				Response: Workspace{
					WorkspaceID:    1234,
					AccountID:      "abc",
					DeploymentName: "900150983cd24fb0",
					WorkspaceName:  "labdata",
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/accounts/abc/workspaces/1234",
				Response: Workspace{
					AccountID:       "abc",
					WorkspaceID:     1234,
					WorkspaceStatus: WorkspaceStatusRunning,
					DeploymentName:  "900150983cd24fb0",
					WorkspaceName:   "labdata",
				},
			},
		},
		Resource: ResourceMwsWorkspaces(),
		HCL: `
		account_id      = "abc"
		workspace_name  = "labdata"
		deployment_name = "900150983cd24fb0"
		location        = "bcd"
		cloud_resource_container {
			gcp {
				project_id = "def"
			}
		}
		network_id = "net_id_a"
		gcp_managed_network_config {
			subnet_cidr = "a"
			gke_cluster_pod_ip_range = "b"
			gke_cluster_service_ip_range = "c"
		}
		gke_config {
			connectivity_type = "d"
			master_ip_range = "e"
		}
		`,
		Gcp:    true,
		Create: true,
	}.ApplyNoError(t)
}

func TestResourceWorkspaceCreateWithIsNoPublicIPEnabledFalse(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/workspaces",
				ExpectedRequest: Workspace{
					AccountID:                           "abc",
					IsNoPublicIPEnabled:                 false,
					WorkspaceName:                       "labdata",
					DeploymentName:                      "900150983cd24fb0",
					AwsRegion:                           "us-east-1",
					CredentialsID:                       "bcd",
					StorageConfigurationID:              "ghi",
					NetworkID:                           "fgh",
					ManagedServicesCustomerManagedKeyID: "def",
					StorageCustomerManagedKeyID:         "def",
				},
				Response: Workspace{
					WorkspaceID:    1234,
					AccountID:      "abc",
					DeploymentName: "900150983cd24fb0",
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/accounts/abc/workspaces/1234",
				Response: Workspace{
					WorkspaceID:                         1234,
					WorkspaceStatus:                     WorkspaceStatusRunning,
					WorkspaceName:                       "labdata",
					DeploymentName:                      "900150983cd24fb0",
					AwsRegion:                           "us-east-1",
					CredentialsID:                       "bcd",
					StorageConfigurationID:              "ghi",
					NetworkID:                           "fgh",
					ManagedServicesCustomerManagedKeyID: "def",
					StorageCustomerManagedKeyID:         "def",
					AccountID:                           "abc",
				},
			},
		},
		Resource: ResourceMwsWorkspaces(),
		State: map[string]any{
			"account_id":     "abc",
			"aws_region":     "us-east-1",
			"credentials_id": "bcd",
			"managed_services_customer_managed_key_id": "def",
			"storage_customer_managed_key_id":          "def",
			"deployment_name":                          "900150983cd24fb0",
			"workspace_name":                           "labdata",
			"is_no_public_ip_enabled":                  false,
			"network_id":                               "fgh",
			"storage_configuration_id":                 "ghi",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/1234", d.Id())
}

func TestResourceWorkspaceCreateLegacyConfig(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/workspaces",
				ExpectedRequest: Workspace{
					AccountID:                           "abc",
					IsNoPublicIPEnabled:                 true,
					WorkspaceName:                       "labdata",
					DeploymentName:                      "900150983cd24fb0",
					AwsRegion:                           "us-east-1",
					CredentialsID:                       "bcd",
					StorageConfigurationID:              "ghi",
					NetworkID:                           "fgh",
					ManagedServicesCustomerManagedKeyID: "def",
				},
				Response: Workspace{
					WorkspaceID:    1234,
					AccountID:      "abc",
					DeploymentName: "900150983cd24fb0",
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/accounts/abc/workspaces/1234",
				Response: Workspace{
					WorkspaceID:                         1234,
					WorkspaceStatus:                     WorkspaceStatusRunning,
					WorkspaceName:                       "labdata",
					DeploymentName:                      "900150983cd24fb0",
					AwsRegion:                           "us-east-1",
					CredentialsID:                       "bcd",
					StorageConfigurationID:              "ghi",
					NetworkID:                           "fgh",
					ManagedServicesCustomerManagedKeyID: "def",
					AccountID:                           "abc",
				},
			},
		},
		Resource: ResourceMwsWorkspaces(),
		State: map[string]any{
			"account_id":               "abc",
			"aws_region":               "us-east-1",
			"credentials_id":           "bcd",
			"customer_managed_key_id":  "def",
			"deployment_name":          "900150983cd24fb0",
			"workspace_name":           "labdata",
			"network_id":               "fgh",
			"storage_configuration_id": "ghi",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/1234", d.Id())
}

func TestResourceWorkspaceCreate_Error(t *testing.T) {
	t.Skipf("Making this test skip until we can configure sleep timings for test purposes")
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/workspaces",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/workspaces",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceMwsWorkspaces(),
		State: map[string]any{
			"account_id":     "abc",
			"aws_region":     "us-east-1",
			"credentials_id": "bcd",
			"managed_services_customer_managed_key_id": "def",
			"storage_customer_managed_key_id":          "def",
			"deployment_name":                          "900150983cd24fb0",
			"workspace_name":                           "labdata",
			"is_no_public_ip_enabled":                  true,
			"network_id":                               "fgh",
			"storage_configuration_id":                 "ghi",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceWorkspaceRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/accounts/abc/workspaces/1234",
				Response: Workspace{
					AccountID:                           "abc",
					WorkspaceStatus:                     WorkspaceStatusRunning,
					WorkspaceName:                       "labdata",
					DeploymentName:                      "900150983cd24fb0",
					AwsRegion:                           "us-east-1",
					CredentialsID:                       "bcd",
					StorageConfigurationID:              "ghi",
					NetworkID:                           "fgh",
					ManagedServicesCustomerManagedKeyID: "def",
					StorageCustomerManagedKeyID:         "def",
					WorkspaceID:                         1234,
				},
			},
		},
		Resource: ResourceMwsWorkspaces(),
		Read:     true,
		New:      true,
		ID:       "abc/1234",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/1234", d.Id(), "Id should not be empty")
	assert.Equal(t, "us-east-1", d.Get("aws_region"))
	assert.Equal(t, "bcd", d.Get("credentials_id"))
	assert.Equal(t, "def", d.Get("managed_services_customer_managed_key_id"))
	assert.Equal(t, "def", d.Get("storage_customer_managed_key_id"))
	assert.Equal(t, "900150983cd24fb0", d.Get("deployment_name"))
	assert.Equal(t, true, d.Get("is_no_public_ip_enabled"))
	assert.Equal(t, "fgh", d.Get("network_id"))
	assert.Equal(t, "ghi", d.Get("storage_configuration_id"))
	assert.Equal(t, 1234, d.Get("workspace_id"))
	assert.Equal(t, "labdata", d.Get("workspace_name"))
	assert.Equal(t, "RUNNING", d.Get("workspace_status"))
	assert.Equal(t, "https://900150983cd24fb0.cloud.databricks.com", d.Get("workspace_url"))
}

func TestResourceWorkspaceRead_Issue382(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/accounts/abc/workspaces/1234",
				Response: Workspace{
					AccountID:                           "abc",
					WorkspaceStatus:                     WorkspaceStatusRunning,
					WorkspaceName:                       "labdata",
					DeploymentName:                      "prefix-900150983cd24fb0",
					AwsRegion:                           "us-east-1",
					CredentialsID:                       "bcd",
					StorageConfigurationID:              "ghi",
					NetworkID:                           "fgh",
					ManagedServicesCustomerManagedKeyID: "def",
					StorageCustomerManagedKeyID:         "def",
					WorkspaceID:                         1234,
				},
			},
		},
		InstanceState: map[string]string{
			"account_id":     "abc",
			"aws_region":     "us-east-1",
			"credentials_id": "bcd",
			"managed_services_customer_managed_key_id": "def",
			"storage_customer_managed_key_id":          "def",
			"deployment_name":                          "900150983cd24fb0",
			"workspace_name":                           "labdata",
			"network_id":                               "fgh",
			"storage_configuration_id":                 "ghi",
		},
		State: map[string]any{
			"account_id":     "abc",
			"aws_region":     "us-east-1",
			"credentials_id": "bcd",
			"managed_services_customer_managed_key_id": "def",
			"storage_customer_managed_key_id":          "def",
			"deployment_name":                          "900150983cd24fb0",
			"workspace_name":                           "labdata",
			"network_id":                               "fgh",
			"storage_configuration_id":                 "ghi",
		},
		Resource: ResourceMwsWorkspaces(),
		Read:     true,
		New:      true,
		ID:       "abc/1234",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/1234", d.Id(), "Id should not be empty")
	assert.Equal(t, "prefix-900150983cd24fb0", d.Get("deployment_name"))
	assert.Equal(t, "https://prefix-900150983cd24fb0.cloud.databricks.com", d.Get("workspace_url"))
}

func TestResourceWorkspaceRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/workspaces/1234",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceMwsWorkspaces(),
		Read:     true,
		Removed:  true,
		ID:       "abc/1234",
	}.ApplyNoError(t)
}

func TestResourceWorkspaceRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for correct url...
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/workspaces/1234",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceMwsWorkspaces(),
		Read:     true,
		ID:       "abc/1234",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/1234", d.Id(), "Id should not be empty for error reads")
}

func TestResourceWorkspaceUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/accounts/abc/workspaces/1234",
				ExpectedRequest: map[string]any{
					"credentials_id":                  "bcd",
					"network_id":                      "fgh",
					"storage_customer_managed_key_id": "def",
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/accounts/abc/workspaces/1234",
				Response: Workspace{
					WorkspaceStatus:                     WorkspaceStatusRunning,
					WorkspaceName:                       "labdata",
					DeploymentName:                      "900150983cd24fb0",
					AwsRegion:                           "us-east-1",
					CredentialsID:                       "bcd",
					StorageConfigurationID:              "ghi",
					NetworkID:                           "fgh",
					ManagedServicesCustomerManagedKeyID: "def",
					StorageCustomerManagedKeyID:         "def",
					AccountID:                           "abc",
					WorkspaceID:                         1234,
				},
			},
		},
		Resource: ResourceMwsWorkspaces(),
		InstanceState: map[string]string{
			"account_id":     "abc",
			"aws_region":     "us-east-1",
			"credentials_id": "__OLDER__",
			"managed_services_customer_managed_key_id": "def",
			"storage_customer_managed_key_id":          "__OLDER__",
			"deployment_name":                          "900150983cd24fb0",
			"workspace_name":                           "labdata",
			"is_no_public_ip_enabled":                  "true",
			"network_id":                               "fgh",
			"storage_configuration_id":                 "ghi",
			"workspace_id":                             "1234",
		},
		State: map[string]any{
			"account_id":     "abc",
			"aws_region":     "us-east-1",
			"credentials_id": "bcd",
			"managed_services_customer_managed_key_id": "def",
			"storage_customer_managed_key_id":          "def",
			"deployment_name":                          "900150983cd24fb0",
			"workspace_name":                           "labdata",
			"is_no_public_ip_enabled":                  true,
			"network_id":                               "fgh",
			"storage_configuration_id":                 "ghi",
			"workspace_id":                             1234,
		},
		Update: true,
		ID:     "abc/1234",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/1234", d.Id(), "Id should be the same as in reading")
}

func TestResourceWorkspaceUpdate_NotAllowed(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceMwsWorkspaces(),
		InstanceState: map[string]string{
			"account_id":     "abc",
			"aws_region":     "us-east-1",
			"credentials_id": "bcd",
			"managed_services_customer_managed_key_id": "def",
			"storage_customer_managed_key_id":          "def",
			"deployment_name":                          "900150983cd24fb0",
			"workspace_name":                           "labdata",
			"is_no_public_ip_enabled":                  "true",
			"network_id":                               "fgh",
			"storage_configuration_id":                 "ghi",
			"workspace_id":                             "1234",
		},
		State: map[string]any{
			"account_id": "THIS_IS_CHANGING",

			"aws_region":     "us-east-1",
			"credentials_id": "bcd",
			"managed_services_customer_managed_key_id": "def",
			"storage_customer_managed_key_id":          "def",
			"deployment_name":                          "900150983cd24fb0",
			"workspace_name":                           "labdata",
			"is_no_public_ip_enabled":                  true,
			"network_id":                               "fgh",
			"storage_configuration_id":                 "ghi",
			"workspace_id":                             1234,
		},
		Update: true,
		ID:     "abc/1234",
	}.ExpectError(t, "changes require new: account_id")
}

func TestResourceWorkspaceUpdateLegacyConfig(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/accounts/abc/workspaces/1234",
				ExpectedRequest: map[string]any{
					"credentials_id": "bcd",
					"network_id":     "fgh",
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/accounts/abc/workspaces/1234",
				Response: Workspace{
					WorkspaceStatus:                     WorkspaceStatusRunning,
					IsNoPublicIPEnabled:                 true,
					WorkspaceName:                       "labdata",
					DeploymentName:                      "900150983cd24fb0",
					AwsRegion:                           "us-east-1",
					CredentialsID:                       "bcd",
					StorageConfigurationID:              "ghi",
					NetworkID:                           "fgh",
					ManagedServicesCustomerManagedKeyID: "def",
					AccountID:                           "abc",
					WorkspaceID:                         1234,
				},
			},
		},
		Resource: ResourceMwsWorkspaces(),
		InstanceState: map[string]string{
			"account_id":               "abc",
			"aws_region":               "us-east-1",
			"credentials_id":           "bcd",
			"customer_managed_key_id":  "def",
			"deployment_name":          "900150983cd24fb0",
			"is_no_public_ip_enabled":  "true",
			"workspace_name":           "labdata",
			"network_id":               "fgh",
			"storage_configuration_id": "ghi",
			"workspace_id":             "1234",
		},
		State: map[string]any{
			"account_id":               "abc",
			"aws_region":               "us-east-1",
			"credentials_id":           "bcd",
			"customer_managed_key_id":  "def",
			"deployment_name":          "900150983cd24fb0",
			"workspace_name":           "labdata",
			"network_id":               "fgh",
			"storage_configuration_id": "ghi",
			"workspace_id":             1234,
		},
		Update: true,
		ID:     "abc/1234",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/1234", d.Id(), "Id should be the same as in reading")
}

func TestResourceWorkspaceUpdate_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/accounts/abc/workspaces/1234",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceMwsWorkspaces(),
		State: map[string]any{
			"account_id":     "abc",
			"aws_region":     "us-east-1",
			"credentials_id": "bcd",
			"managed_services_customer_managed_key_id": "def",
			"storage_customer_managed_key_id":          "def",
			"deployment_name":                          "900150983cd24fb0",
			"workspace_name":                           "labdata",
			"network_id":                               "fgh",
			"storage_configuration_id":                 "ghi",
			"workspace_id":                             1234,
		},
		Update:      true,
		RequiresNew: true,
		ID:          "abc/1234",
	}.ExpectError(t, "Internal error happened")
}

func TestResourceWorkspaceDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/accounts/abc/workspaces/1234",
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/workspaces/1234",
				Response: Workspace{
					WorkspaceName:          "labdata",
					WorkspaceStatus:        WorkspaceStatusCanceled,
					WorkspaceStatusMessage: "Things are being removed",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/workspaces/1234",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Cannot find anything",
				},
				Status: 404,
			},
		},
		Resource: ResourceMwsWorkspaces(),
		Delete:   true,
		ID:       "abc/1234",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/1234", d.Id())
}

func TestResourceWorkspaceDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/accounts/abc/workspaces/1234",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceMwsWorkspaces(),
		Delete:   true,
		ID:       "abc/1234",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/1234", d.Id())
}

func TestWaitForRunning(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/accounts/abc/workspaces",
			ExpectedRequest: Workspace{
				AccountID:                           "abc",
				IsNoPublicIPEnabled:                 true,
				WorkspaceName:                       "labdata",
				DeploymentName:                      "900150983cd24fb0",
				AwsRegion:                           "us-east-1",
				CredentialsID:                       "bcd",
				StorageConfigurationID:              "ghi",
				NetworkID:                           "fgh",
				ManagedServicesCustomerManagedKeyID: "def",
				StorageCustomerManagedKeyID:         "def",
			},
			Response: Workspace{
				WorkspaceID:    1234,
				AccountID:      "abc",
				DeploymentName: "900150983cd24fb0",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/workspaces/1234",
			Response: Workspace{
				WorkspaceID:                         1234,
				WorkspaceStatus:                     WorkspaceStatusProvisioning,
				WorkspaceName:                       "labdata",
				DeploymentName:                      "900150983cd24fb0",
				AwsRegion:                           "us-east-1",
				CredentialsID:                       "bcd",
				StorageConfigurationID:              "ghi",
				NetworkID:                           "fgh",
				ManagedServicesCustomerManagedKeyID: "def",
				StorageCustomerManagedKeyID:         "def",
				AccountID:                           "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/workspaces/1234",
			Response: Workspace{
				WorkspaceID:                         1234,
				WorkspaceStatus:                     WorkspaceStatusRunning,
				WorkspaceName:                       "labdata",
				DeploymentName:                      "900150983cd24fb0",
				AwsRegion:                           "us-east-1",
				CredentialsID:                       "bcd",
				StorageConfigurationID:              "ghi",
				NetworkID:                           "fgh",
				ManagedServicesCustomerManagedKeyID: "def",
				StorageCustomerManagedKeyID:         "def",
				AccountID:                           "abc",
			},
		},
	})
	require.NoError(t, err)
	defer server.Close()

	err = NewWorkspacesAPI(context.Background(), client).Create(&Workspace{
		AccountID:                           "abc",
		IsNoPublicIPEnabled:                 true,
		WorkspaceName:                       "labdata",
		DeploymentName:                      "900150983cd24fb0",
		AwsRegion:                           "us-east-1",
		CredentialsID:                       "bcd",
		StorageConfigurationID:              "ghi",
		NetworkID:                           "fgh",
		ManagedServicesCustomerManagedKeyID: "def",
		StorageCustomerManagedKeyID:         "def",
	}, DefaultProvisionTimeout)
	require.NoError(t, err)
}

func TestCreateFailsAndCleansUp(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/accounts/abc/workspaces",
			ExpectedRequest: Workspace{
				AccountID:                           "abc",
				IsNoPublicIPEnabled:                 true,
				WorkspaceName:                       "labdata",
				DeploymentName:                      "900150983cd24fb0",
				AwsRegion:                           "us-east-1",
				CredentialsID:                       "bcd",
				StorageConfigurationID:              "ghi",
				NetworkID:                           "fgh",
				ManagedServicesCustomerManagedKeyID: "def",
				StorageCustomerManagedKeyID:         "def",
			},
			Response: Workspace{
				WorkspaceID:    1234,
				AccountID:      "abc",
				DeploymentName: "900150983cd24fb0",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/workspaces/1234",
			Response: Workspace{
				WorkspaceID:            1234,
				WorkspaceStatus:        WorkspaceStatusFailed,
				WorkspaceStatusMessage: "Always fails",
				WorkspaceName:          "labdata",
				DeploymentName:         "900150983cd24fb0",
				AwsRegion:              "us-east-1",
				NetworkID:              "fgh",
				AccountID:              "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/networks/fgh",
			Response: Network{
				ErrorMessages: []NetworkHealth{
					{"FAIL", "Message"},
				},
			},
		},
		{
			Method:   "DELETE",
			Resource: "/api/2.0/accounts/abc/workspaces/1234",
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/workspaces/1234",
			Status:   404,
		},
	})
	require.NoError(t, err)
	defer server.Close()

	err = NewWorkspacesAPI(context.Background(), client).Create(&Workspace{
		AccountID:                           "abc",
		IsNoPublicIPEnabled:                 true,
		WorkspaceName:                       "labdata",
		DeploymentName:                      "900150983cd24fb0",
		AwsRegion:                           "us-east-1",
		CredentialsID:                       "bcd",
		StorageConfigurationID:              "ghi",
		NetworkID:                           "fgh",
		ManagedServicesCustomerManagedKeyID: "def",
		StorageCustomerManagedKeyID:         "def",
	}, DefaultProvisionTimeout)
	require.EqualError(t, err, "Workspace failed to create: Always fails, network error message: error: FAIL;error_msg: Message;")
}

func TestListWorkspaces(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/workspaces",
			Response: []Workspace{},
		},
	})
	require.NoError(t, err)
	defer server.Close()

	l, err := NewWorkspacesAPI(context.Background(), client).List("abc")
	require.NoError(t, err)
	assert.Len(t, l, 0)
}

func TestWorkspace_WaitForResolve_Failure(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{},
		func(ctx context.Context, client *common.DatabricksClient) {
			a := NewWorkspacesAPI(ctx, client)
			rerr := a.verifyWorkspaceReachable(Workspace{
				WorkspaceURL: "https://900150983cd24fb0.cloud.databricks.com",
			})
			assert.NotNil(t, rerr)
			assert.True(t, rerr.Retryable)
		})
}

func TestWorkspace_WaitForResolve(t *testing.T) {
	// outer HTTP server is used for inner request for "just created" workspace
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/token/list",
			Response: `{}`, // we just need a JSON for this
		},
	}, func(ctx context.Context, wsClient *common.DatabricksClient) {
		// inner HTTP server is used for outer request for Accounts API
		qa.HTTPFixturesApply(t, []qa.HTTPFixture{
			{
				Method:       "GET",
				Resource:     "/api/2.0/accounts/abc/workspaces/1234",
				ReuseRequest: true,
				Response: Workspace{
					AccountID:       "abc",
					WorkspaceID:     1234,
					WorkspaceStatus: "RUNNING",
					WorkspaceURL:    wsClient.Host,
				},
			},
		}, func(ctx context.Context, client *common.DatabricksClient) {
			a := NewWorkspacesAPI(ctx, client)
			err := a.WaitForRunning(Workspace{
				AccountID:   "abc",
				WorkspaceID: 1234,
			}, 1*time.Second)
			assert.NoError(t, err)
		})
	})
}

func updateWorkspaceTokenFixture(t *testing.T, fixtures []qa.HTTPFixture, state map[string]string, hcl string) {
	accountsAPI := []qa.HTTPFixture{
		{
			Method:   "PATCH",
			Resource: "/api/2.0/accounts/c/workspaces/0",
		},
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.0/accounts/c/workspaces/0",
		},
	}
	tokensAPI := []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/token/list",
			Response: `{}`, // we just need a JSON for this
		},
	}
	tokensAPI = append(tokensAPI, fixtures...)
	// outer HTTP server is used for inner request for "just created" workspace
	qa.HTTPFixturesApply(t, tokensAPI, func(ctx context.Context, wsClient *common.DatabricksClient) {
		// a bit hacky, but the whole thing is more readable
		accountsAPI[1].Response = Workspace{
			WorkspaceStatus: "RUNNING",
			WorkspaceURL:    wsClient.Host,
		}
		state["workspace_url"] = wsClient.Host
		state["workspace_name"] = "b"
		state["account_id"] = "c"
		state["is_no_public_ip_enabled"] = "false"
		qa.ResourceFixture{
			Fixtures:      accountsAPI,
			Resource:      ResourceMwsWorkspaces(),
			InstanceState: state,
			Update:        true,
			ID:            "a",
			HCL: hcl + `
			workspace_name = "b"
			account_id = "c"`,
		}.Apply(t)
	})
}

func TestUpdateWorkspace_AddToken(t *testing.T) {
	updateWorkspaceTokenFixture(t, []qa.HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/token/create",
			ExpectedRequest: Token{
				LifetimeSeconds: 2.592e+06,
				Comment:         "Terraform PAT",
			},
			Response: tokens.TokenResponse{
				TokenValue: "sensitive",
				TokenInfo: &tokens.TokenInfo{
					TokenID: "abcdef",
				},
			},
		},
	}, map[string]string{
		// no token in state
	}, `token {}`)
}

func TestUpdateWorkspace_DeleteToken(t *testing.T) {
	updateWorkspaceTokenFixture(t, []qa.HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/token/delete",
			ExpectedRequest: map[string]any{
				"token_id": "abcdef",
			},
		},
	}, map[string]string{
		"token.#":                  "1",
		"token.0.comment":          "Terraform PAT",
		"token.0.lifetime_seconds": "2592000",
		"token.0.token_id":         "abcdef",
		"token.0.token_value":      "sensitive",
	}, ``)
}

func TestUpdateWorkspace_ReplaceToken(t *testing.T) {
	updateWorkspaceTokenFixture(t, []qa.HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/token/delete",
			ExpectedRequest: map[string]any{
				"token_id": "abcdef",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/token/create",
			ExpectedRequest: Token{
				LifetimeSeconds: 2.592e+06,
				Comment:         "I am Batman!",
			},
			Response: tokens.TokenResponse{
				TokenValue: "new-value",
				TokenInfo: &tokens.TokenInfo{
					TokenID: "new-id",
				},
			},
		},
	}, map[string]string{
		"token.#":                  "1",
		"token.0.comment":          "Terraform PAT",
		"token.0.lifetime_seconds": "2592000",
		"token.0.token_id":         "abcdef",
		"token.0.token_value":      "sensitive",
	}, `token { 
		comment = "I am Batman!"
	}`)
}

func TestEnsureTokenExists(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/token/list",
			Response: `{}`, // we just need a JSON for this
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/token/create",
			ExpectedRequest: Token{
				LifetimeSeconds: 3600,
				Comment:         "test",
			},
			Response: tokens.TokenResponse{
				TokenValue: "new-value",
				TokenInfo: &tokens.TokenInfo{
					TokenID: "new-id",
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		r := ResourceMwsWorkspaces()
		d := r.TestResourceData()
		d.Set("workspace_url", client.Host)
		d.Set("token", []any{
			map[string]any{
				"lifetime_seconds": 3600,
				"comment":          "test",
				"token_id":         "abcdef",
			},
		})
		wsApi := NewWorkspacesAPI(context.Background(), client)
		err := EnsureTokenExistsIfNeeded(wsApi, r.Schema, d)
		assert.NoError(t, err)
	})
}

func TestEnsureTokenExists_NoRecreate(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/token/list",
			Response: tokens.TokenList{
				TokenInfos: []tokens.TokenInfo{
					{
						TokenID: "old-id",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		r := ResourceMwsWorkspaces()
		d := r.TestResourceData()
		d.Set("workspace_url", client.Host)
		d.Set("token", []any{
			map[string]any{
				"lifetime_seconds": 3600,
				"comment":          "test",
				"token_id":         "old-id",
			},
		})
		wsApi := NewWorkspacesAPI(context.Background(), client)
		err := EnsureTokenExistsIfNeeded(wsApi, r.Schema, d)
		assert.NoError(t, err)
	})
}

func TestWorkspaceTokenWrongAuthCornerCase(t *testing.T) {
	defer common.CleanupEnvironment()()
	client := &common.DatabricksClient{}
	r := ResourceMwsWorkspaces()
	d := r.TestResourceData()
	d.Set("workspace_url", client.Host)
	d.Set("token", []any{
		map[string]any{
			"lifetime_seconds": 3600,
			"comment":          "test",
			"token_id":         "old-id",
		},
	})

	wsApi := NewWorkspacesAPI(context.Background(), client)

	noAuth := "cannot authenticate parent client: authentication is not configured " +
		"for provider. Please check https://registry.terraform.io/providers/" +
		"databricks/databricks/latest/docs#authentication for details"
	assert.EqualError(t, CreateTokenIfNeeded(wsApi, r.Schema, d), noAuth, "create")
	assert.EqualError(t, EnsureTokenExistsIfNeeded(wsApi, r.Schema, d), noAuth, "ensure")
	assert.EqualError(t, removeTokenIfNeeded(wsApi, r.Schema, "x", d), noAuth, "remove")

}

func TestWorkspaceTokenHttpCornerCases(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			MatchAny:     true,
			ReuseRequest: true,
			Status:       418,
			Response: common.APIError{
				ErrorCode:  "NONSENSE",
				StatusCode: 418,
				Message:    "I'm a teapot",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		wsApi := NewWorkspacesAPI(context.Background(), client)
		r := ResourceMwsWorkspaces()
		d := r.TestResourceData()
		d.Set("workspace_url", client.Host)
		d.Set("token", []any{
			map[string]any{
				"lifetime_seconds": 3600,
				"comment":          "test",
				"token_id":         "old-id",
			},
		})
		for msg, err := range map[string]error{
			"cannot create token: I'm a teapot": CreateTokenIfNeeded(wsApi, r.Schema, d),
			"cannot read token: I'm a teapot":   EnsureTokenExistsIfNeeded(wsApi, r.Schema, d),
			"cannot remove token: I'm a teapot": removeTokenIfNeeded(wsApi, r.Schema, "x", d),
		} {
			assert.EqualError(t, err, msg)
		}
	})
}

func TestGenerateWorkspaceHostname_CornerCases(t *testing.T) {
	assert.Equal(t, "fallback.cloud.databricks.com",
		generateWorkspaceHostname(&common.DatabricksClient{
			Host: "$%^&*",
		}, Workspace{
			DeploymentName: "fallback",
		}))
	assert.Equal(t, "stuff.is.exaple.com",
		generateWorkspaceHostname(&common.DatabricksClient{
			Host: "https://this.is.exaple.com",
		}, Workspace{
			DeploymentName: "stuff",
		}))
}

func TestExplainWorkspaceFailureCornerCase(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			MatchAny:     true,
			ReuseRequest: true,
			Status:       418,
			Response: common.APIError{
				ErrorCode:  "NONSENSE",
				StatusCode: 418,
				Message:    "🐜",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		wsApi := NewWorkspacesAPI(context.Background(), client)

		assert.EqualError(t, wsApi.explainWorkspaceFailure(Workspace{
			WorkspaceStatusMessage: "🔥",
		}), "🔥")

		assert.EqualError(t, wsApi.explainWorkspaceFailure(Workspace{
			NetworkID: "abc",
		}), "failed to start workspace. Cannot read network: 🐜")
	})
}

func TestResourceWorkspaceUpdatePrivateAccessSettings(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/accounts/abc/workspaces/1234",
				ExpectedRequest: map[string]any{
					"credentials_id":                  "bcd",
					"network_id":                      "fgh",
					"storage_customer_managed_key_id": "def",
					"private_access_settings_id":      "pas",
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/accounts/abc/workspaces/1234",
				Response: Workspace{
					WorkspaceStatus:                     WorkspaceStatusRunning,
					WorkspaceName:                       "labdata",
					DeploymentName:                      "900150983cd24fb0",
					AwsRegion:                           "us-east-1",
					CredentialsID:                       "bcd",
					StorageConfigurationID:              "ghi",
					NetworkID:                           "fgh",
					ManagedServicesCustomerManagedKeyID: "def",
					StorageCustomerManagedKeyID:         "def",
					PrivateAccessSettingsID:             "pas",
					AccountID:                           "abc",
					WorkspaceID:                         1234,
				},
			},
		},
		Resource: ResourceMwsWorkspaces(),
		InstanceState: map[string]string{
			"account_id":     "abc",
			"aws_region":     "us-east-1",
			"credentials_id": "__OLDER__",
			"managed_services_customer_managed_key_id": "def",
			"storage_customer_managed_key_id":          "__OLDER__",
			"deployment_name":                          "900150983cd24fb0",
			"workspace_name":                           "labdata",
			"is_no_public_ip_enabled":                  "true",
			"network_id":                               "fgh",
			"storage_configuration_id":                 "ghi",
			"workspace_id":                             "1234",
		},
		State: map[string]any{
			"account_id":     "abc",
			"aws_region":     "us-east-1",
			"credentials_id": "bcd",
			"managed_services_customer_managed_key_id": "def",
			"storage_customer_managed_key_id":          "def",
			"deployment_name":                          "900150983cd24fb0",
			"workspace_name":                           "labdata",
			"is_no_public_ip_enabled":                  true,
			"network_id":                               "fgh",
			"storage_configuration_id":                 "ghi",
			"private_access_settings_id":               "pas",
			"workspace_id":                             1234,
		},
		Update: true,
		ID:     "abc/1234",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/1234", d.Id(), "Id should be the same as in reading")
}

func TestResourceWorkspaceRemovePAS_NotAllowed(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceMwsWorkspaces(),
		InstanceState: map[string]string{
			"account_id":     "abc",
			"aws_region":     "us-east-1",
			"credentials_id": "bcd",
			"managed_services_customer_managed_key_id": "def",
			"storage_customer_managed_key_id":          "def",
			"deployment_name":                          "900150983cd24fb0",
			"workspace_name":                           "labdata",
			"is_no_public_ip_enabled":                  "true",
			"network_id":                               "fgh",
			"storage_configuration_id":                 "ghi",
			"workspace_id":                             "1234",
			"private_access_settings_id":               "pas",
		},
		State: map[string]any{
			"account_id":     "abc",
			"aws_region":     "us-east-1",
			"credentials_id": "bcd",
			"managed_services_customer_managed_key_id": "def",
			"storage_customer_managed_key_id":          "def",
			"deployment_name":                          "900150983cd24fb0",
			"workspace_name":                           "labdata",
			"is_no_public_ip_enabled":                  true,
			"network_id":                               "fgh",
			"storage_configuration_id":                 "ghi",
			"workspace_id":                             1234,
			"private_access_settings_id":               "",
		},
		Update: true,
		ID:     "abc/1234",
	}.ExpectError(t, "cannot remove private access setting from workspace")
}

func TestResourceWorkspaceCreateGcpManagedVPC(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/workspaces",
				// retreating to raw JSON, as certain fields don't work well together
				ExpectedRequest: map[string]any{
					"account_id": "abc",
					"cloud":      "gcp",
					"cloud_resource_container": map[string]any{
						"gcp": map[string]any{
							"project_id": "def",
						},
					},
					"location":       "bcd",
					"workspace_name": "labdata",
				},
				Response: Workspace{
					WorkspaceID:    1234,
					AccountID:      "abc",
					DeploymentName: "900150983cd24fb0",
					WorkspaceName:  "labdata",
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/accounts/abc/workspaces/1234",
				Response: Workspace{
					AccountID:       "abc",
					WorkspaceID:     1234,
					WorkspaceStatus: WorkspaceStatusRunning,
					DeploymentName:  "900150983cd24fb0",
					WorkspaceName:   "labdata",
					GCPManagedNetworkConfig: &GCPManagedNetworkConfig{
						SubnetCIDR:               "a",
						GKEClusterPodIPRange:     "b",
						GKEClusterServiceIPRange: "c",
					},
					GkeConfig: &GkeConfig{
						ConnectivityType: "d",
						MasterIPRange:    "e",
					},
				},
			},
		},
		Resource: ResourceMwsWorkspaces(),
		HCL: `
		account_id      = "abc"
		workspace_name  = "labdata"
		deployment_name = "900150983cd24fb0"
		location        = "bcd"
		cloud_resource_container {
			gcp {
				project_id = "def"
			}
		}
		`,
		Gcp:    true,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, []any([]any{}), d.Get("network"), "Network configuration should be ignored")
}

func TestResourceWorkspaceUpdateGcpManagedVPCNoChange(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/accounts/abc/workspaces/1234",
				Response: Workspace{
					AccountID:       "abc",
					Cloud:           "gcp",
					WorkspaceID:     1234,
					WorkspaceStatus: WorkspaceStatusRunning,
					DeploymentName:  "900150983cd24fb0",
					WorkspaceName:   "labdata",
					GCPManagedNetworkConfig: &GCPManagedNetworkConfig{
						SubnetCIDR:               "a",
						GKEClusterPodIPRange:     "b",
						GKEClusterServiceIPRange: "c",
					},
					GkeConfig: &GkeConfig{
						ConnectivityType: "d",
						MasterIPRange:    "e",
					},
				},
			},
		},
		Resource: ResourceMwsWorkspaces(),
		InstanceState: map[string]string{
			"account_id":                 "abc",
			"workspace_name":             "labdata",
			"deployment_name":            "900150983cd24fb0",
			"location":                   "bcd",
			"workspace_id":               "1234",
			"is_no_public_ip_enabled":    "false",
			"cloud_resource_container.#": "1",
			"cloud_resource_container.0.gcp.0.project_id": "def",
			"cloud": "gcp",
		},
		HCL: `
		account_id      = "abc"
		workspace_name  = "labdata"
		deployment_name = "900150983cd24fb0"
		location        = "bcd"
		cloud_resource_container {
			gcp {
				project_id = "def"
			}
		}
		`,
		Gcp:    true,
		Update: true,
		ID:     "abc/1234",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/1234", d.Id(), "Id should be the same as in reading")
	assert.Equal(t, []any([]any{}), d.Get("network"), "Network configuration should be ignored")
}
