package tests

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw"
	"github.com/databricks/terraform-provider-databricks/internal/providers/sdkv2"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type providerFixture struct {
	host              string
	token             string
	username          string
	password          string
	configFile        string
	profile           string
	azureClientID     string
	azureClientSecret string
	azureTenantID     string
	azureResourceID   string
	authType          string
	env               map[string]string
	assertError       string
	assertAuth        string
	assertHost        string
	assertAzure       bool
}

func (pf providerFixture) rawConfig() map[string]string {
	rawConfig := map[string]string{}
	if pf.host != "" {
		rawConfig["host"] = pf.host
	}
	if pf.token != "" {
		rawConfig["token"] = pf.token
	}
	if pf.username != "" {
		rawConfig["username"] = pf.username
	}
	if pf.password != "" {
		rawConfig["password"] = pf.password
	}
	if pf.configFile != "" {
		rawConfig["config_file"] = pf.configFile
	}
	if pf.profile != "" {
		rawConfig["profile"] = pf.profile
	}
	if pf.azureClientID != "" {
		rawConfig["azure_client_id"] = pf.azureClientID
	}
	if pf.azureClientSecret != "" {
		rawConfig["azure_client_secret"] = pf.azureClientSecret
	}
	if pf.azureTenantID != "" {
		rawConfig["azure_tenant_id"] = pf.azureTenantID
	}
	if pf.azureResourceID != "" {
		rawConfig["azure_workspace_resource_id"] = pf.azureResourceID
	}
	if pf.authType != "" {
		rawConfig["auth_type"] = pf.authType
	}
	return rawConfig
}

func (pf providerFixture) rawConfigSDKv2() map[string]any {
	rawConfig := pf.rawConfig()
	rawConfigSDKv2 := map[string]any{}
	for k, v := range rawConfig {
		rawConfigSDKv2[k] = v
	}
	return rawConfigSDKv2
}

func (pf providerFixture) rawConfigPluginFramework() tftypes.Value {
	rawConfig := pf.rawConfig()

	rawConfigTypeMap := map[string]tftypes.Type{}
	for k := range rawConfig {
		rawConfigTypeMap[k] = tftypes.String
	}
	rawConfigType := tftypes.Object{
		AttributeTypes: rawConfigTypeMap,
	}

	rawConfigValueMap := map[string]tftypes.Value{}
	for k, v := range rawConfig {
		rawConfigValueMap[k] = tftypes.NewValue(tftypes.String, v)
	}
	rawConfigValue := tftypes.NewValue(rawConfigType, rawConfigValueMap)
	return rawConfigValue
}

func (pf providerFixture) applyWithSDKv2(t *testing.T) *common.DatabricksClient {
	c, err := pf.configureProviderAndReturnClient_SDKv2(t)
	return pf.applyAssertions(c, t, err)
}

func (pf providerFixture) applyWithPluginFramework(t *testing.T) *common.DatabricksClient {
	c, err := pf.configureProviderAndReturnClient_PluginFramework(t)
	return pf.applyAssertions(c, t, err)
}

func (pf providerFixture) applyAssertions(c *common.DatabricksClient, t *testing.T, err error) *common.DatabricksClient {
	if pf.assertError != "" {
		require.NotNilf(t, err, "Expected to have %s error", pf.assertError)
		require.True(t, strings.HasPrefix(err.Error(), pf.assertError),
			"Expected to have '%s' error, but got '%s'", pf.assertError, err)
		return nil
	}
	if err != nil {
		require.NoError(t, err)
		return nil
	}
	assert.Equal(t, pf.assertAzure, c.IsAzure())
	assert.Equal(t, pf.assertAuth, c.Config.AuthType)
	assert.Equal(t, pf.assertHost, c.Config.Host)
	return c
}

func (pf providerFixture) apply(t *testing.T) {
	t.Run("sdkv2", func(t *testing.T) { _ = pf.applyWithSDKv2(t) })
	t.Run("plugin framework", func(t *testing.T) { _ = pf.applyWithPluginFramework(t) })
}

func (pf providerFixture) configureProviderAndReturnClient_SDKv2(t *testing.T) (*common.DatabricksClient, error) {
	for k, v := range pf.env {
		t.Setenv(k, v)
	}
	p := sdkv2.DatabricksProvider()
	ctx := context.Background()
	diags := p.Configure(ctx, terraform.NewResourceConfigRaw(pf.rawConfigSDKv2()))
	if len(diags) > 0 {
		issues := []string{}
		for _, d := range diags {
			issues = append(issues, d.Summary)
		}
		return nil, errors.New(strings.Join(issues, ", "))
	}
	client := p.Meta().(*common.DatabricksClient)
	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		return nil, err
	}
	err = client.Config.Authenticate(r)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (pf providerFixture) configureProviderAndReturnClient_PluginFramework(t *testing.T) (*common.DatabricksClient, error) {
	for k, v := range pf.env {
		t.Setenv(k, v)
	}
	p := pluginfw.GetDatabricksProviderPluginFramework()
	ctx := context.Background()
	rawConfig := pf.rawConfigPluginFramework()
	var providerSchemaResponse provider.SchemaResponse
	p.Schema(ctx, provider.SchemaRequest{}, &providerSchemaResponse)
	configRequest := provider.ConfigureRequest{
		Config: tfsdk.Config{
			Raw:    rawConfig,
			Schema: providerSchemaResponse.Schema,
		},
	}
	configResponse := &provider.ConfigureResponse{}
	p.Configure(ctx, configRequest, configResponse)
	diags := configResponse.Diagnostics
	if len(diags) > 0 {
		issues := []string{}
		for _, d := range diags {
			issues = append(issues, d.Summary())
		}
		return nil, errors.New(strings.Join(issues, ", "))
	}
	client := configResponse.ResourceData.(*common.DatabricksClient)
	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		return nil, err
	}
	err = client.Config.Authenticate(r)
	if err != nil {
		return nil, err
	}
	return client, nil
}
