package common

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
	"google.golang.org/api/impersonate"
	"google.golang.org/api/option"
)

// Configures an authorizer that uses credentials sourced from JSON or file.
// This auth mode does NOT use impersonation and it does NOT use Application
// Default Credentials (ADC).
func (c *DatabricksClient) configureWithGoogleCrendentials(
	ctx context.Context) (func(r *http.Request) error, error) {
	if c.GoogleCredentials == "" || !c.IsGcp() || c.Host == "" {
		return nil, nil
	}
	json, err := readCredentials(c.GoogleCredentials)
	if err != nil {
		err = fmt.Errorf("could not read GoogleCredentials. "+
			"Make sure the file exists, or the JSON content is valid: %w", err)
		return nil, err
	}
	// Obtain token source for creating OIDC token.
	audience := c.Host
	oidcSource, err := idtoken.NewTokenSource(ctx, audience, option.WithCredentialsJSON([]byte(json)))
	if err != nil {
		return nil, fmt.Errorf("could not obtain OIDC token from JSON: %w", err)
	}
	// Obtain token source for creating Google Cloud Platform token.
	creds, err := google.CredentialsFromJSON(ctx, []byte(json),
		"https://www.googleapis.com/auth/cloud-platform",
		"https://www.googleapis.com/auth/compute")
	if err != nil {
		return nil, fmt.Errorf("could not obtain OAuth2 token from JSON: %w", err)
	}
	return newOidcAuthorizerForAccountsAPI(oidcSource, creds.TokenSource), nil
}

// Reads credentials as JSON. Credentials can be either a path to JSON file,
// or actual JSON string.
func readCredentials(credentials string) (string, error) {
	// Try to read credentials as file path.
	if _, err := os.Stat(credentials); err == nil {
		jsonContents, err := ioutil.ReadFile(credentials)
		if err != nil {
			return string(jsonContents), err
		}
		return string(jsonContents), nil
	}
	// Assume that credential is actually JSON string.
	return credentials, nil
}

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
