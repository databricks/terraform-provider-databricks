package terraformtest

import (
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/stretchr/testify/require"
)

const authDefinition = `version: 1
auth:
  type: azure-msi
  azure_use_msi: true
  azure_client_id: replay-azure-sp
steps:
  - name: create
    config: 'provider "databricks" {}'
    calls: []
destroy:
  calls: []
`

func TestAuthReplay(t *testing.T) {
	Run(t, "testdata/auth")
}

func TestDecodeAuth(t *testing.T) {
	d, err := decodeDefinition(strings.NewReader(authDefinition))
	require.NoError(t, err)
	require.Equal(t, journeyAuth{"type": "azure-msi", "azure_use_msi": "true",
		"azure_client_id": "replay-azure-sp"}, d.Auth)
}

func TestDecodeAuthRejectsInvalidAuth(t *testing.T) {
	for _, tc := range []struct {
		name, before, after, message string
	}{
		{"unsupported type", "type: azure-msi", "type: databricks-cli",
			"auth.type: expected one of pat, basic, oauth-m2m, azure-client-secret, " +
				`azure-msi, got "databricks-cli"`},
		{"missing type", "  type: azure-msi\n", "", `auth.type: expected one of`},
		{"empty block", "auth:\n  type: azure-msi\n  azure_use_msi: true\n" +
			"  azure_client_id: replay-azure-sp\n", "auth: {}\n", `got ""`},
		{"not an attribute", "azure_client_id:", "azure_client:",
			"auth.azure_client: not a credential attribute of the provider"},
		{"routing attribute", "azure_client_id: replay-azure-sp", "host: example",
			"auth.host: not a credential attribute of the provider"},
		{"auth type attribute", "azure_client_id: replay-azure-sp", "auth_type: pat",
			"auth.auth_type: not a credential attribute of the provider"},
		{"empty value", "azure_client_id: replay-azure-sp", `azure_client_id: ""`,
			"auth.azure_client_id: value must not be empty"},
		{"invalid boolean", "azure_use_msi: true", "azure_use_msi: maybe",
			`auth.azure_use_msi: strconv.ParseBool: parsing "maybe": invalid syntax`},
		{"nested value", "azure_client_id: replay-azure-sp", "azure_client_id: [a]",
			"cannot unmarshal !!seq into string"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.Replace(authDefinition, tc.before, tc.after, 1)
			require.NotEqual(t, authDefinition, input)
			_, err := decodeDefinition(strings.NewReader(input))
			require.ErrorContains(t, err, tc.message)
		})
	}
}

// Credentials in a provider block are cleared, so auth is their only source, while
// the fixture host, loader and transport are kept.
func TestAuthReplacesProviderBlockCredentials(t *testing.T) {
	server := "http://127.0.0.1:1234"
	customize := journeyAuth{"type": "basic", "username": "replay-user",
		"password": "replay-password"}.withCredentials(fixtureConfig(server))
	cfg := &config.Config{Token: "hcl-token", ClientID: "hcl-client",
		ClientSecret: "hcl-secret", AzureUseMSI: true, AuthType: "oauth-m2m",
		AccountID: "replay-account"}
	require.NoError(t, customize(cfg))
	require.Equal(t, &config.Config{Host: server, AuthType: "basic",
		Username: "replay-user", Password: "replay-password", AccountID: "replay-account",
		Loaders: cfg.Loaders, HTTPTransport: cfg.HTTPTransport}, cfg)
	require.Equal(t, []config.Loader{fixtureLoader{}}, cfg.Loaders)
	require.IsType(t, fixtureTransport{}, cfg.HTTPTransport)
}

func TestNoAuthKeepsFixturePAT(t *testing.T) {
	cfg := &config.Config{Username: "hcl-user"}
	customize := journeyAuth(nil).withCredentials(fixtureConfig("http://127.0.0.1"))
	require.NoError(t, customize(cfg))
	require.Equal(t, "fixture-token", cfg.Token)
	require.Equal(t, "pat", cfg.AuthType)
	require.Equal(t, config.PatCredentials{}, cfg.Credentials)
}

func TestRecordingRejectsAuthAndEnvironment(t *testing.T) {
	for _, d := range []*testDefinition{
		{Auth: journeyAuth{"type": "basic"}},
		{Environment: map[string]string{}},
	} {
		require.EqualError(t, validateRecordingSetup(d), "recording does not support "+
			"auth or environment; update the fixtures by hand")
	}
}
