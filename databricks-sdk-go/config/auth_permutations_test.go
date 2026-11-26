package config

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"testing"

	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/databricks/databricks-sdk-go/internal/env"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type configFixture struct {
	Name              string            `json:"name"`
	Host              string            `json:"host,omitempty"`
	Token             string            `json:"token,omitempty"`
	Username          string            `json:"username,omitempty"`
	Password          string            `json:"password,omitempty"`
	ConfigFile        string            `json:"config_file,omitempty"`
	Profile           string            `json:"profile,omitempty"`
	AzureClientID     string            `json:"azure_client_id,omitempty"`
	AzureClientSecret string            `json:"azure_client_secret,omitempty"`
	AzureTenantID     string            `json:"azure_tenant_id,omitempty"`
	AzureResourceID   string            `json:"azure_workspace_resource_id,omitempty"`
	AuthType          string            `json:"auth_type,omitempty"`
	Env               map[string]string `json:"env,omitempty"`
	AssertError       string            `json:"assertError,omitempty"`
	AssertAuth        string            `json:"assertAuth,omitempty"`
	AssertHost        string            `json:"assertHost,omitempty"`
	AssertAzure       bool              `json:"assertAzure,omitempty"`
}

func init() {
	dumpTo = os.Getenv("DATABRICKS_AUTH_FIXTURES_DUMP")
	if dumpTo != "" {
		os.Remove(dumpTo)
	}
}

// set to a file, where auth test fixtures are going to be dumped
var dumpTo string = ""
var dumpToMu sync.Mutex

var defaultAuthBaseErrorMessage = "default auth: cannot configure default credentials, " +
	"please check https://docs.databricks.com/en/dev-tools/auth.html#databricks-client-unified-authentication " +
	"to configure credentials for your preferred authentication method"

func (cf configFixture) dump(t *testing.T) error {
	if dumpTo == "" {
		return nil
	}
	dumpToMu.Lock()
	defer dumpToMu.Unlock()
	f, err := os.OpenFile(dumpTo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o600)
	if err != nil {
		return err
	}
	cf.Name = t.Name()
	raw, err := json.Marshal(cf)
	if err != nil {
		return err
	}
	_, err = f.Write(raw)
	if err != nil {
		return err
	}
	_, err = f.WriteString("\n")
	if err != nil {
		return err
	}
	return f.Close()
}

func (cf configFixture) apply(t *testing.T) {
	err := cf.dump(t)
	if err != nil {
		t.Fatal(err)
		return
	}
	env.CleanupEnvironment(t)
	c, err := cf.configureProviderAndReturnConfig(t)
	if cf.AssertError != "" {
		require.NotNilf(t, err, "Expected to have %s error", cf.AssertError)
		require.Equal(t, cf.AssertError, err.Error())
		return
	}
	if err != nil {
		require.NoError(t, err)
		return
	}
	assert.Equal(t, cf.AssertAzure, c.IsAzure(), "Auth is not for Azure")
	assert.Equal(t, cf.AssertAuth, c.AuthType)
	assert.Equal(t, cf.AssertHost, c.Host)
}

func (cf configFixture) configureProviderAndReturnConfig(t *testing.T) (*Config, error) {
	// Only tests with an explicit CLI fixture should discover external credentials.
	t.Setenv("PATH", "")
	for k, v := range cf.Env {
		os.Setenv(k, v)
	}
	client := &Config{
		Host:              cf.Host,
		Token:             cf.Token,
		Username:          cf.Username,
		Password:          cf.Password,
		Profile:           cf.Profile,
		ConfigFile:        cf.ConfigFile,
		AzureClientID:     cf.AzureClientID,
		AzureClientSecret: cf.AzureClientSecret,
		AzureTenantID:     cf.AzureTenantID,
		AzureResourceID:   cf.AzureResourceID,
		AuthType:          cf.AuthType,
		HTTPTransport:     fixtures.MappingTransport{},
		azureTenantIdFetchClient: makeClient(&http.Response{
			StatusCode: http.StatusTemporaryRedirect,
			Header:     http.Header{"Location": []string{"https://login.microsoftonline.com/tenant_id/abc"}},
		}),
	}
	err := client.Authenticate(&http.Request{Header: http.Header{}})
	if err != nil {
		return nil, err
	}
	return client, nil
}

