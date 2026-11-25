package config

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/common/environment"
	"github.com/stretchr/testify/assert"
)

func TestDefaultEnvironmentIsReturned(t *testing.T) {
	c := &Config{}
	assert.Equal(t, ".cloud.databricks.com", c.Environment().DnsZone)
}

func TestOverriddenEnvironmentIsReturned(t *testing.T) {
	c := &Config{
		Host: "something.else",
		DatabricksEnvironment: &environment.DatabricksEnvironment{
			Cloud:   environment.CloudAzure,
			DnsZone: "holla",
		},
	}
	assert.Equal(t, "holla", c.Environment().DnsZone)
}

func TestCloudField_AWS(t *testing.T) {
	c := &Config{
		Host:  "https://my-workspace.cloud.databricks.com",
		Cloud: environment.CloudAWS,
	}
	assert.True(t, c.IsAws())
	assert.False(t, c.IsAzure())
	assert.False(t, c.IsGcp())
}

func TestCloudField_Azure(t *testing.T) {
	c := &Config{
		Host:  "https://my-workspace.azuredatabricks.net",
		Cloud: environment.CloudAzure,
	}
	assert.True(t, c.IsAzure())
	assert.False(t, c.IsAws())
	assert.False(t, c.IsGcp())
}

func TestCloudField_GCP(t *testing.T) {
	c := &Config{
		Host:  "https://my-workspace.gcp.databricks.com",
		Cloud: environment.CloudGCP,
	}
	assert.True(t, c.IsGcp())
	assert.False(t, c.IsAws())
	assert.False(t, c.IsAzure())
}

func TestCloudField_FallsBackToEnvironment(t *testing.T) {
	c := &Config{
		Host: "https://my-workspace.azuredatabricks.net",
	}
	assert.True(t, c.IsAzure())
	assert.False(t, c.IsAws())
	assert.False(t, c.IsGcp())
}

func TestCloudField_PrefersCloudOverEnvironment(t *testing.T) {
	c := &Config{
		Host:  "https://my-workspace.cloud.databricks.com",
		Cloud: environment.CloudAzure,
	}
	assert.True(t, c.IsAzure())
	assert.False(t, c.IsAws())
}

func TestCloudField_EmptyStringIsCloudUnknown(t *testing.T) {
	c := &Config{
		Cloud: "",
	}
	assert.Equal(t, environment.CloudUnknown, c.Cloud)
}
