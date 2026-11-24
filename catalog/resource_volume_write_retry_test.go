package catalog

import (
	"context"
	"fmt"
	"io"
	"net"
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

// These tests pin the classic SDK client's HTTP timeout, cancellation and connection
// reset behavior on databricks_volume writes, which replay fixtures cannot express: a
// fixture can neither delay a response nor reset a connection, and a journey cannot
// cancel Terraform mid-call. The write_retry_ replay journeys pin transient-message
// retries and retry exhaustion. Each test has a WithSdkGo mirror that runs it with the
// volume calls through sdk-go uc/volumes/v1.

const (
	volumesAPIPath         = "/api/2.1/unity-catalog/volumes"
	oldVolumeOwner         = "owner-a@example.com"
	newVolumeOwner         = "owner-b@example.com"
	volumeWithNewOwnerJSON = `{"catalog_name": "main", "schema_name": "replay",
		"name": "retry_volume", "full_name": "main.replay.retry_volume",
		"volume_type": "MANAGED", "comment": "first", "owner": "owner-b@example.com"}`
	commentRejectedJSON = `{"error_code": "INVALID_PARAMETER_VALUE",
		"message": "Comment rejected by fixture"}`
)

var (
	createVolumePost = volumeRequest{http.MethodPost, volumesAPIPath,
		`{"catalog_name": "main", "schema_name": "replay", "name": "retry_volume",
		"volume_type": "MANAGED"}`}
	createOwnerPatch = volumeRequest{http.MethodPatch, retryVolumePath,
		`{"owner": "owner-a@example.com"}`}
	updateOwnerPatch = volumeRequest{http.MethodPatch, retryVolumePath,
		`{"owner": "owner-b@example.com"}`}
	updateFieldsPatch = volumeRequest{http.MethodPatch, retryVolumePath,
		`{"comment": "second"}`}
	ownerRollbackPatch = volumeRequest{http.MethodPatch, retryVolumePath,
		`{"owner": "owner-a@example.com"}`}
	deleteVolumeRequest = volumeRequest{http.MethodDelete, retryVolumePath + "?", ""}
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
	// hold withholds the response until the request is cancelled.
	hold bool
	// reset reads the request and then resets the connection, so it needs a real
	// server.
	reset bool
}

// scriptedVolumeAPI answers volume requests with its responses in order and records
// each request. A request past the script gets a 500, which the SDK does not retry.
// Other paths, such as the host metadata the client reads when it is built, get a 404
// and are not recorded.
type scriptedVolumeAPI struct {
	t         *testing.T
	released  chan struct{}
	mu        sync.Mutex
	responses []volumeResponse
	received  []volumeRequest
	// sdkGoReceived counts the received requests that sdk-go sent.
	sdkGoReceived int
}

func newScriptedVolumeAPI(t *testing.T, responses ...volumeResponse) *scriptedVolumeAPI {
	return &scriptedVolumeAPI{t: t, responses: responses, released: make(chan struct{})}
}

// newScriptedVolumeServer serves the API over HTTP, which the hold and reset responses
// need.
func newScriptedVolumeServer(
	t *testing.T, responses ...volumeResponse,
) (*scriptedVolumeAPI, *httptest.Server) {
	api := newScriptedVolumeAPI(t, responses...)
	server := httptest.NewServer(api)
	t.Cleanup(server.Close)
	// Cleanups run last in, first out, so held handlers return before server.Close
	// waits for them.
	t.Cleanup(func() { close(api.released) })
	return api, server
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
	if sentBySdkGoVolumes(r) {
		api.sdkGoReceived++
	}
	response := volumeResponse{
		status: http.StatusInternalServerError,
		body:   `{"error_code": "INTERNAL_ERROR", "message": "unexpected request"}`,
	}
	if len(api.responses) > 0 {
		response, api.responses = api.responses[0], api.responses[1:]
	}
	api.mu.Unlock()
	switch {
	case response.hold:
		select {
		case <-r.Context().Done():
		case <-api.released:
		}
	case response.reset:
		api.resetConnection(w)
	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.status)
		_, _ = w.Write([]byte(response.body))
	}
}

