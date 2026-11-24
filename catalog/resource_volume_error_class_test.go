package catalog

import (
	"net/http"
	"testing"
	"testing/synctest"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/stretchr/testify/assert"
)

// The SDK retries a 503 by its status until the retry timeout, then returns the last
// error. Its RESOURCE_DOES_NOT_EXIST code counts as missing, so the refresh removes the
// volume from state instead of failing. This runs in a synctest bubble because a replay
// journey would race the real clock.
func TestVolumeRefresh503ResourceDoesNotExistUntilRetryTimeoutIsMissing(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		api := &inMemoryVolumeAPI{respond: func(int32) (int, string) {
			return http.StatusServiceUnavailable, `{"error_code": "RESOURCE_DOES_NOT_EXIST",
				"message": "Volume 'main.replay.retry_volume' does not exist."}`
		}}
		// As in TestVolumeRefreshPersistent429FailsAtRetryTimeout, a 2 s limit on the fake
		// clock allows exactly one retry.
		c := newVolumeClient(t, &config.Config{
			Host:                "https://volume-api.test",
			RetryTimeoutSeconds: 2,
			HTTPTransport:       api,
		})

		d, diags := refreshVolume(t.Context(), c)

		assert.Empty(t, diags)
		assert.Equal(t, int32(2), api.attempts.Load())
		assert.Equal(t, "", d.Id(), "the volume must be removed from state")
	})
}
