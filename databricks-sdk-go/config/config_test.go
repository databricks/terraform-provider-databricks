package config

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/databricks/databricks-sdk-go/common/environment"
	"github.com/databricks/databricks-sdk-go/credentials/u2m"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockLoader is a test helper that implements the Loader interface.
type mockLoader func(cfg *Config) error

func (m mockLoader) Name() string {
	return "mockLoader"
}

func (m mockLoader) Configure(cfg *Config) error {
	return m(cfg)
}

// metadataNotFoundTransport returns an HTTPTransport that returns 404 for
// the host metadata endpoint, ensuring tests don't make real HTTP calls.
var metadataNotFoundTransport = fixtures.SliceTransport{
	{
		Method:       "GET",
		Resource:     "/.well-known/databricks-config",
		ReuseRequest: true,
		Status:       404,
	},
}

func TestHostType_AwsAccount(t *testing.T) {
	c := &Config{
		Host:      "https://accounts.cloud.databricks.com",
		AccountID: "123e4567-e89b-12d3-a456-426614174000",
	}
	assert.Equal(t, AccountHost, c.HostType())
}

func TestHostType_AwsDodAccount(t *testing.T) {
	c := &Config{
		Host:      "https://accounts-dod.cloud.databricks.us",
		AccountID: "123e4567-e89b-12d3-a456-426614174000",
	}
	assert.Equal(t, AccountHost, c.HostType())
}

func TestHostType_AwsAccountWithoutScheme(t *testing.T) {
	c := &Config{
		Host:      "accounts.cloud.databricks.com",
		AccountID: "123e4567-e89b-12d3-a456-426614174000",
	}
	assert.Equal(t, AccountHost, c.HostType())
}

func TestHostType_AwsDodAccountWithoutScheme(t *testing.T) {
	c := &Config{
		Host:      "accounts-dod.cloud.databricks.us",
		AccountID: "123e4567-e89b-12d3-a456-426614174000",
	}
	assert.Equal(t, AccountHost, c.HostType())
}

func TestHostType_AwsWorkspace(t *testing.T) {
	c := &Config{
		Host:      "https://my-workspace.cloud.databricks.us",
		AccountID: "123e4567-e89b-12d3-a456-426614174000",
	}
	assert.Equal(t, WorkspaceHost, c.HostType())
}

func TestHostType_NonAccountsHostIsWorkspace(t *testing.T) {
	c := &Config{
		Host:      "https://unified.cloud.databricks.com",
		AccountID: "123e4567-e89b-12d3-a456-426614174000",
	}
	// Host type is determined by URL pattern; a non-accounts host is a
	// workspace host even when AccountID is set.
	assert.Equal(t, WorkspaceHost, c.HostType())
}

func TestHostType_UsesMetadataFirst_Workspace(t *testing.T) {
	c := &Config{
		Host:             "https://accounts.cloud.databricks.com",
		resolvedHostType: WorkspaceHost,
	}
	assert.Equal(t, WorkspaceHost, c.HostType())
}

func TestHostType_UsesMetadataFirst_Account(t *testing.T) {
	c := &Config{
		Host:             "https://my-workspace.cloud.databricks.com",
		resolvedHostType: AccountHost,
	}
	assert.Equal(t, AccountHost, c.HostType())
}

func TestHostType_UsesMetadataFirst_Unified(t *testing.T) {
	c := &Config{
		Host:             "https://my-workspace.cloud.databricks.com",
		resolvedHostType: UnifiedHost,
	}
	assert.Equal(t, UnifiedHost, c.HostType())
}

func TestHostType_FallsBackToURLWhenMetadataUnknown(t *testing.T) {
	c := &Config{
		Host:             "https://accounts.cloud.databricks.com",
		resolvedHostType: HostTypeUnknown,
	}
	assert.Equal(t, AccountHost, c.HostType())
}

func TestHostType_FallsBackToURLWhenMetadataUnknown_Workspace(t *testing.T) {
	c := &Config{
		Host:             "https://my-workspace.cloud.databricks.com",
		resolvedHostType: HostTypeUnknown,
	}
	assert.Equal(t, WorkspaceHost, c.HostType())
}

func TestHostType_EndToEnd_MetadataResolvesHostType(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    "https://my-workspace.cloud.databricks.com",
		Loaders: []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       200,
				Response:     `{"oidc_endpoint": "https://my-workspace.cloud.databricks.com/oidc", "account_id": "` + testHMAccountID + `", "host_type": "account"}`,
			},
		},
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	assert.Equal(t, AccountHost, cfg.HostType())
}

func TestIsAccountClient_DoesNotPanicOnNonAccountsHost(t *testing.T) {
	c := &Config{
		Host:      "https://unified.cloud.databricks.com",
		AccountID: "test-account",
	}
	assert.NotPanics(t, func() { c.IsAccountClient() })
}

