package u2m

import (
	"context"
	"errors"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestBasicOAuthClient_GetAccountOAuthEndpoints(t *testing.T) {
	c := &BasicOAuthEndpointSupplier{}
	s, err := c.GetAccountOAuthEndpoints(context.Background(), "https://abc", "xyz")
	assert.NoError(t, err)
	assert.Equal(t, "https://abc/oidc/accounts/xyz/v1/authorize", s.AuthorizationEndpoint)
	assert.Equal(t, "https://abc/oidc/accounts/xyz/v1/token", s.TokenEndpoint)
}

func TestGetWorkspaceOAuthEndpoints(t *testing.T) {
	p := httpclient.NewApiClient(httpclient.ClientConfig{
		Transport: fixtures.MappingTransport{
			"GET /oidc/.well-known/oauth-authorization-server": {
				Status: 200,
				Response: map[string]string{
					"authorization_endpoint": "a",
					"token_endpoint":         "b",
				},
			},
		},
	})
	c := &BasicOAuthEndpointSupplier{Client: p}
	endpoints, err := c.GetWorkspaceOAuthEndpoints(context.Background(), "https://abc")
	assert.NoError(t, err)
	assert.Equal(t, "a", endpoints.AuthorizationEndpoint)
	assert.Equal(t, "b", endpoints.TokenEndpoint)
}

func TestGetEndpointsFromURL(t *testing.T) {
	p := httpclient.NewApiClient(httpclient.ClientConfig{
		Transport: fixtures.MappingTransport{
			"GET /oidc/v1/authorize-server": {
				Status: 200,
				Response: map[string]string{
					"authorization_endpoint": "https://abc/oidc/v1/authorize",
					"token_endpoint":         "https://abc/oidc/v1/token",
				},
			},
		},
	})
	c := &BasicOAuthEndpointSupplier{Client: p}
	endpoints, err := c.GetEndpointsFromURL(context.Background(), "https://abc/oidc/v1/authorize-server")
	if err != nil {
		t.Fatal(err)
	}
	if endpoints.AuthorizationEndpoint != "https://abc/oidc/v1/authorize" {
		t.Errorf("unexpected AuthorizationEndpoint: %q", endpoints.AuthorizationEndpoint)
	}
	if endpoints.TokenEndpoint != "https://abc/oidc/v1/token" {
		t.Errorf("unexpected TokenEndpoint: %q", endpoints.TokenEndpoint)
	}
}

func TestGetEndpointsFromURL_NotFound(t *testing.T) {
	p := httpclient.NewApiClient(httpclient.ClientConfig{
		Transport: fixtures.MappingTransport{
			"GET /oidc/v1/authorize-server": {
				Status: 404,
			},
		},
		// Mirror the error mapping done by Config.refreshTokenErrorMapper: map
		// HTTP 404 to apierr.ErrNotFound so that errors.Is(err, apierr.ErrNotFound)
		// returns true, matching the production code path.
		ErrorMapper: func(ctx context.Context, resp common.ResponseWrapper) error {
			err := httpclient.DefaultErrorMapper(ctx, resp)
			if err == nil {
				return nil
			}
			var httpErr *httpclient.HttpError
			if errors.As(err, &httpErr) {
				if sdkErr, ok := apierr.ByStatusCode(httpErr.StatusCode); ok {
					return sdkErr
				}
			}
			return err
		},
	})
	c := &BasicOAuthEndpointSupplier{Client: p}
	_, err := c.GetEndpointsFromURL(context.Background(), "https://abc/oidc/v1/authorize-server")
	if !errors.Is(err, ErrOAuthNotSupported) {
		t.Errorf("expected ErrOAuthNotSupported, got %v", err)
	}
}

func TestGetUnifiedOAuthEndpoints(t *testing.T) {
	p := httpclient.NewApiClient(httpclient.ClientConfig{
		Transport: fixtures.MappingTransport{
			"GET /oidc/accounts/xyz/.well-known/oauth-authorization-server": {
				Status: 200,
				Response: map[string]string{
					"authorization_endpoint": "https://abc/oidc/accounts/xyz/v1/authorize",
					"token_endpoint":         "https://abc/oidc/accounts/xyz/v1/token",
				},
			},
		},
	})
	c := &BasicOAuthEndpointSupplier{Client: p}
	endpoints, err := c.GetUnifiedOAuthEndpoints(context.Background(), "https://abc", "xyz")

	assert.NoError(t, err)
	assert.Equal(t, "https://abc/oidc/accounts/xyz/v1/authorize", endpoints.AuthorizationEndpoint)
	assert.Equal(t, "https://abc/oidc/accounts/xyz/v1/token", endpoints.TokenEndpoint)
}
