package catalog

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
)

// These tests pin the classic SDK client's retry timeout, backoff, HTTP timeout, rate
// limiter and cancellation behavior for databricks_volume refreshes, and its parsing of
// a status line without a reason phrase, which replay fixtures cannot express. The SDK
// backs off and rate limits on the real clock, so those tests run in a synctest bubble,
// whose fake clock makes their timing exact.

const (
	retryVolumeID   = "main.replay.retry_volume"
	retryVolumePath = "/api/2.1/unity-catalog/volumes/" + retryVolumeID
	retryVolumeJSON = `{"catalog_name": "main", "schema_name": "replay",
		"name": "retry_volume", "full_name": "main.replay.retry_volume",
		"volume_type": "MANAGED"}`
	tooManyRequestsJSON = `{"error_code": "TOO_MANY_REQUESTS",
		"message": "Too many requests from fixture"}`
	tooManyRequestsDiagnostic = "cannot read volume: Too many requests from fixture"
	unavailableJSON           = `{"error_code": "TEMPORARILY_UNAVAILABLE",
		"message": "Volume service unavailable"}`
	// retryUntilContextEnds is the provider's default retry timeout.
	retryUntilContextEnds = -1
)

// inMemoryVolumeAPI answers volume GETs with respond, numbering attempts from 1, and
// anything else with 404. The backoff tests use it because a qa.HTTPFixture server
// cannot serve a client inside a synctest bubble.
type inMemoryVolumeAPI struct {
	attempts atomic.Int32
	respond  func(attempt int32) (status int, body string)
}

func (api *inMemoryVolumeAPI) RoundTrip(req *http.Request) (*http.Response, error) {
	recorder := httptest.NewRecorder()
	if req.URL.Path == retryVolumePath {
		status, body := api.respond(api.attempts.Add(1))
		recorder.Header().Set("Content-Type", "application/json")
		recorder.WriteHeader(status)
		_, _ = recorder.WriteString(body)
	} else {
		http.NotFound(recorder, req)
	}
	resp := recorder.Result()
	resp.Request = req
	return resp, nil
}

func newVolumeClient(t *testing.T, cfg *config.Config) *common.DatabricksClient {
	t.Helper()
	cfg.Token = "fixture-token"
	inner, err := client.New(cfg)
	require.NoError(t, err)
	c := &common.DatabricksClient{DatabricksClient: inner}
	c.SetCachedWorkspaceID(12345)
	return c
}

// refreshVolume runs the provider's wrapped Read, as a Terraform refresh does.
func refreshVolume(
	ctx context.Context, c *common.DatabricksClient,
) (*schema.ResourceData, diag.Diagnostics) {
	r := ResourceVolume().ToResource()
	d := r.TestResourceData()
	d.SetId(retryVolumeID)
	ctx = context.WithValue(ctx, common.ResourceName, "volume")
	return d, r.ReadContext(ctx, d, c)
}

func diagnosticSummary(t *testing.T, diags diag.Diagnostics) string {
	t.Helper()
	require.Len(t, diags, 1)
	return diags[0].Summary
}

func TestVolumeRefreshPersistent429FailsAtRetryTimeout(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		api := &inMemoryVolumeAPI{respond: func(int32) (int, string) {
			return http.StatusTooManyRequests, tooManyRequestsJSON
		}}
		// Backoff waits 1.05 to 1.75 s before the first retry and at least 2.05 s before
		// the second, so on the fake clock a 2 s limit allows exactly one retry.
		c := newVolumeClient(t, &config.Config{
			Host:                "https://volume-api.test",
			RetryTimeoutSeconds: 2,
			HTTPTransport:       api,
		})

		d, diags := refreshVolume(t.Context(), c)

		assert.Equal(t, tooManyRequestsDiagnostic, diagnosticSummary(t, diags))
		assert.Equal(t, int32(2), api.attempts.Load())
		assert.Equal(t, retryVolumeID, d.Id(), "the volume must stay in state")
	})
}

