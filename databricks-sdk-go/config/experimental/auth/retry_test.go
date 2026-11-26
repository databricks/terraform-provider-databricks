package auth

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/experimental/api"
	"golang.org/x/oauth2"
)

// timeoutError is a test error that implements net.Error with Timeout() == true.
type timeoutError struct{}

func (e timeoutError) Error() string   { return "i/o timeout" }
func (e timeoutError) Timeout() bool   { return true }
func (e timeoutError) Temporary() bool { return false }

// testHTTPError is a test error that implements the httpResponseError
// interface, mirroring httpclient.HttpError without importing it.
type testHTTPError struct {
	code    int
	header  http.Header
	message string
}

func (e *testHTTPError) Error() string       { return fmt.Sprintf("http %d: %s", e.code, e.message) }
func (e *testHTTPError) HTTPStatusCode() int { return e.code }
func (e *testHTTPError) Header() http.Header { return e.header }

func TestRetryingTokenSource(t *testing.T) {
	err401 := &testHTTPError{code: 401, message: "unauthorized"}
	err400 := &testHTTPError{code: 400, message: "bad request"}
	token := &oauth2.Token{AccessToken: "test-token"}

	testCases := []struct {
		name         string
		callErrors   []error
		wantToken    *oauth2.Token
		wantErr      error
		wantNumCalls int
	}{
		{
			name:         "success on first call",
			callErrors:   []error{nil},
			wantToken:    token,
			wantNumCalls: 1,
		},
		{
			name: "retry on http 429 then succeed",
			callErrors: []error{
				&testHTTPError{code: 429, message: "rate limited"},
				nil,
			},
			wantToken:    token,
			wantNumCalls: 2,
		},
		{
			name: "retry on http 503 then succeed",
			callErrors: []error{
				&testHTTPError{code: 503, message: "service unavailable"},
				nil,
			},
			wantToken:    token,
			wantNumCalls: 2,
		},
		{
			name: "retry on oauth2 retrieve error 429",
			callErrors: []error{
				&oauth2.RetrieveError{Response: &http.Response{StatusCode: 429}},
				nil,
			},
			wantToken:    token,
			wantNumCalls: 2,
		},
		{
			name: "retry on transient network error",
			callErrors: []error{
				&url.Error{Op: "Post", URL: "https://host/token", Err: &net.OpError{
					Op: "read", Net: "tcp",
					Err: &os.SyscallError{Syscall: "read", Err: syscall.ECONNRESET},
				}},
				nil,
			},
			wantToken:    token,
			wantNumCalls: 2,
		},
		{
			name: "no retry on http 401",
			callErrors: []error{
				err401,
			},
			wantErr:      err401,
			wantNumCalls: 1,
		},
		{
			name: "no retry on http 400",
			callErrors: []error{
				err400,
			},
			wantErr:      err400,
			wantNumCalls: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotNumCalls := 0
			inner := TokenSourceFn(func(ctx context.Context) (*oauth2.Token, error) {
				err := tc.callErrors[gotNumCalls]
				gotNumCalls++
				if err != nil {
					return nil, err
				}
				return token, nil
			})

			ts := NewRetryingTokenSource(inner)
			gotToken, gotErr := ts.Token(context.Background())

			if gotNumCalls != tc.wantNumCalls {
				t.Errorf("got %d calls, want %d", gotNumCalls, tc.wantNumCalls)
			}
			if !errors.Is(gotErr, tc.wantErr) {
				t.Errorf("got error %v, want %v", gotErr, tc.wantErr)
			}
			if gotToken != tc.wantToken {
				t.Errorf("got token %v, want %v", gotToken, tc.wantToken)
			}
		})
	}
}

