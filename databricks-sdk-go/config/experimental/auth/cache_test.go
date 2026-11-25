package auth

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"golang.org/x/oauth2"
)

func waitForAsyncRefreshToComplete(t *testing.T, cts *cachedTokenSource) {
	t.Helper()

	cts.mu.Lock()
	refresh := cts.inflight
	cts.mu.Unlock()
	if refresh == nil {
		return
	}
	select {
	case <-refresh.done:
	case <-time.After(5 * time.Second):
		t.Fatal("timed out waiting for async refresh to complete")
	}

}

func TestNewCachedTokenSource_noCaching(t *testing.T) {
	want := &cachedTokenSource{}
	got := NewCachedTokenSource(want, nil)
	if got != want {
		t.Errorf("NewCachedTokenSource() = %v, want %v", got, want)
	}
}

func TestNewCachedTokenSource_default(t *testing.T) {
	ts := TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
		return nil, nil
	})

	got, ok := NewCachedTokenSource(ts).(*cachedTokenSource)
	if !ok {
		t.Fatalf("NewCachedTokenSource() = %T, want *cachedTokenSource", got)
	}

	if !got.nextAsyncRefresh.IsZero() {
		t.Errorf("NewCachedTokenSource() nextAsyncRefresh = %v, want zero value", got.nextAsyncRefresh)
	}
	if got.disableAsync != false {
		t.Errorf("NewCachedTokenSource() disableAsync = %v, want %v", got.disableAsync, false)
	}
	if got.cachedToken != nil {
		t.Errorf("NewCachedTokenSource() cachedToken = %v, want nil", got.cachedToken)
	}
}

func TestNewCachedTokenSource_options(t *testing.T) {
	ts := TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
		return nil, nil
	})

	wantDisableAsync := false
	wantCachedToken := &oauth2.Token{Expiry: time.Unix(42, 0)}

	opts := []Option{
		WithAsyncRefresh(!wantDisableAsync),
		WithCachedToken(wantCachedToken),
	}

	got, ok := NewCachedTokenSource(ts, opts...).(*cachedTokenSource)
	if !ok {
		t.Fatalf("NewCachedTokenSource() = %T, want *cachedTokenSource", got)
	}

	if got.disableAsync != wantDisableAsync {
		t.Errorf("NewCachedTokenSource(): disableAsync = %v, want %v", got.disableAsync, wantDisableAsync)
	}
	if got.cachedToken != wantCachedToken {
		t.Errorf("NewCachedTokenSource(): cachedToken = %v, want %v", got.cachedToken, wantCachedToken)
	}
}

func TestCachedTokenSource_updateNextAsyncRefresh(t *testing.T) {
	now := time.Unix(1337, 0)

	testCases := []struct {
		name             string
		tokenTTL         time.Duration
		wantAllowedAfter time.Duration
	}{
		{
			name:             "standard OAuth token with 60-minute TTL",
			tokenTTL:         60 * time.Minute,
			wantAllowedAfter: 40 * time.Minute,
		},
		{
			name:             "short-lived token with 10-minute TTL",
			tokenTTL:         10 * time.Minute,
			wantAllowedAfter: 5 * time.Minute,
		},
		{
			name:             "very short token with 90-second TTL",
			tokenTTL:         90 * time.Second,
			wantAllowedAfter: 45 * time.Second,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cts := &cachedTokenSource{
				cachedToken: &oauth2.Token{
					Expiry: now.Add(tc.tokenTTL),
				},
				timeNow: func() time.Time { return now },
			}

			cts.updateNextAsyncRefresh()

			want := now.Add(tc.wantAllowedAfter)
			if cts.nextAsyncRefresh != want {
				t.Errorf("nextAsyncRefresh = %v, want %v", cts.nextAsyncRefresh, want)
			}
		})
	}
}

