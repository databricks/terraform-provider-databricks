package pluginfw

import (
	"bytes"
	"context"
	"log"
	"testing"

	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
)

// captureFallbackWarnings rewires the package logger to a buffer for the
// duration of the test, restoring the original writer on cleanup.
func captureFallbackWarnings(t *testing.T) *bytes.Buffer {
	t.Helper()
	var buf bytes.Buffer
	origWriter := log.Writer()
	origFlags := log.Flags()
	log.SetOutput(&buf)
	log.SetFlags(0)
	t.Cleanup(func() {
		log.SetOutput(origWriter)
		log.SetFlags(origFlags)
	})
	return &buf
}

func TestGetPluginFrameworkResources_EnvVarFallbackEmitsWarning(t *testing.T) {
	buf := captureFallbackWarnings(t)
	t.Setenv("USE_SDK_V2_RESOURCES", "databricks_library")

	got := getPluginFrameworkResourcesToRegister(nil, nil)

	for _, fn := range got {
		assert.NotEqual(t, "databricks_library", getResourceName(fn),
			"databricks_library must be excluded from PF when fallback is requested")
	}
	assert.Contains(t, buf.String(), `"databricks_library"`)
	assert.Contains(t, buf.String(), "[WARN] resource")
	assert.Contains(t, buf.String(), "next major release")
}

func TestGetPluginFrameworkDataSources_FallbackOptionEmitsWarning(t *testing.T) {
	buf := captureFallbackWarnings(t)

	got := getPluginFrameworkDataSourcesToRegister([]string{"databricks_volumes"}, nil)

	for _, fn := range got {
		assert.NotEqual(t, "databricks_volumes", getDataSourceName(fn),
			"databricks_volumes must be excluded from PF when fallback is requested")
	}
	assert.Contains(t, buf.String(), `"databricks_volumes"`)
	assert.Contains(t, buf.String(), "[WARN] data source")
}

func TestGetPluginFrameworkResources_NoFallbackNoWarning(t *testing.T) {
	buf := captureFallbackWarnings(t)

	_ = getPluginFrameworkResourcesToRegister(nil, nil)

	assert.Empty(t, buf.String(), "no warning should be emitted when no fallback is configured")
}

func TestConfigure(t *testing.T) {
	testCases := []struct {
		name                 string
		config               map[string]tftypes.Value
		validateResourceData func(*common.DatabricksClient)
	}{
		{
			name:   "HTTP timeout set to 65 seconds by default",
			config: map[string]tftypes.Value{},
			validateResourceData: func(dc *common.DatabricksClient) {
				assert.Equal(t, 65, dc.Config.HTTPTimeoutSeconds, "HTTP timeout should be 65 seconds by default")
			},
		},
		{
			name: "HTTP timeout can be overridden in provider config",
			config: map[string]tftypes.Value{
				"http_timeout_seconds": tftypes.NewValue(tftypes.Number, 30),
			},
			validateResourceData: func(dc *common.DatabricksClient) {
				assert.Equal(t, 30, dc.Config.HTTPTimeoutSeconds, "HTTP timeout should be unset by default")
			},
		},
		{
			name: "workspace_id can be set in provider config",
			config: map[string]tftypes.Value{
				"workspace_id": tftypes.NewValue(tftypes.String, "1234567890"),
			},
			validateResourceData: func(dc *common.DatabricksClient) {
				assert.Equal(t, "1234567890", dc.Config.WorkspaceID, "workspace_id should be set when provided")
			},
		},
		{
			name: "user_agent_extra can be set in provider config",
			config: map[string]tftypes.Value{
				"user_agent_extra": tftypes.NewValue(tftypes.String, "pluginfw-config-test/0.0.1"),
			},
			validateResourceData: func(dc *common.DatabricksClient) {
				assert.Contains(t, useragent.FromContext(context.Background()), "pluginfw-config-test/0.0.1",
					"user_agent_extra products should be appended to the user agent")
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a provider instance
			p := GetDatabricksProviderPluginFramework()

			// Create a configure request with no HTTP timeout set
			schema := providerSchemaPluginFramework()
			schemaTfType := schema.Type().TerraformType(context.Background()).(tftypes.Object)
			// TerraformType does not preserve optionality, but all fields are optional
			schemaTfType.OptionalAttributes = map[string]struct{}{}
			for attr := range schemaTfType.AttributeTypes {
				schemaTfType.OptionalAttributes[attr] = struct{}{}
			}
			req := provider.ConfigureRequest{
				Config: tfsdk.Config{
					Raw:    tftypes.NewValue(schemaTfType, tc.config),
					Schema: schema,
				},
			}

			// Create a configure response
			resp := &provider.ConfigureResponse{}

			// Call Configure
			p.Configure(context.Background(), req, resp)

			// Get the client from the response
			client, ok := resp.ResourceData.(*common.DatabricksClient)
			assert.True(t, ok, "ResourceData should be a DatabricksClient")
			tc.validateResourceData(client)
		})
	}
}

func TestConfigure_InvalidUserAgentExtra(t *testing.T) {
	p := GetDatabricksProviderPluginFramework()

	schema := providerSchemaPluginFramework()
	schemaTfType := schema.Type().TerraformType(context.Background()).(tftypes.Object)
	schemaTfType.OptionalAttributes = map[string]struct{}{}
	for attr := range schemaTfType.AttributeTypes {
		schemaTfType.OptionalAttributes[attr] = struct{}{}
	}
	req := provider.ConfigureRequest{
		Config: tfsdk.Config{
			Raw: tftypes.NewValue(schemaTfType, map[string]tftypes.Value{
				"user_agent_extra": tftypes.NewValue(tftypes.String, "not a valid (product)"),
			}),
			Schema: schema,
		},
	}
	resp := &provider.ConfigureResponse{}

	p.Configure(context.Background(), req, resp)

	assert.True(t, resp.Diagnostics.HasError(), "invalid user_agent_extra should return an error")
}
