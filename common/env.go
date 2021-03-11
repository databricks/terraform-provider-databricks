package common

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/mitchellh/go-homedir"
)

var (
	envMutex     sync.Mutex
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
			Environment:    os.Getenv("ARM_ENVIRONMENT"),
		},
		RateLimitPerSecond: 10,
	}
	err := client.Configure()
	if err != nil {
		panic(err)
	}
	return &client
}

// ResetCommonEnvironmentClient resets test dummy
func ResetCommonEnvironmentClient() {
	commonClient = nil
	onceClient = sync.Once{}
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

// CleanupEnvironment backs up environment - use as `defer CleanupEnvironment()()`
// clears it and restores it in the end. It's meant strictly for "unit" tests
// as last resort, because it slows down parallel execution with mutex.
func CleanupEnvironment() func() {
	// make a backed-up pristine environment
	envMutex.Lock()
	prevEnv := os.Environ()
	oldPath := os.Getenv("PATH")
	pwd := os.Getenv("PWD")
	os.Clearenv()
	err := os.Setenv("PATH", oldPath)
	if err != nil {
		log.Printf("[WARN] Cannot bring back PATH: %v", err)
	}
	err = os.Setenv("HOME", pwd)
	if err != nil {
		log.Printf("[WARN] Cannot set HOME to old PWD: %v", err)
	}
	homedir.DisableCache = true
	// and return restore function
	return func() {
		for _, kv := range prevEnv {
			kvs := strings.SplitN(kv, "=", 2)
			os.Setenv(kvs[0], kvs[1])
		}
		envMutex.Unlock()
	}
}