func TestNewWithWorkspaceHost(t *testing.T) {
	c := &Config{
		Host:               "https://accounts.cloud.databricks.com",
		AccountID:          "123e4567-e89b-12d3-a456-426614174000",
		AzureResourceID:    "/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Databricks/workspaces/test",
		ClientID:           "client-id",
		MetadataServiceURL: "http://",
		resolved:           true,
	}
	c2, err := c.NewWithWorkspaceHost("https://my-workspace.cloud.databricks.us")
	assert.NoError(t, err)
	// Host should be updated
	assert.Equal(t, "https://my-workspace.cloud.databricks.us", c2.Host)
	// Account ID and Azure Resource ID should be cleared
	assert.Equal(t, "", c2.AccountID)
	assert.Equal(t, "", c2.AzureResourceID)
	// Other fields should be preserved
	assert.Equal(t, "client-id", c2.ClientID)
	assert.Equal(t, "http://", c2.MetadataServiceURL)
	// The new config will not be automatically resolved.
	assert.False(t, c2.resolved)
}

func TestAuthenticate_InvalidHostSet(t *testing.T) {
	c := &Config{
		Host:  "https://:443",
		Token: "abcdefg",
	}
	req, err := http.NewRequestWithContext(context.Background(), "GET", c.Host, nil)
	require.NoError(t, err)
	err = c.Authenticate(req)
	assert.ErrorIs(t, err, ErrNoHostConfigured)
}

// TestAuthenticateIfNeeded_concurrentLazyInit aims at exercising
// authenticateIfNeeded in parallel to catch potential race conditions when
// running the test with -race (see #1310).
func TestAuthenticateIfNeeded_concurrentLazyInit(t *testing.T) {
	cfg := &Config{
		Host:          "http://localhost",
		Token:         "x",
		Loaders:       []Loader{mockLoader(func(*Config) error { return nil })},
		HTTPTransport: metadataNotFoundTransport,
	}
	if err := cfg.EnsureResolved(); err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	for range 32 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cfg.authenticateIfNeeded()
		}()
	}
	wg.Wait()
}

func TestConfig_getOidcEndpoints_account(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	tests := []struct {
		name      string
		host      string
		accountID string
	}{
		{
			name:      "without trailing slash",
			host:      "https://accounts.cloud.databricks.com",
			accountID: "abc",
		},
		{
			name:      "with trailing slash",
			host:      "https://accounts.cloud.databricks.com/",
			accountID: "abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Host:          tt.host,
				AccountID:     tt.accountID,
				Loaders:       []Loader{noopLoader},
				HTTPTransport: metadataNotFoundTransport,
			}
			got, err := c.getOidcEndpoints(context.Background())
			assert.NoError(t, err)
			assert.Equal(t, &u2m.OAuthAuthorizationServer{
				AuthorizationEndpoint: "https://accounts.cloud.databricks.com/oidc/accounts/abc/v1/authorize",
				TokenEndpoint:         "https://accounts.cloud.databricks.com/oidc/accounts/abc/v1/token",
			}, got)
		})
	}
}

func TestConfig_getOidcEndpoints_workspace(t *testing.T) {
	tests := []struct {
		name string
		host string
	}{
		{
			name: "without trailing slash",
			host: "https://myworkspace.cloud.databricks.com",
		},
		{
			name: "with trailing slash",
			host: "https://myworkspace.cloud.databricks.com/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Host: tt.host,
				HTTPTransport: fixtures.SliceTransport{
					{
						Method:   "GET",
						Resource: "/oidc/.well-known/oauth-authorization-server",
						Status:   200,
						Response: `{"authorization_endpoint": "https://myworkspace.cloud.databricks.com/oidc/v1/authorize", "token_endpoint": "https://myworkspace.cloud.databricks.com/oidc/v1/token"}`,
					},
				},
			}
			got, err := c.getOidcEndpoints(context.Background())
			assert.NoError(t, err)
			assert.Equal(t, &u2m.OAuthAuthorizationServer{
				AuthorizationEndpoint: "https://myworkspace.cloud.databricks.com/oidc/v1/authorize",
				TokenEndpoint:         "https://myworkspace.cloud.databricks.com/oidc/v1/token",
			}, got)
		})
	}
}

func TestConfig_getOidcEndpoints_usesDiscoveryURLFromMetadata(t *testing.T) {
	// Unified hosts resolve DiscoveryURL from metadata during EnsureResolved.
	// The getOidcEndpoints method then uses the DiscoveryURL.
	discoveryURL := "https://unified.cloud.databricks.com/oidc/accounts/abc/.well-known/oauth-authorization-server"
	c := &Config{
		Host:         "https://unified.cloud.databricks.com",
		AccountID:    "abc",
		Token:        "t",
		DiscoveryURL: discoveryURL,
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:   "GET",
				Resource: "/oidc/accounts/abc/.well-known/oauth-authorization-server",
				Status:   200,
				Response: `{"authorization_endpoint": "https://unified.cloud.databricks.com/oidc/accounts/abc/v1/authorize", "token_endpoint": "https://unified.cloud.databricks.com/oidc/accounts/abc/v1/token"}`,
			},
		},
	}
	got, err := c.getOidcEndpoints(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, &u2m.OAuthAuthorizationServer{
		AuthorizationEndpoint: "https://unified.cloud.databricks.com/oidc/accounts/abc/v1/authorize",
		TokenEndpoint:         "https://unified.cloud.databricks.com/oidc/accounts/abc/v1/token",
	}, got)
}

