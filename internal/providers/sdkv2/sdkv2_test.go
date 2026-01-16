package sdkv2

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestConfigureDatabricksClient(t *testing.T) {
	testCases := []struct {
		name                 string
		config               map[string]interface{}
		validateResourceData func(*common.DatabricksClient)
	}{
		{
			name:   "HTTP timeout set to 65 seconds by default",
			config: map[string]interface{}{},
			validateResourceData: func(dc *common.DatabricksClient) {
				assert.Equal(t, 65, dc.Config.HTTPTimeoutSeconds, "HTTP timeout should be 65 seconds by default")
			},
		},
		{
			name: "HTTP timeout can be overridden in provider config",
			config: map[string]interface{}{
				"http_timeout_seconds": 30,
			},
			validateResourceData: func(dc *common.DatabricksClient) {
				assert.Equal(t, 30, dc.Config.HTTPTimeoutSeconds, "HTTP timeout should be overridden when set")
			},
		},
		{
			name:   "Scopes defaults to nil when not set",
			config: map[string]interface{}{},
			validateResourceData: func(dc *common.DatabricksClient) {
				assert.Nil(t, dc.Config.Scopes, "Scopes should be nil by default")
			},
		},
		{
			name: "Scopes can be set in provider config",
			config: map[string]interface{}{
				"scopes": []interface{}{"clusters", "jobs"},
			},
			validateResourceData: func(dc *common.DatabricksClient) {
				assert.Equal(t, []string{"clusters", "jobs"}, dc.Config.Scopes, "Scopes should be set from provider config")
			},
		},
		{
			name: "Single scope can be set in provider config",
			config: map[string]interface{}{
				"scopes": []interface{}{"all-apis"},
			},
			validateResourceData: func(dc *common.DatabricksClient) {
				assert.Equal(t, []string{"all-apis"}, dc.Config.Scopes, "Single scope should be set from provider config")
			},
		},
		{
			name: "Empty scopes list is treated as unset",
			config: map[string]interface{}{
				"scopes": []interface{}{},
			},
			validateResourceData: func(dc *common.DatabricksClient) {
				assert.Nil(t, dc.Config.Scopes, "Empty scopes list should be treated as unset")
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a provider instance
			p := DatabricksProvider()

			// Create a resource data with the test configuration
			rd := schema.TestResourceDataRaw(t, p.Schema, tc.config)

			// Call ConfigureDatabricksClient
			client, diags := ConfigureDatabricksClient(context.Background(), rd, nil)
			assert.False(t, diags.HasError(), "ConfigureDatabricksClient should not return errors")

			// Get the client from the response
			dc, ok := client.(*common.DatabricksClient)
			assert.True(t, ok, "client should be a DatabricksClient")

			// Validate the client configuration
			tc.validateResourceData(dc)
		})
	}
}
