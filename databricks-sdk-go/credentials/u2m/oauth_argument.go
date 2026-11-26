package u2m

// OAuthArgument is an interface that provides the necessary information to
// authenticate with PersistentAuth. Implementations of this interface must
// implement either the WorkspaceOAuthArgument or AccountOAuthArgument
// interface.
type OAuthArgument interface {
	// GetCacheKey returns a unique key for the OAuthArgument. This key is used
	// to store and retrieve the token from the token cache.
	GetCacheKey() string
}

// HostCacheKeyProvider is an interface for OAuthArgument implementations that
// can return a host-based cache key regardless of whether a profile is set.
//
// PersistentAuth itself no longer uses this key; it is exported so that
// external token cache implementations (for example, the CLI's file-based
// cache) can type-assert on it to mirror tokens under the host key for
// cross-SDK compatibility with older SDKs that only know host keys.
type HostCacheKeyProvider interface {
	GetHostCacheKey() string
}
