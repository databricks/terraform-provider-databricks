package mws

import (
	"context"
	"errors"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/databricks/databricks-sdk-go/service/settings"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func mockScimMe(c *mocks.MockWorkspaceClient) {
	c.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "me@hello.com"}, nil)
}

func setConfigHost(host string) func(*mocks.MockWorkspaceClient) {
	return func(c *mocks.MockWorkspaceClient) {
		c.WorkspaceClient.Config = &config.Config{
			Host: host,
		}
	}
}

func setDefaultConfigHost(c *mocks.MockWorkspaceClient) {
	c.WorkspaceClient.Config = &config.Config{
		Host: "900150983cd24fb0.cloud.databricks.com",
	}
}

func basicMockWorkspaceClients(t *testing.T, configs ...func(*mocks.MockWorkspaceClient)) func(map[int64]*mocks.MockWorkspaceClient) {
	return func(m map[int64]*mocks.MockWorkspaceClient) {
		c := mocks.NewMockWorkspaceClient(t)
		for _, config := range configs {
			config(c)
		}
		m[1234] = c
	}
}

func TestResourceWorkspaceCreate(t *testing.T) {
	// Define a mock workspace that can be reused
	mockWorkspace := &provisioning.Workspace{
		WorkspaceId:                         1234,
		WorkspaceStatus:                     provisioning.WorkspaceStatusRunning,
		WorkspaceName:                       "labdata",
		DeploymentName:                      "900150983cd24fb0",
		AwsRegion:                           "us-east-1",
		CredentialsId:                       "bcd",
		StorageConfigurationId:              "ghi",
		NetworkId:                           "fgh",
		ManagedServicesCustomerManagedKeyId: "def",
		StorageCustomerManagedKeyId:         "def",
		AccountId:                           "abc",
		CustomTags: map[string]string{
			"SoldToCode": "1234",
		},
	}

	// Create a mock waiter
	mockWaiter := &provisioning.WaitGetWorkspaceRunning[provisioning.Workspace]{
		WorkspaceId: 1234,
		Poll: func(d time.Duration, f func(*provisioning.Workspace)) (*provisioning.Workspace, error) {
			return mockWorkspace, nil
		},
	}

	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockWorkspacesAPI().EXPECT().Create(mock.Anything, provisioning.CreateWorkspaceRequest{
				IsNoPublicIpEnabled:                 true,
				WorkspaceName:                       "labdata",
				DeploymentName:                      "900150983cd24fb0",
				AwsRegion:                           "us-east-1",
				CredentialsId:                       "bcd",
				StorageConfigurationId:              "ghi",
				NetworkId:                           "fgh",
				ManagedServicesCustomerManagedKeyId: "def",
				StorageCustomerManagedKeyId:         "def",
				CustomTags: map[string]string{
					"SoldToCode": "1234",
				},
			}).Return(mockWaiter, nil)
			a.GetMockWorkspacesAPI().EXPECT().Get(mock.Anything, provisioning.GetWorkspaceRequest{
				WorkspaceId: 1234,
			}).Return(mockWorkspace, nil)
			a.GetMockWorkspacesAPI().EXPECT().WaitGetWorkspaceRunning(mock.Anything, int64(1234), 20*time.Minute, mock.Anything).Return(mockWorkspace, nil)
		},
		MockWorkspaceClientsFunc: basicMockWorkspaceClients(t, setDefaultConfigHost, mockScimMe),
		Resource:                 ResourceMwsWorkspaces(),
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
			"custom_tags": map[string]any{
				"SoldToCode": "1234",
			},
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/1234", d.Id())
}

