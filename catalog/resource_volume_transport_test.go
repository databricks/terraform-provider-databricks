package catalog

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
)

// These tests pin transport behavior of databricks_volume refreshes on both clients that
// replay faults cannot express: transport errors with a given message, a response body
// that keeps arriving past the HTTP timeout, and a panic below the client.

// scriptedVolumeTransport answers volume GETs with respond, numbering attempts from 1,
// and anything else with 404.
type scriptedVolumeTransport struct {
	attempts atomic.Int32
	respond  func(attempt int32, req *http.Request) (*http.Response, error)
}

func (s *scriptedVolumeTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.Path != retryVolumePath {
		recorder := httptest.NewRecorder()
		http.NotFound(recorder, req)
		resp := recorder.Result()
		resp.Request = req
		return resp, nil
	}
	return s.respond(s.attempts.Add(1), req)
}

func volumeResponse200(req *http.Request, body io.Reader) *http.Response {
	return &http.Response{StatusCode: http.StatusOK, Request: req,
		Header: http.Header{"Content-Type": {"application/json"}},
		Body:   io.NopCloser(body)}
}

// refreshVolumeWith runs the provider's wrapped Read, through sdk-go if sdkGo is set.
func refreshVolumeWith(ctx context.Context, c *common.DatabricksClient,
	sdkGo bool) (*schema.ResourceData, diag.Diagnostics) {
	r := volumeResource(sdkGo)
	d := r.TestResourceData()
	d.SetId(retryVolumeID)
	return d, r.ReadContext(volumeContext(ctx), d, c)
}

// refreshVolumeInBubble is refreshVolumeWith inside a synctest bubble. It then ends the
// call's context and waits, so the existing client's timeout goroutine exits before the
// bubble does.
func refreshVolumeInBubble(t *testing.T, c *common.DatabricksClient,
	sdkGo bool) (*schema.ResourceData, diag.Diagnostics) {
	ctx, cancel := context.WithCancel(t.Context())
	d, diags := refreshVolumeWith(ctx, c, sdkGo)
	cancel()
	synctest.Wait()
	return d, diags
}

// refreshedVolumePath is the volume_path that a refresh of retryVolumeJSON sets.
const refreshedVolumePath = "/Volumes/main/replay/retry_volume"

func TestVolumeRefreshRetriesTransientTransportErrors(t *testing.T) {
	forEachVolumeClient(t, testVolumeRefreshRetriesTransientTransportErrors)
}

func testVolumeRefreshRetriesTransientTransportErrors(t *testing.T, sdkGo bool) {
	// The classic client retries a transport error whose message contains one of these.
	for _, message := range []string{
		"net/http: TLS handshake timeout",
		"read tcp 10.0.0.1:443: i/o timeout",
		"Unexpected error from the connection",
	} {
		t.Run(message, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				api := &scriptedVolumeTransport{respond: func(attempt int32,
					req *http.Request) (*http.Response, error) {
					if attempt == 1 {
						return nil, errors.New(message)
					}
					return volumeResponse200(req, strings.NewReader(retryVolumeJSON)), nil
				}}
				c := newVolumeClient(t, &config.Config{Host: "https://volume-api.test",
					RetryTimeoutSeconds: retryUntilContextEnds, HTTPTransport: api})

				d, diags := refreshVolumeInBubble(t, c, sdkGo)

				assert.False(t, diags.HasError(), "%v", diags)
				assert.Equal(t, int32(2), api.attempts.Load())
				assert.Equal(t, refreshedVolumePath, d.Get("volume_path"))
			})
		})
	}
}

func TestVolumeRefreshDoesNotRetryOtherTransportErrors(t *testing.T) {
	forEachVolumeClient(t, testVolumeRefreshDoesNotRetryOtherTransportErrors)
}

