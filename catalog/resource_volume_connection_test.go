package catalog

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
)

// These tests pin how the classic SDK client handles connections that fail on
// databricks_volume calls: which failures it retries, how often, and the error it ends
// with. Replay fixtures cannot break a connection, so these use real servers.
//
// The SDK retries a request whose error contains one of a few connection messages,
// until its retry timeout ends. Independently, net/http resends a GET once, without
// backoff, when it fails on a reused keep-alive connection. Here the first volume
// request reuses the connection of the host metadata request the client makes when it
// is built, so a refresh GET that fails is sent once more than the other calls.

// retryTwiceSeconds is a retry timeout that allows exactly two attempts: the SDK backs
// off 1.05 to 1.75 s before the second and at least 2.05 s before the third.
const retryTwiceSeconds = 3

var refreshVolumeGet = volumeRequest{http.MethodGet, retryVolumePath + "?", ""}

// volumeCall runs one of the provider's wrapped CRUD functions for the volume.
type volumeCall func(context.Context, *testing.T, *common.DatabricksClient) (
	*schema.ResourceData, diag.Diagnostics)

// refreshRetryVolume runs the provider's wrapped Read, as a Terraform refresh does.
func refreshRetryVolume(
	ctx context.Context, _ *testing.T, c *common.DatabricksClient,
) (*schema.ResourceData, diag.Diagnostics) {
	return refreshVolume(ctx, c)
}

// connectionResetError is the error of received request i, whose connection the
// server reset.
func connectionResetError(api *scriptedVolumeAPI, server *httptest.Server, i int) string {
	requests, clients := api.requests(), api.clientAddresses()
	if i >= len(requests) {
		return fmt.Sprintf("request %d was not received", i)
	}
	method := requests[i].method
	return fmt.Sprintf(`%s "%s%s": read tcp %s->%s: read: connection reset by peer`,
		method[:1]+strings.ToLower(method[1:]), server.URL, requests[i].uri, clients[i],
		server.Listener.Addr())
}