func TestVolumeRefreshPastHTTPTimeoutIsNotRetried(t *testing.T) {
	// qa.HTTPFixture cannot delay a response, so this uses httptest. Only the first GET
	// is held past the HTTP timeout; a retry would succeed and fail the assertions
	// below instead of hanging.
	release := make(chan struct{})
	var attempts atomic.Int32
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != retryVolumePath {
			http.NotFound(w, r)
			return
		}
		if attempts.Add(1) == 1 {
			select {
			case <-r.Context().Done():
			case <-release:
			}
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(retryVolumeJSON))
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	// This runs before server.Close, which waits for the held handler to return.
	defer close(release)
	c := newVolumeClient(t, &config.Config{
		Host:                server.URL,
		RetryTimeoutSeconds: retryUntilContextEnds,
		HTTPTimeoutSeconds:  1,
	})

	d, diags := refreshVolume(context.Background(), c)

	expected := fmt.Sprintf(
		`cannot read volume: Get "%s%s?": request timed out after 1s of inactivity`,
		server.URL, retryVolumePath)
	assert.Equal(t, expected, diagnosticSummary(t, diags))
	assert.Equal(t, int32(1), attempts.Load())
	assert.Equal(t, retryVolumeID, d.Id(), "the volume must stay in state")
}

func TestVolumeRefreshCancelledDuringBackoffIsNotRetried(t *testing.T) {
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
		ctx, cancel := context.WithCancel(t.Context())
		defer cancel()
		var d *schema.ResourceData
		var diags diag.Diagnostics
		done := make(chan struct{})
		go func() {
			defer close(done)
			d, diags = refreshVolume(ctx, c)
		}()

		// After the 429, the refresh blocks only in its backoff, so Wait returns there.
		synctest.Wait()
		require.Equal(t, int32(1), api.attempts.Load())
		cancel()
		// Wait does not advance the fake clock, so the refresh must return without its
		// backoff timer firing.
		synctest.Wait()
		select {
		case <-done:
		default:
			t.Fatal("the refresh kept waiting out its backoff after cancellation")
		}

		assert.Equal(t, tooManyRequestsDiagnostic, diagnosticSummary(t, diags))
		assert.Equal(t, int32(1), api.attempts.Load())
		assert.Equal(t, retryVolumeID, d.Id(), "the volume must stay in state")
	})
}

func TestVolumeRefreshBackoffStopsGrowingAtTenSeconds(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// The SDK retries every 503. Retry n waits n seconds, capped at 10, plus 50 to
		// 750 ms of jitter. On the fake clock, each gap between requests is exactly
		// that wait.
		baseWaitSeconds := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10}
		var arrivals []time.Time
		api := &inMemoryVolumeAPI{respond: func(attempt int32) (int, string) {
			arrivals = append(arrivals, time.Now())
			if int(attempt) <= len(baseWaitSeconds) {
				return http.StatusServiceUnavailable, unavailableJSON
			}
			return http.StatusOK, retryVolumeJSON
		}}
		c := newVolumeClient(t, &config.Config{
			Host:                "https://volume-api.test",
			RetryTimeoutSeconds: retryUntilContextEnds,
			HTTPTransport:       api,
		})

		d, diags := refreshVolume(t.Context(), c)

		require.Empty(t, diags)
		assert.Equal(t, "/Volumes/main/replay/retry_volume", d.Get("volume_path"))
		require.Len(t, arrivals, len(baseWaitSeconds)+1)
		for i, seconds := range baseWaitSeconds {
			wait := time.Duration(seconds) * time.Second
			gap := arrivals[i+1].Sub(arrivals[i])
			assert.GreaterOrEqual(t, gap, wait+50*time.Millisecond, "retry %d", i+1)
			assert.LessOrEqual(t, gap, wait+750*time.Millisecond, "retry %d", i+1)
		}
	})
}

// newRateLimitedVolumeClient returns a client with rate_limit = 1, whose limiter holds
// one token and refills it each second, after a first refresh has taken that token.
func newRateLimitedVolumeClient(t *testing.T) (*inMemoryVolumeAPI, *common.DatabricksClient) {
	t.Helper()
	api := &inMemoryVolumeAPI{respond: func(int32) (int, string) {
		return http.StatusOK, retryVolumeJSON
	}}
	c := newVolumeClient(t, &config.Config{
		Host:                "https://volume-api.test",
		RetryTimeoutSeconds: retryUntilContextEnds,
		RateLimitPerSecond:  1,
		HTTPTransport:       api,
	})
	_, diags := refreshVolume(t.Context(), c)
	require.Empty(t, diags)
	require.Equal(t, int32(1), api.attempts.Load())
	return api, c
}

