package common

import (
	"os"
	"sync"
)

var (
	onceClient   sync.Once
	commonClient *DatabricksClient
)

// NewClientFromEnvironment makes very good client for testing purposes
func NewClientFromEnvironment() *DatabricksClient {
	client := DatabricksClient{
		Host:       os.Getenv("DATABRICKS_HOST"),
		Token:      os.Getenv("DATABRICKS_TOKEN"),
		Username:   os.Getenv("DATABRICKS_USERNAME"),
		Password:   os.Getenv("DATABRICKS_PASSWORD"),
		ConfigFile: os.Getenv("DATABRICKS_CONFIG_FILE"),
		Profile:    os.Getenv("DATABRICKS_CONFIG_PROFILE"),
		AzureAuth: AzureAuth{
			ResourceID:     os.Getenv("DATABRICKS_AZURE_WORKSPACE_RESOURCE_ID"),
			WorkspaceName:  os.Getenv("DATABRICKS_AZURE_WORKSPACE_NAME"),
			ResourceGroup:  os.Getenv("DATABRICKS_AZURE_RESOURCE_GROUP"),
			SubscriptionID: os.Getenv("ARM_SUBSCRIPTION_ID"),
			ClientID:       os.Getenv("ARM_CLIENT_ID"),
			ClientSecret:   os.Getenv("ARM_CLIENT_SECRET"),
			TenantID:       os.Getenv("ARM_TENANT_ID"),
		},
	}
	err := client.Configure()
	if err != nil {
		panic(err)
	}
	return &client
}

// CommonEnvironmentClient configured once per run of application
func CommonEnvironmentClient() *DatabricksClient {
	if commonClient != nil {
		return commonClient
	}
	onceClient.Do(func() {
		commonClient = NewClientFromEnvironment()
	})
	return commonClient
}
