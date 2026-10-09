package common

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/internal/retrier"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Body captured from a real account SCIM DELETE /Groups/{id} that lost the race.
const scimConcurrentUpdateBody = `{
  "detail": "Could not handle request. Failed to delete in common handler. Failed due to conflicting concurrent updates. Please retry.",
  "schemas": ["urn:ietf:params:scim:api:messages:2.0:Error"],
  "status": "409"
}`

func newScimStubClient(t *testing.T, handler http.HandlerFunc) *DatabricksClient {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	cfg := &config.Config{
		Host:      server.URL,
		Token:     "dapi-test",
		AccountID: "test-account",
		HostMetadataResolver: func(ctx context.Context, host string) (*config.HostMetadata, error) {
			return nil, nil
		},
	}
	sdkClient, err := client.New(cfg)
	require.NoError(t, err)
	return &DatabricksClient{DatabricksClient: sdkClient}
}

func withFastScimBackoff(t *testing.T) {
	t.Helper()
	orig := scimConcurrentUpdateBackoff
	scimConcurrentUpdateBackoff = func() retrier.BackoffPolicy {
		return retrier.BackoffPolicy{Initial: time.Millisecond, Maximum: time.Millisecond}
	}
	t.Cleanup(func() { scimConcurrentUpdateBackoff = orig })
}

// stubResponses replies with the given status codes in order, then 200 forever.
func stubResponses(calls *int32, failures int32, status int, body string) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		n := atomic.AddInt32(calls, 1)
		if n <= failures {
			rw.WriteHeader(status)
			_, _ = rw.Write([]byte(body))
			return
		}
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write([]byte(`{}`))
	}
}

func TestScim_RetriesConcurrentUpdateConflict(t *testing.T) {
	withFastScimBackoff(t)
	for _, method := range []string{http.MethodPost, http.MethodPatch, http.MethodDelete} {
		t.Run(method, func(t *testing.T) {
			var calls int32
			dc := newScimStubClient(t, stubResponses(&calls, 2, http.StatusConflict, scimConcurrentUpdateBody))

			err := dc.Scim(context.Background(), method, "/preview/scim/v2/Groups/123", nil, nil, ApiLevelAccount)

			require.NoError(t, err)
			assert.Equal(t, int32(3), atomic.LoadInt32(&calls))
		})
	}
}

func TestScim_GivesUpAfterMaxConcurrentUpdateAttempts(t *testing.T) {
	withFastScimBackoff(t)
	var calls int32
	dc := newScimStubClient(t, stubResponses(&calls, 1000, http.StatusConflict, scimConcurrentUpdateBody))

	err := dc.Scim(context.Background(), http.MethodDelete, "/preview/scim/v2/Groups/123", nil, nil, ApiLevelAccount)

	require.ErrorContains(t, err, "conflicting concurrent updates")
	assert.Equal(t, int32(scimConcurrentUpdateMaxAttempts), atomic.LoadInt32(&calls))
}

func TestScim_DoesNotRetryOtherConflicts(t *testing.T) {
	withFastScimBackoff(t)
	var calls int32
	body := `{"detail": "Group with name admins already exists.", "schemas": ["urn:ietf:params:scim:api:messages:2.0:Error"], "status": "409"}`
	dc := newScimStubClient(t, stubResponses(&calls, 1000, http.StatusConflict, body))

	err := dc.Scim(context.Background(), http.MethodPost, "/preview/scim/v2/Groups", nil, nil, ApiLevelAccount)

	require.ErrorContains(t, err, "already exists")
	assert.Equal(t, int32(1), atomic.LoadInt32(&calls))
}
