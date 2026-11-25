package catalog

import (
	"context"
	"net/http"
	"testing"
	"testing/synctest"
	"time"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/stretchr/testify/assert"
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
