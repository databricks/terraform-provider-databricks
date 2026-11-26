package catalog

import (
	"context"
	"net/http"
	"testing"
	"testing/synctest"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// A replay journey cannot cancel Terraform while the SDK's rate limiter is waiting.
func TestVolumeRefreshCancelledInRateLimiterWaitSendsNoRequest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		api := &inMemoryVolumeAPI{respond: func(int32) (int, string) {
			return http.StatusOK, retryVolumeJSON
		}}
		// One request per second with a burst of one, so a refresh right after another
		// waits about 1 s for a token.
		c := newVolumeClient(t, &config.Config{
			Host:               "https://volume-api.test",
			RateLimitPerSecond: 1,
			HTTPTransport:      api,
		})
		_, diags := refreshVolume(t.Context(), c)
		require.Empty(t, diags)
		require.Equal(t, int32(1), api.attempts.Load())

		ctx, cancel := context.WithCancel(t.Context())
		defer cancel()
		var d *schema.ResourceData
		done := make(chan struct{})
		go func() {
			defer close(done)
			d, diags = refreshVolume(ctx, c)
		}()

		// The second refresh blocks only in the rate limiter's wait, so Wait returns
		// there.
		synctest.Wait()
		cancel()
		// Wait does not advance the fake clock, so the refresh must return without the
		// limiter's timer firing.
		synctest.Wait()
		select {
		case <-done:
		default:
			t.Fatal("the refresh kept waiting for a token after cancellation")
		}

		assert.Equal(t, int32(1), api.attempts.Load(), "the second refresh sent a request")
		assert.Equal(t, "cannot read volume: failed in rate limiter: context canceled",
			diagnosticSummary(t, diags))
		assert.Equal(t, retryVolumeID, d.Id(), "the volume must stay in state")
	})
}