func TestCachedTokenSource_tokenExpired(t *testing.T) {
	now := time.Unix(1337, 0) // mock value for time.Now()

	testCases := []struct {
		name  string
		token *oauth2.Token
		want  bool
	}{
		{
			name:  "nil token",
			token: nil,
			want:  true,
		},
		{
			name: "expired token",
			token: &oauth2.Token{
				Expiry: now.Add(-1 * time.Second),
			},
			want: true,
		},
		{
			name: "token expiring now",
			token: &oauth2.Token{
				Expiry: now,
			},
			want: false,
		},
		{
			name: "future token",
			token: &oauth2.Token{
				Expiry: now.Add(1 * time.Hour),
			},
			want: false,
		},
		{
			name: "token without expiry",
			token: &oauth2.Token{
				Expiry: time.Time{},
			},
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cts := &cachedTokenSource{
				cachedToken: tc.token,
				timeNow:     func() time.Time { return now },
			}

			got := cts.tokenExpired()

			if got != tc.want {
				t.Errorf("tokenExpired() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCachedTokenSource_canRefreshAsync(t *testing.T) {
	now := time.Unix(1337, 0) // mock value for time.Now()

	testCases := []struct {
		name             string
		token            *oauth2.Token
		nextAsyncRefresh time.Time
		want             bool
	}{
		{
			name: "nil token",
			want: false,
		},
		{
			name: "token without expiry",
			token: &oauth2.Token{
				Expiry: time.Time{},
			},
			want: false,
		},
		{
			name: "before async refresh window",
			token: &oauth2.Token{
				Expiry: now.Add(1 * time.Hour),
			},
			nextAsyncRefresh: now.Add(1 * time.Minute),
			want:             false,
		},
		{
			name: "exactly at async refresh boundary",
			token: &oauth2.Token{
				Expiry: now.Add(1 * time.Hour),
			},
			nextAsyncRefresh: now,
			want:             false,
		},
		{
			name: "after async refresh boundary",
			token: &oauth2.Token{
				Expiry: now.Add(1 * time.Hour),
			},
			nextAsyncRefresh: now.Add(-1 * time.Second),
			want:             true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cts := &cachedTokenSource{
				cachedToken:      tc.token,
				nextAsyncRefresh: tc.nextAsyncRefresh,
				timeNow:          func() time.Time { return now },
			}

			got := cts.canRefreshAsync()

			if got != tc.want {
				t.Errorf("canRefreshAsync() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCachedTokenSource_AsyncRefreshRetry(t *testing.T) {
	failTime := time.Unix(1337, 0)
	currentToken := &oauth2.Token{
		AccessToken: "current-token",
		Expiry:      failTime.Add(5 * time.Minute),
	}
	refreshedToken := &oauth2.Token{
		AccessToken: "refreshed-token",
		Expiry:      failTime.Add(1 * time.Hour),
	}

	testCases := []struct {
		name      string
		now       time.Time
		wantCalls int32
		wantCache *oauth2.Token
	}{
		{
			name:      "no async refresh allowed during backoff",
			now:       failTime.Add(asyncRefreshRetryBackoff - 1*time.Second),
			wantCalls: 0,
			wantCache: currentToken,
		},
		{
			name:      "async refresh allowed after backoff",
			now:       failTime.Add(asyncRefreshRetryBackoff + 1*time.Second),
			wantCalls: 1,
			wantCache: refreshedToken,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotCalls := int32(0)
			cts := &cachedTokenSource{
				cachedToken:      currentToken,
				nextAsyncRefresh: failTime.Add(asyncRefreshRetryBackoff),
				timeNow:          func() time.Time { return tc.now },
				tokenSource: TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
					atomic.AddInt32(&gotCalls, 1)
					return refreshedToken, nil
				}),
			}

			gotToken, err := cts.Token(context.Background())
			if err != nil {
				t.Fatalf("Token() error = %v", err)
			}
			if gotToken != currentToken {
				t.Errorf("Token() = %v, want %v", gotToken, currentToken)
			}

			waitForAsyncRefreshToComplete(t, cts)

			if got := atomic.LoadInt32(&gotCalls); got != tc.wantCalls {
				t.Fatalf("token source calls = %d, want %d", got, tc.wantCalls)
			}
			if cts.cachedToken != tc.wantCache {
				t.Errorf("cachedToken = %v, want %v", cts.cachedToken, tc.wantCache)
			}
		})
	}
}

func TestCachedTokenSource_Token(t *testing.T) {
	now := time.Unix(1337, 0) // mock value for time.Now()
	nTokenCalls := 10         // number of goroutines calling Token()
	testCases := []struct {
		desc             string        // description of the test case
		cachedToken      *oauth2.Token // token cached before calling Token()
		disableAsync     bool          // whether async refreshes are disabled
		nextAsyncRefresh time.Time     // time after which async refreshes may be attempted

		returnedToken *oauth2.Token // token returned by the token source
		returnedError error         // error returned by the token source

		wantCalls int           // expected number of calls to the token source
		wantToken *oauth2.Token // expected token in the cache
	}{
		{
			desc:          "[Blocking] no cached token",
			disableAsync:  true,
			returnedToken: &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
			wantCalls:     1,
			wantToken:     &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
		},
		{
			desc:          "[Blocking] expired cached token",
			disableAsync:  true,
			cachedToken:   &oauth2.Token{Expiry: now.Add(-1 * time.Second)},
			returnedToken: &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
			wantCalls:     1,
			wantToken:     &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
		},
		{
			desc:         "[Blocking] fresh cached token",
			disableAsync: true,
			cachedToken:  &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
			wantCalls:    0,
			wantToken:    &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
		},
		{
			desc:         "[Blocking] stale cached token",
			disableAsync: true,
			cachedToken:  &oauth2.Token{Expiry: now.Add(1 * time.Minute)},
			wantCalls:    0,
			wantToken:    &oauth2.Token{Expiry: now.Add(1 * time.Minute)},
		},
		{
			desc:          "[Blocking] refresh error",
			disableAsync:  true,
			returnedError: fmt.Errorf("test error"),
			wantCalls:     1,
		},
		{
			desc:          "[Blocking] recover from error",
			disableAsync:  true,
			cachedToken:   &oauth2.Token{Expiry: now.Add(-1 * time.Minute)},
			returnedToken: &oauth2.Token{Expiry: now.Add(-1 * time.Hour)},
			wantCalls:     1,
			wantToken:     &oauth2.Token{Expiry: now.Add(-1 * time.Hour)},
		},
		{
			desc:          "[Async] no cached token",
			returnedToken: &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
			wantCalls:     1,
			wantToken:     &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
		},
		{
			desc:          "[Async] expired cached token",
			cachedToken:   &oauth2.Token{Expiry: now.Add(-1 * time.Second)},
			returnedToken: &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
			wantCalls:     1,
			wantToken:     &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
		},
		{
			desc:             "[Async] fresh cached token",
			cachedToken:      &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
			nextAsyncRefresh: now.Add(1 * time.Minute),
			wantCalls:        0,
			wantToken:        &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
		},
		{
			desc:             "[Async] stale cached token",
			cachedToken:      &oauth2.Token{Expiry: now.Add(1 * time.Minute)},
			nextAsyncRefresh: now.Add(-1 * time.Second),
			returnedToken:    &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
			wantCalls:        1,
			wantToken:        &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
		},
		{
			desc:             "[Async] refresh error",
			cachedToken:      &oauth2.Token{Expiry: now.Add(1 * time.Minute)},
			nextAsyncRefresh: now.Add(-1 * time.Second),
			returnedError:    fmt.Errorf("test error"),
			wantCalls:        1,
			wantToken:        &oauth2.Token{Expiry: now.Add(1 * time.Minute)},
		},
		{
			desc:             "[Async] stale cached token, expired token returned",
			cachedToken:      &oauth2.Token{Expiry: now.Add(1 * time.Minute)},
			nextAsyncRefresh: now.Add(-1 * time.Second),
			returnedToken:    &oauth2.Token{Expiry: now.Add(-1 * time.Second)},
			wantCalls:        1,
			wantToken:        &oauth2.Token{Expiry: now.Add(1 * time.Minute)},
		},
		{
			desc:          "[Async] recover from error",
			cachedToken:   &oauth2.Token{Expiry: now.Add(-1 * time.Minute)},
			returnedToken: &oauth2.Token{Expiry: now.Add(-1 * time.Hour)},
			wantCalls:     1,
			wantToken:     &oauth2.Token{Expiry: now.Add(-1 * time.Hour)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				gotCalls := int32(0)
				release := make(chan struct{})
				cts := &cachedTokenSource{
					disableAsync:     tc.disableAsync,
					cachedToken:      tc.cachedToken,
					nextAsyncRefresh: tc.nextAsyncRefresh,
					timeNow:          func() time.Time { return now },
					tokenSource: TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
						atomic.AddInt32(&gotCalls, 1)
						<-release
						return tc.returnedToken, tc.returnedError
					}),
				}

				wantResult, wantErr := tc.returnedToken, tc.returnedError
				if tc.cachedToken != nil && !now.After(tc.cachedToken.Expiry) {
					wantResult, wantErr = tc.cachedToken, nil
				}
				wg := sync.WaitGroup{}
				for i := 0; i < nTokenCalls; i++ {
					wg.Add(1)
					go func() {
						defer wg.Done()
						got, err := cts.Token(context.Background())
						if got != wantResult || !errors.Is(err, wantErr) {
							t.Errorf("Token() = %v, %v, want %v, %v", got, err, wantResult, wantErr)
						}
					}()
				}

				// Keep the refresh in flight until the callers have joined it, so
				// that they share one refresh instead of each starting their own.
				synctest.Wait()
				close(release)
				wg.Wait()

				waitForAsyncRefreshToComplete(t, cts)

				if got := int(atomic.LoadInt32(&gotCalls)); got != tc.wantCalls {
					t.Errorf("want %d calls to cts.tokenSource.Token(), got %d", tc.wantCalls, got)
				}
				if !reflect.DeepEqual(tc.wantToken, cts.cachedToken) {
					t.Errorf("want cached token %v, got %v", tc.wantToken, cts.cachedToken)
				}
			})
		})
	}
}

func TestCachedTokenSource_BlockingRefreshUpdatesNextAsyncRefresh(t *testing.T) {
	now := time.Unix(1337, 0)
	refreshedToken := &oauth2.Token{Expiry: now.Add(1 * time.Hour)}

	cts := &cachedTokenSource{
		disableAsync:     true,
		nextAsyncRefresh: now.Add(-2 * time.Minute),
		timeNow:          func() time.Time { return now },
		cachedToken:      &oauth2.Token{Expiry: now.Add(-1 * time.Second)}, // expired
		tokenSource: TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
			return refreshedToken, nil
		}),
	}

	_, err := cts.Token(context.Background())
	if err != nil {
		t.Fatalf("Token() error = %v", err)
	}

	want := refreshedToken.Expiry.Add(-maxAsyncRefreshLeadTime)
	if cts.nextAsyncRefresh != want {
		t.Errorf("nextAsyncRefresh = %v, want %v", cts.nextAsyncRefresh, want)
	}
}

