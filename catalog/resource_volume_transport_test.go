package catalog

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
	providerclient "github.com/databricks/terraform-provider-databricks/internal/providers/client"
)

// These tests pin how the classic SDK client's HTTP layer treats databricks_volume calls:
// the inactivity timeout while a body moves, in-flight cancellation and deadlines, the
// retry timeout, the rate limiter, TLS verification and environment proxies. Replay
// journeys can't express them, because each needs control of timing, of client settings
// that replay fixes, or of the network path. Tests on the real clock use httptest; tests
// that must measure time run on the synctest fake clock with an in-memory transport.

const (
	transportVolumeID  = "main.replay.routing_volume"
	transportVolumeGet = "GET " + routingVolumePath + "?"
)

var transportVolumeJSON = routingVolumeJSON("replay-admin@example.com", "")

// transportPriorState is routing_volume's state after an earlier apply and refresh.
func transportPriorState() *terraform.InstanceState {
	return &terraform.InstanceState{ID: transportVolumeID, Attributes: map[string]string{
		"id": transportVolumeID, "name": "routing_volume", "catalog_name": "main",
		"schema_name": "replay", "volume_type": "MANAGED", "comment": "",
		"owner":             "replay-admin@example.com",
		"volume_path":       "/Volumes/main/replay/routing_volume",
		"storage_location":  "s3://replay-metastore/volume-1920",
		"provider_config.#": "1", "provider_config.0.workspace_id": "12345",
	}}
}

// newTransportClient configures cfg as the provider does, without the environment,
// profiles or host metadata, and caches the workspace ID so no call resolves it.
func newTransportClient(t *testing.T, cfg *config.Config) *common.DatabricksClient {
	t.Helper()
	cfg.Token = "fixture-token"
	cfg.AuthType = "pat"
	cfg.Credentials = config.PatCredentials{}
	cfg.Loaders = []config.Loader{workspaceHostLoader{}}
	cfg.HostMetadataResolver = func(context.Context, string) (
		*config.HostMetadata, error) {
		return nil, nil
	}
	c, err := providerclient.PrepareDatabricksClient(context.Background(), cfg, nil)
	require.NoError(t, err)
	c.SetCachedWorkspaceID(12345)
	return c
}

// transportRefresh refreshes routing_volume from its prior state, as terraform plan does.
func transportRefresh(ctx context.Context, c *common.DatabricksClient) (
	*terraform.InstanceState, diag.Diagnostics) {
	ctx = context.WithValue(ctx, common.ResourceName, "volume")
	r := ResourceVolume().ToResource()
	return r.RefreshWithoutUpgrade(ctx, transportPriorState(), c)
}

// assertNoNextPlan asserts that the next plan of v against state has no changes.
func assertNoNextPlan(t *testing.T, c *common.DatabricksClient,
	state *terraform.InstanceState, v routingVolume) {
	t.Helper()
	r := ResourceVolume().ToResource()
	planned := state.DeepCopy()
	planned.RawConfig = v.rawConfig(t, r)
	ctx := context.WithValue(context.Background(), common.ResourceName, "volume")
	diff, err := r.Diff(ctx, planned, v.resourceConfig(), c)
	require.NoError(t, err)
	assert.Nil(t, diff, "the next plan must have no changes")
}

// transportVolumeConfig is routing_volume's configuration without optional fields.
var transportVolumeConfig = routingVolume{attributes: routingVolumeAttributes(nil)}

// newTransportServer creates an unstarted httptest server that logs each call as its
// method and request URI and then passes it to handler. Its cleanup closes the client
// connections first, which ends the request context of any handler still blocked on it.
func newTransportServer(t *testing.T, handler func(http.ResponseWriter, *http.Request),
) (*httptest.Server, *workspaceHostLog) {
	t.Helper()
	calls := &workspaceHostLog{}
	server := httptest.NewUnstartedServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			calls.add(r.Method + " " + r.RequestURI)
			handler(w, r)
		}))
	server.Config.ErrorLog = log.New(io.Discard, "", 0)
	t.Cleanup(func() {
		server.CloseClientConnections()
		server.Close()
	})
	return server, calls
}