func TestVolumeRefreshWaitingOnRateLimiterFailsWhenItsContextEnds(t *testing.T) {
	t.Run("cancelled while waiting", func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			api, c := newRateLimitedVolumeClient(t)
			ctx, cancel := context.WithCancel(t.Context())
			defer cancel()
			var d *schema.ResourceData
			var diags diag.Diagnostics
			done := make(chan struct{})
			go func() {
				defer close(done)
				d, diags = refreshVolume(ctx, c)
			}()

			// The refresh blocks only on the limiter, so Wait returns there.
			synctest.Wait()
			cancel()
			// Wait does not advance the fake clock, so the refresh must return before
			// the limiter has a token.
			synctest.Wait()
			select {
			case <-done:
			default:
				t.Fatal("the refresh kept waiting on the rate limiter after cancellation")
			}

			assert.Equal(t, "cannot read volume: failed in rate limiter: context canceled",
				diagnosticSummary(t, diags))
			assert.Equal(t, int32(1), api.attempts.Load(), "the GET must not be sent")
			assert.Equal(t, retryVolumeID, d.Id(), "the volume must stay in state")
		})
	})
	t.Run("deadline before the next token", func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			api, c := newRateLimitedVolumeClient(t)
			ctx, cancel := context.WithTimeout(t.Context(), 500*time.Millisecond)
			defer cancel()
			start := time.Now()

			d, diags := refreshVolume(ctx, c)

			// The limiter fails at once instead of waiting out the deadline.
			assert.Zero(t, time.Since(start))
			assert.Equal(t, "cannot read volume: failed in rate limiter: "+
				"rate: Wait(n=1) would exceed context deadline", diagnosticSummary(t, diags))
			assert.Equal(t, int32(1), api.attempts.Load(), "the GET must not be sent")
			assert.Equal(t, retryVolumeID, d.Id(), "the volume must stay in state")
		})
	})
}

// writeRawHTMLResponse writes statusLine and page to the hijacked connection, because
// net/http always adds a reason phrase to the status line it writes.
func writeRawHTMLResponse(t *testing.T, w http.ResponseWriter, statusLine, page string) {
	conn, rw, err := http.NewResponseController(w).Hijack()
	if err != nil {
		t.Errorf("cannot hijack the connection: %v", err)
		return
	}
	defer conn.Close()
	_, _ = fmt.Fprintf(rw, "%s\r\nContent-Type: text/html\r\nContent-Length: %d\r\n"+
		"Connection: close\r\n\r\n%s", statusLine, len(page), page)
	if err := rw.Flush(); err != nil {
		t.Errorf("cannot write the response: %v", err)
	}
}

func TestVolumeRefreshAnsweredHTML404StatusLine(t *testing.T) {
	const page = "<html><body><h2>HTTP ERROR 404</h2>" +
		"<pre>Volume main.replay.retry_volume not found.</pre></body></html>"
	tests := []struct {
		name       string
		statusLine string
		// diagnostic is empty when the refresh succeeds.
		diagnostic string
		id         string
	}{
		{
			// The SDK takes the error code from the reason phrase, and uses UNKNOWN
			// without one. UNKNOWN maps to ErrUnknown, which takes precedence over the
			// 404, so the volume does not count as missing.
			name:       "without a reason phrase fails and keeps the volume",
			statusLine: "HTTP/1.1 404",
			diagnostic: "cannot read volume: Volume main.replay.retry_volume not found",
			id:         retryVolumeID,
		},
		{
			name:       "with a reason phrase drops the volume",
			statusLine: "HTTP/1.1 404 Not Found",
			id:         "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var attempts atomic.Int32
			server := httptest.NewServer(http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					if r.URL.Path != retryVolumePath {
						http.NotFound(w, r)
						return
					}
					attempts.Add(1)
					writeRawHTMLResponse(t, w, tt.statusLine, page)
				}))
			defer server.Close()
			c := newVolumeClient(t, &config.Config{
				Host:                server.URL,
				RetryTimeoutSeconds: retryUntilContextEnds,
			})

			d, diags := refreshVolume(context.Background(), c)

			if tt.diagnostic == "" {
				assert.Empty(t, diags)
			} else {
				assert.Equal(t, tt.diagnostic, diagnosticSummary(t, diags))
			}
			assert.Equal(t, int32(1), attempts.Load(), "a 404 must not be retried")
			assert.Equal(t, tt.id, d.Id())
		})
	}
}
