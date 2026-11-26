package config

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"sync/atomic"
	"testing"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/credentials/u2m"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"golang.org/x/oauth2"
)

func TestM2mHappyFlow(t *testing.T) {
	assertHeaders(
		t,
		&Config{
			Host:         "a",
			ClientID:     "b",
			ClientSecret: "c",
			AuthType:     "oauth-m2m",
			HTTPTransport: fixtures.MappingTransport{
				"GET /oidc/.well-known/oauth-authorization-server": {
					Response: u2m.OAuthAuthorizationServer{
						AuthorizationEndpoint: "https://localhost:1234/dummy/auth",
						TokenEndpoint:         "https://localhost:1234/dummy/token",
					},
				},
				"POST /dummy/token": {
					ExpectedHeaders: map[string]string{
						"Authorization": "Basic Yjpj",
						"Content-Type":  "application/x-www-form-urlencoded",
					},
					ExpectedRequest: url.Values{
						"grant_type": {"client_credentials"},
						"scope":      {"all-apis"},
					},
					Response: oauth2.Token{
						TokenType:   "Some",
						AccessToken: "cde",
					},
				},
			},
		},
		map[string]string{
			"Authorization": "Some cde",
		},
	)
}

// m2mGroupRoleTokenFixture provides the common token endpoint expectations
// used by the workspace, unified, and account group-role tests.
func m2mGroupRoleTokenFixture() fixtures.HTTPFixture {
	return fixtures.HTTPFixture{
		ExpectedRequest: url.Values{
			"assume_group": {"group-123"},
			"grant_type":   {"client_credentials"},
			"scope":        {"all-apis"},
		},
		Response: oauth2.Token{
			TokenType:   "Bearer",
			AccessToken: "role-token",
		},
	}
}

// TestM2mCredentials_GroupRoleWorkspace verifies that workspace OAuth
// discovery leads to a token request containing assume_group.
func TestM2mCredentials_GroupRoleWorkspace(t *testing.T) {
	cfg := &Config{
		Host:         "a",
		ClientID:     "b",
		ClientSecret: "c",
		GroupID:      "group-123",
		AuthType:     "oauth-m2m",
		HTTPTransport: fixtures.MappingTransport{
			"GET /oidc/.well-known/oauth-authorization-server": {
				Response: u2m.OAuthAuthorizationServer{
					TokenEndpoint: "https://localhost/dummy/token",
				},
			},
			"POST /dummy/token": m2mGroupRoleTokenFixture(),
		},
		ConfigFile: "/dev/null",
	}

	assertHeaders(t, cfg, map[string]string{
		"Authorization": "Bearer role-token",
	})
}

// TestM2mCredentials_GroupRoleUnified verifies that unified OAuth discovery
// forwards assume_group to the account-scoped token endpoint.
func TestM2mCredentials_GroupRoleUnified(t *testing.T) {
	cfg := &Config{
		Host:         "https://unified.cloud.databricks.com",
		AccountID:    "account-123",
		ClientID:     "b",
		ClientSecret: "c",
		GroupID:      "group-123",
		AuthType:     "oauth-m2m",
		HostMetadataResolver: func(context.Context, string) (*HostMetadata, error) {
			return &HostMetadata{
				HostType:     UnifiedHost,
				OIDCEndpoint: "https://unified.cloud.databricks.com/oidc/accounts/{account_id}",
			}, nil
		},
		HTTPTransport: fixtures.MappingTransport{
			"GET /oidc/accounts/account-123/.well-known/oauth-authorization-server": {
				Response: u2m.OAuthAuthorizationServer{
					TokenEndpoint: "https://localhost/oidc/accounts/account-123/v1/token",
				},
			},
			"POST /oidc/accounts/account-123/v1/token": m2mGroupRoleTokenFixture(),
		},
		ConfigFile: "/dev/null",
	}

	assertHeaders(t, cfg, map[string]string{
		"Authorization": "Bearer role-token",
	})
}

// TestM2mCredentials_GroupRoleAccount verifies that account M2M forwards
// assume_group directly to its account-scoped token endpoint.
func TestM2mCredentials_GroupRoleAccount(t *testing.T) {
	cfg := &Config{
		Host:         "accounts.cloud.databricks.com",
		AccountID:    "account-123",
		ClientID:     "b",
		ClientSecret: "c",
		GroupID:      "group-123",
		AuthType:     "oauth-m2m",
		HTTPTransport: fixtures.MappingTransport{
			"POST /oidc/accounts/account-123/v1/token": m2mGroupRoleTokenFixture(),
		},
		ConfigFile: "/dev/null",
	}

	assertHeaders(t, cfg, map[string]string{
		"Authorization": "Bearer role-token",
	})
}