func writeJSON(w http.ResponseWriter, status int, body string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = io.WriteString(w, body)
}

// writeVolumeInChunks sends the headers, then the volume JSON in eight chunks 200 ms
// apart, so the body takes 1.6 s while it keeps moving.
func writeVolumeInChunks(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	flusher := w.(http.Flusher)
	flusher.Flush()
	body, size := transportVolumeJSON, len(transportVolumeJSON)/8+1
	for body != "" {
		time.Sleep(200 * time.Millisecond)
		n := min(size, len(body))
		_, _ = io.WriteString(w, body[:n])
		flusher.Flush()
		body = body[n:]
	}
}

// TestVolumeRefreshBodyStreamedPastTimeouts pins that a response body that takes longer
// than the HTTP timeout, or than the retry timeout, but keeps moving succeeds: each body
// read extends the inactivity deadline, and the retry timeout only bounds the waits
// between attempts, never an attempt in flight.
func TestVolumeRefreshBodyStreamedPastTimeouts(t *testing.T) {
	tests := []struct {
		name                      string
		httpTimeout, retryTimeout int
	}{
		{"http_timeout_1s", 1, 0},
		{"retry_timeout_1s", 0, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			server, calls := newTransportServer(t, func(w http.ResponseWriter,
				r *http.Request) {
				writeVolumeInChunks(w)
			})
			server.Start()
			c := newTransportClient(t, &config.Config{Host: server.URL,
				HTTPTimeoutSeconds: tt.httpTimeout, RetryTimeoutSeconds: tt.retryTimeout})

			start := time.Now()
			state, diags := transportRefresh(context.Background(), c)

			require.False(t, diags.HasError(), "%v", diags)
			assert.Greater(t, time.Since(start), 1500*time.Millisecond)
			assert.Equal(t, []string{transportVolumeGet}, calls.snapshot())
			assert.Equal(t, transportPriorState().Attributes, state.Attributes)
			assertNoNextPlan(t, c, state, transportVolumeConfig)
		})
	}
}

// Socket buffers small enough that a slow reader holds back the client's body writes,
// instead of the kernels buffering the whole body. A receive buffer below the loopback
// MSS of 64 KiB would stall TCP's window updates.
const (
	transportReceiveBuffer = 64 << 10
	transportSendBuffer    = 16 << 10
)

type smallReadBufferListener struct{ net.Listener }

func (l smallReadBufferListener) Accept() (net.Conn, error) {
	conn, err := l.Listener.Accept()
	if tcp, ok := conn.(*net.TCPConn); ok {
		_ = tcp.SetReadBuffer(transportReceiveBuffer)
	}
	return conn, err
}

// readSlowly reads r in pieces of at most 16 KiB, paced to 640 KiB/s, so a 1 MiB body
// takes 1.6 s with pauses of at most 25 ms.
func readSlowly(r io.Reader) string {
	const bytesPerSecond = 640 << 10
	var body strings.Builder
	piece := make([]byte, 16<<10)
	start := time.Now()
	for {
		n, err := r.Read(piece)
		body.Write(piece[:n])
		if err != nil {
			return body.String()
		}
		time.Sleep(time.Until(start.Add(
			time.Duration(body.Len()) * time.Second / bytesPerSecond)))
	}
}

