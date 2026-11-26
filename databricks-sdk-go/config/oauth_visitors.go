package config

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
)

// cacheOptions returns the auth.Option values derived from the given Config.
func cacheOptions(cfg *Config) []auth.Option {
	return []auth.Option{
		auth.WithAsyncRefresh(!cfg.DisableAsyncTokenRefresh),
	}
}

// serviceToServiceVisitor returns a visitor that sets the Authorization header
// to the token from the primary token source and the provided secondary header
// to the token from the secondary token source. If secondaryOptional is true,
// a failure to get the secondary token logs a warning and skips the header
// instead of returning an error.
func serviceToServiceVisitor(primary, secondary auth.TokenSource, secondaryHeader string, secondaryOptional bool, opts ...auth.Option) func(r *http.Request) error {
	refreshableAuth := auth.NewCachedTokenSource(primary, opts...)
	refreshableSecondary := auth.NewCachedTokenSource(secondary, opts...)
	return func(r *http.Request) error {
		inner, err := refreshableAuth.Token(r.Context())
		if err != nil {
			return fmt.Errorf("inner token: %w", err)
		}
		inner.SetAuthHeader(r)

		cloud, err := refreshableSecondary.Token(r.Context())
		if err != nil {
			if secondaryOptional {
				logger.Warnf(r.Context(), "Failed to get secondary token for %s header: %v. Skipping.", secondaryHeader, err)
				return nil
			}
			return fmt.Errorf("cloud token: %w", err)
		}
		r.Header.Set(secondaryHeader, cloud.AccessToken)
		return nil
	}
}

// refreshableAuthVisitor returns a visitor that sets the Authorization header
// to the token from the given token source.
func refreshableAuthVisitor(inner auth.TokenSource, opts ...auth.Option) func(r *http.Request) error {
	cts := auth.NewCachedTokenSource(inner, opts...)
	return func(r *http.Request) error {
		inner, err := cts.Token(r.Context())
		if err != nil {
			return fmt.Errorf("inner token: %w", err)
		}
		inner.SetAuthHeader(r)
		return nil
	}
}

func azureVisitor(cfg *Config, inner func(*http.Request) error) func(*http.Request) error {
	return func(r *http.Request) error {
		if cfg.AzureResourceID != "" {
			r.Header.Set(xDatabricksAzureWorkspaceResourceId, cfg.AzureResourceID)
		}
		return inner(r)
	}
}

// azureReuseTokenSource returns a cached token source that refreshes token 40
// seconds before they expire. The reason for this is that Azure Databricks
// rejects tokens that expire in 30 seconds or less and we want to give a 10
// second buffer.
func azureReuseTokenSource(t *oauth2.Token, ts auth.TokenSource, opts ...auth.Option) auth.TokenSource {
	early := wrap(ts, func(t *oauth2.Token) *oauth2.Token {
		t.Expiry = t.Expiry.Add(-40 * time.Second)
		return t
	})

	allOpts := append([]auth.Option{auth.WithCachedToken(t)}, opts...)
	return auth.NewCachedTokenSource(early, allOpts...)
}

func wrap(ts auth.TokenSource, fn func(*oauth2.Token) *oauth2.Token) auth.TokenSource {
	return &tokenSourceWrapper{fn: fn, inner: ts}
}

type tokenSourceWrapper struct {
	fn    func(*oauth2.Token) *oauth2.Token
	inner auth.TokenSource
}

func (w *tokenSourceWrapper) Token(ctx context.Context) (*oauth2.Token, error) {
	t, err := w.inner.Token(ctx)
	if err != nil {
		return nil, err
	}
	return w.fn(t), nil
}

// newVisitorOAuthCredentials creates an OAuthCredentialsProvider from a visitor
// function and a token source. The visitor is used to set headers on the
// request, and the token source is used to provide the token.
func newVisitorOAuthCredentials(visitor func(r *http.Request) error, ts auth.TokenSource) credentials.OAuthCredentialsProvider {
	return &visitorOAuthCredentials{visitor: visitor, ts: ts}
}

type visitorOAuthCredentials struct {
	visitor func(r *http.Request) error
	ts      auth.TokenSource
}

func (c *visitorOAuthCredentials) SetHeaders(r *http.Request) error {
	return c.visitor(r)
}

func (c *visitorOAuthCredentials) Token(ctx context.Context) (*oauth2.Token, error) {
	return c.ts.Token(ctx)
}