// TestM2mGroupRoleServerRejectionDoesNotFallback verifies that a rejected role
// request is not retried without assume_group.
func TestM2mGroupRoleServerRejectionDoesNotFallback(t *testing.T) {
	var tokenCalls int32

	transport := &postCountingTransport{
		calls: &tokenCalls,
		inner: fixtures.MappingTransport{
			"GET /oidc/.well-known/oauth-authorization-server": {
				Response: u2m.OAuthAuthorizationServer{TokenEndpoint: "https://localhost/token"},
			},
			"POST /token": {
				Status: http.StatusBadRequest,
				ExpectedRequest: url.Values{
					"assume_group": {"group-123"},
					"grant_type":   {"client_credentials"},
					"scope":        {"all-apis"},
				},
				Response: `{"error":"invalid_target"}`,
			},
		},
	}

	cfg := &Config{
		Host:          "a",
		ClientID:      "b",
		ClientSecret:  "c",
		GroupID:       "group-123",
		AuthType:      "oauth-m2m",
		HTTPTransport: transport,
		ConfigFile:    "/dev/null",
	}

	_, err := authenticateRequest(cfg)
	if err == nil {
		t.Fatal("authenticateRequest() error = nil, want token server rejection")
	}

	if got := atomic.LoadInt32(&tokenCalls); got != 1 {
		t.Errorf("token endpoint calls = %d, want 1", got)
	}
}

// TestM2mCredentials_CachesAreIsolatedByClient verifies that clients assuming
// different group roles do not reuse each other's M2M access tokens.
func TestM2mCredentials_CachesAreIsolatedByClient(t *testing.T) {
	server, tokenCalls := newGroupTokenServer(t)

	testCases := []struct {
		name    string
		groupID string
		want    string
	}{
		{
			name:    "normal access",
			groupID: "",
			want:    "Bearer token-normal",
		},
		{
			name:    "group A",
			groupID: "group-a",
			want:    "Bearer token-group-a",
		},
		{
			name:    "group B",
			groupID: "group-b",
			want:    "Bearer token-group-b",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			cfg := &Config{
				Host:          server.URL,
				ClientID:      "client-id",
				ClientSecret:  "client-secret",
				GroupID:       testCase.groupID,
				AuthType:      "oauth-m2m",
				ConfigFile:    "/dev/null",
				HTTPTransport: server.Client().Transport,
			}

			for range 2 {
				assertHeaders(t, cfg, map[string]string{"Authorization": testCase.want})
			}
		})
	}

	if *tokenCalls != len(testCases) {
		t.Errorf("Token endpoint was called %d times; want one call per client", *tokenCalls)
	}
}

func TestM2mHappyFlowForAccount(t *testing.T) {
	assertHeaders(t,
		&Config{
			Host:         "accounts.cloud.databricks.com",
			AccountID:    "a",
			ClientID:     "b",
			ClientSecret: "c",
			HTTPTransport: fixtures.MappingTransport{
				"POST /oidc/accounts/a/v1/token": {
					ExpectedHeaders: map[string]string{
						"Authorization": "Basic Yjpj",
						"Content-Type":  "application/x-www-form-urlencoded",
					},
					ExpectedRequest: url.Values{
						"grant_type": {"client_credentials"},
						"scope":      {"all-apis"},
					},
					Response: oauth2.Token{
						TokenType:   "Some",
						AccessToken: "cde",
					},
				},
			},
		},
		map[string]string{
			"Authorization": "Some cde",
		},
	)
}

func TestM2mNotSupported(t *testing.T) {
	_, err := authenticateRequest(&Config{
		Host:         "a",
		ClientID:     "b",
		ClientSecret: "c",
		AuthType:     "oauth-m2m",
		HTTPTransport: fixtures.MappingTransport{
			"GET /oidc/.well-known/oauth-authorization-server": {
				Status:   404,
				Response: `...`,
			},
		},
	})
	if !errors.Is(err, u2m.ErrOAuthNotSupported) {
		t.Errorf("got error %v, want %v", err, u2m.ErrOAuthNotSupported)
	}
}

