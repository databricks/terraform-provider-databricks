package auth

import (
	"context"
	"sync"
	"time"

	"golang.org/x/oauth2"
)

// Maximum time before expiry when async refreshes may start. This value is
// chosen to provide robustness for standard OAuth tokens, supporting 99.95%
// availability, adding breathing room over the existing 99.99% SLA.
const maxAsyncRefreshLeadTime = 20 * time.Minute

// Minimum wait time before retrying a failed async refresh. This prevents
// hammering a failing server while still recovering quickly from transient
// errors (e.g., brief network issues). With a 20-minute max refresh window,
// this allows up to ~20 retries before the token expires and forces a
// blocking call.
const asyncRefreshRetryBackoff = 1 * time.Minute

// Maximum time an async refresh is allowed to run before the child context
// is canceled.
const asyncRefreshTimeout = 1 * time.Minute

// computeAsyncRefreshLeadTime calculates how long before expiry async
// refreshes may start for a token with the given remaining TTL.
//
// Formula: min(TTL × 0.5, maxAsyncRefreshLeadTime)
//
// Edge cases:
//   - TTL <= 0 (expired or no expiry): returns 0.
//   - Very short TTL (e.g., 2 seconds): returns TTL / 2 without minimum enforcement.
//   - Standard OAuth (60 minutes): returns 20 minutes (capped at max).
//   - FastPath (10 minutes): returns 5 minutes.
func computeAsyncRefreshLeadTime(ttl time.Duration) time.Duration {
	if ttl <= 0 {
		return 0
	}

	leadTime := ttl / 2
	if leadTime > maxAsyncRefreshLeadTime {
		return maxAsyncRefreshLeadTime
	}
	return leadTime
}

type Option func(*cachedTokenSource)

// WithCachedToken sets the initial token to be used by a cached token source.
func WithCachedToken(t *oauth2.Token) Option {
	return func(cts *cachedTokenSource) {
		cts.cachedToken = t
	}
}

// WithAsyncRefresh enables or disables the asynchronous token refresh.
func WithAsyncRefresh(b bool) Option {
	return func(cts *cachedTokenSource) {
		cts.disableAsync = !b
	}
}

// NewCachedTokenSource wraps a [TokenSource] to cache the tokens it returns.
// By default, the cache will refresh tokens asynchronously a few minutes before
// they expire.
//
// The token cache is safe for concurrent use by multiple goroutines and will
// guarantee that only one token refresh is triggered at a time.
//
// The token cache does not take care of retries in case the token source
// returns and error; it is the responsibility of the provided token source to
// handle retries appropriately.
//
// If the TokenSource is already a cached token source (obtained by calling this
// function), it is returned as is.
func NewCachedTokenSource(ts TokenSource, opts ...Option) TokenSource {
	// This is meant as a niche optimization to avoid double caching of the
	// token source in situations where the user calls needs caching guarantees
	// but does not know if the token source is already cached.
	if cts, ok := ts.(*cachedTokenSource); ok {
		return cts
	}

	cts := &cachedTokenSource{
		tokenSource: ts,
		timeNow:     time.Now,
	}

	for _, opt := range opts {
		opt(cts)
	}

	// If an initial token was provided via WithCachedToken, compute the time
	// after which async refreshes may be attempted based on the token TTL
	// available at construction time.
	if cts.cachedToken != nil {
		cts.updateNextAsyncRefresh()
	}

	return cts
}

type cachedTokenSource struct {
	// The token source to obtain tokens from.
	tokenSource TokenSource

	// If true, only refresh the token with a blocking call when it is expired.
	disableAsync bool

	// Earliest time when an async refresh may be attempted for cachedToken.
	// Computed from the token TTL at the moment the token is cached. If an
	// async refresh fails, this timestamp is pushed forward by
	// asyncRefreshRetryBackoff to delay the next retry.
	nextAsyncRefresh time.Time

	mu          sync.Mutex
	cachedToken *oauth2.Token

	// Indicates that an async refresh is in progress. This is used to prevent
	// multiple async refreshes from being triggered at the same time.
	isRefreshing bool

	timeNow func() time.Time // for testing
}