func TestCachedTokenSource_AsyncRefreshSkipsOlderToken(t *testing.T) {
	now := time.Unix(1337, 0)
	// Simulate a blocking refresh that cached a fresh token while an async
	// refresh was in flight. The async goroutine should not overwrite the
	// fresher cached token with its older result.
	cachedToken := &oauth2.Token{
		AccessToken: "fresh-from-blocking",
		Expiry:      now.Add(1 * time.Hour),
	}
	olderToken := &oauth2.Token{
		AccessToken: "stale-from-async",
		Expiry:      now.Add(30 * time.Minute),
	}

	cts := &cachedTokenSource{
		cachedToken:      &oauth2.Token{AccessToken: "original", Expiry: now.Add(2 * time.Minute)},
		nextAsyncRefresh: now.Add(-1 * time.Second),
		timeNow:          func() time.Time { return now },
	}
	cts.tokenSource = TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
		// Simulate blocking refresh completing first by swapping in the
		// fresh token before this goroutine returns.
		cts.mu.Lock()
		cts.cachedToken = cachedToken
		cts.mu.Unlock()
		return olderToken, nil
	})

	_, err := cts.Token(context.Background())
	if err != nil {
		t.Fatalf("Token() error = %v", err)
	}

	waitForAsyncRefreshToComplete(t, cts)

	if cts.cachedToken != cachedToken {
		t.Errorf("cachedToken = %v, want %v (fresher token from blocking refresh)", cts.cachedToken, cachedToken)
	}
}

