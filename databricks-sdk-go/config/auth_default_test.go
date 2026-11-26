package config

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/config/credentials"
)

func TestDefaultCredentialStrategy(t *testing.T) {
	t.Setenv("PATH", "")
	original := DefaultCredentialStrategyProvider
	t.Cleanup(func() { DefaultCredentialStrategyProvider = original })

	want := &DefaultCredentials{}
	DefaultCredentialStrategyProvider = func() CredentialsStrategy {
		return want
	}

	cfg := &Config{
		Host:          "https://example.databricks.com",
		ConfigFile:    "/dev/null",
		HTTPTransport: metadataNotFoundTransport,
	}
	cfg.Authenticate(&http.Request{Header: http.Header{}})

	if cfg.Credentials != want {
		t.Errorf("Credentials: got %v, want %v", cfg.Credentials, want)
	}
}

// TestBuiltInCredentials_GroupRoleUnsupportedStrategiesDecline verifies that
// built-in normal-access strategies cannot authenticate a group-role client.
func TestBuiltInCredentials_GroupRoleUnsupportedStrategiesDecline(t *testing.T) {
	testCases := []struct {
		strategy CredentialsStrategy
		cfg      *Config
	}{
		{
			strategy: PatCredentials{},
			cfg: &Config{
				Host:  "https://workspace.cloud.databricks.com",
				Token: "pat",
			},
		},
		{
			strategy: BasicCredentials{},
			cfg: &Config{
				Host:     "https://workspace.cloud.databricks.com",
				Username: "user",
				Password: "password",
			},
		},
		{
			strategy: u2mCredentials{},
			cfg: &Config{
				Host: "https://workspace.cloud.databricks.com",
			},
		},
		{
			strategy: MetadataServiceCredentials{},
			cfg: &Config{
				Host:               "https://workspace.cloud.databricks.com",
				MetadataServiceURL: "http://localhost/token",
			},
		},
		{
			strategy: AzureGithubOIDCCredentials{},
			cfg: &Config{
				Host:                       "https://adb-123.azuredatabricks.net",
				AzureClientID:              "client-id",
				AzureTenantID:              "tenant-id",
				ActionsIDTokenRequestURL:   "https://identity.example/token",
				ActionsIDTokenRequestToken: "request-token",
			},
		},
		{
			strategy: AzureMsiCredentials{},
			cfg: &Config{
				Host:        "https://adb-123.azuredatabricks.net",
				AzureUseMSI: true,
			},
		},
		{
			strategy: AzureClientSecretCredentials{},
			cfg: &Config{
				Host:              "https://adb-123.azuredatabricks.net",
				AzureClientID:     "client-id",
				AzureClientSecret: "client-secret",
				AzureTenantID:     "tenant-id",
			},
		},
		{
			strategy: AzureCliCredentials{},
			cfg: &Config{
				Host: "https://adb-123.azuredatabricks.net",
			},
		},
		{
			strategy: GoogleCredentials{},
			cfg: &Config{
				Host:              "https://workspace.gcp.databricks.com",
				GoogleCredentials: `{}`,
			},
		},
		{
			strategy: GoogleDefaultCredentials{},
			cfg: &Config{
				Host:                 "https://workspace.gcp.databricks.com",
				GoogleServiceAccount: "service-account@example.com",
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.strategy.Name(), func(t *testing.T) {
			testCase.cfg.GroupID = "group-123"
			testCase.cfg.resolved = true
			testCase.cfg.resolvedHostType = WorkspaceHost
			provider, err := testCase.strategy.Configure(context.Background(), testCase.cfg)
			if provider != nil {
				t.Errorf("Configure() provider = %v, want nil", provider)
			}
			if err == nil || !strings.Contains(err.Error(), "does not support group role assumption") {
				t.Errorf("Configure() error = %v, want actionable group role error", err)
			}
		})
	}
}

// TestCredentialsChain_GroupRoleFallbackSkipsNormalAccess verifies that the
// credential chain can continue between role-capable strategies without using
// an intervening normal-access strategy.
func TestCredentialsChain_GroupRoleFallbackSkipsNormalAccess(t *testing.T) {
	cfg := &Config{
		Host:             "https://workspace.cloud.databricks.com",
		Username:         "user",
		Password:         "password",
		GroupID:          "group-123",
		resolved:         true,
		resolvedHostType: WorkspaceHost,
	}
	wantProvider := credentials.CredentialsProviderFn(func(req *http.Request) error {
		req.Header.Set("Authorization", "Bearer role-token")
		return nil
	})
	failedRoleStrategy := &testCredentialsStrategy{
		name: "failed-role-capable",
		provider: credentials.CredentialsProviderFn(func(*http.Request) error {
			return errors.New("role token rejected")
		}),
	}
	roleStrategy := &testCredentialsStrategy{name: "role-capable", provider: wantProvider}
	chain := NewCredentialsChain(failedRoleStrategy, BasicCredentials{}, roleStrategy)

	gotProvider, err := chain.Configure(context.Background(), cfg)
	if err != nil {
		t.Fatalf("Configure() error = %v", err)
	}
	if gotProvider == nil {
		t.Fatal("Configure() provider = nil, want role credentials provider")
	}
	if failedRoleStrategy.calls != 1 {
		t.Errorf("failed role strategy calls = %d, want 1", failedRoleStrategy.calls)
	}
	if roleStrategy.calls != 1 {
		t.Errorf("successful role strategy calls = %d, want 1", roleStrategy.calls)
	}

	req := &http.Request{Header: make(http.Header)}
	if err := gotProvider.SetHeaders(req); err != nil {
		t.Fatalf("SetHeaders() error = %v", err)
	}

	if got := req.Header.Get("Authorization"); got != "Bearer role-token" {
		t.Errorf("Authorization header = %q, want %q", got, "Bearer role-token")
	}
}

// TestCredentialsChain_GroupRoleExhaustionKeepsGenericError verifies that
// exhausting the role-aware credential chain preserves the existing error.
func TestCredentialsChain_GroupRoleExhaustionKeepsGenericError(t *testing.T) {
	cfg := &Config{GroupID: "group-123", resolved: true, resolvedHostType: WorkspaceHost}
	chain := NewCredentialsChain(PatCredentials{}, BasicCredentials{}, MetadataServiceCredentials{})

	provider, err := chain.Configure(context.Background(), cfg)
	if provider != nil {
		t.Errorf("Configure() provider = %v, want nil", provider)
	}
	if err != ErrCannotConfigureDefault {
		t.Errorf("Configure() error = %v, want %v", err, ErrCannotConfigureDefault)
	}
}

type testCredentialsStrategy struct {
	name     string
	provider credentials.CredentialsProvider
	calls    int
}

// Name returns the strategy name used by the credential-chain test.
func (s *testCredentialsStrategy) Name() string { return s.name }

// Configure records the attempt and returns the test strategy's provider.
func (s *testCredentialsStrategy) Configure(context.Context, *Config) (credentials.CredentialsProvider, error) {
	s.calls++
	return s.provider, nil
}

func TestDefaultCredentials_Configure(t *testing.T) {
	testCases := []struct {
		desc     string
		authType string
		wantErr  string
	}{
		{
			desc:     "unknown auth type",
			authType: "unknown-auth-type-1337",
			wantErr:  "auth type \"unknown-auth-type-1337\" not found",
		},
		{
			desc:     "not valid auth",
			authType: "",
			wantErr:  "cannot configure default credentials",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := context.Background()
			cfg := &Config{
				AuthType: tc.authType,
				resolved: true, // avoid calling EnsureResolved
			}

			dc := DefaultCredentials{}
			got, gotErr := dc.Configure(ctx, cfg)

			if got != nil {
				t.Errorf("DefaultCredentials.Configure: got %v, want nil", got)
			}
			if gotErr == nil {
				t.Errorf("DefaultCredentials.Configure: got error %v, want non-nil", gotErr)
			}
			if !strings.Contains(gotErr.Error(), tc.wantErr) {
				t.Errorf("DefaultCredentials.Configure: got error %v, want error containing %q", gotErr, tc.wantErr)
			}
		})
	}
}
