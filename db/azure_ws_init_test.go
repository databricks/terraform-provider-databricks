package db

import (
	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/stretchr/testify/assert"
	"os"

	//"os"
	"testing"
)

func GetIntegrationDBClientOptions() *service.DBApiClientConfig {
	var config service.DBApiClientConfig

	return &config
}

func TestAzureAuthCreateApiToken(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	azureAuth := AzureAuth{
		TokenPayload: &TokenPayload{
			ManagedResourceGroup: os.Getenv("TEST_MANAGED_RESOURCE_GROUP"),
			AzureRegion:          "centralus",
			WorkspaceName:        os.Getenv("TEST_WORKSPACE_NAME"),
			ResourceGroup:        os.Getenv("TEST_RESOURCE_GROUP"),
		},
		ManagementToken:        "",
		AdbWorkspaceResourceID: "",
		AdbAccessToken:         "",
		AdbPlatformToken:       "",
	}
	azureAuth.TokenPayload.SubscriptionId = os.Getenv("ARM_SUBSCRIPTION_ID")
	azureAuth.TokenPayload.TenantID = os.Getenv("ARM_TENANT_ID")
	azureAuth.TokenPayload.ClientID = os.Getenv("ARM_CLIENT_ID")
	azureAuth.TokenPayload.ClientSecret = os.Getenv("ARM_CLIENT_SECRET")
	option := GetIntegrationDBClientOptions()
	//Hack
	api, err := azureAuth.initWorkspaceAndGetClient(option)
	//err := azureAuth.getManagementToken(*option)
	assert.NoError(t, err, err)

	instancePoolInfo, instancePoolErr := api.InstancePools().Create(model.InstancePool{
		InstancePoolName:                   "my_instance_pool",
		MinIdleInstances:                   0,
		MaxCapacity:                        10,
		NodeTypeId:                         "Standard_DS3_v2",
		IdleInstanceAutoTerminationMinutes: 20,
		PreloadedSparkVersions: []string{
			"6.3.x-scala2.11",
		},
	})
	defer func() {
		err := api.InstancePools().Delete(instancePoolInfo.InstancePoolId)
		assert.NoError(t, err, err)
	}()

	assert.NoError(t, instancePoolErr, instancePoolErr)
}