func TestConfig_getOAuthArgument_account(t *testing.T) {
	tests := []struct {
		name      string
		host      string
		accountID string
	}{
		{
			name:      "without trailing slash",
			host:      "https://accounts.cloud.databricks.com",
			accountID: "abc",
		},
		{
			name:      "with trailing slash",
			host:      "https://accounts.cloud.databricks.com/",
			accountID: "abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Host:          tt.host,
				AccountID:     tt.accountID,
				HTTPTransport: metadataNotFoundTransport,
			}
			rawGot, err := c.getOAuthArgument()
			assert.NoError(t, err)
			got, ok := rawGot.(u2m.BasicAccountOAuthArgument)
			assert.True(t, ok)
			assert.Equal(t, "https://accounts.cloud.databricks.com", got.GetAccountHost())
			assert.Equal(t, "abc", got.GetAccountId())
		})
	}
}

func TestConfig_getOAuthArgument_workspace(t *testing.T) {
	tests := []struct {
		name string
		host string
	}{
		{
			name: "without trailing slash",
			host: "https://myworkspace.cloud.databricks.com",
		},
		{
			name: "with trailing slash",
			host: "https://myworkspace.cloud.databricks.com/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Host:          tt.host,
				HTTPTransport: metadataNotFoundTransport,
			}
			rawGot, err := c.getOAuthArgument()
			assert.NoError(t, err)
			got, ok := rawGot.(u2m.BasicWorkspaceOAuthArgument)
			assert.True(t, ok)
			assert.Equal(t, "https://myworkspace.cloud.databricks.com", got.GetWorkspaceHost())
		})
	}
}

func TestConfig_getOAuthArgument_NonAccountsHostUsesWorkspaceArgument(t *testing.T) {
	// A non-accounts host uses a BasicWorkspaceOAuthArgument for OAuth
	// even when AccountID is set.
	c := &Config{
		Host:          "https://unified.cloud.databricks.com",
		AccountID:     "account-123",
		HTTPTransport: metadataNotFoundTransport,
	}
	rawGot, err := c.getOAuthArgument()
	assert.NoError(t, err)
	got, ok := rawGot.(u2m.BasicWorkspaceOAuthArgument)
	assert.True(t, ok, "Expected BasicWorkspaceOAuthArgument")
	assert.Equal(t, "https://unified.cloud.databricks.com", got.GetWorkspaceHost())
}

func TestConfig_getOAuthArgument_profileCacheKeys(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	tests := []struct {
		name        string
		config      *Config
		wantKey     string
		wantHostKey string
	}{
		{
			name: "workspace without profile",
			config: &Config{
				Host:    "https://myworkspace.cloud.databricks.com",
				Loaders: []Loader{noopLoader},
			},
			wantKey:     "https://myworkspace.cloud.databricks.com",
			wantHostKey: "https://myworkspace.cloud.databricks.com",
		},
		{
			name: "workspace with profile",
			config: &Config{
				Host:    "https://myworkspace.cloud.databricks.com",
				Profile: "ws-profile",
				Loaders: []Loader{noopLoader},
			},
			wantKey:     "ws-profile",
			wantHostKey: "https://myworkspace.cloud.databricks.com",
		},
		{
			name: "account without profile",
			config: &Config{
				Host:      "https://accounts.cloud.databricks.com",
				AccountID: "abc",
				Loaders:   []Loader{noopLoader},
			},
			wantKey:     "https://accounts.cloud.databricks.com/oidc/accounts/abc",
			wantHostKey: "https://accounts.cloud.databricks.com/oidc/accounts/abc",
		},
		{
			name: "account with profile",
			config: &Config{
				Host:      "https://accounts.cloud.databricks.com",
				AccountID: "abc",
				Profile:   "my-profile",
				Loaders:   []Loader{noopLoader},
			},
			wantKey:     "my-profile",
			wantHostKey: "https://accounts.cloud.databricks.com/oidc/accounts/abc",
		},
		{
			name: "non-accounts host with account_id without profile",
			config: &Config{
				Host:      "https://unified.cloud.databricks.com",
				AccountID: "account-123",
				Loaders:   []Loader{noopLoader},
			},
			wantKey:     "https://unified.cloud.databricks.com",
			wantHostKey: "https://unified.cloud.databricks.com",
		},
		{
			name: "non-accounts host with account_id and profile",
			config: &Config{
				Host:      "https://unified.cloud.databricks.com",
				AccountID: "account-123",
				Profile:   "unified-profile",
				Loaders:   []Loader{noopLoader},
			},
			wantKey:     "unified-profile",
			wantHostKey: "https://unified.cloud.databricks.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.config.HTTPTransport = metadataNotFoundTransport
			rawGot, err := tt.config.getOAuthArgument()
			assert.NoError(t, err)
			assert.Equal(t, tt.wantKey, rawGot.GetCacheKey())
			if hcp, ok := rawGot.(u2m.HostCacheKeyProvider); ok {
				assert.Equal(t, tt.wantHostKey, hcp.GetHostCacheKey())
			}
		})
	}
}

