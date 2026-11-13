package catalog

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"testing/synctest"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
)

// These tests repeat the tests of resource_volume_retry_test.go with reads through
// sdk-go uc/volumes/v1.

// refreshVolumeWithSdkGo runs the provider's wrapped Read with sdk-go selected.
func refreshVolumeWithSdkGo(ctx context.Context,
	c *common.DatabricksClient) (*schema.ResourceData, diag.Diagnostics) {
	r := ResourceVolume().ToResource()
	common.EnableSdkGo(r)
	d := r.TestResourceData()
	d.SetId(retryVolumeID)
	ctx = context.WithValue(ctx, common.ResourceName, "volume")
	return d, r.ReadContext(ctx, d, c)
}

func TestVolumeSdkGoRefreshPersistent429FailsAtRetryTimeout(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		api := &inMemoryVolumeAPI{respond: func(int32) (int, string) {
			return http.StatusTooManyRequests, tooManyRequestsJSON
		}}
		// As on the existing path, a 2 s limit allows exactly one retry.
		c := newVolumeClient(t, &config.Config{
			Host:                "https://volume-api.test",
			RetryTimeoutSeconds: 2,
			HTTPTransport:       api,
		})

		d, diags := refreshVolumeWithSdkGo(t.Context(), c)

		assert.Equal(t, "cannot read volume: Too many requests from fixture", diagnosticSummary(t, diags))
		assert.Equal(t, int32(2), api.attempts.Load())
		assert.Equal(t, retryVolumeID, d.Id(), "a 429 must not remove the volume from state")
	})
}

func TestVolumeSdkGoRefreshPastHTTPTimeoutIsNotRetried(t *testing.T) {
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

	d, diags := refreshVolumeWithSdkGo(context.Background(), c)

	// sdk-go sends no empty query, so the URL has no trailing "?", unlike on the
	// existing path.
	expected := fmt.Sprintf(`cannot read volume: Get "%s%s": request timed out after 1s of inactivity`,
		server.URL, retryVolumePath)
	assert.Equal(t, expected, diagnosticSummary(t, diags))
	assert.Equal(t, int32(1), attempts.Load())
	assert.Equal(t, retryVolumeID, d.Id(), "a timed out refresh must not remove the volume from state")
}

func TestVolumeSdkGoRefreshCancelledDuringBackoffIsNotRetried(t *testing.T) {
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
			d, diags = refreshVolumeWithSdkGo(ctx, c)
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
