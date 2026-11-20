package catalog

import (
	"context"
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

// This test pins the classic SDK client's handling of a databricks_volume refresh
// cancelled during its backoff. The SDK backs off on the real clock, so it runs in a
// synctest bubble, whose fake clock makes its timing exact. The retry_refresh_ replay
// journeys pin retry exhaustion and the HTTP timeout.

const (
	retryVolumeID   = "main.replay.retry_volume"
	retryVolumePath = "/api/2.1/unity-catalog/volumes/" + retryVolumeID
	retryVolumeJSON = `{"catalog_name": "main", "schema_name": "replay",
		"name": "retry_volume", "full_name": "main.replay.retry_volume",
		"volume_type": "MANAGED"}`
	tooManyRequestsJSON = `{"error_code": "TOO_MANY_REQUESTS",
		"message": "Too many requests from fixture"}`
	tooManyRequestsDiagnostic = "cannot read volume: Too many requests from fixture"
	// retryUntilContextEnds is the provider's default retry timeout.
	retryUntilContextEnds = -1
)

// inMemoryVolumeAPI answers volume GETs with respond, numbering attempts from 1, and
// anything else with 404. An HTTP server cannot serve a client inside a synctest
// bubble.
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

// Replay cannot cancel Terraform during a backoff, so this cancellation stays in Go.
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
