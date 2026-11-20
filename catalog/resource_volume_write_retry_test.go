package catalog

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"testing/synctest"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
)

// The write_retry_ replay journeys pin transient-message retries, retry exhaustion, the
// HTTP inactivity timeout and connection resets on databricks_volume writes.

const (
	volumesAPIPath         = "/api/2.1/unity-catalog/volumes"
	oldVolumeOwner         = "owner-a@example.com"
	newVolumeOwner         = "owner-b@example.com"
	volumeWithNewOwnerJSON = `{"catalog_name": "main", "schema_name": "replay",
		"name": "retry_volume", "full_name": "main.replay.retry_volume",
		"volume_type": "MANAGED", "comment": "first", "owner": "owner-b@example.com"}`
)

var (
	updateOwnerPatch = volumeRequest{http.MethodPatch, retryVolumePath,
		`{"owner": "owner-b@example.com"}`}
	updateFieldsPatch = volumeRequest{http.MethodPatch, retryVolumePath,
		`{"comment": "second"}`}
)

// volumeRequest is one request the fake volume API received. An empty body means none.
type volumeRequest struct {
	method string
	uri    string
	body   string
}

// volumeResponse is how the fake volume API answers one request.
type volumeResponse struct {
	status int
	body   string
}

// scriptedVolumeAPI answers volume requests with its responses in order and records
// each request. A request past the script gets a 500, which the SDK does not retry.
// Other paths, such as the host metadata the client reads when it is built, get a 404
// and are not recorded.
type scriptedVolumeAPI struct {
	t         *testing.T
	mu        sync.Mutex
	responses []volumeResponse
	received  []volumeRequest
}

func newScriptedVolumeAPI(t *testing.T, responses ...volumeResponse) *scriptedVolumeAPI {
	return &scriptedVolumeAPI{t: t, responses: responses}
}

func (api *scriptedVolumeAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, volumesAPIPath) {
		http.NotFound(w, r)
		return
	}
	var body []byte
	if r.Body != nil {
		var err error
		if body, err = io.ReadAll(r.Body); err != nil {
			api.t.Errorf("cannot read the %s %s body: %v", r.Method, r.URL, err)
		}
	}
	api.mu.Lock()
	request := volumeRequest{r.Method, r.URL.RequestURI(), string(body)}
	api.received = append(api.received, request)
	response := volumeResponse{
		status: http.StatusInternalServerError,
		body:   `{"error_code": "INTERNAL_ERROR", "message": "unexpected request"}`,
	}
	if len(api.responses) > 0 {
		response, api.responses = api.responses[0], api.responses[1:]
	}
	api.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.status)
	_, _ = w.Write([]byte(response.body))
}

// RoundTrip serves the API in memory, which a synctest bubble needs.
func (api *scriptedVolumeAPI) RoundTrip(req *http.Request) (*http.Response, error) {
	recorder := httptest.NewRecorder()
	api.ServeHTTP(recorder, req)
	resp := recorder.Result()
	resp.Request = req
	return resp, nil
}

func (api *scriptedVolumeAPI) requests() []volumeRequest {
	api.mu.Lock()
	defer api.mu.Unlock()
	return append([]volumeRequest(nil), api.received...)
}

// assertVolumeRequests compares methods and URIs exactly and bodies as JSON.
func assertVolumeRequests(
	t *testing.T, expected []volumeRequest, api *scriptedVolumeAPI,
) {
	t.Helper()
	actual := api.requests()
	require.Len(t, actual, len(expected), "requests: %v", actual)
	for i, want := range expected {
		got := actual[i]
		assert.Equal(t, want.method+" "+want.uri, got.method+" "+got.uri,
			"request %d", i)
		if want.body == "" {
			assert.Empty(t, got.body, "request %d", i)
		} else {
			assert.JSONEq(t, want.body, got.body, "request %d", i)
		}
	}
}

