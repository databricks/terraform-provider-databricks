package provider

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
	"time"

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

func (tc providerFixture) apply(t *testing.T) *common.DatabricksClient {
	c, err := configureProviderAndReturnClient(t, tc)
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
		assertAuth: "pat",
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
		assertError: "default auth: azure-cli: cannot get access token: This is just a failing script.",
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
		assertError: "default auth: cannot configure default credentials.",
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
		assertError: "validate: more than one authorization method configured: azure and pat.",
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
		assertError: "validate: more than one authorization method configured: azure and basic.",
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

func shortLivedOAuthHandler(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/oidc/.well-known/oauth-authorization-server" {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"authorization_endpoint": "http://%s/authorize", "token_endpoint": "http://%s/token"}`, r.Host, r.Host)
		return
	} else if r.RequestURI == "/token" {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"access_token": "x", "token_type": "Bearer", "expires_in": 3}`)
		return
	} else if r.RequestURI == "/api/2.0/clusters/get?cluster_id=123" {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"cluster_id": "123"}`)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

// Start a local webserver to serve tokens. Configure the CLI to fetch a token from this endpoint.
// Then, using a context with a short timeout, call an API. Wait for 2x the timeout, then call
// the API again with a new context with a short timeout. The second call should succeed without
// the context being cancelled.
func TestConfig_OAuthFetchesToken(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(shortLivedOAuthHandler))
	defer ts.Close()

	client := providerFixture{
		env: map[string]string{
			"DATABRICKS_HOST":          ts.URL,
			"DATABRICKS_CLIENT_ID":     "x",
			"DATABRICKS_CLIENT_SECRET": "y",
		},
		assertAuth: "oauth-m2m",
		assertHost: ts.URL,
	}.apply(t)

	ws, err := client.WorkspaceClient()
	require.NoError(t, err)
	bgCtx := context.Background()
	{
		ctx, cancel := context.WithCancel(bgCtx)
		_, err = ws.Clusters.GetByClusterId(ctx, "123")
		require.NoError(t, err)
		cancel()
	}
	log.Printf("[INFO] sleeping for 5 seconds to allow token to expire")
	time.Sleep(5 * time.Second)
	{
		_, err = ws.Clusters.GetByClusterId(bgCtx, "123")
		require.NoError(t, err)
	}
}

func configureProviderAndReturnClient(t *testing.T, tt providerFixture) (*common.DatabricksClient, error) {
	for k, v := range tt.env {
		t.Setenv(k, v)
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

	return client, nil
}