func TestConfig_NoParams(t *testing.T) {
	configFixture{
		AssertError: defaultAuthBaseErrorMessage,
	}.apply(t)
}

func TestConfig_HostEnv(t *testing.T) {
	configFixture{
		Env: map[string]string{
			"DATABRICKS_HOST": "x",
		},
		AssertError: fmt.Sprintf("%s. Config: host=https://x. Env: DATABRICKS_HOST", defaultAuthBaseErrorMessage),
	}.apply(t)
}

func TestConfig_TokenEnv(t *testing.T) {
	configFixture{
		Env: map[string]string{
			"DATABRICKS_TOKEN": "x",
		},
		AssertError: fmt.Sprintf("%s. Config: token=***. Env: DATABRICKS_TOKEN", defaultAuthBaseErrorMessage),
	}.apply(t)
}

func TestConfig_HostTokenEnv(t *testing.T) {
	configFixture{
		Env: map[string]string{
			"DATABRICKS_HOST":  "x",
			"DATABRICKS_TOKEN": "x",
		},
		AssertAuth: "pat",
		AssertHost: "https://x",
	}.apply(t)
}

func TestConfig_HostParamTokenEnv(t *testing.T) {
	configFixture{
		Host: "https://x",
		Env: map[string]string{
			"DATABRICKS_TOKEN": "x",
		},
		AssertAuth: "pat",
		AssertHost: "https://x",
	}.apply(t)
}

func TestConfig_UserPasswordEnv(t *testing.T) {
	configFixture{
		Env: map[string]string{
			"DATABRICKS_USERNAME": "x",
			"DATABRICKS_PASSWORD": "x",
		},
		AssertError: fmt.Sprintf("%s. Config: username=x, password=***. Env: DATABRICKS_USERNAME, DATABRICKS_PASSWORD", defaultAuthBaseErrorMessage),
		AssertHost:  "https://x",
	}.apply(t)
}

func TestConfig_BasicAuth(t *testing.T) {
	configFixture{
		Env: map[string]string{
			"DATABRICKS_HOST":     "x",
			"DATABRICKS_USERNAME": "x",
			"DATABRICKS_PASSWORD": "x",
		},
		AssertAuth: "basic",
		AssertHost: "https://x",
	}.apply(t)
}

func TestConfig_AttributePrecedence(t *testing.T) {
	configFixture{
		Host: "y",
		Env: map[string]string{
			"DATABRICKS_HOST":     "x",
			"DATABRICKS_USERNAME": "x",
			"DATABRICKS_PASSWORD": "x",
		},
		AssertAuth: "basic",
		AssertHost: "https://y",
	}.apply(t)
}

func TestConfig_BasicAuth_Mix(t *testing.T) {
	configFixture{
		Host:     "y",
		Username: "x",
		Env: map[string]string{
			"DATABRICKS_PASSWORD": "x",
		},
		AssertAuth: "basic",
		AssertHost: "https://y",
	}.apply(t)
}

func TestConfig_BasicAuth_Attrs(t *testing.T) {
	configFixture{
		Host:       "y",
		Username:   "x",
		Password:   "x",
		AssertAuth: "basic",
		AssertHost: "https://y",
	}.apply(t)
}

func TestConfig_ConflictingEnvs(t *testing.T) {
	configFixture{
		Env: map[string]string{
			"DATABRICKS_HOST":     "x",
			"DATABRICKS_TOKEN":    "x",
			"DATABRICKS_USERNAME": "x",
			"DATABRICKS_PASSWORD": "x",
		},
		AssertError: "validate: more than one authorization method configured: basic and pat. Config: host=x, token=***, username=x, password=***. Env: DATABRICKS_HOST, DATABRICKS_TOKEN, DATABRICKS_USERNAME, DATABRICKS_PASSWORD",
	}.apply(t)
}

func TestConfig_ConflictingEnvs_AuthType(t *testing.T) {
	configFixture{
		Env: map[string]string{
			"DATABRICKS_HOST":     "x",
			"DATABRICKS_TOKEN":    "x",
			"DATABRICKS_USERNAME": "x",
			"DATABRICKS_PASSWORD": "x",
		},
		AuthType:   "basic",
		AssertAuth: "basic",
		AssertHost: "https://x",
	}.apply(t)
}

func TestConfig_ConfigFile(t *testing.T) {
	configFixture{
		Env: map[string]string{
			"DATABRICKS_CONFIG_FILE": "x",
		},
		AssertError: fmt.Sprintf("%s. Config: config_file=x. Env: DATABRICKS_CONFIG_FILE", defaultAuthBaseErrorMessage),
	}.apply(t)
}

