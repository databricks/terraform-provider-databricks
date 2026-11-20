package catalog

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"sync"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// These tests pin the Azure and Google credential headers and token requests of
// databricks_volume refreshes, with the server and helpers of
// resource_volume_auth_test.go.

// authGoogleKey is a service account key generated once, so that no key is checked in.
var authGoogleKey = sync.OnceValue(func() string {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	der, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		panic(err)
	}
	return string(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: der}))
})

// authGoogleConfig authenticates with a service account key whose token_uri is Google's.
func authGoogleConfig(t *testing.T, authType string) *config.Config {
	t.Helper()
	credentials, err := json.Marshal(map[string]string{
		"type":           "service_account",
		"project_id":     "replay-project",
		"private_key_id": "replay-key",
		"private_key":    authGoogleKey(),
		"client_email":   authGCPAccount,
		"client_id":      "100000000000000000199",
		"token_uri":      "https://" + authGoogleURL,
	})
	require.NoError(t, err)
	return &config.Config{Host: "https://" + authGCPHost, AuthType: authType,
		GoogleCredentials: string(credentials)}
}

// authGoogleToken is a JWT bearer request to Google's token endpoint with the given
// claim, which asks for an ID token for the workspace or for an access token.
func authGoogleToken(claim string, status int, response string) authExchange {
	return authExchange{authRequest{method: "POST", url: authGoogleURL,
		body: `assertion={"aud":"https://` + authGoogleURL + `","iss":"` +
			authGCPAccount + `",` + claim +
			`}&grant_type=urn:ietf:params:oauth:grant-type:jwt-bearer`,
		headers: map[string]string{}}, status, response}
}

func authGoogleIDToken(status int, response string) authExchange {
	return authGoogleToken(`"target_audience":"https://`+authGCPHost+`"`,
		status, response)
}

func authGoogleAccessToken(status int, response string) authExchange {
	return authGoogleToken(`"scope":"https://www.googleapis.com/auth/cloud-platform `+
		`https://www.googleapis.com/auth/compute"`, status, response)
}

// authGoogleSuccess is the script of a google-credentials refresh with both tokens.
var authGoogleSuccess = []authExchange{
	authGoogleIDToken(http.StatusOK, `{"id_token": "`+authIDToken+`"}`),
	authGoogleAccessToken(http.StatusOK, authBearerToken("gcp-sa-token", 3600)),
	authVolumeRead(authGCPHost, map[string]string{
		"Authorization":                    "Bearer " + authIDToken,
		"X-Databricks-Gcp-Sa-Access-Token": "gcp-sa-token",
	}),
}

// authAADClientCredentials is the client_credentials request of Azure service
// principal replay-azure-sp for resource, with its secret in the Authorization header.
func authAADClientCredentials(resource string, status int, response string) authExchange {
	return authExchange{authRequest{method: "POST", url: authAADToken,
		body: "grant_type=client_credentials&resource=" + url.QueryEscape(resource),
		headers: map[string]string{"Authorization": basicAuthorization(
			"replay-azure-sp", "replay-azure-secret")}}, status, response}
}

// authAADRejectedTwice is the rejection of the client_credentials request for
// resource, and of the retry that oauth2 sends with the secret in the body instead.
func authAADRejectedTwice(resource string) []authExchange {
	return []authExchange{
		authAADClientCredentials(resource, http.StatusUnauthorized, authAADRejected),
		{authRequest{method: "POST", url: authAADToken,
			body: "client_id=replay-azure-sp&client_secret=replay-azure-secret" +
				"&grant_type=client_credentials&resource=" + url.QueryEscape(resource),
			headers: map[string]string{}},
			http.StatusUnauthorized, authAADRejected},
	}
}

// authMSIRequest is the Azure instance metadata token request for resource.
func authMSIRequest(resource, token string) authExchange {
	return authExchange{authRequest{method: "GET", url: authMSIToken +
		"?api-version=2018-02-01&resource=" + url.QueryEscape(resource),
		headers: map[string]string{"Metadata": "true"}}, http.StatusOK,
		fmt.Sprintf(`{"token_type": "Bearer", "access_token": %q,
			"expires_on": 4102444800}`, token)}
}

func authAzureClientSecretConfig(*testing.T) *config.Config {
	return &config.Config{Host: "https://" + authAzureHost,
		AuthType: "azure-client-secret", AzureResourceID: authAzureWSID,
		AzureTenantID: "replay-tenant",
		AzureClientID: "replay-azure-sp", AzureClientSecret: "replay-azure-secret"}
}

func authGoogleCredentialsConfig(t *testing.T) *config.Config {
	return authGoogleConfig(t, "google-credentials")
}

func TestVolumeAuthCloudCredentials(t *testing.T) {
	testVolumeAuthCloudCredentials(t, false)
}