func TestVolumeCallsRetryConnectionReset(t *testing.T) {
	reset := volumeResponse{reset: true}
	created := volumeResponse{status: http.StatusOK, body: retryVolumeJSON}
	ownerChanged := volumeResponse{status: http.StatusOK, body: volumeWithNewOwnerJSON}
	deleted := volumeResponse{status: http.StatusOK, body: `{}`}
	tests := []struct {
		name string
		run  volumeCall
		// recovered answers the call's requests, and the read that follows a successful
		// Create or Update, when only the call's first attempt is reset.
		// recoveredRequests are the requests the server then receives.
		recovered         []volumeResponse
		recoveredRequests []volumeRequest
		// exhausted answers them when every attempt is reset.
		exhausted         []volumeResponse
		exhaustedRequests []volumeRequest
		// diagnostic starts the error the exhausted call ends with, which then has the
		// error of request lastReset.
		diagnostic string
		lastReset  int
		// id is the state's ID after the exhausted call.
		id string
	}{
		{
			// If the server created the volume before it reset the connection, the
			// retry gets RESOURCE_ALREADY_EXISTS, as
			// TestVolumeCreateRetriedAfterConnectionResetFailsWithAlreadyExists pins.
			name:      "create POST",
			run:       createVolumeWithOwner,
			recovered: []volumeResponse{reset, created, created, created},
			recoveredRequests: []volumeRequest{
				createVolumePost, createVolumePost, createOwnerPatch, refreshVolumeGet,
			},
			exhausted:         []volumeResponse{reset, reset},
			exhaustedRequests: []volumeRequest{createVolumePost, createVolumePost},
			diagnostic:        "cannot create volume: ",
			lastReset:         1,
			id:                "",
		},
		{
			name:      "owner PATCH after create",
			run:       createVolumeWithOwner,
			recovered: []volumeResponse{created, reset, created, created},
			recoveredRequests: []volumeRequest{
				createVolumePost, createOwnerPatch, createOwnerPatch, refreshVolumeGet,
			},
			exhausted: []volumeResponse{created, reset, reset},
			exhaustedRequests: []volumeRequest{
				createVolumePost, createOwnerPatch, createOwnerPatch,
			},
			diagnostic: "cannot create volume: ",
			lastReset:  2,
			id:         retryVolumeID,
		},
		{
			// net/http resends the GET on a new connection, so one reset costs no SDK
			// retry, and every reset costs three requests.
			name:              "refresh GET",
			run:               refreshRetryVolume,
			recovered:         []volumeResponse{reset, created},
			recoveredRequests: []volumeRequest{refreshVolumeGet, refreshVolumeGet},
			exhausted:         []volumeResponse{reset, reset, reset},
			exhaustedRequests: []volumeRequest{
				refreshVolumeGet, refreshVolumeGet, refreshVolumeGet,
			},
			diagnostic: "cannot read volume: ",
			lastReset:  2,
			id:         retryVolumeID,
		},
		{
			// The owner rollback after the exhausted fields PATCH succeeds, so the
			// error is the fields PATCH's.
			name:      "fields PATCH in Update",
			run:       updateVolumeOwnerAndComment,
			recovered: []volumeResponse{ownerChanged, reset, ownerChanged, created},
			recoveredRequests: []volumeRequest{
				updateOwnerPatch, updateFieldsPatch, updateFieldsPatch, refreshVolumeGet,
			},
			exhausted: []volumeResponse{ownerChanged, reset, reset, created},
			exhaustedRequests: []volumeRequest{
				updateOwnerPatch, updateFieldsPatch, updateFieldsPatch, ownerRollbackPatch,
			},
			diagnostic: "cannot update volume: ",
			lastReset:  2,
			id:         retryVolumeID,
		},
		{
			name:              "DELETE",
			run:               destroyVolume,
			recovered:         []volumeResponse{reset, deleted},
			recoveredRequests: []volumeRequest{deleteVolumeRequest, deleteVolumeRequest},
			exhausted:         []volumeResponse{reset, reset},
			exhaustedRequests: []volumeRequest{deleteVolumeRequest, deleteVolumeRequest},
			diagnostic:        "cannot delete volume: ",
			lastReset:         1,
			id:                retryVolumeID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+" recovers", func(t *testing.T) {
			t.Parallel()
			api, server := newScriptedVolumeServer(t, tt.recovered...)
			c := newVolumeClient(t, &config.Config{
				Host:                server.URL,
				RetryTimeoutSeconds: retryUntilContextEnds,
			})

			d, diags := tt.run(context.Background(), t, c)

			assertVolumeRequests(t, tt.recoveredRequests, api)
			assert.Empty(t, diags)
			assert.Equal(t, retryVolumeID, d.Id())
		})
		t.Run(tt.name+" exhausts the retry timeout", func(t *testing.T) {
			t.Parallel()
			api, server := newScriptedVolumeServer(t, tt.exhausted...)
			c := newVolumeClient(t, &config.Config{
				Host:                server.URL,
				RetryTimeoutSeconds: retryTwiceSeconds,
			})

			d, diags := tt.run(context.Background(), t, c)

			assertVolumeRequests(t, tt.exhaustedRequests, api)
			assert.Equal(t, tt.diagnostic+connectionResetError(api, server, tt.lastReset),
				diagnosticSummary(t, diags))
			assert.Equal(t, tt.id, d.Id())
		})
	}
}

// faultyDialTransport sends requests over connections of its own, and can replace the
// next dial. It records the error of each volume request attempt, which the SDK does
// not report when a retry succeeds.
type faultyDialTransport struct {
	transport *http.Transport
	// failed is closed when an attempt first fails.
	failed     chan struct{}
	failedOnce sync.Once
	mu         sync.Mutex
	fault      func(ctx context.Context, network, address string) (net.Conn, error)
	errs       []string
}

func newFaultyDialTransport(tlsConfig *tls.Config) *faultyDialTransport {
	ft := &faultyDialTransport{failed: make(chan struct{})}
	dialer := &net.Dialer{}
	ft.transport = &http.Transport{
		TLSClientConfig: tlsConfig,
		// A loopback handshake takes milliseconds, so only a stalled one times out.
		TLSHandshakeTimeout: time.Second,
		DialContext: func(ctx context.Context, network, address string) (net.Conn, error) {
			ft.mu.Lock()
			fault := ft.fault
			ft.fault = nil
			ft.mu.Unlock()
			if fault != nil {
				return fault(ctx, network, address)
			}
			return dialer.DialContext(ctx, network, address)
		},
	}
	return ft
}

