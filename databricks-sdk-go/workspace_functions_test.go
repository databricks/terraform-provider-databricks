package databricks

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCurrentWorkspaceIDSendsWorkspaceIdHeaderWhenConfigHasWorkspaceID(t *testing.T) {
	// On unified (SPOG) hosts, requests to /api/2.0/preview/scim/v2/Me must
	// carry an X-Databricks-Workspace-Id header so the gateway can route them
	// to the correct workspace. When Config.WorkspaceID is set we forward it
	// on the request; the server echoes the ID back on the legacy
	// X-Databricks-Org-Id response header (the response side hasn't been
	// migrated yet).
	var meCalls int
	var gotOrgIdHeader string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/2.0/preview/scim/v2/Me" {
			meCalls++
			gotOrgIdHeader = r.Header.Get("X-Databricks-Workspace-Id")
			w.Header().Set("X-Databricks-Org-Id", "7474644166319138")
			w.Write([]byte(`{}`))
			return
		}
		http.NotFound(w, r)
	}))
	defer server.Close()

	w, err := NewWorkspaceClient(&Config{
		Host:        server.URL,
		Token:       "token",
		WorkspaceID: "7474644166319138",
	})
	require.NoError(t, err)

	got, err := w.CurrentWorkspaceID(t.Context())
	require.NoError(t, err)
	assert.Equal(t, int64(7474644166319138), got)
	assert.Equal(t, 1, meCalls)
	assert.Equal(t, "7474644166319138", gotOrgIdHeader)
}

func TestCurrentWorkspaceIDExcludesEntitlements(t *testing.T) {
	var gotRawQuery string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/2.0/preview/scim/v2/Me" {
			gotRawQuery = r.URL.RawQuery
			w.Header().Set("X-Databricks-Org-Id", "7474644166319138")
			w.Write([]byte(`{}`))
			return
		}
		http.NotFound(w, r)
	}))
	defer server.Close()

	w, err := NewWorkspaceClient(&Config{
		Host:  server.URL,
		Token: "token",
	})
	require.NoError(t, err)

	_, err = w.CurrentWorkspaceID(t.Context())
	require.NoError(t, err)
	assert.Equal(t, "excludedAttributes=entitlements", gotRawQuery)
}

func TestCurrentWorkspaceIDOmitsWorkspaceIdHeaderWhenConfigMissingWorkspaceID(t *testing.T) {
	// On legacy workspace hosts the host itself identifies the workspace, so
	// no routing header is needed. When Config.WorkspaceID is empty we send
	// the request without X-Databricks-Workspace-Id and read the ID from the
	// X-Databricks-Org-Id response header.
	var meCalls int
	var gotOrgIdHeader string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/2.0/preview/scim/v2/Me" {
			meCalls++
			gotOrgIdHeader = r.Header.Get("X-Databricks-Workspace-Id")
			w.Header().Set("X-Databricks-Org-Id", "7474644166319138")
			w.Write([]byte(`{}`))
			return
		}
		// Other bootstrap calls (e.g. /.well-known/databricks-config) are not
		// relevant to this test; respond with a benign 404.
		http.NotFound(w, r)
	}))
	defer server.Close()

	w, err := NewWorkspaceClient(&Config{
		Host:  server.URL,
		Token: "token",
	})
	require.NoError(t, err)

	got, err := w.CurrentWorkspaceID(t.Context())
	require.NoError(t, err)
	assert.Equal(t, int64(7474644166319138), got)
	assert.Equal(t, 1, meCalls)
	assert.Empty(t, gotOrgIdHeader)
}