func TestCachedTokenSource_BlockingCallerReturnsWhenContextEnds(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		now := time.Unix(1337, 0)
		refreshedToken := &oauth2.Token{AccessToken: "refreshed-token", Expiry: now.Add(1 * time.Hour)}

		release := make(chan struct{})
		releaseRefresh := sync.OnceFunc(func() { close(release) })
		defer releaseRefresh()
		refreshCtxErr := make(chan error, 1)
		gotCalls := int32(0)
		cts := &cachedTokenSource{
			disableAsync: true,
			timeNow:      func() time.Time { return now },
			tokenSource: TokenSourceFn(func(ctx context.Context) (*oauth2.Token, error) {
				atomic.AddInt32(&gotCalls, 1)
				<-release
				if _, ok := ctx.Deadline(); !ok {
					t.Error("refresh context has no deadline, want it bounded")
				}
				refreshCtxErr <- ctx.Err()
				return refreshedToken, nil
			}),
		}

		// The refresh is held in flight, so the caller's deadline is the only
		// thing that can end the call.
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancel()
		errc := make(chan error, 1)
		go func() {
			_, err := cts.Token(ctx)
			errc <- err
		}()

		select {
		case err := <-errc:
			if !errors.Is(err, context.DeadlineExceeded) {
				t.Fatalf("Token() error = %v, want context.DeadlineExceeded", err)
			}
		case <-time.After(5 * time.Second):
			t.Fatal("Token() did not return after the caller's context deadline")
		}

		// The refresh must not have been canceled with the caller. Its token is
		// served to the next caller, which waits for the same refresh.
		releaseRefresh()
		select {
		case err := <-refreshCtxErr:
			if err != nil {
				t.Errorf("refresh context error = %v, want nil", err)
			}
		case <-time.After(5 * time.Second):
			t.Fatal("token refresh did not complete")
		}

		got, err := cts.Token(context.Background())
		if err != nil {
			t.Fatalf("Token() error = %v", err)
		}
		if got != refreshedToken {
			t.Errorf("Token() = %v, want %v", got, refreshedToken)
		}
		if n := atomic.LoadInt32(&gotCalls); n != 1 {
			t.Errorf("token source calls = %d, want 1", n)
		}
	})
}

