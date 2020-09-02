package provider

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/databrickslabs/databricks-terraform/common"
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
	errPrefix                string
	hasToken                 string
	hasHost                  string
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
	tests := []providerConfigTest{
		{
			errPrefix: "Authentication is not configured for provider",
		},
		{
			env: map[string]string{
				"DATABRICKS_HOST": "x",
			},
			errPrefix: "Authentication is not configured for provider",
		},
		{
			env: map[string]string{
				"DATABRICKS_TOKEN": "x",
			},
			errPrefix: "Host is empty, but is required by token",
		},
		{
			env: map[string]string{
				"DATABRICKS_HOST":  "x",
				"DATABRICKS_TOKEN": "x",
			},
			hasToken: "x",
			hasHost:  "https://x",
		},
		{
			host: "https://x",
			env: map[string]string{
				"DATABRICKS_TOKEN": "x",
			},
			hasToken: "x",
			hasHost:  "https://x",
		},
		{
			env: map[string]string{
				"DATABRICKS_USERNAME": "x",
				"DATABRICKS_PASSWORD": "x",
			},
			errPrefix: "Host is empty, but is required by basic_auth",
			hasToken:  "x",
			hasHost:   "https://x",
		},
		{
			env: map[string]string{
				"DATABRICKS_HOST":     "x",
				"DATABRICKS_USERNAME": "x",
				"DATABRICKS_PASSWORD": "x",
			},
			hasToken: "eDp4",
			hasHost:  "https://x",
		},
		{
			host: "y",
			env: map[string]string{
				"DATABRICKS_HOST":     "x",
				"DATABRICKS_USERNAME": "x",
				"DATABRICKS_PASSWORD": "x",
			},
			hasToken: "eDp4",
			hasHost:  "https://y",
		},
		{
			host:     "y",
			username: "x",
			env: map[string]string{
				"DATABRICKS_PASSWORD": "x",
			},
			hasToken: "eDp4",
			hasHost:  "https://y",
		},
		{
			host:     "y",
			username: "x",
			password: "x",
			hasToken: "eDp4",
			hasHost:  "https://y",
		},
		{
			env: map[string]string{
				"DATABRICKS_HOST":     "x",
				"DATABRICKS_TOKEN":    "x",
				"DATABRICKS_USERNAME": "x",
				"DATABRICKS_PASSWORD": "x",
			},
			errPrefix: "More than one authorization method configured: password and token",
		},
		{
			env: map[string]string{
				"CONFIG_FILE": "x",
			},
			errPrefix: "Authentication is not configured for provider",
		},
		{
			// loading with DEFAULT profile in databrickscfs
			env: map[string]string{
				"HOME": "../common/testdata",
			},
			hasHost:  "https://dbc-XXXXXXXX-YYYY.cloud.databricks.com/",
			hasToken: "PT0+IC9kZXYvdXJhbmRvbSA8PT0KYFZ",
		},
		{
			// loading with nohost profile in databrickscfs
			env: map[string]string{
				"HOME":                      "../common/testdata",
				"DATABRICKS_CONFIG_PROFILE": "nohost",
			},
			errPrefix: "config file ../common/testdata/.databrickscfg is corrupt: cannot find host in nohost profile",
		},
		{
			env: map[string]string{
				"DATABRICKS_TOKEN":          "x",
				"DATABRICKS_CONFIG_PROFILE": "nohost",
				"HOME":                      "../common/testdata",
			},
			errPrefix: "More than one authorization method configured: config profile and token",
		},
		{
			env: map[string]string{
				"DATABRICKS_USERNAME":       "x",
				"DATABRICKS_CONFIG_PROFILE": "nohost",
				"HOME":                      "../common/testdata",
			},
			errPrefix: "More than one authorization method configured: config profile and password",
		},
		{
			// this test will skip ensureWorkspaceUrl
			host:                     "x",
			azureWorkspaceResourceID: "/a/b/c",
			env: map[string]string{
				// // these may fail on windows. use docker container for testing.
				"PATH": "../common/testdata:/bin",
				"HOME": "../common/testdata",
			},
			hasHost:  "https://x",
			hasToken: "",
		},
		{
			azureWorkspaceResourceID: "/a/b/c",
			env: map[string]string{
				// these may fail on windows. use docker container for testing.
				"PATH": "../common/testdata:/bin",
				"FAIL": "yes",
				"HOME": "../common/testdata",
			},
			errPrefix: "Invoking Azure CLI failed with the following error: This is just a failing script.",
		},
		{
			// `az` not installed, which is expected for deployers on other clouds...
			azureWorkspaceResourceID: "/a/b/c",
			env: map[string]string{
				"PATH": "/bin",
				"HOME": "../common/testdata",
			},
			errPrefix: "Most likely Azure CLI is not installed.",
		},
		{
			azureWorkspaceResourceID: "/a/b/c",
			token:                    "x",
			env: map[string]string{
				// these may fail on windows. use docker container for testing.
				"PATH": "../common/testdata:/bin",
				"HOME": "../common/testdata",
			},
			errPrefix: "More than one authorization method configured: azure and token",
		},
		{
			// omit request to management endpoint to get workspace properties
			azureWorkspaceResourceID: "/a/b/c",
			host:                     "x",
			env: map[string]string{
				// these may fail on windows. use docker container for testing.
				"PATH": "../common/testdata:/bin",
				"HOME": "../common/testdata",
			},
			hasHost: "https://x",
		},
		{
			host:                     "x",
			azureWorkspaceResourceID: "/a/b/c",
			env: map[string]string{
				// these may fail on windows. use docker container for testing.
				"PATH":                "../common/testdata:/bin",
				"HOME":                "../common/testdata",
				"DATABRICKS_USERNAME": "x",
			},
			errPrefix: "More than one authorization method configured: azure and password",
		},
		{
			azureWorkspaceResourceID: "/a/b/c",
			azureClientID:            "x",
			azureClientSecret:        "y",
			azureTenantID:            "z",
			env: map[string]string{
				"HOME": "../common/testdata",
			},
			hasHost:  "",
			hasToken: "",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("config:%v env:%v", tt.rawConfig(), tt.env), func(t *testing.T) {
			c, err := configureProviderAndReturnClient(t, tt)
			if tt.errPrefix != "" {
				require.NotNilf(t, err, "Expected to have %s error", tt.errPrefix)
				require.True(t, strings.HasPrefix(err.Error(), tt.errPrefix), err)
				return
			}
			if err != nil {
				require.NoError(t, err)
				return
			}
			assert.Equal(t, tt.hasToken, c.Token)
			assert.Equal(t, tt.hasHost, c.Host)
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

func TestProvider_DurationToSecondsString(t *testing.T) {
	assert.Equal(t, durationToSecondsString(time.Hour), "3600")
}
