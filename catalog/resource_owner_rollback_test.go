package catalog

import (
	"context"
	"fmt"
	"io"
	"maps"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"testing/synctest"
	"time"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/sharing"
)

// These tests pin that the owner rollback of the Unity Catalog and Delta Sharing
// resources other than databricks_volume reaches the server when Terraform cancels
// Update while the fields update backs off after a 429.
// TestVolumeFieldsPatchCancelledInBackoffRollsBackOwner pins the same for the volume.
// The Delta Sharing resources are here too so that one table lists every rollback on
// common.OwnerRollbackContext.

const (
	rollbackOldOwner = "owner-a@example.com"
	rollbackNewOwner = "owner-b@example.com"
)

// ownerRollbackRequest is one request the fake API received.
type ownerRollbackRequest struct {
	method string
	path   string
	body   string
}

// ownerRollbackAPI answers GETs with read, the owner update and the rollback with 200,
// and the fields update with a 429, and records each request. A write past those three
// gets a 500, which the SDK does not retry. Paths outside /api/, such as the host
// metadata the client reads when it is built, get a 404 and are not recorded. It
// serves in memory, which a synctest bubble needs.
type ownerRollbackAPI struct {
	t        *testing.T
	read     string
	mu       sync.Mutex
	received []ownerRollbackRequest
	writes   int
}