func TestCachedTokenSource_ConcurrentCallersShareOneRefresh(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		now := time.Unix(1337, 0)
		refreshedToken := &oauth2.Token{AccessToken: "refreshed-token", Expiry: now.Add(1 * time.Hour)}
		const nCallers = 10

		release := make(chan struct{})
		gotCalls := int32(0)
		cts := &cachedTokenSource{
			disableAsync: true,
			timeNow:      func() time.Time { return now },
			tokenSource: TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
				atomic.AddInt32(&gotCalls, 1)
				<-release
				return refreshedToken, nil
			}),
		}

		tokens := make([]*oauth2.Token, nCallers)
		errs := make([]error, nCallers)
		wg := sync.WaitGroup{}
		for i := range nCallers {
			wg.Add(1)
			go func() {
				defer wg.Done()
				tokens[i], errs[i] = cts.Token(context.Background())
			}()
		}

		// Keep the refresh in flight until the callers have joined it.
		synctest.Wait()
		close(release)
		wg.Wait()

		for i := range nCallers {
			if errs[i] != nil {
				t.Errorf("caller %d: Token() error = %v", i, errs[i])
			}
			if tokens[i] != refreshedToken {
				t.Errorf("caller %d: Token() = %v, want %v", i, tokens[i], refreshedToken)
			}
		}
		if n := atomic.LoadInt32(&gotCalls); n != 1 {
			t.Errorf("token source calls = %d, want 1", n)
		}
	})
}

func TestCachedTokenSource_ExpiredCallerJoinsInFlightAsyncRefresh(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		start := time.Unix(1337, 0)
		var clock atomic.Int64 // Unix nanoseconds, moved forward by the test.
		clock.Store(start.UnixNano())
		timeNow := func() time.Time { return time.Unix(0, clock.Load()) }

		currentToken := &oauth2.Token{AccessToken: "current-token", Expiry: start.Add(4 * time.Minute)}
		refreshedToken := &oauth2.Token{AccessToken: "refreshed-token", Expiry: start.Add(1 * time.Hour)}

		started := make(chan struct{})
		release := make(chan struct{})
		releaseRefresh := sync.OnceFunc(func() { close(release) })
		defer releaseRefresh()
		gotCalls := int32(0)
		cts := &cachedTokenSource{
			cachedToken: currentToken,
			// The token is inside its async refresh window at start.
			nextAsyncRefresh: start.Add(-1 * time.Minute),
			timeNow:          timeNow,
			tokenSource: TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
				if atomic.AddInt32(&gotCalls, 1) == 1 {
					close(started)
				}
				<-release
				return refreshedToken, nil
			}),
		}

		// The valid token is served while the async refresh starts.
		got, err := cts.Token(context.Background())
		if err != nil {
			t.Fatalf("Token() error = %v", err)
		}
		if got != currentToken {
			t.Errorf("Token() = %v, want %v", got, currentToken)
		}
		select {
		case <-started:
		case <-time.After(5 * time.Second):
			t.Fatal("async refresh was not started")
		}

		// The token expires while the refresh is in flight. A caller that now
		// needs a token must wait for that refresh, not start another one.
		clock.Store(currentToken.Expiry.Add(time.Second).UnixNano())
		type result struct {
			token *oauth2.Token
			err   error
		}
		resc := make(chan result, 1)
		go func() {
			tok, err := cts.Token(context.Background())
			resc <- result{tok, err}
		}()

		synctest.Wait()
		releaseRefresh()

		select {
		case res := <-resc:
			if res.err != nil {
				t.Fatalf("Token() error = %v", res.err)
			}
			if res.token != refreshedToken {
				t.Errorf("Token() = %v, want %v", res.token, refreshedToken)
			}
		case <-time.After(5 * time.Second):
			t.Fatal("Token() did not return after the refresh completed")
		}
		if n := atomic.LoadInt32(&gotCalls); n != 1 {
			t.Errorf("token source calls = %d, want 1", n)
		}
	})
}

