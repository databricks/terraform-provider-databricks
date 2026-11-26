package u2m

import (
	"fmt"
)

// AccountOAuthArgument is an interface that provides the necessary information
// to authenticate using OAuth to a specific account.
type AccountOAuthArgument interface {
	OAuthArgument

	// GetAccountHost returns the host of the account to authenticate to.
	GetAccountHost() string

	// GetAccountId returns the account ID of the account to authenticate to.
	GetAccountId() string
}

// BasicAccountOAuthArgument is a basic implementation of the AccountOAuthArgument
// interface that links each account with exactly one OAuth token.
type BasicAccountOAuthArgument struct {
	accountHost string
	accountID   string

	// profile is the optional profile name. When set, GetCacheKey() returns
	// the profile name instead of the host-based key.
	profile string
}

var _ AccountOAuthArgument = BasicAccountOAuthArgument{}
var _ HostCacheKeyProvider = BasicAccountOAuthArgument{}

// NewBasicAccountOAuthArgument creates a new BasicAccountOAuthArgument.
func NewBasicAccountOAuthArgument(accountsHost, accountID string) (BasicAccountOAuthArgument, error) {
	return NewProfileAccountOAuthArgument(accountsHost, accountID, "")
}

// NewProfileAccountOAuthArgument creates a new BasicAccountOAuthArgument with a
// profile name. When a profile is set, GetCacheKey() returns the profile name
// instead of the host-based key.
func NewProfileAccountOAuthArgument(accountsHost, accountID, profile string) (BasicAccountOAuthArgument, error) {
	if err := validateHost(accountsHost); err != nil {
		return BasicAccountOAuthArgument{}, err
	}
	return BasicAccountOAuthArgument{accountHost: accountsHost, accountID: accountID, profile: profile}, nil
}

// GetAccountHost returns the host of the account to authenticate to.
func (a BasicAccountOAuthArgument) GetAccountHost() string {
	return a.accountHost
}

// GetAccountId returns the account ID of the account to authenticate to.
func (a BasicAccountOAuthArgument) GetAccountId() string {
	return a.accountID
}

// GetCacheKey returns a unique key for caching the OAuth token for the account.
// If a profile is set, the profile name is returned as the cache key.
// Otherwise, the key is in the format "<accountHost>/oidc/accounts/<accountID>".
func (a BasicAccountOAuthArgument) GetCacheKey() string {
	if a.profile != "" {
		return a.profile
	}
	return a.GetHostCacheKey()
}

// GetHostCacheKey returns the host-based cache key regardless of whether a
// profile is set. The key is in the format "<accountHost>/oidc/accounts/<accountID>".
func (a BasicAccountOAuthArgument) GetHostCacheKey() string {
	return fmt.Sprintf("%s/oidc/accounts/%s", a.accountHost, a.accountID)
}
