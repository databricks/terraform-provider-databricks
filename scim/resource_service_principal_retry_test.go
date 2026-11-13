package scim

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// These tests pin the classic SDK client's retry, timeout and cancellation behavior for
// service principal reads, which replay fixtures cannot express.

const retryTestServicePrincipalPath = "/api/2.0/preview/scim/v2/ServicePrincipals/abc"

const retryTestServicePrincipalResource = retryTestServicePrincipalPath +
	"?attributes=" + userAttributes

const retryTestServicePrincipalJSON = `{"id": "abc", "applicationId": "bcd",
	"displayName": "Example Service Principal", "active": true}`

const retryTestTooManyRequestsJSON = `{
	"schemas": ["urn:ietf:params:scim:api:messages:2.0:Error"],
	"detail": "Too many requests.", "status": "429"}`

// attemptTransport counts the SDK's attempts at the service principal path. It opens a
// new connection for every request, so connection errors reach the SDK instead of being
// retried by net/http on a reused connection.
type attemptTransport struct {
	attempts atomic.Int32
	base     http.RoundTripper
	// send, if set, sends an attempt, numbered from 1, in place of base.
	send func(attempt int32, r *http.Request) (*http.Response, error)
}

func newAttemptTransport() *attemptTransport {
	return &attemptTransport{base: &http.Transport{DisableKeepAlives: true}}
}

func (a *attemptTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	if r.URL.Path != retryTestServicePrincipalPath {
		return a.base.RoundTrip(r)
	}
	attempt := a.attempts.Add(1)
	if a.send == nil {
		return a.base.RoundTrip(r)
	}
	return a.send(attempt, r)
}

// retryTestLoader keeps environment variables and ~/.databrickscfg out of the client.
type retryTestLoader struct{}

func (retryTestLoader) Name() string { return "retry test configuration" }

func (retryTestLoader) Configure(*config.Config) error { return nil }

func newRetryTestClient(t *testing.T, cfg *config.Config) *common.DatabricksClient {
	cfg.Token, cfg.AuthType, cfg.Credentials = "dapi123", "pat", config.PatCredentials{}
	cfg.Loaders = []config.Loader{retryTestLoader{}}
	inner, err := client.New(cfg)
	require.NoError(t, err)
	c := &common.DatabricksClient{DatabricksClient: inner}
	// Reads validate the workspace ID, which would otherwise be looked up through /Me.
	c.SetCachedWorkspaceID(12345)
	return c
}

// readServicePrincipal refreshes service principal abc through the resource, as
// Terraform does.
func readServicePrincipal(ctx context.Context,
	c *common.DatabricksClient) (*schema.ResourceData, diag.Diagnostics) {
	r := ResourceServicePrincipal().ToResource()
	d := r.TestResourceData()
	d.SetId("abc")
	ctx = context.WithValue(ctx, common.ResourceName, "service_principal")
	return d, r.ReadContext(ctx, d, c)
}

func requireSingleDiagnostic(t *testing.T, diags diag.Diagnostics) string {
	require.Len(t, diags, 1)
	return diags[0].Summary
}

// newPersistent429Server answers every service principal read with 429 and returns the
// server's URL.
func newPersistent429Server(t *testing.T) string {
	_, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:       "GET",
			Resource:     retryTestServicePrincipalResource,
			Status:       http.StatusTooManyRequests,
			ReuseRequest: true,
			Response:     retryTestTooManyRequestsJSON,
		},
	})
	require.NoError(t, err)
	t.Cleanup(server.Close)
	return server.URL
}

// newServicePrincipalServer serves the service principal read with handle and answers
// every other path, such as host metadata, with 404.
func newServicePrincipalServer(t *testing.T, handle http.HandlerFunc) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != retryTestServicePrincipalPath {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			handle(w, r)
		}))
	t.Cleanup(server.Close)
	return server
}

func TestServicePrincipalReadFailsWithLast429AtRetryTimeout(t *testing.T) {
	transport := newAttemptTransport()
	// The SDK backs off 1s and then 2s, each plus 50 to 750ms of jitter, so a 3s retry
	// timeout allows exactly two attempts.
	c := newRetryTestClient(t, &config.Config{
		Host:                newPersistent429Server(t),
		RetryTimeoutSeconds: 3,
		HTTPTransport:       transport,
	})

	d, diags := readServicePrincipal(context.Background(), c)

	// The SDK returns the last 429 rather than a timeout error, and the resource keeps
	// its ID because the error is not treated as missing.
	assert.Equal(t, "cannot read service principal: Too many requests.",
		requireSingleDiagnostic(t, diags))
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, int32(2), transport.attempts.Load())
}

