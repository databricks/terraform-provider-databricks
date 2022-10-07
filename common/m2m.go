package common

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type oauthAuthorizationServer struct {
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
}

func (c *DatabricksClient) getOAuthEndpoints() (*oauthAuthorizationServer, error) {
	err := c.fixHost()
	if err != nil {
		return nil, fmt.Errorf("host: %w", err)
	}
	oidc := fmt.Sprintf("%s/oidc/.well-known/oauth-authorization-server", c.Host)
	oidcResponse, err := http.Get(oidc)
	if err != nil {
		return nil, fmt.Errorf("fetch .well-known: %w", err)
	}
	if oidcResponse.Body == nil {
		return nil, fmt.Errorf("fetch .well-known: empty body")
	}
	defer oidcResponse.Body.Close()
	raw, err := io.ReadAll(oidcResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("read .well-known: %w", err)
	}
	var oauthEndpoints oauthAuthorizationServer
	err = json.Unmarshal(raw, &oauthEndpoints)
	if err != nil {
		return nil, fmt.Errorf("parse .well-known: %w", err)
	}
	return &oauthEndpoints, nil
}

func (c *DatabricksClient) configureWithOAuthM2M(
	ctx context.Context) (func(r *http.Request) error, error) {
	if !c.IsAws() || c.ClientID == "" || c.ClientSecret == "" || c.Host == "" {
		return nil, nil
	}
	// workaround for accounts endpoint not having yet a well-known OIDC alias
	if c.TokenEndpoint == "" {
		endpoints, err := c.getOAuthEndpoints()
		if err != nil {
			return nil, fmt.Errorf("databricks oauth: %w", err)
		}
		c.TokenEndpoint = endpoints.TokenEndpoint
	}
	log.Printf("[INFO] Generating Databricks OAuth token for Service Principal (%s)", c.ClientID)
	ts := (&clientcredentials.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		AuthStyle:    oauth2.AuthStyleInHeader,
		TokenURL:     c.TokenEndpoint,
		Scopes:       []string{"all-apis"},
	}).TokenSource(ctx)
	return newOidcAuthorizerWithJustBearer(ts), nil
}