func testVolumeRefreshDoesNotRetryOtherTransportErrors(t *testing.T, sdkGo bool) {
	api := &scriptedVolumeTransport{respond: func(int32,
		*http.Request) (*http.Response, error) {
		return nil, errors.New("write tcp 10.0.0.1:443: broken pipe")
	}}
	c := newVolumeClient(t, &config.Config{Host: "https://volume-api.test",
		RetryTimeoutSeconds: retryUntilContextEnds, HTTPTransport: api})

	d, diags := refreshVolumeWith(context.Background(), c, sdkGo)

	// The URL is the request's, which only the existing client ends with an empty query.
	summary := diagnosticSummary(t, diags)
	assert.True(t, strings.HasPrefix(summary, `cannot read volume: Get "https://`+
		`volume-api.test`+retryVolumePath), summary)
	assert.True(t, strings.HasSuffix(summary, `": write tcp 10.0.0.1:443: broken pipe`),
		summary)
	assert.Equal(t, int32(1), api.attempts.Load())
	assert.Equal(t, retryVolumeID, d.Id())
}

// trickleBody returns its data in chunks, waiting before each one. Like a network body,
// it fails once its request's context ends.
type trickleBody struct {
	ctx    context.Context
	chunks []string
	wait   time.Duration
}

func (b *trickleBody) Read(p []byte) (int, error) {
	if len(b.chunks) == 0 {
		return 0, io.EOF
	}
	select {
	case <-b.ctx.Done():
		return 0, b.ctx.Err()
	case <-time.After(b.wait):
	}
	n := copy(p, b.chunks[0])
	b.chunks[0] = b.chunks[0][n:]
	if b.chunks[0] == "" {
		b.chunks = b.chunks[1:]
	}
	return n, nil
}

func TestVolumeRefreshBodyProgressKeepsTheAttemptAlive(t *testing.T) {
	forEachVolumeClient(t, testVolumeRefreshBodyProgressKeepsTheAttemptAlive)
}

func testVolumeRefreshBodyProgressKeepsTheAttemptAlive(t *testing.T, sdkGo bool) {
	synctest.Test(t, func(t *testing.T) {
		// The body takes 2.4 s in all, but no chunk takes the full 1 s timeout.
		third := len(retryVolumeJSON) / 3
		api := &scriptedVolumeTransport{respond: func(_ int32,
			req *http.Request) (*http.Response, error) {
			return volumeResponse200(req, &trickleBody{ctx: req.Context(),
				wait: 600 * time.Millisecond,
				chunks: []string{retryVolumeJSON[:third], retryVolumeJSON[third : 2*third],
					retryVolumeJSON[2*third:], ""}}), nil
		}}
		c := newVolumeClient(t, &config.Config{Host: "https://volume-api.test",
			RetryTimeoutSeconds: retryUntilContextEnds, HTTPTimeoutSeconds: 1,
			HTTPTransport: api})

		d, diags := refreshVolumeInBubble(t, c, sdkGo)

		assert.False(t, diags.HasError(), "%v", diags)
		assert.Equal(t, int32(1), api.attempts.Load())
		assert.Equal(t, refreshedVolumePath, d.Get("volume_path"))
	})
}

func TestVolumeRefreshTurnsATransportPanicIntoAnError(t *testing.T) {
	forEachVolumeClient(t, testVolumeRefreshTurnsATransportPanicIntoAnError)
}

func testVolumeRefreshTurnsATransportPanicIntoAnError(t *testing.T, sdkGo bool) {
	api := &scriptedVolumeTransport{respond: func(int32,
		*http.Request) (*http.Response, error) {
		panic("transport exploded")
	}}
	c := newVolumeClient(t, &config.Config{Host: "https://volume-api.test",
		RetryTimeoutSeconds: retryUntilContextEnds, HTTPTransport: api})

	d, diags := refreshVolumeWith(context.Background(), c, sdkGo)

	require.Len(t, diags, 1)
	assert.Equal(t, "cannot read volume: panic: transport exploded", diags[0].Summary)
	assert.Equal(t, int32(1), api.attempts.Load())
	assert.Equal(t, retryVolumeID, d.Id(), "a panic must not remove the volume from state")
}