func testVolumeAuthCloudCredentials(t *testing.T, sdkGo bool) {
	tests := []struct {
		name     string
		metadata string
		config   func(t *testing.T) *config.Config
		script   []authExchange
		// summary is the refresh error, or empty if the refresh succeeds.
		summary string
	}{
		{
			name:     "azure_client_secret",
			metadata: authCloudMetadata("Azure"),
			config:   authAzureClientSecretConfig,
			script: []authExchange{
				authAADClientCredentials(authAzureAppID, http.StatusOK,
					authBearerToken("aad-token", 3600)),
				authAADClientCredentials(authAzureMgmt, http.StatusOK,
					authBearerToken("management-token", 3600)),
				authVolumeRead(authAzureHost, map[string]string{
					"Authorization":                            "Bearer aad-token",
					"X-Databricks-Azure-Sp-Management-Token":   "management-token",
					"X-Databricks-Azure-Workspace-Resource-Id": authAzureWSID,
				}),
			},
		},
		{
			name:     "azure_msi",
			metadata: authCloudMetadata("Azure"),
			config: func(*testing.T) *config.Config {
				return &config.Config{Host: "https://" + authAzureHost,
					AuthType: "azure-msi", AzureResourceID: authAzureWSID,
					AzureUseMSI: true}
			},
			script: []authExchange{
				authMSIRequest(authAzureAppID, "msi-token"),
				authMSIRequest(authAzureMgmt, "msi-management-token"),
				authVolumeRead(authAzureHost, map[string]string{
					"Authorization":                            "Bearer msi-token",
					"X-Databricks-Azure-Sp-Management-Token":   "msi-management-token",
					"X-Databricks-Azure-Workspace-Resource-Id": authAzureWSID,
				}),
			},
		},
		{
			name:     "google_credentials",
			metadata: authCloudMetadata("GCP"),
			config:   authGoogleCredentialsConfig,
			script:   authGoogleSuccess,
		},
		{
			name:     "azure_primary_token_fails",
			metadata: authCloudMetadata("Azure"),
			config:   authAzureClientSecretConfig,
			script:   authAADRejectedTwice(authAzureAppID),
			summary: "cannot read volume: failed during request visitor: inner token: " +
				`Post "https://` + authAADToken + `": ` +
				"AADSTS7000215: Invalid client secret provided",
		},
		{
			name:     "azure_required_secondary_token_fails",
			metadata: authCloudMetadata("Azure"),
			config:   authAzureClientSecretConfig,
			script: slices.Concat([]authExchange{authAADClientCredentials(authAzureAppID,
				http.StatusOK, authBearerToken("aad-token", 3600))},
				authAADRejectedTwice(authAzureMgmt)),
			summary: "cannot read volume: failed during request visitor: cloud token: " +
				`Post "https://` + authAADToken + `": ` +
				"AADSTS7000215: Invalid client secret provided",
		},
		{
			// The Google client fetches the ID token when the strategy is configured, so
			// the error is the chain's, not the visitor's inner token error. A failed
			// configuration isn't kept, so the GET's request visitor configures it again
			// after the workspace client did.
			name:     "google_primary_token_fails",
			metadata: authCloudMetadata("GCP"),
			config:   authGoogleCredentialsConfig,
			script: []authExchange{
				authGoogleIDToken(http.StatusBadRequest, authGoogleRejected),
				authGoogleIDToken(http.StatusBadRequest, authGoogleRejected),
			},
			summary: "cannot read volume: failed during request visitor: default auth: " +
				"could not obtain OIDC token from JSON: oauth2: cannot fetch token: " +
				`Post "https://` + authGoogleURL + `": ` + authGoogleRejected +
				". Config: host=https://" + authGCPHost +
				", workspace_id=12345, google_credentials=***",
		},
		{
			name:     "google_optional_secondary_token_fails",
			metadata: authCloudMetadata("GCP"),
			config:   authGoogleCredentialsConfig,
			script: []authExchange{
				authGoogleIDToken(http.StatusOK, `{"id_token": "`+authIDToken+`"}`),
				authGoogleAccessToken(http.StatusBadRequest, authGoogleRejected),
				authVolumeRead(authGCPHost, map[string]string{
					"Authorization": "Bearer " + authIDToken}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &authServer{metadata: tt.metadata, script: tt.script}

			a := refreshAuthVolume(t, s, tt.config(t), sdkGo)

			if tt.summary == "" {
				a.assertSucceeded(t)
			} else {
				a.assertFailed(t, tt.summary)
			}
		})
	}
}

func TestVolumeAuthDeprecatedGoogleAuthTypes(t *testing.T) {
	testVolumeAuthDeprecatedGoogleAuthTypes(t, false)
}

func testVolumeAuthDeprecatedGoogleAuthTypes(t *testing.T, sdkGo bool) {
	t.Run("google-creds", func(t *testing.T) {
		s := &authServer{metadata: authCloudMetadata("GCP"), script: authGoogleSuccess}

		a := refreshAuthVolume(t, s, authGoogleConfig(t, "google-creds"), sdkGo)

		a.assertSucceeded(t)
		assert.Equal(t, "google-credentials", a.client.Config.AuthType)
	})
	// google-id impersonates the service account through Google's IAM Credentials API,
	// with a client the SDK gives no way to redirect, so these stop before it: the
	// application default credentials the impersonation starts from are missing.
	for _, authType := range []string{"google-accounts", "google-workspace"} {
		t.Run(authType, func(t *testing.T) {
			missing := t.TempDir() + "/missing.json"
			t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", missing)
			s := &authServer{metadata: authCloudMetadata("GCP")}

			a := refreshAuthVolume(t, s, &config.Config{Host: "https://" + authGCPHost,
				AuthType: authType, GoogleServiceAccount: authGCPAccount}, sdkGo)

			a.assertFailed(t, "cannot read volume: failed during request visitor: "+
				"default auth: could not obtain OIDC token. google: error getting "+
				"credentials using GOOGLE_APPLICATION_CREDENTIALS environment variable: "+
				"open "+missing+": no such file or directory Running 'gcloud auth "+
				"application-default login' may help. Config: host=https://"+authGCPHost+
				", workspace_id=12345, google_service_account="+authGCPAccount)
			assert.Equal(t, "google-id", a.client.Config.AuthType)
		})
	}
}