func TestHTTPRetrier_IsRetriable(t *testing.T) {
	testCases := []struct {
		name          string
		err           error
		wantRetriable bool
	}{
		{
			name:          "http 429",
			err:           &testHTTPError{code: 429},
			wantRetriable: true,
		},
		{
			name:          "http 500",
			err:           &testHTTPError{code: 500},
			wantRetriable: false,
		},
		{
			name:          "http 502",
			err:           &testHTTPError{code: 502},
			wantRetriable: true,
		},
		{
			name:          "http 503",
			err:           &testHTTPError{code: 503},
			wantRetriable: true,
		},
		{
			name:          "http 504",
			err:           &testHTTPError{code: 504},
			wantRetriable: true,
		},
		{
			name:          "http 401",
			err:           &testHTTPError{code: 401},
			wantRetriable: false,
		},
		{
			name:          "http 403",
			err:           &testHTTPError{code: 403},
			wantRetriable: false,
		},
		{
			name:          "oauth2 retrieve error 429",
			err:           &oauth2.RetrieveError{Response: &http.Response{StatusCode: 429}},
			wantRetriable: true,
		},
		{
			name:          "oauth2 retrieve error 500",
			err:           &oauth2.RetrieveError{Response: &http.Response{StatusCode: 500}},
			wantRetriable: false,
		},
		{
			name:          "oauth2 retrieve error 400",
			err:           &oauth2.RetrieveError{Response: &http.Response{StatusCode: 400}},
			wantRetriable: false,
		},
		{
			name: "connection reset",
			err: &url.Error{Op: "Post", URL: "https://host/token", Err: &net.OpError{
				Op: "read", Net: "tcp",
				Err: &os.SyscallError{Syscall: "read", Err: syscall.ECONNRESET},
			}},
			wantRetriable: true,
		},
		{
			name: "connection refused",
			err: &url.Error{Op: "Post", URL: "https://host/token", Err: &net.OpError{
				Op: "dial", Net: "tcp",
				Err: &os.SyscallError{Syscall: "connect", Err: syscall.ECONNREFUSED},
			}},
			wantRetriable: true,
		},
		{
			name: "i/o timeout",
			err: &url.Error{Op: "Post", URL: "https://host/token", Err: &net.OpError{
				Op: "read", Net: "tcp",
				Err: timeoutError{},
			}},
			wantRetriable: true,
		},
		{
			name:          "url error non-transient",
			err:           &url.Error{Op: "Post", URL: "https://host/token", Err: fmt.Errorf("no such host")},
			wantRetriable: false,
		},
		{
			name:          "generic error",
			err:           errors.New("something went wrong"),
			wantRetriable: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := &httpRetrier{}
			_, got := r.IsRetriable(tc.err)
			if got != tc.wantRetriable {
				t.Errorf("httpRetrier.IsRetriable(%v) retriable = %v, want %v", tc.err, got, tc.wantRetriable)
			}
		})
	}
}

func TestHTTPRetrier_RetryAfterDelay(t *testing.T) {
	// Backoff returning delays between 0 and 10 seconds.
	backoff := api.BackoffPolicy{Maximum: 10 * time.Second}

	testCases := []struct {
		name    string
		err     error
		minWant time.Duration
		maxWant time.Duration
	}{
		{
			name: "retry-after exceeds backoff maximum",
			err: &testHTTPError{
				code:   429,
				header: http.Header{"Retry-After": []string{"120"}},
			},
			minWant: 120 * time.Second,
			maxWant: 120 * time.Second,
		},
		{
			name: "retry-after below backoff is ignored",
			err: &testHTTPError{
				code:   429,
				header: http.Header{"Retry-After": []string{"0"}},
			},
			minWant: 0,
			maxWant: 10 * time.Second,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := &httpRetrier{backoff: backoff}
			got, gotOk := r.IsRetriable(tc.err)
			if !gotOk {
				t.Fatalf("IsRetriable(%v) = false, want true", tc.err)
			}
			if got < tc.minWant || got > tc.maxWant {
				t.Errorf("delay = %v, want [%v, %v]", got, tc.minWant, tc.maxWant)
			}
		})
	}
}

func TestParseRetryAfter(t *testing.T) {
	testCases := []struct {
		name      string
		header    http.Header
		wantOK    bool
		wantDelay time.Duration
	}{
		{
			name:   "nil header",
			header: nil,
			wantOK: false,
		},
		{
			name:   "empty header",
			header: http.Header{},
			wantOK: false,
		},
		{
			name:      "delay in seconds",
			header:    http.Header{"Retry-After": []string{"30"}},
			wantOK:    true,
			wantDelay: 30 * time.Second,
		},
		{
			name:      "delay zero seconds",
			header:    http.Header{"Retry-After": []string{"0"}},
			wantOK:    true,
			wantDelay: 0,
		},
		{
			name:   "negative seconds",
			header: http.Header{"Retry-After": []string{"-5"}},
			wantOK: false,
		},
		{
			name:   "non-numeric string",
			header: http.Header{"Retry-After": []string{"not-a-number"}},
			wantOK: false,
		},
		{
			name:   "http-date in the past",
			header: http.Header{"Retry-After": []string{"Mon, 01 Jan 2024 00:00:00 GMT"}},
			wantOK: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotDelay, gotOK := parseRetryAfter(tc.header)
			if gotOK != tc.wantOK {
				t.Errorf("ok = %v, want %v", gotOK, tc.wantOK)
			}
			if gotDelay != tc.wantDelay {
				t.Errorf("delay = %v, want %v", gotDelay, tc.wantDelay)
			}
		})
	}
}
