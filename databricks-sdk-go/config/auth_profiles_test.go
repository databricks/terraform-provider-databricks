package config

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/common/environment"
	"github.com/databricks/databricks-sdk-go/credentials/u2m"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/databricks/databricks-sdk-go/internal/env"
	"golang.org/x/oauth2"
)

// Host profile definitions.
//
// Each profile represents a different Databricks deployment shape:
//   - LW    = Legacy Workspace:  host=workspace_url
//   - NW    = New Workspace:     host=workspace_url + account_id + workspace_id
//   - LA    = Legacy Account:    host=accounts_url + account_id
//   - NA    = New Account:       host=accounts_url + account_id (structurally same as LA)
//   - SPOGW = SPOG workspace:   host=unified_url + account_id + workspace_id
//   - SPOGA = SPOG account:     host=unified_url + account_id

type hostProfile struct {
	name string

	// Config fields for the profile.
	host        string
	accountID   string
	workspaceID string

	// Host metadata returned by the resolver.
	metadata *HostMetadata
}

const (
	testWorkspaceHost = "https://test-workspace.cloud.databricks.com"
	testAccountsHost  = "https://accounts.cloud.databricks.com"
	testUnifiedHost   = "https://test-unified.cloud.databricks.com"
	testAccountID     = "00000000-0000-0000-0000-000000000001"
	testWorkspaceID   = "1234567890"
)

var allProfiles = []hostProfile{
	{
		name: "LW",
		host: testWorkspaceHost,
		metadata: &HostMetadata{
			OIDCEndpoint: testWorkspaceHost + "/oidc",
			WorkspaceID:  testWorkspaceID,
			AccountID:    testAccountID,
			Cloud:        environment.CloudAWS,
			HostType:     WorkspaceHost,
		},
	},
	{
		name:        "NW",
		host:        testWorkspaceHost,
		accountID:   testAccountID,
		workspaceID: testWorkspaceID,
		metadata: &HostMetadata{
			OIDCEndpoint: testWorkspaceHost + "/oidc",
			WorkspaceID:  testWorkspaceID,
			AccountID:    testAccountID,
			Cloud:        environment.CloudAWS,
			HostType:     WorkspaceHost,
		},
	},
	{
		name:      "LA",
		host:      testAccountsHost,
		accountID: testAccountID,
		metadata: &HostMetadata{
			OIDCEndpoint: testAccountsHost + "/oidc/accounts/{account_id}",
			AccountID:    testAccountID,
			Cloud:        environment.CloudAWS,
			HostType:     AccountHost,
		},
	},
	{
		name:      "NA",
		host:      testAccountsHost,
		accountID: testAccountID,
		metadata: &HostMetadata{
			OIDCEndpoint: testAccountsHost + "/oidc/accounts/{account_id}",
			AccountID:    testAccountID,
			Cloud:        environment.CloudAWS,
			HostType:     AccountHost,
		},
	},
	{
		name:        "SPOGW",
		host:        testUnifiedHost,
		accountID:   testAccountID,
		workspaceID: testWorkspaceID,
		metadata: &HostMetadata{
			OIDCEndpoint: testUnifiedHost + "/oidc/accounts/{account_id}",
			AccountID:    testAccountID,
			WorkspaceID:  testWorkspaceID,
			Cloud:        environment.CloudAWS,
			HostType:     UnifiedHost,
		},
	},
	{
		name:      "SPOGA",
		host:      testUnifiedHost,
		accountID: testAccountID,
		metadata: &HostMetadata{
			OIDCEndpoint: testUnifiedHost + "/oidc/accounts/{account_id}",
			AccountID:    testAccountID,
			Cloud:        environment.CloudAWS,
			HostType:     AccountHost,
		},
	},
}

// profileConfig returns a Config pre-populated for the given profile, with the
// host metadata resolver injected. Callers add auth-specific fields and
// HTTPTransport before calling Authenticate.
func profileConfig(p hostProfile) *Config {
	return &Config{
		Host:        p.host,
		AccountID:   p.accountID,
		WorkspaceID: p.workspaceID,
		Loaders:     []Loader{noopLoader{}},
		HostMetadataResolver: func(_ context.Context, _ string) (*HostMetadata, error) {
			return p.metadata, nil
		},
	}
}

