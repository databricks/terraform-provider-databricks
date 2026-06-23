package sdkv2

import (
	"context"
	"maps"
	"net/http"
	"net/http/httptest"
	"slices"
	"sync"
	"testing"
	"time"

	databricks "github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// workspaceIDHeader is the routing header injected by common.AddWorkspaceIdHeader.
// Account-scoped requests (/api/<ver>/accounts/...) must never carry it: the
// platform routes such a request to a workspace and the account call fails or
// mis-routes. See common/client.go.
const workspaceIDHeader = "X-Databricks-Workspace-Id"

// capturedAccountRequest records one outbound request so the test can assert on
// its headers.
type capturedAccountRequest struct {
	Method string
	Path   string
	Header http.Header
}

// newAccountHeaderCaptureClient returns a common.DatabricksClient configured as
// an account client (AccountID set) pointed at a header-capturing test server.
//
// Config.WorkspaceID is deliberately set to a non-empty value. That is the only
// condition under which common.AddWorkspaceIdHeader emits the routing header, so
// the test proves account resources never emit it *even when a workspace id is
// present in the provider config* — the worst case for a leak.
func newAccountHeaderCaptureClient(t *testing.T) (*common.DatabricksClient, *[]capturedAccountRequest) {
	t.Helper()
	var mu sync.Mutex
	var captured []capturedAccountRequest
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		mu.Lock()
		captured = append(captured, capturedAccountRequest{
			Method: req.Method,
			Path:   req.URL.Path,
			Header: req.Header.Clone(),
		})
		mu.Unlock()
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write([]byte(`{}`))
	}))
	t.Cleanup(server.Close)

	cfg := &config.Config{
		Host:        server.URL,
		Token:       "dapi-test",
		AccountID:   "test-account-id",
		WorkspaceID: "99999",
		HostMetadataResolver: func(ctx context.Context, host string) (*config.HostMetadata, error) {
			return nil, nil
		},
	}
	sdkClient, err := client.New(cfg)
	require.NoError(t, err)
	dc := &common.DatabricksClient{DatabricksClient: sdkClient}

	acc, err := databricks.NewAccountClient((*databricks.Config)(cfg))
	require.NoError(t, err)
	dc.SetAccountClient(acc)
	return dc, &captured
}

// TestAccountResources_NeverSendWorkspaceIdHeader is a guardrail over the whole
// AccountResources group. This branch added per-callsite injection of the
// X-Databricks-Workspace-Id routing header (common.AddWorkspaceIdHeader) to
// workspace-level resources. Account resources address
// /api/<ver>/accounts/<account_id>/... and must never carry that header.
//
// The test drives every account resource's CRUD against a capturing server with
// a workspace id present in config, and asserts no request ever carries the
// header. It catches a future edit that wires AddWorkspaceIdHeader onto a
// callsite an account resource reaches (and exercises common/client.go's
// /accounts/ suppression guard as defense in depth).
//
// We don't care whether each operation succeeds — only what goes on the wire.
// The common.Resource wrapper recovers panics into diagnostics, which we ignore;
// a short timeout bounds any resource that polls on the empty ({}) responses.
func TestAccountResources_NeverSendWorkspaceIdHeader(t *testing.T) {
	total := 0
	for _, name := range slices.Sorted(maps.Keys(AccountResources)) {
		res := AccountResources[name]
		dc, captured := newAccountHeaderCaptureClient(t)
		require.NotEmpty(t, dc.Config.WorkspaceID,
			"test must run with a workspace id in config to be meaningful")

		exercise := func(fn func(context.Context, *schema.ResourceData, any) diag.Diagnostics) {
			if fn == nil {
				return
			}
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			data := res.TestResourceData()
			// Composite id covers mws account resources that unpack
			// "<account_id>/<resource_id>"; harmless for the rest.
			data.SetId("test-account-id/00000000-0000-0000-0000-000000000000")
			_ = fn(ctx, data, dc)
		}
		exercise(res.CreateContext)
		exercise(res.ReadContext)
		exercise(res.UpdateContext)
		exercise(res.DeleteContext)

		for _, r := range *captured {
			assert.Empty(t, r.Header.Get(workspaceIDHeader),
				"account resource %q issued %s %s carrying a %s header; account requests must never be workspace-tagged",
				name, r.Method, r.Path, workspaceIDHeader)
		}
		total += len(*captured)
	}

	// Sanity: the harness actually put requests on the wire. A zero here would
	// make every assertion above vacuous, so fail loudly instead.
	require.Greater(t, total, 0,
		"expected the account resources to issue at least one captured request")
}

// TestAccountHeaderCaptureClient_DetectsHeader is the positive control for the
// guardrail above: it proves the capture client actually surfaces the header
// when one is sent, so a green TestAccountResources_NeverSendWorkspaceIdHeader
// means "no leak" rather than "harness blind". A raw workspace-path request with
// the visitor must carry the header; the same request on an /accounts/ path must
// not (common/client.go's suppression guard).
func TestAccountHeaderCaptureClient_DetectsHeader(t *testing.T) {
	dc, captured := newAccountHeaderCaptureClient(t)
	ctx := context.Background()

	require.NoError(t, dc.Post(ctx, "/clusters/list", map[string]string{}, nil, dc.AddWorkspaceIdHeader))
	require.NoError(t, dc.Post(ctx, "/accounts/test-account-id/networks", map[string]string{}, nil, dc.AddWorkspaceIdHeader))

	require.Len(t, *captured, 2)
	assert.Equal(t, "99999", (*captured)[0].Header.Get(workspaceIDHeader),
		"workspace-path request with the visitor must carry the header (harness can see it)")
	assert.Empty(t, (*captured)[1].Header.Get(workspaceIDHeader),
		"account-path request must be suppressed by the guard even with the visitor")
}
