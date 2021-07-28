package common

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"google.golang.org/api/impersonate"
)

func (c *DatabricksClient) configureWithGoogleServiceAccount() (func(r *http.Request) error, error) {
	if c.GoogleServiceAccount == "" {
		return nil, nil
	}
	if !c.IsGcp() {
		return nil, nil
	}
	if !c.isAccountsClient() {
		return nil, nil
	}
	// source for generateIdToken
	oidcSource, err := impersonate.IDTokenSource(c.InitContext, impersonate.IDTokenConfig{
		Audience:        "https://accounts.gcp.databricks.com",
		TargetPrincipal: c.GoogleServiceAccount,
		IncludeEmail:    true,
	})
	if err != nil {
		return nil, fmt.Errorf("could not obtain OIDC token. %w. Running 'gcloud auth application-default login' may help", err)
	}
	// TODO: verify that refreshers work...
	oidcSource = oauth2.ReuseTokenSource(nil, oidcSource)
	// source for generateAccessToken
	platformSource, err := impersonate.CredentialsTokenSource(c.InitContext, impersonate.CredentialsConfig{
		TargetPrincipal: c.GoogleServiceAccount,
		Scopes:          []string{
			"https://www.googleapis.com/auth/cloud-platform",
			"https://www.googleapis.com/auth/compute",
		},
	})
	if err != nil {
		return nil, err
	}
	return func(r *http.Request) error {
		oidc, err := oidcSource.Token()
		if err != nil {
			return err
		}
		cloudAccess, err := platformSource.Token()
		if err != nil {
			return err
		}
		oidc.SetAuthHeader(r)
		r.Header.Set("X-Databricks-GCP-SA-Access-Token", cloudAccess.AccessToken)
		return nil
	}, nil
}