// oidcTransportForProfile returns an HTTPTransport that mocks the OIDC
// discovery and token endpoints for OAuth-based auth types.
//
// Host metadata resolution always runs during EnsureResolved, populating
// DiscoveryURL from the metadata OIDCEndpoint. This means getOidcEndpoints
// uses the discovery URL path for all profiles, so every profile needs the
// OIDC discovery endpoint mocked.
func oidcTransportForProfile(p hostProfile) fixtures.MappingTransport {
	if p.metadata.HostType == WorkspaceHost {
		return fixtures.MappingTransport{
			"GET /oidc/.well-known/oauth-authorization-server": {
				Response: u2m.OAuthAuthorizationServer{
					AuthorizationEndpoint: p.host + "/oidc/v1/authorize",
					TokenEndpoint:         p.host + "/oidc/v1/token",
				},
			},
			"POST /oidc/v1/token": {
				Response: oauth2.Token{
					TokenType:   "Bearer",
					AccessToken: "test-token",
				},
			},
		}
	}
	// Account and unified profiles: discovery URL points at account-level OIDC.
	return fixtures.MappingTransport{
		"GET /oidc/accounts/" + testAccountID + "/.well-known/oauth-authorization-server": {
			Response: u2m.OAuthAuthorizationServer{
				AuthorizationEndpoint: p.host + "/oidc/accounts/" + testAccountID + "/v1/authorize",
				TokenEndpoint:         p.host + "/oidc/accounts/" + testAccountID + "/v1/token",
			},
		},
		"POST /oidc/accounts/" + testAccountID + "/v1/token": {
			Response: oauth2.Token{
				TokenType:   "Bearer",
				AccessToken: "test-token",
			},
		},
	}
}

// mergeTransport combines multiple MappingTransport maps.
func mergeTransport(transports ...fixtures.MappingTransport) fixtures.MappingTransport {
	merged := fixtures.MappingTransport{}
	for _, t := range transports {
		for k, v := range t {
			merged[k] = v
		}
	}
	return merged
}

// --- PAT auth ----------------------------------------------------------------

func TestProfileAuth_PAT(t *testing.T) {
	for _, p := range allProfiles {
		t.Run(p.name, func(t *testing.T) {
			cfg := profileConfig(p)
			cfg.Token = "dapi1234567890abcdef"
			cfg.AuthType = "pat"

			req, err := authenticateRequest(cfg)
			if err != nil {
				t.Fatalf("Authenticate(): %v", err)
			}
			if got := req.Header.Get("Authorization"); got != "Bearer dapi1234567890abcdef" {
				t.Errorf("Authorization header: got %q, want %q", got, "Bearer dapi1234567890abcdef")
			}
			if cfg.AuthType != "pat" {
				t.Errorf("AuthType: got %q, want %q", cfg.AuthType, "pat")
			}
		})
	}
}

// --- Basic auth --------------------------------------------------------------

func TestProfileAuth_Basic(t *testing.T) {
	for _, p := range allProfiles {
		t.Run(p.name, func(t *testing.T) {
			cfg := profileConfig(p)
			cfg.Username = "user"
			cfg.Password = "pass"
			cfg.AuthType = "basic"

			req, err := authenticateRequest(cfg)
			if err != nil {
				t.Fatalf("Authenticate(): %v", err)
			}
			got := req.Header.Get("Authorization")
			if got == "" {
				t.Fatal("Authorization header is empty")
			}
			if got != "Basic dXNlcjpwYXNz" {
				t.Errorf("Authorization header: got %q, want %q", got, "Basic dXNlcjpwYXNz")
			}
			if cfg.AuthType != "basic" {
				t.Errorf("AuthType: got %q, want %q", cfg.AuthType, "basic")
			}
		})
	}
}

// --- OAuth M2M ---------------------------------------------------------------

func TestProfileAuth_OAuthM2M(t *testing.T) {
	for _, p := range allProfiles {
		t.Run(p.name, func(t *testing.T) {
			cfg := profileConfig(p)
			cfg.ClientID = "test-client"
			cfg.ClientSecret = "test-secret"
			cfg.AuthType = "oauth-m2m"
			cfg.HTTPTransport = oidcTransportForProfile(p)

			req, err := authenticateRequest(cfg)
			if err != nil {
				t.Fatalf("Authenticate(): %v", err)
			}
			if got := req.Header.Get("Authorization"); got != "Bearer test-token" {
				t.Errorf("Authorization header: got %q, want %q", got, "Bearer test-token")
			}
			if cfg.AuthType != "oauth-m2m" {
				t.Errorf("AuthType: got %q, want %q", cfg.AuthType, "oauth-m2m")
			}
		})
	}
}

// --- GitHub OIDC -------------------------------------------------------------

func TestProfileAuth_GitHubOIDC(t *testing.T) {
	for _, p := range allProfiles {
		t.Run(p.name, func(t *testing.T) {
			cfg := profileConfig(p)
			cfg.ClientID = "test-client"
			cfg.AuthType = "github-oidc"
			cfg.ActionsIDTokenRequestURL = "http://github-actions.test/token?version=1"
			cfg.ActionsIDTokenRequestToken = "github-request-token"
			cfg.TokenAudience = "databricks-test-audience"

			githubTransport := fixtures.MappingTransport{
				"GET /token?version=1&audience=databricks-test-audience": {
					Response: map[string]string{"value": "github-id-token"},
				},
			}
			cfg.HTTPTransport = mergeTransport(oidcTransportForProfile(p), githubTransport)

			req, err := authenticateRequest(cfg)
			if err != nil {
				t.Fatalf("Authenticate(): %v", err)
			}
			if got := req.Header.Get("Authorization"); got != "Bearer test-token" {
				t.Errorf("Authorization header: got %q, want %q", got, "Bearer test-token")
			}
			if cfg.AuthType != "github-oidc" {
				t.Errorf("AuthType: got %q, want %q", cfg.AuthType, "github-oidc")
			}
		})
	}
}

