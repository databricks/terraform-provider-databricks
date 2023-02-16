package provider

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
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

func (tt providerFixture) rawConfig() map[string]any {
	rawConfig := map[string]any{}
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
	assert.Equal(t, tc.assertAuth, c.Config.AuthType)
	assert.Equal(t, tc.assertHost, c.Config.Host)
}

func TestConfig_NoParams(t *testing.T) {
	providerFixture{
		assertError: "default auth: cannot configure default credentials",
	}.apply(t)
}

func TestConfig_HostEnv(t *testing.T) {
	providerFixture{
		env: map[string]string{
			"DATABRICKS_HOST": "x",
		},
		assertError: "default auth: cannot configure default credentials",
	}.apply(t)
}

func TestConfig_TokenEnv(t *testing.T) {
	providerFixture{
		env: map[string]string{
			"DATABRICKS_TOKEN": "x",
		},
		assertError: "default auth: cannot configure default credentials. Config: token=***. Env: DATABRICKS_TOKEN",
	}.apply(t)
}

func TestConfig_HostTokenEnv(t *testing.T) {
	providerFixture{
		env: map[string]string{
			"DATABRICKS_HOST":  "x",
			"DATABRICKS_TOKEN": "x",
		},
		assertAuth: "pat",
		assertHost: "https://x",
	}.apply(t)
}

func TestConfig_HostParamTokenEnv(t *testing.T) {
	providerFixture{
		host: "https://x",
		env: map[string]string{
			"DATABRICKS_TOKEN": "x",
		},
		assertAuth: "pat",
		assertHost: "https://x",
	}.apply(t)
}

func TestConfig_UserPasswordEnv(t *testing.T) {
	providerFixture{
		env: map[string]string{
			"DATABRICKS_USERNAME": "x",
			"DATABRICKS_PASSWORD": "x",
		},
		assertError: "default auth: cannot configure default credentials. " +
			"Config: username=x, password=***. Env: DATABRICKS_USERNAME, DATABRICKS_PASSWORD",
		assertHost: "https://x",
	}.apply(t)
}

func TestConfig_BasicAuth(t *testing.T) {
	providerFixture{
		env: map[string]string{
			"DATABRICKS_HOST":     "x",
			"DATABRICKS_USERNAME": "x",
			"DATABRICKS_PASSWORD": "x",
		},
		assertAuth: "basic",
		assertHost: "https://x",
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
		assertAuth: "basic",
		assertHost: "https://y",
	}.apply(t)
}

func TestConfig_BasicAuth_Mix(t *testing.T) {
	providerFixture{
		host:     "y",
		username: "x",
		env: map[string]string{
			"DATABRICKS_PASSWORD": "x",
		},
		assertAuth: "basic",
		assertHost: "https://y",
	}.apply(t)
}

func TestConfig_BasicAuth_Attrs(t *testing.T) {
	providerFixture{
		host:       "y",
		username:   "x",
		password:   "x",
		assertAuth: "basic",
		assertHost: "https://y",
	}.apply(t)
}

func TestConfig_AzurePAT(t *testing.T) {
	providerFixture{
		// Azure hostnames can support host+token auth, as usual
		host:        "https://adb-xxx.y.azuredatabricks.net/",
		token:       "y",
		assertAzure: true,
		assertHost:  "https://adb-xxx.y.azuredatabricks.net",
		assertAuth:  "pat",
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
		assertError: "validate: more than one authorization method configured: basic and pat. " +
			"Config: host=x, token=***, username=x, password=***. " +
			"Env: DATABRICKS_HOST, DATABRICKS_TOKEN, DATABRICKS_USERNAME, DATABRICKS_PASSWORD",
	}.apply(t)
}

func TestConfig_ConflictingEnvs_AuthType(t *testing.T) {
	providerFixture{
		env: map[string]string{
			"DATABRICKS_HOST":     "x",
			"DATABRICKS_TOKEN":    "x",
			"DATABRICKS_USERNAME": "x",
			"DATABRICKS_PASSWORD": "x",
		},
		authType:   "basic",
		assertAuth: "basic",
		assertHost: "https://x",
	}.apply(t)
}

func TestConfig_ConfigFile(t *testing.T) {
	providerFixture{
		env: map[string]string{
			"CONFIG_FILE": "x",
		},
		assertError: "default auth: cannot configure default credentials",
	}.apply(t)
}

func TestConfig_PatFromDatabricksCfg(t *testing.T) {
	providerFixture{
		// loading with DEFAULT profile in databrickscfs
		env: map[string]string{
			"HOME": "../common/testdata",
		},
		assertHost: "https://dbc-XXXXXXXX-YYYY.cloud.databricks.com",
		assertAuth: "databricks-cli",
	}.apply(t)
}

