package common

import (
	"os"
	"testing"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/stretchr/testify/assert"
)

func TestCommonEnvironmentClient(t *testing.T) {
	ResetCommonEnvironmentClient()
	defer CleanupEnvironment()()
	os.Setenv("DATABRICKS_TOKEN", ".")
	os.Setenv("DATABRICKS_HOST", ".")
	c := CommonEnvironmentClient()
	c2 := CommonEnvironmentClient()
	assert.Equal(t, UserAgent(), c.userAgent)
	assert.Equal(t, c2.userAgent, c.userAgent)
}

func TestGetEnvironment(t *testing.T) {
	defer CleanupEnvironment()()

	os.Setenv("ARM_ENVIRONMENT", "public")
	assert.Equal(t, GetAzureEnvironment(), azure.PublicCloud)
	os.Setenv("ARM_ENVIRONMENT", "china")
	assert.Equal(t, GetAzureEnvironment(), azure.ChinaCloud)
	os.Setenv("ARM_ENVIRONMENT", "german")
	assert.Equal(t, GetAzureEnvironment(), azure.GermanCloud)
	os.Setenv("ARM_ENVIRONMENT", "usgovernment")
	assert.Equal(t, GetAzureEnvironment(), azure.USGovernmentCloud)
	os.Setenv("ARM_ENVIRONMENT", "")
	assert.Equal(t, GetAzureEnvironment(), azure.PublicCloud)
}
