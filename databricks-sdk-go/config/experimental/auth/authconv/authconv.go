package authconv

import (
	"context"

	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"golang.org/x/oauth2"
)

// AuthTokenSource converts an oauth2.TokenSource to an auth.TokenSource.
func AuthTokenSource(ts oauth2.TokenSource) auth.TokenSource {
	return &authTokenSource{ts: ts}
}

type authTokenSource struct {
	ts oauth2.TokenSource
}

func (t *authTokenSource) Token(_ context.Context) (*oauth2.Token, error) {
	return t.ts.Token()
}

// OAuth2TokenSource converts an auth.TokenSource to an oauth2.TokenSource.
func OAuth2TokenSource(ts auth.TokenSource) oauth2.TokenSource {
	return &oauth2TokenSource{ts: ts}
}

type oauth2TokenSource struct {
	ts auth.TokenSource
}

func (t *oauth2TokenSource) Token() (*oauth2.Token, error) {
	return t.ts.Token(context.Background())
}
