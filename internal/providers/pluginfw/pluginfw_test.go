package pluginfw

import (
	"bytes"
	"context"
	"log"
	"strings"
	"sync"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
)

// resetFallbackWarningState rewires the package logger to a buffer and resets
// the once-per-name dedupe map, returning the buffer and a cleanup func.
func resetFallbackWarningState(t *testing.T) *bytes.Buffer {
	t.Helper()
	var buf bytes.Buffer
	origWriter := log.Writer()
	origFlags := log.Flags()
	log.SetOutput(&buf)
	log.SetFlags(0)
	warnedFallbackNames = sync.Map{}
	t.Cleanup(func() {
		log.SetOutput(origWriter)
		log.SetFlags(origFlags)
		warnedFallbackNames = sync.Map{}
	})
	return &buf
}

func TestEmitSdkV2FallbackWarning_LogsOncePerName(t *testing.T) {
	buf := resetFallbackWarningState(t)

	emitSdkV2FallbackWarning("databricks_library", "resource")
	emitSdkV2FallbackWarning("databricks_library", "resource") // duplicate, suppressed
	emitSdkV2FallbackWarning("databricks_shares", "data source")

	out := buf.String()
	assert.Equal(t, 1, strings.Count(out, `"databricks_library"`), "library warning should fire exactly once")
	assert.Contains(t, out, `"databricks_shares"`)
	assert.Contains(t, out, "[WARN]")
	assert.Contains(t, out, "deprecated SDKv2 implementation")
	assert.Contains(t, out, "next major release")
}

func TestGetPluginFrameworkResources_EnvVarFallbackEmitsWarning(t *testing.T) {
	buf := resetFallbackWarningState(t)
	t.Setenv("USE_SDK_V2_RESOURCES", "databricks_library")

	got := getPluginFrameworkResourcesToRegister(nil)

	for _, fn := range got {
		assert.NotEqual(t, "databricks_library", getResourceName(fn),
			"databricks_library must be excluded from PF when fallback is requested")
	}
	assert.Contains(t, buf.String(), `"databricks_library"`)
	assert.Contains(t, buf.String(), "resource")
}

func TestGetPluginFrameworkDataSources_FallbackOptionEmitsWarning(t *testing.T) {
	buf := resetFallbackWarningState(t)

	got := getPluginFrameworkDataSourcesToRegister([]string{"databricks_volumes"})

	for _, fn := range got {
		assert.NotEqual(t, "databricks_volumes", getDataSourceName(fn),
			"databricks_volumes must be excluded from PF when fallback is requested")
	}
	assert.Contains(t, buf.String(), `"databricks_volumes"`)
	assert.Contains(t, buf.String(), "data source")
}

func TestGetPluginFrameworkResources_NoFallbackNoWarning(t *testing.T) {
	buf := resetFallbackWarningState(t)

	_ = getPluginFrameworkResourcesToRegister(nil)

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
