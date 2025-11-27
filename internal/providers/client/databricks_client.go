package client

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/commands"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// HTTPConfig holds configuration for custom HTTP transport behavior.
type HTTPConfig struct {
	// Headers are custom HTTP headers added to all API requests.
	Headers map[string]string
	// PathPrefix is prepended to all API request URL paths.
	PathPrefix string
}

// httpTransportWrapper wraps an http.RoundTripper to add custom headers
// and path prefix to all requests.
type httpTransportWrapper struct {
	base       http.RoundTripper
	headers    map[string]string
	pathPrefix string
}

// RoundTrip implements http.RoundTripper interface.
func (t *httpTransportWrapper) RoundTrip(req *http.Request) (*http.Response, error) {
	// Add custom headers
	for key, val := range t.headers {
		req.Header.Set(key, val)
	}
	// Add path prefix
	if t.pathPrefix != "" {
		req.URL.Path = strings.TrimSuffix(t.pathPrefix, "/") + req.URL.Path
	}
	return t.base.RoundTrip(req)
}

// PrepareDatabricksClient makes some common adjustments to the config that apply in all cases
// and returns a ready-to-use Databricks client. This includes:
// - mapping deprecated auth types to their newer counterparts
// - ensuring the config is resolved
// - setting a default retry timeout if not set
// - setting a default HTTP timeout if not set
// - configuring custom HTTP headers and path prefix if provided
//
// TODO: this should be colocated with the definition of DatabricksClient in common/client.go, but
// this isn't possible without introducing a circular dependency. Fixing this will require refactoring
// DatabricksClient out of the common package.
func PrepareDatabricksClient(ctx context.Context, cfg *config.Config, configCustomizer func(*config.Config) error, httpConfig *HTTPConfig) (*common.DatabricksClient, error) {
	if cfg.AuthType != "" {
		// mapping from previous Google authentication types
		// and current authentication types from Databricks Go SDK
		oldToNewerAuthType := map[string]string{
			"google-creds":     "google-credentials",
			"google-accounts":  "google-id",
			"google-workspace": "google-id",
		}
		newer, ok := oldToNewerAuthType[cfg.AuthType]
		if ok {
			tflog.Info(ctx, fmt.Sprintf("Changing required auth_type from %s to %s", cfg.AuthType, newer))
			cfg.AuthType = newer
		}
	}
	cfg.EnsureResolved()
	// Unless set explicitly, the provider will retry indefinitely until context is cancelled
	// by either a timeout or interrupt.
	if cfg.RetryTimeoutSeconds == 0 {
		cfg.RetryTimeoutSeconds = -1
	}
	// If not set, the default provider timeout is 65 seconds. Most APIs have a server-side timeout of 60 seconds.
	// The additional 5 seconds is to account for network latency.
	if cfg.HTTPTimeoutSeconds == 0 {
		cfg.HTTPTimeoutSeconds = 65
	}
	if configCustomizer != nil {
		err := configCustomizer(cfg)
		if err != nil {
			return nil, err
		}
	}
	// Set up custom HTTP transport if headers or path prefix are specified
	if httpConfig != nil && (len(httpConfig.Headers) > 0 || httpConfig.PathPrefix != "") {
		baseTransport := cfg.HTTPTransport
		if baseTransport == nil {
			baseTransport = http.DefaultTransport
		}
		cfg.HTTPTransport = &httpTransportWrapper{
			base:       baseTransport,
			headers:    httpConfig.Headers,
			pathPrefix: httpConfig.PathPrefix,
		}
		if len(httpConfig.Headers) > 0 {
			tflog.Debug(ctx, fmt.Sprintf("Custom HTTP headers configured: %d header(s)", len(httpConfig.Headers)))
		}
		if httpConfig.PathPrefix != "" {
			tflog.Debug(ctx, fmt.Sprintf("Custom HTTP path prefix configured: %s", httpConfig.PathPrefix))
		}
	}
	client, err := client.New(cfg)
	if err != nil {
		return nil, err
	}
	pc := &common.DatabricksClient{
		DatabricksClient: client,
	}
	pc.WithCommandExecutor(func(ctx context.Context, client *common.DatabricksClient) common.CommandExecutor {
		return commands.NewCommandsAPI(ctx, client)
	})
	return pc, nil
}
