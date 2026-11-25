/*
Package cache provides an interface for storing and looking up OAuth tokens.

The cache should be primarily used for user-to-machine (U2M) OAuth flows. In U2M
OAuth flows, the application needs to store the token for later use, such as in
a separate process, and the cache provides a way to do so without requiring the
user to follow the OAuth flow again.

In machine-to-machine (M2M) OAuth flows, the application is configured with a
secret and can fetch a new token on demand without user interaction, so the
token cache is not necessary.
*/
package cache

import (
	"errors"

	"golang.org/x/oauth2"
)

var ErrNotFound = errors.New("token not found")

// TokenCache is an interface for storing and looking up OAuth tokens.
type TokenCache interface {
	// Store stores the token with the given key, replacing any existing token.
	// If t is nil, it deletes the token.
	Store(key string, t *oauth2.Token) error

	// Lookup looks up the token with the given key. If the token is not found, it
	// returns ErrNotFound.
	Lookup(key string) (*oauth2.Token, error)
}
