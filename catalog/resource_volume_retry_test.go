package catalog

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
)

// These tests pin the classic SDK client's retry timeout, HTTP timeout and cancellation
// behavior for databricks_volume refreshes, which replay fixtures cannot express.

const (
	retryVolumeID   = "main.replay.retry_volume"
	retryVolumePath = "/api/2.1/unity-catalog/volumes/" + retryVolumeID
	retryVolumeJSON = `{"catalog_name": "main", "schema_name": "replay", "name": "retry_volume",
		"full_name": "main.replay.retry_volume", "volume_type": "MANAGED"}`
	tooManyRequestsJSON = `{"error_code": "TOO_MANY_REQUESTS",
		"message": "Too many requests from fixture"}`
)

// volumeCallObserver counts the requests for retryVolumePath. Once each such response
// body is closed, which the SDK client does before deciding whether to retry, it calls
// afterResponse.
type volumeCallObserver struct {
	next          http.RoundTripper
	attempts      atomic.Int32
	afterResponse func()
}

func (o *volumeCallObserver) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.Path != retryVolumePath {
		return o.next.RoundTrip(req)
	}
	o.attempts.Add(1)
	resp, err := o.next.RoundTrip(req)
	if err == nil && o.afterResponse != nil {
		resp.Body = &bodyCloseHook{ReadCloser: resp.Body, onClose: o.afterResponse}
	}
	return resp, err
}

type bodyCloseHook struct {
	io.ReadCloser
	once    sync.Once
	onClose func()
}

func (b *bodyCloseHook) Close() error {
	err := b.ReadCloser.Close()
	b.once.Do(b.onClose)
	return err
}

// observeVolumeCalls installs an observer on a fixture client. The workspace client is
// built from the config on first use, so the observer and later config edits apply to
// every volume call.
func observeVolumeCalls(c *common.DatabricksClient) *volumeCallObserver {
	observer := &volumeCallObserver{next: http.DefaultTransport}
	c.Config.HTTPTransport = observer
	return observer
}

// refreshVolume runs the provider's wrapped Read, as a Terraform refresh does.
func refreshVolume(ctx context.Context, c *common.DatabricksClient) (*schema.ResourceData, diag.Diagnostics) {
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
	c, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:       http.MethodGet,
			Resource:     retryVolumePath + "?",
			Status:       http.StatusTooManyRequests,
			Response:     tooManyRequestsJSON,
			ReuseRequest: true,
		},
	})
	require.NoError(t, err)
	defer server.Close()
	observer := observeVolumeCalls(c)
	// Backoff waits at least 1.05 s before the first retry and 2.05 s before the second,
	// so a 2 s limit allows exactly one retry.
	c.Config.RetryTimeoutSeconds = 2

	d, diags := refreshVolume(context.Background(), c)

	assert.Equal(t, "cannot read volume: Too many requests from fixture", diagnosticSummary(t, diags))
	assert.Equal(t, int32(2), observer.attempts.Load())
	assert.Equal(t, retryVolumeID, d.Id(), "a 429 must not remove the volume from state")
}

func TestVolumeRefreshPastHTTPTimeoutIsNotRetried(t *testing.T) {
	// qa.HTTPFixture cannot delay a response, so this uses httptest. Only the first GET
	// is held past the HTTP timeout; a retry would succeed and fail the assertions
	// below instead of hanging.
	release := make(chan struct{})
	var served atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != retryVolumePath {
			http.NotFound(w, r)
			return
		}
		if served.Add(1) == 1 {
			select {
			case <-r.Context().Done():
			case <-release:
			}
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(retryVolumeJSON))
	}))
	defer server.Close()
	// Runs before server.Close, which waits for the held handler to return.
	defer close(release)
	observer := &volumeCallObserver{next: http.DefaultTransport}
	inner, err := client.New(&config.Config{
		Host:  server.URL,
		Token: "fixture-token",
		// The provider's default, which retries until the context ends.
		RetryTimeoutSeconds: -1,
		HTTPTimeoutSeconds:  1,
		HTTPTransport:       observer,
	})
	require.NoError(t, err)
	c := &common.DatabricksClient{DatabricksClient: inner}
	c.SetCachedWorkspaceID(12345)

	d, diags := refreshVolume(context.Background(), c)

	expected := fmt.Sprintf(`cannot read volume: Get "%s%s?": request timed out after 1s of inactivity`,
		server.URL, retryVolumePath)
	assert.Equal(t, expected, diagnosticSummary(t, diags))
	assert.Equal(t, int32(1), observer.attempts.Load())
	assert.Equal(t, retryVolumeID, d.Id(), "a timed out refresh must not remove the volume from state")
}

func TestVolumeRefreshCancelledDuringBackoffIsNotRetried(t *testing.T) {
	c, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: retryVolumePath + "?",
			Status:   http.StatusTooManyRequests,
			Response: tooManyRequestsJSON,
		},
		{
			Method:   http.MethodGet,
			Resource: retryVolumePath + "?",
			Response: retryVolumeJSON,
		},
	})
	require.NoError(t, err)
	defer server.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	observer := observeVolumeCalls(c)
	// Cancelling once the 429 body is read lands in the backoff before the first retry,
	// which waits at least 1.05 s. The retry would succeed if it were sent.
	observer.afterResponse = cancel
	// The provider's default, which retries until the context ends.
	c.Config.RetryTimeoutSeconds = -1

	d, diags := refreshVolume(ctx, c)

	assert.Equal(t, "cannot read volume: Too many requests from fixture", diagnosticSummary(t, diags))
	assert.Equal(t, int32(1), observer.attempts.Load())
	assert.Equal(t, retryVolumeID, d.Id(), "a cancelled refresh must not remove the volume from state")
}
