package common

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	sdkcommon "github.com/databricks/databricks-sdk-go/common"
	sdkapierr "github.com/databricks/sdk-go/core/apierr"
	"golang.org/x/time/rate"
)

// classicError returns the error the classic client returns for the same failure. API
// errors are parsed again from the response, so apierr.IsMissing, the retry rules and
// the messages stay the same. req is the last request sent, if any.
func (b *sdkGoBridge) classicError(ctx context.Context, err error, req *http.Request) error {
	apiErr, ok := errors.AsType[*sdkapierr.APIError](err)
	if !ok || apiErr.HTTPStatusCode() < 0 {
		var syntaxErr *json.SyntaxError
		var typeErr *json.UnmarshalTypeError
		if errors.As(err, &syntaxErr) || errors.As(err, &typeErr) {
			return fmt.Errorf("failed to unmarshal response body: %w", err)
		}
		return err
	}
	if req == nil {
		req = &http.Request{URL: &url.URL{}}
	}
	body := apiErr.HTTPBody()
	resp := &http.Response{
		StatusCode: apiErr.HTTPStatusCode(),
		Header:     apiErr.HTTPHeader(),
		Body:       io.NopCloser(bytes.NewReader(body)),
		Request:    req,
	}
	mapped := b.httpConfig.ErrorMapper(ctx, sdkcommon.ResponseWrapper{
		Response:   resp,
		ReadCloser: io.NopCloser(bytes.NewReader(body)),
		DebugBytes: body,
	})
	if mapped == nil {
		return err
	}
	return mapped
}

// classicRetryableURLErrors mirrors the unexported list in httpclient.isRetriableUrlError.
var classicRetryableURLErrors = []string{
	"connection reset by peer",
	"TLS handshake timeout",
	"connection refused",
	"Unexpected error",
	"i/o timeout",
}

// isRetriable mirrors httpclient.ApiClient.isRetriable for an error from classicError.
func (b *sdkGoBridge) isRetriable(ctx context.Context, err error) bool {
	if b.httpConfig.ErrorRetriable != nil && b.httpConfig.ErrorRetriable(ctx, err) {
		return true
	}
	message := err.Error()
	if _, ok := errors.AsType[*url.Error](err); ok {
		for _, transient := range classicRetryableURLErrors {
			if strings.Contains(message, transient) {
				return true
			}
		}
	}
	for _, transient := range b.httpConfig.TransientErrors {
		if strings.Contains(message, transient) {
			return true
		}
	}
	return false
}

// classicRetrier retries what the classic client retries, with its backoff, until
// retry_timeout_seconds runs out.
type classicRetrier struct {
	ctx      context.Context
	bridge   *sdkGoBridge
	attempt  *sdkGoAttempt
	deadline time.Time
	retries  int
	lastErr  error
}

func (b *sdkGoBridge) newRetrier(ctx context.Context, attempt *sdkGoAttempt) *classicRetrier {
	r := &classicRetrier{ctx: ctx, bridge: b, attempt: attempt}
	if timeout := b.retry.Timeout(); timeout > 0 {
		r.deadline = time.Now().Add(timeout)
	}
	return r
}

func (r *classicRetrier) IsRetriable(err error) (time.Duration, bool) {
	classic := r.bridge.classicError(r.ctx, err, r.attempt.request)
	if !r.bridge.isRetriable(r.ctx, classic) {
		return 0, false
	}
	r.retries++
	wait := r.bridge.retry.Backoff(r.retries)
	// The classic client also stops when the deadline passes during the backoff.
	if !r.deadline.IsZero() && time.Now().Add(wait).After(r.deadline) {
		return 0, false
	}
	r.lastErr = classic
	return wait, true
}

// classicLimiter reports a rate limiter failure with the classic client's message.
type classicLimiter struct {
	limiter *rate.Limiter
}

func (l classicLimiter) Wait(ctx context.Context) error {
	if err := l.limiter.Wait(ctx); err != nil {
		return fmt.Errorf("failed in rate limiter: %w", err)
	}
	return nil
}
