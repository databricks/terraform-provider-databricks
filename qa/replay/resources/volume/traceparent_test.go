package volume_test

import (
	"net/http"
	"regexp"
	"strings"
	"sync"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestVolumeTraceparent pins the W3C traceparent header of every volume call,
// including both attempts of each retried call.
// It stays in Go because journey matchers cannot compare requests with each other.
func TestVolumeTraceparent(t *testing.T) {
	requests := runVolumeJourney(t, "native_testdata/traceparent_retry.yaml")
	var calls []string
	for _, r := range requests {
		calls = append(calls, r.method+" "+r.path)
	}
	path := "/api/2.1/unity-catalog/volumes"
	volume := path + "/main.replay.traceparent_volume"
	require.Equal(t, []string{
		"POST " + path, "POST " + path, "GET " + volume, "GET " + volume, "GET " + volume,
		"DELETE " + volume, "DELETE " + volume,
	}, calls)
	// Version 00 with the sampled flag set. Every attempt, a retry included, starts a
	// new trace with a new parent ID.
	syntax := regexp.MustCompile(`^00-([0-9a-f]{32})-([0-9a-f]{16})-01$`)
	traceIDs, parentIDs := map[string]bool{}, map[string]bool{}
	for i, r := range requests {
		values := r.header.Values("traceparent")
		require.Len(t, values, 1, calls[i])
		match := syntax.FindStringSubmatch(values[0])
		require.NotNil(t, match, "%s: %s", calls[i], values[0])
		assert.NotEqual(t, strings.Repeat("0", 32), match[1], calls[i])
		assert.NotEqual(t, strings.Repeat("0", 16), match[2], calls[i])
		assert.False(t, traceIDs[match[1]], "%s reuses trace %s", calls[i], match[1])
		assert.False(t, parentIDs[match[2]], "%s reuses parent %s", calls[i], match[2])
		traceIDs[match[1]], parentIDs[match[2]] = true, true
		assert.Empty(t, r.header.Values("tracestate"), calls[i])
	}
}

type volumeRequest struct {
	method string
	path   string
	header http.Header
}

type volumeRequestRecorder struct {
	mu       sync.Mutex
	requests []volumeRequest
}

// runVolumeJourney replays one native fixture and returns its volume requests as the
// server received them.
func runVolumeJourney(t *testing.T, path string) []volumeRequest {
	t.Helper()
	recorder := &volumeRequestRecorder{}
	terraformtest.RunNative(t, path,
		func(tc *resource.TestCase, customize func(*config.Config) error) {
			tc.ProtoV6ProviderFactories = terraformtest.ProviderFactories(
				func(c *config.Config) error {
					if err := customize(c); err != nil {
						return err
					}
					c.HTTPTransport = volumeRequestTransport{recorder, c.HTTPTransport}
					return nil
				})
		})
	recorder.mu.Lock()
	defer recorder.mu.Unlock()
	return recorder.requests
}

type volumeRequestTransport struct {
	recorder  *volumeRequestRecorder
	transport http.RoundTripper
}

func (v volumeRequestTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if strings.HasPrefix(req.URL.Path, "/api/2.1/unity-catalog/volumes") {
		request := volumeRequest{req.Method, req.URL.Path, req.Header.Clone()}
		v.recorder.mu.Lock()
		v.recorder.requests = append(v.recorder.requests, request)
		v.recorder.mu.Unlock()
	}
	return v.transport.RoundTrip(req)
}