// --- Env OIDC ----------------------------------------------------------------

func TestProfileAuth_EnvOIDC(t *testing.T) {
	for _, p := range allProfiles {
		t.Run(p.name, func(t *testing.T) {
			t.Setenv("DATABRICKS_OIDC_TOKEN", "test-oidc-token")

			cfg := profileConfig(p)
			cfg.ClientID = "test-client"
			cfg.AuthType = "env-oidc"
			cfg.HTTPTransport = oidcTransportForProfile(p)

			req, err := authenticateRequest(cfg)
			if err != nil {
				t.Fatalf("Authenticate(): %v", err)
			}
			if got := req.Header.Get("Authorization"); got != "Bearer test-token" {
				t.Errorf("Authorization header: got %q, want %q", got, "Bearer test-token")
			}
			if cfg.AuthType != "env-oidc" {
				t.Errorf("AuthType: got %q, want %q", cfg.AuthType, "env-oidc")
			}
		})
	}
}

// --- File OIDC ---------------------------------------------------------------

func TestProfileAuth_FileOIDC(t *testing.T) {
	for _, p := range allProfiles {
		t.Run(p.name, func(t *testing.T) {
			tokenFile := filepath.Join(t.TempDir(), "oidc_token")
			if err := os.WriteFile(tokenFile, []byte("test-oidc-token"), 0o600); err != nil {
				t.Fatal(err)
			}

			cfg := profileConfig(p)
			cfg.ClientID = "test-client"
			cfg.AuthType = "file-oidc"
			cfg.OIDCTokenFilepath = tokenFile
			cfg.HTTPTransport = oidcTransportForProfile(p)

			req, err := authenticateRequest(cfg)
			if err != nil {
				t.Fatalf("Authenticate(): %v", err)
			}
			if got := req.Header.Get("Authorization"); got != "Bearer test-token" {
				t.Errorf("Authorization header: got %q, want %q", got, "Bearer test-token")
			}
			if cfg.AuthType != "file-oidc" {
				t.Errorf("AuthType: got %q, want %q", cfg.AuthType, "file-oidc")
			}
		})
	}
}

// --- Metadata Service (workspace profiles only) ------------------------------

func TestProfileAuth_MetadataService(t *testing.T) {
	// Metadata service auth is only available for workspace-level profiles.
	var workspaceProfiles []hostProfile
	for _, p := range allProfiles {
		if p.metadata.HostType == WorkspaceHost || p.metadata.HostType == UnifiedHost {
			workspaceProfiles = append(workspaceProfiles, p)
		}
	}
	for _, p := range workspaceProfiles {
		t.Run(p.name, func(t *testing.T) {
			cfg := profileConfig(p)
			cfg.MetadataServiceURL = "http://localhost:1234/metadata/token"
			cfg.AuthType = "metadata-service"
			cfg.HTTPTransport = fixtures.MappingTransport{
				"GET /metadata/token": {
					ExpectedHeaders: map[string]string{
						"Accept":                        "application/json",
						"X-Databricks-Host":             cfg.Host,
						"X-Databricks-Metadata-Version": "1",
					},
					Response: someValidToken("metadata-token"),
				},
			}

			req, err := authenticateRequest(cfg)
			if err != nil {
				t.Fatalf("Authenticate(): %v", err)
			}
			if got := req.Header.Get("Authorization"); got != "Bearer metadata-token" {
				t.Errorf("Authorization header: got %q, want %q", got, "Bearer metadata-token")
			}
			if cfg.AuthType != "metadata-service" {
				t.Errorf("AuthType: got %q, want %q", cfg.AuthType, "metadata-service")
			}
		})
	}
}

// --- Azure profiles -----------------------------------------------------------

const (
	testAzureWorkspaceHost = "https://adb-1234567890.12.azuredatabricks.net"
	testAzureAccountsHost  = "https://accounts.azuredatabricks.net"
	testAzureUnifiedHost   = "https://db-test.azuredatabricks.net"
)

