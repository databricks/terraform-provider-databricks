package provider

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
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

func (tt providerFixture) rawConfig() map[string]string {
	rawConfig := map[string]string{}
	if tt.host != "" {
		rawConfig["host"] = tt.host
	}
	if tt.token != "" {
		rawConfig["token"] = tt.token
	}
	if tt.username != "" {
		rawConfig["username"] = tt.username
	}
	if tt.password != "" {
		rawConfig["password"] = tt.password
	}
	if tt.configFile != "" {
		rawConfig["config_file"] = tt.configFile
	}
	if tt.profile != "" {
		rawConfig["profile"] = tt.profile
	}
	if tt.azureClientID != "" {
		rawConfig["azure_client_id"] = tt.azureClientID
	}
	if tt.azureClientSecret != "" {
		rawConfig["azure_client_secret"] = tt.azureClientSecret
	}
	if tt.azureTenantID != "" {
		rawConfig["azure_tenant_id"] = tt.azureTenantID
	}
	if tt.azureResourceID != "" {
		rawConfig["azure_workspace_resource_id"] = tt.azureResourceID
	}
	if tt.authType != "" {
		rawConfig["auth_type"] = tt.authType
	}
	return rawConfig
}

func (tt providerFixture) rawConfigSDKv2() map[string]any {
	rawConfig := tt.rawConfig()
	rawConfigSDKv2 := map[string]any{}
	for k, v := range rawConfig {
		rawConfigSDKv2[k] = v
	}
	return rawConfigSDKv2
}

func (tt providerFixture) rawConfigPluginFramework() tftypes.Value {
	rawConfig := tt.rawConfig()

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

func (tc providerFixture) applyWithSDKv2(t *testing.T) *common.DatabricksClient {
	c, err := configureProviderAndReturnClient_SDKv2(t, tc)
	return tc.applyAssertions(c, t, err)
}

func (tc providerFixture) applyWithPluginFramework(t *testing.T) *common.DatabricksClient {
	c, err := configureProviderAndReturnClient_PluginFramework(t, tc)
	return tc.applyAssertions(c, t, err)
}

func (tc providerFixture) applyAssertions(c *common.DatabricksClient, t *testing.T, err error) *common.DatabricksClient {
	if tc.assertError != "" {
		require.NotNilf(t, err, "Expected to have %s error", tc.assertError)
		require.True(t, strings.HasPrefix(err.Error(), tc.assertError),
			"Expected to have '%s' error, but got '%s'", tc.assertError, err)
		return nil
	}
	if err != nil {
		require.NoError(t, err)
		return nil
	}
	assert.Equal(t, tc.assertAzure, c.IsAzure())
	assert.Equal(t, tc.assertAuth, c.Config.AuthType)
	assert.Equal(t, tc.assertHost, c.Config.Host)
	return c
}

func (tc providerFixture) apply(t *testing.T) {
	_ = tc.applyWithSDKv2(t)
	_ = tc.applyWithPluginFramework(t)
}

func configureProviderAndReturnClient_SDKv2(t *testing.T, tt providerFixture) (*common.DatabricksClient, error) {
	for k, v := range tt.env {
		t.Setenv(k, v)
	}
	p := DatabricksProvider()
	ctx := context.Background()
	diags := p.Configure(ctx, terraform.NewResourceConfigRaw(tt.rawConfigSDKv2()))
	if len(diags) > 0 {
		issues := []string{}
		for _, d := range diags {
			issues = append(issues, d.Summary)
		}
		return nil, fmt.Errorf(strings.Join(issues, ", "))
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

func configureProviderAndReturnClient_PluginFramework(t *testing.T, tt providerFixture) (*common.DatabricksClient, error) {
	for k, v := range tt.env {
		t.Setenv(k, v)
	}
	p := GetDatabricksProviderPluginFramework()
	ctx := context.Background()
	rawConfig := tt.rawConfigPluginFramework()
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
		return nil, fmt.Errorf(strings.Join(issues, ", "))
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