// TestVolumeCreateRequestBodyReadSlowlyPastHTTPTimeout pins that a create whose 1 MiB
// request body the server reads for longer than the HTTP timeout, with steady progress,
// succeeds: each body read extends the inactivity deadline. Small socket buffers on both
// ends make the client's body reads follow the server's pace. The SDK can't set its own
// transport's dialer, so the test passes a transport with one that shrinks the buffer.
func TestVolumeCreateRequestBodyReadSlowlyPastHTTPTimeout(t *testing.T) {
	comment := strings.Repeat("c", 1<<20)
	volumeJSON := routingVolumeJSON("replay-admin@example.com", comment)
	type received struct {
		body string
		took time.Duration
	}
	posts := make(chan received, 1)
	server, calls := newTransportServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			start := time.Now()
			posts <- received{readSlowly(r.Body), time.Since(start)}
		}
		writeJSON(w, http.StatusOK, volumeJSON)
	})
	server.Listener = smallReadBufferListener{server.Listener}
	server.Start()
	dialer := &net.Dialer{}
	transport := &http.Transport{DialContext: func(ctx context.Context, network,
		address string) (net.Conn, error) {
		conn, err := dialer.DialContext(ctx, network, address)
		if tcp, ok := conn.(*net.TCPConn); ok {
			_ = tcp.SetWriteBuffer(transportSendBuffer)
		}
		return conn, err
	}}
	defer transport.CloseIdleConnections()
	c := newTransportClient(t, &config.Config{Host: server.URL, HTTPTimeoutSeconds: 1,
		HTTPTransport: transport})
	v := routingVolume{attributes: routingVolumeAttributes(
		map[string]string{"comment": comment})}

	state, diags := routingApplyVolume(t, c, nil, v)

	require.False(t, diags.HasError(), "%v", diags)
	post := <-posts
	assert.Greater(t, post.took, 1500*time.Millisecond)
	wantBody := fmt.Sprintf(`{"catalog_name":"main","comment":%q,`+
		`"name":"routing_volume","schema_name":"replay",`+
		`"volume_type":"MANAGED"}`, comment)
	// A 1 MiB mismatch would flood the output, so only the length is reported.
	assert.True(t, wantBody == post.body, "unexpected POST body of %d bytes",
		len(post.body))
	assert.Equal(t, []string{"POST " + workspaceHostVolumesPath, transportVolumeGet},
		calls.snapshot())
	assert.Equal(t, transportVolumeID, state.ID)
	assert.True(t, comment == state.Attributes["comment"], "the comment must be in state")
	assertNoNextPlan(t, c, state, v)
}

// TestVolumeRefreshBodyStalledAfterHeadersFails pins a GET whose response stalls after
// the headers and half the body. After 1 s without progress the client cancels the body
// read. That surfaces as a bare context cancellation, not as the inactivity timeout a
// stall before the headers reports, and the refresh isn't retried, though a retry would
// succeed.
func TestVolumeRefreshBodyStalledAfterHeadersFails(t *testing.T) {
	var gets atomic.Int32
	server, calls := newTransportServer(t, func(w http.ResponseWriter,
		r *http.Request) {
		if gets.Add(1) > 1 {
			writeJSON(w, http.StatusOK, transportVolumeJSON)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, transportVolumeJSON[:len(transportVolumeJSON)/2])
		w.(http.Flusher).Flush()
		<-r.Context().Done()
	})
	server.Start()
	c := newTransportClient(t, &config.Config{Host: server.URL, HTTPTimeoutSeconds: 1})

	state, diags := transportRefresh(context.Background(), c)

	assert.Equal(t, "cannot read volume: response body: context canceled",
		diagnosticSummary(t, diags))
	assert.Equal(t, []string{transportVolumeGet}, calls.snapshot())
	assert.Equal(t, transportPriorState().Attributes, state.Attributes)
	assertNoNextPlan(t, c, state, transportVolumeConfig)
}