// replaceNextDial drops idle connections, so that the next request dials, and makes
// fault that dial.
func (ft *faultyDialTransport) replaceNextDial(
	fault func(ctx context.Context, network, address string) (net.Conn, error),
) {
	ft.transport.CloseIdleConnections()
	ft.mu.Lock()
	defer ft.mu.Unlock()
	ft.fault = fault
}

func (ft *faultyDialTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := ft.transport.RoundTrip(req)
	if !strings.HasPrefix(req.URL.Path, volumesAPIPath) {
		return resp, err
	}
	ft.mu.Lock()
	defer ft.mu.Unlock()
	if err == nil {
		ft.errs = append(ft.errs, "")
		return resp, nil
	}
	ft.errs = append(ft.errs, err.Error())
	ft.failedOnce.Do(func() { close(ft.failed) })
	return resp, err
}

// attempts returns the error of each volume request attempt, empty if it got a
// response.
func (ft *faultyDialTransport) attempts() []string {
	ft.mu.Lock()
	defer ft.mu.Unlock()
	return append([]string(nil), ft.errs...)
}

// connectionRefusedError is the error of a dial to the address of server.
func connectionRefusedError(server *httptest.Server) string {
	return fmt.Sprintf("dial tcp %s: connect: connection refused", server.Listener.Addr())
}

func TestVolumeRefreshRetriesConnectionRefused(t *testing.T) {
	t.Run("recovers", func(t *testing.T) {
		t.Parallel()
		api, server := newScriptedVolumeServer(t,
			volumeResponse{status: http.StatusOK, body: retryVolumeJSON})
		transport := newFaultyDialTransport(nil)
		c := newVolumeClient(t, &config.Config{
			Host:                server.URL,
			RetryTimeoutSeconds: retryUntilContextEnds,
			HTTPTransport:       transport,
		})
		refused := connectionRefusedError(server)
		address := server.Listener.Addr().String()
		// Once the server and the idle connection are closed, the GET dials and is
		// refused.
		server.Close()
		transport.transport.CloseIdleConnections()
		var diags diag.Diagnostics
		done := make(chan struct{})
		go func() {
			defer close(done)
			_, diags = refreshVolume(context.Background(), c)
		}()

		select {
		case <-transport.failed:
		case <-done:
			t.Fatal("the refresh returned without a refused attempt")
		}
		// The SDK backs off at least 1.05 s, so the server is back before the retry.
		listener, err := net.Listen("tcp", address)
		require.NoError(t, err)
		reopened := httptest.NewUnstartedServer(api)
		reopened.Listener.Close()
		reopened.Listener = listener
		reopened.Start()
		defer reopened.Close()
		<-done

		assert.Equal(t, []string{refused, ""}, transport.attempts())
		assertVolumeRequests(t, []volumeRequest{refreshVolumeGet}, api)
		assert.Empty(t, diags)
	})
	t.Run("exhausts the retry timeout", func(t *testing.T) {
		t.Parallel()
		api, server := newScriptedVolumeServer(t)
		transport := newFaultyDialTransport(nil)
		c := newVolumeClient(t, &config.Config{
			Host:                server.URL,
			RetryTimeoutSeconds: retryTwiceSeconds,
			HTTPTransport:       transport,
		})
		refused := connectionRefusedError(server)
		server.Close()
		transport.transport.CloseIdleConnections()

		d, diags := refreshVolume(context.Background(), c)

		assert.Equal(t, []string{refused, refused}, transport.attempts())
		assertVolumeRequests(t, nil, api)
		assert.Equal(t, fmt.Sprintf(`cannot read volume: Get "%s%s?": %s`,
			server.URL, retryVolumePath, refused), diagnosticSummary(t, diags))
		assert.Equal(t, retryVolumeID, d.Id(), "the volume must stay in state")
	})
}

// silentListener accepts connections and never writes to them, so a TLS handshake over
// one stalls. It returns the listener's address.
func silentListener(t *testing.T) string {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	var mu sync.Mutex
	var accepted []net.Conn
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			mu.Lock()
			accepted = append(accepted, conn)
			mu.Unlock()
		}
	}()
	t.Cleanup(func() {
		_ = listener.Close()
		mu.Lock()
		defer mu.Unlock()
		for _, conn := range accepted {
			_ = conn.Close()
		}
	})
	return listener.Addr().String()
}

