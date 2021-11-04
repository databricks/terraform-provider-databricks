package common

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"google.golang.org/api/impersonate"
)

func (c *DatabricksClient) getGoogleOIDCSource(ctx context.Context) (oauth2.TokenSource, error) {
	// source for generateIdToken
	ts, err := impersonate.IDTokenSource(ctx, impersonate.IDTokenConfig{
		Audience:        c.Host,
		TargetPrincipal: c.GoogleServiceAccount,
		IncludeEmail:    true,
	}, c.googleAuthOptions...)
	if err != nil {
		err = fmt.Errorf("could not obtain OIDC token. %w Running 'gcloud auth application-default login' may help", err)
		return nil, err
	}
	// TODO: verify that refreshers work...
	ts = oauth2.ReuseTokenSource(nil, ts)
	return ts, nil
}

func (c *DatabricksClient) configureWithGoogleForAccountsAPI(ctx context.Context) (func(*http.Request) error, error) {
	if c.GoogleServiceAccount == "" || !c.IsGcp() || !c.isAccountsClient() {
		return nil, nil
	}
	oidcSource, err := c.getGoogleOIDCSource(ctx)
	if err != nil {
		return nil, err
	}
	// source for generateAccessToken
	platformSource, err := impersonate.CredentialsTokenSource(ctx, impersonate.CredentialsConfig{
		TargetPrincipal: c.GoogleServiceAccount,
		Scopes: []string{
			"https://www.googleapis.com/auth/cloud-platform",
			"https://www.googleapis.com/auth/compute",
		},
	}, c.googleAuthOptions...)
	if err != nil {
		return nil, err
	}
	return newOidcAuthorizerForAccountsAPI(oidcSource, platformSource), nil
}

func newOidcAuthorizerForAccountsAPI(oidcSource oauth2.TokenSource,
	platformSource oauth2.TokenSource) func(r *http.Request) error {
	return func(r *http.Request) error {
		oidc, err := oidcSource.Token()
		if err != nil {
			return fmt.Errorf("failed to get oidc token: %w", err)
		}
		cloudAccess, err := platformSource.Token()
		if err != nil {
			return fmt.Errorf("failed to get access token: %w", err)
		}
		oidc.SetAuthHeader(r)
		r.Header.Set("X-Databricks-GCP-SA-Access-Token", cloudAccess.AccessToken)
		return nil
	}
}

func (c *DatabricksClient) configureWithGoogleForWorkspace(ctx context.Context) (func(r *http.Request) error, error) {
	if c.GoogleServiceAccount == "" || !c.IsGcp() || c.isAccountsClient() {
		return nil, nil
	}
	oidcSource, err := c.getGoogleOIDCSource(ctx)
	if err != nil {
		return nil, err
	}
	return newOidcAuthorizerForWorkspace(oidcSource), nil
}

func newOidcAuthorizerForWorkspace(oidcSource oauth2.TokenSource) func(r *http.Request) error {
	return func(r *http.Request) error {
		oidc, err := oidcSource.Token()
		if err != nil {
			return err
		}
		oidc.SetAuthHeader(r)
		return nil
	}
}
