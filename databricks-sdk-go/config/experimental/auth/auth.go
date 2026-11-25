// Package auth is an internal package that provides authentication utilities.
//
// IMPORTANT: This package is not meant to be used directly by consumers of the
// SDK and is subject to change without notice.
package auth

import (
	"context"

	"golang.org/x/oauth2"
)

// Credentials is anything that can return authentication headers.
type Credentials interface {
	AuthHeaders(context.Context) (map[string]string, error)
}

// CredentialsFn is an adapter to allow the use of ordinary functions as
// Credentials.
//
// Example:
//
//	   creds := CredentialsFn(func(ctx context.Context) (map[string]string, error) {
//			return map[string]string{"Authorization": "Bearer " + token}, nil
//	   })
type CredentialsFn func(context.Context) (map[string]string, error)

func (fn CredentialsFn) AuthHeaders(ctx context.Context) (map[string]string, error) {
	return fn(ctx)
}

// A TokenSource is anything that can return a token.
type TokenSource interface {
	// Token returns a token or an error. Token must be safe for concurrent use
	// by multiple goroutines. The returned Token must not be modified.
	Token(context.Context) (*oauth2.Token, error)
}

// TokenSourceFn is an adapter to allow the use of ordinary functions as
// TokenSource.
//
// Example:
//
//	   ts := TokenSourceFn(func(ctx context.Context) (*oauth2.Token, error) {
//			return &oauth2.Token{}, nil
//	   })
type TokenSourceFn func(context.Context) (*oauth2.Token, error)

func (fn TokenSourceFn) Token(ctx context.Context) (*oauth2.Token, error) {
	return fn(ctx)
}

// OAuthCredentials is a Credentials and TokenSource that can be used to
// authenticate with OAuth.
type OAuthCredentials interface {
	Credentials
	TokenSource
}
