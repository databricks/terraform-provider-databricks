package auth

import (
	"context"
	"fmt"
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

// Maximum time a token refresh is allowed to run before its context is
// canceled. It applies to blocking and async refreshes alike, since neither is
// tied to the context of the caller that started it.
const refreshTimeout = 1 * time.Minute

// Allow a replacement for one timed-out source call. Some providers ignore
// cancellation and can keep running after the cache releases their waiters.
const maxRunningRefreshes = 2

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
// The token cache is safe for concurrent use by multiple goroutines and
// shares one refresh attempt among callers.
//
// Each caller waits only while its context is live. A refresh already in
// progress can continue and populate the cache using its own one-minute deadline.
// After the deadline, the cache can start a new attempt even if the token
// source is still handling the canceled attempt.
//
// The token cache does not take care of retries in case the token source
// returns an error; it is the responsibility of the provided token source to
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

	// The active token refresh, if any. Callers waiting for a token share it.
	inflight *tokenRefresh

	// Includes calls whose waiters were released at the refresh deadline.
	runningRefreshes int

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

// tokenRefresh records a refresh that callers can wait on. Its result is set
// before done closes.
type tokenRefresh struct {
	done      chan struct{}
	proactive bool
	token     *oauth2.Token
	err       error
}

func (r *tokenRefresh) wait(ctx context.Context) error {
	select {
	case <-r.done:
		return nil
	case <-ctx.Done():
		return fmt.Errorf("waiting for token refresh: %w", ctx.Err())
	}
}

func (cts *cachedTokenSource) blockingToken(ctx context.Context) (*oauth2.Token, error) {
	for {
		cts.mu.Lock()
		if !cts.tokenExpired() {
			t := cts.cachedToken
			cts.mu.Unlock()
			return t, nil
		}
		if err := ctx.Err(); err != nil {
			cts.mu.Unlock()
			return nil, fmt.Errorf("waiting for token refresh: %w", err)
		}
		r := cts.inflight
		if r == nil {
			if cts.runningRefreshes >= maxRunningRefreshes {
				cts.mu.Unlock()
				return nil, fmt.Errorf("previous token refreshes are still running: %w", context.DeadlineExceeded)
			}
			r = cts.startRefresh(ctx, false)
		}
		cts.mu.Unlock()

		if err := r.wait(ctx); err != nil {
			return nil, err
		}
		if !r.proactive {
			return r.token, r.err
		}
		// A proactive attempt can fail or yield a token that the cache
		// discards. Recheck the cache before deciding to refresh again.
	}
}

// startRefresh starts a token refresh in the background and records it as the
// refresh in flight. The caller must hold cts.mu.
//
// The refresh runs on a context detached from ctx, so it is not canceled when
// the caller that started it gives up. Its result is cached for the next caller.
func (cts *cachedTokenSource) startRefresh(ctx context.Context, proactive bool) *tokenRefresh {
	r := &tokenRefresh{
		done:      make(chan struct{}),
		proactive: proactive,
	}
	cts.inflight = r
	cts.runningRefreshes++
	refreshCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), refreshTimeout)
	finish := func(t *oauth2.Token, err error, returned bool) {
		cts.mu.Lock()
		defer cts.mu.Unlock()
		if returned {
			cts.runningRefreshes--
		}
		// A result that arrives after the deadline must not replace a later token.
		if cts.inflight != r {
			return
		}
		defer close(r.done)
		cts.inflight = nil
		if ctxErr := refreshCtx.Err(); ctxErr != nil {
			t, err = nil, ctxErr
		}
		r.token, r.err = t, err
		if err != nil {
			if proactive {
				cts.nextAsyncRefresh = cts.timeNow().Add(asyncRefreshRetryBackoff)
			}
			return
		}
		// A proactive attempt must not replace the cached token with an older one.
		if proactive && cts.cachedToken != nil && !cts.cachedToken.Expiry.IsZero() && t.Expiry.Before(cts.cachedToken.Expiry) {
			return
		}
		cts.cachedToken = t
		cts.updateNextAsyncRefresh()
	}
	// Some token sources cannot honor cancellation. Release the shared
	// attempt at the deadline even if the source hasn't returned.
	stop := context.AfterFunc(refreshCtx, func() {
		finish(nil, refreshCtx.Err(), false)
	})
	go func() {
		defer cancel()
		defer stop()
		t, err := cts.tokenSource.Token(refreshCtx)
		finish(t, err, true)
	}()
	return r
}

func (cts *cachedTokenSource) triggerAsyncRefresh(ctx context.Context) {
	cts.mu.Lock()
	defer cts.mu.Unlock()
	if ctx.Err() != nil || cts.inflight != nil || cts.runningRefreshes >= maxRunningRefreshes ||
		cts.tokenExpired() || !cts.canRefreshAsync() {
		return
	}
	cts.startRefresh(ctx, true)
}