func TestResourceWorkspaceCreateGcp(t *testing.T) {
	// Define a mock workspace that can be reused
	mockWorkspace := &provisioning.Workspace{
		WorkspaceId:     1234,
		WorkspaceStatus: provisioning.WorkspaceStatusRunning,
		WorkspaceName:   "labdata",
		DeploymentName:  "900150983cd24fb0",
		AccountId:       "abc",
		Cloud:           "gcp",
		Location:        "bcd",
		GcpManagedNetworkConfig: &provisioning.GcpManagedNetworkConfig{
			SubnetCidr: "a",
		},
	}

	// Create a mock waiter
	mockWaiter := &provisioning.WaitGetWorkspaceRunning[provisioning.Workspace]{
		WorkspaceId: 1234,
		Poll: func(d time.Duration, f func(*provisioning.Workspace)) (*provisioning.Workspace, error) {
			return mockWorkspace, nil
		},
	}

	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockWorkspacesAPI().EXPECT().Create(mock.Anything, provisioning.CreateWorkspaceRequest{
				CloudResourceContainer: &provisioning.CloudResourceContainer{
					Gcp: &provisioning.CustomerFacingGcpCloudResourceContainer{
						ProjectId: "def",
					},
				},
				DeploymentName:      "900150983cd24fb0",
				IsNoPublicIpEnabled: true,
				Location:            "bcd",
				NetworkId:           "net_id_a",
				GcpManagedNetworkConfig: &provisioning.GcpManagedNetworkConfig{
					SubnetCidr: "a",
				},
				WorkspaceName: "labdata",
			}).Return(mockWaiter, nil)
			a.GetMockWorkspacesAPI().EXPECT().Get(mock.Anything, provisioning.GetWorkspaceRequest{
				WorkspaceId: 1234,
			}).Return(mockWorkspace, nil)
			a.GetMockWorkspacesAPI().EXPECT().WaitGetWorkspaceRunning(mock.Anything, int64(1234), 20*time.Minute, mock.Anything).Return(mockWorkspace, nil)
		},
		MockWorkspaceClientsFunc: basicMockWorkspaceClients(t, setDefaultConfigHost, mockScimMe),
		Resource:                 ResourceMwsWorkspaces(),
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
		}
		`,
		Gcp:    true,
		Create: true,
	}.ApplyNoError(t)
}

func TestResourceWorkspaceCreate_Error_Custom_tags(t *testing.T) {
	qa.ResourceFixture{
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
		private_access_settings_id = "pas_id_a"
		network_id = "net_id_a"
		gcp_managed_network_config {
			subnet_cidr = "a"
		}
		custom_tags = {
			SoldToCode = "1234"
		}
		`,
		Gcp:    true,
		Create: true,
	}.ExpectError(t, "custom_tags are only allowed for AWS workspaces")
}

func TestResourceWorkspaceCreateGcpPsc(t *testing.T) {
	// Define a mock workspace that can be reused
	mockWorkspace := &provisioning.Workspace{
		WorkspaceId:             1234,
		WorkspaceStatus:         provisioning.WorkspaceStatusRunning,
		WorkspaceName:           "labdata",
		DeploymentName:          "900150983cd24fb0",
		AccountId:               "abc",
		Cloud:                   "gcp",
		Location:                "bcd",
		PrivateAccessSettingsId: "pas_id_a",
		GcpManagedNetworkConfig: &provisioning.GcpManagedNetworkConfig{
			SubnetCidr: "a",
		},
	}

	// Create a mock waiter
	mockWaiter := &provisioning.WaitGetWorkspaceRunning[provisioning.Workspace]{
		WorkspaceId: 1234,
		Poll: func(d time.Duration, f func(*provisioning.Workspace)) (*provisioning.Workspace, error) {
			return mockWorkspace, nil
		},
	}

	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockWorkspacesAPI().EXPECT().Create(mock.Anything, provisioning.CreateWorkspaceRequest{
				CloudResourceContainer: &provisioning.CloudResourceContainer{
					Gcp: &provisioning.CustomerFacingGcpCloudResourceContainer{
						ProjectId: "def",
					},
				},
				DeploymentName:          "900150983cd24fb0",
				IsNoPublicIpEnabled:     true,
				Location:                "bcd",
				PrivateAccessSettingsId: "pas_id_a",
				NetworkId:               "net_id_a",
				GcpManagedNetworkConfig: &provisioning.GcpManagedNetworkConfig{
					SubnetCidr: "a",
				},
				WorkspaceName: "labdata",
			}).Return(mockWaiter, nil)
			a.GetMockWorkspacesAPI().EXPECT().Get(mock.Anything, provisioning.GetWorkspaceRequest{
				WorkspaceId: 1234,
			}).Return(mockWorkspace, nil)
			a.GetMockWorkspacesAPI().EXPECT().WaitGetWorkspaceRunning(mock.Anything, int64(1234), 20*time.Minute, mock.Anything).Return(mockWorkspace, nil)

		},
		MockWorkspaceClientsFunc: basicMockWorkspaceClients(t, setDefaultConfigHost, mockScimMe),
		Resource:                 ResourceMwsWorkspaces(),
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
		private_access_settings_id = "pas_id_a"
		network_id = "net_id_a"
		gcp_managed_network_config {
			subnet_cidr = "a"
		}
		`,
		Gcp:    true,
		Create: true,
	}.ApplyNoError(t)
}

func TestResourceWorkspaceCreateGcpCmk(t *testing.T) {
	// Define a mock workspace that can be reused
	mockWorkspace := &provisioning.Workspace{
		WorkspaceId:             1234,
		WorkspaceStatus:         provisioning.WorkspaceStatusRunning,
		WorkspaceName:           "labdata",
		AccountId:               "abc",
		DeploymentName:          "900150983cd24fb0",
		Cloud:                   "gcp",
		Location:                "bcd",
		PrivateAccessSettingsId: "pas_id_a",
		GcpManagedNetworkConfig: &provisioning.GcpManagedNetworkConfig{
			SubnetCidr: "a",
		},
		ManagedServicesCustomerManagedKeyId: "managed_services_cmk",
		StorageCustomerManagedKeyId:         "storage_cmk",
	}

	// Create a mock waiter
	mockWaiter := &provisioning.WaitGetWorkspaceRunning[provisioning.Workspace]{
		WorkspaceId: 1234,
		Poll: func(d time.Duration, f func(*provisioning.Workspace)) (*provisioning.Workspace, error) {
			return mockWorkspace, nil
		},
	}

	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockWorkspacesAPI().EXPECT().Create(mock.Anything, provisioning.CreateWorkspaceRequest{
				CloudResourceContainer: &provisioning.CloudResourceContainer{
					Gcp: &provisioning.CustomerFacingGcpCloudResourceContainer{
						ProjectId: "def",
					},
				},
				DeploymentName:          "900150983cd24fb0",
				IsNoPublicIpEnabled:     true,
				Location:                "bcd",
				PrivateAccessSettingsId: "pas_id_a",
				NetworkId:               "net_id_a",
				GcpManagedNetworkConfig: &provisioning.GcpManagedNetworkConfig{
					SubnetCidr: "a",
				},
				WorkspaceName:                       "labdata",
				ManagedServicesCustomerManagedKeyId: "managed_services_cmk",
				StorageCustomerManagedKeyId:         "storage_cmk",
			}).Return(mockWaiter, nil)
			a.GetMockWorkspacesAPI().EXPECT().Get(mock.Anything, provisioning.GetWorkspaceRequest{
				WorkspaceId: 1234,
			}).Return(mockWorkspace, nil)
			a.GetMockWorkspacesAPI().EXPECT().WaitGetWorkspaceRunning(mock.Anything, int64(1234), 20*time.Minute, mock.Anything).Return(mockWorkspace, nil)
		},
		MockWorkspaceClientsFunc: basicMockWorkspaceClients(t, setDefaultConfigHost, mockScimMe),
		Resource:                 ResourceMwsWorkspaces(),
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
		private_access_settings_id = "pas_id_a"
		network_id = "net_id_a"
		gcp_managed_network_config {
			subnet_cidr = "a"
		}
		managed_services_customer_managed_key_id = "managed_services_cmk"
		storage_customer_managed_key_id = "storage_cmk"
		`,
		Gcp:    true,
		Create: true,
	}.ApplyNoError(t)
}