var azureProfiles = []hostProfile{
	{
		name: "AZ_LW",
		host: testAzureWorkspaceHost,
		metadata: &HostMetadata{
			OIDCEndpoint: testAzureWorkspaceHost + "/oidc",
			WorkspaceID:  testWorkspaceID,
			AccountID:    testAccountID,
			Cloud:        environment.CloudAzure,
			HostType:     WorkspaceHost,
		},
	},
	{
		name:        "AZ_NW",
		host:        testAzureWorkspaceHost,
		accountID:   testAccountID,
		workspaceID: testWorkspaceID,
		metadata: &HostMetadata{
			OIDCEndpoint: testAzureWorkspaceHost + "/oidc",
			WorkspaceID:  testWorkspaceID,
			AccountID:    testAccountID,
			Cloud:        environment.CloudAzure,
			HostType:     WorkspaceHost,
		},
	},
	{
		name:      "AZ_LA",
		host:      testAzureAccountsHost,
		accountID: testAccountID,
		metadata: &HostMetadata{
			OIDCEndpoint: testAzureAccountsHost + "/oidc/accounts/{account_id}",
			AccountID:    testAccountID,
			Cloud:        environment.CloudAzure,
			HostType:     AccountHost,
		},
	},
	{
		name:      "AZ_NA",
		host:      testAzureAccountsHost,
		accountID: testAccountID,
		metadata: &HostMetadata{
			OIDCEndpoint: testAzureAccountsHost + "/oidc/accounts/{account_id}",
			AccountID:    testAccountID,
			Cloud:        environment.CloudAzure,
			HostType:     AccountHost,
		},
	},
	{
		name:        "AZ_SPOGW",
		host:        testAzureUnifiedHost,
		accountID:   testAccountID,
		workspaceID: testWorkspaceID,
		metadata: &HostMetadata{
			OIDCEndpoint: testAzureUnifiedHost + "/oidc/accounts/{account_id}",
			AccountID:    testAccountID,
			WorkspaceID:  testWorkspaceID,
			Cloud:        environment.CloudAzure,
			HostType:     UnifiedHost,
		},
	},
	{
		name:      "AZ_SPOGA",
		host:      testAzureUnifiedHost,
		accountID: testAccountID,
		metadata: &HostMetadata{
			OIDCEndpoint: testAzureUnifiedHost + "/oidc/accounts/{account_id}",
			AccountID:    testAccountID,
			Cloud:        environment.CloudAzure,
			HostType:     AccountHost,
		},
	},
}

// --- Azure Client Secret -----------------------------------------------------

func TestProfileAuth_AzureClientSecret(t *testing.T) {
	tokenExpiry := time.Now().Add(time.Hour).Unix()
	for _, p := range azureProfiles {
		t.Run(p.name, func(t *testing.T) {
			cfg := profileConfig(p)
			cfg.AzureClientID = "test-azure-client"
			cfg.AzureClientSecret = "test-azure-secret"
			cfg.AzureTenantID = "test-tenant-id"
			cfg.AuthType = "azure-client-secret"
			cfg.HTTPTransport = fixtures.MappingTransport{
				"POST /test-tenant-id/oauth2/token": {
					Response: map[string]any{
						"token_type":   "Bearer",
						"access_token": "workspace-token",
						"expires_on":   fmt.Sprintf("%d", tokenExpiry),
					},
				},
			}

			req, err := authenticateRequest(cfg)
			if err != nil {
				t.Fatalf("Authenticate(): %v", err)
			}
			if got := req.Header.Get("Authorization"); got != "Bearer workspace-token" {
				t.Errorf("Authorization header: got %q, want %q", got, "Bearer workspace-token")
			}
			if cfg.AuthType != "azure-client-secret" {
				t.Errorf("AuthType: got %q, want %q", cfg.AuthType, "azure-client-secret")
			}
		})
	}
}

// --- Azure GitHub OIDC -------------------------------------------------------

func TestProfileAuth_AzureGithubOIDC(t *testing.T) {
	for _, p := range azureProfiles {
		t.Run(p.name, func(t *testing.T) {
			cfg := profileConfig(p)
			cfg.AzureClientID = "test-azure-client"
			cfg.AzureTenantID = "test-tenant-id"
			cfg.ActionsIDTokenRequestURL = "http://github-actions.test/token?version=1"
			cfg.ActionsIDTokenRequestToken = "github-request-token"
			cfg.AuthType = "github-oidc-azure"
			cfg.DatabricksEnvironment = &environment.DatabricksEnvironment{
				Cloud:              environment.CloudAzure,
				AzureApplicationID: "test-azure-app-id",
				AzureEnvironment:   &environment.AzurePublicCloud,
			}
			cfg.HTTPTransport = fixtures.MappingTransport{
				"GET /token?version=1&audience=api://AzureADTokenExchange": {
					Response: `{"value": "github-id-token"}`,
				},
				"POST /test-tenant-id/oauth2/token": {
					Response: map[string]string{
						"token_type":    "Bearer",
						"access_token":  "azure-token",
						"refresh_token": "refresh",
						"expires_on":    "0",
					},
				},
			}

			req, err := authenticateRequest(cfg)
			if err != nil {
				t.Fatalf("Authenticate(): %v", err)
			}
			if got := req.Header.Get("Authorization"); got != "Bearer azure-token" {
				t.Errorf("Authorization header: got %q, want %q", got, "Bearer azure-token")
			}
			if cfg.AuthType != "github-oidc-azure" {
				t.Errorf("AuthType: got %q, want %q", cfg.AuthType, "github-oidc-azure")
			}
		})
	}
}