// TestM2mCredentials_DirectTokenSource verifies that M2mCredentials.Configure
// plumbs a direct token source (ccfg.Token) through cachedTokenSource rather
// than wrapping clientcredentials.Config.TokenSource, which returns an
// oauth2.ReuseTokenSource that adds a second cache layer. With double-caching,
// the proactive async refresh in cachedTokenSource is silently suppressed until
// ~10 s before expiry, causing bursts of 401 errors at token rotation boundaries.
// See https://github.com/databricks/databricks-sdk-go/issues/1549.
func TestM2mCredentials_DirectTokenSource(t *testing.T) {
	var tokenCalls int32
	transport := &postCountingTransport{
		calls: &tokenCalls,
		inner: fixtures.MappingTransport{
			"GET /oidc/.well-known/oauth-authorization-server": {
				Response: u2m.OAuthAuthorizationServer{
					TokenEndpoint: "https://localhost/token",
				},
			},
			"POST /token": {
				Response: oauth2.Token{
					TokenType:   "Bearer",
					AccessToken: "test-token",
				},
			},
		},
	}

	cfg := &Config{
		Host:          "a",
		ClientID:      "b",
		ClientSecret:  "c",
		AuthType:      "oauth-m2m",
		ConfigFile:    "/dev/null",
		HTTPTransport: transport,
	}

	err := cfg.EnsureResolved()
	if err != nil {
		t.Fatalf("EnsureResolved(): %v", err)
	}

	ctx := cfg.refreshClient.InContextForOAuth2(cfg.refreshCtx)
	provider, err := M2mCredentials{}.Configure(ctx, cfg)
	if err != nil {
		t.Fatalf("Configure(): %v", err)
	}

	oauthProvider := provider.(credentials.OAuthCredentialsProvider)

	// Token() goes through cachedTokenSource, which fetches once and caches.
	// Verify the endpoint is reached (not short-circuited by an inner cache).
	tok, err := oauthProvider.Token(context.Background())
	if err != nil {
		t.Fatalf("Token(): %v", err)
	}
	if tok.AccessToken == "" {
		t.Fatalf("Token(): empty access token")
	}

	if got := int(atomic.LoadInt32(&tokenCalls)); got != 1 {
		t.Errorf("token endpoint calls = %d, want 1", got)
	}
}

type postCountingTransport struct {
	calls *int32
	inner http.RoundTripper
}

func (t *postCountingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Method == "POST" {
		atomic.AddInt32(t.calls, 1)
	}
	return t.inner.RoundTrip(req)
}

func (t *postCountingTransport) SkipRetryOnIO() bool { return true }

func TestM2M_Scopes(t *testing.T) {
	tests := []struct {
		name   string
		scopes []string
		want   string
	}{
		{
			name:   "nil scopes uses default",
			scopes: nil,
			want:   "all-apis",
		},
		{
			name:   "empty scopes uses default",
			scopes: []string{},
			want:   "all-apis",
		},
		{
			name:   "single scope",
			scopes: []string{"dashboards"},
			want:   "dashboards",
		},
		{
			name:   "multiple scopes are sorted",
			scopes: []string{"jobs", "files:read", "mlflow"},
			want:   "files:read jobs mlflow",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transport := fixtures.MappingTransport{
				"GET /oidc/.well-known/oauth-authorization-server": {
					Response: u2m.OAuthAuthorizationServer{
						AuthorizationEndpoint: "https://localhost/auth",
						TokenEndpoint:         "https://localhost/token",
					},
				},
				"POST /token": {
					// The scope assertion: verifies the SDK sends the correct scope parameter.
					ExpectedRequest: url.Values{
						"grant_type": {"client_credentials"},
						"scope":      {tt.want},
					},
					Response: oauth2.Token{
						TokenType:   "Bearer",
						AccessToken: "test-token",
					},
				},
			}

			cfg := &Config{
				Host:          "a",
				ClientID:      "b",
				ClientSecret:  "c",
				AuthType:      "oauth-m2m",
				Scopes:        tt.scopes,
				HTTPTransport: transport,
				ConfigFile:    "/dev/null",
			}

			req, err := http.NewRequest("GET", "http://localhost", nil)
			if err != nil {
				t.Fatalf("http.NewRequest(): unexpected error: %v", err)
			}
			err = cfg.Authenticate(req)
			if err != nil {
				t.Fatalf("Authenticate(): unexpected error: %v", err)
			}
			if got, want := req.Header.Get("Authorization"), "Bearer test-token"; got != want {
				t.Errorf("Authorization header: got %q, want %q", got, want)
			}
		})
	}
}
