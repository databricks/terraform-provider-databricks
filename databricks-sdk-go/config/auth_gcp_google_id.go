package config

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth/authconv"
	"github.com/databricks/databricks-sdk-go/logger"
	"google.golang.org/api/impersonate"
	"google.golang.org/api/option"
)

type GoogleDefaultCredentials struct {
	// options used to enable unit testing mode for OIDC
	opts []option.ClientOption
}

func (c GoogleDefaultCredentials) Name() string {
	return "google-id"
}

func (c GoogleDefaultCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if err := unsupportedGroupRoleAssumption(cfg, c.Name()); err != nil {
		return nil, err
	}
	if cfg.GoogleServiceAccount == "" || !cfg.IsGcp() {
		return nil, nil
	}
	inner, err := c.idTokenSource(ctx, cfg.Host, cfg.GoogleServiceAccount, c.opts...)
	if err != nil {
		return nil, err
	}

	// Disable async token refresh. Google's token sources cache tokens
	// internally via oauth2.ReuseTokenSource, and there is no way to
	// bypass this caching for unexported token source types. Async
	// refresh would be unnecessary work since Google's cache already
	// handles token renewal.
	opts := append(cacheOptions(cfg), auth.WithAsyncRefresh(false))

	// Always attempt to create SA token source for the secondary header.
	// If it fails, fall back to refreshableAuthVisitor with a warning.
	platform, err := impersonate.CredentialsTokenSource(ctx, impersonate.CredentialsConfig{
		TargetPrincipal: cfg.GoogleServiceAccount,
		Scopes: []string{
			"https://www.googleapis.com/auth/cloud-platform",
			"https://www.googleapis.com/auth/compute",
		},
	}, c.opts...)
	if err != nil {
		logger.Warnf(ctx, "Failed to create GCP SA access token source: %v. Proceeding without SA token.", err)
		visitor := refreshableAuthVisitor(inner, opts...)
		return credentials.CredentialsProviderFn(visitor), nil
	}
	logger.Infof(ctx, "Using Google Default Application Credentials")
	visitor := serviceToServiceVisitor(inner, authconv.AuthTokenSource(platform), "X-Databricks-GCP-SA-Access-Token", true, opts...)
	return newVisitorOAuthCredentials(visitor, inner), nil
}

func (c GoogleDefaultCredentials) idTokenSource(ctx context.Context, host, serviceAccount string,
	opts ...option.ClientOption) (auth.TokenSource, error) {
	ts, err := impersonate.IDTokenSource(ctx, impersonate.IDTokenConfig{
		Audience:        host,
		TargetPrincipal: serviceAccount,
		IncludeEmail:    true,
	}, opts...)
	if err != nil {
		err = fmt.Errorf("could not obtain OIDC token. %w Running 'gcloud auth application-default login' may help", err)
		return nil, err
	}
	return authconv.AuthTokenSource(ts), nil
}
