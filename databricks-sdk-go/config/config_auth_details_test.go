package config

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/internal/env"
	"github.com/stretchr/testify/assert"
)

func TestConfigAuthDetails(t *testing.T) {
	env.CleanupEnvironment(t)

	c := &Config{
		Host:     "https://cloud.databricks.com",
		AuthType: "pat",
		Token:    "test-token",
	}
	t.Setenv("ARM_USE_MSI", "true")
	ConfigAttributes.Configure(c)

	ConfigAttributes.ResolveFromStringMapWithSource(c, map[string]string{
		"azure_client_id": "1234",
		"group_id":        "group-A",
	}, Source{Type: SourceFile, Name: "test.file"})

	authDetails := c.GetAuthDetails()
	assert.Equal(t, "https://cloud.databricks.com", authDetails.Host)
	assert.Equal(t, "pat", authDetails.AuthType)
	assert.Equal(t, "********", authDetails.Configuration["token"].Value)
	assert.Equal(t, "dynamic configuration", authDetails.Configuration["token"].Source.String())
	assert.False(t, authDetails.Configuration["token"].AuthTypeMismatch)

	assert.Equal(t, "true", authDetails.Configuration["azure_use_msi"].Value)
	assert.Equal(t, "ARM_USE_MSI environment variable", authDetails.Configuration["azure_use_msi"].Source.String())
	assert.True(t, authDetails.Configuration["azure_use_msi"].AuthTypeMismatch)

	assert.Equal(t, "1234", authDetails.Configuration["azure_client_id"].Value)
	assert.Equal(t, "test.file config file", authDetails.Configuration["azure_client_id"].Source.String())
	assert.True(t, authDetails.Configuration["azure_client_id"].AuthTypeMismatch)

	assert.Equal(t, "group-A", authDetails.Configuration["group_id"].Value)
	assert.Equal(t, "test.file config file", authDetails.Configuration["group_id"].Source.String())
	assert.False(t, authDetails.Configuration["group_id"].AuthTypeMismatch)
}

func TestConfigAuthDetailsToString(t *testing.T) {
	env.CleanupEnvironment(t)

	c := &Config{
		Host:     "https://cloud.databricks.com",
		AuthType: "pat",
		Token:    "test-token",
	}
	t.Setenv("ARM_USE_MSI", "true")
	ConfigAttributes.Configure(c)

	ConfigAttributes.ResolveFromStringMapWithSource(c, map[string]string{
		"azure_client_id": "1234",
	}, Source{Type: SourceFile, Name: "test.file"})

	authDetails := c.GetAuthDetails()
	expected := `host=https://cloud.databricks.com, auth_type=pat
Configuration:
- host=https://cloud.databricks.com (from dynamic configuration)
- token=******** (from dynamic configuration)
- azure_use_msi=true (from ARM_USE_MSI environment variable) (not used by the current auth type)
- azure_client_id=1234 (from test.file config file) (not used by the current auth type)
- auth_type=pat (from dynamic configuration)`

	assert.Equal(t, expected, authDetails.String())
}
