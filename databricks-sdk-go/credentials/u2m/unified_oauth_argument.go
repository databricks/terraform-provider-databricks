package u2m

import (
	"fmt"
)

// UnifiedOAuthArgument is an interface that provides the necessary information
// to authenticate using OAuth to a host that supports both account and workspace APIs.
type UnifiedOAuthArgument interface {
	OAuthArgument

	// GetHost returns the host to authenticate to.
	GetHost() string

	// GetAccountId returns the account ID of the account to authenticate to.
	GetAccountId() string
}

// BasicUnifiedOAuthArgument is a basic implementation of the UnifiedOAuthArgument
// interface that links each account with exactly one OAuth token.
type BasicUnifiedOAuthArgument struct {
	host      string
	accountID string

	// profile is the optional profile name. When set, GetCacheKey() returns
	// the profile name instead of the host-based key.
	profile string
}

var _ UnifiedOAuthArgument = BasicUnifiedOAuthArgument{}
var _ HostCacheKeyProvider = BasicUnifiedOAuthArgument{}

// NewBasicUnifiedOAuthArgument creates a new BasicUnifiedOAuthArgument.
func NewBasicUnifiedOAuthArgument(accountsHost, accountID string) (BasicUnifiedOAuthArgument, error) {
	return NewProfileUnifiedOAuthArgument(accountsHost, accountID, "")
}

// NewProfileUnifiedOAuthArgument creates a new BasicUnifiedOAuthArgument with a
// profile name. When a profile is set, GetCacheKey() returns the profile name
// instead of the host-based key.
func NewProfileUnifiedOAuthArgument(accountsHost, accountID, profile string) (BasicUnifiedOAuthArgument, error) {
	if err := validateHost(accountsHost); err != nil {
		return BasicUnifiedOAuthArgument{}, err
	}
	return BasicUnifiedOAuthArgument{host: accountsHost, accountID: accountID, profile: profile}, nil
}

// GetHost returns the host to authenticate to.
func (a BasicUnifiedOAuthArgument) GetHost() string {
	return a.host
}

// GetAccountId returns the account ID of the account to authenticate to.
func (a BasicUnifiedOAuthArgument) GetAccountId() string {
	return a.accountID
}

// GetCacheKey returns a unique key for caching the OAuth token for the account.
// If a profile is set, the profile name is returned as the cache key.
// Otherwise, the key is in the format "<hostName>/oidc/accounts/<accountID>".
func (a BasicUnifiedOAuthArgument) GetCacheKey() string {
	if a.profile != "" {
		return a.profile
	}
	return a.GetHostCacheKey()
}

// GetHostCacheKey returns the host-based cache key regardless of whether a
// profile is set. The key is in the format "<hostName>/oidc/accounts/<accountID>".
func (a BasicUnifiedOAuthArgument) GetHostCacheKey() string {
	return fmt.Sprintf("%s/oidc/accounts/%s", a.host, a.accountID)
}
