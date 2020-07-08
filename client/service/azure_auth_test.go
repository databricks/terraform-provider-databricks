package service

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/stretchr/testify/assert"
)

// getAndAssertEnv fetches the env for testing and also asserts that the env value is not Zero i.e ""
func getAndAssertEnv(t *testing.T, key string) string {
	value, present := os.LookupEnv(key)
	assert.True(t, present, fmt.Sprintf("Env variable %s is not set", key))
	return value
}

func TestAccAzureAuth_TestPatTokenDuration(t *testing.T) {
	if _, ok := os.LookupEnv("TF_ACC"); !ok {
		t.Skip("Acceptance tests skipped unless env 'TF_ACC' set")
	}

	patTokenSeconds := (time.Duration(1) * time.Hour).Seconds()
	client := DatabricksClient{
		AzureAuth: AzureAuth{
			ManagedResourceGroup: getAndAssertEnv(t, "DATABRICKS_AZURE_MANAGED_RESOURCE_GROUP"),
			AzureRegion:          getAndAssertEnv(t, "AZURE_REGION"),
			WorkspaceName:        getAndAssertEnv(t, "DATABRICKS_AZURE_WORKSPACE_NAME"),
			ResourceGroup:        getAndAssertEnv(t, "DATABRICKS_AZURE_RESOURCE_GROUP"),
			SubscriptionID:       getAndAssertEnv(t, "DATABRICKS_AZURE_SUBSCRIPTION_ID"),
			TenantID:             getAndAssertEnv(t, "DATABRICKS_AZURE_TENANT_ID"),
			ClientID:             getAndAssertEnv(t, "DATABRICKS_AZURE_CLIENT_ID"),
			ClientSecret:         getAndAssertEnv(t, "DATABRICKS_AZURE_CLIENT_SECRET"),
			PatTokenSeconds:      int32(patTokenSeconds),
		},
	}
	err := client.Configure("dev-integration")
	assert.NoError(t, err, err)

	err = client.AzureAuth.initWorkspaceAndGetClient()
	assert.NoError(t, err, err)
	// Time in milliseconds
	tokenActualDuration := client.tokenExpiryTime - client.tokenCreateTime
	assert.Equal(t, patTokenSeconds, (time.Duration(tokenActualDuration) * time.Millisecond).Seconds(),
		"duration should be the same")
}

func TestAzureAuthCreateApiToken(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	patTokenSeconds := (time.Duration(1) * time.Hour).Seconds()
	client := DatabricksClient{
		AzureAuth: AzureAuth{
			ManagedResourceGroup: getAndAssertEnv(t, "DATABRICKS_AZURE_MANAGED_RESOURCE_GROUP"),
			AzureRegion:          getAndAssertEnv(t, "AZURE_REGION"),
			WorkspaceName:        getAndAssertEnv(t, "DATABRICKS_AZURE_WORKSPACE_NAME"),
			ResourceGroup:        getAndAssertEnv(t, "DATABRICKS_AZURE_RESOURCE_GROUP"),
			SubscriptionID:       getAndAssertEnv(t, "DATABRICKS_AZURE_SUBSCRIPTION_ID"),
			TenantID:             getAndAssertEnv(t, "DATABRICKS_AZURE_TENANT_ID"),
			ClientID:             getAndAssertEnv(t, "DATABRICKS_AZURE_CLIENT_ID"),
			ClientSecret:         getAndAssertEnv(t, "DATABRICKS_AZURE_CLIENT_SECRET"),
			PatTokenSeconds:      int32(patTokenSeconds),
		},
	}
	err := client.Configure("dev-integration")
	assert.NoError(t, err, err)
	instancePoolInfo, instancePoolErr := client.InstancePools().Create(model.InstancePool{
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
		err := client.InstancePools().Delete(instancePoolInfo.InstancePoolID)
		assert.NoError(t, err, err)
	}()

	assert.NoError(t, instancePoolErr, instancePoolErr)
}