func TestCachedTokenSource_ExpiredCallerRetriesUnusableProactiveRefresh(t *testing.T) {
	start := time.Unix(1337, 0)
	current := &oauth2.Token{AccessToken: "current", Expiry: start.Add(time.Minute)}
	fresh := &oauth2.Token{AccessToken: "fresh", Expiry: start.Add(time.Hour)}
	for _, tc := range []struct {
		name  string
		token *oauth2.Token
		err   error
	}{
		{name: "refresh failed", err: errors.New("temporary error")},
		{name: "older token discarded", token: &oauth2.Token{AccessToken: "older", Expiry: start}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				var clock atomic.Int64
				clock.Store(start.UnixNano())
				release := make(chan struct{})
				releaseRefresh := sync.OnceFunc(func() { close(release) })
				defer releaseRefresh()
				var calls atomic.Int32
				cts := &cachedTokenSource{
					cachedToken:      current,
					nextAsyncRefresh: start.Add(-time.Second),
					timeNow:          func() time.Time { return time.Unix(0, clock.Load()) },
					tokenSource: TokenSourceFn(func(context.Context) (*oauth2.Token, error) {
						if calls.Add(1) == 1 {
							<-release
							return tc.token, tc.err
						}
						return fresh, nil
					}),
				}
				if got, err := cts.Token(t.Context()); err != nil || got != current {
					t.Fatalf("Token() = %v, %v, want the current token", got, err)
				}
				clock.Store(current.Expiry.Add(time.Second).UnixNano())
				done := make(chan struct{})
				go func() {
					defer close(done)
					if got, err := cts.Token(t.Context()); err != nil || got != fresh {
						t.Errorf("Token() = %v, %v, want a fresh token", got, err)
					}
				}()
				synctest.Wait()
				if n := calls.Load(); n != 1 {
					t.Errorf("calls while proactive refresh is in flight = %d, want 1", n)
				}
				releaseRefresh()
				select {
				case <-done:
				case <-time.After(5 * time.Second):
					t.Fatal("caller did not finish after the proactive attempt")
				}
				if n := calls.Load(); n != 2 {
					t.Errorf("token source calls = %d, want 2", n)
				}
			})
		})
	}
}

func TestCachedTokenSource_RefreshDeadlineReleasesWaiters(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		start := time.Now()
		current := &oauth2.Token{AccessToken: "current", Expiry: start.Add(time.Second)}
		fresh := &oauth2.Token{AccessToken: "fresh", Expiry: start.Add(time.Hour)}
		late := &oauth2.Token{AccessToken: "late", Expiry: start.Add(2 * time.Hour)}
		release := make(chan struct{})
		releaseRefresh := sync.OnceFunc(func() { close(release) })
		defer releaseRefresh()
		var calls atomic.Int32
		source := NewCachedTokenSource(TokenSourceFn(func(context.Context) (*oauth2.Token, error) {
			if calls.Add(1) == 1 {
				<-release
				return late, nil
			}
			return fresh, nil
		}), WithCachedToken(current))

		time.Sleep(750 * time.Millisecond)
		if got, err := source.Token(t.Context()); err != nil || got != current {
			t.Fatalf("Token() = %v, %v, want the current token", got, err)
		}
		time.Sleep(time.Second)
		ctx, cancel := context.WithTimeout(t.Context(), refreshTimeout+time.Second)
		defer cancel()
		if got, err := source.Token(ctx); err != nil || got != fresh {
			t.Fatalf("Token() = %v, %v, want the fresh token", got, err)
		}

		releaseRefresh()
		synctest.Wait()
		if got, err := source.Token(t.Context()); err != nil || got != fresh {
			t.Errorf("Token() after the overdue refresh returned = %v, %v, want the fresh token", got, err)
		}
	})
}

