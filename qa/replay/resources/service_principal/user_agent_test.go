package service_principal_test

import (
	"net/http"
	"strings"
	"sync"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/assert"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestServicePrincipalUserAgent pins the User-Agent of service principal calls.
// YAML header matchers compare exact values, so the substrings are checked in Go.
func TestServicePrincipalUserAgent(t *testing.T) {
	calls := recordUserAgents(t, "native_testdata/user_agent.yaml",
		terraformtest.ProviderFactories)
	methods := map[string]bool{}
	for _, call := range calls {
		methods[call.method] = true
		request := call.method + " " + call.path
		assert.Contains(t, call.userAgent, "databricks-tf-provider/", request)
		assert.Contains(t, call.userAgent, "resource/service_principal", request)
	}
	assert.Equal(t, map[string]bool{"POST": true, "GET": true, "DELETE": true}, methods)
}

// TestServicePrincipalSdkGoUserAgent pins that sdk-go reads keep the existing
// User-Agent fields and add the sdk-go module, while writes keep the existing one.
func TestServicePrincipalSdkGoUserAgent(t *testing.T) {
	calls := recordUserAgents(t, "native_testdata/sdk_go_user_agent.yaml",
		func(customize func(*config.Config) error) providerFactories {
			return terraformtest.ProviderFactoriesWithSdkGo(customize,
				[]string{"databricks_service_principal"})
		})
	methods := map[string]bool{}
	for _, call := range calls {
		methods[call.method] = true
		request := call.method + " " + call.path
		assert.Contains(t, call.userAgent, "databricks-tf-provider/", request)
		assert.Contains(t, call.userAgent, "resource/service_principal", request)
		if call.method == "GET" {
			assert.Contains(t, call.userAgent, " sdk-go-scim/", request)
		} else {
			assert.NotContains(t, call.userAgent, "sdk-go-", request)
		}
	}
	assert.Equal(t, map[string]bool{"POST": true, "GET": true, "DELETE": true}, methods)
}

type providerFactories = map[string]func() (tfprotov6.ProviderServer, error)

func recordUserAgents(t *testing.T, path string,
	factories func(func(*config.Config) error) providerFactories) []userAgentCall {
	recorder := &userAgentRecorder{}
	terraformtest.RunNative(t, path,
		func(tc *resource.TestCase, customize func(*config.Config) error) {
			tc.ProtoV6ProviderFactories = factories(
				func(c *config.Config) error {
					if err := customize(c); err != nil {
						return err
					}
					c.HTTPTransport = userAgentTransport{recorder, c.HTTPTransport}
					return nil
				})
		})
	return recorder.snapshot()
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
	if !strings.Contains(req.URL.Path, "/scim/v2/ServicePrincipals") {
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
