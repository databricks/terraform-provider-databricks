package config

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadAzureTenantId(t *testing.T) {
	c := &Config{
		Host: "https://adb-xyz.c.azuredatabricks.net/",
		azureTenantIdFetchClient: makeClient(&http.Response{
			StatusCode: 302,
			Header:     http.Header{"Location": []string{"https://login.microsoftonline.com/123-abc/oauth2/token"}},
		}),
	}
	err := c.loadAzureTenantId(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, c.AzureTenantID, "123-abc")
}

func TestLoadAzureTenantId_Failure(t *testing.T) {
	testErr := errors.New("Failed to fetch login page")
	c := &Config{
		Host:                     "https://adb-xyz.c.azuredatabricks.net/",
		azureTenantIdFetchClient: makeFailingClient(testErr),
	}
	err := c.loadAzureTenantId(context.Background())
	assert.ErrorIs(t, err, testErr)
}

func TestLoadAzureTenantId_SkipNotInAzure(t *testing.T) {
	testErr := errors.New("Failed to fetch login page")
	c := &Config{
		Host:                     "https://test.cloud.databricks.com/",
		azureTenantIdFetchClient: makeFailingClient(testErr),
	}
	err := c.loadAzureTenantId(context.Background())
	assert.NoError(t, err)
	assert.Empty(t, c.AzureTenantID)
}