func TestResourceWorkspaceCreateWithIsNoPublicIPEnabledFalse(t *testing.T) {
	// Define a mock workspace that can be reused
	mockWorkspace := &provisioning.Workspace{
		WorkspaceId:                         1234,
		WorkspaceStatus:                     provisioning.WorkspaceStatusRunning,
		WorkspaceName:                       "labdata",
		DeploymentName:                      "900150983cd24fb0",
		AwsRegion:                           "us-east-1",
		CredentialsId:                       "bcd",
		StorageConfigurationId:              "ghi",
		NetworkId:                           "fgh",
		ManagedServicesCustomerManagedKeyId: "def",
		StorageCustomerManagedKeyId:         "def",
		AccountId:                           "abc",
	}

	// Create a mock waiter
	mockWaiter := &provisioning.WaitGetWorkspaceRunning[provisioning.Workspace]{
		WorkspaceId: 1234,
		Poll: func(d time.Duration, f func(*provisioning.Workspace)) (*provisioning.Workspace, error) {
			return mockWorkspace, nil
		},
	}

	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockWorkspacesAPI().EXPECT().Create(mock.Anything, provisioning.CreateWorkspaceRequest{
				IsNoPublicIpEnabled:                 false,
				WorkspaceName:                       "labdata",
				DeploymentName:                      "900150983cd24fb0",
				AwsRegion:                           "us-east-1",
				CredentialsId:                       "bcd",
				StorageConfigurationId:              "ghi",
				NetworkId:                           "fgh",
				ManagedServicesCustomerManagedKeyId: "def",
				StorageCustomerManagedKeyId:         "def",
			}).Return(mockWaiter, nil)
			a.GetMockWorkspacesAPI().EXPECT().Get(mock.Anything, provisioning.GetWorkspaceRequest{
				WorkspaceId: 1234,
			}).Return(mockWorkspace, nil)
			a.GetMockWorkspacesAPI().EXPECT().WaitGetWorkspaceRunning(mock.Anything, int64(1234), 20*time.Minute, mock.Anything).Return(mockWorkspace, nil)
		},
		MockWorkspaceClientsFunc: basicMockWorkspaceClients(t, setDefaultConfigHost, mockScimMe),
		Resource:                 ResourceMwsWorkspaces(),
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
	assert.NoError(t, err)
	assert.Equal(t, "abc/1234", d.Id())
}

