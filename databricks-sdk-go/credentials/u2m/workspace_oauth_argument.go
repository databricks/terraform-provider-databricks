package u2m

import (
	"fmt"
	"strings"
)

// WorkspaceOAuthArgument is an interface that provides the necessary information
// to authenticate using OAuth to a specific workspace.
type WorkspaceOAuthArgument interface {
	OAuthArgument

	// GetWorkspaceHost returns the host of the workspace to authenticate to.
	GetWorkspaceHost() string
}

// BasicWorkspaceOAuthArgument is a basic implementation of the WorkspaceOAuthArgument
// interface that links each host with exactly one OAuth token.
type BasicWorkspaceOAuthArgument struct {
	// host is the host of the workspace to authenticate to. This must start
	// with "https://" and must not have a trailing slash.
	host string

	// profile is the optional profile name. When set, GetCacheKey() returns
	// the profile name instead of the host-based key.
	profile string
}

func validateHost(host string) error {
	// Allow http for localhost. This is necessary for local end to end testing
	// of the `databricks auth login` command using a test server on localhost.
	if strings.HasPrefix(host, "http://127.0.0.1") {
		return nil
	}
	if !strings.HasPrefix(host, "https://") {
		return fmt.Errorf("host must start with 'https://': %s", host)
	}
	if strings.HasSuffix(host, "/") {
		return fmt.Errorf("host must not have a trailing slash: %s", host)
	}
	return nil
}

// NewBasicWorkspaceOAuthArgument creates a new BasicWorkspaceOAuthArgument.
func NewBasicWorkspaceOAuthArgument(host string) (BasicWorkspaceOAuthArgument, error) {
	return NewProfileWorkspaceOAuthArgument(host, "")
}

// NewProfileWorkspaceOAuthArgument creates a new BasicWorkspaceOAuthArgument
// with a profile name. When a profile is set, GetCacheKey() returns the profile
// name instead of the host-based key.
func NewProfileWorkspaceOAuthArgument(host, profile string) (BasicWorkspaceOAuthArgument, error) {
	if err := validateHost(host); err != nil {
		return BasicWorkspaceOAuthArgument{}, err
	}
	return BasicWorkspaceOAuthArgument{host: host, profile: profile}, nil
}

// GetWorkspaceHost returns the host of the workspace to authenticate to.
func (a BasicWorkspaceOAuthArgument) GetWorkspaceHost() string {
	return a.host
}

// GetCacheKey returns a unique key for caching the OAuth token for the workspace.
// If a profile is set, the profile name is returned as the cache key.
// Otherwise, the key is in the format "<host>".
func (a BasicWorkspaceOAuthArgument) GetCacheKey() string {
	if a.profile != "" {
		return a.profile
	}
	return a.GetHostCacheKey()
}

// GetHostCacheKey returns the host-based cache key regardless of whether a
// profile is set. The key is in the format "<host>".
func (a BasicWorkspaceOAuthArgument) GetHostCacheKey() string {
	host := strings.TrimSuffix(a.host, "/")
	if !strings.HasPrefix(host, "http") {
		host = fmt.Sprintf("https://%s", host)
	}
	return host
}

var _ WorkspaceOAuthArgument = BasicWorkspaceOAuthArgument{}
var _ HostCacheKeyProvider = BasicWorkspaceOAuthArgument{}
