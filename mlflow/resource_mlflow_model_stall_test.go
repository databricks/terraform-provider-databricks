package mlflow

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/config"
	sdkgoutility "github.com/databricks/terraform-provider-databricks/sdk-go-utility"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mlflowModelStallServer is a workspace that reads every request and then stalls.
// With an empty body it sends nothing. Otherwise it sends the status line, the
// headers with the full Content-Length and the first half of body, and then nothing.
type mlflowModelStallServer struct {
	*httptest.Server
	body  string
	calls mlflowModelWireCalls

	mu sync.Mutex
	// requestBody is the body of the last request.
	requestBody string
	// release ends the stalled handlers, so that Close doesn't wait for them.
	release chan struct{}
}

func mlflowModelNewStallServer(t *testing.T, body string) *mlflowModelStallServer {
	s := &mlflowModelStallServer{body: body, release: make(chan struct{})}
	s.Server = httptest.NewServer(http.HandlerFunc(s.serve))
	t.Cleanup(func() {
		close(s.release)
		s.Close()
	})
	return s
}

func (s *mlflowModelStallServer) serve(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == mlflowModelWireMetadataPath {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w,
			`{"host_type": "workspace", "cloud": "AWS", "workspace_id": "12345"}`)
		return
	}
	// The server sees that the client hung up only once it has read the request.
	sent, _ := io.ReadAll(r.Body)
	s.mu.Lock()
	s.requestBody = string(sent)
	s.mu.Unlock()
	s.calls.add(r.Method + " " + r.URL.RequestURI())
	if s.body != "" {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", strconv.Itoa(len(s.body)))
		_, _ = io.WriteString(w, s.body[:len(s.body)/2])
		w.(http.Flusher).Flush()
	}
	select {
	case <-r.Context().Done():
	case <-s.release:
	}
}

func (s *mlflowModelStallServer) lastRequestBody() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.requestBody
}

// A request that the workspace accepts and never answers ends after
// http_timeout_seconds, here 1, its smallest value, with the inactivity error of the
// classic client. It is sent once: the timeout is not retried. A failed create leaves
// nothing for the state, although the workspace may have created the model, and a
// failed refresh keeps the model with its attributes.
//
// When the workspace sends the headers and then stalls in the body, the two paths
// report the same timeout with another text. The classic client ends the read of the
// body with the bare cancellation, and the AttemptTimeout helper of the modular path
// names the timeout. The responses are synthetic: inferred from the MLflow Model
// Registry API, not recorded.
func TestModelStalledRequestTimesOutWithoutRetry(t *testing.T) {
	mlflowModelOnBothPaths(t, testModelStalledRequestTimesOutWithoutRetry)
}

func testModelStalledRequestTimesOutWithoutRetry(t *testing.T) {
	const (
		name       = "stalled_model"
		createCall = "POST /api/2.0/mlflow/registered-models/create"
		readCall   = "GET " + mlflowModelWireGetPath + "?name=" + name
		timedOut   = "request timed out after 1s of inactivity"
		// The classic client reports a stall in the body without its timeout.
		classicBodyStall = "response body: context canceled"
	)
	tests := []struct {
		name   string
		create bool
		// body is the response that the workspace starts to send. Empty sends
		// nothing.
		body string
		// call is the one request, and requestBody its JSON body.
		call        string
		requestBody string
		// diagnostic is the exact error after the URL of the workspace is put in
		// for HOST, and sdkGoDiagnostic the one on the modular SDK path where it
		// differs.
		diagnostic      string
		sdkGoDiagnostic string
	}{
		{
			name:        "create_answered_with_nothing",
			create:      true,
			call:        createCall,
			requestBody: `{"name": "stalled_model", "description": "Pins a stalled request"}`,
			diagnostic: `cannot create mlflow model: Post "HOST/api/2.0/mlflow/` +
				`registered-models/create": ` + timedOut,
		},
		{
			name: "read_answered_with_nothing",
			call: readCall,
			diagnostic: `cannot read mlflow model: Get "HOST` + mlflowModelWireGetPath +
				`?name=stalled_model": ` + timedOut,
		},
		{
			name:   "create_answered_with_headers_and_half_a_body",
			create: true,
			body: `{"registered_model": {"name": "stalled_model",
				"description": "Pins a stalled request",
				"creation_timestamp": 1700000000000,
				"last_updated_timestamp": 1700000000000,
				"user_id": "tester@example.com"}}`,
			call:            createCall,
			requestBody:     `{"name": "stalled_model", "description": "Pins a stalled request"}`,
			diagnostic:      "cannot create mlflow model: " + classicBodyStall,
			sdkGoDiagnostic: "cannot create mlflow model: " + timedOut,
		},
		{
			name: "read_answered_with_headers_and_half_a_body",
			body: `{"registered_model_databricks": {"name": "stalled_model",
				"id": "0a1b2c3d4e5f60718293a4b5c6d7ea51",
				"description": "Pins a stalled request",
				"creation_timestamp": 1700000000000,
				"last_updated_timestamp": 1700000000000,
				"user_id": "tester@example.com", "permission_level": "CAN_MANAGE"}}`,
			call:            readCall,
			diagnostic:      "cannot read mlflow model: " + classicBodyStall,
			sdkGoDiagnostic: "cannot read mlflow model: " + timedOut,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := mlflowModelNewStallServer(t, tt.body)
			c := mlflowModelWireClient(t, &config.Config{
				Host:               server.URL,
				HTTPTimeoutSeconds: 1,
			})
			// The provider records that http_timeout_seconds is set in its block.
			c.SetHTTPTimeoutSetByUser(true)
			r := ResourceMlflowModel().ToResource()

			var d *schema.ResourceData
			var before map[string]string
			var diags diag.Diagnostics
			if tt.create {
				d = schema.TestResourceDataRaw(t, r.Schema, map[string]any{
					"name":        name,
					"description": "Pins a stalled request",
				})
				diags = r.CreateContext(mlflowModelWireContext(), d, c)
			} else {
				d = mlflowModelWireInState(t, r, name)
				before = d.State().Attributes
				diags = r.ReadContext(mlflowModelWireContext(), d, c)
			}

			want := tt.diagnostic
			if tt.sdkGoDiagnostic != "" && sdkgoutility.Enabled(mlflowModelResourceType) {
				want = tt.sdkGoDiagnostic
			}
			want = strings.ReplaceAll(want, "HOST", server.URL)
			assert.Equal(t, diag.Diagnostics{{Severity: diag.Error, Summary: want}}, diags)
			// The server records a request when it has read it, which a busy
			// machine can delay past the timeout.
			require.Eventually(t, func() bool { return len(server.calls.snapshot()) > 0 },
				10*time.Second, 10*time.Millisecond)
			// Exactly one request: the attempt that timed out is not retried.
			assert.Equal(t, []string{tt.call}, server.calls.snapshot())
			if tt.requestBody == "" {
				assert.Empty(t, server.lastRequestBody())
			} else {
				assert.JSONEq(t, tt.requestBody, server.lastRequestBody())
			}
			if tt.create {
				// Without an ID Terraform records no model.
				assert.Empty(t, d.Id())
				assert.Nil(t, d.State())
				return
			}
			assert.Equal(t, name, d.Id())
			require.NotNil(t, d.State())
			assert.Equal(t, before, d.State().Attributes)
		})
	}
}
