package catalog

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
)

// These tests pin how a databricks_volume refresh is retried after connection errors that
// the replay fixture server can't produce, because each fails before the server could
// answer: a refused connection, a connect timeout, a TLS handshake timeout and a proxy
// that refuses the tunnel. Each error's message contains one of the classic SDK's
// transient URL error messages, so the SDK retries it until retry_timeout_seconds ends.

// connectionRetryTimeoutSeconds allows exactly two attempts: the retry starts 1.05 to
// 1.75 s after the first attempt, and a third could start no sooner than 3.1 s after it.
const connectionRetryTimeoutSeconds = 3

// connectionAttempt is a transport whose first attempts fail with one connection error
// and whose later attempts reach the server. attempts counts them, and failure is the
// refresh diagnostic for that error.
type connectionAttempt struct {
	transport *http.Transport
	attempts  *atomic.Int32
	failure   string
}

// connectionCalls records each call's method, request URI and Authorization header.
type connectionCalls struct {
	mu    sync.Mutex
	calls []string
}

func (c *connectionCalls) snapshot() []string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return append([]string(nil), c.calls...)
}

// newConnectionServer answers the volume GET and logs every call it receives.
func newConnectionServer(t *testing.T, useTLS bool) (*httptest.Server, *connectionCalls) {
	t.Helper()
	calls := &connectionCalls{}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.mu.Lock()
		calls.calls = append(calls.calls, fmt.Sprintf("%s %s Authorization=%s",
			r.Method, r.RequestURI, r.Header.Get("Authorization")))
		calls.mu.Unlock()
		if r.Method != http.MethodGet || r.URL.Path != retryVolumePath {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(retryVolumeJSON))
	})
	server := httptest.NewUnstartedServer(handler)
	if useTLS {
		server.StartTLS()
	} else {
		server.Start()
	}
	t.Cleanup(server.Close)
	return server, calls
}

type dialFunc func(ctx context.Context, network, addr string) (net.Conn, error)

// failingDialer counts dials and sends the first failures of them to fail, the rest to
// the dialed address.
func failingDialer(failures int32, fail dialFunc) (dialFunc, *atomic.Int32) {
	attempts := &atomic.Int32{}
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		if attempts.Add(1) <= failures {
			return fail(ctx, network, addr)
		}
		return (&net.Dialer{}).DialContext(ctx, network, addr)
	}, attempts
}

// closedAddress returns a loopback address that nothing listens on, so connecting to it
// is refused.
func closedAddress(t *testing.T) string {
	t.Helper()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	addr := listener.Addr().String()
	require.NoError(t, listener.Close())
	return addr
}

// silentListener accepts connections and never writes to them, so a TLS handshake over
// one waits until the client gives up. It closes them when the test ends.
func silentListener(t *testing.T) string {
	t.Helper()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	var conns sync.WaitGroup
	t.Cleanup(func() {
		_ = listener.Close()
		conns.Wait()
	})
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			conns.Add(1)
			go func() {
				defer conns.Done()
				_, _ = io.Copy(io.Discard, conn)
				_ = conn.Close()
			}()
		}
	}()
	return listener.Addr().String()
}

// refusingProxy answers the first failures CONNECT requests with a 502 whose reason
// phrase is "Unexpected error" and tunnels later ones to their target. It returns the
// proxy URL and the number of CONNECT requests so far.
func refusingProxy(t *testing.T, failures int32) (*url.URL, *atomic.Int32) {
	t.Helper()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	t.Cleanup(func() { _ = listener.Close() })
	connects := &atomic.Int32{}
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			go tunnel(conn, connects.Add(1) <= failures)
		}
	}()
	return &url.URL{Scheme: "http", Host: listener.Addr().String()}, connects
}

func tunnel(conn net.Conn, refuse bool) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	req, err := http.ReadRequest(reader)
	if err != nil || req.Method != http.MethodConnect {
		return
	}
	if refuse {
		_, _ = io.WriteString(conn, "HTTP/1.1 502 Unexpected error\r\n\r\n")
		return
	}
	target, err := net.Dial("tcp", req.Host)
	if err != nil {
		return
	}
	defer target.Close()
	_, _ = io.WriteString(conn, "HTTP/1.1 200 Connection established\r\n\r\n")
	go func() { _, _ = io.Copy(target, reader) }()
	_, _ = io.Copy(conn, target)
}

// tlsTransport trusts server's certificate.
func tlsTransport(server *httptest.Server) *http.Transport {
	trusted := server.Client().Transport.(*http.Transport).TLSClientConfig
	return &http.Transport{TLSClientConfig: trusted.Clone()}
}

// readFailure is the refresh diagnostic for a volume GET to server failing with err.
func readFailure(server *httptest.Server, err string) string {
	return fmt.Sprintf(`cannot read volume: Get "%s%s?": %s`,
		server.URL, retryVolumePath, err)
}