func TestConfig_ConfigFileSkipDefaultProfileIfHostSpecified(t *testing.T) {
	configFixture{
		Host: "x",
		Env: map[string]string{
			// This directory has a DEFAULT profile in databrickscfg
			"HOME": "testdata",
		},
		AssertError: fmt.Sprintf("%s. Config: host=https://x", defaultAuthBaseErrorMessage),
	}.apply(t)
}

func TestConfig_ConfigFileWithEmptyDefaultProfileSelectDefault(t *testing.T) {
	configFixture{
		Env: map[string]string{
			"HOME": "testdata/empty_default",
		},
		AssertError: defaultAuthBaseErrorMessage,
	}.apply(t)
}

func TestConfig_ConfigFileWithEmptyDefaultProfileSelectAbc(t *testing.T) {
	configFixture{
		Env: map[string]string{
			"HOME":                      "testdata/empty_default",
			"DATABRICKS_CONFIG_PROFILE": "abc",
		},
		AssertAuth: "pat",
		AssertHost: "https://foo",
	}.apply(t)
}

func TestConfig_PatFromDatabricksCfg(t *testing.T) {
	configFixture{
		// loading with DEFAULT profile in databrickscfs
		Env: map[string]string{
			"HOME": "testdata",
		},
		AssertHost: "https://dbc-XXXXXXXX-YYYY.cloud.databricks.com",
		AssertAuth: "pat",
	}.apply(t)
}

func TestConfig_PatFromDatabricksCfg_DotProfile(t *testing.T) {
	configFixture{
		// loading with nohost profile in databrickscfs
		Env: map[string]string{
			"HOME":                      "testdata",
			"DATABRICKS_CONFIG_PROFILE": "pat.with.dot",
		},
		AssertAuth: "pat",
		AssertHost: "https://dbc-XXXXXXXX-YYYY.cloud.databricks.com",
	}.apply(t)
}

func TestConfig_PatFromDatabricksCfg_NohostProfile(t *testing.T) {
	configFixture{
		// loading with nohost profile in databrickscfs
		Env: map[string]string{
			"HOME":                      "testdata",
			"DATABRICKS_CONFIG_PROFILE": "nohost",
		},
		AssertError: fmt.Sprintf("%s. Config: token=***, profile=nohost. Env: DATABRICKS_CONFIG_PROFILE", defaultAuthBaseErrorMessage),
	}.apply(t)
}

func TestConfig_ConfigProfileAndToken(t *testing.T) {
	configFixture{
		Env: map[string]string{
			"DATABRICKS_TOKEN":          "x",
			"DATABRICKS_CONFIG_PROFILE": "nohost",
			"HOME":                      "testdata",
		},
		AssertError: fmt.Sprintf("%s. Config: token=***, profile=nohost. Env: DATABRICKS_TOKEN, DATABRICKS_CONFIG_PROFILE", defaultAuthBaseErrorMessage),
	}.apply(t)
}

func TestConfig_ConfigProfileAndPassword(t *testing.T) {
	configFixture{
		Env: map[string]string{
			"DATABRICKS_USERNAME":       "x",
			"DATABRICKS_CONFIG_PROFILE": "nohost",
			"HOME":                      "testdata",
		},
		AssertError: "validate: more than one authorization method configured: basic and pat. Config: token=***, username=x, profile=nohost. Env: DATABRICKS_USERNAME, DATABRICKS_CONFIG_PROFILE",
	}.apply(t)
}

var azResourceID = "/sub/rg/ws"

func TestConfig_AzurePAT(t *testing.T) {
	configFixture{
		// Azure hostnames can support host+token auth, as usual
		Host:        "https://adb-xxx.y.azuredatabricks.net/",
		Token:       "y",
		AssertAzure: true,
		AssertHost:  "https://adb-xxx.y.azuredatabricks.net",
		AssertAuth:  "pat",
	}.apply(t)
}

func TestConfig_AzureCliHost(t *testing.T) {
	configFixture{
		Host:            "https://adb-123.4.azuredatabricks.net",
		AzureResourceID: azResourceID, // skips ensureWorkspaceUrl
		Env: map[string]string{
			"PATH": testdataPath(),
			"HOME": "testdata/azure",
		},
		AssertAzure: true,
		AssertHost:  "https://adb-123.4.azuredatabricks.net",
		AssertAuth:  "azure-cli",
	}.apply(t)
}