// TestVolumeRefreshEndedInFlight pins a refresh whose context ends while its GET waits
// for the response headers. Neither ending is reported as the inactivity timeout, and
// neither is retried, though a retry would succeed.
func TestVolumeRefreshEndedInFlight(t *testing.T) {
	tests := []struct {
		name string
		// end returns the refresh context, which ends once the GET has arrived.
		end  func(arrived <-chan struct{}) (context.Context, context.CancelFunc)
		want string
	}{
		{
			name: "cancelled",
			end: func(arrived <-chan struct{}) (context.Context, context.CancelFunc) {
				ctx, cancel := context.WithCancel(context.Background())
				go func() {
					<-arrived
					cancel()
				}()
				return ctx, cancel
			},
			want: "context canceled",
		},
		{
			name: "parent_deadline",
			end: func(<-chan struct{}) (context.Context, context.CancelFunc) {
				return context.WithTimeout(context.Background(), 500*time.Millisecond)
			},
			want: "context deadline exceeded",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var gets atomic.Int32
			arrived := make(chan struct{})
			server, calls := newTransportServer(t, func(w http.ResponseWriter,
				r *http.Request) {
				if gets.Add(1) > 1 {
					writeJSON(w, http.StatusOK, transportVolumeJSON)
					return
				}
				close(arrived)
				<-r.Context().Done()
			})
			server.Start()
			c := newTransportClient(t, &config.Config{Host: server.URL})
			ctx, cancel := tt.end(arrived)
			defer cancel()

			state, diags := transportRefresh(ctx, c)

			assert.Equal(t, fmt.Sprintf(`cannot read volume: Get "%s%s?": %s`, server.URL,
				routingVolumePath, tt.want), diagnosticSummary(t, diags))
			assert.Equal(t, []string{transportVolumeGet}, calls.snapshot())
			assert.Equal(t, transportPriorState().Attributes, state.Attributes)
			assertNoNextPlan(t, c, state, transportVolumeConfig)
		})
	}
}

// transportCall is one call to a transportVolumeAPI: its method and request URI, its
// body, and when it was sent and its context deadline as offsets from the API's
// creation. A call without a deadline has 0.
type transportCall struct {
	route, body    string
	sent, deadline time.Duration
}

// transportVolumeAPI is an in-memory volume API for synctest bubbles, which an httptest
// server can't serve. It answers each call with respond, numbering a route's calls from
// 1, and records it.
type transportVolumeAPI struct {
	start   time.Time
	respond func(route string, attempt int) (status int, body string)
	mu      sync.Mutex
	calls   []transportCall
}

func newTransportVolumeAPI(respond func(string, int) (int, string)) *transportVolumeAPI {
	return &transportVolumeAPI{start: time.Now(), respond: respond}
}

func (api *transportVolumeAPI) RoundTrip(req *http.Request) (*http.Response, error) {
	call := transportCall{route: req.Method + " " + req.URL.RequestURI(),
		sent: time.Since(api.start)}
	if req.Body != nil {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		call.body = string(body)
	}
	if deadline, ok := req.Context().Deadline(); ok {
		call.deadline = deadline.Sub(api.start)
	}
	api.mu.Lock()
	attempt := 1
	for _, earlier := range api.calls {
		if earlier.route == call.route {
			attempt++
		}
	}
	api.calls = append(api.calls, call)
	api.mu.Unlock()
	status, body := api.respond(call.route, attempt)
	recorder := httptest.NewRecorder()
	writeJSON(recorder, status, body)
	resp := recorder.Result()
	resp.Request = req
	return resp, nil
}

func (api *transportVolumeAPI) snapshot() []transportCall {
	api.mu.Lock()
	defer api.mu.Unlock()
	return append([]transportCall(nil), api.calls...)
}

// requests lists each call's method, request URI and body.
func (api *transportVolumeAPI) requests() []string {
	var requests []string
	for _, call := range api.snapshot() {
		requests = append(requests, strings.TrimSpace(call.route+" "+call.body))
	}
	return requests
}

