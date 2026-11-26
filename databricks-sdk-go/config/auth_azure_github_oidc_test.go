package config

import (
	"net/http"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/common/environment"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/google/go-cmp/cmp"
)

func errPrefix(prefix string) *string {
	return &prefix
}

func hasPrefix(err error, prefix string) bool {
	if err == nil {
		return false
	}
	return strings.HasPrefix(err.Error(), prefix)
}

func TestAzureGithubOIDCCredentials(t *testing.T) {
	testCases := []struct {
		desc          string
		cfg           *Config
		wantHeaders   map[string]string
		wantErrPrefix *string
	}{
		{
			desc: "not an azure config",
			cfg: &Config{
				DatabricksEnvironment: &environment.DatabricksEnvironment{
					Cloud: "foo-bar-cloud",
				},
			},
			wantErrPrefix: errPrefix("github-oidc-azure auth: not configured"),
		},
		{
			desc: "missing azure client ID",
			cfg: &Config{
				DatabricksEnvironment: &environment.DatabricksEnvironment{
					Cloud:              environment.CloudAzure,
					AzureApplicationID: "test-azure-app-id",
					AzureEnvironment:   &environment.AzurePublicCloud,
				},
				Host:                       "http://host.com/test",
				AzureTenantID:              "test-tenant-id",
				ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
				ActionsIDTokenRequestToken: "token-1337",
				HTTPTransport:              fixtures.MappingTransport{},
			},
			wantErrPrefix: errPrefix("github-oidc-azure auth: not configured"),
		},
		{
			desc: "missing host",
			cfg: &Config{
				DatabricksEnvironment: &environment.DatabricksEnvironment{
					Cloud:              environment.CloudAzure,
					AzureApplicationID: "test-azure-app-id",
					AzureEnvironment:   &environment.AzurePublicCloud,
				},
				AzureClientID:              "test-client-id",
				AzureTenantID:              "test-tenant-id",
				ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
				ActionsIDTokenRequestToken: "token-1337",
			},
			wantErrPrefix: errPrefix("github-oidc-azure auth: not configured"),
		},
		{
			desc: "missing tenant ID",
			cfg: &Config{
				DatabricksEnvironment: &environment.DatabricksEnvironment{
					Cloud:              environment.CloudAzure,
					AzureApplicationID: "test-azure-app-id",
					AzureEnvironment:   &environment.AzurePublicCloud,
				},
				Host:                       "http://host.com/test",
				AzureClientID:              "test-client-id",
				ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
				ActionsIDTokenRequestToken: "token-1337",
				HTTPTransport:              fixtures.MappingTransport{},
			},
			wantErrPrefix: errPrefix("github-oidc-azure auth: not configured"),
		},
		{
			desc: "missing env ACTIONS_ID_TOKEN_REQUEST_TOKEN",
			cfg: &Config{
				DatabricksEnvironment: &environment.DatabricksEnvironment{
					Cloud:              environment.CloudAzure,
					AzureApplicationID: "test-azure-app-id",
					AzureEnvironment:   &environment.AzurePublicCloud,
				},
				AzureClientID:            "test-client-id",
				AzureTenantID:            "test-tenant-id",
				Host:                     "http://host.com/test",
				ActionsIDTokenRequestURL: "http://endpoint.com/test?version=1",
				HTTPTransport:            fixtures.MappingTransport{},
			},
			wantErrPrefix: errPrefix("github-oidc-azure auth: not configured"),
		},
		{
			desc: "missing env ACTIONS_ID_TOKEN_REQUEST_URL",
			cfg: &Config{
				DatabricksEnvironment: &environment.DatabricksEnvironment{
					Cloud:              environment.CloudAzure,
					AzureApplicationID: "test-azure-app-id",
					AzureEnvironment:   &environment.AzurePublicCloud,
				},
				AzureClientID:              "test-client-id",
				AzureTenantID:              "test-tenant-id",
				Host:                       "http://host.com/test",
				ActionsIDTokenRequestToken: "token-1337",
				HTTPTransport:              fixtures.MappingTransport{},
			},
			wantErrPrefix: errPrefix("github-oidc-azure auth: not configured"),
		},
		{
			desc: "azure aad token exchange server error",
			cfg: &Config{
				DatabricksEnvironment: &environment.DatabricksEnvironment{
					Cloud:              environment.CloudAzure,
					AzureApplicationID: "test-azure-app-id",
					AzureEnvironment:   &environment.AzurePublicCloud,
				},
				AzureClientID:              "test-client-id",
				AzureTenantID:              "test-tenant-id",
				ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
				ActionsIDTokenRequestToken: "token-1337",
				Host:                       "http://host.com/test",
				HTTPTransport: fixtures.MappingTransport{
					"GET /test?version=1&audience=api://AzureADTokenExchange": {
						Status: http.StatusInternalServerError,
						ExpectedHeaders: map[string]string{
							"Authorization": "Bearer token-1337",
							"Accept":        "application/json",
						},
					},
				},
			},
			wantErrPrefix: errPrefix("github-oidc-azure"),
		},
		{
			desc: "azure auth server error",
			cfg: &Config{
				DatabricksEnvironment: &environment.DatabricksEnvironment{
					Cloud:              environment.CloudAzure,
					AzureApplicationID: "test-azure-app-id",
					AzureEnvironment:   &environment.AzurePublicCloud,
				},
				AzureClientID:              "test-client-id",
				AzureTenantID:              "test-tenant-id",
				ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
				ActionsIDTokenRequestToken: "token-1337",
				Host:                       "http://host.com/test",
				HTTPTransport: fixtures.MappingTransport{
					"GET /test?version=1&audience=api://AzureADTokenExchange": {
						Status: http.StatusOK,
						ExpectedHeaders: map[string]string{
							"Authorization": "Bearer token-1337",
							"Accept":        "application/json",
						},
						Response: `{"value": "id-token-42"}`,
					},
					"POST /test-tenant-id/oauth2/token": {
						Status: http.StatusInternalServerError,
						ExpectedHeaders: map[string]string{
							"Accept":       "application/json",
							"Content-Type": "application/x-www-form-urlencoded",
						},
					},
				},
			},
			wantErrPrefix: errPrefix("error getting token: http 500"),
		},
		{
			desc: "invalid auth token",
			cfg: &Config{
				DatabricksEnvironment: &environment.DatabricksEnvironment{
					Cloud:              environment.CloudAzure,
					AzureApplicationID: "test-azure-app-id",
					AzureEnvironment:   &environment.AzurePublicCloud,
				},
				AzureClientID:              "test-client-id",
				AzureTenantID:              "test-tenant-id",
				ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
				ActionsIDTokenRequestToken: "token-1337",
				Host:                       "http://host.com/test",
				HTTPTransport: fixtures.MappingTransport{
					"GET /test?version=1&audience=api://AzureADTokenExchange": {
						Status: http.StatusOK,
						ExpectedHeaders: map[string]string{
							"Authorization": "Bearer token-1337",
							"Accept":        "application/json",
						},
						Response: `{"value": "id-token-42"}`,
					},
					"POST /test-tenant-id/oauth2/token": {
						Status: http.StatusOK,
						ExpectedHeaders: map[string]string{
							"Accept":       "application/json",
							"Content-Type": "application/x-www-form-urlencoded",
						},
						Response: map[string]string{
							"foo": "bar",
						},
					},
				},
			},
			wantErrPrefix: errPrefix("error getting token: invalid token"),
		},
		{
			desc: "success",
			cfg: &Config{
				DatabricksEnvironment: &environment.DatabricksEnvironment{
					Cloud:              environment.CloudAzure,
					AzureApplicationID: "test-azure-app-id",
					AzureEnvironment:   &environment.AzurePublicCloud,
				},
				AzureClientID:              "test-client-id",
				AzureTenantID:              "test-tenant-id",
				ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
				ActionsIDTokenRequestToken: "token-1337",
				Host:                       "http://host.com/test",
				HTTPTransport: fixtures.MappingTransport{
					"GET /test?version=1&audience=api://AzureADTokenExchange": {
						Status: http.StatusOK,
						ExpectedHeaders: map[string]string{
							"Authorization": "Bearer token-1337",
							"Accept":        "application/json",
						},
						Response: `{"value": "id-token-42"}`,
					},
					"POST /test-tenant-id/oauth2/token": {
						Status: http.StatusOK,
						ExpectedHeaders: map[string]string{
							"Accept":       "application/json",
							"Content-Type": "application/x-www-form-urlencoded",
						},
						Response: map[string]string{
							"token_type":    "access-token",
							"access_token":  "test-auth-token",
							"refresh_token": "refresh",
							"expires_on":    "0",
						},
					},
				},
			},
			wantHeaders: map[string]string{
				"Authorization": "access-token test-auth-token",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// Clear OIDC env vars that GitHub Actions injects when id-token: write
			// is set. Without this, Config picks them up via the env: struct tag
			// and tests that assert "not configured" fail.
			if tc.cfg.ActionsIDTokenRequestURL == "" {
				t.Setenv("ACTIONS_ID_TOKEN_REQUEST_URL", "")
			}
			if tc.cfg.ActionsIDTokenRequestToken == "" {
				t.Setenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN", "")
			}
			tc.cfg.Credentials = &AzureGithubOIDCCredentials{} // only test this credential strategy
			tc.cfg.DebugHeaders = true
			if tc.wantHeaders == nil {
				tc.wantHeaders = map[string]string{}
			}

			req, _ := http.NewRequest("GET", "http://localhost", nil)
			gotErr := tc.cfg.Authenticate(req)
			gotHeaders := map[string]string{}
			for h := range req.Header {
				gotHeaders[h] = req.Header.Get(h)
			}

			if tc.wantErrPrefix == nil && gotErr != nil {
				t.Errorf("Authenticate(): got error %q, want none", gotErr)
			}
			if tc.wantErrPrefix != nil && !hasPrefix(gotErr, *tc.wantErrPrefix) {
				t.Errorf("Authenticate(): got error %q, want error with prefix %q", gotErr, *tc.wantErrPrefix)
			}
			if diff := cmp.Diff(tc.wantHeaders, gotHeaders); diff != "" {
				t.Errorf("Authenticate(): mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestAzureGithubOIDCCredentials_Name(t *testing.T) {
	c := AzureGithubOIDCCredentials{}
	want := "github-oidc-azure"

	if got := c.Name(); got != want {
		t.Errorf("Name(): got %s, want %s", got, want)
	}
}
