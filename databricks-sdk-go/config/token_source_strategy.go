package config

import (
	"context"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
)

// Creates a CredentialsStrategy from a TokenSource.
func NewTokenSourceStrategy(name string, ts auth.TokenSource) CredentialsStrategy {
	return &tokenSourceStrategy{name: name, ts: ts}
}

// tokenSourceStrategy is wrapper on a auth.TokenSource which converts it into
// a CredentialsStrategy.
type tokenSourceStrategy struct {
	name string
	ts   auth.TokenSource
}

// Configure implements [CredentialsStrategy.Configure].
func (tss *tokenSourceStrategy) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	cp := credentials.NewOAuthCredentialsProviderFromTokenSource(
		auth.NewCachedTokenSource(tss.ts, cacheOptions(cfg)...),
	)

	// Sanity check that a token can be obtained.
	//
	// TODO: Move this outside of this function. If credentials providers have
	// to be tested, this should be done in the main default loop, not here.
	if _, err := cp.Token(ctx); err != nil {
		return nil, err
	}

	return cp, nil
}

// Name implements [CredentialsStrategy.Name].
func (t *tokenSourceStrategy) Name() string {
	return t.name
}
