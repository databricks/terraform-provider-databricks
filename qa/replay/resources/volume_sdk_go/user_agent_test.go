package volume_sdk_go_test

import (
	"net/http"
	"strings"
	"sync"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/assert"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestVolumeSdkGoUserAgent pins that sdk-go volume calls keep the existing User-Agent
// fields and add the sdk-go module. YAML header matchers compare exact values, so the
// substrings are checked in Go.
func TestVolumeSdkGoUserAgent(t *testing.T) {
	recorder := &userAgentRecorder{}
	terraformtest.RunNative(t, "native_testdata/user_agent.yaml",
		func(tc *resource.TestCase, customize func(*config.Config) error) {
			tc.ProtoV6ProviderFactories = terraformtest.ProviderFactoriesWithSdkGo(
				func(c *config.Config) error {
					if err := customize(c); err != nil {
						return err
					}
					c.HTTPTransport = userAgentTransport{recorder, c.HTTPTransport}
					return nil
				}, []string{"databricks_volume"})
		})
	methods := map[string]bool{}
	for _, call := range recorder.snapshot() {
		methods[call.method] = true
		request := call.method + " " + call.path
		assert.Contains(t, call.userAgent, "databricks-tf-provider/", request)
		assert.Contains(t, call.userAgent, "resource/volume", request)
		assert.Contains(t, call.userAgent, " sdk-go-uc-volumes/", request)
	}
	want := map[string]bool{"POST": true, "PATCH": true, "GET": true, "DELETE": true}
	assert.Equal(t, want, methods)
}

type userAgentCall struct {
	method    string
	path      string
	userAgent string
}

type userAgentRecorder struct {
	mu    sync.Mutex
	calls []userAgentCall
}

func (r *userAgentRecorder) record(req *http.Request) {
	if !strings.HasPrefix(req.URL.Path, "/api/2.1/unity-catalog/volumes") {
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	call := userAgentCall{req.Method, req.URL.Path, req.Header.Get("User-Agent")}
	r.calls = append(r.calls, call)
}

func (r *userAgentRecorder) snapshot() []userAgentCall {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]userAgentCall(nil), r.calls...)
}

type userAgentTransport struct {
	recorder  *userAgentRecorder
	transport http.RoundTripper
}

func (u userAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	u.recorder.record(req)
	return u.transport.RoundTrip(req)
}
