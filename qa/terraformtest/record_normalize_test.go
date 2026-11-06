package terraformtest

import (
	"encoding/base64"
	"net/url"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestRecordingNormalizesProfileHostRoundTrip(t *testing.T) {
	for _, profileHost := range []string{
		"https://recording-workspace.cloud.databricks.com",
		"https://recording-unified.databricks.com",
		"http://127.0.0.1:45678",
		"http://[::1]:45678",
	} {
		t.Run(profileHost, func(t *testing.T) {
			host, err := url.Parse(profileHost)
			require.NoError(t, err)
			n, err := newRecordingNormalizer(&testDefinition{}, &config.Config{Host: profileHost})
			require.NoError(t, err)
			endpoint := profileHost + "/oidc/accounts/account-id"
			original := map[string]any{
				"host":     host.Host,
				"endpoint": endpoint,
				"audiences": []any{
					profileHost + "/oidc/v1/token",
					"https://unrelated.databricks.com/oidc",
				},
				"text":   "call " + endpoint + " then " + endpoint,
				"base64": base64.StdEncoding.EncodeToString([]byte(endpoint)),
			}
			data, err := yaml.Marshal(original)
			require.NoError(t, err)
			normalized, err := n.transform(data, false, true)
			require.NoError(t, err)
			var values map[string]any
			require.NoError(t, yaml.Unmarshal(normalized, &values))
			replayURL := host.Scheme + "://example.databricks.com"
			require.Equal(t, "example.databricks.com", values["host"])
			require.Equal(t, replayURL+"/oidc/accounts/account-id", values["endpoint"])
			require.Equal(t, []any{replayURL + "/oidc/v1/token", "https://unrelated.databricks.com/oidc"}, values["audiences"])
			require.Equal(t, "call "+replayURL+"/oidc/accounts/account-id then "+replayURL+"/oidc/accounts/account-id", values["text"])
			decoded, err := base64.StdEncoding.DecodeString(values["base64"].(string))
			require.NoError(t, err)
			require.Equal(t, replayURL+"/oidc/accounts/account-id", string(decoded))
			live, err := n.transform(normalized, true, false)
			require.NoError(t, err)
			require.NoError(t, yaml.Unmarshal(live, &values))
			require.Equal(t, original, values)
		})
	}
}

func TestRecordingHostNormalizationKeepsOtherHosts(t *testing.T) {
	n, err := newRecordingNormalizer(&testDefinition{}, &config.Config{Host: "https://workspace.databricks.com"})
	require.NoError(t, err)
	for _, value := range []string{
		"https://other.workspace.databricks.com/oidc",
		"https://workspace.databricks.com.other/oidc",
		"https://prefix-workspace.databricks.com/oidc",
		"https://workspace.databricks.com:8443/oidc",
		"https://other.example.databricks.com/oidc",
		"https://example.databricks.com.other/oidc",
	} {
		require.Equal(t, value, n.text(value, false))
		require.Equal(t, value, n.text(value, true))
	}
}
