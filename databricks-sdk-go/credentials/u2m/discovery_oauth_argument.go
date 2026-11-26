package u2m

import "fmt"

// DiscoveryOAuthArgument is an OAuthArgument for bootstrapping the
// login.databricks.com discovery flow. Unlike other OAuthArgument types, it
// has no host at construction time. The host is discovered from the iss
// parameter in the OAuth callback after the user authenticates and selects a
// workspace.
//
// DiscoveryOAuthArgument is only intended for use with [WithDiscoveryLogin]
// during [PersistentAuth.Challenge]. Once the host has been discovered,
// callers should construct the usual host-based OAuthArgument
// (for example [WorkspaceOAuthArgument]) for future PersistentAuth instances.
type DiscoveryOAuthArgument interface {
	OAuthArgument

	// SetDiscoveredHost stores the workspace host derived from the iss
	// callback parameter.
	SetDiscoveredHost(host string)

	// GetDiscoveredHost returns the workspace host discovered from the
	// callback, or empty string if not yet discovered.
	GetDiscoveredHost() string
}

// BasicDiscoveryOAuthArgument is a basic implementation of
// [DiscoveryOAuthArgument] that uses the profile name as the cache key during
// discovery login bootstrap.
type BasicDiscoveryOAuthArgument struct {
	profile        string
	discoveredHost string
}

// NewBasicDiscoveryOAuthArgument creates a new [BasicDiscoveryOAuthArgument].
// The profile name is required and used as the cache key during discovery
// login. Use it with [WithDiscoveryLogin]. After discovery, construct the
// usual host-based OAuthArgument with the discovered host.
func NewBasicDiscoveryOAuthArgument(profile string) (*BasicDiscoveryOAuthArgument, error) {
	if profile == "" {
		return nil, fmt.Errorf("profile name must not be empty for discovery login")
	}
	return &BasicDiscoveryOAuthArgument{profile: profile}, nil
}

// GetCacheKey returns the profile name as the cache key.
func (a *BasicDiscoveryOAuthArgument) GetCacheKey() string {
	return a.profile
}

// SetDiscoveredHost stores the workspace host derived from the iss callback
// parameter.
func (a *BasicDiscoveryOAuthArgument) SetDiscoveredHost(host string) {
	a.discoveredHost = host
}

// GetDiscoveredHost returns the workspace host discovered from the callback,
// or empty string if not yet discovered.
func (a *BasicDiscoveryOAuthArgument) GetDiscoveredHost() string {
	return a.discoveredHost
}

var _ DiscoveryOAuthArgument = &BasicDiscoveryOAuthArgument{}