func TestConfig_EnsureResolved_scopeNormalization(t *testing.T) {
	testCases := []struct {
		desc   string
		scopes []string
		want   []string
	}{
		{
			desc:   "nil scopes",
			scopes: nil,
			want:   nil,
		},
		{
			desc:   "empty scopes",
			scopes: []string{},
			want:   []string{},
		},
		{
			desc:   "single scope",
			scopes: []string{"clusters"},
			want:   []string{"clusters"},
		},
		{
			desc:   "already sorted no duplicates",
			scopes: []string{"a", "b", "c"},
			want:   []string{"a", "b", "c"},
		},
		{
			desc:   "unsorted scopes are sorted",
			scopes: []string{"jobs", "clusters", "pipelines"},
			want:   []string{"clusters", "jobs", "pipelines"},
		},
		{
			desc:   "duplicate scopes are removed",
			scopes: []string{"clusters", "jobs", "clusters", "pipelines:read", "jobs"},
			want:   []string{"clusters", "jobs", "pipelines:read"},
		},
		{
			desc:   "all duplicates reduced to one",
			scopes: []string{"all-apis", "all-apis", "all-apis"},
			want:   []string{"all-apis"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			cfg := &Config{
				Host: "https://example.cloud.databricks.com",
				Loaders: []Loader{mockLoader(func(cfg *Config) error {
					cfg.Scopes = tc.scopes
					return nil
				})},
				HTTPTransport: metadataNotFoundTransport,
			}

			err := cfg.EnsureResolved()
			if err != nil {
				t.Fatalf("EnsureResolved() error: %v", err)
			}

			if diff := cmp.Diff(tc.want, cfg.Scopes); diff != "" {
				t.Errorf("EnsureResolved() scopes mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestConfig_DiscoveryURL_FromEnv(t *testing.T) {
	t.Setenv("DATABRICKS_DISCOVERY_URL", "https://custom.idp.example.com/oidc")
	cfg := &Config{Host: testHMHost, Token: "t", HTTPTransport: metadataNotFoundTransport}
	if err := cfg.EnsureResolved(); err != nil {
		t.Fatal(err)
	}
	if cfg.DiscoveryURL != "https://custom.idp.example.com/oidc" {
		t.Errorf("unexpected DiscoveryURL: %q", cfg.DiscoveryURL)
	}
}

func TestConfig_getOidcEndpoints_UsesDiscoveryURL(t *testing.T) {
	discoveryURL := testHMHost + "/oidc/.well-known/oauth-authorization-server"
	cfg := &Config{
		Host:         testHMHost,
		Token:        "t",
		DiscoveryURL: discoveryURL,
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:   "GET",
				Resource: "/oidc/.well-known/oauth-authorization-server",
				Status:   200,
				Response: `{"authorization_endpoint": "` + testHMHost + `/oidc/v1/authorize", "token_endpoint": "` + testHMHost + `/oidc/v1/token"}`,
			},
		},
	}
	endpoints, err := cfg.getOidcEndpoints(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	want := &u2m.OAuthAuthorizationServer{
		AuthorizationEndpoint: testHMHost + "/oidc/v1/authorize",
		TokenEndpoint:         testHMHost + "/oidc/v1/token",
	}
	if diff := cmp.Diff(want, endpoints); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestConfig_ResolveHostMetadata_AccountSubstitutesAccountID(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:      testHMAccHost,
		AccountID: testHMAccountID,
		Loaders:   []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       200,
				Response:     `{"oidc_endpoint": "` + testHMAccHost + `/oidc/accounts/{account_id}"}`,
			},
		},
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	want := testHMAccHost + "/oidc/accounts/" + testHMAccountID + "/.well-known/oauth-authorization-server"
	assert.Equal(t, want, cfg.DiscoveryURL)
}

func TestConfig_ResolveHostMetadata_DoesNotOverwriteExistingFields(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:        testHMHost,
		AccountID:   testHMAccountID,
		WorkspaceID: testHMWorkspaceID,
		Loaders:     []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       200,
				Response:     `{"oidc_endpoint": "` + testHMHost + `/oidc", "account_id": "other-account", "workspace_id": "other-ws"}`,
			},
		},
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	assert.Equal(t, testHMAccountID, cfg.AccountID)
	assert.Equal(t, testHMWorkspaceID, cfg.WorkspaceID)
}

func TestConfig_ResolveHostMetadata_PopulatesCloudFromAPI(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    testHMHost,
		Loaders: []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       200,
				Response:     `{"oidc_endpoint": "` + testHMHost + `/oidc", "account_id": "` + testHMAccountID + `", "cloud": "Azure"}`,
			},
		},
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	assert.Equal(t, environment.Cloud("Azure"), cfg.Cloud)
}

func TestConfig_ResolveHostMetadata_CloudFallbackToDNS(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    "https://my-workspace.azuredatabricks.net",
		Loaders: []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       200,
				Response:     `{"oidc_endpoint": "https://my-workspace.azuredatabricks.net/oidc", "account_id": "` + testHMAccountID + `"}`,
			},
		},
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	assert.Equal(t, environment.Cloud("Azure"), cfg.Cloud)
}

