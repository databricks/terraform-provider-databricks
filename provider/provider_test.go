package provider

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type providerConfigTest struct {
	host                     string
	token                    string
	username                 string
	password                 string
	configFile               string
	profile                  string
	azureClientID            string
	azureClientSecret        string
	azureTenantID            string
	azureResourceGroup       string
	azureWorkspaceName       string
	azureSubscriptionID      string
	azureWorkspaceResourceID string
	env                      map[string]string
	assertError              string
	assertToken              string
	assertHost               string
	assertAzure              bool
}

func (tt providerConfigTest) rawConfig() map[string]interface{} {
	rawConfig := map[string]interface{}{}
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
	if tt.azureResourceGroup != "" {
		rawConfig["azure_resource_group"] = tt.azureResourceGroup
	}
	if tt.azureWorkspaceName != "" {
		rawConfig["azure_workspace_name"] = tt.azureWorkspaceName
	}
	if tt.azureSubscriptionID != "" {
		rawConfig["azure_subscription_id"] = tt.azureSubscriptionID
	}
	if tt.azureWorkspaceResourceID != "" {
		rawConfig["azure_workspace_resource_id"] = tt.azureWorkspaceResourceID
	}
	return rawConfig
}

func TestProviderConfigurationOptions(t *testing.T) {
	azResourceID := "/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c"
	tests := []providerConfigTest{
		{
			assertError: "Authentication is not configured for provider",
		},
		{
			env: map[string]string{
				"DATABRICKS_HOST": "x",
			},
			assertError: "Authentication is not configured for provider",
		},
		{
			env: map[string]string{
				"DATABRICKS_TOKEN": "x",
			},
			assertError: "Host is empty, but is required by token",
		},
		{
			env: map[string]string{
				"DATABRICKS_HOST":  "x",
				"DATABRICKS_TOKEN": "x",
			},
			assertToken: "x",
			assertHost:  "https://x",
		},
		{
			host: "https://x",
			env: map[string]string{
				"DATABRICKS_TOKEN": "x",
			},
			assertToken: "x",
			assertHost:  "https://x",
		},
		{
			env: map[string]string{
				"DATABRICKS_USERNAME": "x",
				"DATABRICKS_PASSWORD": "x",
			},
			assertError: "Host is empty, but is required by basic_auth",
			assertToken: "x",
			assertHost:  "https://x",
		},
		{
			env: map[string]string{
				"DATABRICKS_HOST":     "x",
				"DATABRICKS_USERNAME": "x",
				"DATABRICKS_PASSWORD": "x",
			},
			assertToken: "eDp4",
			assertHost:  "https://x",
		},
		{
			host: "y",
			env: map[string]string{
				"DATABRICKS_HOST":     "x",
				"DATABRICKS_USERNAME": "x",
				"DATABRICKS_PASSWORD": "x",
			},
			assertToken: "eDp4",
			assertHost:  "https://y",
		},
		{
			host:     "y",
			username: "x",
			env: map[string]string{
				"DATABRICKS_PASSWORD": "x",
			},
			assertToken: "eDp4",
			assertHost:  "https://y",
		},
		{
			host:        "y",
			username:    "x",
			password:    "x",
			assertToken: "eDp4",
			assertHost:  "https://y",
		},
		{
			// Azure hostnames can support host+token auth, as usual
			host:        "https://adb-xxx.y.azuredatabricks.net/",
			token:       "y",
			assertAzure: true,
			assertHost:  "https://adb-xxx.y.azuredatabricks.net/",
			assertToken: "y",
		},
		{
			env: map[string]string{
				"DATABRICKS_HOST":     "x",
				"DATABRICKS_TOKEN":    "x",
				"DATABRICKS_USERNAME": "x",
				"DATABRICKS_PASSWORD": "x",
			},
			assertError: "More than one authorization method configured: password and token",
		},
		{
			env: map[string]string{
				"CONFIG_FILE": "x",
			},
			assertError: "Authentication is not configured for provider",
		},
		{
			// loading with DEFAULT profile in databrickscfs
			env: map[string]string{
				"HOME": "../common/testdata",
			},
			assertHost:  "https://dbc-XXXXXXXX-YYYY.cloud.databricks.com/",
			assertToken: "PT0+IC9kZXYvdXJhbmRvbSA8PT0KYFZ",
		},
		{
			// loading with nohost profile in databrickscfs
			env: map[string]string{
				"HOME":                      "../common/testdata",
				"DATABRICKS_CONFIG_PROFILE": "nohost",
			},
			assertError: "config file ../common/testdata/.databrickscfg is corrupt: cannot find host in nohost profile",
		},
		{
			env: map[string]string{
				"DATABRICKS_TOKEN":          "x",
				"DATABRICKS_CONFIG_PROFILE": "nohost",
				"HOME":                      "../common/testdata",
			},
			assertError: "More than one authorization method configured: config profile and token",
		},
		{
			env: map[string]string{
				"DATABRICKS_USERNAME":       "x",
				"DATABRICKS_CONFIG_PROFILE": "nohost",
				"HOME":                      "../common/testdata",
			},
			assertError: "More than one authorization method configured: config profile and password",
		},
		{
			// this test will skip ensureWorkspaceUrl
			host:                     "x",
			azureWorkspaceResourceID: azResourceID,
			env: map[string]string{
				// // these may fail on windows. use docker container for testing.
				"PATH": "../common/testdata",
				"HOME": "../common/testdata",
			},
			assertAzure: true,
			assertHost:  "https://x",
			assertToken: "",
		},
		{
			azureWorkspaceResourceID: azResourceID,
			env: map[string]string{
				// these may fail on windows. use docker container for testing.
				"PATH": "../common/testdata",
				"HOME": "../common/testdata",
				"FAIL": "yes",
			},
			assertError: "Invoking Azure CLI failed with the following error: This is just a failing script.",
		},
		{
			// `az` not installed, which is expected for deployers on other clouds...
			azureWorkspaceResourceID: azResourceID,
			env: map[string]string{
				"PATH": "whatever",
				"HOME": "../common/testdata",
			},
			assertError: "Most likely Azure CLI is not installed.",
		},
		{
			azureWorkspaceResourceID: azResourceID,
			token:                    "x",
			env: map[string]string{
				// these may fail on windows. use docker container for testing.
				"PATH": "../common/testdata",
				"HOME": "../common/testdata",
			},
			assertError: "More than one authorization method configured: azure and token",
		},
		{
			// omit request to management endpoint to get workspace properties
			azureWorkspaceResourceID: azResourceID,
			host:                     "x",
			env: map[string]string{
				// these may fail on windows. use docker container for testing.
				"PATH": "../common/testdata",
				"HOME": "../common/testdata",
			},
			assertAzure: true,
			assertHost:  "https://x",
		},
		{
			host:                     "x",
			azureWorkspaceResourceID: azResourceID,
			env: map[string]string{
				// these may fail on windows. use docker container for testing.
				"PATH":                "../common/testdata",
				"HOME":                "../common/testdata",
				"DATABRICKS_USERNAME": "x",
			},
			assertError: "More than one authorization method configured: azure and password",
		},
		{
			azureWorkspaceResourceID: azResourceID,
			azureClientID:            "x",
			azureClientSecret:        "y",
			azureTenantID:            "z",
			env: map[string]string{
				"PATH": "../common/testdata",
				"HOME": "../common/testdata",
			},
			assertAzure: true,
			assertHost:  "",
			assertToken: "",
		},
		{
			// https://github.com/databrickslabs/terraform-provider-databricks/issues/294
			azureResourceGroup: "b",
			azureWorkspaceName: "c",
			env: map[string]string{
				"ARM_CLIENT_ID":       "x",
				"ARM_CLIENT_SECRET":   "y",
				"ARM_SUBSCRIPTION_ID": "q",
				"ARM_TENANT_ID":       "z",
				"HOME":                "../common/testdata",
				"PATH":                "../common/testdata",
			},
			assertAzure: true,
			assertHost:  "",
			assertToken: "",
		},
		{
			env: map[string]string{
				"HOME": "../common/testdata/corrupt",
			},
			assertError: "../common/testdata/corrupt/.databrickscfg has no DEFAULT profile configured",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("config:%v env:%v", tt.rawConfig(), tt.env), func(t *testing.T) {
			c, err := configureProviderAndReturnClient(t, tt)
			if tt.assertError != "" {
				require.NotNilf(t, err, "Expected to have %s error", tt.assertError)
				require.True(t, strings.HasPrefix(err.Error(), tt.assertError), err)
				return
			}
			if err != nil {
				require.NoError(t, err)
				return
			}
			assert.Equal(t, tt.assertAzure, c.IsAzure())
			assert.Equal(t, tt.assertToken, c.Token)
			assert.Equal(t, tt.assertHost, c.Host)
		})
	}
}

