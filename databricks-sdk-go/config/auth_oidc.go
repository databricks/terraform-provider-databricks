package config

import (
	"context"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth/oidc"
)

// GitHubOIDCCredentials authenticates with Databricks using an OIDC token
// from GitHub Actions.
type GitHubOIDCCredentials struct{}

func (c GitHubOIDCCredentials) Name() string {
	return "github-oidc"
}

func (c GitHubOIDCCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	return githubOIDC(cfg).Configure(ctx, cfg)
}

// AzureDevOpsOIDCCredentials authenticates with Databricks using an OIDC token
// from Azure DevOps pipelines.
type AzureDevOpsOIDCCredentials struct{}

func (c AzureDevOpsOIDCCredentials) Name() string {
	return "azure-devops-oidc"
}

func (c AzureDevOpsOIDCCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	return azureDevOpsOIDC(cfg).Configure(ctx, cfg)
}

// EnvOIDCCredentials authenticates with Databricks using an OIDC token read
// from an environment variable.
type EnvOIDCCredentials struct{}

func (c EnvOIDCCredentials) Name() string {
	return "env-oidc"
}

func (c EnvOIDCCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	return envOIDC(cfg).Configure(ctx, cfg)
}

// FileOIDCCredentials authenticates with Databricks using an OIDC token read
// from a file.
type FileOIDCCredentials struct{}

func (c FileOIDCCredentials) Name() string {
	return "file-oidc"
}

func (c FileOIDCCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	return fileOIDC(cfg).Configure(ctx, cfg)
}

func githubOIDC(cfg *Config) CredentialsStrategy {
	return oidcStrategy(cfg, "github-oidc", oidc.NewGithubIDTokenSource(
		cfg.refreshClient,
		cfg.ActionsIDTokenRequestURL,
		cfg.ActionsIDTokenRequestToken,
	))
}

func azureDevOpsOIDC(cfg *Config) CredentialsStrategy {
	// Return a systemically failed strategy if the Azure DevOps ID token source
	// cannot be created.
	idts, err := oidc.NewAzureDevOpsIDTokenSource(cfg.refreshClient)
	if err != nil {
		return &failedStrategy{name: "azure-devops-oidc", err: err}
	}
	return oidcStrategy(cfg, "azure-devops-oidc", idts)
}

func envOIDC(cfg *Config) CredentialsStrategy {
	v := cfg.OIDCTokenEnv
	if v == "" {
		v = "DATABRICKS_OIDC_TOKEN"
	}
	return oidcStrategy(cfg, "env-oidc", oidc.NewEnvIDTokenSource(v))
}

func fileOIDC(cfg *Config) CredentialsStrategy {
	return oidcStrategy(cfg, "file-oidc", oidc.NewFileTokenSource(cfg.OIDCTokenFilepath))
}

// oidcStrategy returns a new CredentialsStrategy to authenticate with
// Databricks using the given OIDC IDTokenSource.
func oidcStrategy(cfg *Config, name string, ts oidc.IDTokenSource) CredentialsStrategy {
	oidcConfig := oidc.DatabricksOIDCTokenSourceConfig{
		ClientID:              cfg.ClientID,
		Host:                  cfg.CanonicalHostName(),
		GroupID:               cfg.GroupID,
		TokenEndpointProvider: cfg.getOidcEndpoints,
		Audience:              cfg.TokenAudience,
		IDTokenSource:         ts,
	}
	if cfg.HostType() != WorkspaceHost {
		oidcConfig.AccountID = cfg.AccountID
	}
	oidcConfig.SetScopes(cfg.GetScopes())
	tokenSource := oidc.NewDatabricksOIDCTokenSource(oidcConfig)
	return NewTokenSourceStrategy(name, auth.NewRetryingTokenSource(tokenSource))
}

// failedStrategy is a CredentialsStrategy that always fails.
//
// TODO: This type will likely be removed in the future. It is needed at moment
// to provide a way to fail early when an OIDC credentials strategy cannot be
// configured.
type failedStrategy struct {
	name string
	err  error
}

func (s *failedStrategy) Name() string {
	return s.name
}

func (s *failedStrategy) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	return nil, s.err
}
