package databricks

import (
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/stretchr/testify/assert"
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
			ManagedResourceGroup: os.Getenv("DATABRICKS_AZURE_MANAGED_RESOURCE_GROUP"),
			AzureRegion:          os.Getenv("AZURE_REGION"),
			WorkspaceName:        os.Getenv("DATABRICKS_AZURE_WORKSPACE_NAME"),
			ResourceGroup:        os.Getenv("DATABRICKS_AZURE_RESOURCE_GROUP"),
		},
		ManagementToken:        "",
		AdbWorkspaceResourceID: "",
		AdbAccessToken:         "",
		AdbPlatformToken:       "",
	}
	azureAuth.TokenPayload.SubscriptionID = os.Getenv("DATABRICKS_AZURE_SUBSCRIPTION_ID")
	azureAuth.TokenPayload.TenantID = os.Getenv("DATABRICKS_AZURE_TENANT_ID")
	azureAuth.TokenPayload.ClientID = os.Getenv("DATABRICKS_AZURE_CLIENT_ID")
	azureAuth.TokenPayload.ClientSecret = os.Getenv("DATABRICKS_AZURE_CLIENT_SECRET")
	config := GetIntegrationDBClientOptions()
	err := azureAuth.initWorkspaceAndGetClient(config)
	assert.NoError(t, err, err)
	api := service.DBApiClient{}
	api.SetConfig(config)
	instancePoolInfo, instancePoolErr := api.InstancePools().Create(model.InstancePool{
		InstancePoolName:                   "my_instance_pool",
		MinIdleInstances:                   0,
		MaxCapacity:                        10,
		NodeTypeID:                         "Standard_DS3_v2",
		IdleInstanceAutoTerminationMinutes: 20,
		PreloadedSparkVersions: []string{
			"6.3.x-scala2.11",
		},
	})
	defer func() {
		err := api.InstancePools().Delete(instancePoolInfo.InstancePoolID)
		assert.NoError(t, err, err)
	}()

	assert.NoError(t, instancePoolErr, instancePoolErr)
}