func configureProviderAndReturnClient(t *testing.T, tt providerConfigTest) (*common.DatabricksClient, error) {
	defer common.CleanupEnvironment()()
	for k, v := range tt.env {
		os.Setenv(k, v)
	}
	p := DatabricksProvider()
	diags := p.Configure(context.Background(), terraform.NewResourceConfigRaw(tt.rawConfig()))
	if len(diags) > 0 {
		issues := []string{}
		for _, d := range diags {
			issues = append(issues, d.Summary)
		}
		return nil, fmt.Errorf(strings.Join(issues, ", "))
	}
	client := p.Meta().(*common.DatabricksClient)
	err := client.Authenticate()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func TestProvider_NoHostGivesError(t *testing.T) {
	var raw = make(map[string]interface{})
	raw["config_file"] = "testdata/.databrickscfg"
	raw["profile"] = "nohost"
	p := DatabricksProvider()
	err := p.Configure(context.Background(), terraform.NewResourceConfigRaw(raw))
	assert.NotNil(t, err)
}

func TestProvider_NoTokenGivesError(t *testing.T) {
	var raw = make(map[string]interface{})
	raw["config_file"] = "testdata/.databrickscfg"
	raw["profile"] = "notoken"
	p := DatabricksProvider()
	err := p.Configure(context.Background(), terraform.NewResourceConfigRaw(raw))
	assert.NotNil(t, err)
}

func TestProvider_InvalidProfileGivesError(t *testing.T) {
	var raw = make(map[string]interface{})
	raw["config_file"] = "testdata/.databrickscfg"
	raw["profile"] = "invalidhost"
	p := DatabricksProvider()
	err := p.Configure(context.Background(), terraform.NewResourceConfigRaw(raw))
	assert.NotNil(t, err)
}

func TestAllResourcesMustHaveImport(t *testing.T) {
	t.Skip("databricks_mws_* are currently not importable")
	p := DatabricksProvider()
	for name, r := range p.ResourcesMap {
		if r.Importer == nil {
			t.Logf("Missing importer: %s", name)
		}
	}
}