// --- GCP profiles ------------------------------------------------------------

const (
	testGCPWorkspaceHost = "https://1234567890.1.gcp.databricks.com"
	testGCPAccountsHost  = "https://accounts.gcp.databricks.com"
	testGCPUnifiedHost   = "https://db-test.gcp.databricks.com"
)

var gcpProfiles = []hostProfile{
	{
		name: "GCP_LW",
		host: testGCPWorkspaceHost,
		metadata: &HostMetadata{
			OIDCEndpoint: testGCPWorkspaceHost + "/oidc",
			WorkspaceID:  testWorkspaceID,
			AccountID:    testAccountID,
			Cloud:        environment.CloudGCP,
			HostType:     WorkspaceHost,
		},
	},
	{
		name:        "GCP_NW",
		host:        testGCPWorkspaceHost,
		accountID:   testAccountID,
		workspaceID: testWorkspaceID,
		metadata: &HostMetadata{
			OIDCEndpoint: testGCPWorkspaceHost + "/oidc",
			WorkspaceID:  testWorkspaceID,
			AccountID:    testAccountID,
			Cloud:        environment.CloudGCP,
			HostType:     WorkspaceHost,
		},
	},
	{
		name:      "GCP_LA",
		host:      testGCPAccountsHost,
		accountID: testAccountID,
		metadata: &HostMetadata{
			OIDCEndpoint: testGCPAccountsHost + "/oidc/accounts/{account_id}",
			AccountID:    testAccountID,
			Cloud:        environment.CloudGCP,
			HostType:     AccountHost,
		},
	},
	{
		name:      "GCP_NA",
		host:      testGCPAccountsHost,
		accountID: testAccountID,
		metadata: &HostMetadata{
			OIDCEndpoint: testGCPAccountsHost + "/oidc/accounts/{account_id}",
			AccountID:    testAccountID,
			Cloud:        environment.CloudGCP,
			HostType:     AccountHost,
		},
	},
	{
		name:        "GCP_SPOGW",
		host:        testGCPUnifiedHost,
		accountID:   testAccountID,
		workspaceID: testWorkspaceID,
		metadata: &HostMetadata{
			OIDCEndpoint: testGCPUnifiedHost + "/oidc/accounts/{account_id}",
			AccountID:    testAccountID,
			WorkspaceID:  testWorkspaceID,
			Cloud:        environment.CloudGCP,
			HostType:     UnifiedHost,
		},
	},
	{
		name:      "GCP_SPOGA",
		host:      testGCPUnifiedHost,
		accountID: testAccountID,
		metadata: &HostMetadata{
			OIDCEndpoint: testGCPUnifiedHost + "/oidc/accounts/{account_id}",
			AccountID:    testAccountID,
			Cloud:        environment.CloudGCP,
			HostType:     AccountHost,
		},
	},
}

// --- PAT across all clouds ---------------------------------------------------

func TestProfileAuth_PAT_Azure(t *testing.T) {
	for _, p := range azureProfiles {
		t.Run(p.name, func(t *testing.T) {
			cfg := profileConfig(p)
			cfg.Token = "dapi1234567890abcdef"
			cfg.AuthType = "pat"

			req, err := authenticateRequest(cfg)
			if err != nil {
				t.Fatalf("Authenticate(): %v", err)
			}
			if got := req.Header.Get("Authorization"); got != "Bearer dapi1234567890abcdef" {
				t.Errorf("Authorization header: got %q, want %q", got, "Bearer dapi1234567890abcdef")
			}
		})
	}
}

func TestProfileAuth_PAT_GCP(t *testing.T) {
	for _, p := range gcpProfiles {
		t.Run(p.name, func(t *testing.T) {
			cfg := profileConfig(p)
			cfg.Token = "dapi1234567890abcdef"
			cfg.AuthType = "pat"

			req, err := authenticateRequest(cfg)
			if err != nil {
				t.Fatalf("Authenticate(): %v", err)
			}
			if got := req.Header.Get("Authorization"); got != "Bearer dapi1234567890abcdef" {
				t.Errorf("Authorization header: got %q, want %q", got, "Bearer dapi1234567890abcdef")
			}
		})
	}
}

// --- Basic across all clouds -------------------------------------------------

func TestProfileAuth_Basic_Azure(t *testing.T) {
	for _, p := range azureProfiles {
		t.Run(p.name, func(t *testing.T) {
			cfg := profileConfig(p)
			cfg.Username = "user"
			cfg.Password = "pass"
			cfg.AuthType = "basic"

			req, err := authenticateRequest(cfg)
			if err != nil {
				t.Fatalf("Authenticate(): %v", err)
			}
			if got := req.Header.Get("Authorization"); got != "Basic dXNlcjpwYXNz" {
				t.Errorf("Authorization header: got %q, want %q", got, "Basic dXNlcjpwYXNz")
			}
		})
	}
}