func TestConfig_ResolveHostMetadata_DoesNotOverwriteExistingCloud(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    testHMHost,
		Cloud:   "GCP",
		Loaders: []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       200,
				Response:     `{"oidc_endpoint": "` + testHMHost + `/oidc", "account_id": "` + testHMAccountID + `", "cloud": "AWS"}`,
			},
		},
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	assert.Equal(t, environment.Cloud("GCP"), cfg.Cloud)
}

func TestEnsureResolved_ResolvesHostMetadata_WhenHostSet(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    testHMHost,
		Loaders: []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       200,
				Response:     `{"oidc_endpoint": "` + testHMHost + `/oidc", "account_id": "` + testHMAccountID + `", "workspace_id": "` + testHMWorkspaceID + `", "cloud": "AWS"}`,
			},
		},
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	// Metadata is now always resolved when host is set
	assert.Equal(t, testHMAccountID, cfg.AccountID)
	assert.Equal(t, testHMWorkspaceID, cfg.WorkspaceID)
}

func TestEnsureResolved_HostMetadataFailure_NonFatal(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    testHMHost,
		Loaders: []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       500,
				Response:     `{"error": "internal error"}`,
			},
		},
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	assert.True(t, cfg.resolved)
}

func TestEnsureResolved_HostMetadata_NoOidcEndpoint_NonFatal(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    testHMHost,
		Loaders: []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       200,
				Response:     `{"account_id": "` + testHMAccountID + `"}`,
			},
		},
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	assert.Equal(t, testHMAccountID, cfg.AccountID)
	assert.Empty(t, cfg.DiscoveryURL)
}

func TestEnsureResolved_HostMetadata_MissingAccountIdWithPlaceholder_Warns(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    testHMHost,
		Loaders: []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       200,
				Response:     `{"oidc_endpoint": "` + testHMHost + `/oidc/accounts/{account_id}"}`,
			},
		},
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	// DiscoveryURL should not be set because account_id is empty and placeholder can't be substituted
	assert.Empty(t, cfg.DiscoveryURL)
}

func TestApplyHostMetadata_SetsTokenAudienceForAccountHost(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    testHMHost,
		Loaders: []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       200,
				Response:     `{"oidc_endpoint": "` + testHMHost + `/oidc", "account_id": "` + testHMAccountID + `", "cloud": "AWS"}`,
			},
		},
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	// No workspace_id and has account_id → token audience should be set
	assert.Equal(t, testHMAccountID, cfg.TokenAudience)
}

func TestApplyHostMetadata_NoTokenAudienceForWorkspaceHost(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    testHMHost,
		Loaders: []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       200,
				Response:     `{"oidc_endpoint": "` + testHMHost + `/oidc", "account_id": "` + testHMAccountID + `", "workspace_id": "` + testHMWorkspaceID + `", "cloud": "AWS"}`,
			},
		},
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	// Has workspace_id → token audience should NOT be set
	assert.Empty(t, cfg.TokenAudience)
}

func TestApplyHostMetadata_DoesNotOverrideExistingTokenAudience(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:          testHMHost,
		TokenAudience: "custom-audience",
		Loaders:       []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       200,
				Response:     `{"oidc_endpoint": "` + testHMHost + `/oidc", "account_id": "` + testHMAccountID + `", "cloud": "AWS"}`,
			},
		},
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	assert.Equal(t, "custom-audience", cfg.TokenAudience)
}

func TestApplyHostMetadata_SetsTokenAudienceFromTokenFederationDefaultOIDCAudiences(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    testHMHost,
		Loaders: []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       200,
				Response:     `{"oidc_endpoint": "` + testHMHost + `/oidc", "account_id": "` + testHMAccountID + `", "workspace_id": "` + testHMWorkspaceID + `", "cloud": "AWS", "token_federation_default_oidc_audiences": ["` + testHMHost + `/oidc/v1/token"]}`,
			},
		},
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	assert.Equal(t, testHMHost+"/oidc/v1/token", cfg.TokenAudience)
}

func TestApplyHostMetadata_TokenFederationDefaultOIDCAudiencesTakesPriorityOverAccountIDFallback(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    testHMHost,
		Loaders: []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       200,
				Response:     `{"oidc_endpoint": "` + testHMHost + `/oidc", "account_id": "` + testHMAccountID + `", "cloud": "AWS", "token_federation_default_oidc_audiences": ["custom-audience-from-server"]}`,
			},
		},
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	// token_federation_default_oidc_audiences should take priority over the account_id fallback
	assert.Equal(t, "custom-audience-from-server", cfg.TokenAudience)
}

func TestApplyHostMetadata_TokenFederationDefaultOIDCAudiencesDoesNotOverrideExisting(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:          testHMHost,
		TokenAudience: "user-set-audience",
		Loaders:       []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       200,
				Response:     `{"oidc_endpoint": "` + testHMHost + `/oidc", "account_id": "` + testHMAccountID + `", "cloud": "AWS", "token_federation_default_oidc_audiences": ["` + testHMHost + `/oidc/v1/token"]}`,
			},
		},
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	assert.Equal(t, "user-set-audience", cfg.TokenAudience)
}

