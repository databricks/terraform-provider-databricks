package exporter

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/mws"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Helper functions to set up account configs

func setupAwsAccountConfig(ma *mocks.MockAccountClient) {
	ma.AccountClient.Config = &config.Config{
		AccountID: testAccountID,
		Host:      "https://accounts.cloud.databricks.com",
	}
}

func setupAzureAccountConfig(ma *mocks.MockAccountClient) {
	ma.AccountClient.Config = &config.Config{
		AccountID: testAccountID,
		Host:      "https://accounts.azuredatabricks.net",
	}
}

// Tests for MWS Credentials

func TestListMwsCredentials(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		ma.GetMockCredentialsAPI().EXPECT().List(mock.Anything).Return([]provisioning.Credential{
			{
				CredentialsId:   "cred-1",
				CredentialsName: "Test Credential 1",
				CreationTime:    1700000000000,
			},
			{
				CredentialsId:   "cred-2",
				CredentialsName: "Test Credential 2",
				CreationTime:    1700000000000,
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "mws")

		err := resourcesMap["databricks_mws_credentials"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_mws_credentials[Test Credential 1] (id: "+testAccountID+"/cred-1)"])
		assert.True(t, ic.testEmits["databricks_mws_credentials[Test Credential 2] (id: "+testAccountID+"/cred-2)"])
	})
}

func TestListMwsCredentialsWithMatch(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		ma.GetMockCredentialsAPI().EXPECT().List(mock.Anything).Return([]provisioning.Credential{
			{
				CredentialsId:   "cred-1",
				CredentialsName: "Test Credential",
				CreationTime:    1700000000000,
			},
			{
				CredentialsId:   "cred-2",
				CredentialsName: "Production Credential",
				CreationTime:    1700000000000,
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "mws")
		ic.match = "Test"

		err := resourcesMap["databricks_mws_credentials"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_mws_credentials[Test Credential] (id: "+testAccountID+"/cred-1)"])
		assert.False(t, ic.testEmits["databricks_mws_credentials[Production Credential] (id: "+testAccountID+"/cred-2)"])
	})
}

func TestListMwsCredentialsIncremental(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		ma.GetMockCredentialsAPI().EXPECT().List(mock.Anything).Return([]provisioning.Credential{
			{
				CredentialsId:   "cred-1",
				CredentialsName: "Old Credential",
				CreationTime:    1600000000000, // Old
			},
			{
				CredentialsId:   "cred-2",
				CredentialsName: "New Credential",
				CreationTime:    1700000000000, // New
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "mws")
		ic.incremental = true
		ic.updatedSinceMs = 1650000000000

		err := resourcesMap["databricks_mws_credentials"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_mws_credentials[New Credential] (id: "+testAccountID+"/cred-2)"])
	})
}

func TestListMwsCredentialsAzure(t *testing.T) {
	qa.MockAccountsApply(t, setupAzureAccountConfig,
		func(ctx context.Context, client *common.DatabricksClient) {
			ic := importContextForAccountTestWithClient(ctx, client, "mws")

			err := resourcesMap["databricks_mws_credentials"].List(ic)
			assert.NoError(t, err)
			require.Equal(t, 0, len(ic.testEmits))
		})
}

// Tests for MWS Storage Configurations

func TestListMwsStorageConfigurations(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		ma.GetMockStorageAPI().EXPECT().List(mock.Anything).Return([]provisioning.StorageConfiguration{
			{
				StorageConfigurationId:   "storage-1",
				StorageConfigurationName: "Test Storage 1",
				CreationTime:             1700000000000,
			},
			{
				StorageConfigurationId:   "storage-2",
				StorageConfigurationName: "Test Storage 2",
				CreationTime:             1700000000000,
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "mws")

		err := resourcesMap["databricks_mws_storage_configurations"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_mws_storage_configurations[Test Storage 1] (id: "+testAccountID+"/storage-1)"])
		assert.True(t, ic.testEmits["databricks_mws_storage_configurations[Test Storage 2] (id: "+testAccountID+"/storage-2)"])
	})
}

func TestListMwsStorageConfigurationsIncremental(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		ma.GetMockStorageAPI().EXPECT().List(mock.Anything).Return([]provisioning.StorageConfiguration{
			{
				StorageConfigurationId:   "storage-1",
				StorageConfigurationName: "Old Storage",
				CreationTime:             1600000000000,
			},
			{
				StorageConfigurationId:   "storage-2",
				StorageConfigurationName: "New Storage",
				CreationTime:             1700000000000,
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "mws")
		ic.incremental = true
		ic.updatedSinceMs = 1650000000000

		err := resourcesMap["databricks_mws_storage_configurations"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_mws_storage_configurations[New Storage] (id: "+testAccountID+"/storage-2)"])
	})
}

func TestListMwsStorageConfigurationsAzure(t *testing.T) {
	qa.MockAccountsApply(t, setupAzureAccountConfig,
		func(ctx context.Context, client *common.DatabricksClient) {
			ic := importContextForAccountTestWithClient(ctx, client, "mws")

			err := resourcesMap["databricks_mws_storage_configurations"].List(ic)
			assert.NoError(t, err)
			require.Equal(t, 0, len(ic.testEmits))
		})
}

// Tests for MWS VPC Endpoints

func TestListMwsVpcEndpoints(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		ma.GetMockVpcEndpointsAPI().EXPECT().List(mock.Anything).Return([]provisioning.VpcEndpoint{
			{
				VpcEndpointId:   "vpce-1",
				VpcEndpointName: "Test VPC Endpoint 1",
			},
			{
				VpcEndpointId:   "vpce-2",
				VpcEndpointName: "Test VPC Endpoint 2",
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "mws")

		err := resourcesMap["databricks_mws_vpc_endpoint"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_mws_vpc_endpoint[<unknown>] (id: "+testAccountID+"/vpce-1)"])
		assert.True(t, ic.testEmits["databricks_mws_vpc_endpoint[<unknown>] (id: "+testAccountID+"/vpce-2)"])
	})
}

func TestListMwsVpcEndpointsWithMatch(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		ma.GetMockVpcEndpointsAPI().EXPECT().List(mock.Anything).Return([]provisioning.VpcEndpoint{
			{
				VpcEndpointId:   "vpce-1",
				VpcEndpointName: "Test VPC Endpoint",
			},
			{
				VpcEndpointId:   "vpce-2",
				VpcEndpointName: "Production VPC Endpoint",
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "mws")
		ic.match = "Test"

		err := resourcesMap["databricks_mws_vpc_endpoint"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_mws_vpc_endpoint[<unknown>] (id: "+testAccountID+"/vpce-1)"])
	})
}

func TestListMwsVpcEndpointsAzure(t *testing.T) {
	qa.MockAccountsApply(t, setupAzureAccountConfig,
		func(ctx context.Context, client *common.DatabricksClient) {
			ic := importContextForAccountTestWithClient(ctx, client, "mws")

			err := resourcesMap["databricks_mws_vpc_endpoint"].List(ic)
			assert.NoError(t, err)
			require.Equal(t, 0, len(ic.testEmits))
		})
}

// Tests for MWS Private Access Settings

func TestListMwsPrivateAccessSettings(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		ma.GetMockPrivateAccessAPI().EXPECT().List(mock.Anything).Return([]provisioning.PrivateAccessSettings{
			{
				PrivateAccessSettingsId:   "pas-1",
				PrivateAccessSettingsName: "Test PAS 1",
			},
			{
				PrivateAccessSettingsId:   "pas-2",
				PrivateAccessSettingsName: "Test PAS 2",
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "mws")

		err := resourcesMap["databricks_mws_private_access_settings"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_mws_private_access_settings[Test PAS 1] (id: "+testAccountID+"/pas-1)"])
		assert.True(t, ic.testEmits["databricks_mws_private_access_settings[Test PAS 2] (id: "+testAccountID+"/pas-2)"])
	})
}

func TestListMwsPrivateAccessSettingsAzure(t *testing.T) {
	qa.MockAccountsApply(t, setupAzureAccountConfig,
		func(ctx context.Context, client *common.DatabricksClient) {
			ic := importContextForAccountTestWithClient(ctx, client, "mws")

			err := resourcesMap["databricks_mws_private_access_settings"].List(ic)
			assert.NoError(t, err)
			require.Equal(t, 0, len(ic.testEmits))
		})
}

func TestImportMwsPrivateAccessSettings(t *testing.T) {
	qa.MockAccountsApply(t, setupAwsAccountConfig,
		func(ctx context.Context, client *common.DatabricksClient) {
			ic := importContextForAccountTestWithClient(ctx, client, "mws")

			d := mws.ResourceMwsPrivateAccessSettings().ToResource().TestResourceData()
			d.SetId(testAccountID + "/pas-1")
			d.Set("private_access_settings_id", "pas-1")
			d.Set("allowed_vpc_endpoint_ids", []interface{}{"vpce-1", "vpce-2"})

			r := &resource{
				ID:       testAccountID + "/pas-1",
				Resource: "databricks_mws_private_access_settings",
				Data:     d,
			}

			err := resourcesMap["databricks_mws_private_access_settings"].Import(ic, r)
			assert.NoError(t, err)
			require.Equal(t, 2, len(ic.testEmits))
			assert.True(t, ic.testEmits["databricks_mws_vpc_endpoint[<unknown>] (id: "+testAccountID+"/vpce-1)"])
			assert.True(t, ic.testEmits["databricks_mws_vpc_endpoint[<unknown>] (id: "+testAccountID+"/vpce-2)"])
		})
}

func TestImportMwsPrivateAccessSettingsNoEndpoints(t *testing.T) {
	qa.MockAccountsApply(t, setupAwsAccountConfig,
		func(ctx context.Context, client *common.DatabricksClient) {
			ic := importContextForAccountTestWithClient(ctx, client, "mws")

			d := mws.ResourceMwsPrivateAccessSettings().ToResource().TestResourceData()
			d.SetId(testAccountID + "/pas-1")
			d.Set("private_access_settings_id", "pas-1")

			r := &resource{
				ID:       testAccountID + "/pas-1",
				Resource: "databricks_mws_private_access_settings",
				Data:     d,
			}

			err := resourcesMap["databricks_mws_private_access_settings"].Import(ic, r)
			assert.NoError(t, err)
			require.Equal(t, 0, len(ic.testEmits))
		})
}

// Tests for MWS Customer Managed Keys

func TestListMwsCustomerManagedKeys(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		ma.GetMockEncryptionKeysAPI().EXPECT().List(mock.Anything).Return([]provisioning.CustomerManagedKey{
			{
				CustomerManagedKeyId: "cmk-1",
				CreationTime:         1700000000000,
			},
			{
				CustomerManagedKeyId: "cmk-2",
				CreationTime:         1700000000000,
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "mws")

		err := resourcesMap["databricks_mws_customer_managed_keys"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_mws_customer_managed_keys[cmk-1] (id: "+testAccountID+"/cmk-1)"])
		assert.True(t, ic.testEmits["databricks_mws_customer_managed_keys[cmk-2] (id: "+testAccountID+"/cmk-2)"])
	})
}

func TestListMwsCustomerManagedKeysIncremental(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		ma.GetMockEncryptionKeysAPI().EXPECT().List(mock.Anything).Return([]provisioning.CustomerManagedKey{
			{
				CustomerManagedKeyId: "cmk-1",
				CreationTime:         1600000000000,
			},
			{
				CustomerManagedKeyId: "cmk-2",
				CreationTime:         1700000000000,
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "mws")
		ic.incremental = true
		ic.updatedSinceMs = 1650000000000

		err := resourcesMap["databricks_mws_customer_managed_keys"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_mws_customer_managed_keys[cmk-2] (id: "+testAccountID+"/cmk-2)"])
	})
}

func TestListMwsCustomerManagedKeysAzure(t *testing.T) {
	qa.MockAccountsApply(t, setupAzureAccountConfig,
		func(ctx context.Context, client *common.DatabricksClient) {
			ic := importContextForAccountTestWithClient(ctx, client, "mws")

			err := resourcesMap["databricks_mws_customer_managed_keys"].List(ic)
			assert.NoError(t, err)
			require.Equal(t, 0, len(ic.testEmits))
		})
}

// Tests for MWS Networks

func TestListMwsNetworks(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		ma.GetMockNetworksAPI().EXPECT().List(mock.Anything).Return([]provisioning.Network{
			{
				NetworkId:   "network-1",
				NetworkName: "Test Network 1",
			},
			{
				NetworkId:   "network-2",
				NetworkName: "Test Network 2",
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "mws")

		err := resourcesMap["databricks_mws_networks"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_mws_networks[Test Network 1] (id: "+testAccountID+"/network-1)"])
		assert.True(t, ic.testEmits["databricks_mws_networks[Test Network 2] (id: "+testAccountID+"/network-2)"])
	})
}

func TestListMwsNetworksAzure(t *testing.T) {
	qa.MockAccountsApply(t, setupAzureAccountConfig,
		func(ctx context.Context, client *common.DatabricksClient) {
			ic := importContextForAccountTestWithClient(ctx, client, "mws")

			err := resourcesMap["databricks_mws_networks"].List(ic)
			assert.NoError(t, err)
			require.Equal(t, 0, len(ic.testEmits))
		})
}

func TestImportMwsNetworks(t *testing.T) {
	qa.MockAccountsApply(t, setupAwsAccountConfig,
		func(ctx context.Context, client *common.DatabricksClient) {
			ic := importContextForAccountTestWithClient(ctx, client, "mws")

			d := mws.ResourceMwsNetworks().ToResource().TestResourceData()
			d.SetId(testAccountID + "/network-1")

			// Set VPC endpoints
			vpcEndpoints := []interface{}{
				map[string]interface{}{
					"dataplane_relay": []interface{}{"vpce-relay-1", "vpce-relay-2"},
					"rest_api":        []interface{}{"vpce-api-1"},
				},
			}
			d.Set("vpc_endpoints", vpcEndpoints)

			r := &resource{
				ID:       testAccountID + "/network-1",
				Resource: "databricks_mws_networks",
				Data:     d,
			}

			err := resourcesMap["databricks_mws_networks"].Import(ic, r)
			assert.NoError(t, err)
			require.Equal(t, 3, len(ic.testEmits))
			assert.True(t, ic.testEmits["databricks_mws_vpc_endpoint[<unknown>] (id: "+testAccountID+"/vpce-relay-1)"])
			assert.True(t, ic.testEmits["databricks_mws_vpc_endpoint[<unknown>] (id: "+testAccountID+"/vpce-relay-2)"])
			assert.True(t, ic.testEmits["databricks_mws_vpc_endpoint[<unknown>] (id: "+testAccountID+"/vpce-api-1)"])
		})
}

func TestImportMwsNetworksNoEndpoints(t *testing.T) {
	qa.MockAccountsApply(t, setupAwsAccountConfig,
		func(ctx context.Context, client *common.DatabricksClient) {
			ic := importContextForAccountTestWithClient(ctx, client, "mws")

			d := mws.ResourceMwsNetworks().ToResource().TestResourceData()
			d.SetId(testAccountID + "/network-1")

			r := &resource{
				ID:       testAccountID + "/network-1",
				Resource: "databricks_mws_networks",
				Data:     d,
			}

			err := resourcesMap["databricks_mws_networks"].Import(ic, r)
			assert.NoError(t, err)
			require.Equal(t, 0, len(ic.testEmits))
		})
}

// Tests for MWS Workspaces

func TestListMwsWorkspaces(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		ma.GetMockWorkspacesAPI().EXPECT().List(mock.Anything).Return([]provisioning.Workspace{
			{
				WorkspaceId:     123,
				WorkspaceName:   "Test Workspace 1",
				WorkspaceStatus: "RUNNING",
				CreationTime:    1700000000000,
			},
			{
				WorkspaceId:     456,
				WorkspaceName:   "Test Workspace 2",
				WorkspaceStatus: "RUNNING",
				CreationTime:    1700000000000,
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "mws")

		err := resourcesMap["databricks_mws_workspaces"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_mws_workspaces[Test Workspace 1_123] (id: "+testAccountID+"/123)"])
		assert.True(t, ic.testEmits["databricks_mws_workspaces[Test Workspace 2_456] (id: "+testAccountID+"/456)"])
	})
}

func TestListMwsWorkspacesSkipNonRunning(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		ma.GetMockWorkspacesAPI().EXPECT().List(mock.Anything).Return([]provisioning.Workspace{
			{
				WorkspaceId:     123,
				WorkspaceName:   "Running Workspace",
				WorkspaceStatus: "RUNNING",
				CreationTime:    1700000000000,
			},
			{
				WorkspaceId:     456,
				WorkspaceName:   "Stopped Workspace",
				WorkspaceStatus: "STOPPED",
				CreationTime:    1700000000000,
			},
			{
				WorkspaceId:     789,
				WorkspaceName:   "Failed Workspace",
				WorkspaceStatus: "FAILED",
				CreationTime:    1700000000000,
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "mws")

		err := resourcesMap["databricks_mws_workspaces"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_mws_workspaces[Running Workspace_123] (id: "+testAccountID+"/123)"])
		assert.False(t, ic.testEmits["databricks_mws_workspaces[Stopped Workspace_456] (id: "+testAccountID+"/456)"])
	})
}

func TestListMwsWorkspacesIncremental(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		ma.GetMockWorkspacesAPI().EXPECT().List(mock.Anything).Return([]provisioning.Workspace{
			{
				WorkspaceId:     123,
				WorkspaceName:   "Old Workspace",
				WorkspaceStatus: "RUNNING",
				CreationTime:    1600000000000,
			},
			{
				WorkspaceId:     456,
				WorkspaceName:   "New Workspace",
				WorkspaceStatus: "RUNNING",
				CreationTime:    1700000000000,
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "mws")
		ic.incremental = true
		ic.updatedSinceMs = 1650000000000

		err := resourcesMap["databricks_mws_workspaces"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_mws_workspaces[New Workspace_456] (id: "+testAccountID+"/456)"])
	})
}

func TestListMwsWorkspacesAzure(t *testing.T) {
	qa.MockAccountsApply(t, setupAzureAccountConfig,
		func(ctx context.Context, client *common.DatabricksClient) {
			ic := importContextForAccountTestWithClient(ctx, client, "mws")

			err := resourcesMap["databricks_mws_workspaces"].List(ic)
			assert.NoError(t, err)
			require.Equal(t, 0, len(ic.testEmits))
		})
}

func TestImportMwsWorkspaces(t *testing.T) {
	qa.MockAccountsApply(t, setupAwsAccountConfig,
		func(ctx context.Context, client *common.DatabricksClient) {
			ic := importContextForAccountTestWithClient(ctx, client, "mws")

			d := mws.ResourceMwsWorkspaces().ToResource().TestResourceData()
			d.SetId(testAccountID + "/123")
			d.Set("workspace_id", 123)
			d.Set("network_id", "network-1")
			d.Set("private_access_settings_id", "pas-1")
			d.Set("storage_configuration_id", "storage-1")
			d.Set("managed_services_customer_managed_key_id", "cmk-2")
			d.Set("credentials_id", "cred-1")

			r := &resource{
				ID:       testAccountID + "/123",
				Resource: "databricks_mws_workspaces",
				Data:     d,
			}

			err := resourcesMap["databricks_mws_workspaces"].Import(ic, r)
			assert.NoError(t, err)
			require.Equal(t, 5, len(ic.testEmits))
			assert.True(t, ic.testEmits["databricks_mws_networks[<unknown>] (id: "+testAccountID+"/network-1)"])
			assert.True(t, ic.testEmits["databricks_mws_private_access_settings[<unknown>] (id: "+testAccountID+"/pas-1)"])
			assert.True(t, ic.testEmits["databricks_mws_storage_configurations[<unknown>] (id: "+testAccountID+"/storage-1)"])
			assert.True(t, ic.testEmits["databricks_mws_customer_managed_keys[<unknown>] (id: "+testAccountID+"/cmk-2)"])
			assert.True(t, ic.testEmits["databricks_mws_credentials[<unknown>] (id: "+testAccountID+"/cred-1)"])
		})
}

func TestImportMwsWorkspacesMinimal(t *testing.T) {
	qa.MockAccountsApply(t, setupAwsAccountConfig,
		func(ctx context.Context, client *common.DatabricksClient) {
			ic := importContextForAccountTestWithClient(ctx, client, "mws")

			d := mws.ResourceMwsWorkspaces().ToResource().TestResourceData()
			d.SetId(testAccountID + "/123")
			d.Set("workspace_id", 123)

			r := &resource{
				ID:       testAccountID + "/123",
				Resource: "databricks_mws_workspaces",
				Data:     d,
			}

			err := resourcesMap["databricks_mws_workspaces"].Import(ic, r)
			assert.NoError(t, err)
			require.Equal(t, 0, len(ic.testEmits))
		})
}

func TestImportMwsWorkspacesWithNccBinding(t *testing.T) {
	qa.MockAccountsApply(t, setupAwsAccountConfig,
		func(ctx context.Context, client *common.DatabricksClient) {
			ic := importContextForAccountTestWithClient(ctx, client, "mws,nccs")

			d := mws.ResourceMwsWorkspaces().ToResource().TestResourceData()
			d.SetId(testAccountID + "/123")
			d.Set("workspace_id", 123)
			d.Set("network_connectivity_config_id", "ncc-1")

			r := &resource{
				ID:       testAccountID + "/123",
				Resource: "databricks_mws_workspaces",
				Data:     d,
			}

			err := resourcesMap["databricks_mws_workspaces"].Import(ic, r)
			assert.NoError(t, err)
			require.Equal(t, 2, len(ic.testEmits))
			assert.True(t, ic.testEmits["databricks_mws_ncc_binding[ws_123_ncc-1] (id: 123/ncc-1)"])
			assert.True(t, ic.testEmits["databricks_mws_network_connectivity_config[<unknown>] (id: "+testAccountID+"/ncc-1)"])
		})
}