// TestVolumeRefreshWithoutRetryTimeoutRetriesUntilReadTimeout pins the effective retry
// bound when retry_timeout_seconds is unset. The provider sets the SDK's retry timeout
// to -1, retry until the context ends, so a GET that keeps getting 429 is retried until
// the Terraform SDK's default 20 minute read timeout, and the refresh fails with the
// last error. The volume resource declares no timeouts, so no setting changes the bound.
func TestVolumeRefreshWithoutRetryTimeoutRetriesUntilReadTimeout(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		api := newTransportVolumeAPI(func(string, int) (int, string) {
			return http.StatusTooManyRequests, tooManyRequestsJSON
		})
		c := newTransportClient(t, &config.Config{Host: "https://volume-api.test",
			HTTPTransport: api})
		require.Equal(t, -1, c.Config.RetryTimeoutSeconds)

		state, diags := transportRefresh(t.Context(), c)

		assert.Equal(t, tooManyRequestsDiagnostic, diagnosticSummary(t, diags))
		assert.Equal(t, 20*time.Minute, time.Since(api.start))
		calls := api.snapshot()
		// Backoff waits grow by 1 s per attempt to a 10 s cap, plus 50 to 750 ms of
		// jitter, so 20 minutes fit 116 to 124 attempts.
		assert.GreaterOrEqual(t, len(calls), 116)
		assert.LessOrEqual(t, len(calls), 124)
		for _, call := range calls {
			assert.Equal(t, transportCall{transportVolumeGet, "", call.sent,
				20 * time.Minute}, call)
		}
		assert.Equal(t, transportPriorState().Attributes, state.Attributes)
		assertNoNextPlan(t, c, state, transportVolumeConfig)
	})
}

// TestVolumeCreateCallsShareOneOperationDeadline pins that the POST, the owner PATCH and
// the GET of one create run under the Terraform SDK's single 20 minute create timeout:
// time the POST spends on retries is taken from the PATCH and the GET.
func TestVolumeCreateCallsShareOneOperationDeadline(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		post := "POST " + workspaceHostVolumesPath
		api := newTransportVolumeAPI(func(route string, attempt int) (int, string) {
			if route == post && attempt < 3 {
				return http.StatusTooManyRequests, tooManyRequestsJSON
			}
			return http.StatusOK, routingVolumeJSON("data-engineers", "")
		})
		c := newTransportClient(t, &config.Config{Host: "https://volume-api.test",
			HTTPTransport: api})
		v := routingVolume{attributes: routingVolumeAttributes(
			map[string]string{"owner": "data-engineers"})}

		state, diags := routingApplyVolume(t, c, nil, v)

		require.False(t, diags.HasError(), "%v", diags)
		postBody := post + ` {"catalog_name":"main","name":"routing_volume",` +
			`"schema_name":"replay","volume_type":"MANAGED"}`
		assert.Equal(t, []string{postBody, postBody, postBody,
			"PATCH " + routingVolumePath + ` {"owner":"data-engineers"}`,
			transportVolumeGet}, api.requests())
		calls := api.snapshot()
		for _, call := range calls {
			assert.Equal(t, 20*time.Minute, call.deadline, call.route)
		}
		// Two backoffs of at least 1.05 s and 2.05 s precede the third POST.
		assert.GreaterOrEqual(t, calls[4].sent, 3100*time.Millisecond)
		assert.Equal(t, transportVolumeID, state.ID)
		assert.Equal(t, "data-engineers", state.Attributes["owner"])
		assertNoNextPlan(t, c, state, v)
	})
}

// TestVolumeRefreshRefusedByRateLimiter pins two refreshes that fail in the rate limiter
// before sending anything: one whose context was cancelled before the call, and one with
// a negative rate_limit, which the provider accepts and the limiter then refuses.
func TestVolumeRefreshRefusedByRateLimiter(t *testing.T) {
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	tests := []struct {
		name      string
		ctx       context.Context
		rateLimit int
		want      string
	}{
		{"cancelled_context", cancelled, 0, "context canceled"},
		{"negative_rate_limit", context.Background(), -1,
			"rate: Wait(n=1) exceeds limiter's burst -1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := newTransportVolumeAPI(func(string, int) (int, string) {
				return http.StatusOK, transportVolumeJSON
			})
			c := newTransportClient(t, &config.Config{Host: "https://volume-api.test",
				HTTPTransport: api, RateLimitPerSecond: tt.rateLimit})

			state, diags := transportRefresh(tt.ctx, c)

			assert.Equal(t, "cannot read volume: failed in rate limiter: "+tt.want,
				diagnosticSummary(t, diags))
			assert.Empty(t, api.snapshot())
			assert.Equal(t, transportPriorState().Attributes, state.Attributes)
			assertNoNextPlan(t, c, state, transportVolumeConfig)
		})
	}
}