func TestProfileAuth_Basic_GCP(t *testing.T) {
	for _, p := range gcpProfiles {
		t.Run(p.name, func(t *testing.T) {
			cfg := profileConfig(p)
			cfg.Username = "user"
			cfg.Password = "pass"
			cfg.AuthType = "basic"

			req, err := authenticateRequest(cfg)
			if err != nil {
				t.Fatalf("Authenticate(): %v", err)
			}
			if got := req.Header.Get("Authorization"); got != "Basic dXNlcjpwYXNz" {
				t.Errorf("Authorization header: got %q, want %q", got, "Basic dXNlcjpwYXNz")
			}
		})
	}
}

// --- OAuth M2M across all clouds ---------------------------------------------

func oidcTransportForAzureProfile(p hostProfile) fixtures.MappingTransport {
	if p.metadata.HostType == WorkspaceHost {
		return fixtures.MappingTransport{
			"GET /oidc/.well-known/oauth-authorization-server": {
				Response: u2m.OAuthAuthorizationServer{
					AuthorizationEndpoint: p.host + "/oidc/v1/authorize",
					TokenEndpoint:         p.host + "/oidc/v1/token",
				},
			},
			"POST /oidc/v1/token": {
				Response: oauth2.Token{
					TokenType:   "Bearer",
					AccessToken: "test-token",
				},
			},
		}
	}
	return fixtures.MappingTransport{
		"GET /oidc/accounts/" + testAccountID + "/.well-known/oauth-authorization-server": {
			Response: u2m.OAuthAuthorizationServer{
				AuthorizationEndpoint: p.host + "/oidc/accounts/" + testAccountID + "/v1/authorize",
				TokenEndpoint:         p.host + "/oidc/accounts/" + testAccountID + "/v1/token",
			},
		},
		"POST /oidc/accounts/" + testAccountID + "/v1/token": {
			Response: oauth2.Token{
				TokenType:   "Bearer",
				AccessToken: "test-token",
			},
		},
	}
}

func oidcTransportForGCPProfile(p hostProfile) fixtures.MappingTransport {
	if p.metadata.HostType == WorkspaceHost {
		return fixtures.MappingTransport{
			"GET /oidc/.well-known/oauth-authorization-server": {
				Response: u2m.OAuthAuthorizationServer{
					AuthorizationEndpoint: p.host + "/oidc/v1/authorize",
					TokenEndpoint:         p.host + "/oidc/v1/token",
				},
			},
			"POST /oidc/v1/token": {
				Response: oauth2.Token{
					TokenType:   "Bearer",
					AccessToken: "test-token",
				},
			},
		}
	}
	return fixtures.MappingTransport{
		"GET /oidc/accounts/" + testAccountID + "/.well-known/oauth-authorization-server": {
			Response: u2m.OAuthAuthorizationServer{
				AuthorizationEndpoint: p.host + "/oidc/accounts/" + testAccountID + "/v1/authorize",
				TokenEndpoint:         p.host + "/oidc/accounts/" + testAccountID + "/v1/token",
			},
		},
		"POST /oidc/accounts/" + testAccountID + "/v1/token": {
			Response: oauth2.Token{
				TokenType:   "Bearer",
				AccessToken: "test-token",
			},
		},
	}
}

func TestProfileAuth_OAuthM2M_Azure(t *testing.T) {
	for _, p := range azureProfiles {
		t.Run(p.name, func(t *testing.T) {
			cfg := profileConfig(p)
			cfg.ClientID = "test-client"
			cfg.ClientSecret = "test-secret"
			cfg.AuthType = "oauth-m2m"
			cfg.HTTPTransport = oidcTransportForAzureProfile(p)

			req, err := authenticateRequest(cfg)
			if err != nil {
				t.Fatalf("Authenticate(): %v", err)
			}
			if got := req.Header.Get("Authorization"); got != "Bearer test-token" {
				t.Errorf("Authorization header: got %q, want %q", got, "Bearer test-token")
			}
		})
	}
}

func TestProfileAuth_OAuthM2M_GCP(t *testing.T) {
	for _, p := range gcpProfiles {
		t.Run(p.name, func(t *testing.T) {
			cfg := profileConfig(p)
			cfg.ClientID = "test-client"
			cfg.ClientSecret = "test-secret"
			cfg.AuthType = "oauth-m2m"
			cfg.HTTPTransport = oidcTransportForGCPProfile(p)

			req, err := authenticateRequest(cfg)
			if err != nil {
				t.Fatalf("Authenticate(): %v", err)
			}
			if got := req.Header.Get("Authorization"); got != "Bearer test-token" {
				t.Errorf("Authorization header: got %q, want %q", got, "Bearer test-token")
			}
		})
	}
}

