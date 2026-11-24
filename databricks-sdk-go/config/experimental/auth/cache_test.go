package auth

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

func waitForAsyncRefreshToComplete(t *testing.T, cts *cachedTokenSource) {
	t.Helper()

	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		cts.mu.Lock()
		isRefreshing := cts.isRefreshing
		cts.mu.Unlock()

		if !isRefreshing {
			return
		}

		time.Sleep(1 * time.Millisecond)
	}

	t.Fatal("timed out waiting for async refresh to complete")
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
			wantCalls:     10,
		},
		{
			desc:          "[Blocking] recover from error",
			disableAsync:  true,
			cachedToken:   &oauth2.Token{Expiry: now.Add(-1 * time.Minute)},
			returnedToken: &oauth2.Token{Expiry: now.Add(-1 * time.Hour)},
			wantCalls:     10,
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
			wantCalls:     10,
			wantToken:     &oauth2.Token{Expiry: now.Add(-1 * time.Hour)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			gotCalls := int32(0)
			cts := &cachedTokenSource{
				disableAsync:     tc.disableAsync,
				cachedToken:      tc.cachedToken,
				nextAsyncRefresh: tc.nextAsyncRefresh,
				timeNow:          func() time.Time { return now },
				tokenSource: TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
					atomic.AddInt32(&gotCalls, 1)
					time.Sleep(10 * time.Millisecond)
					return tc.returnedToken, tc.returnedError
				}),
			}

			wg := sync.WaitGroup{}
			for i := 0; i < nTokenCalls; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					cts.Token(context.Background())
				}()
			}

			wg.Wait()

			waitForAsyncRefreshToComplete(t, cts)

			if got := int(atomic.LoadInt32(&gotCalls)); got != tc.wantCalls {
				t.Errorf("want %d calls to cts.tokenSource.Token(), got %d", tc.wantCalls, got)
			}
			if !reflect.DeepEqual(tc.wantToken, cts.cachedToken) {
				t.Errorf("want cached token %v, got %v", tc.wantToken, cts.cachedToken)
			}
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