func TestVolumeRefreshRetriesOtherTransientConnectionErrors(t *testing.T) {
	tests := []struct {
		name string
		tls  bool
		// fault replaces the dial of the GET's first attempt.
		fault func(t *testing.T) func(context.Context, string, string) (net.Conn, error)
		// attemptError returns the error of that attempt for the server's address.
		attemptError func(address string) string
	}{
		{
			name: "i/o timeout",
			fault: func(*testing.T) func(context.Context, string, string) (net.Conn, error) {
				// A dial whose deadline has passed times out.
				expired := &net.Dialer{Deadline: time.Unix(1, 0)}
				return expired.DialContext
			},
			attemptError: func(address string) string {
				return "dial tcp " + address + ": i/o timeout"
			},
		},
		{
			name: "TLS handshake timeout",
			tls:  true,
			fault: func(t *testing.T) func(context.Context, string, string) (net.Conn, error) {
				silent := silentListener(t)
				return func(ctx context.Context, network, _ string) (net.Conn, error) {
					return (&net.Dialer{}).DialContext(ctx, network, silent)
				}
			},
			attemptError: func(string) string { return "net/http: TLS handshake timeout" },
		},
		{
			// No Go standard library package returns this message, so a dial
			// function stands in for the network stack or proxy that would.
			name: "Unexpected error",
			fault: func(*testing.T) func(context.Context, string, string) (net.Conn, error) {
				return func(context.Context, string, string) (net.Conn, error) {
					return nil, errors.New("Unexpected error")
				}
			},
			attemptError: func(string) string { return "Unexpected error" },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			api := newScriptedVolumeAPI(t,
				volumeResponse{status: http.StatusOK, body: retryVolumeJSON})
			server := httptest.NewUnstartedServer(api)
			var tlsConfig *tls.Config
			if tt.tls {
				server.StartTLS()
				tlsConfig = server.Client().Transport.(*http.Transport).TLSClientConfig
			} else {
				server.Start()
			}
			t.Cleanup(server.Close)
			transport := newFaultyDialTransport(tlsConfig)
			c := newVolumeClient(t, &config.Config{
				Host:                server.URL,
				RetryTimeoutSeconds: retryUntilContextEnds,
				HTTPTransport:       transport,
			})
			transport.replaceNextDial(tt.fault(t))

			d, diags := refreshVolume(context.Background(), c)

			assert.Equal(t, []string{tt.attemptError(server.Listener.Addr().String()), ""},
				transport.attempts())
			assertVolumeRequests(t, []volumeRequest{refreshVolumeGet}, api)
			assert.Empty(t, diags)
			assert.Equal(t, retryVolumeID, d.Id())
		})
	}
}

func TestVolumeCallsDoNotRetryConnectionClosedWithoutResponse(t *testing.T) {
	eof := volumeResponse{eof: true}
	tests := []struct {
		name     string
		run      volumeCall
		requests []volumeRequest
		// diagnostic returns the error for the server's URL.
		diagnostic func(serverURL string) string
		id         string
	}{
		{
			name:     "create POST",
			run:      createVolumeWithOwner,
			requests: []volumeRequest{createVolumePost},
			diagnostic: func(u string) string {
				return `cannot create volume: Post "` + u + volumesAPIPath + `": EOF`
			},
			id: "",
		},
		{
			// net/http resends the GET once on a new connection, where the server
			// closes it again.
			name:     "refresh GET",
			run:      refreshRetryVolume,
			requests: []volumeRequest{refreshVolumeGet, refreshVolumeGet},
			diagnostic: func(u string) string {
				return `cannot read volume: Get "` + u + retryVolumePath + `?": EOF`
			},
			id: retryVolumeID,
		},
		{
			name:     "DELETE",
			run:      destroyVolume,
			requests: []volumeRequest{deleteVolumeRequest},
			diagnostic: func(u string) string {
				return `cannot delete volume: Delete "` + u + retryVolumePath + `?": EOF`
			},
			id: retryVolumeID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// A retry would get a 200 and fail the request assertion.
			api, server := newScriptedVolumeServer(t, eof, eof,
				volumeResponse{status: http.StatusOK, body: retryVolumeJSON})
			c := newVolumeClient(t, &config.Config{
				Host:                server.URL,
				RetryTimeoutSeconds: retryUntilContextEnds,
			})

			d, diags := tt.run(context.Background(), t, c)

			assertVolumeRequests(t, tt.requests, api)
			assert.Equal(t, tt.diagnostic(server.URL), diagnosticSummary(t, diags))
			assert.Equal(t, tt.id, d.Id())
		})
	}
}
