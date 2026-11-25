package httpfixture

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"testing"
)

// NewOfflineAuthTransport stubs remote host and OAuth discovery as unavailable.
// Local fixture servers remain reachable so authentication and token refresh
// tests can exercise their real HTTP flows. Other remote requests fail the test.
func NewOfflineAuthTransport(t testing.TB) http.RoundTripper {
	t.Helper()
	local := &http.Transport{}
	t.Cleanup(local.CloseIdleConnections)
	return offlineAuthTransport{t: t, local: local}
}

type offlineAuthTransport struct {
	t     testing.TB
	local http.RoundTripper
}

func (transport offlineAuthTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	if net.ParseIP(request.URL.Hostname()).IsLoopback() {
		return transport.local.RoundTrip(request)
	}
	if request.Method == http.MethodGet &&
		(request.URL.Path == "/.well-known/databricks-config" ||
			strings.HasSuffix(request.URL.Path, "/.well-known/oauth-authorization-server")) {
		return &http.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Header:     make(http.Header),
			Body:       http.NoBody,
			Request:    request,
		}, nil
	}
	err := fmt.Errorf("unexpected request outside local auth fixtures: %s %s", request.Method, request.URL)
	transport.t.Error(err)
	return nil, err
}
