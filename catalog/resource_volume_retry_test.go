package catalog

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"testing/synctest"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
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
	// The provider's default retry timeout, which retries until the context ends.
	retryUntilContextEnds = -1
)

// inMemoryVolumeAPI answers volume GETs with respond, numbering attempts from 1, and 404
// to anything else. The SDK backs off on the real clock, so the backoff tests run in a
// synctest bubble, whose fake clock makes it deterministic. A qa.HTTPFixture server
// cannot serve a client inside a bubble, so requests are served in memory instead.
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

		assert.Equal(t, "cannot read volume: Too many requests from fixture", diagnosticSummary(t, diags))
		assert.Equal(t, int32(2), api.attempts.Load())
		assert.Equal(t, retryVolumeID, d.Id(), "a 429 must not remove the volume from state")
	})
}

func TestVolumeRefreshPastHTTPTimeoutIsNotRetried(t *testing.T) {
	// qa.HTTPFixture cannot delay a response, so this uses httptest. Only the first GET
	// is held past the HTTP timeout; a retry would succeed and fail the assertions
	// below instead of hanging.
	release := make(chan struct{})
	var attempts atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	}))
	defer server.Close()
	// Runs before server.Close, which waits for the held handler to return.
	defer close(release)
	c := newVolumeClient(t, &config.Config{
		Host:                server.URL,
		RetryTimeoutSeconds: retryUntilContextEnds,
		HTTPTimeoutSeconds:  1,
	})

	d, diags := refreshVolume(context.Background(), c)

	expected := fmt.Sprintf(`cannot read volume: Get "%s%s?": request timed out after 1s of inactivity`,
		server.URL, retryVolumePath)
	assert.Equal(t, expected, diagnosticSummary(t, diags))
	assert.Equal(t, int32(1), attempts.Load())
	assert.Equal(t, retryVolumeID, d.Id(), "a timed out refresh must not remove the volume from state")
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

		// Wait returns once the refresh blocks, which after the 429 is in the backoff.
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

		assert.Equal(t, "cannot read volume: Too many requests from fixture", diagnosticSummary(t, diags))
		assert.Equal(t, int32(1), api.attempts.Load())
		assert.Equal(t, retryVolumeID, d.Id(), "a cancelled refresh must not remove the volume from state")
	})
}