func TestConfig_AzureCliHost_Fail(t *testing.T) {
	configFixture{
		AzureResourceID: azResourceID,
		Env: map[string]string{
			"PATH": testdataPath(),
			"HOME": "testdata/azure",
			"FAIL": "yes",
		},
		AssertError: fmt.Sprintf("%s. Config: azure_workspace_resource_id=/sub/rg/ws", defaultAuthBaseErrorMessage),
	}.apply(t)
}

func TestConfig_AzureCliHost_AzNotInstalled(t *testing.T) {
	configFixture{
		// `az` not installed, which is expected for deployers on other clouds...
		AzureResourceID: azResourceID,
		Env: map[string]string{
			"PATH": "whatever",
			"HOME": "testdata/azure",
		},
		AssertError: fmt.Sprintf("%s. Config: azure_workspace_resource_id=/sub/rg/ws", defaultAuthBaseErrorMessage),
	}.apply(t)
}

func TestConfig_AzureCliHost_PatConflict_WithConfigFilePresentWithoutDefaultProfile(t *testing.T) {
	configFixture{
		AzureResourceID: azResourceID,
		Token:           "x",
		Env: map[string]string{
			"PATH": testdataPath(),
			"HOME": "testdata/azure",
		},
		AssertError: "validate: more than one authorization method configured: azure and pat. Config: token=***, azure_workspace_resource_id=/sub/rg/ws",
	}.apply(t)
}

func TestConfig_AzureCliHostAndResourceID(t *testing.T) {
	configFixture{
		// omit request to management endpoint to get workspace properties
		AzureResourceID: azResourceID,
		Host:            "https://adb-123.4.azuredatabricks.net",
		Env: map[string]string{
			"PATH": testdataPath(),
			"HOME": "testdata", // .databrickscfg has DEFAULT profile
		},
		AssertAzure: true,
		AssertHost:  "https://adb-123.4.azuredatabricks.net",
		AssertAuth:  "azure-cli",
	}.apply(t)
}

func TestConfig_AzureCliHostAndResourceID_ConfigurationPrecedence(t *testing.T) {
	configFixture{
		// omit request to management endpoint to get workspace properties
		AzureResourceID: azResourceID,
		Host:            "https://adb-123.4.azuredatabricks.net",
		Env: map[string]string{
			"PATH":                      testdataPath(),
			"HOME":                      "testdata/azure",
			"DATABRICKS_CONFIG_PROFILE": "justhost",
		},
		AssertAzure: true,
		AssertHost:  "https://adb-123.4.azuredatabricks.net",
		AssertAuth:  "azure-cli",
	}.apply(t)
}

func TestConfig_AzureAndPasswordConflict(t *testing.T) { // TODO: this breaks
	configFixture{
		Host:            "https://adb-123.4.azuredatabricks.net",
		AzureResourceID: azResourceID,
		Env: map[string]string{
			"PATH":                testdataPath(),
			"HOME":                "testdata/azure",
			"DATABRICKS_USERNAME": "x",
		},
		AssertError: "validate: more than one authorization method configured: azure and basic. Config: host=https://adb-123.4.azuredatabricks.net, username=x, azure_workspace_resource_id=/sub/rg/ws. Env: DATABRICKS_USERNAME",
	}.apply(t)
}

func TestConfig_CorruptConfig(t *testing.T) {
	configFixture{
		Env: map[string]string{
			"HOME":                      "testdata/corrupt",
			"DATABRICKS_CONFIG_PROFILE": "DEFAULT",
		},
		AssertError: "resolve: testdata/corrupt/.databrickscfg has no DEFAULT profile configured. Config: profile=DEFAULT. Env: DATABRICKS_CONFIG_PROFILE",
	}.apply(t)
}

func TestConfig_AuthTypeFromEnv(t *testing.T) {
	configFixture{
		Host: "x",
		Env: map[string]string{
			"DATABRICKS_TOKEN":     "token",
			"DATABRICKS_USERNAME":  "user",
			"DATABRICKS_PASSWORD":  "password",
			"DATABRICKS_AUTH_TYPE": "basic",
		},
		AssertAuth: "basic",
		AssertHost: "https://x",
	}.apply(t)
}