func TestApplyHostMetadata_FallsBackToAccountIDWhenNoTokenFederationDefaultOIDCAudiences(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    testHMHost,
		Loaders: []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       200,
				Response:     `{"oidc_endpoint": "` + testHMHost + `/oidc", "account_id": "` + testHMAccountID + `", "cloud": "AWS"}`,
			},
		},
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	// No token_federation_default_oidc_audiences and no workspace_id → falls back to account_id
	assert.Equal(t, testHMAccountID, cfg.TokenAudience)
}

func TestEnsureResolved_UsesCustomHostMetadataResolver(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    testHMHost,
		Loaders: []Loader{noopLoader},
		HostMetadataResolver: func(ctx context.Context, host string) (*HostMetadata, error) {
			return &HostMetadata{
				OIDCEndpoint: testHMHost + "/oidc",
				AccountID:    testHMAccountID,
				WorkspaceID:  testHMWorkspaceID,
				Cloud:        "AWS",
			}, nil
		},
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	assert.Equal(t, testHMAccountID, cfg.AccountID)
	assert.Equal(t, testHMWorkspaceID, cfg.WorkspaceID)
}

func TestEnsureResolved_CustomResolver_FullReplacement(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    testHMHost,
		Loaders: []Loader{noopLoader},
		HostMetadataResolver: func(ctx context.Context, host string) (*HostMetadata, error) {
			assert.Equal(t, testHMHost, host)
			return &HostMetadata{
				AccountID:   testHMAccountID,
				WorkspaceID: testHMWorkspaceID,
			}, nil
		},
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	assert.Equal(t, testHMAccountID, cfg.AccountID)
	assert.Equal(t, testHMWorkspaceID, cfg.WorkspaceID)
}

func TestEnsureResolved_CustomResolver_NilMetadata_NoBackfill(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    testHMHost,
		Loaders: []Loader{noopLoader},
		HostMetadataResolver: func(ctx context.Context, host string) (*HostMetadata, error) {
			return nil, nil
		},
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	assert.Empty(t, cfg.AccountID)
	assert.Empty(t, cfg.WorkspaceID)
}

func TestEnsureResolved_CustomResolver_Error_NonFatal(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    testHMHost,
		Loaders: []Loader{noopLoader},
		HostMetadataResolver: func(ctx context.Context, host string) (*HostMetadata, error) {
			return nil, fmt.Errorf("resolver error")
		},
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	assert.Empty(t, cfg.AccountID)
	assert.Empty(t, cfg.WorkspaceID)
}

func TestEnsureResolved_DefaultHostMetadataResolver_Fallback(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    testHMHost,
		Loaders: []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       200,
				Response:     `{"oidc_endpoint": "` + testHMHost + `/oidc", "account_id": "` + testHMAccountID + `", "workspace_id": "` + testHMWorkspaceID + `", "cloud": "AWS"}`,
			},
		},
	}
	// Simulate cache-with-fallback: cache miss delegates to the SDK's default fetch.
	cfg.HostMetadataResolver = func(ctx context.Context, host string) (*HostMetadata, error) {
		// Cache miss: fall back to the SDK's built-in HTTP fetch.
		return cfg.DefaultHostMetadataResolver()(ctx, host)
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	assert.Equal(t, testHMAccountID, cfg.AccountID)
	assert.Equal(t, testHMWorkspaceID, cfg.WorkspaceID)
}

func TestConfig_ResolveHostMetadata_Clouds(t *testing.T) {
	tests := []struct {
		name      string
		cloudJSON string
		wantCloud string
	}{
		{
			name:      "AWS",
			cloudJSON: "AWS",
			wantCloud: "AWS",
		},
		{
			name:      "Azure",
			cloudJSON: "Azure",
			wantCloud: "Azure",
		},
		{
			name:      "GCP",
			cloudJSON: "GCP",
			wantCloud: "GCP",
		},
		{
			name:      "aws lowercase",
			cloudJSON: "aws",
			wantCloud: "AWS",
		},
		{
			name:      "AWS uppercase",
			cloudJSON: "AWS",
			wantCloud: "AWS",
		},
		{
			name:      "azure lowercase",
			cloudJSON: "azure",
			wantCloud: "Azure",
		},
		{
			name:      "AZURE uppercase",
			cloudJSON: "AZURE",
			wantCloud: "Azure",
		},
		{
			name:      "Azure title case",
			cloudJSON: "Azure",
			wantCloud: "Azure",
		},
		{
			name:      "gcp lowercase",
			cloudJSON: "gcp",
			wantCloud: "GCP",
		},
		{
			name:      "GCP uppercase",
			cloudJSON: "GCP",
			wantCloud: "GCP",
		},
		{
			name:      "Another cloud is supported",
			cloudJSON: "Another",
			wantCloud: "Another",
		},
		{
			name:      "Unknown cloud falls back to DNS",
			cloudJSON: "",
			wantCloud: "AWS", // Falls back to DNS-based detection for testHMHost
		},
	}
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cfg := &Config{
				Host:    testHMHost,
				Loaders: []Loader{noopLoader},
				HTTPTransport: fixtures.SliceTransport{
					{
						Method:       "GET",
						Resource:     "/.well-known/databricks-config",
						ReuseRequest: true,
						Status:       200,
						Response:     `{"oidc_endpoint": "` + testHMHost + `/oidc", "account_id": "` + testHMAccountID + `", "cloud": "` + tc.cloudJSON + `"}`,
					},
				},
			}
			err := cfg.EnsureResolved()
			require.NoError(t, err)
			assert.Equal(t, tc.wantCloud, string(cfg.Cloud))
		})
	}
}