func TestResourceWorkspaceCreateLegacyConfig(t *testing.T) {
	// Define a mock workspace that can be reused
	mockWorkspace := &provisioning.Workspace{
		WorkspaceId:                         1234,
		WorkspaceStatus:                     provisioning.WorkspaceStatusRunning,
		WorkspaceName:                       "labdata",
		DeploymentName:                      "900150983cd24fb0",
		AwsRegion:                           "us-east-1",
		CredentialsId:                       "bcd",
		StorageConfigurationId:              "ghi",
		NetworkId:                           "fgh",
		ManagedServicesCustomerManagedKeyId: "def",
		AccountId:                           "abc",
	}

	// Create a mock waiter
	mockWaiter := &provisioning.WaitGetWorkspaceRunning[provisioning.Workspace]{
		WorkspaceId: 1234,
		Poll: func(d time.Duration, f func(*provisioning.Workspace)) (*provisioning.Workspace, error) {
			return mockWorkspace, nil
		},
	}

	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockWorkspacesAPI().EXPECT().Create(mock.Anything, provisioning.CreateWorkspaceRequest{
				IsNoPublicIpEnabled:                 true,
				WorkspaceName:                       "labdata",
				DeploymentName:                      "900150983cd24fb0",
				AwsRegion:                           "us-east-1",
				CredentialsId:                       "bcd",
				StorageConfigurationId:              "ghi",
				NetworkId:                           "fgh",
				ManagedServicesCustomerManagedKeyId: "def",
			}).Return(mockWaiter, nil)
			a.GetMockWorkspacesAPI().EXPECT().Get(mock.Anything, provisioning.GetWorkspaceRequest{
				WorkspaceId: 1234,
			}).Return(mockWorkspace, nil)
			a.GetMockWorkspacesAPI().EXPECT().WaitGetWorkspaceRunning(mock.Anything, int64(1234), 20*time.Minute, mock.Anything).Return(mockWorkspace, nil)
		},
		MockWorkspaceClientsFunc: basicMockWorkspaceClients(t, setDefaultConfigHost, mockScimMe),
		Resource:                 ResourceMwsWorkspaces(),
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
	assert.NoError(t, err)
	assert.Equal(t, "abc/1234", d.Id())
}

func TestResourceWorkspaceRead(t *testing.T) {
	// Define a mock workspace that can be reused
	mockWorkspace := &provisioning.Workspace{
		WorkspaceId:                         1234,
		WorkspaceStatus:                     provisioning.WorkspaceStatusRunning,
		WorkspaceName:                       "labdata",
		DeploymentName:                      "900150983cd24fb0",
		AwsRegion:                           "us-east-1",
		CredentialsId:                       "bcd",
		StorageConfigurationId:              "ghi",
		NetworkId:                           "fgh",
		ManagedServicesCustomerManagedKeyId: "def",
		StorageCustomerManagedKeyId:         "def",
		AccountId:                           "abc",
	}

	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockWorkspacesAPI().EXPECT().Get(mock.Anything, provisioning.GetWorkspaceRequest{
				WorkspaceId: 1234,
			}).Return(mockWorkspace, nil)
			a.GetMockWorkspacesAPI().EXPECT().WaitGetWorkspaceRunning(mock.Anything, int64(1234), 20*time.Minute, mock.Anything).Return(mockWorkspace, nil)
		},
		MockWorkspaceClientsFunc: basicMockWorkspaceClients(t, setDefaultConfigHost),
		Resource:                 ResourceMwsWorkspaces(),
		Read:                     true,
		New:                      true,
		ID:                       "abc/1234",
	}.Apply(t)
	assert.NoError(t, err)
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
	// Define a mock workspace that can be reused
	mockWorkspace := &provisioning.Workspace{
		WorkspaceId:                         1234,
		WorkspaceStatus:                     provisioning.WorkspaceStatusRunning,
		WorkspaceName:                       "labdata",
		DeploymentName:                      "prefix-900150983cd24fb0",
		AwsRegion:                           "us-east-1",
		CredentialsId:                       "bcd",
		StorageConfigurationId:              "ghi",
		NetworkId:                           "fgh",
		ManagedServicesCustomerManagedKeyId: "def",
		StorageCustomerManagedKeyId:         "def",
		AccountId:                           "abc",
	}

	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockWorkspacesAPI().EXPECT().Get(mock.Anything, provisioning.GetWorkspaceRequest{
				WorkspaceId: 1234,
			}).Return(mockWorkspace, nil)
			a.GetMockWorkspacesAPI().EXPECT().WaitGetWorkspaceRunning(mock.Anything, int64(1234), 20*time.Minute, mock.Anything).Return(mockWorkspace, nil)
		},
		MockWorkspaceClientsFunc: basicMockWorkspaceClients(t, setConfigHost("prefix-900150983cd24fb0.cloud.databricks.com")),
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
	assert.NoError(t, err)
	assert.Equal(t, "abc/1234", d.Id(), "Id should not be empty")
	assert.Equal(t, "prefix-900150983cd24fb0", d.Get("deployment_name"))
	assert.Equal(t, "https://prefix-900150983cd24fb0.cloud.databricks.com", d.Get("workspace_url"))
}

func TestResourceWorkspaceRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockWorkspacesAPI().EXPECT().Get(mock.Anything, provisioning.GetWorkspaceRequest{
				WorkspaceId: 1234,
			}).Return(nil, &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				Message:    "Item not found",
				StatusCode: 404,
			})
		},
		Resource: ResourceMwsWorkspaces(),
		Read:     true,
		Removed:  true,
		ID:       "abc/1234",
	}.ApplyNoError(t)
}

func TestResourceWorkspaceRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockWorkspacesAPI().EXPECT().Get(mock.Anything, provisioning.GetWorkspaceRequest{
				WorkspaceId: 1234,
			}).Return(nil, &apierr.APIError{
				ErrorCode:  "INVALID_REQUEST",
				Message:    "Internal error happened",
				StatusCode: 400,
			})
		},
		Resource: ResourceMwsWorkspaces(),
		Read:     true,
		ID:       "abc/1234",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/1234", d.Id(), "Id should not be empty for error reads")
}

func TestResourceWorkspaceUpdate(t *testing.T) {
	// Define a mock workspace that can be reused
	mockWorkspace := &provisioning.Workspace{
		WorkspaceId:                         1234,
		WorkspaceStatus:                     provisioning.WorkspaceStatusRunning,
		WorkspaceName:                       "labdata",
		DeploymentName:                      "900150983cd24fb0",
		AwsRegion:                           "us-east-1",
		CredentialsId:                       "bcd",
		StorageConfigurationId:              "ghi",
		NetworkId:                           "fgh",
		ManagedServicesCustomerManagedKeyId: "def",
		StorageCustomerManagedKeyId:         "def",
		AccountId:                           "abc",
	}

	// Create a mock waiter
	mockWaiter := &provisioning.WaitGetWorkspaceRunning[struct{}]{
		WorkspaceId: 1234,
		Poll: func(d time.Duration, f func(*provisioning.Workspace)) (*provisioning.Workspace, error) {
			return mockWorkspace, nil
		},
	}

	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockWorkspacesAPI().EXPECT().Update(mock.Anything, provisioning.UpdateWorkspaceRequest{
				WorkspaceId:                         1234,
				AwsRegion:                           "us-east-1",
				ManagedServicesCustomerManagedKeyId: "def",
				CredentialsId:                       "bcd",
				NetworkId:                           "fgh",
				StorageCustomerManagedKeyId:         "def",
				StorageConfigurationId:              "ghi",
			}).Return(mockWaiter, nil)
			a.GetMockWorkspacesAPI().EXPECT().Get(mock.Anything, provisioning.GetWorkspaceRequest{
				WorkspaceId: 1234,
			}).Return(mockWorkspace, nil)
			a.GetMockWorkspacesAPI().EXPECT().WaitGetWorkspaceRunning(mock.Anything, int64(1234), 20*time.Minute, mock.Anything).Return(mockWorkspace, nil)
		},
		MockWorkspaceClientsFunc: basicMockWorkspaceClients(t, setDefaultConfigHost),
		Resource:                 ResourceMwsWorkspaces(),
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
	assert.NoError(t, err)
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
	// Define a mock workspace that can be reused
	mockWorkspace := &provisioning.Workspace{
		WorkspaceId:                         1234,
		WorkspaceStatus:                     provisioning.WorkspaceStatusRunning,
		IsNoPublicIpEnabled:                 true,
		WorkspaceName:                       "labdata",
		DeploymentName:                      "900150983cd24fb0",
		AwsRegion:                           "us-east-1",
		CredentialsId:                       "bcd",
		StorageConfigurationId:              "ghi",
		NetworkId:                           "fgh",
		ManagedServicesCustomerManagedKeyId: "def",
		AccountId:                           "abc",
	}

	// Create a mock waiter
	mockWaiter := &provisioning.WaitGetWorkspaceRunning[struct{}]{
		WorkspaceId: 1234,
		Poll: func(d time.Duration, f func(*provisioning.Workspace)) (*provisioning.Workspace, error) {
			return mockWorkspace, nil
		},
	}

	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockWorkspacesAPI().EXPECT().Update(mock.Anything, provisioning.UpdateWorkspaceRequest{
				WorkspaceId:                         1234,
				AwsRegion:                           "us-east-1",
				ManagedServicesCustomerManagedKeyId: "def",
				StorageConfigurationId:              "ghi",
				CredentialsId:                       "bcd",
				NetworkId:                           "fgh",
			}).Return(mockWaiter, nil)
			a.GetMockWorkspacesAPI().EXPECT().Get(mock.Anything, provisioning.GetWorkspaceRequest{
				WorkspaceId: 1234,
			}).Return(mockWorkspace, nil)
			a.GetMockWorkspacesAPI().EXPECT().WaitGetWorkspaceRunning(mock.Anything, int64(1234), 20*time.Minute, mock.Anything).Return(mockWorkspace, nil)
		},
		MockWorkspaceClientsFunc: basicMockWorkspaceClients(t, setDefaultConfigHost),
		Resource:                 ResourceMwsWorkspaces(),
		InstanceState: map[string]string{
			"account_id":               "abc",
			"aws_region":               "us-east-1",
			"credentials_id":           "old",
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
	assert.NoError(t, err)
	assert.Equal(t, "abc/1234", d.Id(), "Id should be the same as in reading")
}