// --- GitHub OIDC across all clouds -------------------------------------------

func TestProfileAuth_GitHubOIDC_Azure(t *testing.T) {
	for _, p := range azureProfiles {
		t.Run(p.name, func(t *testing.T) {
			cfg := profileConfig(p)
			cfg.ClientID = "test-client"
			cfg.AuthType = "github-oidc"
			cfg.ActionsIDTokenRequestURL = "http://github-actions.test/token?version=1"
			cfg.ActionsIDTokenRequestToken = "github-request-token"
			cfg.TokenAudience = "databricks-test-audience"

			githubTransport := fixtures.MappingTransport{
				"GET /token?version=1&audience=databricks-test-audience": {
					Response: map[string]string{"value": "github-id-token"},
				},
			}
			cfg.HTTPTransport = mergeTransport(oidcTransportForAzureProfile(p), githubTransport)

			req, err := authenticateRequest(cfg)
			if err != nil {
				t.Fatalf("Authenticate(): %v", err)
			}
			if got := req.Header.Get("Authorization"); got != "Bearer test-token" {
				t.Errorf("Authorization header: got %q, want %q", got, "Bearer test-token")
			}
		})
	}
}

func TestProfileAuth_GitHubOIDC_GCP(t *testing.T) {
	for _, p := range gcpProfiles {
		t.Run(p.name, func(t *testing.T) {
			cfg := profileConfig(p)
			cfg.ClientID = "test-client"
			cfg.AuthType = "github-oidc"
			cfg.ActionsIDTokenRequestURL = "http://github-actions.test/token?version=1"
			cfg.ActionsIDTokenRequestToken = "github-request-token"
			cfg.TokenAudience = "databricks-test-audience"

			githubTransport := fixtures.MappingTransport{
				"GET /token?version=1&audience=databricks-test-audience": {
					Response: map[string]string{"value": "github-id-token"},
				},
			}
			cfg.HTTPTransport = mergeTransport(oidcTransportForGCPProfile(p), githubTransport)

			req, err := authenticateRequest(cfg)
			if err != nil {
				t.Fatalf("Authenticate(): %v", err)
			}
			if got := req.Header.Get("Authorization"); got != "Bearer test-token" {
				t.Errorf("Authorization header: got %q, want %q", got, "Bearer test-token")
			}
		})
	}
}

// --- Env OIDC across all clouds ----------------------------------------------

func TestProfileAuth_EnvOIDC_Azure(t *testing.T) {
	for _, p := range azureProfiles {
		t.Run(p.name, func(t *testing.T) {
			t.Setenv("DATABRICKS_OIDC_TOKEN", "test-oidc-token")

			cfg := profileConfig(p)
			cfg.ClientID = "test-client"
			cfg.AuthType = "env-oidc"
			cfg.HTTPTransport = oidcTransportForAzureProfile(p)

			req, err := authenticateRequest(cfg)
			if err != nil {
				t.Fatalf("Authenticate(): %v", err)
			}
			if got := req.Header.Get("Authorization"); got != "Bearer test-token" {
				t.Errorf("Authorization header: got %q, want %q", got, "Bearer test-token")
			}
		})
	}
}

func TestProfileAuth_EnvOIDC_GCP(t *testing.T) {
	for _, p := range gcpProfiles {
		t.Run(p.name, func(t *testing.T) {
			t.Setenv("DATABRICKS_OIDC_TOKEN", "test-oidc-token")

			cfg := profileConfig(p)
			cfg.ClientID = "test-client"
			cfg.AuthType = "env-oidc"
			cfg.HTTPTransport = oidcTransportForGCPProfile(p)

			req, err := authenticateRequest(cfg)
			if err != nil {
				t.Fatalf("Authenticate(): %v", err)
			}
			if got := req.Header.Get("Authorization"); got != "Bearer test-token" {
				t.Errorf("Authorization header: got %q, want %q", got, "Bearer test-token")
			}
		})
	}
}

// --- File OIDC across all clouds ---------------------------------------------

func TestProfileAuth_FileOIDC_Azure(t *testing.T) {
	for _, p := range azureProfiles {
		t.Run(p.name, func(t *testing.T) {
			tokenFile := filepath.Join(t.TempDir(), "oidc_token")
			if err := os.WriteFile(tokenFile, []byte("test-oidc-token"), 0o600); err != nil {
				t.Fatal(err)
			}

			cfg := profileConfig(p)
			cfg.ClientID = "test-client"
			cfg.AuthType = "file-oidc"
			cfg.OIDCTokenFilepath = tokenFile
			cfg.HTTPTransport = oidcTransportForAzureProfile(p)

			req, err := authenticateRequest(cfg)
			if err != nil {
				t.Fatalf("Authenticate(): %v", err)
			}
			if got := req.Header.Get("Authorization"); got != "Bearer test-token" {
				t.Errorf("Authorization header: got %q, want %q", got, "Bearer test-token")
			}
		})
	}
}