// ownerRollbackDiagnostic is the diagnostic for a failed owner rollback after the
// fields PATCH failed with updateErr.
func ownerRollbackDiagnostic(updateErr, rollbackErr string) string {
	return "cannot update volume: " + updateErr + ". Owner rollback also failed: " +
		rollbackErr + ". Owner of this resource was updated but other fields " +
		"couldn't be updated and owner couldn't be rolled back. \n As a result, the " +
		"owner of this resource is updated to " + newVolumeOwner + " but other " +
		"attributes aren't. To revert the owner change, please manually change the " +
		"owner to " + oldVolumeOwner + ". \n\n You can also use the databricks cli " +
		"(https://docs.databricks.com/en/dev-tools/cli/install.html) to update the " +
		"owner."
}

func volumeContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, common.ResourceName, "volume")
}

// ownerAndCommentUpdateData is the data Terraform passes to Update for a plan that
// changes the owner from oldVolumeOwner to newVolumeOwner and the comment from first
// to second.
func ownerAndCommentUpdateData(t *testing.T, r *schema.Resource) *schema.ResourceData {
	t.Helper()
	state := &terraform.InstanceState{ID: retryVolumeID, Attributes: map[string]string{
		"id":           retryVolumeID,
		"name":         "retry_volume",
		"catalog_name": "main",
		"schema_name":  "replay",
		"volume_type":  "MANAGED",
		"comment":      "first",
		"owner":        oldVolumeOwner,
	}}
	diff := &terraform.InstanceDiff{Attributes: map[string]*terraform.ResourceAttrDiff{
		"comment": {Old: "first", New: "second"},
		"owner":   {Old: oldVolumeOwner, New: newVolumeOwner},
	}}
	d, err := schema.InternalMap(r.Schema).Data(state, diff)
	require.NoError(t, err)
	return d
}

// A journey cannot cancel Terraform during a backoff, so this case is a Go test.
func TestVolumeFieldsPatchCancelledInBackoffRollsBackOnCancelledContext(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// The fields PATCH retry and the rollback would both succeed if they were sent.
		api := newScriptedVolumeAPI(t,
			volumeResponse{status: http.StatusOK, body: volumeWithNewOwnerJSON},
			volumeResponse{status: http.StatusTooManyRequests, body: tooManyRequestsJSON},
			volumeResponse{status: http.StatusOK, body: retryVolumeJSON},
			volumeResponse{status: http.StatusOK, body: retryVolumeJSON},
		)
		c := newVolumeClient(t, &config.Config{
			Host:                "https://volume-api.test",
			RetryTimeoutSeconds: retryUntilContextEnds,
			HTTPTransport:       api,
		})
		r := ResourceVolume().ToResource()
		d := ownerAndCommentUpdateData(t, r)
		ctx, cancel := context.WithCancel(t.Context())
		defer cancel()
		var diags diag.Diagnostics
		done := make(chan struct{})
		go func() {
			defer close(done)
			diags = r.UpdateContext(volumeContext(ctx), d, c)
		}()

		// After the 429, Update blocks only in the fields PATCH's backoff, so Wait
		// returns there.
		synctest.Wait()
		assertVolumeRequests(t, []volumeRequest{updateOwnerPatch, updateFieldsPatch}, api)
		cancel()
		// Wait does not advance the fake clock, so Update must return without the
		// backoff timer firing.
		synctest.Wait()
		select {
		case <-done:
		default:
			t.Fatal("Update kept waiting out its backoff after cancellation")
		}

		// The rollback fails in the SDK's rate limiter, so it never reaches the server
		// and the owner change stays.
		assertVolumeRequests(t, []volumeRequest{updateOwnerPatch, updateFieldsPatch}, api)
		assert.Equal(t, ownerRollbackDiagnostic("Too many requests from fixture",
			"failed in rate limiter: context canceled"), diagnosticSummary(t, diags))
		assert.Equal(t, retryVolumeID, d.Id(), "the volume must stay in state")
	})
}
