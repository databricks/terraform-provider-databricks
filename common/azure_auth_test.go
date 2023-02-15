package common

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetJWTProperty_AzureCLI_SP(t *testing.T) {
	defer CleanupEnvironment()()
	p, _ := filepath.Abs("./testdata")
	os.Setenv("PATH", p+":/bin")

	aa := DatabricksClient{
		AzureClientID:     "a",
		AzureClientSecret: "b",
		AzureTenantID:     "c",
		Host:              "https://adb-1232.azuredatabricks.net",
	}
	tid, err := aa.GetAzureJwtProperty("tid")
	assert.NoError(t, err)
	assert.Equal(t, "c", tid)
}

func TestGetJWTProperty_NonAzure(t *testing.T) {
	defer CleanupEnvironment()()
	p, _ := filepath.Abs("./testdata")
	os.Setenv("PATH", p+":/bin")

	aa := DatabricksClient{
		Host:  "https://abc.cloud.databricks.com",
		Token: "abc",
	}
	_, err := aa.GetAzureJwtProperty("tid")
	require.EqualError(t, err, "can't get Azure JWT token in non-Azure environment")
}

func TestGetJWTProperty_Authenticate_Fail(t *testing.T) {
	defer CleanupEnvironment()()
	p, _ := filepath.Abs("./testdata")
	os.Setenv("PATH", p+":/bin")
	os.Setenv("FAIL", "yes")

	client := &DatabricksClient{
		Config: &config.Config{
			Host: "https://adb-1232.azuredatabricks.net",
		},
	}
	_, err := client.GetAzureJwtProperty("tid")
	require.EqualError(t, err, "cannot configure azure-cli auth: "+
		"Invoking Azure CLI failed with the following error: "+
		"This is just a failing script.\n. "+
		"Please check https://registry.terraform.io/providers/"+
		"databricks/databricks/latest/docs#authentication for details")
}
