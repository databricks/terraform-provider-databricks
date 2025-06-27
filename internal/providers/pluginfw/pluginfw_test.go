package pluginfw

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
)

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
