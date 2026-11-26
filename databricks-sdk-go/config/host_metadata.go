package config

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/environment"
	"github.com/databricks/databricks-sdk-go/httpclient"
)

// HostMetadata holds the parsed response from the /.well-known/databricks-config discovery endpoint.
type HostMetadata struct {
	// OIDCEndpoint is the OIDC discovery URL for this host. For account hosts,
	// this may contain an {account_id} placeholder that callers must substitute.
	OIDCEndpoint string `json:"oidc_endpoint"`

	// AccountID is the Databricks account ID associated with this host, if available.
	AccountID string `json:"account_id"`

	// WorkspaceID is the Databricks workspace ID associated with this host, if available.
	WorkspaceID string `json:"workspace_id"`

	// Cloud is the cloud provider for this Databricks deployment (AWS, Azure, or GCP).
	Cloud environment.Cloud `json:"cloud"`

	// HostType is the type of host (WORKSPACE_HOST, ACCOUNT_HOST, or UNIFIED_HOST).
	HostType HostType `json:"host_type"`

	// TokenFederationDefaultOIDCAudiences is the default OIDC audience for token requests.
	// For workspace hosts: "https://<workspace_host>/oidc/v1/token"
	// For account/unified hosts: the resolved account ID.
	TokenFederationDefaultOIDCAudiences []string `json:"token_federation_default_oidc_audiences"`
}

// HostMetadataResolver, when set on [Config], overrides the default HTTP fetch
// of /.well-known/databricks-config during config resolution. The function
// receives the canonical host and returns metadata (or nil, nil if unavailable).
// This allows callers to provide cached metadata without the SDK making an HTTP call.
type HostMetadataResolver func(ctx context.Context, host string) (*HostMetadata, error)

// DefaultHostMetadataResolverFactory is consulted by [Config.EnsureResolved]
// when [Config.HostMetadataResolver] is nil. When set, the factory is invoked
// with the resolving Config and must return the resolver to use for that
// Config (or nil to fall through to the SDK's default HTTP fetch).
//
// Intended for programs that want a single hook to install a caching or
// otherwise-customised resolver across every Config they construct, without
// per-site wiring. Set once from an init() block in a package that is
// blank-imported by the main binary. Callers needing a per-Config resolver
// should use [Config.HostMetadataResolver] instead.
//
// Experimental: subject to change.
var DefaultHostMetadataResolverFactory func(*Config) HostMetadataResolver

// getHostMetadata fetches the raw Databricks well-known configuration from
// {host}/.well-known/databricks-config. The returned HostMetadata contains
// raw values with no substitution (e.g., {account_id} placeholders are left
// as-is). Callers are responsible for interpreting the result.
func getHostMetadata(ctx context.Context, host string, client *httpclient.ApiClient) (*HostMetadata, error) {
	discoveryURL := host + "/.well-known/databricks-config"
	var meta HostMetadata
	if err := client.Do(ctx, "GET", discoveryURL, httpclient.WithResponseUnmarshal(&meta)); err != nil {
		return nil, fmt.Errorf("fetching host metadata from %q: %w", discoveryURL, err)
	}
	return &meta, nil
}