func TestProfileAuth_FileOIDC_GCP(t *testing.T) {
	for _, p := range gcpProfiles {
		t.Run(p.name, func(t *testing.T) {
			tokenFile := filepath.Join(t.TempDir(), "oidc_token")
			if err := os.WriteFile(tokenFile, []byte("test-oidc-token"), 0o600); err != nil {
				t.Fatal(err)
			}

			cfg := profileConfig(p)
			cfg.ClientID = "test-client"
			cfg.AuthType = "file-oidc"
			cfg.OIDCTokenFilepath = tokenFile
			cfg.HTTPTransport = oidcTransportForGCPProfile(p)

			req, err := authenticateRequest(cfg)
			if err != nil {
				t.Fatalf("Authenticate(): %v", err)
			}
			if got := req.Header.Get("Authorization"); got != "Bearer test-token" {
				t.Errorf("Authorization header: got %q, want %q", got, "Bearer test-token")
			}
		})
	}
}

// --- Databricks CLI ----------------------------------------------------------

func TestProfileAuth_DatabricksCli(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skipping on Windows")
	}
	for _, p := range allProfiles {
		t.Run(p.name, func(t *testing.T) {
			cfg := profileConfig(p)
			cfg.AuthType = "databricks-cli"
			cfg.DatabricksCliPath = createMockCli(t, mockCliTokenResponse(t))

			req, err := authenticateRequest(cfg)
			if err != nil {
				t.Fatalf("Authenticate(): %v", err)
			}
			if got := req.Header.Get("Authorization"); got != "Bearer test-token" {
				t.Errorf("Authorization header: got %q, want %q", got, "Bearer test-token")
			}
			if cfg.AuthType != "databricks-cli" {
				t.Errorf("AuthType: got %q, want %q", cfg.AuthType, "databricks-cli")
			}
		})
	}
}

// --- Azure CLI ---------------------------------------------------------------

func TestProfileAuth_AzureCli(t *testing.T) {
	for _, p := range azureProfiles {
		t.Run(p.name, func(t *testing.T) {
			env.CleanupEnvironment(t)
			os.Setenv("PATH", testdataPath())

			cfg := profileConfig(p)
			cfg.AuthType = "azure-cli"
			cfg.azureTenantIdFetchClient = makeClient(&http.Response{
				StatusCode: http.StatusTemporaryRedirect,
				Header:     http.Header{"Location": []string{"https://login.microsoftonline.com/test-tenant-id/oauth2/token"}},
			})

			req, err := authenticateRequest(cfg)
			if err != nil {
				t.Fatalf("Authenticate(): %v", err)
			}
			if got := req.Header.Get("Authorization"); got == "" {
				t.Error("Authorization header is empty")
			}
			if cfg.AuthType != "azure-cli" {
				t.Errorf("AuthType: got %q, want %q", cfg.AuthType, "azure-cli")
			}
		})
	}
}

// --- Azure DevOps OIDC -------------------------------------------------------

func TestProfileAuth_AzureDevOpsOIDC(t *testing.T) {
	for _, p := range azureProfiles {
		t.Run(p.name, func(t *testing.T) {
			env.CleanupEnvironment(t)
			os.Setenv("SYSTEM_ACCESSTOKEN", "test-access-token")
			os.Setenv("SYSTEM_TEAMFOUNDATIONCOLLECTIONURI", "http://devops.test")
			os.Setenv("SYSTEM_PLANID", "plan-123")
			os.Setenv("SYSTEM_JOBID", "job-456")
			os.Setenv("SYSTEM_TEAMPROJECTID", "project-789")
			os.Setenv("SYSTEM_HOSTTYPE", "build")

			devOpsTransport := fixtures.MappingTransport{
				"POST /project-789/_apis/distributedtask/hubs/build/plans/plan-123/jobs/job-456/oidctoken?api-version=7.2-preview.1": {
					Response: map[string]string{"oidcToken": "devops-oidc-token"},
				},
			}

			cfg := profileConfig(p)
			cfg.ClientID = "test-client"
			cfg.AuthType = "azure-devops-oidc"
			cfg.HTTPTransport = mergeTransport(oidcTransportForAzureProfile(p), devOpsTransport)

			req, err := authenticateRequest(cfg)
			if err != nil {
				t.Fatalf("Authenticate(): %v", err)
			}
			if got := req.Header.Get("Authorization"); got != "Bearer test-token" {
				t.Errorf("Authorization header: got %q, want %q", got, "Bearer test-token")
			}
			if cfg.AuthType != "azure-devops-oidc" {
				t.Errorf("AuthType: got %q, want %q", cfg.AuthType, "azure-devops-oidc")
			}
		})
	}
}