func (api *scriptedVolumeAPI) resetConnection(w http.ResponseWriter) {
	conn, _, err := http.NewResponseController(w).Hijack()
	if err != nil {
		api.t.Errorf("cannot hijack the connection: %v", err)
		return
	}
	// A zero linger makes Close send a TCP reset instead of a FIN.
	if err := conn.(*net.TCPConn).SetLinger(0); err != nil {
		api.t.Errorf("cannot set linger: %v", err)
	}
	_ = conn.Close()
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

func (api *scriptedVolumeAPI) sdkGoRequestCount() int {
	api.mu.Lock()
	defer api.mu.Unlock()
	return api.sdkGoReceived
}

// assertVolumeRequests compares methods and URIs exactly and bodies as JSON, and checks
// that sdk-go sent all requests if sdkGo is set, and none otherwise.
func assertVolumeRequests(
	t *testing.T, sdkGo bool, expected []volumeRequest, api *scriptedVolumeAPI,
) {
	t.Helper()
	actual := api.requests()
	assert.Equal(t, sdkGoCalls(sdkGo, len(actual)), api.sdkGoRequestCount(),
		"requests sent by sdk-go")
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

// createVolumeWithOwner runs the provider's wrapped Create for a volume owned by
// oldVolumeOwner, through sdk-go if sdkGo is set.
func createVolumeWithOwner(
	ctx context.Context, t *testing.T, c *common.DatabricksClient, sdkGo bool,
) (*schema.ResourceData, diag.Diagnostics) {
	r := volumeResource(sdkGo)
	d := schema.TestResourceDataRaw(t, r.Schema, map[string]any{
		"name":         "retry_volume",
		"catalog_name": "main",
		"schema_name":  "replay",
		"volume_type":  "MANAGED",
		"owner":        oldVolumeOwner,
	})
	d.MarkNewResource()
	return d, r.CreateContext(volumeContext(ctx), d, c)
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

// updateVolumeOwnerAndComment runs the provider's wrapped Update for that plan, through
// sdk-go if sdkGo is set.
func updateVolumeOwnerAndComment(
	ctx context.Context, t *testing.T, c *common.DatabricksClient, sdkGo bool,
) (*schema.ResourceData, diag.Diagnostics) {
	r := volumeResource(sdkGo)
	d := ownerAndCommentUpdateData(t, r)
	return d, r.UpdateContext(volumeContext(ctx), d, c)
}

// destroyVolume runs the provider's wrapped Delete, through sdk-go if sdkGo is set.
func destroyVolume(
	ctx context.Context, _ *testing.T, c *common.DatabricksClient, sdkGo bool,
) (*schema.ResourceData, diag.Diagnostics) {
	r := volumeResource(sdkGo)
	d := r.TestResourceData()
	d.SetId(retryVolumeID)
	return d, r.DeleteContext(volumeContext(ctx), d, c)
}

func TestVolumeWritesPastHTTPTimeoutAreNotRetried(t *testing.T) {
	testVolumeWritesPastHTTPTimeoutAreNotRetried(t, false)
}

func TestVolumeWritesPastHTTPTimeoutAreNotRetriedWithSdkGo(t *testing.T) {
	testVolumeWritesPastHTTPTimeoutAreNotRetried(t, true)
}

func testVolumeWritesPastHTTPTimeoutAreNotRetried(t *testing.T, sdkGo bool) {
	held := volumeResponse{hold: true}
	created := volumeResponse{status: http.StatusOK, body: retryVolumeJSON}
	ownerChanged := volumeResponse{status: http.StatusOK, body: volumeWithNewOwnerJSON}
	rejected := volumeResponse{status: http.StatusBadRequest, body: commentRejectedJSON}
	timedOut := func(method, url string) string {
		return fmt.Sprintf(`%s "%s": request timed out after 1s of inactivity`,
			method, url)
	}
	tests := []struct {
		name      string
		responses []volumeResponse
		run       func(context.Context, *testing.T, *common.DatabricksClient, bool) (
			*schema.ResourceData, diag.Diagnostics)
		requests []volumeRequest
		// diagnostic returns the expected diagnostic for the server's URL.
		diagnostic func(serverURL string) string
		id         string
	}{
		{
			name:      "create POST",
			responses: []volumeResponse{held},
			run:       createVolumeWithOwner,
			requests:  []volumeRequest{createVolumePost},
			diagnostic: func(u string) string {
				return "cannot create volume: " + timedOut("Post", u+volumesAPIPath)
			},
			id: "",
		},
		{
			name:      "owner PATCH after create",
			responses: []volumeResponse{created, held},
			run:       createVolumeWithOwner,
			requests:  []volumeRequest{createVolumePost, createOwnerPatch},
			diagnostic: func(u string) string {
				return "cannot create volume: " + timedOut("Patch", u+retryVolumePath)
			},
			id: retryVolumeID,
		},
		{
			name:      "owner PATCH in Update",
			responses: []volumeResponse{held},
			run:       updateVolumeOwnerAndComment,
			requests:  []volumeRequest{updateOwnerPatch},
			diagnostic: func(u string) string {
				return "cannot update volume: " + timedOut("Patch", u+retryVolumePath)
			},
			id: retryVolumeID,
		},
		{
			// The rollback runs on Update's context, which the timeout did not end.
			name:      "fields PATCH after an owner change",
			responses: []volumeResponse{ownerChanged, held, created},
			run:       updateVolumeOwnerAndComment,
			requests: []volumeRequest{
				updateOwnerPatch, updateFieldsPatch, ownerRollbackPatch,
			},
			diagnostic: func(u string) string {
				return "cannot update volume: " + timedOut("Patch", u+retryVolumePath)
			},
			id: retryVolumeID,
		},
		{
			name:      "rollback PATCH",
			responses: []volumeResponse{ownerChanged, rejected, held},
			run:       updateVolumeOwnerAndComment,
			requests: []volumeRequest{
				updateOwnerPatch, updateFieldsPatch, ownerRollbackPatch,
			},
			diagnostic: func(u string) string {
				return ownerRollbackDiagnostic("Comment rejected by fixture",
					timedOut("Patch", u+retryVolumePath))
			},
			id: retryVolumeID,
		},
		{
			name:      "DELETE",
			responses: []volumeResponse{held},
			run:       destroyVolume,
			requests:  []volumeRequest{deleteVolumeRequest},
			diagnostic: func(u string) string {
				return "cannot delete volume: " +
					timedOut("Delete", u+retryVolumePath+"?")
			},
			id: retryVolumeID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// A retry of the held request would get a 500 and fail the request
			// assertion.
			api, server := newScriptedVolumeServer(t, tt.responses...)
			c := newVolumeClient(t, &config.Config{
				Host:                server.URL,
				RetryTimeoutSeconds: retryUntilContextEnds,
				HTTPTimeoutSeconds:  1,
			})

			d, diags := tt.run(context.Background(), t, c, sdkGo)

			assertVolumeRequests(t, sdkGo, tt.requests, api)
			assert.Equal(t, tt.diagnostic(server.URL), diagnosticSummary(t, diags))
			assert.Equal(t, tt.id, d.Id())
		})
	}
}

func TestVolumeFieldsPatchCancelledInBackoffRollsBackOnCancelledContext(t *testing.T) {
	testVolumeFieldsPatchCancelledInBackoffRollsBackOnCancelledContext(t, false)
}

func TestVolumeFieldsPatchCancelledInBackoffRollsBackOnCancelledContextWithSdkGo(
	t *testing.T) {
	testVolumeFieldsPatchCancelledInBackoffRollsBackOnCancelledContext(t, true)
}

func testVolumeFieldsPatchCancelledInBackoffRollsBackOnCancelledContext(t *testing.T,
	sdkGo bool) {
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
		r := volumeResource(sdkGo)
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
		assertVolumeRequests(t, sdkGo,
			[]volumeRequest{updateOwnerPatch, updateFieldsPatch}, api)
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
		assertVolumeRequests(t, sdkGo,
			[]volumeRequest{updateOwnerPatch, updateFieldsPatch}, api)
		assert.Equal(t, ownerRollbackDiagnostic("Too many requests from fixture",
			"failed in rate limiter: context canceled"), diagnosticSummary(t, diags))
		assert.Equal(t, retryVolumeID, d.Id(), "the volume must stay in state")
	})
}