func TestConfig_ResolveHostMetadata_PopulatesHostTypeFromAPI(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    testHMHost,
		Loaders: []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       200,
				Response:     `{"oidc_endpoint": "` + testHMHost + `/oidc", "account_id": "` + testHMAccountID + `", "host_type": "unified"}`,
			},
		},
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	assert.Equal(t, UnifiedHost, cfg.resolvedHostType)
}

func TestConfig_ResolveHostMetadata_HostTypeEmptyWhenNotInResponse(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    "https://accounts.cloud.databricks.com",
		Loaders: []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       200,
				Response:     `{"oidc_endpoint": "https://accounts.cloud.databricks.com/oidc", "account_id": "` + testHMAccountID + `"}`,
			},
		},
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	assert.Equal(t, HostTypeUnknown, cfg.resolvedHostType)
}

func TestConfig_ResolveHostMetadata_DoesNotOverwriteExistingHostType(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    testHMHost,
		Loaders: []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       200,
				Response:     `{"oidc_endpoint": "` + testHMHost + `/oidc", "account_id": "` + testHMAccountID + `", "host_type": "account"}`,
			},
		},
	}
	cfg.resolvedHostType = UnifiedHost
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	assert.Equal(t, UnifiedHost, cfg.resolvedHostType)
}

func TestConfig_ResolveHostMetadata_HostTypes(t *testing.T) {
	tests := []struct {
		name         string
		hostTypeJSON string
		wantHostType string
	}{
		{
			name:         "workspace",
			hostTypeJSON: "workspace",
			wantHostType: "WORKSPACE_HOST",
		},
		{
			name:         "account",
			hostTypeJSON: "account",
			wantHostType: "ACCOUNT_HOST",
		},
		{
			name:         "unified",
			hostTypeJSON: "unified",
			wantHostType: "UNIFIED_HOST",
		},
		{
			name:         "Workspace uppercase",
			hostTypeJSON: "WORKSPACE",
			wantHostType: "WORKSPACE_HOST",
		},
		{
			name:         "Account uppercase",
			hostTypeJSON: "ACCOUNT",
			wantHostType: "ACCOUNT_HOST",
		},
		{
			name:         "Unified uppercase",
			hostTypeJSON: "UNIFIED",
			wantHostType: "UNIFIED_HOST",
		},
		{
			name:         "Unknown host type string returns HostTypeUnknown",
			hostTypeJSON: "CUSTOM_HOST",
			wantHostType: "",
		},
		{
			name:         "Empty host type returns HostTypeUnknown",
			hostTypeJSON: "",
			wantHostType: "",
		},
	}
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cfg := &Config{
				Host:    testHMHost,
				Loaders: []Loader{noopLoader},
				HTTPTransport: fixtures.SliceTransport{
					{
						Method:       "GET",
						Resource:     "/.well-known/databricks-config",
						ReuseRequest: true,
						Status:       200,
						Response:     `{"oidc_endpoint": "` + testHMHost + `/oidc", "account_id": "` + testHMAccountID + `", "host_type": "` + tc.hostTypeJSON + `"}`,
					},
				},
			}
			err := cfg.EnsureResolved()
			require.NoError(t, err)
			assert.Equal(t, tc.wantHostType, string(cfg.resolvedHostType))
		})
	}
}

// withDefaultHostMetadataResolverFactory installs factory for the duration of
// the current test, restoring whatever was previously set on cleanup.
// Capture/set/restore are not atomic — do not use with t.Parallel across
// multiple tests that touch the package-level default.
func withDefaultHostMetadataResolverFactory(t *testing.T, factory func(*Config) HostMetadataResolver) {
	t.Helper()
	prev := DefaultHostMetadataResolverFactory
	DefaultHostMetadataResolverFactory = factory
	t.Cleanup(func() { DefaultHostMetadataResolverFactory = prev })
}

func TestDefaultHostMetadataResolverFactory_UsedWhenConfigHasNoResolver(t *testing.T) {
	var factoryCalls atomic.Int32
	withDefaultHostMetadataResolverFactory(t, func(c *Config) HostMetadataResolver {
		factoryCalls.Add(1)
		return func(ctx context.Context, host string) (*HostMetadata, error) {
			return &HostMetadata{AccountID: testHMAccountID, WorkspaceID: testHMWorkspaceID}, nil
		}
	})

	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{Host: testHMHost, Loaders: []Loader{noopLoader}}
	require.NoError(t, cfg.EnsureResolved())

	assert.Equal(t, int32(1), factoryCalls.Load(), "factory must be invoked exactly once per resolve")
	assert.Equal(t, testHMAccountID, cfg.AccountID)
	assert.Equal(t, testHMWorkspaceID, cfg.WorkspaceID)
}

