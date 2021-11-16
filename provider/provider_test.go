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

type providerFixture struct {
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

func (tt providerFixture) rawConfig() map[string]interface{} {
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

func (tc providerFixture) apply(t *testing.T) {
	c, err := configureProviderAndReturnClient(t, tc)
	if tc.assertError != "" {
		require.NotNilf(t, err, "Expected to have %s error", tc.assertError)
		require.True(t, strings.HasPrefix(err.Error(), tc.assertError),
			"Expected to have '%s' error, but got '%s'", tc.assertError, err)
		return
	}
	if err != nil {
		require.NoError(t, err)
		return
	}
	assert.Equal(t, tc.assertAzure, c.IsAzure())
	assert.Equal(t, tc.assertToken, c.Token)
	assert.Equal(t, tc.assertHost, c.Host)
}

func TestConfig_NoParams(t *testing.T) {
	providerFixture{
		assertError: "authentication is not configured for provider",
	}.apply(t)
}

func TestConfig_HostEnv(t *testing.T) {
	providerFixture{
		env: map[string]string{
			"DATABRICKS_HOST": "x",
		},
		assertError: "authentication is not configured for provider",
	}.apply(t)
}

func TestConfig_TokenEnv(t *testing.T) {
	providerFixture{
		env: map[string]string{
			"DATABRICKS_TOKEN": "x",
		},
		assertError: "cannot configure direct auth: host is empty, but is required by token. Environment variables used: DATABRICKS_TOKEN",
	}.apply(t)
}

func TestConfig_HostTokenEnv(t *testing.T) {
	providerFixture{
		env: map[string]string{
			"DATABRICKS_HOST":  "x",
			"DATABRICKS_TOKEN": "x",
		},
		assertToken: "x",
		assertHost:  "https://x",
	}.apply(t)
}

func TestConfig_HostParamTokenEnv(t *testing.T) {
	providerFixture{
		host: "https://x",
		env: map[string]string{
			"DATABRICKS_TOKEN": "x",
		},
		assertToken: "x",
		assertHost:  "https://x",
	}.apply(t)
}

func TestConfig_UserPasswordEnv(t *testing.T) {
	providerFixture{
		env: map[string]string{
			"DATABRICKS_USERNAME": "x",
			"DATABRICKS_PASSWORD": "x",
		},
		assertError: "cannot configure direct auth: host is empty, but is required by basic_auth. " +
			"Environment variables used: DATABRICKS_USERNAME, DATABRICKS_PASSWORD. " +
			"Please check https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs#authentication for details",
		assertToken: "x",
		assertHost:  "https://x",
	}.apply(t)
}

func TestConfig_BasicAuth(t *testing.T) {
	providerFixture{
		env: map[string]string{
			"DATABRICKS_HOST":     "x",
			"DATABRICKS_USERNAME": "x",
			"DATABRICKS_PASSWORD": "x",
		},
		assertToken: "eDp4",
		assertHost:  "https://x",
	}.apply(t)
}

func TestConfig_AttributePrecedence(t *testing.T) {
	providerFixture{
		host: "y",
		env: map[string]string{
			"DATABRICKS_HOST":     "x",
			"DATABRICKS_USERNAME": "x",
			"DATABRICKS_PASSWORD": "x",
		},
		assertToken: "eDp4",
		assertHost:  "https://y",
	}.apply(t)
}

func TestConfig_BasicAuth_Mix(t *testing.T) {
	providerFixture{
		host:     "y",
		username: "x",
		env: map[string]string{
			"DATABRICKS_PASSWORD": "x",
		},
		assertToken: "eDp4",
		assertHost:  "https://y",
	}.apply(t)
}

func TestConfig_BasicAuth_Attrs(t *testing.T) {
	providerFixture{
		host:        "y",
		username:    "x",
		password:    "x",
		assertToken: "eDp4",
		assertHost:  "https://y",
	}.apply(t)
}

func TestConfig_AzurePAT(t *testing.T) {
	providerFixture{
		// Azure hostnames can support host+token auth, as usual
		host:        "https://adb-xxx.y.azuredatabricks.net/",
		token:       "y",
		assertAzure: true,
		assertHost:  "https://adb-xxx.y.azuredatabricks.net/",
		assertToken: "y",
	}.apply(t)
}

func TestConfig_ConflictingEnvs(t *testing.T) {
	providerFixture{
		env: map[string]string{
			"DATABRICKS_HOST":     "x",
			"DATABRICKS_TOKEN":    "x",
			"DATABRICKS_USERNAME": "x",
			"DATABRICKS_PASSWORD": "x",
		},
		assertError: "More than one authorization method configured: password and token",
	}.apply(t)
}

func TestConfig_ConfigFile(t *testing.T) {
	providerFixture{
		env: map[string]string{
			"CONFIG_FILE": "x",
		},
		assertError: "authentication is not configured for provider",
	}.apply(t)
}

func TestConfig_PatFromDatabricksCfg(t *testing.T) {
	providerFixture{
		// loading with DEFAULT profile in databrickscfs
		env: map[string]string{
			"HOME": "../common/testdata",
		},
		assertHost:  "https://dbc-XXXXXXXX-YYYY.cloud.databricks.com/",
		assertToken: "PT0+IC9kZXYvdXJhbmRvbSA8PT0KYFZ",
	}.apply(t)
}

func TestConfig_PatFromDatabricksCfg_NohostProfile(t *testing.T) {
	providerFixture{
		// loading with nohost profile in databrickscfs
		env: map[string]string{
			"HOME":                      "../common/testdata",
			"DATABRICKS_CONFIG_PROFILE": "nohost",
		},
		assertError: "cannot configure Databricks CLI auth: config file ../common/testdata/.databrickscfg is corrupt: cannot find host in nohost profile",
	}.apply(t)
}

func TestConfig_ConfigProfileAndToken(t *testing.T) {
	providerFixture{
		env: map[string]string{
			"DATABRICKS_TOKEN":          "x",
			"DATABRICKS_CONFIG_PROFILE": "nohost",
			"HOME":                      "../common/testdata",
		},
		assertError: "More than one authorization method configured: config profile and token",
	}.apply(t)
}

func TestConfig_ConfigProfileAndPassword(t *testing.T) {
	providerFixture{
		env: map[string]string{
			"DATABRICKS_USERNAME":       "x",
			"DATABRICKS_CONFIG_PROFILE": "nohost",
			"HOME":                      "../common/testdata",
		},
		assertError: "More than one authorization method configured: config profile and password",
	}.apply(t)
}

var azResourceID = "/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c"

func TestConfig_AzureCliHost(t *testing.T) {
	providerFixture{
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
	}.apply(t)
}

func TestConfig_AzureCliHost_Fail(t *testing.T) {
	providerFixture{
		azureWorkspaceResourceID: azResourceID,
		env: map[string]string{
			// these may fail on windows. use docker container for testing.
			"PATH": "../common/testdata",
			"HOME": "../common/testdata",
			"FAIL": "yes",
		},
		assertError: "cannot configure Azure CLI auth: Invoking Azure CLI failed with the following error: This is just a failing script.",
	}.apply(t)
}

func TestConfig_AzureCliHost_AzNotInstalled(t *testing.T) {
	providerFixture{
		// `az` not installed, which is expected for deployers on other clouds...
		azureWorkspaceResourceID: azResourceID,
		env: map[string]string{
			"PATH": "whatever",
			"HOME": "../common/testdata",
		},
		assertError: "cannot configure Azure CLI auth: most likely Azure CLI is not installed.",
	}.apply(t)
}

func TestConfig_AzureCliHost_PatConflict(t *testing.T) {
	providerFixture{
		azureWorkspaceResourceID: azResourceID,
		token:                    "x",
		env: map[string]string{
			// these may fail on windows. use docker container for testing.
			"PATH": "../common/testdata",
			"HOME": "../common/testdata",
		},
		assertError: "More than one authorization method configured: azure and token",
	}.apply(t)
}

func TestConfig_AzureCliHostAndResourceID(t *testing.T) {
	providerFixture{
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
	}.apply(t)
}

func TestConfig_AzureAndPasswordConflict(t *testing.T) {
	providerFixture{
		host:                     "x",
		azureWorkspaceResourceID: azResourceID,
		env: map[string]string{
			// these may fail on windows. use docker container for testing.
			"PATH":                "../common/testdata",
			"HOME":                "../common/testdata",
			"DATABRICKS_USERNAME": "x",
		},
		assertError: "More than one authorization method configured: azure and password",
	}.apply(t)
}

func TestConfig_AzureResourceIDViaSpn(t *testing.T) {
	providerFixture{
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
	}.apply(t)
}

func TestConfig_Bug294(t *testing.T) {
	providerFixture{
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
	}.apply(t)
}

func TestConfig_CorruptConfig(t *testing.T) {
	providerFixture{
		env: map[string]string{
			"HOME": "../common/testdata/corrupt",
		},
		assertError: "cannot configure Databricks CLI auth: ../common/testdata/corrupt/.databrickscfg has no DEFAULT profile configured",
	}.apply(t)
}

func configureProviderAndReturnClient(t *testing.T, tt providerFixture) (*common.DatabricksClient, error) {
	defer common.CleanupEnvironment()()
	for k, v := range tt.env {
		os.Setenv(k, v)
	}
	p := DatabricksProvider()
	ctx := context.Background()
	diags := p.Configure(ctx, terraform.NewResourceConfigRaw(tt.rawConfig()))
	if len(diags) > 0 {
		issues := []string{}
		for _, d := range diags {
			issues = append(issues, d.Summary)
		}
		return nil, fmt.Errorf(strings.Join(issues, ", "))
	}
	client := p.Meta().(*common.DatabricksClient)
	err := client.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func TestAllResourcesMustHaveImport(t *testing.T) {
	p := DatabricksProvider()
	for name, r := range p.ResourcesMap {
		if r.Importer == nil {
			assert.Fail(t, "Missing importer: %s", name)
		}
	}
}