func TestVolumeCreateRetriedAfterConnectionResetFailsWithAlreadyExists(t *testing.T) {
	testVolumeCreateRetriedAfterConnectionResetFailsWithAlreadyExists(t, false)
}

func TestVolumeCreateRetriedAfterConnectionResetFailsWithAlreadyExistsWithSdkGo(
	t *testing.T) {
	testVolumeCreateRetriedAfterConnectionResetFailsWithAlreadyExists(t, true)
}

func testVolumeCreateRetriedAfterConnectionResetFailsWithAlreadyExists(t *testing.T,
	sdkGo bool) {
	// The first POST reaches the server, which creates the volume and then resets the
	// connection before responding. The SDK retries the reset with the same body, and
	// the server rejects the retry because the volume now exists.
	api, server := newScriptedVolumeServer(t,
		volumeResponse{reset: true},
		volumeResponse{status: http.StatusConflict, body: `{
			"error_code": "RESOURCE_ALREADY_EXISTS",
			"message": "Volume 'main.replay.retry_volume' already exists"}`},
	)
	c := newVolumeClient(t, &config.Config{
		Host:                server.URL,
		RetryTimeoutSeconds: retryUntilContextEnds,
	})

	d, diags := createVolumeWithOwner(context.Background(), t, c, sdkGo)

	assertVolumeRequests(t, sdkGo, []volumeRequest{createVolumePost, createVolumePost},
		api)
	assert.Equal(t,
		"cannot create volume: Volume 'main.replay.retry_volume' already exists",
		diagnosticSummary(t, diags))
	assert.Empty(t, d.Id(), "the volume the first POST created must not be in state")
}
