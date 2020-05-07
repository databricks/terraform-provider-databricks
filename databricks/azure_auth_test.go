package databricks

import (
	"errors"
	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
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
	azureAuth.TokenPayload.SubscriptionID = os.Getenv("DATABRICKS_AZURE_SUBSCRIPTION_ID")
	azureAuth.TokenPayload.TenantID = os.Getenv("DATABRICKS_AZURE_TENANT_ID")
	azureAuth.TokenPayload.ClientID = os.Getenv("DATABRICKS_AZURE_CLIENT_ID")
	azureAuth.TokenPayload.ClientSecret = os.Getenv("DATABRICKS_AZURE_CLIENT_SECRET")
	option := GetIntegrationDBClientOptions()
	api, err := azureAuth.initWorkspaceAndGetClient(option)
	assert.NoError(t, err, err)

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

func TestValidateWorkspaceApis(t *testing.T) {

	// Eventually will pass after a few attempts
	err := validateWorkspaceApis(0, 1, func(attempt int) error {
		log.Printf("Attempt: %v hello world\n", attempt)
		if attempt == 2 {
			return nil
		}
		return errors.New("com.databricks.backend.manager.util.UnknownWorkerEnvironmentException: Unknown worker environment WorkerEnvId")
	})
	assert.NoError(t, err, err)

	// Eventually will fail with an error after a few attempts
	expectedError := errors.New("failing valid error")
	err = validateWorkspaceApis(0, 1, func(attempt int) error {
		log.Printf("Attempt: %v hello world\n", attempt)
		if attempt == 2 {
			return expectedError
		}
		return errors.New("com.databricks.backend.manager.util.UnknownWorkerEnvironmentException: Unknown worker environment WorkerEnvId")
	})
	assert.Error(t, err, err)

	// Eventually will timeout with an error after 0 attempts
	//expectedError := errors.New("failing valid error")
	err = validateWorkspaceApis(0, 0, func(attempt int) error {
		log.Printf("Attempt: %v hello world\n", attempt)
		return errors.New("com.databricks.backend.manager.util.UnknownWorkerEnvironmentException: Unknown worker environment WorkerEnvId")
	})
	assert.Error(t, err, err)
}