// Token returns a token from the cache or fetches a new one if the current
// token is expired.
func (cts *cachedTokenSource) Token(ctx context.Context) (*oauth2.Token, error) {
	if cts.disableAsync {
		return cts.blockingToken(ctx)
	}
	return cts.asyncToken(ctx)
}

// updateNextAsyncRefresh recomputes nextAsyncRefresh from the current
// cachedToken's TTL. It must be called immediately after cachedToken is set.
// When called on a shared instance (after construction), the lock must be
// held.
func (cts *cachedTokenSource) updateNextAsyncRefresh() {
	if cts.cachedToken == nil || cts.cachedToken.Expiry.IsZero() {
		cts.nextAsyncRefresh = time.Time{}
		return
	}
	ttl := cts.cachedToken.Expiry.Sub(cts.timeNow())
	cts.nextAsyncRefresh = cts.cachedToken.Expiry.Add(-computeAsyncRefreshLeadTime(ttl))
}

// tokenExpired reports whether the current token is missing or unusable. The
// function is not thread-safe and should be called with the lock held.
func (cts *cachedTokenSource) tokenExpired() bool {
	if cts.cachedToken == nil {
		return true
	}
	if cts.cachedToken.Expiry.IsZero() {
		return false // zero expiry means that token is permanently valid.
	}
	return cts.timeNow().After(cts.cachedToken.Expiry)
}

// canRefreshAsync reports whether the current token has entered the async
// refresh window. The function is not thread-safe and should be called with
// the lock held.
func (cts *cachedTokenSource) canRefreshAsync() bool {
	if cts.cachedToken == nil || cts.cachedToken.Expiry.IsZero() {
		return false
	}
	return cts.timeNow().After(cts.nextAsyncRefresh)
}

func (cts *cachedTokenSource) asyncToken(ctx context.Context) (*oauth2.Token, error) {
	cts.mu.Lock()
	t := cts.cachedToken
	tokenExpired := cts.tokenExpired()
	canRefreshAsync := !tokenExpired && cts.canRefreshAsync()
	cts.mu.Unlock()

	if tokenExpired {
		return cts.blockingToken(ctx)
	}
	if canRefreshAsync {
		cts.triggerAsyncRefresh(ctx)
	}
	return t, nil
}

func (cts *cachedTokenSource) blockingToken(ctx context.Context) (*oauth2.Token, error) {
	cts.mu.Lock()

	// The lock is kept for the entire operation to ensure that only one
	// blockingToken operation is running at a time.
	defer cts.mu.Unlock()

	// It's possible that the token got refreshed (either by a blockingToken or
	// an asyncRefresh call) while this particular call was waiting to acquire
	// the mutex. This check avoids refreshing the token again in such cases.
	if !cts.tokenExpired() {
		return cts.cachedToken, nil
	}

	t, err := cts.tokenSource.Token(ctx)
	if err != nil {
		return nil, err
	}
	cts.cachedToken = t
	cts.updateNextAsyncRefresh()
	return t, nil
}

func (cts *cachedTokenSource) triggerAsyncRefresh(ctx context.Context) {
	cts.mu.Lock()
	defer cts.mu.Unlock()
	if cts.isRefreshing || cts.tokenExpired() || !cts.canRefreshAsync() {
		return
	}

	cts.isRefreshing = true
	go func() {
		refreshCtx, cancel := context.WithTimeout(ctx, asyncRefreshTimeout)
		defer cancel()

		t, err := cts.tokenSource.Token(refreshCtx)

		cts.mu.Lock()
		defer cts.mu.Unlock()
		cts.isRefreshing = false
		if err != nil {
			cts.nextAsyncRefresh = cts.timeNow().Add(asyncRefreshRetryBackoff)
			return
		}
		// Skip the update if a blocking refresh cached a newer token while
		// this async refresh was in flight.
		if cts.cachedToken != nil && !cts.cachedToken.Expiry.IsZero() && t.Expiry.Before(cts.cachedToken.Expiry) {
			return
		}
		cts.cachedToken = t
		cts.updateNextAsyncRefresh()
	}()
}
