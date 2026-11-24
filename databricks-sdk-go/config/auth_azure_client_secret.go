package config

import (
	"context"
	"fmt"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"github.com/databricks/databricks-sdk-go/logger"
)

type AzureClientSecretCredentials struct {
}

func (c AzureClientSecretCredentials) Name() string {
	return "azure-client-secret"
}

// tokenSourceFor returns an auth.TokenSource that creates a fresh
// clientcredentials.Config on each call and uses the provided context for the
// token exchange. This does not include the inner caching layer from
// clientcredentials.Config.TokenSource so that caching is purely controlled
// upstream.
func (c AzureClientSecretCredentials) tokenSourceFor(
	_ context.Context, cfg *Config, aadEndpoint, resource string) auth.TokenSource {
	return auth.TokenSourceFn(func(ctx context.Context) (*oauth2.Token, error) {
		ctx = cfg.refreshClient.InContextForOAuth2(ctx)
		return (&clientcredentials.Config{
			ClientID:     cfg.AzureClientID,
			ClientSecret: cfg.AzureClientSecret,
			TokenURL:     fmt.Sprintf("%s%s/oauth2/token", aadEndpoint, cfg.AzureTenantID),
			EndpointParams: url.Values{
				"resource": []string{resource},
			},
		}).Token(ctx)
	})
}

// TODO: We need to expose which authentication mechanism is used to Terraform,
// as we cannot create AKV backed secret scopes when authenticated as SP.
// If we are authenticated as SP and wish to create one we want to fail early.
// Also see https://github.com/databricks/terraform-provider-databricks/issues/1490.
func (c AzureClientSecretCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if err := unsupportedGroupRoleAssumption(cfg, c.Name()); err != nil {
		return nil, err
	}
	if cfg.AzureClientID == "" || cfg.AzureClientSecret == "" || cfg.AzureTenantID == "" {
		return nil, nil
	}
	if !cfg.IsAzure() {
		return nil, nil
	}
	err := cfg.loadAzureTenantId(ctx)
	if err != nil {
		return nil, fmt.Errorf("load tenant id: %w", err)
	}
	err = cfg.azureEnsureWorkspaceUrl(ctx, c)
	if err != nil {
		return nil, fmt.Errorf("resolve host: %w", err)
	}
	logger.Infof(ctx, "Generating AAD token for Service Principal (%s)", cfg.AzureClientID)
	env := cfg.Environment()
	aadEndpoint := env.AzureActiveDirectoryEndpoint()
	managementEndpoint := env.AzureServiceManagementEndpoint()
	opts := cacheOptions(cfg)
	inner := azureReuseTokenSource(nil, auth.NewRetryingTokenSource(c.tokenSourceFor(ctx, cfg, aadEndpoint, env.AzureApplicationID)), opts...)
	management := azureReuseTokenSource(nil, auth.NewRetryingTokenSource(c.tokenSourceFor(ctx, cfg, aadEndpoint, managementEndpoint)), opts...)
	visitor := azureVisitor(cfg, serviceToServiceVisitor(inner, management, xDatabricksAzureSpManagementToken, false, opts...))
	return newVisitorOAuthCredentials(visitor, inner), nil
}
