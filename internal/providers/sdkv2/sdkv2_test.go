package sdkv2

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/client"
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

func TestProviderSchema_HttpHeadersAndPathPrefix(t *testing.T) {
	p := DatabricksProvider()

	// Verify http_headers schema exists with correct configuration
	httpHeadersSchema, ok := p.Schema["http_headers"]
	assert.True(t, ok, "http_headers schema should exist")
	assert.Equal(t, schema.TypeMap, httpHeadersSchema.Type)
	assert.True(t, httpHeadersSchema.Optional)
	assert.True(t, httpHeadersSchema.Sensitive)
	assert.Equal(t, "Custom HTTP headers to add to all API requests. Useful for HTTP proxies that require custom authentication headers.", httpHeadersSchema.Description)

	// Verify http_path_prefix schema exists with correct configuration
	httpPathPrefixSchema, ok := p.Schema["http_path_prefix"]
	assert.True(t, ok, "http_path_prefix schema should exist")
	assert.Equal(t, schema.TypeString, httpPathPrefixSchema.Type)
	assert.True(t, httpPathPrefixSchema.Optional)
	assert.False(t, httpPathPrefixSchema.Sensitive)
	assert.Equal(t, "Path prefix to prepend to all API request URLs. Useful for HTTP proxies that use path-based routing.", httpPathPrefixSchema.Description)
}

func TestConfigureDatabricksClient_HttpConfig(t *testing.T) {
	testCases := []struct {
		name                   string
		config                 map[string]interface{}
		expectHTTPTransportSet bool
	}{
		{
			name:                   "No HTTP config by default",
			config:                 map[string]interface{}{},
			expectHTTPTransportSet: false,
		},
		{
			name: "HTTP headers set configures transport",
			config: map[string]interface{}{
				"http_headers": map[string]interface{}{
					"X-Custom-Header": "value",
				},
			},
			expectHTTPTransportSet: true,
		},
		{
			name: "HTTP path prefix set configures transport",
			config: map[string]interface{}{
				"http_path_prefix": "/proxy/env",
			},
			expectHTTPTransportSet: true,
		},
		{
			name: "Both HTTP headers and path prefix set",
			config: map[string]interface{}{
				"http_headers": map[string]interface{}{
					"X-Auth": "token",
				},
				"http_path_prefix": "/proxy",
			},
			expectHTTPTransportSet: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := DatabricksProvider()
			rd := schema.TestResourceDataRaw(t, p.Schema, tc.config)

			result, diags := ConfigureDatabricksClient(context.Background(), rd, nil)
			assert.False(t, diags.HasError(), "ConfigureDatabricksClient should not return errors")

			dc, ok := result.(*common.DatabricksClient)
			assert.True(t, ok, "result should be a DatabricksClient")

			if tc.expectHTTPTransportSet {
				assert.NotNil(t, dc.Config.HTTPTransport, "HTTPTransport should be set")
			} else {
				assert.Nil(t, dc.Config.HTTPTransport, "HTTPTransport should not be set")
			}
		})
	}
}

func TestHTTPConfig_Struct(t *testing.T) {
	// Test that HTTPConfig struct can be created and used
	cfg := &client.HTTPConfig{
		Headers: map[string]string{
			"X-Header-1": "value1",
			"X-Header-2": "value2",
		},
		PathPrefix: "/proxy/path",
	}

	assert.Len(t, cfg.Headers, 2)
	assert.Equal(t, "value1", cfg.Headers["X-Header-1"])
	assert.Equal(t, "/proxy/path", cfg.PathPrefix)
}