func TestResourceWorkspaceUpdate_Error(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockWorkspacesAPI().EXPECT().Update(mock.Anything, provisioning.UpdateWorkspaceRequest{
				WorkspaceId:                         1234,
				AwsRegion:                           "us-east-1",
				ManagedServicesCustomerManagedKeyId: "def",
				StorageConfigurationId:              "ghi",
				CredentialsId:                       "bcd",
				NetworkId:                           "fgh",
				StorageCustomerManagedKeyId:         "def",
			}).Return(nil, &apierr.APIError{
				ErrorCode:  "INVALID_REQUEST",
				Message:    "Internal error happened",
				StatusCode: 400,
			})
		},
		Resource: ResourceMwsWorkspaces(),
		InstanceState: map[string]string{
			"account_id":     "abc",
			"aws_region":     "us-east-1",
			"credentials_id": "old",
			"managed_services_customer_managed_key_id": "def",
			"storage_customer_managed_key_id":          "def",
			"is_no_public_ip_enabled":                  "true",
			"deployment_name":                          "900150983cd24fb0",
			"workspace_name":                           "labdata",
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
			"network_id":                               "fgh",
			"storage_configuration_id":                 "ghi",
			"workspace_id":                             1234,
		},
		Update: true,
		ID:     "abc/1234",
	}.ExpectError(t, "Internal error happened")
}

func TestResourceWorkspaceDelete(t *testing.T) {
	// Define a mock workspace that can be reused for the first GET call
	mockWorkspace := &provisioning.Workspace{
		WorkspaceName:          "labdata",
		WorkspaceStatus:        provisioning.WorkspaceStatusCancelling,
		WorkspaceStatusMessage: "Things are being removed",
	}

	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockWorkspacesAPI().EXPECT().Delete(mock.Anything, provisioning.DeleteWorkspaceRequest{
				WorkspaceId: 1234,
			}).Return(nil)
			a.GetMockWorkspacesAPI().EXPECT().Get(mock.Anything, provisioning.GetWorkspaceRequest{
				WorkspaceId: 1234,
			}).Return(mockWorkspace, nil).Once()
			a.GetMockWorkspacesAPI().EXPECT().Get(mock.Anything, provisioning.GetWorkspaceRequest{
				WorkspaceId: 1234,
			}).Return(nil, &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				Message:    "Cannot find anything",
				StatusCode: 404,
			})
		},
		Resource: ResourceMwsWorkspaces(),
		Delete:   true,
		ID:       "abc/1234",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/1234", d.Id())
}

func TestCreateFailsAndCleansUp(t *testing.T) {
	// Define a mock workspace that represents the failed state
	mockFailedWorkspace := &provisioning.Workspace{
		WorkspaceId:            1234,
		WorkspaceStatus:        provisioning.WorkspaceStatusFailed,
		WorkspaceStatusMessage: "Always fails",
		WorkspaceName:          "labdata",
		DeploymentName:         "900150983cd24fb0",
		AwsRegion:              "us-east-1",
		NetworkId:              "fgh",
		AccountId:              "abc",
	}

	// Create a mock waiter
	mockWaiter := &provisioning.WaitGetWorkspaceRunning[provisioning.Workspace]{
		WorkspaceId: 1234,
		Response:    mockFailedWorkspace,
		Poll: func(d time.Duration, f func(*provisioning.Workspace)) (*provisioning.Workspace, error) {
			return nil, errors.New("failed to reach RUNNING, got FAILED")
		},
	}

	// Define a mock network with error messages
	mockNetwork := &provisioning.Network{
		NetworkId: "fgh",
		ErrorMessages: []provisioning.NetworkHealth{
			{
				ErrorType:    provisioning.ErrorTypeCredentials,
				ErrorMessage: "Message",
			},
		},
	}

	_, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			// Expect the Create call
			a.GetMockWorkspacesAPI().EXPECT().Create(mock.Anything, provisioning.CreateWorkspaceRequest{
				IsNoPublicIpEnabled:                 true,
				WorkspaceName:                       "labdata",
				DeploymentName:                      "900150983cd24fb0",
				AwsRegion:                           "us-east-1",
				CredentialsId:                       "bcd",
				StorageConfigurationId:              "ghi",
				NetworkId:                           "fgh",
				ManagedServicesCustomerManagedKeyId: "def",
				StorageCustomerManagedKeyId:         "def",
			}).Return(mockWaiter, nil)

			// Expect the Get call to retrieve the failed workspace
			a.GetMockWorkspacesAPI().EXPECT().Get(mock.Anything, provisioning.GetWorkspaceRequest{
				WorkspaceId: 1234,
			}).Return(mockFailedWorkspace, nil)

			// Expect the Get call to retrieve the network with errors
			a.GetMockNetworksAPI().EXPECT().Get(mock.Anything, provisioning.GetNetworkRequest{
				NetworkId: "fgh",
			}).Return(mockNetwork, nil)

			// Expect the Delete call to clean up the failed workspace
			a.GetMockWorkspacesAPI().EXPECT().Delete(mock.Anything, provisioning.DeleteWorkspaceRequest{
				WorkspaceId: 1234,
			}).Return(nil)

			// Expect the final Get call to confirm the workspace is gone
			a.GetMockWorkspacesAPI().EXPECT().Get(mock.Anything, provisioning.GetWorkspaceRequest{
				WorkspaceId: 1234,
			}).Return(nil, &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				Message:    "Item not found",
				StatusCode: 404,
			})
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
	assert.ErrorContains(t, err, "workspace status message: Always fails, network error message: error: credentials;error_msg: Message;")
}