func TestDefaultHostMetadataResolverFactory_PerConfigResolverTakesPrecedence(t *testing.T) {
	var factoryCalls atomic.Int32
	withDefaultHostMetadataResolverFactory(t, func(c *Config) HostMetadataResolver {
		factoryCalls.Add(1)
		return func(ctx context.Context, host string) (*HostMetadata, error) {
			return &HostMetadata{AccountID: "factory-account"}, nil
		}
	})

	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    testHMHost,
		Loaders: []Loader{noopLoader},
		HostMetadataResolver: func(ctx context.Context, host string) (*HostMetadata, error) {
			return &HostMetadata{AccountID: testHMAccountID}, nil
		},
	}
	require.NoError(t, cfg.EnsureResolved())

	assert.Equal(t, int32(0), factoryCalls.Load(), "factory must not be consulted when Config has its own resolver")
	assert.Equal(t, testHMAccountID, cfg.AccountID)
}

func TestDefaultHostMetadataResolverFactory_NilResolverFromFactoryFallsThroughToHTTP(t *testing.T) {
	withDefaultHostMetadataResolverFactory(t, func(c *Config) HostMetadataResolver {
		return nil
	})

	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    testHMHost,
		Loaders: []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       200,
				Response:     `{"oidc_endpoint": "` + testHMHost + `/oidc", "account_id": "` + testHMAccountID + `"}`,
			},
		},
	}
	require.NoError(t, cfg.EnsureResolved())

	assert.Equal(t, testHMAccountID, cfg.AccountID)
}

func TestConfig_fixHostIfNeeded_extractsWorkspaceIDFromQuery(t *testing.T) {
	tests := []struct {
		name            string
		host            string
		workspaceID     string
		accountID       string
		wantHost        string
		wantWorkspaceID string
		wantAccountID   string
	}{
		{
			name:            "?o= promoted to WorkspaceID",
			host:            "https://acme.databricks.net/?o=12345",
			wantHost:        "https://acme.databricks.net",
			wantWorkspaceID: "12345",
		},
		{
			name:            "?workspace_id= promoted to WorkspaceID",
			host:            "https://acme.databricks.net/?workspace_id=12345",
			wantHost:        "https://acme.databricks.net",
			wantWorkspaceID: "12345",
		},
		{
			name:          "?a= promoted to AccountID",
			host:          "https://acme.databricks.net/?a=abc",
			wantHost:      "https://acme.databricks.net",
			wantAccountID: "abc",
		},
		{
			name:            "?o= and ?a= both promoted",
			host:            "https://acme.databricks.net/?o=12345&a=abc",
			wantHost:        "https://acme.databricks.net",
			wantWorkspaceID: "12345",
			wantAccountID:   "abc",
		},
		{
			name:            "existing WorkspaceID is preserved",
			host:            "https://acme.databricks.net/?o=12345",
			workspaceID:     "99999",
			wantHost:        "https://acme.databricks.net",
			wantWorkspaceID: "99999",
		},
		{
			name:          "existing AccountID is preserved",
			host:          "https://acme.databricks.net/?a=other",
			accountID:     "kept",
			wantHost:      "https://acme.databricks.net",
			wantAccountID: "kept",
		},
		{
			name:     "non-numeric ?o= is dropped",
			host:     "https://acme.databricks.net/?o=notanumber",
			wantHost: "https://acme.databricks.net",
		},
		{
			name:            "?w= with numeric value promoted to WorkspaceID",
			host:            "https://acme.databricks.net/?w=12345",
			wantHost:        "https://acme.databricks.net",
			wantWorkspaceID: "12345",
		},
		{
			name:            "?w= with non-numeric value promoted to WorkspaceID",
			host:            "https://acme.databricks.net/?w=7a99b43c-b46c-432b-b0a7-814217701909",
			wantHost:        "https://acme.databricks.net",
			wantWorkspaceID: "7a99b43c-b46c-432b-b0a7-814217701909",
		},
		{
			name:            "?w= takes precedence over ?o=",
			host:            "https://acme.databricks.net/?w=12345&o=99999",
			wantHost:        "https://acme.databricks.net",
			wantWorkspaceID: "12345",
		},
		{
			name:            "existing WorkspaceID is preserved over ?w=",
			host:            "https://acme.databricks.net/?w=12345",
			workspaceID:     "99999",
			wantHost:        "https://acme.databricks.net",
			wantWorkspaceID: "99999",
		},
		{
			name:     "host without query is unchanged",
			host:     "https://acme.databricks.net",
			wantHost: "https://acme.databricks.net",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &Config{
				Host:        tt.host,
				WorkspaceID: tt.workspaceID,
				AccountID:   tt.accountID,
			}
			require.NoError(t, cfg.fixHostIfNeeded())
			assert.Equal(t, tt.wantHost, cfg.Host)
			assert.Equal(t, tt.wantWorkspaceID, cfg.WorkspaceID)
			assert.Equal(t, tt.wantAccountID, cfg.AccountID)
		})
	}
}
