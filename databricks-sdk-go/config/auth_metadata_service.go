package config

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
)

// timeout to wait for metadata service to respond
const metadataServiceTimeout = 10 * time.Second

const MetadataServiceVersion = "1"
const MetadataServiceVersionHeader = "X-Databricks-Metadata-Version"
const MetadataServiceHostHeader = "X-Databricks-Host"

var errMetadataServiceMalformed = errors.New("invalid auth server URL")
var errMetadataServiceNotLocalhost = errors.New("only localhost URLs are allowed")

// Credentials provider that fetches a token from a locally running HTTP server
//
// The credentials provider will perform a GET request to the configured URL.
//
// The MUST return 4xx response if the "X-Databricks-Metadata-Version" header
// is not set or set to a version that the server doesn't support.
//
// The server MUST guarantee stable sessions per URL path. That is, if the
// server returns a token for a Host on a given URL path, it MUST continue to return
// tokens for the same Host.
//
// The server MUST return a 4xx response if the Host passed in the "X-Databricks-Host"
// header doesn't match the token.
//
// The server is expected to return a JSON response with the following fields:
//
// - access_token: The requested access token.
// - token_type: The type of token, which is a "Bearer" access token.
// - expires_on: Unix timestamp when the access token expires.
type MetadataServiceCredentials struct{}

func (c MetadataServiceCredentials) Name() string {
	return "metadata-service"
}

func (c MetadataServiceCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if err := unsupportedGroupRoleAssumption(cfg, c.Name()); err != nil {
		return nil, err
	}
	if cfg.MetadataServiceURL == "" || cfg.Host == "" {
		return nil, nil
	}
	parsedMetadataServiceURL, err := url.Parse(cfg.MetadataServiceURL)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errMetadataServiceMalformed, err)
	}

	// Only allow localhost URLs.
	if parsedMetadataServiceURL.Hostname() != "localhost" && parsedMetadataServiceURL.Hostname() != "127.0.0.1" {
		return nil, fmt.Errorf("%w: %s", errMetadataServiceNotLocalhost, cfg.MetadataServiceURL)
	}

	ms := metadataService{
		metadataServiceURL: parsedMetadataServiceURL,
		config:             cfg,
	}

	// Sanity check that a token can be obtained.
	//
	// TODO: Move this outside of this function. If credentials providers have
	// to be tested, this should be done in the main default loop, not here.
	if _, err := ms.Token(ctx); err != nil {
		return nil, err
	}

	cts := auth.NewCachedTokenSource(ms, cacheOptions(cfg)...)
	return credentials.NewOAuthCredentialsProviderFromTokenSource(cts), nil
}

type metadataService struct {
	metadataServiceURL *url.URL
	config             *Config
}

func (ms metadataService) Token(ctx context.Context) (*oauth2.Token, error) {
	ctx, cancel := context.WithTimeout(ctx, metadataServiceTimeout)
	defer cancel()

	var mt msiToken
	err := ms.config.refreshClient.Do(ctx, http.MethodGet,
		ms.metadataServiceURL.String(),
		httpclient.WithRequestHeader(MetadataServiceVersionHeader, MetadataServiceVersion),
		httpclient.WithRequestHeader(MetadataServiceHostHeader, ms.config.Host),
		httpclient.WithResponseUnmarshal(&mt),
	)
	if err != nil {
		return nil, fmt.Errorf("token request: %w", err)
	}

	token, err := mt.Token()
	if err != nil {
		return nil, fmt.Errorf("token parse: %w", err)
	}

	logger.Debugf(ctx, "Refreshed access token from local metadata service, which expires on %s", token.Expiry.Format(time.RFC3339))
	return token, nil
}
