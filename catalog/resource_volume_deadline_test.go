package catalog

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/synctest"
	"time"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// A refresh whose context deadline passes during the backoff after a 429 fails with the
// 429, not with the deadline, and sends no retry. Terraform gives every operation such
// a deadline, its timeout.
func TestVolumeRefreshPastItsDeadlineDuringBackoffFailsWithTheLastError(t *testing.T) {
	forEachVolumeClient(t, testVolumeRefreshPastItsDeadlineDuringBackoffFailsWithTheLastError)
}

func testVolumeRefreshPastItsDeadlineDuringBackoffFailsWithTheLastError(t *testing.T,
	sdkGo bool) {
	synctest.Test(t, func(t *testing.T) {
		// The retry would succeed if it were sent.
		api := &inMemoryVolumeAPI{respond: func(attempt int32) (int, string) {
			if attempt == 1 {
				return http.StatusTooManyRequests, tooManyRequestsJSON
			}
			return http.StatusOK, retryVolumeJSON
		}}
		c := newVolumeClient(t, &config.Config{
			Host:                "https://volume-api.test",
			RetryTimeoutSeconds: retryUntilContextEnds,
			HTTPTransport:       api,
		})
		// The first backoff waits at least 1.05 s.
		ctx, cancel := context.WithTimeout(t.Context(), time.Second)
		defer cancel()

		d, diags := refreshVolumeWith(ctx, c, sdkGo)

		assert.Equal(t, tooManyRequestsDiagnostic, diagnosticSummary(t, diags))
		assert.Equal(t, int32(1), api.attempts.Load())
		assert.Equal(t, retryVolumeID, d.Id(),
			"a refresh past its deadline must not remove the volume from state")
		cancel()
		synctest.Wait()
	})
}

// A refresh keeps retrying 429s for longer than five minutes, sdk-go's default call
// timeout, until retry_timeout_seconds ends, and then succeeds.
func TestVolumeRefreshRetriesPastFiveMinutesWithinRetryTimeout(t *testing.T) {
	forEachVolumeClient(t, testVolumeRefreshRetriesPastFiveMinutesWithinRetryTimeout)
}

func testVolumeRefreshRetriesPastFiveMinutesWithinRetryTimeout(t *testing.T, sdkGo bool) {
	synctest.Test(t, func(t *testing.T) {
		start := time.Now()
		api := &inMemoryVolumeAPI{respond: func(int32) (int, string) {
			if time.Since(start) < 6*time.Minute {
				return http.StatusTooManyRequests, tooManyRequestsJSON
			}
			return http.StatusOK, retryVolumeJSON
		}}
		c := newVolumeClient(t, &config.Config{
			Host:                "https://volume-api.test",
			RetryTimeoutSeconds: 600,
			HTTPTransport:       api,
		})

		d, diags := refreshVolumeInBubble(t, c, sdkGo)

		require.False(t, diags.HasError(), "%v", diags)
		assert.Equal(t, refreshedVolumePath, d.Get("volume_path"))
		assert.GreaterOrEqual(t, time.Since(start), 6*time.Minute)
	})
}

// cancelOnTimeoutTransport holds each volume request until its attempt's inactivity
// timeout cancels it, then ends the caller's context too before returning. It answers
// anything else, such as host metadata, with 404.
type cancelOnTimeoutTransport struct {
	cancelCaller context.CancelFunc
}

func (c cancelOnTimeoutTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.Path != retryVolumePath {
		recorder := httptest.NewRecorder()
		http.NotFound(recorder, req)
		resp := recorder.Result()
		resp.Request = req
		return resp, nil
	}
	<-req.Context().Done()
	c.cancelCaller()
	return nil, req.Context().Err()
}

// When the caller's context ends while an attempt times out, the refresh fails with
// the cancellation, not with the inactivity timeout, and isn't retried.
func TestVolumeRefreshCancelledAsItTimesOutKeepsTheCancellation(t *testing.T) {
	forEachVolumeClient(t, testVolumeRefreshCancelledAsItTimesOutKeepsTheCancellation)
}

func testVolumeRefreshCancelledAsItTimesOutKeepsTheCancellation(t *testing.T,
	sdkGo bool) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(t.Context())
		defer cancel()
		c := newVolumeClient(t, &config.Config{
			Host:                "https://volume-api.test",
			RetryTimeoutSeconds: retryUntilContextEnds,
			HTTPTimeoutSeconds:  1,
			HTTPTransport:       cancelOnTimeoutTransport{cancelCaller: cancel},
		})

		d, diags := refreshVolumeWith(ctx, c, sdkGo)
		synctest.Wait()

		assert.Equal(t, `cannot read volume: Get "https://volume-api.test`+retryVolumePath+
			`?": context canceled`, diagnosticSummary(t, diags))
		assert.Equal(t, retryVolumeID, d.Id())
	})
}