// This test uses httptest because qa.HTTPFixture cannot reset a connection.
func TestServicePrincipalReadRetriesConnectionReset(t *testing.T) {
	var served atomic.Int32
	server := newServicePrincipalServer(t, func(w http.ResponseWriter, r *http.Request) {
		if served.Add(1) == 1 {
			resetConnection(t, w)
			return
		}
		_, err := w.Write([]byte(retryTestServicePrincipalJSON))
		assert.NoError(t, err)
	})
	transport := newAttemptTransport()
	c := newRetryTestClient(t, &config.Config{Host: server.URL, HTTPTransport: transport})

	d, diags := readServicePrincipal(context.Background(), c)

	require.Empty(t, diags)
	assert.Equal(t, "bcd", d.Get("application_id"))
	assert.Equal(t, int32(2), transport.attempts.Load())
	assert.Equal(t, int32(2), served.Load())
}

// resetConnection closes the client's connection with a TCP RST, so the client reads
// "connection reset by peer".
func resetConnection(t *testing.T, w http.ResponseWriter) {
	conn, _, err := http.NewResponseController(w).Hijack()
	if !assert.NoError(t, err) {
		return
	}
	assert.NoError(t, conn.(*net.TCPConn).SetLinger(0))
	assert.NoError(t, conn.Close())
}

// This test uses httptest because qa.HTTPFixture always accepts connections.
func TestServicePrincipalReadRetriesRefusedConnection(t *testing.T) {
	var served atomic.Int32
	server := newServicePrincipalServer(t, func(w http.ResponseWriter, r *http.Request) {
		served.Add(1)
		_, err := w.Write([]byte(retryTestServicePrincipalJSON))
		assert.NoError(t, err)
	})
	refusedAddress := closedListenerAddress(t)
	transport := newAttemptTransport()
	transport.send = func(attempt int32, r *http.Request) (*http.Response, error) {
		if attempt == 1 {
			r = r.Clone(r.Context())
			r.URL.Host = refusedAddress
		}
		return transport.base.RoundTrip(r)
	}
	c := newRetryTestClient(t, &config.Config{Host: server.URL, HTTPTransport: transport})

	d, diags := readServicePrincipal(context.Background(), c)

	require.Empty(t, diags)
	assert.Equal(t, "bcd", d.Get("application_id"))
	assert.Equal(t, int32(2), transport.attempts.Load())
	assert.Equal(t, int32(1), served.Load())
}

// closedListenerAddress returns a local address that refuses connections.
func closedListenerAddress(t *testing.T) string {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	address := listener.Addr().String()
	require.NoError(t, listener.Close())
	return address
}

// This test uses httptest because qa.HTTPFixture cannot delay a response.
func TestServicePrincipalReadDoesNotRetryHTTPTimeout(t *testing.T) {
	server := newServicePrincipalServer(t, func(w http.ResponseWriter, r *http.Request) {
		// Hold the response until the client gives up on it.
		<-r.Context().Done()
	})
	transport := newAttemptTransport()
	// The retry timeout would allow a second attempt if the HTTP timeout were retried.
	c := newRetryTestClient(t, &config.Config{
		Host:                server.URL,
		HTTPTimeoutSeconds:  1,
		RetryTimeoutSeconds: 3,
		HTTPTransport:       transport,
	})

	d, diags := readServicePrincipal(context.Background(), c)

	timeout := fmt.Sprintf(`Get "%s%s": request timed out after 1s of inactivity`,
		server.URL, retryTestServicePrincipalResource)
	assert.Equal(t, "cannot read service principal: "+timeout,
		requireSingleDiagnostic(t, diags))
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, int32(1), transport.attempts.Load())
}

func TestServicePrincipalReadStopsRetryingWhenCancelledDuringBackoff(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	transport := newAttemptTransport()
	// The context is cancelled once the first 429 has been read, so it is done when the
	// backoff before the second attempt starts. The body is read first so the cancel
	// cannot interrupt the response itself.
	transport.send = func(_ int32, r *http.Request) (*http.Response, error) {
		resp, err := transport.base.RoundTrip(r)
		if err != nil {
			return nil, err
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		resp.Body = io.NopCloser(bytes.NewReader(body))
		cancel()
		return resp, err
	}
	c := newRetryTestClient(t, &config.Config{
		Host:          newPersistent429Server(t),
		HTTPTransport: transport,
	})

	_, diags := readServicePrincipal(ctx, c)

	// The SDK returns the 429 that preceded the backoff, not a cancellation error.
	assert.Equal(t, "cannot read service principal: Too many requests.",
		requireSingleDiagnostic(t, diags))
	assert.Equal(t, int32(1), transport.attempts.Load())
}