func (api *ownerRollbackAPI) RoundTrip(req *http.Request) (*http.Response, error) {
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
			ownerRollbackRequest{req.Method, req.URL.Path, string(body)})
		status, response := http.StatusOK, api.read
		if req.Method != http.MethodGet {
			api.writes++
			switch {
			case api.writes == 2:
				status, response = http.StatusTooManyRequests, tooManyRequestsJSON
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

func (api *ownerRollbackAPI) requests() []ownerRollbackRequest {
	api.mu.Lock()
	defer api.mu.Unlock()
	return append([]ownerRollbackRequest(nil), api.received...)
}

func (api *ownerRollbackAPI) writeCount() int {
	api.mu.Lock()
	defer api.mu.Unlock()
	return api.writes
}

// ownerRollbackResource is a resource whose Update changes the owner and other fields.
type ownerRollbackResource struct {
	// name is the resource name, which the update diagnostic carries.
	name     string
	resource func() common.Resource
	// account makes the client an account-level one.
	account bool
	id      string
	// state holds the attributes besides the ID and the owner.
	state map[string]string
	// change is the change besides the owner.
	change map[string]*terraform.ResourceAttrDiff
	// read is the body of the GET that Update sends first, or empty if it sends none.
	read string
	// endpoint is the method and path of the owner update, the fields update and the
	// rollback.
	endpoint string
	// ownerBody is the format of an owner update body, with %q for the owner.
	ownerBody string
}

// cancelInFieldsUpdateBackoff runs Update for the owner and fields change, cancels it in
// the fields update's backoff, and waits for it to return. It returns the time Update
// ran after cancellation, and must run in a synctest bubble.
func (rr ownerRollbackResource) cancelInFieldsUpdateBackoff(t *testing.T) (
	*ownerRollbackAPI, diag.Diagnostics, time.Duration,
) {
	t.Helper()
	api := &ownerRollbackAPI{t: t, read: rr.read}
	cfg := &config.Config{
		Host:                "https://workspace.test",
		RetryTimeoutSeconds: retryUntilContextEnds,
		HTTPTransport:       api,
	}
	if rr.account {
		cfg.Host, cfg.AccountID = "https://accounts.cloud.databricks.com", "100"
	}
	c := newVolumeClient(t, cfg)
	r := rr.resource().ToResource()
	state := map[string]string{"id": rr.id, "owner": rollbackOldOwner}
	maps.Copy(state, rr.state)
	change := map[string]*terraform.ResourceAttrDiff{
		"owner": {Old: rollbackOldOwner, New: rollbackNewOwner},
	}
	maps.Copy(change, rr.change)
	d, err := schema.InternalMap(r.Schema).Data(
		&terraform.InstanceState{ID: rr.id, Attributes: state},
		&terraform.InstanceDiff{Attributes: change})
	require.NoError(t, err)
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()
	var diags diag.Diagnostics
	done := make(chan struct{})
	go func() {
		defer close(done)
		diags = r.UpdateContext(context.WithValue(ctx, common.ResourceName, rr.name), d, c)
	}()

	// After the 429, Update blocks only in the fields update's backoff, so Wait returns
	// there.
	synctest.Wait()
	require.Equal(t, 2, api.writeCount(), "requests: %v", api.requests())
	cancelledAt := time.Now()
	cancel()
	<-done
	return api, diags, time.Since(cancelledAt)
}

// commentChange changes the comment from first to second.
func commentChange() map[string]*terraform.ResourceAttrDiff {
	return map[string]*terraform.ResourceAttrDiff{"comment": {Old: "first", New: "second"}}
}

func TestFieldsUpdateCancelledInBackoffRollsBackOwner(t *testing.T) {
	for _, rr := range []ownerRollbackResource{
		{
			name:      "catalog",
			resource:  ResourceCatalog,
			id:        "main",
			state:     map[string]string{"name": "main", "comment": "first"},
			change:    commentChange(),
			endpoint:  "PATCH /api/2.1/unity-catalog/catalogs/main",
			ownerBody: `{"owner": %q}`,
		},
		{
			name:     "external_location",
			resource: ResourceExternalLocation,
			id:       "landing",
			state: map[string]string{"name": "landing", "url": "s3://landing",
				"credential_name": "landing", "comment": "first"},
			change:    commentChange(),
			endpoint:  "PATCH /api/2.1/unity-catalog/external-locations/landing",
			ownerBody: `{"owner": %q}`,
		},
		{
			name:      "credential",
			resource:  ResourceCredential,
			id:        "service",
			state:     map[string]string{"name": "service", "comment": "first"},
			change:    commentChange(),
			endpoint:  "PATCH /api/2.1/unity-catalog/credentials/service",
			ownerBody: `{"owner": %q}`,
		},
		{
			name:     "metastore",
			resource: ResourceMetastore,
			id:       "abc",
			state: map[string]string{"name": "primary",
				"delta_sharing_organization_name": "first"},
			change: map[string]*terraform.ResourceAttrDiff{
				"delta_sharing_organization_name": {Old: "first", New: "second"},
			},
			endpoint:  "PATCH /api/2.1/unity-catalog/metastores/abc",
			ownerBody: `{"owner": %q}`,
		},
		{
			name:     "metastore",
			resource: ResourceMetastore,
			account:  true,
			id:       "abc",
			state: map[string]string{"name": "primary",
				"delta_sharing_organization_name": "first"},
			change: map[string]*terraform.ResourceAttrDiff{
				"delta_sharing_organization_name": {Old: "first", New: "second"},
			},
			endpoint:  "PUT /api/2.0/accounts/100/metastores/abc",
			ownerBody: `{"metastore_info": {"owner": %q}}`,
		},
		{
			name:     "registered_model",
			resource: ResourceRegisteredModel,
			id:       "main.default.model",
			state: map[string]string{"name": "model", "catalog_name": "main",
				"schema_name": "default", "comment": "first"},
			change:    commentChange(),
			endpoint:  "PATCH /api/2.1/unity-catalog/models/main.default.model",
			ownerBody: `{"owner": %q}`,
		},
		{
			name:     "share",
			resource: sharing.ResourceShare,
			id:       "abc",
			// Update sends only object changes, so the comment of the object changes.
			state: map[string]string{"name": "abc", "object.#": "1",
				"object.0.name": "main.default.t", "object.0.data_object_type": "TABLE",
				"object.0.comment": "first"},
			change: map[string]*terraform.ResourceAttrDiff{
				"object.0.comment": {Old: "first", New: "second"},
			},
			read: `{"name": "abc", "objects": [{"name": "main.default.t",
				"data_object_type": "TABLE", "comment": "first"}]}`,
			endpoint:  "PATCH /api/2.1/unity-catalog/shares/abc",
			ownerBody: `{"owner": %q}`,
		},
		{
			name:     "recipient",
			resource: sharing.ResourceRecipient,
			id:       "partner",
			state: map[string]string{"name": "partner", "authentication_type": "TOKEN",
				"comment": "first"},
			change:    commentChange(),
			endpoint:  "PATCH /api/2.1/unity-catalog/recipients/partner",
			ownerBody: `{"owner": %q}`,
		},
	} {
		test := rr.name
		if rr.account {
			test += " at account level"
		}
		t.Run(test, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				api, diags, elapsed := rr.cancelInFieldsUpdateBackoff(t)

				// The fake clock advances only on a timer, so the fields update's
				// backoff ended at cancellation and the rollback was not retried.
				assert.Zero(t, elapsed)
				requests := api.requests()
				if rr.read != "" {
					require.NotEmpty(t, requests)
					assert.Equal(t, http.MethodGet, requests[0].method)
					requests = requests[1:]
				}
				require.Len(t, requests, 3, "requests: %v", requests)
				for i, request := range requests {
					assert.Equal(t, rr.endpoint, request.method+" "+request.path,
						"write %d", i)
				}
				assert.JSONEq(t, fmt.Sprintf(rr.ownerBody, rollbackNewOwner),
					requests[0].body, "owner update")
				assert.JSONEq(t, fmt.Sprintf(rr.ownerBody, rollbackOldOwner),
					requests[2].body, "owner rollback")
				assert.Equal(t, "cannot update "+strings.ReplaceAll(rr.name, "_", " ")+
					": Too many requests from fixture", diagnosticSummary(t, diags))
			})
		})
	}
}