func TestWorkspace_verifyWorkspaceReachable(t *testing.T) {
	// Create a mock client
	mockClient := mocks.NewMockWorkspaceClient(t)

	// Set up expectations for the first call (DNS error)
	mockClient.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(nil, &net.OpError{
		Op:  "dial",
		Net: "tcp",
		Err: &net.DNSError{
			Name: "900150983cd24fb0.cloud.databricks.com",
			Err:  "no such host",
		},
	}).Once()

	// Set up expectations for the second call (success)
	mockClient.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{
		Id: "12345",
	}, nil).Once()

	// Create a context
	ctx := context.Background()

	// Call the function with the mock client
	err := verifyWorkspaceReachable(ctx, mockClient.WorkspaceClient)

	// The function should retry and eventually succeed
	assert.NoError(t, err)
}

func TestEnsureTokenExists(t *testing.T) {
	// Create a mock workspace client
	mockClient := mocks.NewMockWorkspaceClient(t)
	mockTokensAPI := mockClient.GetMockTokensAPI()

	// Set up expectations for token list and create
	mockTokensAPI.EXPECT().
		List(mock.Anything).
		Return(&listing.SliceIterator[settings.PublicTokenInfo]{}).
		Times(1)

	mockTokensAPI.EXPECT().
		Create(mock.Anything, settings.CreateTokenRequest{
			LifetimeSeconds: 3600,
			Comment:         "test",
		}).
		Return(&settings.CreateTokenResponse{
			TokenValue: "new-value",
			TokenInfo: &settings.PublicTokenInfo{
				TokenId: "new-id",
			},
		}, nil).
		Times(1)

	// Test the function
	token := &Token{
		LifetimeSeconds: 3600,
		Comment:         "test",
	}
	err := ensureTokenExists(context.Background(), mockClient.WorkspaceClient, token)
	assert.NoError(t, err)
	assert.Equal(t, token.TokenID, "new-id")
	assert.Equal(t, token.TokenValue, SensitiveString("new-value"))
}

func TestEnsureTokenExists_NoRecreate(t *testing.T) {
	// Create a mock workspace client
	mockClient := mocks.NewMockWorkspaceClient(t)
	mockTokensAPI := mockClient.GetMockTokensAPI()

	// Set up expectations for token list
	mockTokensAPI.EXPECT().
		List(mock.Anything).
		Return(&listing.SliceIterator[settings.PublicTokenInfo]{
			{
				TokenId: "old-id",
			},
		}).
		Times(1)

	// Test the function
	token := &Token{
		LifetimeSeconds: 3600,
		Comment:         "test",
		TokenID:         "old-id",
	}
	err := ensureTokenExists(context.Background(), mockClient.WorkspaceClient, token)
	assert.NoError(t, err)
}

func TestExplainWorkspaceFailureCornerCase(t *testing.T) {
	t.Run("no network ID", func(t *testing.T) {
		assert.EqualError(t, explainWorkspaceFailure(context.Background(), nil, &provisioning.Workspace{
			WorkspaceStatusMessage: "🔥",
		}), "workspace status message: 🔥")
	})

	t.Run("network error", func(t *testing.T) {
		mockClient := mocks.NewMockAccountClient(t)
		mockNetworksClient := mockClient.GetMockNetworksAPI()

		mockNetworksClient.EXPECT().
			Get(context.Background(), provisioning.GetNetworkRequest{NetworkId: "abc"}).
			Return(nil, errors.New("🐜")).
			Times(1)

		ws := &provisioning.Workspace{
			NetworkId:              "abc",
			WorkspaceStatusMessage: "🔥",
		}
		assert.EqualError(t, explainWorkspaceFailure(context.Background(), mockClient.AccountClient, ws), "workspace status message: 🔥; network error message: cannot read network: 🐜")
	})
}