func TestCachedTokenSource_LimitsOverdueRefreshes(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		release := make(chan struct{})
		releaseRefresh := sync.OnceFunc(func() { close(release) })
		defer releaseRefresh()
		var calls atomic.Int32
		fresh := &oauth2.Token{AccessToken: "fresh"}
		source := NewCachedTokenSource(TokenSourceFn(func(context.Context) (*oauth2.Token, error) {
			calls.Add(1)
			<-release
			return fresh, nil
		}))

		for range maxRunningRefreshes + 1 {
			ctx, cancel := context.WithTimeout(t.Context(), refreshTimeout+time.Second)
			_, err := source.Token(ctx)
			cancel()
			if !errors.Is(err, context.DeadlineExceeded) {
				t.Fatalf("Token() error = %v, want DeadlineExceeded", err)
			}
		}
		if n := calls.Load(); n != maxRunningRefreshes {
			t.Fatalf("token source calls = %d, want %d", n, maxRunningRefreshes)
		}

		releaseRefresh()
		synctest.Wait()
		if got, err := source.Token(t.Context()); err != nil || got != fresh {
			t.Errorf("Token() after overdue refreshes returned = %v, %v, want a fresh token", got, err)
		}
	})
}

func TestCachedTokenSource_CanceledCallerDoesNotStartRefresh(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var calls atomic.Int32
		source := NewCachedTokenSource(TokenSourceFn(func(context.Context) (*oauth2.Token, error) {
			calls.Add(1)
			return &oauth2.Token{AccessToken: "fresh"}, nil
		}))
		ctx, cancel := context.WithCancel(t.Context())
		cancel()
		if _, err := source.Token(ctx); !errors.Is(err, context.Canceled) {
			t.Errorf("Token() error = %v, want context.Canceled", err)
		}
		synctest.Wait()
		if n := calls.Load(); n != 0 {
			t.Errorf("token source calls = %d, want 0", n)
		}
	})
}

func TestCachedTokenSource_AsyncRefreshOutlivesRequestContext(t *testing.T) {
	now := time.Unix(1337, 0)
	currentToken := &oauth2.Token{
		AccessToken: "current-token",
		Expiry:      now.Add(4 * time.Minute),
	}
	refreshedToken := &oauth2.Token{
		AccessToken: "refreshed-token",
		Expiry:      now.Add(1 * time.Hour),
	}

	started := make(chan struct{})
	release := make(chan struct{})
	releaseRefresh := sync.OnceFunc(func() { close(release) })
	defer releaseRefresh()
	cts := &cachedTokenSource{
		cachedToken: currentToken,
		// The token is inside its async refresh window at now+3m.
		nextAsyncRefresh: now.Add(2 * time.Minute),
		timeNow:          func() time.Time { return now.Add(3 * time.Minute) },
		tokenSource: TokenSourceFn(func(ctx context.Context) (*oauth2.Token, error) {
			close(started)
			<-release
			if err := ctx.Err(); err != nil {
				return nil, err
			}
			return refreshedToken, nil
		}),
	}

	reqCtx, cancelReq := context.WithCancel(context.Background())
	got, err := cts.Token(reqCtx)
	if err != nil {
		t.Fatalf("Token() error = %v", err)
	}
	if got != currentToken {
		t.Errorf("Token() = %v, want the current token while refreshing in the background", got)
	}

	select {
	case <-started:
	case <-time.After(5 * time.Second):
		t.Fatal("async refresh was not started")
	}

	// The request ends while the background refresh is still in flight.
	cancelReq()
	releaseRefresh()

	waitForAsyncRefreshToComplete(t, cts)

	if cts.cachedToken != refreshedToken {
		t.Errorf("cachedToken = %v, want %v (async refresh should not be canceled with the request)", cts.cachedToken, refreshedToken)
	}
}
