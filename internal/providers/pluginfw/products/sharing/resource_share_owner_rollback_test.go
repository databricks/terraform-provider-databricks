package sharing

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"testing/synctest"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	rollbackShareName = "abc"
	rollbackSharePath = "/api/2.1/unity-catalog/shares/" + rollbackShareName
	rollbackOldOwner  = "owner-a@example.com"
	rollbackNewOwner  = "owner-b@example.com"
)

// shareRollbackRequest is one request the fake API received.
type shareRollbackRequest struct {
	method string
	path   string
	body   string
}

// shareRollbackAPI answers GETs with the share, the owner update and the rollback with
// 200, and the fields update with a 429, and records each request. A write past those
// three gets a 500, which the SDK does not retry. Paths outside /api/, such as the host
// metadata the client reads when it is built, get a 404 and are not recorded. It
// serves in memory, which a synctest bubble needs.
type shareRollbackAPI struct {
	t        *testing.T
	mu       sync.Mutex
	received []shareRollbackRequest
	writes   int
}

func (api *shareRollbackAPI) RoundTrip(req *http.Request) (*http.Response, error) {
	recorder := httptest.NewRecorder()
	if strings.HasPrefix(req.URL.Path, "/api/") {
		var body []byte
		if req.Body != nil {
			var err error
			if body, err = io.ReadAll(req.Body); err != nil {
				api.t.Errorf("cannot read the %s %s body: %v", req.Method, req.URL, err)
			}
		}
		api.mu.Lock()
		api.received = append(api.received,
			shareRollbackRequest{req.Method, req.URL.Path, string(body)})
		status := http.StatusOK
		response := fmt.Sprintf(`{"name": %q, "owner": %q, "comment": "first"}`,
			rollbackShareName, rollbackOldOwner)
		if req.Method != http.MethodGet {
			api.writes++
			switch {
			case api.writes == 2:
				status = http.StatusTooManyRequests
				response = `{"error_code": "TOO_MANY_REQUESTS",
					"message": "Too many requests from fixture"}`
			case api.writes > 3:
				status = http.StatusInternalServerError
				response = `{"error_code": "INTERNAL_ERROR", "message": "unexpected request"}`
			default:
				response = `{}`
			}
		}
		api.mu.Unlock()
		recorder.Header().Set("Content-Type", "application/json")
		recorder.WriteHeader(status)
		_, _ = recorder.WriteString(response)
	} else {
		http.NotFound(recorder, req)
	}
	resp := recorder.Result()
	resp.Request = req
	return resp, nil
}

func (api *shareRollbackAPI) requests() []shareRollbackRequest {
	api.mu.Lock()
	defer api.mu.Unlock()
	return append([]shareRollbackRequest(nil), api.received...)
}

func (api *shareRollbackAPI) writeCount() int {
	api.mu.Lock()
	defer api.mu.Unlock()
	return api.writes
}

// shareTfState converts share to the resource's model and sets it on a state with the
// resource's schema.
func shareTfState(
	t *testing.T, ctx context.Context, r *ShareResource, share sharing.ShareInfo,
) tfsdk.State {
	t.Helper()
	var schemaResp resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &schemaResp)
	require.False(t, schemaResp.Diagnostics.HasError(), "%v", schemaResp.Diagnostics)
	var model ShareInfoExtended
	require.False(t, converters.GoSdkToTfSdkStruct(ctx, share, &model).HasError())
	model.ID = types.StringValue(share.Name)
	model.ProviderConfig = types.ListNull(tfschema.ProviderConfig{}.Type(ctx))
	state := tfsdk.State{Schema: schemaResp.Schema}
	diags := state.Set(ctx, &model)
	require.False(t, diags.HasError(), "%v", diags)
	return state
}

// TestUpdateCancelledInFieldsUpdateBackoffRollsBackOwner pins that the owner rollback
// reaches the server when Terraform cancels Update while the fields update backs off
// after a 429.
func TestUpdateCancelledInFieldsUpdateBackoffRollsBackOwner(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		api := &shareRollbackAPI{t: t}
		inner, err := client.New(&config.Config{
			Host:  "https://workspace.test",
			Token: "fixture-token",
			// The provider's default retry timeout, which retries until the context ends.
			RetryTimeoutSeconds: -1,
			HTTPTransport:       api,
		})
		require.NoError(t, err)
		r := &ShareResource{Client: &common.DatabricksClient{DatabricksClient: inner}}

		ctx, cancel := context.WithCancel(t.Context())
		defer cancel()
		state := shareTfState(t, ctx, r, sharing.ShareInfo{
			Name: rollbackShareName, Owner: rollbackOldOwner, Comment: "first",
		})
		plan := shareTfState(t, ctx, r, sharing.ShareInfo{
			Name: rollbackShareName, Owner: rollbackNewOwner, Comment: "second",
		})
		req := resource.UpdateRequest{State: state, Plan: tfsdk.Plan(plan)}
		resp := resource.UpdateResponse{State: plan}
		done := make(chan struct{})
		go func() {
			defer close(done)
			r.Update(ctx, req, &resp)
		}()

		// After the 429, Update blocks only in the fields update's backoff, so Wait
		// returns there.
		synctest.Wait()
		require.Equal(t, 2, api.writeCount(), "requests: %v", api.requests())
		cancelledAt := time.Now()
		cancel()
		<-done

		// The fake clock advances only on a timer, so the fields update's backoff ended
		// at cancellation and the rollback was not retried.
		assert.Zero(t, time.Since(cancelledAt))
		requests := api.requests()
		require.Len(t, requests, 4, "requests: %v", requests)
		assert.Equal(t, http.MethodGet+" "+rollbackSharePath,
			requests[0].method+" "+requests[0].path, "current share")
		for i, request := range requests[1:] {
			assert.Equal(t, http.MethodPatch+" "+rollbackSharePath,
				request.method+" "+request.path, "write %d", i)
		}
		assert.JSONEq(t, fmt.Sprintf(`{"owner": %q}`, rollbackNewOwner), requests[1].body,
			"owner update")
		assert.JSONEq(t, `{"comment": "second"}`, requests[2].body, "fields update")
		assert.JSONEq(t, fmt.Sprintf(`{"owner": %q}`, rollbackOldOwner), requests[3].body,
			"owner rollback")
		assert.Equal(t, diag.Diagnostics{diag.NewErrorDiagnostic(
			"failed to update share", "Too many requests from fixture")},
			resp.Diagnostics)
	})
}