func TestConfig_PatFromDatabricksCfg_NohostProfile(t *testing.T) {
	providerFixture{
		// loading with nohost profile in databrickscfs
		env: map[string]string{
			"HOME":                      "../common/testdata",
			"DATABRICKS_CONFIG_PROFILE": "nohost",
		},
		assertError: "default auth: cannot configure default credentials. " +
			"Config: token=***, profile=nohost. Env: DATABRICKS_CONFIG_PROFILE",
	}.apply(t)
}

func TestConfig_ConfigProfileAndToken(t *testing.T) {
	providerFixture{
		env: map[string]string{
			"DATABRICKS_TOKEN":          "x",
			"DATABRICKS_CONFIG_PROFILE": "nohost",
			"HOME":                      "../common/testdata",
		},
		assertError: "default auth: cannot configure default credentials. " +
			"Config: token=***, profile=nohost. Env: DATABRICKS_TOKEN, DATABRICKS_CONFIG_PROFILE",
	}.apply(t)
}

func TestConfig_ConfigProfileAndPassword(t *testing.T) {
	providerFixture{
		env: map[string]string{
			"DATABRICKS_USERNAME":       "x",
			"DATABRICKS_CONFIG_PROFILE": "nohost",
			"HOME":                      "../common/testdata",
		},
		assertError: "validate: more than one authorization method configured: basic and pat. " +
			"Config: token=***, username=x, profile=nohost. Env: DATABRICKS_USERNAME, DATABRICKS_CONFIG_PROFILE",
	}.apply(t)
}

var azResourceID = "/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c"

func TestConfig_AzureCliHost(t *testing.T) {
	p, _ := filepath.Abs("../common/testdata")
	providerFixture{
		// this test will skip ensureWorkspaceUrl
		host:            "x",
		azureResourceID: azResourceID,
		env: map[string]string{
			// // these may fail on windows. use docker container for testing.
			"PATH": p,
			"HOME": p,
		},
		assertAzure: true,
		assertHost:  "https://x",
		assertAuth:  "azure-cli",
	}.apply(t)
}

func TestConfig_AzureCliHost_Fail(t *testing.T) {
	p, _ := filepath.Abs("../common/testdata")
	providerFixture{
		azureResourceID: azResourceID,
		env: map[string]string{
			// these may fail on windows. use docker container for testing.
			"PATH": p,
			"HOME": p,
			"FAIL": "yes",
		},
		assertError: "cannot configure azure-cli auth: Invoking Azure CLI " +
			"failed with the following error: This is just a failing script.",
	}.apply(t)
}

func TestConfig_AzureCliHost_AzNotInstalled(t *testing.T) {
	providerFixture{
		// `az` not installed, which is expected for deployers on other clouds...
		azureResourceID: azResourceID,
		env: map[string]string{
			"PATH": "whatever",
			"HOME": "../common/testdata",
		},
		assertError: "cannot configure azure-cli auth: most likely Azure CLI is not installed.",
	}.apply(t)
}

func TestConfig_AzureCliHost_PatConflict(t *testing.T) {
	p, _ := filepath.Abs("../common/testdata")
	providerFixture{
		azureResourceID: azResourceID,
		token:           "x",
		env: map[string]string{
			// these may fail on windows. use docker container for testing.
			"PATH": p,
			"HOME": p,
		},
		assertError: "validate: more than one authorization method configured: azure and pat. " +
			"Config: host=https://dbc-XXXXXXXX-YYYY.cloud.databricks.com/, token=***, " +
			"azure_workspace_resource_id=/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c",
	}.apply(t)
}

func TestConfig_AzureCliHostAndResourceID(t *testing.T) {
	p, _ := filepath.Abs("../common/testdata")
	providerFixture{
		// omit request to management endpoint to get workspace properties
		azureResourceID: azResourceID,
		host:            "x",
		env: map[string]string{
			// these may fail on windows. use docker container for testing.
			"PATH": p,
			"HOME": p,
		},
		assertAzure: true,
		assertHost:  "https://x",
		assertAuth:  "azure-cli",
	}.apply(t)
}

func TestConfig_AzureAndPasswordConflict(t *testing.T) {
	p, _ := filepath.Abs("../common/testdata")
	providerFixture{
		host:            "x",
		azureResourceID: azResourceID,
		env: map[string]string{
			// these may fail on windows. use docker container for testing.
			"PATH":                p,
			"HOME":                p,
			"DATABRICKS_USERNAME": "x",
		},
		assertError: "validate: more than one authorization method configured: azure and basic and pat. " + 
			"Config: host=x, token=***, username=x, " +
			"azure_workspace_resource_id=/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c. " +
			"Env: DATABRICKS_USERNAME",
	}.apply(t)
}

func TestConfig_CorruptConfig(t *testing.T) {
	providerFixture{
		env: map[string]string{
			"HOME": "../common/testdata/corrupt",
		},
		assertError: "default auth: cannot configure default credentials",
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
	req, _ := http.NewRequest("GET", client.Config.Host, nil)
	err := client.Config.Authenticate(req)
	if err != nil {
		return nil, err
	}

	return client, nil
}