func TestResourceWorkspaceUpdatePrivateAccessSettings(t *testing.T) {
	// Define a mock workspace that can be reused
	mockWorkspace := &provisioning.Workspace{
		WorkspaceId:                         1234,
		WorkspaceStatus:                     provisioning.WorkspaceStatusRunning,
		WorkspaceName:                       "labdata",
		DeploymentName:                      "900150983cd24fb0",
		AwsRegion:                           "us-east-1",
		CredentialsId:                       "bcd",
		StorageConfigurationId:              "ghi",
		NetworkId:                           "fgh",
		ManagedServicesCustomerManagedKeyId: "def",
		StorageCustomerManagedKeyId:         "def",
		PrivateAccessSettingsId:             "pas",
		AccountId:                           "abc",
	}

	// Create a mock waiter
	mockWaiter := &provisioning.WaitGetWorkspaceRunning[struct{}]{
		WorkspaceId: 1234,
		Poll: func(d time.Duration, f func(*provisioning.Workspace)) (*provisioning.Workspace, error) {
			return mockWorkspace, nil
		},
	}

	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			// Expect the Update call
			a.GetMockWorkspacesAPI().EXPECT().Update(mock.Anything, provisioning.UpdateWorkspaceRequest{
				WorkspaceId:                         1234,
				AwsRegion:                           "us-east-1",
				ManagedServicesCustomerManagedKeyId: "def",
				StorageConfigurationId:              "ghi",
				CredentialsId:                       "bcd",
				NetworkId:                           "fgh",
				StorageCustomerManagedKeyId:         "def",
				PrivateAccessSettingsId:             "pas",
			}).Return(mockWaiter, nil)

			// Expect the Get call to retrieve the updated workspace
			a.GetMockWorkspacesAPI().EXPECT().Get(mock.Anything, provisioning.GetWorkspaceRequest{
				WorkspaceId: 1234,
			}).Return(mockWorkspace, nil)
			a.GetMockWorkspacesAPI().EXPECT().WaitGetWorkspaceRunning(mock.Anything, int64(1234), 20*time.Minute, mock.Anything).Return(mockWorkspace, nil)
		},
		MockWorkspaceClientsFunc: basicMockWorkspaceClients(t, setDefaultConfigHost),
		Resource:                 ResourceMwsWorkspaces(),
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
	assert.NoError(t, err)
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
	// Define a mock workspace that can be reused
	mockWorkspace := &provisioning.Workspace{
		WorkspaceId:     1234,
		WorkspaceStatus: provisioning.WorkspaceStatusRunning,
		WorkspaceName:   "labdata",
		DeploymentName:  "900150983cd24fb0",
		AccountId:       "abc",
		Cloud:           "gcp",
		Location:        "bcd",
		GcpManagedNetworkConfig: &provisioning.GcpManagedNetworkConfig{
			SubnetCidr: "a",
		},
	}

	// Create a mock waiter
	mockWaiter := &provisioning.WaitGetWorkspaceRunning[provisioning.Workspace]{
		WorkspaceId: 1234,
		Poll: func(d time.Duration, f func(*provisioning.Workspace)) (*provisioning.Workspace, error) {
			return mockWorkspace, nil
		},
	}

	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			// Expect the Create call
			a.GetMockWorkspacesAPI().EXPECT().Create(mock.Anything, provisioning.CreateWorkspaceRequest{
				CloudResourceContainer: &provisioning.CloudResourceContainer{
					Gcp: &provisioning.CustomerFacingGcpCloudResourceContainer{
						ProjectId: "def",
					},
				},
				IsNoPublicIpEnabled: true,
				DeploymentName:      "900150983cd24fb0",
				Location:            "bcd",
				WorkspaceName:       "labdata",
			}).Return(mockWaiter, nil)

			// Expect the Get call to retrieve the workspace
			a.GetMockWorkspacesAPI().EXPECT().Get(mock.Anything, provisioning.GetWorkspaceRequest{
				WorkspaceId: 1234,
			}).Return(mockWorkspace, nil)
			a.GetMockWorkspacesAPI().EXPECT().WaitGetWorkspaceRunning(mock.Anything, int64(1234), 20*time.Minute, mock.Anything).Return(mockWorkspace, nil)
		},
		MockWorkspaceClientsFunc: basicMockWorkspaceClients(t, setDefaultConfigHost, mockScimMe),
		Resource:                 ResourceMwsWorkspaces(),
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
	}.ApplyNoError(t)
}

func TestSensitiveDataInLogs(t *testing.T) {
	tk := Token{
		Comment:         "comment",
		LifetimeSeconds: 123,
		TokenID:         "tokenID",
		TokenValue:      "sensitive",
	}
	assert.Contains(t, fmt.Sprintf("%v", tk), "comment")
	assert.Contains(t, fmt.Sprintf("%#v", tk), "comment")
	assert.Contains(t, fmt.Sprintf("%+v", tk), "comment")
	assert.NotContains(t, fmt.Sprintf("%v", tk), "sensitive")
	assert.NotContains(t, fmt.Sprintf("%#v", tk), "sensitive")
	assert.NotContains(t, fmt.Sprintf("%+v", tk), "sensitive")
}
