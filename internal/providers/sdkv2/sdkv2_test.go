package sdkv2

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

// TestPluginFrameworkOptInsExistInSdkV2 asserts that every resource and data
// source in pluginfw's pluginFwOptInResources / pluginFwOptInDataSources lists
// has a matching entry in the SDK V2 provider's maps. Otherwise, a user
// opting in via DATABRICKS_TF_ENABLED_PF_RESOURCES would crash the provider
// at startup when GetSdkV2ResourcesToRemove's caller tries to remove a name
// that isn't there.
//
// The test lives here rather than in pluginfw because pluginfw cannot import
// sdkv2 (sdkv2 imports pluginfw); this is the closest reachable seam.
func TestPluginFrameworkOptInsExistInSdkV2(t *testing.T) {
	p := DatabricksProvider()
	for _, name := range pluginfw.PluginFrameworkOptInResourceNames() {
		if _, ok := p.ResourcesMap[name]; !ok {
			t.Errorf("resource %q is in pluginFwOptInResources but has no SDKv2 implementation; opting in would panic at startup", name)
		}
	}
	for _, name := range pluginfw.PluginFrameworkOptInDataSourceNames() {
		if _, ok := p.DataSourcesMap[name]; !ok {
			t.Errorf("data source %q is in pluginFwOptInDataSources but has no SDKv2 implementation; opting in would panic at startup", name)
		}
	}
}

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
			name: "workspace_id can be set in provider config",
			config: map[string]interface{}{
				"host":         "https://accounts.cloud.databricks.com",
				"account_id":   "00000000-0000-0000-0000-000000000001",
				"workspace_id": "1234567890",
			},
			validateResourceData: func(dc *common.DatabricksClient) {
				assert.Equal(t, "1234567890", dc.Config.WorkspaceID, "workspace_id should be set when provided")
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
