package tests

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/sdkv2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfig_NoParams(t *testing.T) {
	if f, err := os.Stat("~/.databrickscfg"); err == nil && !f.IsDir() {
		// the provider should fail to configure if no configuration options are available,
		// either through environment or config file. However, many developers have a
		// ~/.databrickscfg file, so we skip this test if that file exists.
		t.Skip("Skipping no-params auth test because ~/.databrickscfg exists")
	}
	providerFixture{
		assertError: common.NoAuth,
	}.apply(t)
}

func TestConfig_HostEnv(t *testing.T) {
	providerFixture{
		env: map[string]string{
			"DATABRICKS_HOST": "x",
		},
		assertError: common.NoAuth,
	}.apply(t)
}

func TestConfig_TokenEnv(t *testing.T) {
	providerFixture{
		env: map[string]string{
			"DATABRICKS_TOKEN": "x",
		},
		assertError: common.NoAuth + ". Config: token=***. Env: DATABRICKS_TOKEN",
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
		assertError: common.NoAuth +
			". Config: username=x, password=***. Env: DATABRICKS_USERNAME, DATABRICKS_PASSWORD",
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
			"DATABRICKS_CONFIG_FILE": "x",
		},
		assertError: common.NoAuth,
	}.apply(t)
}

func TestConfig_PatFromDatabricksCfg(t *testing.T) {
	providerFixture{
		// loading with DEFAULT profile in databrickscfs
		env: map[string]string{
			"HOME": getTestDataPath(),
		},
		assertHost: "https://dbc-XXXXXXXX-YYYY.cloud.databricks.com",
		assertAuth: "pat",
	}.apply(t)
}

func TestConfig_PatFromDatabricksCfg_NohostProfile(t *testing.T) {
	providerFixture{
		// loading with nohost profile in databrickscfs
		env: map[string]string{
			"HOME":                      getTestDataPath(),
			"DATABRICKS_CONFIG_PROFILE": "nohost",
		},
		assertError: common.NoAuth +
			". Config: token=***, profile=nohost. Env: DATABRICKS_CONFIG_PROFILE",
	}.apply(t)
}

func TestConfig_ConfigProfileAndToken(t *testing.T) {
	providerFixture{
		env: map[string]string{
			"DATABRICKS_TOKEN":          "x",
			"DATABRICKS_CONFIG_PROFILE": "nohost",
			"HOME":                      getTestDataPath(),
		},
		assertError: common.NoAuth +
			". Config: token=***, profile=nohost. Env: DATABRICKS_TOKEN, DATABRICKS_CONFIG_PROFILE",
	}.apply(t)
}

func TestConfig_ConfigProfileAndPassword(t *testing.T) {
	providerFixture{
		env: map[string]string{
			"DATABRICKS_USERNAME":       "x",
			"DATABRICKS_CONFIG_PROFILE": "nohost",
			"HOME":                      getTestDataPath(),
		},
		assertError: "validate: more than one authorization method configured: basic and pat. " +
			"Config: token=***, username=x, profile=nohost. Env: DATABRICKS_USERNAME, DATABRICKS_CONFIG_PROFILE",
	}.apply(t)
}

var azResourceID = "/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c"

func TestConfig_AzureCliHost(t *testing.T) {
	p, _ := filepath.Abs(getTestDataPath())
	providerFixture{
		// this test will skip ensureWorkspaceUrl
		host:            "x",
		azureResourceID: azResourceID,
		env: map[string]string{
			// // these may fail on windows. use docker container for testing.
			"PATH": p,
			"HOME": p,
		},
		assertAzure:   true,
		assertHost:    "https://x",
		assertAuth:    "azure-cli",
		azureTenantID: "tenant-id",
	}.apply(t)
}

func TestConfig_AzureCliHost_Fail(t *testing.T) {
	p, _ := filepath.Abs(getTestDataPath())
	providerFixture{
		azureResourceID: azResourceID,
		env: map[string]string{
			// these may fail on windows. use docker container for testing.
			"PATH": p,
			"HOME": p,
			"FAIL": "yes",
		},
		assertError: "default auth: azure-cli: cannot get account info",
	}.apply(t)
}