// connectionSetup makes a transport whose first failures attempts fail.
type connectionSetup func(t *testing.T, server *httptest.Server, failures int32,
) connectionAttempt

var connectionCases = map[string]struct {
	tls   bool
	setup connectionSetup
}{
	// A dial to a closed port stands in for a host that refuses the connection.
	"connection_refused": {
		setup: func(t *testing.T, server *httptest.Server, failures int32,
		) connectionAttempt {
			closed := closedAddress(t)
			dial, attempts := failingDialer(failures,
				func(ctx context.Context, network, _ string) (net.Conn, error) {
					return (&net.Dialer{}).DialContext(ctx, network, closed)
				})
			return connectionAttempt{&http.Transport{DialContext: dial}, attempts,
				readFailure(server, "dial tcp "+closed+": connect: connection refused")}
		},
	},
	// The SDK's dialer gives up after 30 s; a dialer whose deadline has passed fails
	// with the same i/o timeout at once.
	"connect_timeout": {
		setup: func(t *testing.T, server *httptest.Server, failures int32,
		) connectionAttempt {
			dial, attempts := failingDialer(failures,
				(&net.Dialer{Deadline: time.Unix(1, 0)}).DialContext)
			return connectionAttempt{&http.Transport{DialContext: dial}, attempts,
				readFailure(server, "dial tcp "+server.Listener.Addr().String()+
					": i/o timeout")}
		},
	},
	// The SDK's transport waits 30 s for a TLS handshake; this one waits 100 ms.
	"tls_handshake_timeout": {
		tls: true,
		setup: func(t *testing.T, server *httptest.Server, failures int32,
		) connectionAttempt {
			silent := silentListener(t)
			dial, attempts := failingDialer(failures,
				func(ctx context.Context, network, _ string) (net.Conn, error) {
					return (&net.Dialer{}).DialContext(ctx, network, silent)
				})
			transport := tlsTransport(server)
			transport.DialContext = dial
			transport.TLSHandshakeTimeout = 100 * time.Millisecond
			return connectionAttempt{transport, attempts,
				readFailure(server, "net/http: TLS handshake timeout")}
		},
	},
	// net/http reports a refused CONNECT by the reason phrase of the proxy's answer.
	"proxy_refuses_the_tunnel": {
		tls: true,
		setup: func(t *testing.T, server *httptest.Server, failures int32,
		) connectionAttempt {
			proxy, connects := refusingProxy(t, failures)
			transport := tlsTransport(server)
			transport.Proxy = http.ProxyURL(proxy)
			return connectionAttempt{transport, connects,
				readFailure(server, "Unexpected error")}
		},
	},
}

// newConnectionClient configures a client without the environment, profiles or host
// metadata, so every attempt is a volume call.
func newConnectionClient(t *testing.T, host string, transport http.RoundTripper,
) *common.DatabricksClient {
	t.Helper()
	return newVolumeClient(t, &config.Config{
		Host:                 host,
		RetryTimeoutSeconds:  connectionRetryTimeoutSeconds,
		HTTPTransport:        transport,
		Loaders:              []config.Loader{workspaceHostLoader{}},
		HostMetadataResolver: noHostMetadata,
	})
}

func noHostMetadata(context.Context, string) (*config.HostMetadata, error) {
	return nil, nil
}

// Replay can't refuse, time out or proxy a connection, so these use real sockets.
func TestVolumeRefreshRetriedOnceAfterConnectionError(t *testing.T) {
	for name, tc := range connectionCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			server, calls := newConnectionServer(t, tc.tls)
			attempt := tc.setup(t, server, 1)
			c := newConnectionClient(t, server.URL, attempt.transport)

			d, diags := refreshVolume(context.Background(), c)

			require.Empty(t, diags)
			assert.Equal(t, int32(2), attempt.attempts.Load())
			assert.Equal(t, []string{
				"GET " + retryVolumePath + "? Authorization=Bearer fixture-token",
			}, calls.snapshot())
			assert.Equal(t, retryVolumeID, d.Id())
			assert.Equal(t, "/Volumes/main/replay/retry_volume", d.Get("volume_path"))
		})
	}
}

// Replay can't refuse, time out or proxy a connection, so these use real sockets.
func TestVolumeRefreshConnectionErrorUntilRetryTimeout(t *testing.T) {
	for name, tc := range connectionCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			server, calls := newConnectionServer(t, tc.tls)
			attempt := tc.setup(t, server, math.MaxInt32)
			c := newConnectionClient(t, server.URL, attempt.transport)

			d, diags := refreshVolume(context.Background(), c)

			assert.Equal(t, attempt.failure, diagnosticSummary(t, diags))
			assert.Equal(t, int32(2), attempt.attempts.Load())
			assert.Empty(t, calls.snapshot())
			assert.Equal(t, retryVolumeID, d.Id(), "the volume must stay in state")
		})
	}
}
