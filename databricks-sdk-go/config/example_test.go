package config_test

import (
	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/config"
)

func ExampleConfig_pat() {
	databricks.Must(databricks.NewWorkspaceClient(&databricks.Config{
		Host:  "https://abc.cloud.databricks.com", // env: DATABRICKS_HOST
		Token: "dapi0c2a3f4e...",                  // env: DATABRICKS_TOKEN
	}))
}

func ExampleConfig_basic() {
	databricks.Must(databricks.NewWorkspaceClient(&databricks.Config{
		Host:     "https://abc.cloud.databricks.com", // env: DATABRICKS_HOST
		Username: "me@example.com",                   // env: DATABRICKS_USERNAME
		Password: "som3thing!S@cret",                 // env: DATABRICKS_PASSWORD
	}))
}

func ExampleConfig_accounts() {
	databricks.Must(databricks.NewWorkspaceClient(&databricks.Config{
		Host:      "https://accounts.cloud.databricks.com", // env: DATABRICKS_HOST
		AccountID: "00000000-0000-0000-0000-111122223333",  // env: DATABRICKS_ACCOUNT_ID
		Username:  "me@example.com",                        // env: DATABRICKS_USERNAME
		Password:  "som3thing!S@cret",                      // env: DATABRICKS_PASSWORD
	}))
}

func ExampleConfig_customProfile() {
	databricks.Must(databricks.NewWorkspaceClient(&databricks.Config{
		Profile: "production", // env: DATABRICKS_CONFIG_PROFILE
	}))
}

func ExampleConfig_customConfigFile() {
	databricks.Must(databricks.NewWorkspaceClient(&databricks.Config{
		ConfigFile: "/path/to/.databrickscfg", // env: DATABRICKS_CONFIG_FILE
	}))
}

func ExampleConfig_azureActiveDirectoryServicePrincipal() {
	databricks.Must(databricks.NewWorkspaceClient(&databricks.Config{
		Host:              "https://adb-123.4.azuredatabricks.net", // env: DATABRICKS_HOST
		AzureResourceID:   "/subscriptions/../resourceGroups/...",  // env: DATABRICKS_AZURE_RESOURCE_ID
		AzureTenantID:     "00000000-0000-0000-0000-111122223334",  // env: ARM_TENANT_ID
		AzureClientID:     "00000000-0000-0000-0000-111122223335",  // env: ARM_CLIENT_ID
		AzureClientSecret: "som3thing!S@cret",                      // env: ARM_CLIENT_SECRET
	}))
}

func ExampleConfig_forceAzureActiveDirectoryServicePrincipal() {
	databricks.Must(databricks.NewWorkspaceClient(&databricks.Config{
		Host:              "https://adb-123.4.azuredatabricks.net", // env: DATABRICKS_HOST
		AzureResourceID:   "/subscriptions/../resourceGroups/...",  // env: DATABRICKS_AZURE_RESOURCE_ID
		AzureTenantID:     "00000000-0000-0000-0000-111122223334",  // env: ARM_TENANT_ID
		AzureClientID:     "00000000-0000-0000-0000-111122223335",  // env: ARM_CLIENT_ID
		AzureClientSecret: "som3thing!S@cret",                      // env: ARM_CLIENT_SECRET
		Credentials:       config.AzureClientSecretCredentials{},
	}))
}

func ExampleConfig_debugging() {
	databricks.Must(databricks.NewWorkspaceClient(&databricks.Config{
		HTTPTimeoutSeconds:  60,
		DebugTruncateBytes:  96,    // env: DATABRICKS_DEBUG_TRUNCATE_BYTES
		DebugHeaders:        false, // env: DATABRICKS_DEBUG_HEADERS
		RateLimitPerSecond:  15,    // env: DATABRICKS_RATE_LIMIT
		RetryTimeoutSeconds: 300,
	}))
}