func TestConfig_AzureCliHost_AzNotInstalled(t *testing.T) {
	providerFixture{
		// `az` not installed, which is expected for deployers on other clouds...
		azureResourceID: azResourceID,
		env: map[string]string{
			"PATH": "whatever",
			"HOME": "../../common/testdata",
		},
		assertError: common.NoAuth,
	}.apply(t)
}

func TestConfig_AzureCliHost_PatConflict(t *testing.T) {
	p, _ := filepath.Abs(getTestDataPath())
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
	p, _ := filepath.Abs(getTestDataPath())
	providerFixture{
		// omit request to management endpoint to get workspace properties
		azureResourceID: azResourceID,
		host:            "x",
		env: map[string]string{
			// these may fail on windows. use docker container for testing.
			"PATH": p,
			"HOME": p,
		},
		assertAzure:   true,
		azureTenantID: "tenant-id",
		assertHost:    "https://x",
		assertAuth:    "azure-cli",
	}.apply(t)
}

func TestConfig_AzureAndPasswordConflict(t *testing.T) {
	p, _ := filepath.Abs(getTestDataPath())
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
			"HOME": getTestDataPath() + "/corrupt",
		},
		assertError: common.NoAuth,
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
	} else if r.RequestURI == "/api/2.1/clusters/get?cluster_id=123" {
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

	testFixture := providerFixture{
		env: map[string]string{
			"DATABRICKS_HOST":          ts.URL,
			"DATABRICKS_CLIENT_ID":     "x",
			"DATABRICKS_CLIENT_SECRET": "y",
		},
		assertAuth: "oauth-m2m",
		assertHost: ts.URL,
	}

	client := testFixture.applyWithSDKv2(t)
	testOAuthFetchesToken(t, client)

	client = testFixture.applyWithPluginFramework(t)
	testOAuthFetchesToken(t, client)

}

func testOAuthFetchesToken(t *testing.T, c *common.DatabricksClient) {
	ws, err := c.WorkspaceClient()
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

type parseUserAgentTestCase struct {
	name string
	env  string
	err  error
	out  []sdkv2.UserAgentExtra
}

func Test_ParseUserAgentExtra(t *testing.T) {
	testCases := []parseUserAgentTestCase{
		{
			name: "single product",
			env:  "databricks-cli/0.1.2",
			err:  nil,
			out: []sdkv2.UserAgentExtra{
				{Key: "databricks-cli", Value: "0.1.2"},
			},
		},
		{
			name: "multiple products",
			env:  "databricks-cli/0.1.2 custom-thing/0.0.1",
			err:  nil,
			out: []sdkv2.UserAgentExtra{
				{Key: "databricks-cli", Value: "0.1.2"},
				{Key: "custom-thing", Value: "0.0.1"},
			},
		},
		{
			name: "multiple products with many separators",
			env:  "\ta/0.0.1\tb/0.0.2 \t c/0.0.3",
			err:  nil,
			out: []sdkv2.UserAgentExtra{
				{Key: "a", Value: "0.0.1"},
				{Key: "b", Value: "0.0.2"},
				{Key: "c", Value: "0.0.3"},
			},
		},
		{
			name: "empty string",
			env:  "",
			err:  nil,
			out:  []sdkv2.UserAgentExtra{},
		},
		{
			name: "product with comment",
			env:  "a/0.0.1 (my comment)",
			err:  fmt.Errorf("product string must follow RFC 9110: (my"),
			out:  nil,
		},
		{
			name: "no version",
			env:  "cli",
			err:  fmt.Errorf("product string must include version: cli"),
			out:  nil,
		},
		{
			name: "invalid format",
			env:  "no/no/no",
			err:  fmt.Errorf("product string must follow RFC 9110: no/no/no"),
			out:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out, err := sdkv2.ParseUserAgentExtra(tc.env)

			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.out, out)
		})
	}
}