// TestVolumeRefreshRateLimit pins the rate limiter's burst and refill on the synctest
// clock. rate_limit sets both the burst and the refill rate per second, and unset it
// defaults to 15. Refreshes on one provider client share its workspace client's limiter.
func TestVolumeRefreshRateLimit(t *testing.T) {
	tests := []struct {
		name      string
		rateLimit int
		// sent is when each consecutive refresh's GET is sent.
		sent []time.Duration
	}{
		{"rate_limit_2", 2, []time.Duration{0, 0, 500 * time.Millisecond, time.Second,
			1500 * time.Millisecond}},
		{"unset_defaults_to_15", 0, append(make([]time.Duration, 15),
			time.Second/15, 2*time.Second/15)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				api := newTransportVolumeAPI(func(string, int) (int, string) {
					return http.StatusOK, transportVolumeJSON
				})
				c := newTransportClient(t, &config.Config{Host: "https://volume-api.test",
					HTTPTransport: api, RateLimitPerSecond: tt.rateLimit})

				for range tt.sent {
					_, diags := transportRefresh(t.Context(), c)
					require.False(t, diags.HasError(), "%v", diags)
				}

				var sent []time.Duration
				for _, call := range api.snapshot() {
					require.Equal(t, transportVolumeGet, call.route)
					sent = append(sent, call.sent)
				}
				assert.Equal(t, tt.sent, sent)
			})
		})
	}
}

// TestVolumeRefreshOverHTTPS pins TLS verification on the SDK's own transport: a server
// certificate from an untrusted issuer fails the refresh at the handshake, without a
// retry, and skip_verify accepts it.
func TestVolumeRefreshOverHTTPS(t *testing.T) {
	tests := []struct {
		name       string
		skipVerify bool
		// want is the refresh error after the URL, or empty for success.
		want  string
		calls []string
	}{
		{"untrusted_certificate", false, "tls: failed to verify certificate: " +
			"x509: certificate signed by unknown authority", nil},
		{"skip_verify", true, "", []string{transportVolumeGet}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, calls := newTransportServer(t, func(w http.ResponseWriter,
				r *http.Request) {
				writeJSON(w, http.StatusOK, transportVolumeJSON)
			})
			var handshakes atomic.Int32
			server.TLS = &tls.Config{GetConfigForClient: func(*tls.ClientHelloInfo) (
				*tls.Config, error) {
				handshakes.Add(1)
				return nil, nil
			}}
			server.StartTLS()
			c := newTransportClient(t, &config.Config{Host: server.URL,
				InsecureSkipVerify: tt.skipVerify})

			state, diags := transportRefresh(context.Background(), c)

			if tt.want == "" {
				require.False(t, diags.HasError(), "%v", diags)
			} else {
				assert.Equal(t, fmt.Sprintf(`cannot read volume: Get "%s%s?": %s`,
					server.URL, routingVolumePath, tt.want), diagnosticSummary(t, diags))
			}
			assert.Equal(t, int32(1), handshakes.Load())
			assert.Equal(t, tt.calls, calls.snapshot())
			assert.Equal(t, transportPriorState().Attributes, state.Attributes)
			assertNoNextPlan(t, c, state, transportVolumeConfig)
		})
	}
}

// transportProxyChildHost names the volume host when this test binary runs as the
// proxy test's child.
const transportProxyChildHost = "VOLUME_TRANSPORT_PROXY_CHILD_HOST"

