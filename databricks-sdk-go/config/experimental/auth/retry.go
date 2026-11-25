package auth

import (
	"context"
	"errors"
	"net"
	"net/http"
	"slices"
	"strconv"
	"syscall"
	"time"

	"github.com/databricks/databricks-sdk-go/experimental/api"
	"golang.org/x/oauth2"
)

// retryingTokenSource wraps a TokenSource with retry logic for transient
// failures during token acquisition. Each retry calls the underlying Token()
// method, which creates a fresh HTTP request.
type retryingTokenSource struct {
	inner TokenSource
	opts  []api.Option
}

// NewRetryingTokenSource wraps inner with retry logic for transient failures.
// The provided options are applied after the default options, allowing callers
// to override the default timeout and retry behavior.
func NewRetryingTokenSource(inner TokenSource, opts ...api.Option) TokenSource {
	defaultOptions := []api.Option{
		api.WithTimeout(1 * time.Minute),
		api.WithRetrier(func() api.Retrier { return &httpRetrier{} }),
	}
	return &retryingTokenSource{
		inner: inner,
		opts:  append(defaultOptions, opts...),
	}
}

// Token returns a token from the underlying source, retrying on transient
// errors.
func (r *retryingTokenSource) Token(ctx context.Context) (*oauth2.Token, error) {
	return api.ExecuteWithResult(ctx, r.inner.Token, r.opts...)
}

// httpRetrier classifies errors from HTTP-based token endpoints and determines
// whether they should be retried and with what delay.
//
// An error is retriable if it carries an HTTP status code in retriableCodes
// (429, 502, 503, 504) or is a transient network error (ECONNRESET,
// ECONNREFUSED, timeout). All other errors, including 4xx client errors, are
// not retriable regardless of headers.
//
// For retriable HTTP errors, the Retry-After header is used as a hint: the
// delay is the maximum of the exponential backoff and the Retry-After value,
// so we never retry sooner than the server asked. No upper cap is applied,
// the context timeout already bounds total retry duration.
//
// TODO: this retrier encodes logic (retriable status codes, Retry-After
// parsing, network error classification) that is also needed by the main HTTP
// retry loop in httpclient/. Consider moving it to a shared location so both
// token acquisition and regular API calls use the same retry classification.
type httpRetrier struct {
	backoff api.BackoffPolicy
}

// IsRetriable reports whether err is a transient failure that should be
// retried, and if so, how long the caller should wait before retrying.
func (r *httpRetrier) IsRetriable(err error) (time.Duration, bool) {
	if code, header := httpResponse(err); code != 0 {
		if !slices.Contains(retriableCodes, code) {
			return 0, false
		}
		delay := r.backoff.Delay()
		if d, ok := parseRetryAfter(header); ok && d > delay {
			delay = d // context timeout already bounds total retry duration
		}
		return delay, true
	}
	if isTransientNetworkError(err) {
		return r.backoff.Delay(), true
	}
	return 0, false
}

// parseRetryAfter parses a Retry-After header value as either a delay in
// seconds or an HTTP-date.
func parseRetryAfter(header http.Header) (time.Duration, bool) {
	if header == nil {
		return 0, false
	}
	v := header.Get("Retry-After")
	if v == "" {
		return 0, false
	}
	if seconds, parseErr := strconv.Atoi(v); parseErr == nil && seconds >= 0 {
		return time.Duration(seconds) * time.Second, true
	}
	if t, parseErr := http.ParseTime(v); parseErr == nil {
		if d := time.Until(t); d > 0 {
			return d, true
		}
	}
	return 0, false
}

// httpResponseError is implemented by errors that carry HTTP response
// metadata. This interface avoids importing httpclient (which would create a
// cycle) while still allowing classification of httpclient.HttpError by status
// code and headers.
//
// TODO: This is meant to be temporary and should move this to a shared location
// so both token acquisition and regular API calls use the same classification.
type httpResponseError interface {
	HTTPStatusCode() int
	Header() http.Header
}

var retriableCodes = []int{
	http.StatusTooManyRequests,    // 429
	http.StatusBadGateway,         // 502
	http.StatusServiceUnavailable, // 503
	http.StatusGatewayTimeout,     // 504
}

// httpResponse extracts HTTP response metadata from an error, if available.
func httpResponse(err error) (statusCode int, header http.Header) {
	var retrieveErr *oauth2.RetrieveError
	if errors.As(err, &retrieveErr) && retrieveErr.Response != nil {
		return retrieveErr.Response.StatusCode, retrieveErr.Response.Header
	}
	var re httpResponseError
	if errors.As(err, &re) {
		return re.HTTPStatusCode(), re.Header()
	}
	return 0, nil
}

// isTransientNetworkError reports whether err represents a transient network
// condition that is likely to resolve on retry.
func isTransientNetworkError(err error) bool {
	if errors.Is(err, syscall.ECONNRESET) {
		return true
	}
	if errors.Is(err, syscall.ECONNREFUSED) {
		return true
	}
	var netErr net.Error
	if errors.As(err, &netErr) && netErr.Timeout() {
		return true
	}
	return false
}
