package config

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/databricks/databricks-sdk-go/config/experimental/auth/oidc"
	"github.com/databricks/databricks-sdk-go/credentials/u2m"
)

func TestGithubOIDC_Scopes(t *testing.T) {
	const oidcTokenPath = "/oidc/v1/token"

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
			scopes: []string{"clusters"},
			want:   "clusters",
		},
		{
			name:   "multiple scopes are sorted",
			scopes: []string{"jobs", "clusters", "files:read"},
			want:   "clusters files:read jobs",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock GitHub server for OIDC token requests.
			githubServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(map[string]string{"value": "github-id-token"})
			}))
			defer githubServer.Close()

			// Mock Databricks server to verify the SDK passes the correct scopes.
			var databricksServer *httptest.Server
			databricksServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch r.URL.Path {
				case "/oidc/.well-known/oauth-authorization-server":
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(u2m.OAuthAuthorizationServer{
						AuthorizationEndpoint: "https://host.com/oidc/v1/authorize",
						TokenEndpoint:         databricksServer.URL + oidcTokenPath,
					})

				case oidcTokenPath:
					if err := r.ParseForm(); err != nil {
						t.Fatalf("Failed to parse form: %v", err)
					}
					// The scope assertion: verifies the SDK sends the correct scope parameter.
					if got := r.Form.Get("scope"); got != tt.want {
						t.Errorf("scope: got %q, want %q", got, tt.want)
					}
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(map[string]interface{}{
						"token_type":   "Bearer",
						"access_token": "databricks-access-token",
						"expires_in":   3600,
					})

				case "/.well-known/databricks-config":
					http.Error(w, "Not found", http.StatusNotFound)

				default:
					t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
					http.Error(w, "Not found", http.StatusNotFound)
				}
			}))
			defer databricksServer.Close()

			cfg := &Config{
				Host:                       databricksServer.URL,
				ClientID:                   "test-client-id",
				ActionsIDTokenRequestURL:   githubServer.URL + "/github-token?version=1",
				ActionsIDTokenRequestToken: "github-request-token",
				TokenAudience:              "databricks-test-audience",
				AuthType:                   "github-oidc",
				Scopes:                     tt.scopes,
			}

			req, err := http.NewRequest("GET", databricksServer.URL+"/api/test", nil)
			if err != nil {
				t.Fatalf("http.NewRequest(): unexpected error: %v", err)
			}
			err = cfg.Authenticate(req)
			if err != nil {
				t.Fatalf("Authenticate(): unexpected error: %v", err)
			}
			wantAuthHeader := "Bearer databricks-access-token"
			if got := req.Header.Get("Authorization"); got != wantAuthHeader {
				t.Errorf("Authorization header: got %q, want %q", got, wantAuthHeader)
			}
		})
	}
}

// TestOIDCCredentials_CachesAreIsolatedByClient verifies that clients assuming
// different group roles do not reuse each other's WIF access tokens.
func TestOIDCCredentials_CachesAreIsolatedByClient(t *testing.T) {
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
				GroupID:       testCase.groupID,
				TokenAudience: "audience",
				ConfigFile:    "/dev/null",
				HTTPTransport: server.Client().Transport,
			}
			cfg.Credentials = oidcStrategy(cfg, "test-oidc", oidc.IDTokenSourceFn(func(context.Context, string) (*oidc.IDToken, error) {
				return &oidc.IDToken{Value: "id-token"}, nil
			}))

			for range 2 {
				assertHeaders(t, cfg, map[string]string{"Authorization": testCase.want})
			}
		})
	}

	if *tokenCalls != len(testCases) {
		t.Errorf("Token endpoint was called %d times; want one call per client", *tokenCalls)
	}
}