// TestVolumeRefreshThroughEnvironmentProxy pins that volume calls go through the proxy
// that HTTP_PROXY or HTTPS_PROXY names. net/http reads those variables once per process,
// so each case refreshes in a child run of this test binary with a fresh environment.
// The child refreshes routing_volume on volume-api.test, which only the proxy reaches.
// An HTTPS call tunnels through CONNECT to a TLS server, with skip_verify set because
// the server's certificate isn't trusted.
func TestVolumeRefreshThroughEnvironmentProxy(t *testing.T) {
	if host := os.Getenv(transportProxyChildHost); host != "" {
		refreshThroughEnvironmentProxy(t, host)
		return
	}
	tests := []struct {
		name, variable, host    string
		proxyCalls, originCalls []string
	}{
		{"http_proxy", "HTTP_PROXY", "http://volume-api.test",
			[]string{"GET http://volume-api.test" + routingVolumePath + "?"}, nil},
		{"https_proxy", "HTTPS_PROXY", "https://volume-api.test",
			[]string{"CONNECT volume-api.test:443"}, []string{transportVolumeGet}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			origin, originCalls := newTransportServer(t, func(w http.ResponseWriter,
				r *http.Request) {
				writeJSON(w, http.StatusOK, transportVolumeJSON)
			})
			origin.StartTLS()
			proxy, proxyCalls := newTransportServer(t, func(w http.ResponseWriter,
				r *http.Request) {
				if r.Method == http.MethodConnect {
					tunnel(w, origin.Listener.Addr().String())
					return
				}
				writeJSON(w, http.StatusOK, transportVolumeJSON)
			})
			proxy.Start()

			child := exec.Command(os.Args[0],
				"-test.run=^TestVolumeRefreshThroughEnvironmentProxy$")
			child.Env = append(environmentWithoutProxies(),
				transportProxyChildHost+"="+tt.host, tt.variable+"="+proxy.URL)
			output, err := child.CombinedOutput()

			require.NoError(t, err, "%s", output)
			assert.Equal(t, tt.proxyCalls, proxyCalls.snapshot())
			assert.Equal(t, tt.originCalls, originCalls.snapshot())
		})
	}
}

// refreshThroughEnvironmentProxy is the proxy test's child: a successful refresh of
// routing_volume on host.
func refreshThroughEnvironmentProxy(t *testing.T, host string) {
	c := newTransportClient(t, &config.Config{Host: host,
		InsecureSkipVerify: strings.HasPrefix(host, "https://")})

	state, diags := transportRefresh(context.Background(), c)

	require.False(t, diags.HasError(), "%v", diags)
	assert.Equal(t, transportPriorState().Attributes, state.Attributes)
	assertNoNextPlan(t, c, state, transportVolumeConfig)
}

// environmentWithoutProxies is this process's environment without the variables that
// configure net/http's proxies.
func environmentWithoutProxies() []string {
	var env []string
	for _, variable := range os.Environ() {
		name, _, _ := strings.Cut(variable, "=")
		switch strings.ToUpper(name) {
		case "HTTP_PROXY", "HTTPS_PROXY", "NO_PROXY", "REQUEST_METHOD",
			transportProxyChildHost:
			continue
		}
		env = append(env, variable)
	}
	return env
}

// tunnel answers a CONNECT request by piping its connection to address.
func tunnel(w http.ResponseWriter, address string) {
	upstream, err := net.Dial("tcp", address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer upstream.Close()
	conn, buffered, err := http.NewResponseController(w).Hijack()
	if err != nil {
		return
	}
	defer conn.Close()
	_, _ = io.WriteString(conn, "HTTP/1.1 200 Connection established\r\n\r\n")
	go func() {
		_, _ = io.Copy(upstream, buffered)
		_ = upstream.Close()
	}()
	_, _ = io.Copy(conn, upstream)
}
